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

	"github.com/mitchellh/mapstructure"
)

type baseUriWithUris struct {
	baseUri string
	uris    []string
}
type storage interface {
	Get(uri string) ([]byte, error)
	Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
	UploadBatch(data []interface{}, fileStartNumber int, contractAddress string, signerAddress string) (*baseUriWithUris, error)
}

type uploadResponse struct {
	IpfsHash    string   `json:"ipfsHash"`
	PinSize     *big.Int `json:"PinSize"`
	Timestamp   string   `json:"Timestamp"`
	IsDuplicate bool     `json:"isDuplicate"`
	IpfsUri     string   `json:"IpfsUri"`
}

type IpfsStorage struct {
	Url string
}

func newIpfsStorage(uri string) *IpfsStorage {
	return &IpfsStorage{
		Url: uri,
	}
}

// Get
//
// Get IPFS data at a given hash and return it as byte data
//
// uri: the IPFS URI to fetch data from
//
// returns: byte data of the IPFS data at the URI
func (ipfs *IpfsStorage) Get(uri string) ([]byte, error) {
	gatewayUrl := replaceIpfsPrefixWithGateway(uri, ipfs.Url)
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
	baseUriWithUris, err := ipfs.uploadBatchWithCid(data, fileStartNumber, contractAddress, signerAddress)
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

// TODO: Take map as inputs instead of structs
func (ipfs *IpfsStorage) _uploadBatchWithCid(
	data []map[string]interface{},
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

	for i, metadata := range data {
		jsonData, err := json.Marshal(metadata)
		if err != nil {
			return nil, err
		}

		fileName := fmt.Sprintf("%v", i+fileStartNumber)
		fileNames = append(fileNames, fileName)

		part, err := writer.CreateFormFile("file", fmt.Sprintf("files/%v", fileName))
		part.Write(jsonData)
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

// TODO: Take map as inputs instead of structs
func (ipfs *IpfsStorage) uploadBatchWithCid(
	data []map[string]interface{},
	fileStartNumber int,
	contractAddress string,
	signerAddress string,
) (*baseUriWithUris, error) {
	uploadToken, err := ipfs.getUploadToken(contractAddress)

	client := &http.Client{}
	fileNames := []string{}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if err != nil {
		return nil, err
	}

	for i, asset := range data {

		file, ok := asset.(io.Reader)
		if ok {
			fileName := fmt.Sprintf("%v", i+fileStartNumber)
			fileNames = append(fileNames, fileName)

			part, _ := writer.CreateFormFile("file", fmt.Sprintf("files/%v", fileName))
			io.Copy(part, file)
		} else {
			assetToUpload := map[string]interface{}{}
			err := mapstructure.Decode(asset, &assetToUpload)
			if err != nil {
				return nil, err
			}

			// Omit null fields from map
			assetMetadata := map[string]interface{}{}
			for k, v := range assetToUpload {
				if v != nil {
					assetMetadata[k] = v
				}
			}

			// If there is an image field that has a file, upload image to IPFS
			if image, ok := assetMetadata["image"]; ok && image != nil {
				if _, ok := image.(io.Reader); ok {
					uri, err := ipfs.Upload(assetMetadata["image"], contractAddress, signerAddress)
					if err != nil {
						return nil, err
					}
					assetMetadata["image"] = uri
					assetToUpload = assetMetadata
				} else if reflect.TypeOf(image).String() != "string" {
					assetMetadata["image"] = ""
				}
			}

			// Necessary fix for dashboard bug on deployment, but should be outside when
			// we refactor this function to take a map instead
			if externalLink, ok := assetMetadata["external_link"]; ok && externalLink == "" {
				delete(assetMetadata, "external_link")
				assetToUpload = assetMetadata
			}

			jsonData, err := json.Marshal(assetToUpload)
			if err != nil {
				return nil, err
			}

			fileName := fmt.Sprintf("%v", i+fileStartNumber)
			fileNames = append(fileNames, fileName)

			part, err := writer.CreateFormFile("file", fmt.Sprintf("files/%v", fileName))
			part.Write(jsonData)
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

func replaceIpfsPrefixWithGateway(ipfsUrl string, gatewayUrl string) string {
	if ipfsUrl == "" {
		return ""
	}

	gateway := gatewayUrl
	if !strings.HasSuffix(gatewayUrl, "/") {
		gateway = gatewayUrl + "/"
	}

	return strings.Replace(ipfsUrl, "ipfs://", gateway, 1)
}
