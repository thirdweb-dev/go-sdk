package nftlabs

import "math/big"

type CollectionMetadata struct {
	NftMetadata
	Creator string   `json:"creator"`
	Supply  *big.Int `json:"supply"`
}

type CreateCollectionArgs struct {
	Supply   *big.Int    `json:"supply"`
	Metadata interface{} `json:"metadata"`
}

type MintCollectionArgs struct {
	TokenId *big.Int `json:"tokenId"`
	Amount  *big.Int `json:"amount"`
}

type collectionAssetMetadata struct {
	Uri    string
	Supply *big.Int
}
