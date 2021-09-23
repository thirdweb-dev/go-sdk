package nftlabs

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type MarketSdk interface {
	CommonModule
	GetListing(listingId *big.Int) (Listing, error)
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
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	caller *abi.MarketCaller

	privateKey *ecdsa.PrivateKey
	signerAddress common.Address
}

func NewMarketSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*MarketSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewMarketCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	gw := NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &MarketSdkModule{
		Client: client,
		Address: address,
		Options: opt,
		gateway: gw,
		caller: caller,
	}, nil
}

func (m *MarketSdkModule) GetListing(listingId *big.Int) (Listing, error) {
	if result, err := m.caller.Listings(&bind.CallOpts{}, listingId); err != nil {
		return Listing{}, err
	} else {
		return m.transformResultToListing(result)
	}
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

func (m *MarketSdkModule) transformResultToListing(listing abi.MarketListing) (Listing, error) {
	listingCurrency := listing.Currency

	var currencyMetadata *CurrencyValue
	if strings.HasPrefix(listingCurrency.String(), "0x000000000000") {
		currencyMetadata = nil
	} else {
		// TODO: this is bad, don't want to create an instance of the module every time but idk how else to get it in here
		// damages testability
		log.Printf("Getting listing currency at address %v\n", listingCurrency)
		currency, err := NewCurrencySdkModule(m.Client, listingCurrency.Hex())
		if err != nil {
			// TODO: return better error
			return Listing{}, err
		}

		
		if currencyValue, err := currency.GetValue(listing.PricePerToken); err != nil {
			// TODO: return better error
			return Listing{}, err
		} else {
			currencyMetadata = currencyValue
		}
	}

	var nftMetadata *NftMetadata
	if !strings.HasPrefix(listing.AssetContract.String(), "0x000000000000") {
		log.Printf("Getting nft module at %v", listing.AssetContract)
		// TODO: again, bad, need to create this in the function because we don't know the nft contract when we get here
		// damages testability
		nftModule, err := NewNftSdkModule(m.Client, listing.AssetContract.Hex(), &SdkOptions{})
		if err != nil {
			// TODO: return better error
			return Listing{}, err
		}
		
		if meta, err := nftModule.Get(listing.TokenId); err != nil {
			// TODO: return better error
			return Listing{}, err
		} else {
			nftMetadata = &meta
		}
	}

	var saleStart *time.Time
	// TODO: should I be doing Int64() here ??? is there data loss ???
	if listing.SaleStart.Int64() > 0 {
		time.Unix(listing.SaleStart.Int64() * 1000, 0)
	} else {
		saleStart = nil
	}

	var saleEnd *time.Time
	// TODO: should I be doing Int64() here ??? is there data loss ???
	if listing.SaleEnd.Int64() > 0 && listing.SaleEnd.Int64() < big.MaxExp - 1 {
		time.Unix(listing.SaleEnd.Int64() * 1000, 0)
	} else {
		saleEnd = nil
	}

	return Listing{
		Id:               listing.ListingId,
		Seller:           listing.Seller,
		TokenContract:    listing.AssetContract,
		TokenId:          listing.TokenId,
		TokenMetadata:    nftMetadata,
		Quantity:         listing.Quantity,
		CurrentContract:  listingCurrency,
		CurrencyMetadata: currencyMetadata,
		Price:            listing.PricePerToken,
		SaleStart:        saleStart,
		SaleEnd:          saleEnd,
	}, nil
}

func (sdk *MarketSdkModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}
