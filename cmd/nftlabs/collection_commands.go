package main

import (
	"github.com/spf13/cobra"
	"log"
)

var (
	collectionContractAddress string
)

var collectionCmd = &cobra.Command{
	Use:   "collection [command]",
	Short: "Interact with a collection contract",
	Args: cobra.MinimumNArgs(1),
}

var collectionGetAllCmd = &cobra.Command {
	Use: "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
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
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Name, nft.Description, nft.Id)
		}
	},
}

func init() {
	collectionCmd.PersistentFlags().StringVarP(&collectionContractAddress, "address", "a", "", "collection contract address")
	collectionCmd.AddCommand(collectionGetAllCmd)
}
