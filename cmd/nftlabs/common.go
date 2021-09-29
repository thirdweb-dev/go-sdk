package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	sdk "github.com/nftlabs/nftlabs-sdk-go/nftlabs"
	"log"
)

func getNftModule() (sdk.Nft, error) {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating an NFT sdk module on chain %v, contract %v\n", chainRpcUrl, contractAddress)

	// You can mock the sdk.Nft interface when writing tests against the SDK
	var module sdk.Nft
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

func getPackModule() (sdk.Pack, error) {
	// TODO: need to reuse these
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating a pack sdk module on chain %v, contract %v\n", chainRpcUrl, packContractAddress)

	// You can mock the sdk.Nft interface when writing tests against the SDK
	var module sdk.Pack
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

func getCollectionModule() (sdk.NftCollection, error) {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating an NFT sdk module on chain %v, contract %v\n", chainRpcUrl, collectionContractAddress)

	// You can mock the sdk.Nft interface when writing tests against the SDK
	var module sdk.NftCollection
	if caller, err := sdk.NewNftCollectionModule(client, collectionContractAddress, &sdk.SdkOptions{}); err != nil {
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

func getMarketplaceModule() (sdk.Market, error) {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating a Marketplace sdk module on chain %v, contract %v\n", chainRpcUrl, marketplaceContractAddress)

	// You can mock the sdk.Nft interface when writing tests against the SDK
	var module sdk.Market
	if caller, err := sdk.NewMarketSdkModule(client, marketplaceContractAddress, &sdk.SdkOptions{}); err != nil {
		log.Println("Failed to create an Marketplace caller object")
		panic(err)
	} else {
		module = caller
	}

	if err := module.SetPrivateKey(privateKey); err != nil {
		panic(err)
	}

	return module, nil
}

func getCurrencyModule() (sdk.Currency, error) {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating a CurrencyMetadata sdk module on chain %v, contract %v\n", chainRpcUrl, currencyContractAddress)

	// You can mock the sdk.Nft interface when writing tests against the SDK
	var module sdk.Currency
	if caller, err := sdk.NewCurrencySdkModule(client, currencyContractAddress); err != nil {
		log.Println("Failed to create an currency caller object")
		panic(err)
	} else {
		module = caller
	}

	if err := module.SetPrivateKey(privateKey); err != nil {
		panic(err)
	}

	return module, nil
}

