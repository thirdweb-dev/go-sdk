package nftlabs

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
	"math/big"
)

type MarketSdk interface {
	Get(listingId *big.Int) (Listing, error)
	GetAll() ([]Listing, error)
	List(
		assetContract string,
		tokenId string,
		currentContract string,
		price *big.Int,
		quantity *big.Int,
		secondsUntilStart uint64,
		secondsUntilEnd uint64) (Listing, error)
	UnlistAll(listingId *big.Int) error
	Unlist(listingId *big.Int, quantity *big.Int) error
	Buy(listingId *big.Int, quantity *big.Int) error
}

type MarketSdkModule struct {
	CurrencySdk *CurrencySdk
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	caller *abi.MarketCaller
}

func NewMarketSdkModule(client *ethclient.Client, currencySdk *CurrencySdk, address string, opt *SdkOptions) (*MarketSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewMarketCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	var gw Gateway
	gw = NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &MarketSdkModule{
		CurrencySdk: currencySdk,
		Client: client,
		Address: address,
		Options: opt,
		gateway: gw,
		caller: caller,
	}, nil
}

func (m *MarketSdkModule) Get(listingId *big.Int) (Listing, error) {
	if result, err := m.caller.Listings(&bind.CallOpts{}, listingId); err != nil {
		return Listing{}, err
	} else {
		return m.transformResultToListing(result)
	}
}

func (m *MarketSdkModule) transformResultToListing(listing abi.MarketListing) (Listing, error) {
	//var currency CurrencyValue
	//
	//if c, err :=

	return Listing{}, nil
}

func (m *MarketSdkModule) GetAll() ([]Listing, error) {
	panic("implement me")
}

func (m *MarketSdkModule) List(assetContract string, tokenId string, currentContract string, price *big.Int, quantity *big.Int, secondsUntilStart uint64, secondsUntilEnd uint64) (Listing, error) {
	panic("implement me")
}

func (m *MarketSdkModule) UnlistAll(listingId *big.Int) error {
	panic("implement me")
}

func (m *MarketSdkModule) Unlist(listingId *big.Int, quantity *big.Int) error {
	panic("implement me")
}

func (m *MarketSdkModule) Buy(listingId *big.Int, quantity *big.Int) error {
	panic("implement me")
}
