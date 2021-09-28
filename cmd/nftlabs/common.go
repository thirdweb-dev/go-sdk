package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	sdk "github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"log"
)

func getNftModule() (sdk.NftSdk, error) {
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
	return module, nil
}
