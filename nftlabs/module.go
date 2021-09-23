package nftlabs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SigningMethod = func(common.Address, *types.Transaction) (*types.Transaction, error)

type CommonModule interface {
	SetSigningMethod(signer SigningMethod)
	SetSigningAddress(publicKey string)
}