package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [command]",
	Short: "Deploy a contract",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var deployNftCmd = &cobra.Command{
	Use:   "nft",
	Short: "Deploy an nft collection",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployNFTCollection(&thirdweb.DeployNFTCollectionMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionCmd = &cobra.Command{
	Use:   "edition",
	Short: "Deploy an edition",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployEdition(&thirdweb.DeployEditionMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Deploy an token",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployToken(&thirdweb.DeployTokenMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployNFTDropCmd = &cobra.Command{
	Use:   "nftdrop",
	Short: "Deploy an nft drop",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployNFTDrop(&thirdweb.DeployNFTDropMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployEditionDropCmd = &cobra.Command{
	Use:   "editiondrop",
	Short: "Deploy an edition drop",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployEditionDrop(&thirdweb.DeployEditionDropMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMultiwrapCmd = &cobra.Command{
	Use:   "multiwrap",
	Short: "Deploy a multiwrap",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployMultiwrap(&thirdweb.DeployMultiwrapMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

var deployMarketplaceCmd = &cobra.Command{
	Use:   "marketplace",
	Short: "Deploy a marketplace",
	Run: func(cmd *cobra.Command, args []string) {
		if thirdwebSDK == nil {
			initSdk()
		}

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		address, err := thirdwebSDK.Deployer.DeployMarketplace(&thirdweb.DeployMarketplaceMetadata{
			Name: "Go SDK",
		})
		if err != nil {
			panic(err)
		}

		log.Println("Address:")
		log.Println(address)
	},
}

func init() {
	deployCmd.AddCommand(deployNftCmd)
	deployCmd.AddCommand(deployEditionCmd)
	deployCmd.AddCommand(deployTokenCmd)
	deployCmd.AddCommand(deployNFTDropCmd)
	deployCmd.AddCommand(deployEditionDropCmd)
	deployCmd.AddCommand(deployMultiwrapCmd)
	deployCmd.AddCommand(deployMarketplaceCmd)
}
