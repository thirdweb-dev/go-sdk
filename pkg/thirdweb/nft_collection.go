package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTCollection struct {
	*ERC721
}

func NewNFTCollection(provider *ethclient.Client, address common.Address, privateKey string, storage Storage) (*NFTCollection, error) {
	if abi, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := NewContractWrapper[abi.TokenERC721](abi, provider, privateKey); err != nil {
			return nil, err
		} else {
			erc721 := NewERC721(contractWrapper, storage)
			nftCollection := &NFTCollection{
				erc721,
			}
			return nftCollection, nil
		}
	}
}
