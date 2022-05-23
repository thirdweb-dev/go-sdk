package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

var storageCmd = &cobra.Command{
	Use:   "storage [command]",
	Short: "Interact with the IPFS storage interface",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please input a command to run")
	},
}

var storageUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload data with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		uri, err := storage.Upload(&thirdweb.NFTMetadataInput{
			Name:        "Test NFT 1",
			Description: "Description 1",
		}, "", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uri)
	},
}

var storageUploadBatchCmd = &cobra.Command{
	Use:   "uploadBatch",
	Short: "Upload data with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		uriWithBaseUris, err := storage.UploadBatch(
			[]interface{}{
				&thirdweb.NFTMetadataInput{Name: "Test NFT 2", Description: "Description 2"},
				&thirdweb.NFTMetadataInput{Name: "Test NFT 3", Description: "Description 3"},
			},
			0,
			"",
			"",
		)

		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uriWithBaseUris)
	},
}

var storageUploadImageCmd = &cobra.Command{
	Use:   "uploadImage",
	Short: "Upload image with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		uri, err := storage.Upload(&thirdweb.NFTMetadataInput{
			Name:  "Test NFT 1",
			Image: imageFile,
		}, "", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uri)
	},
}

var storageUploadImageLinkCmd = &cobra.Command{
	Use:   "uploadImageLink",
	Short: "Upload image with link with storage interface",
	Run: func(cmd *cobra.Command, args []string) {
		storage := getStorage()

		imageFile, err := os.Open("internal/test/0.jpg")
		if err != nil {
			panic(err)
		}
		defer imageFile.Close()

		uri, err := storage.Upload(&thirdweb.NFTMetadataInput{
			Name:  "Test NFT 1",
			Image: "ipfs://QmcCJC4T37rykDjR6oorM8hpB9GQWHKWbAi2YR1uTabUZu/0",
		}, "", "")
		if err != nil {
			panic(err)
		}

		fmt.Println("Successfully uploaded to URI:", uri)
	},
}

func init() {
	storageCmd.AddCommand(storageUploadCmd)
	storageCmd.AddCommand(storageUploadBatchCmd)
	storageCmd.AddCommand(storageUploadImageCmd)
	storageCmd.AddCommand(storageUploadImageLinkCmd)
}
