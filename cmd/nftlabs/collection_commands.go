package main

import (
	"fmt"
	"github.com/nftlabs/nftlabs-sdk-go/pkg/nftlabs"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
)

const (
	createCollectionSupplyFlag = "supply"
	nftMetadataFlag = "nftMetadata"
)

var (
	collectionContractAddress string
	createCollectionArgs nftlabs.CreateCollectionArgs
	createCollectionSupply int64
	createCollectionNftMetadata string
)

var collectionCmd = &cobra.Command{
	Use:   "collection [command]",
	Short: "Interact with a collection contract",
}

var collectionGetAllCmd = &cobra.Command {
	Use: "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Example:"nftlabs -k $PKEY collection -a $COLLECTION_ADDRESS getAll",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getCollectionModule()
		if err != nil {
			panic(err)
		}

		allNfts, err := module.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts in collection %v\n", len(allNfts), collectionContractAddress)
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d', supply '%v'\n", nft.Name, nft.Description, nft.Id, nft.Supply)
		}
	},
}

var collectionBalanceCmd = &cobra.Command {
	Use: "balance [tokenId]",
	Short: "Check your balance of [tokenId]",
	ValidArgs: []string{"tokenId"},
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tokenId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic(err)
		}

		module, err := getCollectionModule()
		if err != nil {
			panic(err)
		}
		balance, err := module.Balance(big.NewInt(tokenId))
		if err != nil {
			panic(err)
		}
		fmt.Println(balance)
	},
}

func init() {
	collectionCmd.PersistentFlags().StringVarP(&collectionContractAddress, "address", "a", "", "collection contract address")
	collectionCmd.AddCommand(collectionGetAllCmd)

	collectionCmd.AddCommand(collectionBalanceCmd)
}
