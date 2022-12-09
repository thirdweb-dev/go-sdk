package thirdweb

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"mime/multipart"
	"net/http"
	"reflect"
	"strings"
)

type baseUriWithUris struct {
	baseUri string
	uris    []string
}
type storage interface {
	Get(ctx context.Context, uri string) ([]byte, error)
	Upload(ctx context.Context, data map[string]interface{}, contractAddress string, signerAddress string) (string, error)
	UploadBatch(ctx context.Context, data []map[string]interface{}, fileStartNumber int, contractAddress string, signerAddress string) (*baseUriWithUris, error)
}

type uploadResponse struct {
	IpfsHash    string   `json:"ipfsHash"`
	PinSize     *big.Int `json:"PinSize"`
	Timestamp   string   `json:"Timestamp"`
	IsDuplicate bool     `json:"isDuplicate"`
	IpfsUri     string   `json:"IpfsUri"`
}

type IpfsStorage struct {
	gatewayUrl string
	httpClient *http.Client
}

func newIpfsStorage(gatewayUrl string, httpClient *http.Client) *IpfsStorage {
	return &IpfsStorage{
		gatewayUrl: gatewayUrl,
		httpClient: httpClient,
	}
}

// Get
//
// # Get IPFS data at a given hash and return it as byte data
//
// uri: the IPFS URI to fetch data from
//
// returns: byte data of the IPFS data at the URI
func (ipfs *IpfsStorage) Get(ctx context.Context, uri string) ([]byte, error) {
	gatewayUrl := replaceHashWithGatewayUrl(uri, ipfs.gatewayUrl)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, gatewayUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := ipfs.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Bad status code, %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Upload
//
// Upload method can be used to upload a generic payload to IPFS.
//
// data: the individual data to upload to IPFS
//
// contractAddress: the optional contractAddress upload is being called from
//
// signerAddress: the optional signerAddress upload is being called from
//
// returns: the URI of the IPFS upload
func (ipfs *IpfsStorage) Upload(ctx context.Context, data map[string]interface{}, contractAddress string, signerAddress string) (string, error) {
	baseUriWithUris, err := ipfs.UploadBatch(ctx, []map[string]interface{}{data}, 0, contractAddress, signerAddress)
	if err != nil {
		return "", err
	}

	baseUri := baseUriWithUris.baseUri + "0"
	return baseUri, nil
}

// UploadBatch
//
// UploadBatch method can be used to upload a batch of generic payloads to IPFS.
//
// data: the array of data to upload to IPFS
//
// contractAddress: the optional contractAddress upload is being called from
//
// signerAddress: the optional signerAddress upload is being called from
//
// returns: the base URI of the IPFS upload folder with the URIs of each subfile
func (ipfs *IpfsStorage) UploadBatch(ctx context.Context, data []map[string]interface{}, fileStartNumber int, contractAddress string, signerAddress string) (*baseUriWithUris, error) {
	preparedData, err := ipfs.batchUploadProperties(ctx, data)
	if err != nil {
		return nil, err
	}

	dataToUpload := []interface{}{}
	dataValue := reflect.ValueOf(preparedData)
	switch dataValue.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < dataValue.Len(); i++ {
			jsonData, err := json.Marshal(dataValue.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			dataToUpload = append(dataToUpload, jsonData)
		}
		break
	default:
		return nil, errors.New("data must be an array or slice")
	}

	baseUriWithUris, err := ipfs.uploadBatchWithCid(ctx, dataToUpload, fileStartNumber, contractAddress, signerAddress)
	if err != nil {
		return nil, err
	}

	return baseUriWithUris, nil
}

func (ipfs *IpfsStorage) getUploadToken(ctx context.Context, contractAddress string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("%v/grant", twIpfsServerUrl), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("X-App-Name", fmt.Sprintf("CONSOLE-GO-SDK-%v", contractAddress))
	result, err := ipfs.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	if result.StatusCode != http.StatusOK {
		return "", &failedToUploadError{
			statusCode: result.StatusCode,
		}
	}

	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return "", err
	}
	text := string(body)

	return text, nil
}

