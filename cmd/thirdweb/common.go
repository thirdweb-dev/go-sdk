package main

import (
	"log"

	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

var (
	thirdwebSDK *thirdweb.ThirdwebSDK
)

func initSdk() {
	if sdk, err := thirdweb.NewThirdwebSDK(
		chainRpcUrl,
		&thirdweb.SDKOptions{
			PrivateKey: privateKey,
		},
	); err != nil {
		panic(err)
	} else {
		thirdwebSDK = sdk
	}
}

func getNftCollection() (*thirdweb.NFTCollection, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Collection on chain %v, contract %v\n", chainRpcUrl, nftContractAddress)

	if contract, err := thirdwebSDK.GetNFTCollection(nftContractAddress); err != nil {
		log.Println("Failed to create an NFT Collection object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getEdition() (*thirdweb.Edition, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition on chain %v, contract %v\n", chainRpcUrl, editionAddress)

	if contract, err := thirdwebSDK.GetEdition(editionAddress); err != nil {
		log.Println("Failed to create an Edition object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getNftDrop() (*thirdweb.NFTDrop, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Drop on chain %v, contract %v\n", chainRpcUrl, nftDropContractAddress)

	if contract, err := thirdwebSDK.GetNFTDrop(nftDropContractAddress); err != nil {
		log.Println("Failed to create an NFT Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getStorage() thirdweb.Storage {
	return thirdweb.NewIpfsStorage(thirdweb.DEFAULT_IPFS_GATEWAY_URL)
}