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

	if err := module.SetPrivateKey(privateKey); err != nil {
		panic(err)
	}

	return module, nil
}

func getPackModule() (sdk.PackSdk, error) {
	// TODO: need to reuse these
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating a pack sdk module on chain %v, contract %v\n", chainRpcUrl, packContractAddress)

	// You can mock the sdk.NftSdk interface when writing tests against the SDK
	var module sdk.PackSdk
	if caller, err := sdk.NewPackSdkModule(client, packContractAddress, &sdk.SdkOptions{}); err != nil {
		log.Println("Failed to create a pack sdk module")
		panic(err)
	} else {
		module = caller
	}

	if err := module.SetPrivateKey(privateKey); err != nil {
		panic(err)
	}

	return module, nil
}
