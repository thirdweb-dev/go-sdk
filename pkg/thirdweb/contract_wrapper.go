package thirdweb

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type contractWrapper struct {
	address common.Address
	*ProviderHandler
}

const (
	txWaitTimeBetweenAttempts = time.Second * 1
	txMaxAttempts             = 20
)

func newContractWrapper(address common.Address, provider *ethclient.Client, privateKey string) (*contractWrapper, error) {
	if handler, err := NewProviderHandler(provider, privateKey); err != nil {
		return nil, err
	} else {
		wrapper := &contractWrapper{
			address,
			handler,
		}
		return wrapper, nil
	}
}

func (wrapper *contractWrapper) getAddress() common.Address {
	return wrapper.address
}

func (wrapper *contractWrapper) getTxOptions() *bind.TransactOpts {
	var tipCap, feeCap *big.Int

	provider := wrapper.GetProvider()
	block, err := provider.BlockByNumber(context.Background(), nil)
	if err == nil && block.BaseFee() != nil {
		tipCap, _ = big.NewInt(0).SetString("2500000000", 10)
		baseFee := big.NewInt(0).Mul(block.BaseFee(), big.NewInt(2))
		feeCap = big.NewInt(0).Add(baseFee, tipCap)
	}

	return &bind.TransactOpts{
		NoSend:    false,
		From:      wrapper.GetSignerAddress(),
		Signer:    wrapper.getSigner(),
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
	}
}

func (wrapper *contractWrapper) awaitTx(hash common.Hash) (*types.Transaction, error) {
	provider := wrapper.GetProvider()
	wait := txWaitTimeBetweenAttempts
	maxAttempts := uint8(txMaxAttempts)
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
