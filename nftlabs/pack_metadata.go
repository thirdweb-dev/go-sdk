package nftlabs

import (
	"math/big"
	"time"
)

type Pack struct {
	NftMetadata
	Creator string
	CurrentSupply big.Int
	OpenStart time.Time
	OpenEnd time.Time
}

type PackNft struct {
	NftMetadata
	Supply big.Int
}