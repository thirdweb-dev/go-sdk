package thirdweb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
	"strings"
)

type BaseUriWithUris struct {
	baseUri string
	uris    []string
}
type Storage interface {
	Get(uri string) ([]byte, error)
	Upload(data any, contractAddress string, signerAddress string) (string, error)
	UploadBatch(data []any, contractAddress string, signerAddress string) (*BaseUriWithUris, error)
}

type uploadResponse struct {
	IpfsHash    string   `json:"ipfsHash"`
	PinSize     *big.Int `json:"PinSize"`
	Timestamp   string   `json:"Timestamp"`
	IsDuplicate bool     `json:"isDuplicate"`
	IpfsUri     string   `json:"IpfsUri"`
}

const (
	nftLabsApiUrl = "https://upload.nftlabs.co"
)

type IpfsStorage struct {
	Url string
}

func NewIpfsStorage(uri string) Storage {
	return &IpfsStorage{
		Url: uri,
	}
}

func (gw *IpfsStorage) Get(uri string) ([]byte, error) {
	gatewayUrl := replaceIpfsPrefixWithGateway(uri, gw.Url)
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

// Upload method can be used to upload a generic payload to IPFS. NftLabs provides a default proxy
// in the SDK. You can override this with the ISdk.SetStorage
func (ipfs *IpfsStorage) Upload(data any, contractAddress string, signerAddress string) (string, error) {
	baseUriWithUris, err := ipfs.UploadBatch([]any{data}, contractAddress, signerAddress)
	if err != nil {
		return "", err
	}

	baseUri := baseUriWithUris.baseUri + "0"
	return baseUri, nil
}

// UploadBatch uploads a list of arbitrary objects and returns their URIs *in the order they were passed*
func (ipfs *IpfsStorage) UploadBatch(data []any, contractAddress string, signerAddress string) (*BaseUriWithUris, error) {
	baseUriWithUris, err := ipfs.uploadBatchWithCid(data, contractAddress, signerAddress)
	if err != nil {
		return nil, err
	}

	return baseUriWithUris, nil
}

func (ipfs *IpfsStorage) uploadBatchWithCid(
	data []any,
	contractAddress string,
	signerAddress string,
) (*BaseUriWithUris, error) {
	client := &http.Client{}
	fileNames := []string{}

	body := &bytes.Buffer{}
	mp := multipart.NewWriter(body)
	for i, asset := range data {
		jsonData, err := json.Marshal(asset)
		if err != nil {
			return nil, err
		}

		fileName := fmt.Sprintf("%v", i)
		fileNames = append(fileNames, fileName)
		if err := mp.WriteField(fileName, string(jsonData)); err != nil {
			return nil, err
		}
	}

	_ = mp.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/upload", nftLabsApiUrl), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-App-Name", fmt.Sprintf("CONSOLE-GO-SDK-%v", contractAddress))
	req.Header.Set("X-Public-Address", signerAddress)
	req.Header.Set("Content-Type", mp.FormDataContentType())

	if result, err := client.Do(req); err != nil {
		return nil, err
	} else {
		if result.StatusCode != http.StatusOK {
			return nil, &FailedToUploadError{
				statusCode: result.StatusCode,
				Payload:    data,
			}
		}

		var uploadMeta uploadResponse
		bodyBytes, err := ioutil.ReadAll(result.Body)
		if err != nil {
			return nil, &FailedToUploadError{
				statusCode:      result.StatusCode,
				Payload:         data,
				UnderlyingError: err,
			}
		}

		if err := json.Unmarshal(bodyBytes, &uploadMeta); err != nil {
			return nil, &UnmarshalError{
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

		return &BaseUriWithUris{
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
