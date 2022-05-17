package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTDrop struct {
	*ContractWrapper[*abi.DropERC721]
	*ERC721
}

func NewNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage Storage) (*NFTDrop, error) {
	if dropAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := NewContractWrapper(dropAbi, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := NewERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				nftCollection := &NFTDrop{
					contractWrapper,
					erc721,
				}
				return nftCollection, nil
			}
		}
	}
}
