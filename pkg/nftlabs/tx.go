package nftlabs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func waitForTx(client *ethclient.Client, hash common.Hash, wait time.Duration, maxAttempts uint8) error {
	attempts := uint8(0)
	var syncError error
	for {
		if attempts >= maxAttempts {
			fmt.Println("Retry attempts to get tx exhausted, tx might have failed")
			return syncError
		}

		if tx, isPending, err := client.TransactionByHash(context.Background(), hash); err != nil {
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
			break
		}
	}

	return nil
}
