package nftlabs

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type Listing struct {
	Id *big.Int
	Seller common.Address
	TokenContract common.Address
	TokenId *big.Int
	TokenMetadata NftMetadata
	Quantity *big.Int
	CurrentContract common.Address
	CurrencyMetadata CurrencyValue // TODO: use currency type here
	Price *big.Int
	SaleStart *time.Time
	SaleEnd *time.Time
}
