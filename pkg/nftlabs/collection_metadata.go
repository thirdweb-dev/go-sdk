package nftlabs

import (
	"math/big"
)

type CollectionMetadata struct {
	NftMetadata
	Creator string   `json:"creator"`
	Supply  *big.Int `json:"supply"`
}

type MintCollectionArgs struct {
	TokenId *big.Int `json:"tokenId"`
	Amount  *big.Int `json:"amount"`
}

type collectionAssetMetadata struct {
	Uri    string
	Supply *big.Int
}

type NftCollectionBatchArgs struct {
	Amount *big.Int `json:"amount"`
	TokenId *big.Int `json:"tokenId"`
}

type CreateCollectionArgs struct {
	Metadata Metadata
	Supply *big.Int
}
