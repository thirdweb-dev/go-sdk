package thirdweb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type Token struct {
	abi    *abi.TokenERC20
	helper *contractHelper
	*ERC20
}

func newToken(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Token, error) {
	if contractAbi, err := abi.NewTokenERC20(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		if erc20, err := newERC20(provider, address, privateKey, storage); err != nil {
			return nil, err
		} else {
			token := &Token{
				contractAbi,
				helper,
				erc20,
			}
			return token, nil
		}
	}
}

// GetVoteBalance
//
// Get the connected wallets voting power in this token
//
// returns: vote balance of the connected wallet
func (token *Token) GetVoteBalance() (*CurrencyValue, error) {
	return token.GetVoteBalanceOf(token.helper.GetSignerAddress().String())
}

// GetVoteBalanceOf
//
// Get the voting power of the specified wallet in this token
//
// address: wallet address to check the vote balance of
//
// returns: vote balance of the specified wallet
func (token *Token) GetVoteBalanceOf(address string) (*CurrencyValue, error) {
	votes, err := token.abi.GetVotes(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return token.getValue(votes)
}

// GetDelegation
//
// Get the connected wallets delegatee address for this token.
//
// returns: delegation address of the connected wallet
func (token *Token) GetDelegation() (string, error) {
	return token.GetDelegationOf(token.helper.GetSignerAddress().String())
}

// GetDelegationOf
//
// Get a specified wallets delegatee for this token
//
// returns: delegation address of the connected wallet
func (token *Token) GetDelegationOf(address string) (string, error) {
	delegation, err := token.abi.Delegates(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return "", err
	}

	return delegation.String(), nil
}

// Mint
//
// Mint tokens to the connected wallet
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
func (token *Token) Mint(amount float64) (*types.Transaction, error) {
	return token.MintTo(token.helper.GetSignerAddress().String(), amount)
}

// MintTo
//
// Mint tokens to a specified wallet
//
// to: wallet address to mint tokens to
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
func (token *Token) MintTo(to string, amount float64) (*types.Transaction, error) {
	amountWithDecimals, err := token.normalizeAmount(amount)
	if err != nil {
		return nil, err
	}

	tx, err := token.abi.MintTo(token.helper.getTxOptions(), common.HexToAddress(to), amountWithDecimals)
	if err != nil {
		return nil, err
	}

	return token.helper.awaitTx(tx.Hash())
}

// MintBatchTo
//
// Mint tokens to a list of wallets
//
// args: list of wallet addresses and amounts to mint
//
// returns: transaction receipt of the mint
func (token *Token) MintBatchTo(args []*TokenAmount) (*types.Transaction, error) {
	encoded := [][]byte{}

	for _, arg := range args {
		amountWithDecimals, err := token.normalizeAmount(arg.Amount)
		if err != nil {
			return nil, err
		}

		tx, err := token.abi.MintTo(token.helper.getTxOptions(), common.HexToAddress(arg.ToAddress), amountWithDecimals)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := token.abi.Multicall(token.helper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return token.helper.awaitTx(tx.Hash())
}

// DelegateTo
//
// Delegate the connected wallets tokens to a specified wallet
//
// delegatreeAddress: wallet address to delegate tokens to
//
// returns: transaction receipt of the delegation
func (token *Token) DelegateTo(delegatreeAddress string) (*types.Transaction, error) {
	tx, err := token.abi.Delegate(token.helper.getTxOptions(), common.HexToAddress(delegatreeAddress))
	if err != nil {
		return nil, err
	}

	return token.helper.awaitTx(tx.Hash())
}
