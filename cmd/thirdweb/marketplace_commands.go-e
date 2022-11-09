package main

import (
	"context"
	"encoding/hex"
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

		tx, err := marketplace.CreateListing(context.Background(), &thirdweb.NewDirectListing{})
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

var marketplaceEncodeCancelCmd = &cobra.Command{
	Use:   "encode",
	Short: "Encode cancel listing `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		marketplace, err := getMarketplace()
		if err != nil {
			panic(err)
		}

		tx, err := marketplace.Encoder.CancelListing(context.Background(), "0x0000000000000000000000000000000000000000", 0)

		fmt.Println("Nonce:", tx.Nonce())
		fmt.Println("To:", tx.To())
		fmt.Println("GasLimit:", tx.Gas())
		fmt.Println("GasPrice:", tx.GasPrice())
		fmt.Println("Value:", tx.Value())
		fmt.Println("Data:", hex.EncodeToString(tx.Data()))
	},
}

var marketplaceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get listing `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		marketplace, err := getMarketplace()
		if err != nil {
			panic(err)
		}

		listing, err := marketplace.GetListing(context.Background(), 0)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v", listing)
		fmt.Printf("%#v", listing.Asset)
	},
}

func init() {
	marketplaceCmd.PersistentFlags().StringVarP(&marketplaceAddress, "address", "a", "", "marketplace contract address")
	marketplaceCmd.AddCommand(marketplaceListCmd)
	marketplaceCmd.AddCommand(marketplaceEncodeCancelCmd)
	marketplaceCmd.AddCommand(marketplaceGetCmd)
}
