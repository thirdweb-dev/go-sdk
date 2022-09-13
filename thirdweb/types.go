package thirdweb

import (
	"math/big"
	"time"
)

type SDKOptions struct {
	PrivateKey string
	GatewayUrl string
}

type Metadata struct {
	MetadataUri    string
	MetadataObject interface{}
}

type NFTMetadata struct {
	Id              *big.Int    `json:"id"`
	Uri             string      `json:"uri"`
	Name            string      `json:"name"`
	Description     string      `json:"description"`
	Image           interface{} `json:"image"`
	ExternalUrl     string      `json:"external_url"`
	AnimationUrl    string      `json:"animation_url"`
	BackgroundColor string      `json:"background_color"`
	Properties      interface{} `json:"properties,omitempty"`
	Attributes      interface{} `json:"attributes,omitempty"`
}

type NFTMetadataInput struct {
	Name            string      `mapstructure:"name" json:"name"`
	Description     string      `mapstructure:"description,omitempty" json:"description"`
	Image           interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalUrl     string      `mapstructure:"external_url,omitempty" json:"external_url"`
	AnimationUrl    string      `mapstructure:"animation_url,omitempty" json:"animation_url"`
	BackgroundColor string      `mapstructure:"background_color,omitempty" json:"background_color"`
	Properties      interface{} `mapstructure:"properties,omitempty" json:"properties,omitempty"`
	Attributes      interface{} `mapstructure:"attributes,omitempty" json:"attributes,omitempty"`
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
	price                     *big.Int
	currencyAddress           string
	value                     *big.Int
}

type ClaimConditionOutput struct {
	Price                       *big.Int
	MaxQuantity                 *big.Int
	QuantityLimitPerTransaction *big.Int
	WaitInSeconds               *big.Int
	StartTime                   *big.Int
	AvailableSupply             *big.Int
	CurrencyAddress             string
	CurrencyMetadata            *CurrencyValue
}

type Currency struct {
	Name     string
	Symbol   string
	Decimals int
}

type CurrencyValue struct {
	Name         string
	Symbol       string
	Decimals     int
	Value        *big.Int
	DisplayValue float64
}

