package nftlabs

import "math/big"

type NftMetadata struct {
	Id          *big.Int `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
}

type MintNftMetadata struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Image                string   `json:"image"`
	Attributes  interface{} `json:"attributes"`
}
