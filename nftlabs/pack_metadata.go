package nftlabs

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
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
	Supply *big.Int
}

type PackNftAddition struct {
	NftId *big.Int
	Supply *big.Int
}
