package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var (
	marketplaceAddress string
)

var marketplaceCmd = &cobra.Command{
	Use:   "marketplace [command]",
	Short: "Interact with a marketplace contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var marketplaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "Mint tokens on `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		marketplace, err := getMarketplace()
		if err != nil {
			panic(err)
		}

		tx, err := marketplace.CreateListing(&thirdweb.NewDirectListing{})
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

func init() {
	marketplaceCmd.PersistentFlags().StringVarP(&marketplaceAddress, "address", "a", "", "marketplace contract address")
	marketplaceCmd.AddCommand(marketplaceListCmd)
}
