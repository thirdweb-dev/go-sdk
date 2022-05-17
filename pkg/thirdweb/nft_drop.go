package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTDrop struct {
	*DropERC721
}

func NewNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage Storage) (*NFTDrop, error) {
	if dropErc721, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := NewContractWrapper(dropErc721, provider, privateKey); err != nil {
			return nil, err
		} else {
			erc721 := NewDropERC721(contractWrapper, storage)
			nftDrop := &NFTDrop{
				erc721,
			}
			return nftDrop, nil
		}
	}
}
