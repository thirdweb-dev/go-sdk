package nftlabs

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"time"
)

type Pack struct {
	NftMetadata
	Creator common.Address
	CurrentSupply big.Int
	OpenStart time.Time
	OpenEnd time.Time
}

type PackNft struct {
	NftMetadata
	Supply big.Int
}