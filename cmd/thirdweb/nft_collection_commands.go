package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

var (
	nftContractAddress string
)

var nftCmd = &cobra.Command{
	Use:   "nft [command]",
	Short: "Interact with an nft contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var nftGetAllCmd = &cobra.Command{
	Use:   "getAll",
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

var nftGetOwnedCmd = &cobra.Command{
	Use:   "getOwned",
	Short: "Get owned nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		allNfts, err := nftCollection.GetOwned("")
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var nftMintCmd = &cobra.Command{
	Use:   "mint",
	Short: "Get all available nfts in a contract",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		if tx, err := nftCollection.Mint(&thirdweb.NFTMetadataInput{
			Name:  "NFT Test",
			Image: imageFile,
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

var nftMintLinkCmd = &cobra.Command{
	Use:   "mintLink",
	Short: "Get all available nfts in a contract",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		if tx, err := nftCollection.Mint(&thirdweb.NFTMetadataInput{
			Name:  "NFT Test",
			Image: "ipfs://QmcCJC4T37rykDjR6oorM8hpB9GQWHKWbAi2YR1uTabUZu/0",
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
	nftCmd.PersistentFlags().StringVarP(&nftContractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
	nftCmd.AddCommand(nftGetOwnedCmd)
	nftCmd.AddCommand(nftMintCmd)
	nftCmd.AddCommand(nftMintLinkCmd)
}
