package main

import (
	"encoding/json"
	"fmt"
	"github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
)

const (
	newListingAssetContractAddressFlag = "assetAddress"
	newListingCurrencyContractAddressFlag = "currencyAddress"
)

var (
	marketplaceContractAddress string
	newMarketplaceListing nftlabs.NewListingArgs
)

var marketplaceCmd = &cobra.Command{
	Use:   "marketplace [command]",
	Short: "Interact with a marketplace contract",
}

var marketplaceGetAllCmd = &cobra.Command{
	Use: "getAll",
	Short: "Gets all listings in a marketplace",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getMarketplaceModule()
		if err != nil {
			panic(err)
		}

		allListings, err := module.GetAll(nftlabs.ListingFilter{})
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d listings in collection %v\n", len(allListings), marketplaceContractAddress)
		for _, listing := range allListings {
			log.Printf("Got listing with ID '%v', price '%d', quantity '%d'\n", listing.Id, listing.Price, listing.Quantity)
			log.Printf("Listing asset metadata: %v\n", listing.TokenMetadata)
		}
	},
}

var marketplaceListCmd = &cobra.Command{
	Use: "list [tokenId] [quantity] [price] [tokensPerBuyer] [secondsUntilOpenStart] [secondsUntilOpenEnd]",
	Short: "Lists an asset in the marketplace",
	Args: cobra.ExactArgs(6),
	ValidArgs: []string{
		"tokenId",
		"quantity",
		"price",
		"tokensPerBuyer",
		"secondsUntilOpenStart",
		"secondsUntilOpenEnd",
	},
	Run: func(cmd *cobra.Command, args []string) {
		if listingId, err := strconv.ParseInt(args[0], 10, 64); err != nil {
			panic(err)
		} else {
			newMarketplaceListing.TokenId = big.NewInt(listingId)
		}

		if quantity, err := strconv.ParseInt(args[1], 10, 64); err != nil {
			panic(err)
		} else {
			newMarketplaceListing.Quantity = big.NewInt(quantity)
		}

		if price, err := strconv.ParseInt(args[2], 10, 64); err != nil {
			panic(err)
		} else {
			newMarketplaceListing.Price = big.NewInt(price)
		}

		if tokensPerBuyer, err := strconv.ParseInt(args[3], 10, 64); err != nil {
			panic(err)
		} else {
			newMarketplaceListing.RewardsPerOpen = big.NewInt(tokensPerBuyer)
		}

		if secondsUntilOpenStart, err := strconv.ParseInt(args[4], 10, 64); err != nil {
			panic(err)
		} else {
			newMarketplaceListing.SecondsUntilOpenStart = big.NewInt(secondsUntilOpenStart)
		}

		if secondsUntilOpenEnd, err := strconv.ParseInt(args[5], 10, 64); err != nil {
			panic(err)
		} else {
			newMarketplaceListing.SecondsUntilOpenEnd = big.NewInt(secondsUntilOpenEnd)
		}

		module, err := getMarketplaceModule()
		if err != nil {
			panic(err)
		}

		newListing, err := module.List(newMarketplaceListing)
		if err != nil {
			panic(err)
		}

		jsonStr, err := json.Marshal(newListing)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonStr))
	},
}

func init() {
	marketplaceListCmd.PersistentFlags().StringVar(&newMarketplaceListing.AssetContractAddress, newListingAssetContractAddressFlag, "", "contract address of the asset to be listed")
	marketplaceListCmd.PersistentFlags().StringVar(&newMarketplaceListing.CurrencyContractAddress, newListingCurrencyContractAddressFlag, "", "contract address of the currency to use for the listing")
	_ = marketplaceListCmd.MarkPersistentFlagRequired(newListingAssetContractAddressFlag)
	_ = marketplaceListCmd.MarkPersistentFlagRequired(newListingCurrencyContractAddressFlag)

	marketplaceCmd.AddCommand(marketplaceGetAllCmd)
	marketplaceCmd.AddCommand(marketplaceListCmd)

	marketplaceCmd.PersistentFlags().StringVarP(&marketplaceContractAddress, "address", "a", "", "marketplace contract address")
}
