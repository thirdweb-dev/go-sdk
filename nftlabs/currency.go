package nftlabs

import (
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
}

type CurrencySdkModule struct {
	Client *ethclient.Client
	Address string
	caller *abi.CurrencyCaller
}

func NewCurrencySdkModule(client *ethclient.Client, asset string) (*CurrencySdkModule, error) {
	caller, err := abi.NewCurrencyCaller(common.HexToAddress(asset), client)
	if err != nil {
		return nil, err
	}

	return &CurrencySdkModule{
		Client: client,
		Address: asset,
		caller: caller,
	}, nil
}

func (c *CurrencySdkModule) Get() (Currency, error) {
	panic("implement me")
}

func (c *CurrencySdkModule) GetValue(value *big.Int) (*CurrencyValue, error) {
	if c.Address == common.HexToAddress("0").Hex() {
		return &CurrencyValue{}, nil
	}

	name, err := c.caller.Name(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	symbol, err := c.caller.Symbol(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	decimals, err := c.caller.Decimals(&bind.CallOpts{})
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

func (c *CurrencySdkModule) Balance() (CurrencyValue, error) {
	panic("implement me")
}

func (c *CurrencySdkModule) BalanceOf(address string, tokenId string) (CurrencyValue, error) {
	panic("implement me")
}

func (c *CurrencySdkModule) Transfer(to string, amount *big.Int) error {
	panic("implement me")
}

