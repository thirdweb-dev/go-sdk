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

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		if tx, err := edition.Mint(&thirdweb.EditionMetadataInput{
			Metadata: &thirdweb.NFTMetadataInput{
				Name:  "Edition Test",
				Image: imageFile,
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

var editionSigmintCmd = &cobra.Command{
	Use:   "sigmint",
	Short: "Sign and mint an nft",
	Run: func(cmd *cobra.Command, args []string) {
		edition, err := getEdition()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/1.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		payload, err := edition.Signature.Generate(
			&thirdweb.Signature1155PayloadInput{
				To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
				Price:                0,
				CurrencyAddress:      "0x0000000000000000000000000000000000000000",
				MintStartTime:        0,
				MintEndTime:          100000000000000,
				PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
				Metadata: &thirdweb.NFTMetadataInput{
					Name:  "ERC1155 Sigmint!",
					Image: imageFile,
				},
				RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
				RoyaltyBps:       0,
				Quantity:         1,
			},
		)
		if err != nil {
			panic(err)
		}

		valid, err := edition.Signature.Verify(payload)
		if err != nil {
			panic(err)
		} else if !valid {
			panic("Invalid signature")
		}

		tx, err := edition.Signature.Mint(payload)
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

var editionSigmintTokenIdCmd = &cobra.Command{
	Use:   "sigmint-tokenid",
	Short: "Sign and mint an nft",
	Run: func(cmd *cobra.Command, args []string) {
		edition, err := getEdition()
		if err != nil {
			panic(err)
		}

		imageFile, err := os.Open("internal/test/1.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		payload, err := edition.Signature.GenerateFromTokenId(
			&thirdweb.Signature1155PayloadInputWithTokenId{
				To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
				Price:                0,
				CurrencyAddress:      "0x0000000000000000000000000000000000000000",
				MintStartTime:        0,
				MintEndTime:          100000000000000,
				PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
				Metadata: &thirdweb.NFTMetadataInput{
					Name:  "ERC1155 Sigmint!",
					Image: imageFile,
				},
				RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
				RoyaltyBps:       0,
				Quantity:         11,
				TokenId:          0,
			},
		)
		if err != nil {
			panic(err)
		}

		valid, err := edition.Signature.Verify(payload)
		if err != nil {
			panic(err)
		} else if !valid {
			panic("Invalid signature")
		}

		tx, err := edition.Signature.Mint(payload)
		if err != nil {
			panic(err)
		}

		result, _ := json.Marshal(&tx)
		fmt.Println(string(result))
	},
}

func init() {
	editionCmd.PersistentFlags().StringVarP(&editionAddress, "address", "a", "", "edition contract address")
	editionCmd.AddCommand(editionGetAllCmd)
	editionCmd.AddCommand(editionGetOwnedCmd)
	editionCmd.AddCommand(editionMintCmd)
	editionCmd.AddCommand(editionSigmintCmd)
	editionCmd.AddCommand(editionSigmintTokenIdCmd)
}
