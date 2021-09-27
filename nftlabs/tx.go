package nftlabs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"time"
)

func waitForTx(client *ethclient.Client, hash common.Hash, wait time.Duration, maxAttempts uint8) (error) {
	attempts := uint8(0)
	var syncError error
	for {
		if attempts >= maxAttempts {
			fmt.Println("Retry attempts to get tx exhausted, tx might have failed")
			return syncError
		}

		if tx, isPending, err := client.TransactionByHash(context.Background(), hash); err != nil {
			syncError = err
			log.Printf("Failed to get tx, err = %v\n", err)
			attempts += 1
			time.Sleep(wait)
			continue
		} else {
			if isPending {
				log.Println("Transaction still pending...")
				time.Sleep(wait)
				continue
			}

			log.Printf("Sync tx = %v, isPending = %v\n", tx, isPending)
			break
		}
	}

	return nil
}
