package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

var (
	editionAddress string
)

var editionCmd = &cobra.Command{
	Use:   "edition [command]",
	Short: "Interact with an edition contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var editionGetAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		edition, err := getEdition()
		if err != nil {
			panic(err)
		}

		allNfts, err := edition.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and supply '%v' and id '%d'\n", nft.Metadata.Name, nft.Supply, nft.Metadata.Id)
		}
	},
}

var editionGetOwnedCmd = &cobra.Command{
	Use:   "getOwned",
	Short: "Get owned nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		edition, err := getEdition()
		if err != nil {
			panic(err)
		}

		allNfts, err := edition.GetOwned("")
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and supply '%v' and id '%d'\n", nft.Metadata.Name, nft.Supply, nft.Metadata.Id)
		}
	},
}

var editionMintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Get all available nfts in a contract",
	Run: func(cmd *cobra.Command, args []string) {
		edition, err := getEdition()
		if err != nil {
			panic(err)
		}

		if tx, err := edition.Mint(&thirdweb.EditionMetadataInput{
			Metadata: &thirdweb.NFTMetadataInput{
				Name: "Edition Test",
			},
			Supply: 1,
		}); err != nil {
			panic(err)
		} else {
			// TODO return the minted token ID
			log.Printf("Minted nft successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

func init() {
	editionCmd.PersistentFlags().StringVarP(&editionAddress, "address", "a", "", "nft contract address")
	editionCmd.AddCommand(editionGetAllCmd)
	editionCmd.AddCommand(editionGetOwnedCmd)
	editionCmd.AddCommand(editionMintCmd)
}
