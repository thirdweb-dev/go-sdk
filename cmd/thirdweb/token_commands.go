package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var (
	tokenAddress string
)

var tokenCmd = &cobra.Command{
	Use:   "token [command]",
	Short: "Interact with a token contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var tokenGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get token data and balance on `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := getToken()
		if err != nil {
			panic(err)
		}

		data, err := token.Get()
		if err != nil {
			panic(err)
		}

		log.Println("Name: ", data.Name)
		log.Println("Symbols: ", data.Symbol)
		log.Println("Decimals: ", data.Decimals)

		balance, err := token.Balance()
		if err != nil {
			panic(err)
		}

		log.Println("Balance: ", balance.DisplayValue)
	},
}

var tokenMintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Mint tokens on `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := getToken()
		if err != nil {
			panic(err)
		}

		tx, err := token.Mint(1)
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

var tokenMintBatchCmd = &cobra.Command{
	Use:   "mintBatch",
	Short: "Mint tokens on `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := getToken()
		if err != nil {
			panic(err)
		}

		tx, err := token.MintBatchTo([]*thirdweb.TokenAmount{
			{
				ToAddress: "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
				Amount:    1,
			},
			{
				ToAddress: "0x0f14090Dc6BB0Eb36E6d386176047cbC6BF1D077",
				Amount:    1,
			},
		})
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

func init() {
	tokenCmd.PersistentFlags().StringVarP(&tokenAddress, "address", "a", "", "token contract address")
	tokenCmd.AddCommand(tokenGetCmd)
	tokenCmd.AddCommand(tokenMintCmd)
	tokenCmd.AddCommand(tokenMintBatchCmd)
}
