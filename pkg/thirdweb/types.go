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

type ClaimVerification struct {
	proofs                    [][32]byte
	maxQuantityPerTransaction int
	price                     float64
	currencyAddress           string
}

type ClaimCondition struct {
	startTime                   int
	quantityLimitPerTransaction int
	maxQuantity                 int
	waitInSeconds               int
	currencyAddress             string
	price                       float64
}

type ClaimConditionOutput struct {
	price                       float64
	maxQuantity                 int
	quantityLimitPerTransaction int
	waitInSeconds               int
	startTime                   int
	availableSupply             int
	currencyAddress             string
	currencyMetadata            *CurrencyValue
}

type Currency struct {
	name     string
	symbol   string
	decimals int
}

type CurrencyValue struct {
	name         string
	symbol       string
	decimals     int
	value        *big.Int
	displayValue float64
}

type WrappedToken struct {
	address string
	name    string
	symbol  string
}

type NativeToken struct {
	name     string
	symbol   string
	decimals int
	wrapper  *WrappedToken
}

// Contract Metadata

type NFTCollectionContractMetadata struct {
	Name                   string
	Description            string
	Image                  string
	ExternalLink           string
	SellerFeeBasisPoints   int
	FeeRecipient           string
	Symbol                 string
	PlatformFeeBasisPoints int
	PlatformFeeRecipient   string
	TrustedForwarder       string
}

type EditionContractMetadata struct {
	Name                   string
	Description            string
	Image                  string
	ExternalLink           string
	SellerFeeBasisPoints   int
	FeeRecipient           string
	Symbol                 string
	PlatformFeeBasisPoints int
	PlatformFeeRecipient   string
	TrustedForwarder       string
}

type NFTDropContractMetadata struct {
	Name                   string
	Description            string
	Image                  string
	ExternalLink           string
	SellerFeeBasisPoints   int
	FeeRecipient           string
	Symbol                 string
	PrimarySaleRecipient   string
	PlatformFeeBasisPoints int
	PlatformFeeRecipient   string
	TrustedForwarder       string
}
