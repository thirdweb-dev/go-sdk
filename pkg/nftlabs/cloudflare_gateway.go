package nftlabs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

const (
	nftLabsApiUrl = "https://upload.nftlabs.co"
)

type CloudflareGateway struct {
	Url string
}

func newCloudflareGateway(uri string) Gateway {
	return &CloudflareGateway{
		Url: uri,
	}
}

func (gw *CloudflareGateway) Get(uri string) ([]byte, error) {
	gatewayUrl := replaceIpfsWithGateway(uri, gw.Url)
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

func (gw *CloudflareGateway) Upload(data interface{}, contractAddress string, signerAddress string) (string, error) {
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