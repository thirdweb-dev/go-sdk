package nftlabs

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"errors"
)

type CloudflareGateway struct {
	Url string
}

func NewCloudflareGateway(uri string) *CloudflareGateway {
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