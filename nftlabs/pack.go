package nftlabs

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
	"math/big"
)

type PackSdk interface {
	Open(packId big.Int) (NftMetadata, error)
	Get(tokenId big.Int) (NftMetadata, error)
	GetAll() ([]Pack, error)
	GetNfts(packId big.Int) ([]PackNft, error)
	Balance(tokenId big.Int) (big.Int, error)
	BalanceOf(address string, tokenId string) uint64
	Transfer(to string, tokenId string, quantity big.Int) (error)
}

type PackSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	caller *abi.NFTCaller
}