package thirdweb

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

// You can access the Marketplace interface from the SDK as follows:
//
// 	import (
// 		"github.com/thirdweb-dev/go-sdk/thirdweb"
// 	)
//
// 	privateKey = "..."
//
// 	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
// 	})
//
//	contract, err := sdk.GetMarketplace("{{contract_address}}")
type Marketplace struct {
	abi     *abi.Marketplace
	helper  *contractHelper
	storage storage
}

func newMarketplace(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Marketplace, error) {
	if contractAbi, err := abi.NewMarketplace(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		marketplace := &Marketplace{
			contractAbi,
			helper,
			storage,
		}
		return marketplace, nil
	}
}

// Get a single listing from the marketplace.
//
// listingId: listing ID to get
//
// returns: listing at the given listing ID
//
// Example
//
// 	listingId := 0
// 	listing, err := marketplace.GetListing(listingId)
func (marketplace *Marketplace) GetListing(listingId int) (*DirectListing, error) {
	listing, err := marketplace.abi.Listings(&bind.CallOpts{}, big.NewInt(int64(listingId)))
	if err != nil {
		return nil, err
	}

	if listing.AssetContract.String() == zeroAddress {
		return nil, fmt.Errorf("Failed to find listing with ID %d", listingId)
	}

	if listing.ListingType == 0 {
		return marketplace.mapListing(listing)
	} else {
		return nil, fmt.Errorf("Unknown listing type: %d", listingId)
	}
}

// Get all active listings from the marketplace.
//
// returns: all active listings in the marketplace
//
// Example
//
// 	listings, err := marketplace.GetActiveListings()
// 	// Price per token of the first listing
// 	listings[0].BuyoutCurrencyValuePerToken.DisplayValue
func (marketplace *Marketplace) GetActiveListings() ([]*DirectListing, error) {
	listings, err := marketplace.getAllListingsNoFilter()
	if err != nil {
		return nil, err
	}

	return listings, err
}

// Get all the listings from the marketplace.
//
// returns: all listings in the marketplace
//
// Example
// 	listings, err := marketplace.GetAllListings()
// 	// Price per token of the first listing
// 	listings[0].BuyoutCurrencyValuePerToken.DisplayValue
func (marketplace *Marketplace) GetAllListings() ([]*DirectListing, error) {
	return marketplace.getAllListingsNoFilter()
}

// Get the total number of listings in the marketplace.
//
// returns: total number of listings in the marketplace
func (marketplace *Marketplace) GetTotalCount() (int, error) {
	total, err := marketplace.abi.TotalListings(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return int(total.Int64()), nil
}

func (marketplace *Marketplace) CreateListing(listing *NewDirectListing) (int, error) {
	err := handleTokenApproval(
		marketplace.helper.GetProvider(),
		marketplace.helper,
		marketplace.helper.getAddress().Hex(),
		listing.AssetContractAddress,
		listing.TokenId,
		marketplace.helper.GetSignerAddress().Hex(),
	)
	if err != nil {
		return 0, err
	}

	normalizedPricePerToken, err := normalizePriceValue(
		marketplace.helper.GetProvider(),
		listing.BuyoutPricePerToken,
		listing.CurrencyContractAddress,
	)
	if err != nil {
		return 0, err
	}

	txOpts, err := marketplace.helper.getTxOptions()
	if err != nil {
		return 0, err
	}
	tx, err := marketplace.abi.CreateListing(txOpts, abi.IMarketplaceListingParameters{
		AssetContract:        common.HexToAddress(listing.AssetContractAddress),
		TokenId:              big.NewInt(int64(listing.TokenId)),
		StartTime:            big.NewInt(int64(listing.StartTimeInEpochSeconds)),
		SecondsUntilEndTime:  big.NewInt(int64(listing.ListingDurationInSeconds)),
		QuantityToList:       big.NewInt(int64(listing.Quantity)),
		CurrencyToAccept:     common.HexToAddress(listing.CurrencyContractAddress),
		ReservePricePerToken: normalizedPricePerToken,
		BuyoutPricePerToken:  normalizedPricePerToken,
		ListingType:          0,
	})

	receipt, err := marketplace.helper.awaitTx(tx.Hash())

	txReceipt, err := marketplace.helper.GetProvider().TransactionReceipt(context.Background(), receipt.Hash())
	if err != nil {
		return 0, err
	}

	for _, log := range txReceipt.Logs {
		event, err := marketplace.abi.ParseListingAdded(*log)
		if err != nil {
			continue
		}

		return int(event.ListingId.Int64()), nil
	}

	return 0, errors.New("No ListingAdded event found")
}

func (marketplace *Marketplace) mapListing(listing abi.IMarketplaceListing) (*DirectListing, error) {
	currencyValue, err := fetchCurrencyValue(
		marketplace.helper.GetProvider(),
		listing.Currency.String(),
		listing.BuyoutPricePerToken,
	)
	if err != nil {
		return nil, err
	}

	asset, err := fetchTokenMetadataForContract(
		listing.AssetContract.String(),
		marketplace.helper.GetProvider(),
		int(listing.TokenId.Int64()),
		marketplace.storage,
	)
	if err != nil {
		return nil, err
	}

	return &DirectListing{
		AssetContractAddress:        listing.AssetContract.String(),
		BuyoutPrice:                 int(listing.BuyoutPricePerToken.Int64()),
		CurrencyContractAddress:     listing.Currency.String(),
		BuyoutCurrencyValuePerToken: currencyValue,
		Id:                          listing.ListingId.String(),
		TokenId:                     int(listing.TokenId.Int64()),
		Quantity:                    int(listing.Quantity.Int64()),
		StartTimeInEpochSeconds:     int(listing.StartTime.Int64()),
		EndTimeInEpochSeconds:       int(listing.EndTime.Int64()),
		SellerAddress:               listing.TokenOwner.String(),
		Asset:                       asset,
	}, nil
}

func (marketplace *Marketplace) getAllListingsNoFilter() ([]*DirectListing, error) {
	totalCount, err := marketplace.abi.TotalListings(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	listings := []*DirectListing{}
	for id := 0; id < int(totalCount.Int64()); id++ {
		listing, err := marketplace.GetListing(id)
		if err != nil {
			return nil, err
		}
		listings = append(listings, listing)
	}

	return listings, nil
}
