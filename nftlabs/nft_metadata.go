package nftlabs

import "math/big"

type NftMetadata struct {
	Id *big.Int
	Uri string `json:"external_url"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Attributes interface{} `json:"attributes"`
}