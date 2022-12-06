package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
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

		allNfts, err := nftDrop.GetAll(context.Background())
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d nfts\n", len(allNfts))
		for _, nft := range allNfts {
			log.Printf("Got drop nft with name '%v' and description '%v' and id '%d'\n", nft.Metadata.Name, nft.Metadata.Description, nft.Metadata.Id)
		}
	},
}

var nftDropEncoderCmd = &cobra.Command{
	Use:   "encoder",
	Short: "Get encoded claim transaction",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		// tx, err := nftDrop.Encoder.ApproveClaimTo(
		// 	context.Background(),
		// 	thirdwebSDK.GetSignerAddress().String(),
		// 	1,
		// )

		tx, err := nftDrop.Encoder.ClaimTo(
			context.Background(),
			thirdwebSDK.GetSignerAddress().String(),
			thirdwebSDK.GetSignerAddress().String(),
			1,
		)
		if err != nil {
			panic(err)
		}

		log.Println(tx.Data())
		log.Println(tx.To())
	},
}

var nftDropGetActiveCmd = &cobra.Command{
	Use:   "getActive",
	Short: "Get the active claim condition in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		nftDrop, err := getNftDrop()
		if err != nil {
			panic(err)
		}

		active, err := nftDrop.ClaimConditions.GetActive(context.Background())
		if err != nil {
			panic(err)
		}

		fmt.Println("Start Time:", active.StartTime)
		fmt.Println("Available:", active.AvailableSupply)
		fmt.Println("Quantity:", active.MaxClaimableSupply)
		fmt.Println("Quantity Limit:", active.MaxClaimablePerWallet)
		fmt.Println("Price:", active.Price)
		fmt.Println("Wait In Seconds", active.WaitInSeconds)

		all, err := nftDrop.ClaimConditions.GetAll(context.Background())
		if err != nil {
			panic(err)
		}

		for i, c := range all {
			fmt.Printf(fmt.Sprintf("\n\nClaim Condition %d\n================\n", i))
			fmt.Println("Start Time:", c.StartTime)
			fmt.Println("Available:", c.AvailableSupply)
			fmt.Println("Quantity:", c.MaxClaimableSupply)
			fmt.Println("Quantity Limit:", c.MaxClaimablePerWallet)
			fmt.Println("Price:", c.Price)
			fmt.Println("Wait In Seconds", c.WaitInSeconds)
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

		emptySdk, err := thirdweb.NewThirdwebSDK(
			chainRpcUrl,
			nil,
		)
		if err != nil {
			panic(err)
		}

		emptyDrop, err := emptySdk.GetNFTDrop(nftDropContractAddress)

		address := thirdwebSDK.GetSignerAddress().String()
		claimArgs, err := emptyDrop.GetClaimArguments(context.Background(), address, 1)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v\n", claimArgs)

		txOpts, err := nftDrop.Helper.GetTxOptions(context.Background())
		if err != nil {
			panic(err)
		}

		txOpts.Value = claimArgs.TxValue

		tx, err := nftDrop.Abi.Claim(
			txOpts,
			claimArgs.Receiver,
			claimArgs.Quantity,
			claimArgs.Currency,
			claimArgs.PricePerToken,
			claimArgs.AllowlistProof,
			claimArgs.Data,
		)
		if err != nil {
			panic(err)
		}

		awaitTx(tx.Hash())

		unclaimed, err := nftDrop.TotalUnclaimedSupply()
		if err != nil {
			panic(err)
		}

		fmt.Println(unclaimed)

		// tx, err := nftDrop.Claim(context.Background(), 1)
		// if err != nil {
		// 	panic(err)
		// }

		// log.Println("Claimed nft with tx hash", tx.Hash().String())
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
			context.Background(),
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
	nftDropCmd.AddCommand(nftDropEncoderCmd)
	nftDropCmd.AddCommand(nftDropGetActiveCmd)
	nftDropCmd.AddCommand(nftDropClaimCmd)
	nftDropCmd.AddCommand(nftDropCreateBatchCmd)
}
