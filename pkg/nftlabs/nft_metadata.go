package nftlabs

import "math/big"

type NftMetadata struct {
	Id          *big.Int    `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Properties  interface{} `json:"properties"`
}

type MintNftMetadata struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Properties  interface{} `json:"properties"`
}
