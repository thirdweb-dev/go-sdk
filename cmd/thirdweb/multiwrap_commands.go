package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var (
	multiwrapContractAddress string
	multiwrapNft             string
	multiwrapEdition         string
	multiwrapToken           string
)

var multiwrapCmd = &cobra.Command{
	Use:   "multiwrap [command]",
	Short: "Interact with a multiwrap contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var multiwrapGetAllCmd = &cobra.Command{
	Use:   "getAll",
	Short: "Get all tokens of `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		multiwrap, err := getMultiwrap()
		if err != nil {
			panic(err)
		}
		tokens, err := multiwrap.GetAll()
		if err != nil {
			panic(err)
		}
		log.Printf("%d tokens found", len(tokens))

		for _, token := range tokens {
			log.Println("ID: ", token.Metadata.Id)
			log.Println("Owner: ", token.Owner)
			log.Println("Name: ", token.Metadata.Name)
		}
	},
}

var multiwrapGetContentsCmd = &cobra.Command{
	Use:   "getContents",
	Short: "Get contents of a token `ADDRESS`",
	Run: func(cmd *cobra.Command, args []string) {
		multiwrap, err := getMultiwrap()
		if err != nil {
			panic(err)
		}
		contents, err := multiwrap.GetWrappedContents(1)
		if err != nil {
			panic(err)
		}
		log.Printf("%d erc20s", len(contents.ERC20Tokens))
		log.Printf("%d erc721s", len(contents.ERC721Tokens))
		log.Printf("%d erc1155s", len(contents.ERC1155Tokens))
	},
}

var multiwrapWrapCmd = &cobra.Command{
	Use:   "wrap",
	Short: "Wrap tokens",
	Run: func(cmd *cobra.Command, args []string) {
		multiwrap, err := getMultiwrap()
		if err != nil {
			panic(err)
		}

		contents := &thirdweb.MultiwrapBundle{
			ERC20Tokens: []*thirdweb.MultiwrapERC20{
				{
					ContractAddress: multiwrapToken,
					Quantity:        1,
				},
			},
			ERC721Tokens: []*thirdweb.MultiwrapERC721{
				{
					ContractAddress: multiwrapNft,
					TokenId:         5,
				},
			},
			ERC1155Tokens: []*thirdweb.MultiwrapERC1155{
				{
					ContractAddress: multiwrapEdition,
					TokenId:         0,
					Quantity:        1,
				},
			},
		}

		log.Println("Token: ", multiwrapToken)
		log.Println("NFT: ", multiwrapNft)
		log.Println("Edition: ", multiwrapEdition)

		if tx, err := multiwrap.Wrap(
			contents,
			&thirdweb.NFTMetadataInput{
				Name: "Wrapped Token",
			},
			"",
		); err != nil {
			panic(err)
		} else {
			log.Printf("Wrapped tokens successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

var multiwrapUnwrapCmd = &cobra.Command{
	Use:   "unwrap",
	Short: "Unwrap tokens",
	Run: func(cmd *cobra.Command, args []string) {
		multiwrap, err := getMultiwrap()
		if err != nil {
			panic(err)
		}

		if tx, err := multiwrap.Unwrap(3, ""); err != nil {
			panic(err)
		} else {
			log.Printf("Wrapped tokens successfully")

			result, _ := json.Marshal(&tx)
			fmt.Println(string(result))
		}
	},
}

func init() {
	multiwrapCmd.PersistentFlags().StringVarP(&multiwrapContractAddress, "address", "a", "", "multiwrap contract address")
	multiwrapCmd.PersistentFlags().StringVarP(&multiwrapNft, "nft", "n", "", "nft contract address")
	multiwrapCmd.PersistentFlags().StringVarP(&multiwrapEdition, "edition", "e", "", "edition contract address")
	multiwrapCmd.PersistentFlags().StringVarP(&multiwrapToken, "token", "t", "", "token contract address")
	multiwrapCmd.AddCommand(multiwrapWrapCmd)
	multiwrapCmd.AddCommand(multiwrapUnwrapCmd)
	multiwrapCmd.AddCommand(multiwrapGetContentsCmd)
	multiwrapCmd.AddCommand(multiwrapGetAllCmd)
}
