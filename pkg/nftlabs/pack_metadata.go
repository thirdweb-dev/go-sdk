package nftlabs

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type PackMetadata struct {
	NftMetadata
	Creator       common.Address
	CurrentSupply *big.Int
	OpenStart     time.Time
}

type PackNft struct {
	NftMetadata
	Supply *big.Int
}

/**
  assetContract: string;
  metadata: MetadataURIOrObject;
  assets: {
    tokenId: BigNumberish;
    amount: BigNumberish;
  }[];
  secondsUntilOpenStart?: number;
  secondsUntilOpenEnd?: number;
  rewardsPerOpen?: number;
*/

type CreatePackArgs struct {
	AssetContractAddress  string
	Assets                []PackAssetAddition
	SecondsUntilOpenStart *big.Int
	SecondsUntilOpenEnd   *big.Int
	RewardsPerOpen        *big.Int
	Metadata              interface{}
}

type PackAssetAddition struct {
	NftId  *big.Int
	Supply *big.Int
}
