package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ISdk interface {
	GetNftModule(address string) (Nft, error)

	getSignerAddress() common.Address
	getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error)
}

type Sdk struct {
	client *ethclient.Client
	opt *SdkOptions

	privateKey    *ecdsa.PrivateKey
	rawPrivateKey string
	signerAddress common.Address

	nftModule Nft
}

func NewSdk(client *ethclient.Client, opt *SdkOptions) (*Sdk, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	sdk := &Sdk{
		client: client,
		opt: opt,
	}

	if opt.PrivateKey != "" {
		if err := sdk.setPrivateKey(opt.PrivateKey); err != nil {
			return nil, err
		}
	}

	return sdk, nil
}

func (sdk *Sdk) GetNftModule(address string) (Nft, error) {
	if sdk.nftModule != nil {
		return sdk.nftModule, nil
	}

	module, err := NewNftSdkModule(sdk.client, address, sdk.opt, sdk)
	if err != nil {
		return nil, err
	}

	sdk.nftModule = module
	return module, nil
}

func (sdk *Sdk) setPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}

func (sdk *Sdk) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}

func (sdk *Sdk) getSignerAddress() common.Address {
	return sdk.signerAddress
}
