package main

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

var (
	thirdwebSDK *thirdweb.ThirdwebSDK
)

func initSdk() {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	if sdk, err := thirdweb.NewThirdwebSDK(client, privateKey, ""); err != nil {
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
		log.Println("Failed to create an NFT collection object")
		return nil, err
	} else {
		return contract, nil
	}	
}

func getNftDrop() (*thirdweb.NFTDrop, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a NFT Collection on chain %v, contract %v\n", chainRpcUrl, nftDropContractAddress)

	if contract, err := thirdwebSDK.GetNFTDrop(nftDropContractAddress); err != nil {
		log.Println("Failed to create an NFT collection object")
		return nil, err
	} else {
		return contract, nil
	}	
}


