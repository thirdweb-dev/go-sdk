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

// Cancel a listing on the marketplace.
//
// listingId: listing ID to cancel
//
// returns: transaction receipt of the cancellation
//
// Example
//
// 	listingId := 0
// 	receipt, err := marketplace.CancelListing(listingId)
func (marketplace *Marketplace) CancelListing(listingId int) (*types.Transaction, error) {
	txOpts, err := marketplace.helper.getTxOptions()
	if err != nil {
		return nil, err
	}

	tx, err := marketplace.abi.CancelDirectListing(txOpts, big.NewInt(int64(listingId)))
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
func (marketplace *Marketplace) BuyoutListing(listingId int, quantityDesired int) (*types.Transaction, error) {
	return marketplace.BuyoutListingTo(listingId, quantityDesired, marketplace.helper.GetSignerAddress().Hex())
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
// 	listingId := 0
// 	quantityDesired := 1
// 	receiver := "0x..."
// 	receipt, err := marketplace.BuyoutListingTo(listingId, quantityDesired, receiver)
func (marketplace *Marketplace) BuyoutListingTo(listingId int, quantityDesired int, receiver string) (*types.Transaction, error) {
	listing, err := marketplace.validateListing(listingId)
	if err != nil {
		return nil, err
	}

	valid, err := marketplace.isStillValidListing(listing, quantityDesired)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("The asset on this listing has been moved from the lister's wallet, this listing is now invalid")
	}

	quantity := big.NewInt(int64(quantityDesired))
	value := big.NewInt(0).Mul(big.NewInt(int64(listing.BuyoutPrice)), quantity)

	txOpts, err := marketplace.helper.getTxOptions()
	if err != nil {
		return nil, err
	}

	setErc20Allowance(
		marketplace.helper,
		value,
		listing.CurrencyContractAddress,
		txOpts,
	)

	tx, err := marketplace.abi.Buy(
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
// 	listing := &NewDirectListing{
// 		AssetContractAddress: "0x...", // Address of the asset contract
// 		TokenId: 0, // Token ID of the asset to list
// 		StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
// 		ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
// 		Quantity: 1, // Quantity of the asset to list
// 		CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
// 		BuyoutPricePerToken: 1, // Price per token of the asset to list
// 	}
//
// 	listingId, err := marketplace.CreateListing(listing)
func (marketplace *Marketplace) CreateListing(listing *NewDirectListing) (int, error) {
	listing.fillDefaults()

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

func (marketplace *Marketplace) validateListing(listingId int) (*DirectListing, error) {
	listing, err := marketplace.GetListing(listingId)
	if err != nil {
		return nil, fmt.Errorf("Error getting the listing with ID %d", listingId)
	}

	return listing, nil
}

func (marketplace *Marketplace) isStillValidListing(listing *DirectListing, quantity int) (bool, error) {
	approved, err := isTokenApprovedForTransfer(
		marketplace.helper.GetProvider(),
		marketplace.helper.getAddress().Hex(),
		listing.AssetContractAddress,
		listing.TokenId,
		listing.SellerAddress,
	)
	if err != nil {
		return false, err
	}

	if !approved {
		return false, nil
	}

	erc165, err := abi.NewIERC165(common.HexToAddress(listing.AssetContractAddress), marketplace.helper.GetProvider())
	if err != nil {
		return false, err
	}

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return false, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return false, err
	}

	if isErc721 {
		contract, err := abi.NewTokenERC721(
			common.HexToAddress(listing.AssetContractAddress),
			marketplace.helper.GetProvider(),
		)
		if err != nil {
			return false, err
		}

		ownerOf, err := contract.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(listing.TokenId)))
		if err != nil {
			return false, err
		}

		return strings.ToLower(ownerOf.Hex()) == strings.ToLower(listing.SellerAddress), nil
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(
			common.HexToAddress(listing.AssetContractAddress),
			marketplace.helper.GetProvider(),
		)
		if err != nil {
			return false, err
		}

		balance, err := contract.BalanceOf(
			&bind.CallOpts{},
			common.HexToAddress(listing.SellerAddress),
			big.NewInt(int64(listing.TokenId)),
		)

		return balance.Int64() >= int64(quantity), nil
	} else {
		return false, errors.New("Contract does not implement ERC721 or ERC1155")
	}
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
