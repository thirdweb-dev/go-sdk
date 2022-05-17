package main

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	nftDropContractAddress string
)

var nftDropCmd = &cobra.Command{
	Use:   "nftdrop [command]",
	Short: "Interact with an NFT Drop contract",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var nftDropGetAllCmd = &cobra.Command {
	Use: "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftDrop()
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
	nftDropCmd.PersistentFlags().StringVarP(&nftDropContractAddress, "address", "a", "", "nft contract address")
	nftDropCmd.AddCommand(nftDropGetAllCmd)
}
