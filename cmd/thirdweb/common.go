package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
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

func awaitTx(hash common.Hash) (*types.Transaction, error) {
	provider := thirdwebSDK.GetProvider()
	wait := time.Second * 1
	maxAttempts := uint8(20)
	attempts := uint8(0)

	var syncError error
	for {
		if attempts >= maxAttempts {
			fmt.Println("Retry attempts to get tx exhausted, tx might have failed")
			return nil, syncError
		}

		if tx, isPending, err := provider.TransactionByHash(context.Background(), hash); err != nil {
			syncError = err
			log.Printf("Failed to get tx %v, err = %v\n", hash.String(), err)
			attempts += 1
			time.Sleep(wait)
			continue
		} else {
			if isPending {
				log.Println("Transaction still pending...")
				time.Sleep(wait)
				continue
			}
			log.Printf("Transaction with hash %v mined successfully\n", tx.Hash())
			return tx, nil
		}
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

func getToken() (*thirdweb.Token, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Token on chain %v, contract %v\n", chainRpcUrl, tokenAddress)

	if contract, err := thirdwebSDK.GetToken(tokenAddress); err != nil {
		log.Println("Failed to create an Token object")
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

func getEditionDrop() (*thirdweb.EditionDrop, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Edition Drop on chain %v, contract %v\n", chainRpcUrl, editionDropContractAddress)

	if contract, err := thirdwebSDK.GetEditionDrop(editionDropContractAddress); err != nil {
		log.Println("Failed to create an Edition Drop object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMultiwrap() (*thirdweb.Multiwrap, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Multiwrap on chain %v, contract %v\n", chainRpcUrl, multiwrapContractAddress)

	if contract, err := thirdwebSDK.GetMultiwrap(multiwrapContractAddress); err != nil {
		log.Println("Failed to create a Multiwrap object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getMarketplace() (*thirdweb.Marketplace, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Marketplace on chain %v, contract %v\n", chainRpcUrl, marketplaceAddress)

	if contract, err := thirdwebSDK.GetMarketplace(marketplaceAddress); err != nil {
		log.Println("Failed to create a Marketplace object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getCustom() (*thirdweb.SmartContract, error) {
	if thirdwebSDK == nil {
		initSdk()
	}

	log.Printf("Obtaining a Custom on chain %v, contract %v\n", chainRpcUrl, customContractAddress)

	if contract, err := thirdwebSDK.GetContract(customContractAddress); err != nil {
		log.Println("Failed to create an Custom object")
		return nil, err
	} else {
		return contract, nil
	}
}

func getStorage() thirdweb.IpfsStorage {
	if thirdwebSDK == nil {
		initSdk()
	}

	return thirdwebSDK.Storage
}
