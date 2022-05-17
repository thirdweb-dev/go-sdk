package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/nftlabs/nftlabs-sdk-go/pkg/nftlabs"
	"github.com/spf13/cobra"
)

const (
	nameFlag        = "name"
	descriptionFlag = "description"
	imageFlag       = "image"
	externalUrlFlag = "externalUrl"
	sellerFeeBasisPointsFlag = "sellerFeeBasisPoints"
	feeRecipientFlag = "feeRecipient"
	backgroundColorFlag = "backgroundColor"
)

var (
	nftContractAddress string
	nftMetadata        nftlabs.MintNftMetadata
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
	Use: "getAll",
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
	Use: "mint",
	Short: "Get all available nfts in a contract",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getNftModule()
		if err != nil {
			panic(err)
		}

		if nft, err := module.Mint(nftMetadata); err != nil {
			panic(err)
		} else {
			log.Printf("Minted nft with info = %v\n", nft)

			result, _ := json.Marshal(&nft)
			fmt.Println(string(result))
		}
	},
}

var nftGetOwnedCmd = &cobra.Command {
	Use: "getOwned [address]",
	Short: "Get all nfts owned by wallet `address`",
	Args: cobra.ExactArgs(1),
	ValidArgs: []string{"address"},
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getNftModule()
		if err != nil {
			panic(err)
		}

		allNfts, err := module.GetOwned(args[0])
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got nft with name '%v' and description '%v' and id '%d'\n", nft.Name, nft.Description, nft.Id)
		}
	},
}

func init() {
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.Name, nameFlag, "", "name for nft")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.Description, descriptionFlag, "", "description for nft")
	nftMintCmd.PersistentFlags().StringVar(&nftMetadata.Image, imageFlag, "", "image for nft")
	nftMintCmd.PersistentFlags().Int64Var(&sellerFeeBasisPoints, sellerFeeBasisPointsFlag, 0, "basis points to collect (to feeRecipient) on each sale")
	_ = nftMintCmd.MarkPersistentFlagRequired(nameFlag)
	nftCmd.AddCommand(nftMintCmd)
	nftCmd.AddCommand(nftGetOwnedCmd)

	nftCmd.PersistentFlags().StringVarP(&nftContractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
}
