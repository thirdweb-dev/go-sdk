package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	nftContractAddress string
)

var nftCmd = &cobra.Command{
	Use:   "nft [command]",
	Short: "Interact with an nft contract",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Hello from the inside nft")
	},
}

var nftGetAllCmd = &cobra.Command {
	Use: "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		allNfts, err := nftCollection.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

func init() {
	nftCmd.PersistentFlags().StringVarP(&nftContractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
}
