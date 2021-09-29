package main

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	currencyContractAddress string
)

var currencyCmd = &cobra.Command{
	Use:   "currency [command]",
	Short: "Interact with a currency contract",
}

var currencyGetCmd = &cobra.Command{
	Use: "get",
	Short: "Gets the currency at the specified address",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getCurrencyModule()
		if err != nil {
			panic(err)
		}

		currency, err := module.Get()
		if err != nil {
			panic(err)
		}
		log.Printf("Got currency  %v\n", currency)
	},
}

func init() {
	currencyCmd.AddCommand(currencyGetCmd)

	currencyCmd.PersistentFlags().StringVarP(&currencyContractAddress, "address", "a", "", "currency contract address")
}