package nftlabs

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type Listing struct {
	Id               *big.Int `json:"id"`
	Seller           common.Address `json:"seller"`
	TokenContract    common.Address `json:"tokenContract"`
	TokenId          *big.Int `json:"tokenId"`
	TokenMetadata    *NftMetadata `json:"tokenMetadata"`
	Quantity         *big.Int `json:"quantity"`
	CurrentContract  common.Address `json:"currentContract"`
	CurrencyMetadata *CurrencyValue `json:"currencyMetadata"`
	Price            *big.Int `json:"price"`
	SaleStart        *time.Time `json:"saleStart"`
	SaleEnd          *time.Time `json:"saleEnd"`
}

type ListingFilter struct {
	Seller        string   `json:"seller"`
	TokenContract string   `json:"tokenContract"`
	TokenId       *big.Int `json:"tokenId"`
}

type NewListingArgs struct {
	AssetContractAddress    string `json:"assetContractAddress"`
	TokenId                 *big.Int `json:"tokenId"`
	CurrencyContractAddress string `json:"currencyContractAddress"`
	Price                   *big.Int `json:"price"`
	Quantity                *big.Int `json:"quantity"`
	SecondsUntilOpenStart   *big.Int `json:"secondsUntilOpenStart"`
	SecondsUntilOpenEnd     *big.Int `json:"secondsUntilOpenEnd"`
	RewardsPerOpen          *big.Int `json:"rewardsPerOpen"`
}
