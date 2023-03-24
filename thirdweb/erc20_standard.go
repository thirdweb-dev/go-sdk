package thirdweb

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// This interface is currently support by the Token contract. You can access
// all of its functions through a Token contract instance.
type ERC20Standard struct {
	erc20 *ERC20
}

func newERC20Standard(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC20Standard, error) {
	erc20, err := newERC20(provider, address, privateKey, storage)
	if err != nil {
		return nil, err
	}

	return &ERC20Standard{
		erc20,
	}, nil
}

// Get token metadata including name, symbol, decimals, etc.
//
// returns: the metadata for the token
//
// Example
//
//	currency, err := contract.Get()
//	symbol := currency.Symbol
func (erc20 *ERC20Standard) Get(ctx context.Context) (*Currency, error) {
	return erc20.erc20.Get(ctx)
}

// Get the token balance of the connected wallet.
//
// returns: balance of the connected wallet
//
// Example
//
//		balance, err := contract.Balance()
//	 balanceValue := balance.DisplayValue
func (erc20 *ERC20Standard) Balance(ctx context.Context) (*CurrencyValue, error) {
	return erc20.erc20.Balance(ctx)
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
func (erc20 *ERC20Standard) BalanceOf(ctx context.Context, address string) (*CurrencyValue, error) {
	return erc20.erc20.BalanceOf(ctx, address)
}

// Get the total minted supply of the token.
//
// returns: total minted supply of the token
func (erc20 *ERC20Standard) TotalSupply(ctx context.Context) (*CurrencyValue, error) {
	return erc20.erc20.TotalSupply(ctx)
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
func (erc20 *ERC20Standard) Allowance(ctx context.Context, spender string) (*CurrencyValue, error) {
	return erc20.erc20.Allowance(ctx, spender)
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
func (erc20 *ERC20Standard) AllowanceOf(ctx context.Context, owner string, spender string) (*CurrencyValue, error) {
	return erc20.erc20.AllowanceOf(ctx, owner, spender)
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
func (erc20 *ERC20Standard) Transfer(ctx context.Context, to string, amount float64) (*types.Transaction, error) {
	return erc20.erc20.Transfer(ctx, to, amount)
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
func (erc20 *ERC20Standard) TransferFrom(ctx context.Context, from string, to string, amount float64) (*types.Transaction, error) {
	return erc20.erc20.TransferFrom(ctx, from, to, amount)
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
func (erc20 *ERC20Standard) SetAllowance(ctx context.Context, spender string, amount float64) (*types.Transaction, error) {
	return erc20.erc20.SetAllowance(ctx, spender, amount)
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
func (erc20 *ERC20Standard) TransferBatch(ctx context.Context, args []*TokenAmount) (*types.Transaction, error) {
	return erc20.erc20.TransferBatch(ctx, args)
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
func (erc20 *ERC20Standard) Burn(ctx context.Context, amount float64) (*types.Transaction, error) {
	return erc20.erc20.Burn(ctx, amount)
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
func (erc20 *ERC20Standard) BurnFrom(ctx context.Context, holder string, amount float64) (*types.Transaction, error) {
	return erc20.erc20.BurnFrom(ctx, holder, amount)
}