package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type ERC20 struct {
	abi     *abi.TokenERC20
	helper  *contractHelper
	storage storage
}

func newERC20(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC20, error) {
	if contractAbi, err := abi.NewTokenERC20(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		return &ERC20{
			contractAbi,
			helper,
			storage,
		}, nil
	}
}

// Get
//
// Get token metadata including name, symbol, decimals, etc.
//
// returns: the metadata for the token
func (erc20 *ERC20) Get() (*Currency, error) {
	return fetchCurrencyMetadata(erc20.helper.GetProvider(), erc20.helper.getAddress().String())
}

// Balance
//
// Get the token balance of the connected wallet
//
// returns: balance of the connected wallet
func (erc20 *ERC20) Balance() (*CurrencyValue, error) {
	return erc20.BalanceOf(erc20.helper.GetSignerAddress().String())
}

// BalanceOf
//
// Get the balance of the specified wallet
//
// address: wallet address to check the balance of
//
// returns: balance of the specified wallet
func (erc20 *ERC20) BalanceOf(address string) (*CurrencyValue, error) {
	balanceOf, err := erc20.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return erc20.getValue(balanceOf)
}

// TotalSupply
//
// Get the total minted supply of the token
//
// returns: total minted supply of the token
func (erc20 *ERC20) TotalSupply() (*CurrencyValue, error) {
	totalySupply, err := erc20.abi.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return erc20.getValue(totalySupply)
}

// Allowance
//
// Get a specified spenders allowance for the connected wallets tokens
//
// spender: wallet address to check the allowance of
//
// returns: allowance of the spender for the connected wallets tokens
func (erc20 *ERC20) Allowance(spender string) (*CurrencyValue, error) {
	return erc20.AllowanceOf(erc20.helper.GetSignerAddress().String(), spender)
}

// AllowanceOf
//
// Get a specified spenders allowance for the a specific wallets tokens
//
// owner: wallet address who owns the assets
//
// spender: wallet address to check the allowance of
//
// returns: allowance of the spender for the connected wallets tokens
func (erc20 *ERC20) AllowanceOf(owner string, spender string) (*CurrencyValue, error) {
	allowance, err := erc20.abi.Allowance(&bind.CallOpts{}, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		return nil, err
	}

	return erc20.getValue(allowance)
}

// Transfer
//
// Transfer a specified amount of tokens from the connected wallet to a specified address.
//
// to: address to transfer the tokens to
//
// amount: amount of tokens to transfer
//
// returns: transaction receipt of the transfer
func (erc20 *ERC20) Transfer(to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	tx, err := erc20.abi.Transfer(erc20.helper.getTxOptions(), common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// TransferFrom
//
// Transfer a specified amount of tokens from one specified address to another.
//
// from: address to transfer the tokens from
//
// to: address to transfer the tokens to
//
// amount: amount of tokens to transfer
//
// returns: transaction receipt of the transfer
func (erc20 *ERC20) TransferFrom(from string, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	tx, err := erc20.abi.TransferFrom(erc20.helper.getTxOptions(), common.HexToAddress(from), common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// SetAllowance
//
// Sets the allowance of a wallet to spend the connected wallets funds
//
// spender: wallet address to set the allowance of
//
// amount: amount of tokens to grant the spender allowance of
//
// returns: transaction receipt of the allowance set
func (erc20 *ERC20) SetAllowance(spender string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	tx, err := erc20.abi.Approve(erc20.helper.getTxOptions(), common.HexToAddress(spender), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// TransferBatch
//
// Transfer tokens from the connected wallet to many wallets
//
// args: list of token amounts with amounts and addresses to transfer to
//
// returns: transaction receipt of the transfers
func (erc20 *ERC20) TransferBatch(args []*TokenAmount) (*types.Transaction, error) {
	encoded := [][]byte{}

	for _, arg := range args {
		amountWithDecimals, err := erc20.normalizeAmount(arg.Amount)
		if err != nil {
			return nil, err
		}

		tx, err := erc20.abi.Transfer(erc20.helper.getTxOptions(), common.HexToAddress(arg.ToAddress), amountWithDecimals)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := erc20.abi.Multicall(erc20.helper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// Burn
//
// Burn a specified amount of tokens from the connected wallet
//
// amount: amount of tokens to burn
//
// returns: transaction receipt of the burn
func (erc20 *ERC20) Burn(amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	tx, err := erc20.abi.Burn(erc20.helper.getTxOptions(), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// BurnFrom
//
// Burn a specified amount of tokens from a specific wallet
//
// holder: wallet address to burn the tokens from
//
// amount: amount of tokens to burn
//
// returns: transaction receipt of the burn
func (erc20 *ERC20) BurnFrom(holder string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	tx, err := erc20.abi.BurnFrom(erc20.helper.getTxOptions(), common.HexToAddress(holder), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

func (erc20 *ERC20) getValue(value *big.Int) (*CurrencyValue, error) {
	return fetchCurrencyValue(
		erc20.helper.GetProvider(),
		erc20.helper.getAddress().String(),
		value,
	)
}

func (erc20 *ERC20) normalizeAmount(amount float64) (*big.Int, error) {
	currency, err := erc20.Get()
	if err != nil {
		return nil, err
	}

	return parseUnits(amount, currency.Decimals), nil
}
