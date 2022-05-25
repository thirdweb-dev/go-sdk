package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
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

var nftSigmintCmd = &cobra.Command{
	Use:   "sigmint",
	Short: "Sign and mint an nft",
	Run: func(cmd *cobra.Command, args []string) {
		nftCollection, err := getNftCollection()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/1.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		payload, err := nftCollection.Signature.Generate(
			&thirdweb.Signature721PayloadInput{
				To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
				Price:                0,
				CurrencyAddress:      "0x0000000000000000000000000000000000000000",
				MintStartTime:        0,
				MintEndTime:          100000000000000,
				PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
				Metadata: &thirdweb.NFTMetadataInput{
					Name:  "ERC721 Sigmint!",
					Image: imageFile,
				},
				RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
				RoyaltyBps:       0,
			},
		)
		if err != nil {
			panic(err)
		}

		valid, err := nftCollection.Signature.Verify(payload)
		if err != nil {
			panic(err)
		} else if !valid {
			panic("Invalid signature")
		}

		tx, err := nftCollection.Signature.Mint(payload)
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

func init() {
	nftCmd.PersistentFlags().StringVarP(&nftContractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
	nftCmd.AddCommand(nftGetOwnedCmd)
	nftCmd.AddCommand(nftMintCmd)
	nftCmd.AddCommand(nftMintLinkCmd)
	nftCmd.AddCommand(nftSigmintCmd)
}
