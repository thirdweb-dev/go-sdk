package nftlabs

import (
	"math/big"
	"time"
)

type Listing struct {
	Id string
	Seller string
	TokenContract string
	TokenId string
	TokenMetadata NftMetadata
	Quantity big.Int
	CurrentContract string
	CurrencyMetadata interface {} // TODO: use currency type here
	Price big.Int
	SaleStart time.Time
	SaleEnd time.Time
}
