package thirdweb

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// You can access the Token interface from the SDK as follows:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	contract, err := sdk.GetToken("{{contract_address}}")
type Token struct {
	*ERC20Standard
	abi     *abi.TokenERC20
	Helper  *contractHelper
	Encoder *ContractEncoder
	Events  *ContractEvents
}

func newToken(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Token, error) {
	if contractAbi, err := abi.NewTokenERC20(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		if erc20, err := newERC20Standard(provider, address, privateKey, storage); err != nil {
			return nil, err
		} else {
			encoder, err := newContractEncoder(abi.TokenERC20ABI, helper)
			if err != nil {
				return nil, err
			}

			events, err := newContractEvents(abi.TokenERC20ABI, helper)
			if err != nil {
				return nil, err
			}

			token := &Token{
				erc20,
				contractAbi,
				helper,
				encoder,
				events,
			}
			return token, nil
		}
	}
}

// Get the connected wallets voting power in this token.
//
// returns: vote balance of the connected wallet
func (token *Token) GetVoteBalance(ctx context.Context) (*CurrencyValue, error) {
	return token.GetVoteBalanceOf(ctx, token.Helper.GetSignerAddress().String())
}

// Get the voting power of the specified wallet in this token.
//
// address: wallet address to check the vote balance of
//
// returns: vote balance of the specified wallet
func (token *Token) GetVoteBalanceOf(ctx context.Context, address string) (*CurrencyValue, error) {
	votes, err := token.abi.GetVotes(&bind.CallOpts{Context: ctx}, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	return token.erc20.getValue(ctx, votes)
}

// Get the connected wallets delegatee address for this token.
//
// returns: delegation address of the connected wallet
func (token *Token) GetDelegation(ctx context.Context) (string, error) {
	return token.GetDelegationOf(ctx, token.Helper.GetSignerAddress().String())
}

// Get a specified wallets delegatee for this token.
//
// returns: delegation address of the connected wallet
func (token *Token) GetDelegationOf(ctx context.Context, address string) (string, error) {
	delegation, err := token.abi.Delegates(&bind.CallOpts{Context: ctx}, common.HexToAddress(address))
	if err != nil {
		return "", err
	}

	return delegation.String(), nil
}

// Mint tokens to the connected wallet.
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
func (token *Token) Mint(ctx context.Context, amount float64) (*types.Transaction, error) {
	return token.erc20.Mint(ctx, amount)
}

// Mint tokens to a specified wallet.
//
// to: wallet address to mint tokens to
//
// amount: amount of tokens to mint
//
// returns: transaction receipt of the mint
//
// Example
//
//	tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", 1)
func (token *Token) MintTo(ctx context.Context, to string, amount float64) (*types.Transaction, error) {
	return token.erc20.MintTo(ctx, to, amount)
}

// Mint tokens to a list of wallets.
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
//	tx, err := contract.MintBatchTo(context.Background(), args)
func (token *Token) MintBatchTo(ctx context.Context, args []*TokenAmount) (*types.Transaction, error) {
	return token.erc20.MintBatchTo(ctx, args)
}

// Delegate the connected wallets tokens to a specified wallet.
//
// delegateeAddress: wallet address to delegate tokens to
//
// returns: transaction receipt of the delegation
func (token *Token) DelegateTo(ctx context.Context, delegatreeAddress string) (*types.Transaction, error) {
	txOpts, err := token.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := token.abi.Delegate(txOpts, common.HexToAddress(delegatreeAddress))
	if err != nil {
		return nil, err
	}

	return token.Helper.AwaitTx(ctx, tx.Hash())
}
