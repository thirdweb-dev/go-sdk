package main

import (
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

		tx, err := contract.Call("setProfile", "Akira", "Fudo")
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))

		data, err := contract.Call("getProfile", "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6")
		if err != nil {
			panic(err)
		}

		log.Println(data)
	},
}

func init() {
	customCmd.PersistentFlags().StringVarP(&customContractAddress, "address", "a", "", "nft contract address")
	customCmd.AddCommand(customSetCmd)
}
