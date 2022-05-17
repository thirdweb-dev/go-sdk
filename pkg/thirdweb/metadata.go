package thirdweb

import (
	"math/big"
)

type Metadata struct {
	MetadataUri    string
	MetadataObject interface{}
}

type NFTMetadata struct {
	Id              *big.Int    `json:"id"`
	Uri             string      `json:"uri"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Image           string      `json:"image"`
	ExternalUrl     string      `json:"external_url"`
	AnimationUrl    string      `json:"animation_url"`
	BackgroundColor string      `json:"background_color"`
	Properties      interface{} `json:"properties"`
}

type NFTMetadataInput struct {
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Image           string      `json:"image"`
	ExternalUrl     string      `json:"external_url"`
	AnimationUrl    string      `json:"animation_url"`
	BackgroundColor string      `json:"background_color"`
	Properties      interface{} `json:"properties"`
}

type NFTMetadataOwner struct {
	Metadata *NFTMetadata
	Owner    string
}

type EditionMetadata struct {
	Metadata *NFTMetadata
	Supply   int
}

type EditionMetadataOwner struct {
	Metadata      *NFTMetadata
	Supply        int
	Owner         string
	QuantityOwned int
}

type EditionMetadataInput struct {
	Metadata *NFTMetadataInput
	Supply   int
}
