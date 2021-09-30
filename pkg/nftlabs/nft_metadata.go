package nftlabs

import "math/big"

type NftMetadata struct {
	Id          *big.Int `json:"id"`
	Uri         string      `json:"external_url"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
}

type MintNftMetadata struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Image                string   `json:"image"`
	ExternalUrl          string   `json:"external_url"`
	SellerFeeBasisPoints *big.Int `json:"seller_fee_basis_points"`
	FeeRecipient         string   `json:"fee_recipient"`
	BackgroundColor      string   `json:"background_color"`
}
