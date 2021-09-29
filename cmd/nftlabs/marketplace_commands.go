package main

import (
	"encoding/json"
	"fmt"
	"github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"math/big"
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
		jsonStr, err := json.Marshal(allListings)
		if err != nil {
			panic(err)
		}
		println(string(jsonStr))
	},
}

var marketplaceGetCmd = &cobra.Command{
	Use: "get [listingId]",
	Short: "Gets listing [listingId] from marketplace",
	Args: cobra.ExactArgs(1),
	ValidArgs: []string{"listingId"},
	Run: func(cmd *cobra.Command, args []string) {
		listingId := big.NewInt(0)
		if _, success := listingId.SetString(args[0], 10); !success {
			panic("Failed to parse ID")
		}
		newMarketplaceListing.TokenId = listingId

		module, err := getMarketplaceModule()
		if err != nil {
			panic(err)
		}

		listing, err := module.GetListing(listingId)
		if err != nil {
			panic(err)
		}
		jsonStr, err := json.Marshal(listing)
		if err != nil {
			panic(err)
		}
		println(string(jsonStr))
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
		listingId := big.NewInt(0)
		if _, success := listingId.SetString(args[0], 10); !success {
			panic("Failed to parse ID")
		}
		newMarketplaceListing.TokenId = listingId

		quantity := big.NewInt(0)
		if _, success := quantity.SetString(args[1], 10); !success {
			panic("Failed to parse quantity")
		}
		newMarketplaceListing.Quantity = quantity

		price := big.NewInt(0)
		if _, success := price.SetString(args[2], 10); !success {
			panic("Failed to parse price")
		}
		newMarketplaceListing.Price = price

		tokensPerBuyer := big.NewInt(0)
		if _, success := tokensPerBuyer.SetString(args[3], 10); !success {
			panic("Failed to parse tokensPerBuyer")
		}
		newMarketplaceListing.RewardsPerOpen = tokensPerBuyer

		secondsUntilOpenStart := big.NewInt(0)
		if _, success := secondsUntilOpenStart.SetString(args[4], 10); !success {
			panic("Failed to parse secondsUntilOpenStart")
		}
		newMarketplaceListing.SecondsUntilOpenStart = secondsUntilOpenStart

		secondsUntilOpenEnd := big.NewInt(0)
		if _, success := secondsUntilOpenEnd.SetString(args[5], 10); !success {
			panic("Failed to parse secondsUntilOpenEnd")
		}
		newMarketplaceListing.SecondsUntilOpenEnd = secondsUntilOpenEnd

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
	marketplaceCmd.AddCommand(marketplaceGetCmd)

	marketplaceCmd.PersistentFlags().StringVarP(&marketplaceContractAddress, "address", "a", "", "marketplace contract address")
}
