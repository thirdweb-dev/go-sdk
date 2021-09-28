package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	sdk "github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"log"
)

var (
	contractAddress string
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
		client, err := ethclient.Dial(chainRpcUrl)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Creating an NFT sdk module on chain %v, contract %v\n", chainRpcUrl, contractAddress)

		// You can mock the sdk.NftSdk interface when writing tests against the SDK
		var module sdk.NftSdk
		if caller, err := sdk.NewNftSdkModule(client, contractAddress, &sdk.SdkOptions{}); err != nil {
			log.Println("Failed to create an NFT caller object")
			panic(err)
		} else {
			module = caller
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

func init() {
	nftCmd.PersistentFlags().StringVarP(&contractAddress, "address", "a", "", "nft contract address")
	nftCmd.AddCommand(nftGetAllCmd)
}