type TokenAmount struct {
	ToAddress string
	Amount    float64
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

type Signature721PayloadInput struct {
	To                   string
	Price                float64
	CurrencyAddress      string
	MintStartTime        int
	MintEndTime          int
	PrimarySaleRecipient string
	Metadata             *NFTMetadataInput
	RoyaltyRecipient     string
	RoyaltyBps           int
}

type Signature721PayloadOutput struct {
	To                   string
	Price                float64
	CurrencyAddress      string
	MintStartTime        int
	MintEndTime          int
	PrimarySaleRecipient string
	Metadata             *NFTMetadataInput
	RoyaltyRecipient     string
	RoyaltyBps           int
	Uri                  string
	Uid                  [32]byte
}

type SignedPayload721 struct {
	Payload   *Signature721PayloadOutput
	Signature []byte
}

type Signature1155PayloadInput struct {
	To                   string
	Price                float64
	CurrencyAddress      string
	MintStartTime        int
	MintEndTime          int
	PrimarySaleRecipient string
	Metadata             *NFTMetadataInput
	RoyaltyRecipient     string
	RoyaltyBps           int
	Quantity             int
}

type Signature1155PayloadInputWithTokenId struct {
	To                   string
	Price                float64
	CurrencyAddress      string
	MintStartTime        int
	MintEndTime          int
	PrimarySaleRecipient string
	Metadata             *NFTMetadataInput
	TokenId              int
	RoyaltyRecipient     string
	RoyaltyBps           int
	Quantity             int
}

type Signature1155PayloadOutput struct {
	To                   string
	Price                float64
	CurrencyAddress      string
	MintStartTime        int
	MintEndTime          int
	PrimarySaleRecipient string
	Metadata             *NFTMetadataInput
	RoyaltyRecipient     string
	RoyaltyBps           int
	TokenId              int
	Quantity             int
	Uri                  string
	Uid                  [32]byte
}

type SignedPayload1155 struct {
	Payload   *Signature1155PayloadOutput
	Signature []byte
}

type MultiwrapERC20 struct {
	ContractAddress string
	Quantity        float64
}

type MultiwrapERC721 struct {
	ContractAddress string
	TokenId         int
}

type MultiwrapERC1155 struct {
	ContractAddress string
	TokenId         int
	Quantity        int
}

type MultiwrapBundle struct {
	ERC20Tokens   []*MultiwrapERC20
	ERC721Tokens  []*MultiwrapERC721
	ERC1155Tokens []*MultiwrapERC1155
}

// Wallet Authenticator

type WalletLoginOptions struct {
	Nonce          string
	ExpirationTime time.Time
	ChainId        int
}

type WalletLoginPayloadData struct {
	Domain         string    `json:"domain"`
	Address        string    `json:"address"`
	Nonce          string    `json:"nonce"`
	ExpirationTime time.Time `json:"expiration_time"`
	ChainId        int       `json:"chain_id"`
}

type WalletLoginPayload struct {
	Payload   *WalletLoginPayloadData `json:"payload"`
	Signature string                  `json:"signature"`
}

type WalletVerifyOptions struct {
	ChainId int
}

type WalletAuthenticationOptions struct {
	InvalidBefore  time.Time
	ExpirationTime time.Time
}

type WalletAuthenticationPayloadData struct {
	Iss string `json:"iss"`
	Sub string `json:"sub"`
	Aud string `json:"aud"`
	Exp int64  `json:"exp"`
	Nbf int64  `json:"nbf"`
	Iat int64  `json:"iat"`
	Jti string `json:"jti"`
}

// Contract Metadata

type DeployNFTCollectionMetadata struct {
	Name                   string      `mapstructure:"name" json:"name"`
	Description            string      `mapstructure:"description" json:"description"`
	Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink           string      `mapstructure:"external_link,omitempty" json:"external_link"`
	SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
	FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
	Symbol                 string      `mapstructure:"symbol" json:"symbol"`
	PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
	PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
	PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
	TrustedForwarders      []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployNFTCollectionMetadata) fillDefaults() {
	if metadata.FeeRecipient == "" {
		metadata.FeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PrimarySaleRecipient == "" {
		metadata.PrimarySaleRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PlatformFeeRecipient == "" {
		metadata.PlatformFeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DeployEditionMetadata struct {
	Name                   string      `mapstructure:"name" json:"name"`
	Description            string      `mapstructure:"description" json:"description"`
	Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink           string      `mapstructure:"external_link,omitempty" json:"external_link"`
	SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
	FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
	Symbol                 string      `mapstructure:"symbol" json:"symbol"`
	PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
	PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
	PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
	TrustedForwarders      []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployEditionMetadata) fillDefaults() {
	if metadata.FeeRecipient == "" {
		metadata.FeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PrimarySaleRecipient == "" {
		metadata.PrimarySaleRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PlatformFeeRecipient == "" {
		metadata.PlatformFeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DeployTokenMetadata struct {
	Name                   string      `mapstructure:"name" json:"name"`
	Description            string      `mapstructure:"description" json:"description"`
	Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink           string      `mapstructure:"external_link,omitempty" json:"external_link"`
	Symbol                 string      `mapstructure:"symbol" json:"symbol"`
	PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
	PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
	PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
	TrustedForwarders      []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployTokenMetadata) fillDefaults() {
	if metadata.PrimarySaleRecipient == "" {
		metadata.PrimarySaleRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PlatformFeeRecipient == "" {
		metadata.PlatformFeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DeployNFTDropMetadata struct {
	Name                   string      `mapstructure:"name" json:"name"`
	Description            string      `mapstructure:"description" json:"description"`
	Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink           string      `mapstructure:"external_link,omitempty" json:"external_link"`
	SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
	FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
	Merkle                 interface{} `mapstructure:"merkle" json:"merkle"`
	Symbol                 string      `mapstructure:"symbol" json:"symbol"`
	PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
	PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
	PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
	TrustedForwarders      []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployNFTDropMetadata) fillDefaults() {
	if metadata.FeeRecipient == "" {
		metadata.FeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PrimarySaleRecipient == "" {
		metadata.PrimarySaleRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PlatformFeeRecipient == "" {
		metadata.PlatformFeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DeployEditionDropMetadata struct {
	Name                   string      `mapstructure:"name" json:"name"`
	Description            string      `mapstructure:"description" json:"description"`
	Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink           string      `mapstructure:"external_link,omitempty" json:"external_link"`
	SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
	FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
	Merkle                 interface{} `mapstructure:"merkle" json:"merkle"`
	Symbol                 string      `mapstructure:"symbol" json:"symbol"`
	PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
	PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
	PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
	TrustedForwarders      []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployEditionDropMetadata) fillDefaults() {
	if metadata.FeeRecipient == "" {
		metadata.FeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PrimarySaleRecipient == "" {
		metadata.PrimarySaleRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.PlatformFeeRecipient == "" {
		metadata.PlatformFeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DeployMultiwrapMetadata struct {
	Name                 string      `mapstructure:"name" json:"name"`
	Description          string      `mapstructure:"description" json:"description"`
	Image                interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink         string      `mapstructure:"external_link,omitempty" json:"external_link"`
	SellerFeeBasisPoints int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
	FeeRecipient         string      `mapstructure:"fee_recipient" json:"fee_recipient"`
	Symbol               string      `mapstructure:"symbol" json:"symbol"`
	TrustedForwarders    []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployMultiwrapMetadata) fillDefaults() {
	if metadata.FeeRecipient == "" {
		metadata.FeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DeployMarketplaceMetadata struct {
	Name                   string      `mapstructure:"name" json:"name"`
	Description            string      `mapstructure:"description" json:"description"`
	Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
	ExternalLink           string      `mapstructure:"external_link,omitempty" json:"external_link"`
	PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
	PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
	TrustedForwarders      []string    `mapstructure:"trusted_forwarders" json:"trusted_forwarders"`
}

func (metadata *DeployMarketplaceMetadata) fillDefaults() {
	if metadata.PlatformFeeRecipient == "" {
		metadata.PlatformFeeRecipient = "0x0000000000000000000000000000000000000000"
	}

	if metadata.TrustedForwarders == nil {
		metadata.TrustedForwarders = []string{}
	}
}

type DirectListing struct {
	Id                          string
	AssetContractAddress        string
	TokenId                     int
	Asset                       *NFTMetadata
	StartTimeInEpochSeconds     int
	EndTimeInEpochSeconds       int
	Quantity                    int
	CurrencyContractAddress     string
	BuyoutCurrencyValuePerToken *CurrencyValue
	BuyoutPrice                 string
	SellerAddress               string
}

type NewDirectListing struct {
	AssetContractAddress     string
	TokenId                  int
	StartTimeInEpochSeconds  int
	ListingDurationInSeconds int
	Quantity                 int
	CurrencyContractAddress  string
	BuyoutPricePerToken      float64
}

func (listing *NewDirectListing) fillDefaults() {
	if listing.CurrencyContractAddress == "" {
		listing.CurrencyContractAddress = "0x0000000000000000000000000000000000000000"
	}

	if listing.StartTimeInEpochSeconds == 0 {
		listing.StartTimeInEpochSeconds = int(time.Now().Unix())
	}

	if listing.ListingDurationInSeconds == 0 {
		listing.ListingDurationInSeconds = int(time.Now().Unix() + 3600)
	}
}

type AuctionListing struct {
	Id                                string
	AssetContractAddress              string
	TokenId                           int
	Asset                             *NFTMetadata
	StartTimeInEpochSeconds           int
	EndTimeInEpochSeconds             int
	Quantity                          int
	CurrencyContractAddress           string
	ReservePrice                      string
	BuyoutPrice                       string
	BuyoutCurrencyValuePerToken       *CurrencyValue
	ReservePriceCurrencyValuePerToken *CurrencyValue
	SellerAddress                     string
}

type MarketplaceFilter struct {
	Start         int
	Count         int
	Seller        string
	TokenContract string
}
