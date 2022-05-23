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
	nftDropContractAddress string
)

var nftDropCmd = &cobra.Command{
	Use:   "nftdrop [command]",
	Short: "Interact with an NFT Drop contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var nftDropGetAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all available nfts in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		allNfts, err := nftDrop.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got drop nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var nftDropClaimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Claim an nft",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		if tx, err := nftDrop.Claim(1); err != nil {
			panic(err)
		} else {
			log.Printf("Claimed nft successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

var nftDropCreateBatchCmd = &cobra.Command{
	Use:   "createBatch",
	Short: "Create a batch of nfts",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		image0, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer image0.Close()

		image1, err := os.Open("internal/test/1.jpg")
		if err != nil {
			panic(err)
		}
		defer image1.Close()

		if tx, err := nftDrop.CreateBatch(
			[]*thirdweb.NFTMetadataInput{
				{
					Name:  "Drop NFT 1",
					Image: image0,
				},
				{
					Name:  "Drop NFT 2",
					Image: image1,
				},
				{
					Name:  "Drop NFT 3",
					Image: "ipfs://QmcCJC4T37rykDjR6oorM8hpB9GQWHKWbAi2YR1uTabUZu/0",
				},
				{
					Name:  "Drop NFT 4",
					Image: "ipfs://QmRCGCu9uyo2deiTFRUc5aMFB6AYUapCCxvF4QLUJbK474/0",
				},
			},
		); err != nil {
			panic(err)
		} else {
			log.Printf("Created batch of nfts successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

func init() {
	nftDropCmd.PersistentFlags().StringVarP(&nftDropContractAddress, "address", "a", "", "nft drop contract address")
	nftDropCmd.AddCommand(nftDropGetAllCmd)
	nftDropCmd.AddCommand(nftDropClaimCmd)
	nftDropCmd.AddCommand(nftDropCreateBatchCmd)
}
