package thirdweb

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
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

// Get token metadata
//
// @extension: ERC20
//
// returns: the metadata for the token
//
// Example
//
//	currency, err := contract.ERC20.Get()
//	symbol := currency.Symbol
func (erc20 *ERC20) Get(ctx context.Context) (*Currency, error) {
	return fetchCurrencyMetadata(ctx, erc20.helper.GetProvider(), erc20.helper.getAddress().String())
}

// Get token balance
//
// @extension: ERC20
//
// returns: balance of the connected wallet
//
// Example
//
// 	balance, err := contract.ERC20.Balance()
// 	balanceValue := balance.DisplayValue
func (erc20 *ERC20) Balance(ctx context.Context) (*CurrencyValue, error) {
	return erc20.BalanceOf(ctx, erc20.helper.GetSignerAddress().String())
}

// Get token balance of a specific wallet
//
// @extension: ERC20
//
// address: wallet address to check the balance of
//
// returns: balance of the specified wallet
//
// Example
//
// 	address := "{{wallet_address}}"
// 	balance, err := contract.ERC20.BalanceOf()
// 	balanceValue := balance.DisplayValue
func (erc20 *ERC20) BalanceOf(ctx context.Context, address string) (*CurrencyValue, error) {
	balanceOf, err := erc20.abi.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return erc20.getValue(ctx, balanceOf)
}

// Get the total minted supply
//
// @extension: ERC20
//
// returns: total minted supply of the token
//
// Example
//
// 	supply, err := contract.ERC20.TotalSupply(context.Background())
func (erc20 *ERC20) TotalSupply(ctx context.Context) (*CurrencyValue, error) {
	totalySupply, err := erc20.abi.TotalSupply(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	return erc20.getValue(ctx, totalySupply)
}

// Get token allowance for a specific spender
//
// @extension: ERC20
//
// spender: wallet address to check the allowance of
//
// returns: allowance of the spender for the connected wallets tokens
//
// Example
//
//	spender := "0x..."
//
//	allowance, err := contract.ERC20.Allowance(spender)
//	allowanceValue := allowance.DisplayValue
func (erc20 *ERC20) Allowance(ctx context.Context, spender string) (*CurrencyValue, error) {
	return erc20.AllowanceOf(ctx, erc20.helper.GetSignerAddress().String(), spender)
}

// Get token allowance for a specific spender and owner
//
// @extension: ERC20
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
//	allowance, err := contract.ERC20.AllowanceOf(address, spender)
//	allowanceValue := allowance.DisplayValue
func (erc20 *ERC20) AllowanceOf(ctx context.Context, owner string, spender string) (*CurrencyValue, error) {
	allowance, err := erc20.abi.Allowance(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner), common.HexToAddress(spender))
	if err != nil {
		return nil, err
	}

	return erc20.getValue(ctx, allowance)
}

// Transfer tokens
//
// @extension: ERC20
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
//	tx, err := contract.ERC20.Transfer(context.Background(), to, amount)
func (erc20 *ERC20) Transfer(ctx context.Context, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Transfer(txOpts, common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

// Transfer tokens from a specific wallet
//
// @extension: ERC20
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
//	tx, err := contract.ERC20.TransferFrom(context.Background(), from, to, amount)
func (erc20 *ERC20) TransferFrom(ctx context.Context, from string, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.TransferFrom(txOpts, common.HexToAddress(from), common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

// Set token allowance
//
// @extension: ERC20
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
//	tx, err := contract.ERC20.SetAllowance(context.Background(), spender, amount)
func (erc20 *ERC20) SetAllowance(ctx context.Context, spender string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Approve(txOpts, common.HexToAddress(spender), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

// Transfer many tokens
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
//	tx, err := contract.ERC20.TransferBatch(context.Background(), args)
func (erc20 *ERC20) TransferBatch(ctx context.Context, args []*TokenAmount) (*types.Transaction, error) {
	encoded := [][]byte{}

	for _, arg := range args {
		amountWithDecimals, err := erc20.normalizeAmount(ctx, arg.Amount)
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

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

// Burn tokens
//
// @extension: ERC20Burnable
//
// amount: amount of tokens to burn
//
// returns: transaction receipt of the burn
//
// Example
//
//	amount := 1
//	tx, err := contract.ERC20.Burn(context.Background(), amount)
func (erc20 *ERC20) Burn(ctx context.Context, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Burn(txOpts, amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

// Burn tokens from a specific wallet
//
// @extension: ERC20Burnable
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
//	tx, err := contract.ERC20.BurnFrom(context.Background(), holder, amount)
func (erc20 *ERC20) BurnFrom(ctx context.Context, holder string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.BurnFrom(txOpts, common.HexToAddress(holder), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}


// Mint tokens
//
// @extension: ERC20Mintable
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
//
// Example
//
// 	tx, err := contract.ERC20.Mint(context.Background(), 1)
func (erc20 *ERC20) Mint(ctx context.Context, amount float64) (*types.Transaction, error) {
	return erc20.MintTo(ctx, erc20.helper.GetSignerAddress().String(), amount)
}

// Mint tokens to a specific wallet
//
// @extension: ERC20Mintable
//
// to: wallet address to mint tokens to
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
//
// Example
//
//	tx, err := contract.ERC20.MintTo(context.Background(), "{{wallet_address}}", 1)
func (erc20 *ERC20) MintTo(ctx context.Context, to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := erc20.normalizeAmount(ctx, amount)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.MintTo(txOpts, common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

// Mint tokens to many wallets
//
// @extension: ERC20BatchMintable
//
// args: list of wallet addresses and amounts to mint
//
// returns: transaction receipt of the mint
//
// Example
//
//	args = []*thirdweb.TokenAmount{
//		&thirdweb.TokenAmount{
//			ToAddress: "{{wallet_address}}",
//			Amount:    1
//		}
//		&thirdweb.TokenAmount{
//			ToAddress: "{{wallet_address}}",
//			Amount:    2
//		}
//	}
//
//	tx, err := contract.ERC20.MintBatchTo(context.Background(), args)
func (erc20 *ERC20) MintBatchTo(ctx context.Context, args []*TokenAmount) (*types.Transaction, error) {
	encoded := [][]byte{}

	for _, arg := range args {
		amountWithDecimals, err := erc20.normalizeAmount(ctx, arg.Amount)
		if err != nil {
			return nil, err
		}

		txOpts, err := erc20.helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := erc20.abi.MintTo(txOpts, common.HexToAddress(arg.ToAddress), amountWithDecimals)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := erc20.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc20.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return erc20.helper.AwaitTx(ctx, tx.Hash())
}

func (erc20 *ERC20) getValue(ctx context.Context, value *big.Int) (*CurrencyValue, error) {
	return fetchCurrencyValue(
		ctx,
		erc20.helper.GetProvider(),
		erc20.helper.getAddress().String(),
		value,
	)
}

func (erc20 *ERC20) normalizeAmount(ctx context.Context, amount float64) (*big.Int, error) {
	currency, err := erc20.Get(ctx)
	if err != nil {
		return nil, err
	}

	return parseUnits(amount, currency.Decimals)
}