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

type contractHelper struct {
	address common.Address
	*ProviderHandler
}

const (
	txWaitTimeBetweenAttempts = time.Second * 1
	txMaxAttempts             = 20
)

func newContractHelper(address common.Address, provider *ethclient.Client, privateKey string) (*contractHelper, error) {
	if handler, err := NewProviderHandler(provider, privateKey); err != nil {
		return nil, err
	} else {
		helper := &contractHelper{
			address,
			handler,
		}
		return helper, nil
	}
}

func (helper *contractHelper) getAddress() common.Address {
	return helper.address
}

func (helper *contractHelper) getUnsignedTxOptions(ctx context.Context, signerAddress string) (*bind.TransactOpts, error) {
	var tipCap, feeCap *big.Int

	provider := helper.GetProvider()
	block, err := provider.BlockByNumber(context.Background(), nil)
	if err == nil && block.BaseFee() != nil {
		tipCap, _ = big.NewInt(0).SetString("2500000000", 10)
		baseFee := big.NewInt(0).Mul(block.BaseFee(), big.NewInt(2))
		feeCap = big.NewInt(0).Add(baseFee, tipCap)
	}

	txOpts := &bind.TransactOpts{
		Context:   ctx,
		NoSend:    true,
		From:      common.HexToAddress(signerAddress),
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
	}

	txOpts.Signer = func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return transaction, nil
	}

	return txOpts, nil
}

func (helper *contractHelper) getEncodedTxOptions(ctx context.Context) (*bind.TransactOpts, error) {
	return helper.getRawTxOptions(ctx, true)
}

func (helper *contractHelper) getTxOptions(ctx context.Context) (*bind.TransactOpts, error) {
	return helper.getRawTxOptions(ctx, false)
}

func (helper *contractHelper) getRawTxOptions(ctx context.Context, noSend bool) (*bind.TransactOpts, error) {
	if helper.GetRawPrivateKey() == "" {
		return nil, fmt.Errorf("You need to set a private key to use this function!")
	}

	var tipCap, feeCap *big.Int

	provider := helper.GetProvider()
	block, err := provider.BlockByNumber(context.Background(), nil)
	if err == nil && block.BaseFee() != nil {
		tipCap, _ = big.NewInt(0).SetString("2500000000", 10)
		baseFee := big.NewInt(0).Mul(block.BaseFee(), big.NewInt(2))
		feeCap = big.NewInt(0).Add(baseFee, tipCap)
	}

	txOpts := &bind.TransactOpts{
		Context:   ctx,
		NoSend:    noSend,
		From:      helper.GetSignerAddress(),
		Signer:    helper.getSigner(),
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
	}

	return txOpts, nil
}

func (helper *contractHelper) awaitTx(hash common.Hash) (*types.Transaction, error) {
	provider := helper.GetProvider()
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
