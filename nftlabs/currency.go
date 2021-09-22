package nftlabs

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
	"math/big"
)

type CurrencySdk interface {
	Get() (Currency, error)
	GetValue(value *big.Int) (CurrencyValue, error)
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

func (c CurrencySdkModule) Get() (Currency, error) {
	panic("implement me")
}

func (c CurrencySdkModule) GetValue(value *big.Int) (CurrencyValue, error) {
	panic("implement me")
}

func (c CurrencySdkModule) Balance() (CurrencyValue, error) {
	panic("implement me")
}

func (c CurrencySdkModule) BalanceOf(address string, tokenId string) (CurrencyValue, error) {
	panic("implement me")
}

func (c CurrencySdkModule) Transfer(to string, amount *big.Int) error {
	panic("implement me")
}

