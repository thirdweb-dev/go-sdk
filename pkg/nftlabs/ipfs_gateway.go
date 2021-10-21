package nftlabs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"log"
	"math/big"
	"mime/multipart"
	"net/http"
)

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

func newIpfsStorage(uri string) Storage {
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

func (gw *IpfsStorage) Upload(data interface{}, contractAddress string, signerAddress string) (string, error) {
	if meta, ok := data.(Metadata); ok && meta.MetadataUri != "" {
		return meta.MetadataUri, nil
	}

	client := &http.Client{}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	body := &bytes.Buffer{}
	mp := multipart.NewWriter(body)
	if err := mp.WriteField("file", string(jsonData)); err != nil {
		return "", err
	}

	_ = mp.Close()

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/upload", nftLabsApiUrl), body)
	if err != nil {
		return "", err
	}
	req.Header.Set("X-App-Name", fmt.Sprintf("CONSOLE-GO-SDK-%v", contractAddress))
	req.Header.Set("X-Public-Address", signerAddress)
	req.Header.Set("Content-Type", mp.FormDataContentType())

	if result, err := client.Do(req); err != nil {
		return "", err
	} else {
		if result.StatusCode != http.StatusOK {
			//TODO: return better error
			return "", errors.New(fmt.Sprintf("HTTP request failed, status code %v", result.StatusCode))
		}

		var uploadMeta uploadResponse
		bodyBytes, err := ioutil.ReadAll(result.Body)
		if err != nil {
			return "", err
		}

		if err := json.Unmarshal(bodyBytes, &uploadMeta); err != nil {
			return "", &UnmarshalError{
				body:            string(bodyBytes),
				typeName:        "UploadResponse",
				underlyingError: err,
			}
		}
		return uploadMeta.IpfsUri, nil
	}
}

// UploadBatch uploads a list of arbitrary objects and returns their URIs *in the order they were passed*
func (gw *IpfsStorage) UploadBatch(data []interface{}, contractAddress string, signerAddress string) ([]string, error) {
	wg := new(errgroup.Group)

	results := make([]string, len(data))
	for i, asset := range data {
		func(meta interface{}, index int) {
			wg.Go(func() error {
				toUpload := meta
				if meta, ok := meta.(Metadata); ok {
					if meta.MetadataUri != "" {
						 results[index] = meta.MetadataUri
						 return nil
					} else {
						toUpload = meta.MetadataObject
					}
				}

				uri, err := gw.Upload(toUpload, contractAddress, signerAddress)
				if err != nil {
					return err
				} else {
					results[index] = uri
					return nil
				}
			})
		}(asset, i)
	}

	if err := wg.Wait(); err != nil {
		log.Println("Failed to upload a portion of the metadata, aborting")
		return nil, err
	}

	return results, nil
}
