package thirdweb

import (
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type ERC721 struct {
	contractWrapper *ContractWrapper[*abi.TokenERC721]
}

func NewERC721(contractWrapper *ContractWrapper[*abi.TokenERC721]) *ERC721 {
	return &ERC721{
		contractWrapper,
	}
}
