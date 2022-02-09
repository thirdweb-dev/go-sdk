package nftlabs

import (
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/internal/abi"
)

type Marketplace interface {
	GetListing(listingId *big.Int) (Listing, error)
	GetAll(filter ListingFilter) ([]Listing, error)
	GetMarketFeeBps() (*big.Int, error)
	SetMarketFeeBps(fee *big.Int) error
}

type MarketplaceModule struct {
	Client  *ethclient.Client
	Address string
	module  *abi.Marketplace

	main ISdk
}

func newMarketplaceModule(client *ethclient.Client, address string, main ISdk) (*MarketplaceModule, error) {
	module, err := abi.NewMarketplace(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	return &MarketplaceModule{
		Client:  client,
		Address: address,
		module:  module,
		main: main,
	}, nil
}

func (sdk *MarketplaceModule) GetMarketFeeBps() (*big.Int, error) {
	if result, err := sdk.module.MarketplaceCaller.MarketFeeBps(&bind.CallOpts{}); err != nil {
		return nil, err
	} else {
		return big.NewInt(0).SetUint64(result), nil
	}
}

func (sdk *MarketplaceModule) SetMarketFeeBps(fee *big.Int) error {
	if tx, err := sdk.module.SetMarketFeeBps(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, fee); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *MarketplaceModule) GetListing(listingId *big.Int) (Listing, error) {
	if result, err := sdk.module.MarketplaceCaller.Listings(&bind.CallOpts{}, listingId); err != nil {
		return Listing{}, err
	} else {
		return sdk.transformResultToListing(result)
	}
}

func (sdk *MarketplaceModule) GetAll(filter ListingFilter) ([]Listing, error) {
	listings := make([]abi.IMarketplaceListing, 0)
	filteredListings := make([]abi.IMarketplaceListing, 0)
	results := make([]Listing, 0)

	hasFilter := filter.TokenContract != "" || filter.TokenId != nil || filter.Seller != ""

	totalListings, err := sdk.module.MarketplaceCaller.TotalListings(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	zero := big.NewInt(0)
	one := big.NewInt(1)
	for i := new(big.Int).Set(zero); i.Cmp(totalListings) < 0; i.Add(i, one) {
		listing, err := sdk.module.MarketplaceCaller.Listings(&bind.CallOpts{}, i)
		if err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	if hasFilter {
		for _, listing := range listings {
			if listing.Quantity.Cmp(zero) == 0 {
				continue
			}
			if !sdk.isStillValidDirectListing(listing) {
				continue
			}
			if filter.TokenContract != "" && common.HexToAddress(filter.TokenContract) != listing.AssetContract {
				continue
			}
			if filter.TokenId != nil && filter.TokenId.Cmp(listing.TokenId) != 0 {
				continue
			}
			if filter.Seller != "" && common.HexToAddress(filter.Seller) != listing.TokenOwner {
				continue
			}
			filteredListings = append(filteredListings, listing)
		}
	}

	for _, listing := range filteredListings {
		transformed, err := sdk.transformResultToListing(listing)
		if err != nil {
			return nil, err
		}
		results = append(results, transformed)
	}

	return results, nil
}

func (sdk *MarketplaceModule) isStillValidDirectListing(listing abi.IMarketplaceListing) bool {
	// TODO: check token owner balance
	// TODO: check token owner approval
	return true
}

func (sdk *MarketplaceModule) transformResultToListing(listing abi.IMarketplaceListing) (Listing, error) {
	listingCurrency := listing.Currency

	var currencyMetadata *CurrencyValue
	if listingCurrency.String() == "0x0000000000000000000000000000000000000000" || strings.ToLower(listingCurrency.String()) == "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee" {
		currencyMetadata = nil
	} else {
		// TODO: this is bad, don't want to create an instance of the module every time
		// but idk how else to get it in here given that the address is dynamic per listing
		// damages testability
		currency, err := newCurrencyModule(sdk.Client, listingCurrency.Hex(), sdk.main)
		if err != nil {
			// TODO: return better error
			return Listing{}, err
		}

		if currencyValue, err := currency.GetValue(listing.BuyoutPricePerToken); err != nil {
			// TODO: return better error
			return Listing{}, err
		} else {
			currencyMetadata = &currencyValue
		}
	}

	var nftMetadata *NftMetadata
	if listing.AssetContract.String() != "0x0000000000000000000000000000000000000000" {
		// TODO: again, bad, need to create this in the function because we don't know the nft contract when we get here
		// damages testability
		nftModule, err := newNftModule(sdk.Client, listing.AssetContract.Hex(), sdk.main)
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
	if listing.StartTime.Int64() > 0 {
		time.Unix(listing.StartTime.Int64()*1000, 0)
	} else {
		saleStart = nil
	}

	var saleEnd *time.Time
	if listing.EndTime.Int64() > 0 && listing.EndTime.Int64() < math.MaxInt64 - 1 {
		time.Unix(listing.EndTime.Int64()*1000, 0)
	} else {
		saleEnd = nil
	}

	return Listing{
		Id:               listing.ListingId,
		Seller:           listing.TokenOwner,
		TokenContract:    listing.AssetContract,
		TokenId:          listing.TokenId,
		TokenMetadata:    nftMetadata,
		Quantity:         listing.Quantity,
		CurrentContract:  listingCurrency,
		CurrencyMetadata: currencyMetadata,
		Price:            listing.BuyoutPricePerToken,
		SaleStart:        saleStart,
		SaleEnd:          saleEnd,
	}, nil
}

