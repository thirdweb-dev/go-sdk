package main

import (
	"encoding/json"
	"fmt"
	"github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"log"
	"math/big"
)

const (
	nameFlag = "name"
	descriptionFlag = "description"
	imageFlag = "image"
	externalUrlFlag = "externalUrl"
	sellerFeeBasisPointsFlag = "sellerFeeBasisPoints"
	feeRecipientFlag = "feeRecipient"
	backgroundColorFlag = "backgroundColor"
)

var (
	contractAddress string
	nftMetadata nftlabs.MintNftMetadata
	sellerFeeBasisPoints int64
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
	Use: "getAll [tokenId]",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getNftModule()
		if err != nil {
			panic(err)
		}

		allNfts, err := module.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Name, nft.Description, nft.Id)
		}
	},
}

var nftMintCmd = &cobra.Command {
	Use: "mint [tokenId]",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getNftModule()
		if err != nil {
			panic(err)
		}
		if err := module.SetPrivateKey(privateKey); err != nil {
			panic(err)
		}

		nftMetadata.SellerFeeBasisPoints = big.NewInt(sellerFeeBasisPoints)
		if nft, err := module.Mint(nftMetadata); err != nil {
			panic(err)
		} else {
			log.Printf("Minted nft with info = %v\n", nft)

			result, _ := json.Marshal(&nft)
			fmt.Println(string(result))
		}
	},
}

func init() {
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.Name, nameFlag, "", "name for nft")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.Description, descriptionFlag, "", "description for nft")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.ExternalUrl, externalUrlFlag, "", "external url for nft")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.Image, imageFlag, "", "image for nft")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.FeeRecipient, feeRecipientFlag, "", "fee recipient for nft royalties")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.BackgroundColor, backgroundColorFlag, "", "hex value for background color")
	nftMintCmd.PersistentFlags().Int64Var(&sellerFeeBasisPoints, sellerFeeBasisPointsFlag, 0, "basis points to collect (to feeRecipient) on each sale")
	_ = nftMintCmd.MarkPersistentFlagRequired(nameFlag)
	nftCmd.AddCommand(nftMintCmd)

	nftCmd.PersistentFlags().StringVarP(&contractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
}
