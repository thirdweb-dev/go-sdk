package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	customContractAddress string
)

var customCmd = &cobra.Command{
	Use:   "custom [command]",
	Short: "Interact with an custom contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var customSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set custom contract profile in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		contract, err := getCustom()
		if err != nil {
			panic(err)
		}

		data, err := contract.Call(context.Background(), "tokenURI", 0)
		if err != nil {
			panic(err)
		}

		log.Println(data)

		tx, err := contract.Call(context.Background(), "mintTo", thirdwebSDK.GetSignerAddress().Hex(), "ipfs://QmXCKX8MHHCXU62UWdiU38cjK3vbQ4MeL9FgXKo2hAR6Yz/0")
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

func init() {
	customCmd.PersistentFlags().StringVarP(&customContractAddress, "address", "a", "", "nft contract address")
	customCmd.AddCommand(customSetCmd)
}
