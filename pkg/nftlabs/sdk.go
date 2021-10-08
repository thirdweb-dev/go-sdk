package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type ISdk interface {
	GetNftModule(address string) (Nft, error)
	GetMarketModule(address string) (Market, error)
	GetCurrencyModule(address string) (Currency, error)
	GetPackModule(address string) (Pack, error)
	GetNftCollectionModule(address string) (NftCollection, error)
	GetStorage() (Storage, error)
	TransferNativeToken(to string, amount *big.Int) error

	SetStorage(gateway Storage)

	getSignerAddress() common.Address
	getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error)
	getRawPrivateKey() string
	getOptions() *SdkOptions
	getGateway() Storage
	getTransactOpts(send bool) *bind.TransactOpts
}

type Sdk struct {
	client *ethclient.Client
	opt *SdkOptions

	privateKey    *ecdsa.PrivateKey
	rawPrivateKey string
	signerAddress common.Address

	nftModule Nft
	marketModule Market
	currencyModule Currency
	packModule Pack
	nftCollectionModule NftCollection

	gateway Storage
}

func NewSdk(client *ethclient.Client, opt *SdkOptions) (*Sdk, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	if opt.GasLimit == 0 {
		opt.GasLimit = 10000 // TODO: what should this be?
	}

	defaultGateway := newIpfsStorage(opt.IpfsGatewayUrl)
	sdk := &Sdk{
		client: client,
		opt: opt,
		gateway: defaultGateway,
	}

	if opt.PrivateKey != "" {
		if err := sdk.setPrivateKey(opt.PrivateKey); err != nil {
			return nil, err
		}
	}

	return sdk, nil
}

func (sdk *Sdk) GetCurrencyModule(address string) (Currency, error) {
	if sdk.currencyModule != nil {
		return sdk.currencyModule, nil
	}

	module, err := newCurrencyModule(sdk.client, address, sdk)
	if err != nil {
		return nil, err
	}

	sdk.currencyModule = module
	return module, nil
}


func (sdk *Sdk) GetMarketModule(address string) (Market, error) {
	if sdk.marketModule != nil {
		return sdk.marketModule, nil
	}

	module, err := newMarketModule(sdk.client, address, sdk)
	if err != nil {
		return nil, err
	}

	sdk.marketModule = module
	return module, nil
}

func (sdk *Sdk) GetNftModule(address string) (Nft, error) {
	if sdk.nftModule != nil {
		return sdk.nftModule, nil
	}

	module, err := newNftModule(sdk.client, address, sdk)
	if err != nil {
		return nil, err
	}

	sdk.nftModule = module
	return module, nil
}

func (sdk *Sdk) GetPackModule(address string) (Pack, error) {
	if sdk.packModule != nil {
		return sdk.packModule, nil
	}

	module, err := newPackModule(sdk.client, address, sdk)
	if err != nil {
		return nil, err
	}

	sdk.packModule = module
	return module, nil
}


func (sdk *Sdk) GetStorage() (Storage, error) {
	if sdk.gateway != nil {
		return sdk.gateway, nil
	}

	module := newIpfsStorage(sdk.opt.IpfsGatewayUrl)

	sdk.gateway = module
	return module, nil
}

func (sdk *Sdk) GetNftCollectionModule(address string) (NftCollection, error) {
	if sdk.nftCollectionModule != nil {
		return sdk.nftCollectionModule, nil
	}

	module, err := newNftCollectionModule(sdk.client, address, sdk)
	if err != nil {
		return nil, err
	}

	sdk.nftCollectionModule = module
	return module, nil
}

func (sdk *Sdk) setPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
		sdk.rawPrivateKey = privateKey
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

func (sdk *Sdk) getRawPrivateKey() string {
	return sdk.rawPrivateKey
}

func (sdk *Sdk) getOptions() *SdkOptions {
	return sdk.opt
}

func (sdk *Sdk) SetStorage(gateway Storage) {
	sdk.gateway = gateway
}

func (sdk *Sdk) getGateway() Storage {
	return sdk.gateway
}

func (sdk *Sdk) getTransactOpts(send bool) *bind.TransactOpts {
	return &bind.TransactOpts{
		NoSend: !send,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
		GasPrice: sdk.opt.MaxGasPriceInGwei,
	}
}

func (sdk *Sdk) TransferNativeToken(to string, amount *big.Int) error {
	if sdk.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "SDK"}
	}

	nonce, err := sdk.client.PendingNonceAt(context.Background(), sdk.getSignerAddress())
	if err != nil {
		return err
	}

	gasLimit := uint64(21000)                // in units
	gasPrice, err := sdk.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(to), amount, gasLimit, gasPrice, data)

	chainID, err := sdk.client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), sdk.privateKey)
	if err != nil {
		return err
	}

	err = sdk.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	return nil
}

