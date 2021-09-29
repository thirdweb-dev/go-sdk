package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type erc1155 interface {
	CommonModule
}

type erc1155Module struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	module  *abi.ERC1155

	privateKey    *ecdsa.PrivateKey
	signerAddress common.Address
}

func newErc1155SdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*erc1155Module, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	module, err := abi.NewERC1155(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	var gw Gateway
	gw = NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &erc1155Module{
		Client:  client,
		Address: address,
		Options: opt,
		gateway: gw,
		module:  module,
	}, nil
}

func (sdk *erc1155Module) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return &NoSignerError{typeName: "erc1155", Err: err}
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}
func (sdk *erc1155Module) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}
