package thirdweb

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/imdario/mergo"
)

type contractHelper struct {
	address       common.Address
	nextOverrides *bind.TransactOpts
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
			nil,
			handler,
		}
		return helper, nil
	}
}

func (helper *contractHelper) getAddress() common.Address {
	return helper.address
}

func (helper *contractHelper) mergeOverrides(opts *bind.TransactOpts) (*bind.TransactOpts, error) {
	if (helper.nextOverrides != nil) {
		if err := mergo.Merge(opts, helper.nextOverrides); err != nil {
			return nil, err
		}

		helper.nextOverrides = nil
	}

	return opts, nil
}

func (helper *contractHelper) OverrideNextTransaction(opts *bind.TransactOpts) {
	helper.nextOverrides = opts
}

func (helper *contractHelper) getUnsignedTxOptions(ctx context.Context, signerAddress string) (*bind.TransactOpts, error) {
	var tipCap, feeCap *big.Int

	provider := helper.GetProvider()
	block, err := provider.BlockByNumber(ctx, nil)
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

	txOpts, err = helper.mergeOverrides(txOpts)
	if err != nil {
		return nil, err
	}

	return txOpts, nil
}

func (helper *contractHelper) getEncodedTxOptions(ctx context.Context) (*bind.TransactOpts, error) {
	return helper.getRawTxOptions(ctx, true)
}

func (helper *contractHelper) GetTxOptions(ctx context.Context) (*bind.TransactOpts, error) {
	return helper.getRawTxOptions(ctx, false)
}

func (helper *contractHelper) getRawTxOptions(ctx context.Context, noSend bool) (*bind.TransactOpts, error) {
	if helper.GetRawPrivateKey() == "" {
		return nil, fmt.Errorf("You need to set a private key to use this function!")
	}

	var tipCap, feeCap *big.Int

	provider := helper.GetProvider()
	chainId, err := provider.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	block, err := provider.BlockByNumber(ctx, nil)

	if err == nil && block.BaseFee() != nil {
		// Use gas station to get maxPriorityFeePerGas if we're on polygon
		if chainId.Cmp(big.NewInt(137)) == 0 {
			tipCap, _ = helper.getPolygonGasPriorityFee(ctx)
		} else {
			tipCap, _ = big.NewInt(0).SetString("2500000000", 10) // maxPriorityFeePerGas
		}

		baseFee := big.NewInt(0).Mul(block.BaseFee(), big.NewInt(2))
		feeCap = big.NewInt(0).Add(baseFee, tipCap) // maxFeePerGas
	}

	signer, err := helper.getSigner(ctx)
	if err != nil {
		return nil, err
	}
	txOpts := &bind.TransactOpts{
		Context:   ctx,
		NoSend:    noSend,
		From:      helper.GetSignerAddress(),
		Signer:    signer,
		GasTipCap: tipCap, // maxPriorityFeePerGas
		GasFeeCap: feeCap, // maxFeePerGas
	}

	txOpts, err = helper.mergeOverrides(txOpts)
	if err != nil {
		return nil, err
	}

	return txOpts, nil
}

func (helper *contractHelper) AwaitTx(ctx context.Context, hash common.Hash) (*types.Transaction, error) {
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

		if tx, isPending, err := provider.TransactionByHash(ctx, hash); err != nil {
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

func (helper *contractHelper) getPolygonGasPriorityFee(ctx context.Context) (*big.Int, error) {
	getTipCap := func() (*big.Int, error) {
		req, err := http.NewRequestWithContext(ctx, "GET", "https://gasstation-mainnet.matic.network/v2", nil)
		if err != nil {
			return nil, err
		}

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}

		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		var feeData map[string]map[string]interface{}
		json.Unmarshal(bodyBytes, &feeData)

		priorityFee, ok := feeData["standard"]["maxPriorityFee"].(float64)
		if !ok {
			return nil, err
		}

		return parseUnits(priorityFee, 9) // maxPriorityFeePerGas
	}

	tipCap, err := getTipCap()
	if err == nil {
		return tipCap, nil
	}

	return parseUnits(31, 9)
}