func (ipfs *IpfsStorage) uploadBatchWithCid(
	ctx context.Context,
	// data (string | io.Reader)[] - file or JSON string
	data []interface{},
	fileStartNumber int,
	contractAddress string,
	signerAddress string,
) (*baseUriWithUris, error) {
	uploadToken, err := ipfs.getUploadToken(ctx, contractAddress)
	if err != nil {
		return nil, err
	}

	fileNames := []string{}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for i, obj := range data {
		if jsonData, ok := obj.([]byte); ok {
			fileName := fmt.Sprintf("%v", i+fileStartNumber)
			fileNames = append(fileNames, fileName)

			part, err := writer.CreateFormFile("file", fmt.Sprintf("files/%v", fileName))
			if err != nil {
				return nil, err
			}

			if _, err := part.Write(jsonData); err != nil {
				return nil, err
			}
		} else if fileData, ok := obj.(io.Reader); ok {
			fileName := fmt.Sprintf("%v", i+fileStartNumber)
			fileNames = append(fileNames, fileName)

			part, err := writer.CreateFormFile("file", fmt.Sprintf("files/%v", fileName))
			if err != nil {
				return nil, err
			}

			_, err = io.Copy(part, fileData)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("Data to upload must be either JSON ([]byte) or a file (io.Reader)")
		}
	}

	_ = writer.Close()

	req, err := http.NewRequestWithContext(ctx, "POST", pinataIpfsUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", uploadToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if result, err := ipfs.httpClient.Do(req); err != nil {
		return nil, err
	} else {
		if result.StatusCode != http.StatusOK {
			return nil, &failedToUploadError{
				statusCode: result.StatusCode,
				Payload:    data,
			}
		}

		var uploadMeta uploadResponse
		bodyBytes, err := ioutil.ReadAll(result.Body)
		if err != nil {
			return nil, &failedToUploadError{
				statusCode:      result.StatusCode,
				Payload:         data,
				UnderlyingError: err,
			}
		}

		if err := json.Unmarshal(bodyBytes, &uploadMeta); err != nil {
			return nil, &unmarshalError{
				body:            string(bodyBytes),
				typeName:        "UploadResponse",
				UnderlyingError: err,
			}
		}

		baseUri := "ipfs://" + uploadMeta.IpfsHash + "/"

		uris := []string{}
		for _, fileName := range fileNames {
			uri := baseUri + fileName
			uris = append(uris, uri)
		}

		return &baseUriWithUris{
			baseUri: baseUri,
			uris:    uris,
		}, nil
	}
}

// returns - map[string]interface{}
func (ipfs *IpfsStorage) batchUploadProperties(ctx context.Context, data []map[string]interface{}) (interface{}, error) {
	sanitizedMetadatas, err := replaceGatewayUrlWithHash(data, "ipfs://", ipfs.gatewayUrl)
	if err != nil {
		return nil, err
	}

	filesToUpload, err := buildFilePropertiesMap(sanitizedMetadatas, []interface{}{})
	if err != nil {
		return nil, err
	}

	if len(filesToUpload) == 0 {
		return sanitizedMetadatas, nil
	}

	baseUriWithUris, err := ipfs.uploadBatchWithCid(ctx, filesToUpload, 0, "", "")
	if err != nil {
		return nil, err
	}

	replacedMetadatas, err := replaceFilePropertiesWithHashes(sanitizedMetadatas, baseUriWithUris.uris)
	if err != nil {
		return nil, err
	}

	return replacedMetadatas, nil
}

// data - array or map or strings
// Returns []io.Reader files to upload
func buildFilePropertiesMap(data interface{}, files []interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			builtFiles, err := buildFilePropertiesMap(v.Index(i).Interface(), files)
			if err != nil {
				return nil, err
			}

			files = builtFiles
		}
		break
	case reflect.Map:
		for _, k := range v.MapKeys() {
			builtFiles, err := buildFilePropertiesMap(v.MapIndex(k).Interface(), files)
			if err != nil {
				return nil, err
			}

			files = builtFiles
		}
		break
	default:
		file, ok := data.(io.Reader)
		if ok {
			files = append(files, file)
		}
	}

	return files, nil
}

func replaceFilePropertiesWithHashes(data interface{}, cids []string) (interface{}, error) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		updated := []interface{}{}
		for i := 0; i < v.Len(); i++ {
			val, err := replaceFilePropertiesWithHashes(v.Index(i).Interface(), cids)
			if err != nil {
				return nil, err
			}

			updated = append(updated, val)
		}

		return updated, nil
	case reflect.Map:
		updated := map[string]interface{}{}
		for _, k := range v.MapKeys() {
			val, err := replaceFilePropertiesWithHashes(v.MapIndex(k).Interface(), cids)
			if err != nil {
				return nil, err
			}

			updated[k.String()] = val
		}

		return updated, nil
	default:
		_, ok := data.(io.Reader)
		if ok {
			data, cids = cids[0], cids[1:]
		}

		return data, nil
	}
}

func replaceGatewayUrlWithHash(data interface{}, scheme string, gatewayUrl string) (interface{}, error) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if _, err := replaceGatewayUrlWithHash(v.Index(i), scheme, gatewayUrl); err != nil {
				return nil, err
			}
		}
		break
	case reflect.Map:
		for _, k := range v.MapKeys() {
			if _, err := replaceGatewayUrlWithHash(v.MapIndex(k), scheme, gatewayUrl); err != nil {
				return nil, err
			}
		}
		break
	case reflect.String:
		if strings.Contains(v.String(), gatewayUrl) {
			data = strings.Replace(v.String(), gatewayUrl, scheme, 1)
		}
	}

	return data, nil
}

func resolveGatewayUrl(data interface{}, scheme string, gatewayUrl string) interface{} {
	return data
}

func replaceHashWithGatewayUrl(ipfsUrl string, gatewayUrl string) string {
	if ipfsUrl == "" {
		return ""
	}

	gateway := gatewayUrl
	if !strings.HasSuffix(gatewayUrl, "/") {
		gateway = gatewayUrl + "/"
	}

	return strings.Replace(ipfsUrl, "ipfs://", gateway, 1)
}
