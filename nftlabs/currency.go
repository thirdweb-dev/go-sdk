package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type CurrencySdk interface {
	Get() (Currency, error)
	GetValue(value *big.Int) (*CurrencyValue, error)
	Balance() (CurrencyValue, error)
	BalanceOf(address string, tokenId string) (CurrencyValue, error)
	Transfer(to string, amount *big.Int) error
	Allowance(spender string) (*big.Int, error)
	SetAllowance(spender string, amount *big.Int) error
	Mint(amount *big.Int) error
	Burn(amount *big.Int) error
	BurnFrom(from string, amount *big.Int) error
	TransferFrom(from string, to string, amount *big.Int) error
	GrantRole(role Role, address string) error
	RevokeRole(role Role, address string) error
	TotalSupply() (*big.Int, error)
}

type CurrencySdkModule struct {
	CommonModule
	Client *ethclient.Client
	Address string
	module *abi.Currency

	privateKey *ecdsa.PrivateKey
	rawPrivateKey string
	signerAddress common.Address
}

func (sdk *CurrencySdkModule) TotalSupply() (*big.Int, error) {
	return sdk.module.TotalSupply(&bind.CallOpts{})
}

func (sdk *CurrencySdkModule) Allowance(spender string) (*big.Int, error) {
	return sdk.module.Allowance(&bind.CallOpts{}, sdk.getSignerAddress(), common.HexToAddress(spender))
}

func (sdk *CurrencySdkModule) SetAllowance(spender string, amount *big.Int) error {
	if tx, err := sdk.module.Approve(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, common.HexToAddress(spender), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencySdkModule) Mint(amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.Mint(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, sdk.signerAddress, amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencySdkModule) Burn(amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.Burn(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencySdkModule) BurnFrom(from string, amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.BurnFrom(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, common.HexToAddress(from), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencySdkModule) TransferFrom(from string, to string, amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.TransferFrom(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, common.HexToAddress(from), common.HexToAddress(to), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencySdkModule) GrantRole(role Role, address string) error {
	if tx, err := sdk.module.CurrencyTransactor.GrantRole(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, [32]byte{}, common.HexToAddress(address)); err != nil { // TODO: fill in role in [32]byte
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencySdkModule) RevokeRole(role Role, address string) error {
	if tx, err := sdk.module.CurrencyTransactor.RevokeRole(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, [32]byte{}, common.HexToAddress(address)); err != nil { // TODO: fill in role in [32]byte
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func NewCurrencySdkModule(client *ethclient.Client, asset string) (*CurrencySdkModule, error) {
	module, err := abi.NewCurrency(common.HexToAddress(asset), client)
	if err != nil {
		return nil, err
	}

	return &CurrencySdkModule{
		Client: client,
		Address: asset,
		module: module,
	}, nil
}

func (sdk *CurrencySdkModule) Get() (Currency, error) {
	panic("implement me")
}

func (sdk *CurrencySdkModule) GetValue(value *big.Int) (*CurrencyValue, error) {
	if sdk.Address == common.HexToAddress("0").Hex() {
		return &CurrencyValue{}, nil
	}

	name, err := sdk.module.CurrencyCaller.Name(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	symbol, err := sdk.module.CurrencyCaller.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	decimals, err := sdk.module.CurrencyCaller.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return &CurrencyValue{
		Currency: Currency{
			Name: name,
			Symbol: symbol,
			Decimals: decimals,
		},
	}, nil
}

func (sdk *CurrencySdkModule) Balance() (CurrencyValue, error) {
	panic("implement me")
}

func (sdk *CurrencySdkModule) BalanceOf(address string, tokenId string) (CurrencyValue, error) {
	panic("implement me")
}

func (sdk *CurrencySdkModule) Transfer(to string, amount *big.Int) error {
	panic("implement me")
}

func (sdk *CurrencySdkModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.rawPrivateKey = privateKey
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}

func (sdk *CurrencySdkModule) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}

func (sdk *CurrencySdkModule) getSignerAddress() common.Address {
	if sdk.signerAddress == common.HexToAddress("0") {
		return common.HexToAddress(sdk.Address)
	} else {
		return sdk.signerAddress
	}
}
