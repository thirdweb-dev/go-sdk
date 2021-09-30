package nftlabs

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type Currency interface {
	Get() (CurrencyMetadata, error)
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

type CurrencyModule struct {
	Client  *ethclient.Client
	Address string
	module  *abi.Currency

	main ISdk
}

func newCurrencyModule(client *ethclient.Client, asset string, main ISdk) (*CurrencyModule, error) {
	module, err := abi.NewCurrency(common.HexToAddress(asset), client)
	if err != nil {
		return nil, err
	}

	return &CurrencyModule{
		Client:  client,
		Address: asset,
		module:  module,
		main: main,
	}, nil
}


func (sdk *CurrencyModule) TotalSupply() (*big.Int, error) {
	return sdk.module.TotalSupply(&bind.CallOpts{})
}

func (sdk *CurrencyModule) Allowance(spender string) (*big.Int, error) {
	return sdk.module.Allowance(&bind.CallOpts{}, sdk.main.getSignerAddress(), common.HexToAddress(spender))
}

func (sdk *CurrencyModule) SetAllowance(spender string, amount *big.Int) error {
	if tx, err := sdk.module.Approve(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, common.HexToAddress(spender), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) Mint(amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.Mint(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, sdk.main.getSignerAddress(), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) Burn(amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.Burn(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) BurnFrom(from string, amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.BurnFrom(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, common.HexToAddress(from), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) TransferFrom(from string, to string, amount *big.Int) error {
	if tx, err := sdk.module.CurrencyTransactor.TransferFrom(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, common.HexToAddress(from), common.HexToAddress(to), amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) GrantRole(role Role, address string) error {
	if tx, err := sdk.module.CurrencyTransactor.GrantRole(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, [32]byte{}, common.HexToAddress(address)); err != nil { // TODO: fill in role in [32]byte
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) RevokeRole(role Role, address string) error {
	if tx, err := sdk.module.CurrencyTransactor.RevokeRole(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, [32]byte{}, common.HexToAddress(address)); err != nil { // TODO: fill in role in [32]byte
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *CurrencyModule) Get() (CurrencyMetadata, error) {
	if strings.HasPrefix(sdk.Address, "0x0000000") {
		return CurrencyMetadata{}, nil
	}

	erc20Module, err := newErc20SdkModule(sdk.Client, sdk.Address, &SdkOptions{})
	if err != nil {
		return CurrencyMetadata{}, err
	}

	name, err := erc20Module.module.Name(&bind.CallOpts{})
	if err != nil {
		return CurrencyMetadata{}, err
	}

	symbol, err := erc20Module.module.Symbol(&bind.CallOpts{})
	if err != nil {
		return CurrencyMetadata{}, err
	}

	decimals, err := erc20Module.module.Decimals(&bind.CallOpts{})
	if err != nil {
		return CurrencyMetadata{}, err
	}

	return CurrencyMetadata{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
	}, nil
}

func (sdk *CurrencyModule) formatUnits(value *big.Int, units *big.Int) string {
	if value.Int64() == 0 {
		return "0"
	}

	unit := big.NewInt(18)
	if units != nil {
		unit.Set(units)
	}

	ten := big.NewInt(10)
	ten.Exp(ten, unit, big.NewInt(0))
	v := big.NewInt(0)
	v.SetString(value.String(), 10)
	return v.Div(v, ten).String()
}

func (sdk *CurrencyModule) GetValue(value *big.Int) (*CurrencyValue, error) {
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
		CurrencyMetadata: CurrencyMetadata{
			Name:     name,
			Symbol:   symbol,
			Decimals: decimals,
		},
		Value: value,
		DisplayValue: sdk.formatUnits(value, big.NewInt(int64(decimals))),
	}, nil
}

func (sdk *CurrencyModule) Balance() (CurrencyValue, error) {
	panic("implement me")
}

func (sdk *CurrencyModule) BalanceOf(address string, tokenId string) (CurrencyValue, error) {
	panic("implement me")
}

func (sdk *CurrencyModule) Transfer(to string, amount *big.Int) error {
	panic("implement me")
}

