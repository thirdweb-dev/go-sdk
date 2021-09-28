package main

import (
	"errors"
	"github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
)

const (
	createCollectionSupplyFlag = "supply"
)

var (
	collectionContractAddress string
	createCollectionArgs nftlabs.CreateCollectionArgs
	createCollectionSupply int64
	createCollectionNftMetadata nftlabs.MintNftMetadata
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
			log.Printf("Got nft with name '%v' and description '%v' and id '%d', supply '%v'\n", nft.Name, nft.Description, nft.Id, nft.Supply)
		}
	},
}

var collectionCreateCmd = &cobra.Command {
	Use: "create [name] [supply]",
	Short: "Create a collection with a single asset in the contract `ADDRESS`",
	ValidArgs: []string{
		createCollectionSupplyFlag,
	},
	ArgAliases: []string{
		"s",
	},
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		supply, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			panic(errors.New("supply is not a valid number"))
		}

		module, err := getCollectionModule()
		if err != nil {
			panic(err)
		}

		createCollectionNftMetadata.Name = name
		createCollectionArgs.Supply = big.NewInt(supply)
		createCollectionArgs.Metadata = createCollectionNftMetadata

		if result, err := module.Create([]nftlabs.CreateCollectionArgs{createCollectionArgs}); err != nil {
			panic(err)
		} else {
			log.Printf("Created collection with %d nfts \n", len(result))

			for _, nft := range result {
				log.Printf("Nft in collection with ID %d, name %v \n", nft.Id, nft.NftMetadata.Name)
			}
		}
	},
}

func init() {
	collectionCreateCmd.PersistentFlags().StringVar(&createCollectionNftMetadata.Description, descriptionFlag, "", "description for nft")
	collectionCreateCmd.PersistentFlags().StringVar(&createCollectionNftMetadata.ExternalUrl, externalUrlFlag, "", "external url for nft")
	collectionCreateCmd.PersistentFlags().StringVar(&createCollectionNftMetadata.Image, imageFlag, "", "image for nft")
	collectionCreateCmd.PersistentFlags().StringVar(&createCollectionNftMetadata.FeeRecipient, feeRecipientFlag, "", "fee recipient for nft royalties")
	collectionCreateCmd.PersistentFlags().StringVar(&createCollectionNftMetadata.BackgroundColor, backgroundColorFlag, "", "hex value for background color")
	collectionCreateCmd.PersistentFlags().Int64Var(&sellerFeeBasisPoints, sellerFeeBasisPointsFlag, 0, "basis points to collect (to feeRecipient) on each sale")

	collectionCreateCmd.PersistentFlags().Int64Var(&createCollectionSupply, createCollectionSupplyFlag,  0, "number of supply for the collection to create")

	collectionCmd.AddCommand(collectionCreateCmd)

	collectionCmd.PersistentFlags().StringVarP(&collectionContractAddress, "address", "a", "", "collection contract address")
	collectionCmd.AddCommand(collectionGetAllCmd)
}
