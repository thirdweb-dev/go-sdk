package thirdweb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	Get(uri string) ([]byte, error)
	Upload(data map[string]interface{}, contractAddress string, signerAddress string) (string, error)
	UploadBatch(data []map[string]interface{}, fileStartNumber int, contractAddress string, signerAddress string) (*baseUriWithUris, error)
}

type uploadResponse struct {
	IpfsHash    string   `json:"ipfsHash"`
	PinSize     *big.Int `json:"PinSize"`
	Timestamp   string   `json:"Timestamp"`
	IsDuplicate bool     `json:"isDuplicate"`
	IpfsUri     string   `json:"IpfsUri"`
}

type IpfsStorage struct {
	GatewayUrl string
}

func newIpfsStorage(gatewayUrl string) *IpfsStorage {
	return &IpfsStorage{
		GatewayUrl: gatewayUrl,
	}
}

// Get
//
// # Get IPFS data at a given hash and return it as byte data
//
// uri: the IPFS URI to fetch data from
//
// returns: byte data of the IPFS data at the URI
func (ipfs *IpfsStorage) Get(uri string) ([]byte, error) {
	gatewayUrl := replaceHashWithGatewayUrl(uri, ipfs.GatewayUrl)
	resp, err := http.Get(gatewayUrl)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("Bad status code, %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
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
func (ipfs *IpfsStorage) Upload(data map[string]interface{}, contractAddress string, signerAddress string) (string, error) {
	baseUriWithUris, err := ipfs.UploadBatch([]map[string]interface{}{data}, 0, contractAddress, signerAddress)
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
func (ipfs *IpfsStorage) UploadBatch(data []map[string]interface{}, fileStartNumber int, contractAddress string, signerAddress string) (*baseUriWithUris, error) {
	preparedData, err := ipfs.batchUploadProperties(data)
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

	baseUriWithUris, err := ipfs.uploadBatchWithCid(dataToUpload, fileStartNumber, contractAddress, signerAddress)
	if err != nil {
		return nil, err
	}

	return baseUriWithUris, nil
}

func (ipfs *IpfsStorage) getUploadToken(contractAddress string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf("%v/grant", twIpfsServerUrl), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("X-App-Name", fmt.Sprintf("CONSOLE-GO-SDK-%v", contractAddress))
	result, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if result.StatusCode != http.StatusOK {
		return "", &failedToUploadError{
			statusCode: result.StatusCode,
		}
	}

	body, err := ioutil.ReadAll(result.Body)
	text := string(body)

	return text, nil
}

func (ipfs *IpfsStorage) uploadBatchWithCid(
	// data (string | io.Reader)[] - file or JSON string
	data []interface{},
	fileStartNumber int,
	contractAddress string,
	signerAddress string,
) (*baseUriWithUris, error) {
	uploadToken, err := ipfs.getUploadToken(contractAddress)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
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

			part.Write(jsonData)
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

	req, err := http.NewRequest("POST", pinataIpfsUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", uploadToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if result, err := client.Do(req); err != nil {
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
func (ipfs *IpfsStorage) batchUploadProperties(data []map[string]interface{}) (interface{}, error) {
	sanitizedMetadatas, err := ipfs.replaceGatewayUrlWithHash(data, "ipfs://", ipfs.GatewayUrl)
	if err != nil {
		return nil, err
	}

	filesToUpload, err := ipfs.buildFilePropertiesMap(sanitizedMetadatas, []interface{}{})
	if err != nil {
		return nil, err
	}

	if len(filesToUpload) == 0 {
		return sanitizedMetadatas, nil
	}

	baseUriWithUris, err := ipfs.uploadBatchWithCid(filesToUpload, 0, "", "")
	if err != nil {
		return nil, err
	}

	replacedMetadatas, err := ipfs.replaceFilePropertiesWithHashes(sanitizedMetadatas, baseUriWithUris.uris)
	if err != nil {
		return nil, err
	}

	return replacedMetadatas, nil
}

// data - array or map or strings
// Returns []io.Reader files to upload
func (ipfs *IpfsStorage) buildFilePropertiesMap(data interface{}, files []interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			builtFiles, err := ipfs.buildFilePropertiesMap(v.Index(i).Interface(), files)
			if err != nil {
				return nil, err
			}

			files = builtFiles
		}
		break
	case reflect.Map:
		for _, k := range v.MapKeys() {
			builtFiles, err := ipfs.buildFilePropertiesMap(v.MapIndex(k).Interface(), files)
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

func (ipfs *IpfsStorage) replaceFilePropertiesWithHashes(data interface{}, cids []string) (interface{}, error) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		updated := []interface{}{}
		for i := 0; i < v.Len(); i++ {
			val, err := ipfs.replaceFilePropertiesWithHashes(v.Index(i).Interface(), cids)
			if err != nil {
				return nil, err
			}

			updated = append(updated, val)
		}

		return updated, nil
	case reflect.Map:
		updated := map[string]interface{}{}
		for _, k := range v.MapKeys() {
			val, err := ipfs.replaceFilePropertiesWithHashes(v.MapIndex(k).Interface(), cids)
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

func (ipfs *IpfsStorage) replaceGatewayUrlWithHash(data interface{}, scheme string, gatewayUrl string) (interface{}, error) {
	v := reflect.ValueOf(data)
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			ipfs.replaceGatewayUrlWithHash(v.Index(i), scheme, gatewayUrl)
		}
		break
	case reflect.Map:
		for _, k := range v.MapKeys() {
			ipfs.replaceGatewayUrlWithHash(v.MapIndex(k), scheme, gatewayUrl)
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
