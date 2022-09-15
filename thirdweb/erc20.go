package thirdweb

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// This interface is currently support by the Token contract. You can access
// all of its functions through a Token contract instance.
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

// Get token metadata including name, symbol, decimals, etc.
//
// returns: the metadata for the token
//
// Example
//
//	currency, err := contract.Get()
//	symbol := currency.Symbol
func (erc20 *ERC20) Get() (*Currency, error) {
	return fetchCurrencyMetadata(erc20.helper.GetProvider(), erc20.helper.getAddress().String())
}

// Get the token balance of the connected wallet.
//
// returns: balance of the connected wallet
//
// Example
//
//		balance, err := contract.Balance()
//	 balanceValue := balance.DisplayValue
func (erc20 *ERC20) Balance() (*CurrencyValue, error) {
	return erc20.BalanceOf(erc20.helper.GetSignerAddress().String())
}

// Get the balance of the specified wallet.
//
// address: wallet address to check the balance of
//
// returns: balance of the specified wallet
//
// Example
//
//		address := "{{wallet_address}}"
//		balance, err := contract.BalanceOf()
//	 balanceValue := balance.DisplayValue
func (erc20 *ERC20) BalanceOf(address string) (*CurrencyValue, error) {
	balanceOf, err := erc20.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return erc20.getValue(balanceOf)
}

// Get the total minted supply of the token.
//
// returns: total minted supply of the token
func (erc20 *ERC20) TotalSupply() (*CurrencyValue, error) {
	totalySupply, err := erc20.abi.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	return erc20.getValue(totalySupply)
}

// Get a specified spenders allowance for the connected wallets tokens.
//
// spender: wallet address to check the allowance of
//
// returns: allowance of the spender for the connected wallets tokens
//
// Example
//
//	spender := "0x..."
//
//	allowance, err := contract.Allowance(spender)
//	allowanceValue := allowance.DisplayValue
func (erc20 *ERC20) Allowance(spender string) (*CurrencyValue, error) {
	return erc20.AllowanceOf(erc20.helper.GetSignerAddress().String(), spender)
}

// Get a specified spenders allowance for the a specific wallets tokens.
//
// owner: wallet address who owns the assets
//
// spender: wallet address to check the allowance of
//
// returns: allowance of the spender for the connected wallets tokens
//
// Example
//
//	address := "{{wallet_address}}"
//	spender := "0x..."
//
//	allowance, err := contract.AllowanceOf(address, spender)
//	allowanceValue := allowance.DisplayValue
func (erc20 *ERC20) AllowanceOf(owner string, spender string) (*CurrencyValue, error) {
	allowance, err := erc20.abi.Allowance(&bind.CallOpts{}, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		return nil, err
	}

	return erc20.getValue(allowance)
}

// Transfer a specified amount of tokens from the connected wallet to a specified address.
//
// to: address to transfer the tokens to
//
// amount: amount of tokens to transfer
//
// returns: transaction receipt of the transfer
//
// Example
//
//	to := "0x..."
//	amount := 1
//
//	tx, err := contract.Transfer(context.Background(), to, amount)
func (erc20 *ERC20) Transfer(ctx context.Context, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Transfer(txOpts, common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// Transfer a specified amount of tokens from one specified address to another.
//
// from: address to transfer the tokens from
//
// to: address to transfer the tokens to
//
// amount: amount of tokens to transfer
//
// returns: transaction receipt of the transfer
//
// Example
//
//	from := "{{wallet_address}}"
//	to := "0x..."
//	amount := 1
//
//	tx, err := contract.TransferFrom(context.Background(), from, to, amount)
func (erc20 *ERC20) TransferFrom(ctx context.Context, from string, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.TransferFrom(txOpts, common.HexToAddress(from), common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// Sets the allowance of a wallet to spend the connected wallets funds.
//
// spender: wallet address to set the allowance of
//
// amount: amount of tokens to grant the spender allowance of
//
// returns: transaction receipt of the allowance set
//
// Example
//
//	spender := "0x..."
//	amount := 1
//
//	tx, err := contract.SetAllowance(context.Background(), spender, amount)
func (erc20 *ERC20) SetAllowance(ctx context.Context, spender string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Approve(txOpts, common.HexToAddress(spender), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// Transfer tokens from the connected wallet to many wallets.
//
// args: list of token amounts with amounts and addresses to transfer to
//
// returns: transaction receipt of the transfers
//
// Example
//
//	args = []*thirdweb.TokenAmount{
//		&thirdweb.TokenAmount{
//			ToAddress: "0x...",
//			Amount:    1
//		}
//		&thirdweb.TokenAmount{
//			ToAddress: "0x...",
//			Amount:    2
//		}
//	}
//
//	tx, err := contract.TransferBatch(context.Background(), args)
func (erc20 *ERC20) TransferBatch(ctx context.Context, args []*TokenAmount) (*types.Transaction, error) {
	encoded := [][]byte{}

	for _, arg := range args {
		amountWithDecimals, err := erc20.normalizeAmount(arg.Amount)
		if err != nil {
			return nil, err
		}

		txOpts, err := erc20.helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := erc20.abi.Transfer(txOpts, common.HexToAddress(arg.ToAddress), amountWithDecimals)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := erc20.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// Burn a specified amount of tokens from the connected wallet.
//
// amount: amount of tokens to burn
//
// returns: transaction receipt of the burn
//
// Example
//
//	amount := 1
//	tx, err := contract.Burn(context.Background(), amount)
func (erc20 *ERC20) Burn(ctx context.Context, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Burn(txOpts, amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.awaitTx(tx.Hash())
}

// Burn a specified amount of tokens from a specific wallet.
//
// holder: wallet address to burn the tokens from
//
// amount: amount of tokens to burn
//
// returns: transaction receipt of the burn
//
// Example
//
//	holder := "0x..."
//	amount := 1
//
//	tx, err := contract.BurnFrom(context.Background(), holder, amount)
func (erc20 *ERC20) BurnFrom(ctx context.Context, holder string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.BurnFrom(txOpts, common.HexToAddress(holder), amountWithDecimals)
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

	return parseUnits(amount, currency.Decimals)
}
