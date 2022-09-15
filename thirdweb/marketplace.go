package thirdweb

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// You can access the Marketplace interface from the SDK as follows:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	contract, err := sdk.GetMarketplace("{{contract_address}}")
type Marketplace struct {
	Abi     *abi.Marketplace
	helper  *contractHelper
	storage storage
	Encoder *MarketplaceEncoder
}

func newMarketplace(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Marketplace, error) {
	if contractAbi, err := abi.NewMarketplace(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		encoder, err := newMarketplaceEncoder(contractAbi, helper, storage)
		if err != nil {
			return nil, err
		}

		marketplace := &Marketplace{
			Abi:     contractAbi,
			helper:  helper,
			storage: storage,
			Encoder: encoder,
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
//	listingId := 0
//	listing, err := marketplace.GetListing(context.Background(), listingId)
func (marketplace *Marketplace) GetListing(ctx context.Context, listingId int) (*DirectListing, error) {
	listing, err := marketplace.Abi.Listings(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(listingId)))
	if err != nil {
		return nil, err
	}

	// If listing does not exist or is cancelled, return nil as the listing
	if listing.AssetContract.String() == zeroAddress {
		return nil, fmt.Errorf("Failed to find listing with ID %d", listingId)
	}

	if listing.ListingType == 0 {
		return mapListing(marketplace.helper, marketplace.storage, listing)
	} else {
		return nil, fmt.Errorf("Unsupported listing type: %d. Currently only direct listings are supported.", listingId)
	}
}

// Get all active listings from the marketplace.
//
// filter: optional filter parameters
//
// returns: all active listings in the marketplace
//
// Example
//
//	listings, err := marketplace.GetActiveListings(context.Background(), nil)
//	// Price per token of the first listing
//	listings[0].BuyoutCurrencyValuePerToken.DisplayValue
func (marketplace *Marketplace) GetActiveListings(ctx context.Context, filter *MarketplaceFilter) ([]*DirectListing, error) {
	listings, err := marketplace.getAllListingsNoFilter(ctx)
	if err != nil {
		return nil, err
	}

	listings, err = marketplace.applyFilter(listings, filter)
	if err != nil {
		return nil, err
	}

	activeListings := []*DirectListing{}
	for _, listing := range listings {
		if listing.Quantity > 0 {
			activeListings = append(activeListings, listing)
		}
	}

	return activeListings, nil
}

// Get all the listings from the marketplace.
//
// filter: optional filter parameters
//
// returns: all listings in the marketplace
//
// Example
//
//	listings, err := marketplace.GetAllListings(context.Background(), nil)
//	// Price per token of the first listing
//	listings[0].BuyoutCurrencyValuePerToken.DisplayValue
func (marketplace *Marketplace) GetAllListings(ctx context.Context, filter *MarketplaceFilter) ([]*DirectListing, error) {
	listings, err := marketplace.getAllListingsNoFilter(ctx)
	if err != nil {
		return nil, err
	}

	return marketplace.applyFilter(listings, filter)
}

// Get the total number of listings in the marketplace.
//
// returns: total number of listings in the marketplace
func (marketplace *Marketplace) GetTotalCount(ctx context.Context) (int, error) {
	total, err := marketplace.Abi.TotalListings(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return 0, err
	}

	return int(total.Int64()), nil
}

// Cancel a listing on the marketplace.
//
// listingId: listing ID to cancel
//
// returns: transaction receipt of the cancellation
//
// Example
//
//	listingId := 0
//	receipt, err := marketplace.CancelListing(context.Background(), listingId)
func (marketplace *Marketplace) CancelListing(ctx context.Context, listingId int) (*types.Transaction, error) {
	txOpts, err := marketplace.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := marketplace.Abi.CancelDirectListing(txOpts, big.NewInt(int64(listingId)))
	if err != nil {
		return nil, err
	}

	return marketplace.helper.awaitTx(tx.Hash())
}

// Buy a specific listing from the marketplace.
//
// listingId: listing ID of the asset you want to buy
//
// quantityDesired: the quantity of the asset to buy from the listing
//
// returns: transaction receipt of the purchase
func (marketplace *Marketplace) BuyoutListing(ctx context.Context, listingId int, quantityDesired int) (*types.Transaction, error) {
	return marketplace.BuyoutListingTo(ctx, listingId, quantityDesired, marketplace.helper.GetSignerAddress().Hex())
}

// Buy a specific listing from the marketplace to a specific address.
//
// listingId: listing ID of the asset you want to buy
//
// quantityDesired: the quantity of the asset to buy from the listing
//
// receiver: specific address to receive the assets from the listing
//
// returns: transaction receipt of the purchase
//
// Example
//
//	listingId := 0
//	quantityDesired := 1
//	receiver := "0x..."
//	receipt, err := marketplace.BuyoutListingTo(context.Background(), listingId, quantityDesired, receiver)
func (marketplace *Marketplace) BuyoutListingTo(ctx context.Context, listingId int, quantityDesired int, receiver string) (*types.Transaction, error) {
	listing, err := marketplace.validateListing(ctx, listingId)
	if err != nil {
		return nil, err
	}

	valid, err := isStillValidListing(marketplace.helper, listing, quantityDesired)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("The asset on this listing has been moved from the lister's wallet, this listing is now invalid")
	}

	quantity := big.NewInt(int64(quantityDesired))
	value := listing.BuyoutCurrencyValuePerToken.Value.Mul(listing.BuyoutCurrencyValuePerToken.Value, quantity)

	txOpts, err := marketplace.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	err = setErc20Allowance(
		marketplace.helper,
		value,
		listing.CurrencyContractAddress,
		txOpts,
	)
	if err != nil {
		return nil, err
	}

	tx, err := marketplace.Abi.Buy(
		txOpts,
		big.NewInt(int64(listingId)),
		common.HexToAddress(receiver),
		big.NewInt(int64(quantityDesired)),
		common.HexToAddress(listing.CurrencyContractAddress),
		value,
	)
	if err != nil {
		return nil, err
	}

	return marketplace.helper.awaitTx(tx.Hash())
}

// Create a new listing on the marketplace where people can buy an asset directly.
//
// listing: the data for the listing to create
//
// returns: the ID of the listing that was created
//
// Example
//
//	listing := &NewDirectListing{
//		AssetContractAddress: "0x...", // Address of the asset contract
//		TokenId: 0, // Token ID of the asset to list
//		StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
//		ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
//		Quantity: 1, // Quantity of the asset to list
//		CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
//		BuyoutPricePerToken: 1, // Price per token of the asset to list
//	}
//
//	listingId, err := marketplace.CreateListing(context.Background(), listing)
func (marketplace *Marketplace) CreateListing(ctx context.Context, listing *NewDirectListing) (int, error) {
	listing.fillDefaults()

	err := handleTokenApproval(
		ctx,
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

	txOpts, err := marketplace.helper.getTxOptions(ctx)
	if err != nil {
		return 0, err
	}
	tx, err := marketplace.Abi.CreateListing(txOpts, abi.IMarketplaceListingParameters{
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
		event, err := marketplace.Abi.ParseListingAdded(*log)
		if err != nil {
			continue
		}

		return int(event.ListingId.Int64()), nil
	}

	return 0, errors.New("No ListingAdded event found")
}

func (marketplace *Marketplace) validateListing(ctx context.Context, listingId int) (*DirectListing, error) {
	listing, err := marketplace.GetListing(ctx, listingId)
	if err != nil {
		return nil, fmt.Errorf("Error getting the listing with ID %d", listingId)
	}

	return listing, nil
}

func (marketplace *Marketplace) getAllListingsNoFilter(ctx context.Context) ([]*DirectListing, error) {
	totalCount, err := marketplace.Abi.TotalListings(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return nil, err
	}

	listings := []*DirectListing{}
	for id := 0; id < int(totalCount.Int64()); id++ {
		listing, err := marketplace.GetListing(ctx, id)
		if err != nil {
			if strings.Contains(err.Error(), "Failed to find listing") || strings.Contains(err.Error(), "Unsupported listing type") {
				continue
			} else {
				return nil, err
			}
		}
		listings = append(listings, listing)
	}

	return listings, nil
}

func (marketplace *Marketplace) applyFilter(listings []*DirectListing, filter *MarketplaceFilter) ([]*DirectListing, error) {
	if filter == nil {
		return listings, nil
	}

	filteredListings := listings

	if filter.Seller != "" {
		rawListings := []*DirectListing{}
		for _, listing := range filteredListings {
			if strings.ToLower(listing.SellerAddress) == strings.ToLower(filter.Seller) {
				rawListings = append(rawListings, listing)
			}
		}

		filteredListings = rawListings
	}

	if filter.TokenContract != "" {
		rawListings := []*DirectListing{}
		for _, listing := range filteredListings {
			if strings.ToLower(listing.AssetContractAddress) == strings.ToLower(filter.TokenContract) {
				rawListings = append(rawListings, listing)
			}
		}

		filteredListings = rawListings
	}

	if len(filteredListings) == 0 {
		return filteredListings, nil
	}

	start := 0
	count := 100

	if filter.Start != 0 {
		start = filter.Start
	}

	if filter.Count != 0 {
		count = filter.Count
	}

	if start > len(filteredListings)-1 {
		return nil, fmt.Errorf("Start index %d is out of bounds for %d total listings", start, len(filteredListings))
	}

	end := start + count
	if start+count > len(filteredListings) {
		end = len(filteredListings)
	}

	filteredListings = filteredListings[start:end]
	return filteredListings, nil
}
