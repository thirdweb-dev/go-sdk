package main

import (
	"fmt"
	"github.com/nftlabs/nftlabs-sdk-go/pkg/nftlabs"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
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

var currencyMintCmd = &cobra.Command{
	Use: "mint [amount]",
	Short: "Mints [amount] of currency at the specified address",
	Args: cobra.ExactArgs(1),
	ValidArgs: []string{"amount"},
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}

		module, err := getCurrencyModule()
		if err != nil {
			panic(err)
		}

		currency, err := module.Get()
		if err != nil {
			panic(err)
		}

		if err := module.Mint(big.NewInt(amount)); err != nil {
			panic(err)
		}
		fmt.Printf("Minted %d %v\n", amount, currency.Symbol)
	},
}

var currencyTotalSupplyCmd = &cobra.Command{
	Use: "totalSupply",
	Short: "Gets the total supply of the currency at the specified address",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getCurrencyModule()
		if err != nil {
			panic(err)
		}

		totalSupply, err := module.TotalSupply()
		if err != nil {
			panic(err)
		}

		fmt.Println(totalSupply)
	},
}

var currencyBurnCmd = &cobra.Command{
	Use: "burn [amount]",
	Short: "Burns [amount] of the currency at the specified address",
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}

		module, err := getCurrencyModule()
		if err != nil {
			panic(err)
		}

		if err := module.Burn(big.NewInt(amount)); err != nil {
			panic(err)
		}

		fmt.Printf("%d burned successfully\n", amount)
	},
}

var currencyGrantRoleCmd = &cobra.Command{
	Use: "grantRole [role] [address]",
	Short: "Grants the role `role` to the wallet at `address`",
	Args: cobra.ExactArgs(2),
	ArgAliases: []string{"role, address"},
	Example: "grantRole admin|minter `address`",
	Run: func(cmd *cobra.Command, args []string) {
		role := args[0]
		address := args[1]

		module, err := getCurrencyModule()
		if err != nil {
			panic(err)
		}

		if r, err := nftlabs.RoleFromString(role); err != nil {
			panic(err)
		} else {
			if err := module.GrantRole(r, address); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	currencyCmd.AddCommand(currencyTotalSupplyCmd)
	currencyCmd.AddCommand(currencyMintCmd)
	currencyCmd.AddCommand(currencyGetCmd)
	currencyCmd.AddCommand(currencyBurnCmd)

	currencyCmd.AddCommand(currencyGrantRoleCmd)

	currencyCmd.PersistentFlags().StringVarP(&currencyContractAddress, "address", "a", "", "currency contract address")
}
