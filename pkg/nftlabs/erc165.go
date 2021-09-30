package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/internal/abi"
)

type erc165 interface {
	CommonModule
}

type erc165Module struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
	module  *abi.ERC165

	privateKey    *ecdsa.PrivateKey
	signerAddress common.Address
}

func newErc165SdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*erc165Module, error) {
	module, err := abi.NewERC165(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	return &erc165Module{
		Client:  client,
		Address: address,
		Options: opt,
		module:  module,
	}, nil
}

func (sdk *erc165Module) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return &NoSignerError{typeName: "erc165", Err: err}
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}
func (sdk *erc165Module) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}
