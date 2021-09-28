package main

import (
	"github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"log"
)

var (
	marketplaceContractAddress string
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

func init() {
	marketplaceCmd.AddCommand(marketplaceGetAllCmd)

	marketplaceCmd.PersistentFlags().StringVarP(&marketplaceContractAddress, "address", "a", "", "marketplace contract address")
}