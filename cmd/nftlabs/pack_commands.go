package main

import (
	"github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"github.com/spf13/cobra"
	"log"
	"math/big"
	"strconv"
)

const (
	assetContractFlag = "assetContract"
	secondsUntilOpenStartFlag = "secondsUntilOpenStart"
	secondsUntilOpenEndFlag = "secondsUntilOpenEnd"
	rewardsPerOpenFlag = "rewardsPerOpen"
	assetSuppliesFlag = "assetSupplies"
	//descriptionFlag = "description"
	//imageFlag = "image"
	//externalUrlFlag = "externalUrl"
	//sellerFeeBasisPointsFlag = "sellerFeeBasisPoints"
	//feeRecipientFlag = "feeRecipient"
	//backgroundColorFlag = "backgroundColor"
)

var (
	packContractAddress string
	createPackArgs nftlabs.CreatePackArgs
	secondsUntilOpenStart int64
	secondsUntilOpenEnd int64
	rewardsPerOpen int64
	assetPairs map[string]int64
)

var packCmd = &cobra.Command{
	Use:   "pack [command]",
	Short: "Interact with a pack contract",
	Args: cobra.MinimumNArgs(1),
}

var packGetAllCmd = &cobra.Command {
	Use: "getAll [tokenId]",
	Short: "Get all available packs in a contract `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getPackModule()
		if err != nil {
			panic(err)
		}

		allPacks, err := module.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("Recieved %d packs\n", len(allPacks))
		for _, nft := range allPacks {
			log.Printf("Got pack with name '%v' and description '%v' and id '%d'\n", nft.Name, nft.Description, nft.Id)
		}
	},
}

var packCreateCmd = &cobra.Command {
	Use: "create",
	Short: "Create a new pack",
	Run: func(cmd *cobra.Command, args []string) {
		module, err := getPackModule()
		if err != nil {
			panic(err)
		}


		assets := make([]nftlabs.PackNftAddition, len(assetPairs))
		cnt := 0
		for tokenId, quantity := range assetPairs {
			tokenIdParse, err := strconv.ParseInt(tokenId, 10, 64)
			if err != nil {
				// TODO: return better error
				panic(err)
			}
			assets[cnt] = nftlabs.PackNftAddition{
				NftId: big.NewInt(tokenIdParse),
				Supply: big.NewInt(quantity),
			}
			cnt += 1
		}

		createPackArgs.Assets = assets
		createPackArgs.SecondsUntilOpenStart = big.NewInt(secondsUntilOpenStart)
		createPackArgs.SecondsUntilOpenEnd = big.NewInt(secondsUntilOpenEnd)
		createPackArgs.RewardsPerOpen = big.NewInt(rewardsPerOpen)
		log.Printf("Args = %v\n", createPackArgs)
		if nft, err := module.Create(createPackArgs); err != nil {
			panic(err)
		} else {
			log.Printf("Minted nft with info = %v\n", nft)
		}
	},
}

func init() {
	packCreateCmd.PersistentFlags().StringVar(&createPackArgs.AssetContractAddress, assetContractFlag, "", "address of the asset going in the pack")
	packCreateCmd.PersistentFlags().Int64Var(&secondsUntilOpenStart, secondsUntilOpenStartFlag, 0, "seconds until open start (default 0)")
	packCreateCmd.PersistentFlags().Int64Var(&secondsUntilOpenEnd, secondsUntilOpenEndFlag, 0, "seconds until open end (default 0)")
	packCreateCmd.PersistentFlags().Int64Var(&rewardsPerOpen, rewardsPerOpenFlag, 0, "rewards per open")
	packCreateCmd.PersistentFlags().StringToInt64Var(&assetPairs, assetSuppliesFlag, map[string]int64{}, "pairs of (tokenid, quantity) to transfer to the pack")
	_ = packCreateCmd.MarkPersistentFlagRequired(rewardsPerOpenFlag)
	_ = packCreateCmd.MarkPersistentFlagRequired(assetContractFlag)
	_ = packCreateCmd.MarkPersistentFlagRequired(assetSuppliesFlag)
	packCmd.AddCommand(packCreateCmd)

	packCmd.AddCommand(packGetAllCmd)
	packCmd.PersistentFlags().StringVarP(&packContractAddress, "address", "a", "", "pack contract address")
}
