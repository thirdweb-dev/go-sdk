package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/pkg/nftlabs"
	"log"
)

var (
	nftlabsSdk nftlabs.ISdk
)

func initSdk() {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	if sdk, err := nftlabs.NewSdk(client, &nftlabs.SdkOptions{
		PrivateKey: privateKey,
	}); err != nil {
		panic(err)
	} else {
		nftlabsSdk = sdk
	}
}

func getNftModule() (nftlabs.Nft, error) {
	if nftlabsSdk == nil {
		initSdk()
	}

	log.Printf("Creating an NFT nftlabs module on chain %v, contract %v\n", chainRpcUrl, nftContractAddress)

	// You can mock the nftlabs.Nft interface when writing tests against the SDK
	if caller, err := nftlabsSdk.GetNftModule(nftContractAddress); err != nil {
		log.Println("Failed to create an NFT caller object")
		panic(err)
	} else {
		return caller, nil
	}
}

func getPackModule() (nftlabs.Pack, error) {
	if nftlabsSdk == nil {
		initSdk()
	}

	log.Printf("Creating a pack nftlabs module on chain %v, contract %v\n", chainRpcUrl, packContractAddress)

	if caller, err := nftlabsSdk.GetPackModule(packContractAddress); err != nil {
		log.Println("Failed to create a Pack module")
		panic(err)
	} else {
		return caller, nil
	}
}

func getCollectionModule() (nftlabs.NftCollection, error) {
	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Creating an NFT nftlabs module on chain %v, contract %v\n", chainRpcUrl, collectionContractAddress)

	// You can mock the nftlabs.Nft interface when writing tests against the SDK
	var module nftlabs.NftCollection
	if caller, err := nftlabs.NewNftCollectionModule(client, collectionContractAddress, &nftlabs.SdkOptions{}); err != nil {
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

func getMarketplaceModule() (nftlabs.Market, error) {
	if nftlabsSdk == nil {
		initSdk()
	}

	log.Printf("Creating a Marketplace nftlabs module on chain %v, contract %v\n", chainRpcUrl, marketplaceContractAddress)

	// You can mock the nftlabs.Nft interface when writing tests against the SDK
	if caller, err := nftlabsSdk.GetMarketModule(marketplaceContractAddress); err != nil {
		log.Println("Failed to create a Marketplace module")
		panic(err)
	} else {
		return caller, nil
	}
}

func getCurrencyModule() (nftlabs.Currency, error) {
	if nftlabsSdk == nil {
		initSdk()
	}

	log.Printf("Creating a CurrencyMetadata nftlabs module on chain %v, contract %v\n", chainRpcUrl, currencyContractAddress)

	if caller, err := nftlabsSdk.GetCurrencyModule(currencyContractAddress); err != nil {
		log.Println("Failed to create an Currency module")
		panic(err)
	} else {
		return caller, nil
	}
}

