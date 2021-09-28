package nftlabs

import "math/big"

type uploadResponse struct {
	IpfsHash string `json:"ipfsHash"`
	PinSize *big.Int `json:"PinSize"`
	Timestamp string `json:"Timestamp"`
	IsDuplicate bool `json:"isDuplicate"`
	IpfsUri string `json:"IpfsUri"`
}

type Gateway interface {
	Get(uri string) ([]byte, error)
	Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
}
