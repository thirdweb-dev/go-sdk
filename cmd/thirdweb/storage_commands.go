package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
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
			[]any{
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

func init() {
	storageCmd.AddCommand(storageUploadCmd)
	storageCmd.AddCommand(storageUploadBatchCmd)
}
