package thirdweb

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// The nft drop encoder class is used to get the unsigned transaction data for nft drop contract
// contract calls that can be signed at a later time after generation.
//
// It can be accessed from the SDK through the `Encoder` namespace of the nft drop contract:
//
// You can access the NFTDrop interface from the SDK as follows:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	contract, err := sdk.GetNFTDrop("{{contract_address}}")
//
//	// Now the encoder can be accessed from the contract
//	contract.Encoder.ClaimTo(...)
type NFTDropEncoder struct {
	abi             *abi.DropERC721
	helper          *contractHelper
	storage         storage
	claimConditions *NFTDropClaimConditions
	*ContractEncoder
}

func newNFTDropEncoder(
	contractAbi *abi.DropERC721,
	helper *contractHelper,
	claimConditions *NFTDropClaimConditions,
	storage storage,
) (*NFTDropEncoder, error) {
	encoder, err := newContractEncoder(abi.DropERC721ABI, helper)
	if err != nil {
		return nil, err
	}

	return &NFTDropEncoder{
		abi:             contractAbi,
		helper:          helper,
		claimConditions: claimConditions,
		storage:         storage,
		ContractEncoder: encoder,
	}, nil
}

// Get the data for the transaction data required to approve the ERC20 token transfers
// necessary to claim NFTs from this contract.
//
// signerAddress: the address intended to sign the transaction
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction data of the token approval for the claim
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//	// Address of the wallet we want to claim the NFTs to
//	destinationAddress := "{{wallet_address}}"
//	// Number of NFTs to claim
//	quantity = 1
//
//	tx, err := contract.Encoder.ApproveClaimTo(context.Background(), signerAddress, destinationAddress, quantity)
//
//	// Now you can get all the standard transaction data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *NFTDropEncoder) ApproveClaimTo(ctx context.Context, signerAddress string, destinationAddress string, quantity int) (*types.Transaction, error) {
	claimVerification, err := encoder.prepareClaim(quantity)
	if err != nil {
		return nil, err
	}

	return setErc20AllowanceEncoder(
		ctx,
		encoder.helper,
		signerAddress,
		claimVerification.value,
		claimVerification.currencyAddress,
	)
}

// Get the data for the transaction required to claim NFTs from this contract.
//
// signerAddress: the address intended to sign the transaction
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction data of the claim
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//	// Address of the wallet we want to claim the NFTs to
//	destinationAddress := "{{wallet_address}}"
//	// Number of NFTs to claim
//	quantity = 1
//
//	tx, err := contract.Encoder.ClaimTo(context.Background(), signerAddress, destinationAddress, quantity)
//
//	// Now you can get all the standard transaction data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *NFTDropEncoder) ClaimTo(ctx context.Context, signerAddress string, destinationAddress string, quantity int) (*types.Transaction, error) {
	claimVerification, err := encoder.prepareClaim(quantity)
	if err != nil {
		return nil, err
	}

	txOpts, err := encoder.helper.getUnsignedTxOptions(ctx, signerAddress)
	if err != nil {
		return nil, err
	}

	txOpts.Value = claimVerification.value

	err = encoder.checkErc20Allowance(
		signerAddress,
		claimVerification.value,
		claimVerification.currencyAddress,
	)
	if err != nil {
		return nil, err
	}

	return encoder.abi.Claim(
		txOpts,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(claimVerification.currencyAddress),
		claimVerification.price,
		claimVerification.proofs,
		big.NewInt(int64(claimVerification.maxQuantityPerTransaction)),
	)
}

func (encoder *NFTDropEncoder) prepareClaim(quantity int) (*ClaimVerification, error) {
	claimCondition, err := encoder.claimConditions.GetActive()
	if err != nil {
		return nil, err
	}

	claimVerification := &ClaimVerification{
		proofs:                    [][32]byte{},
		maxQuantityPerTransaction: 0,
		price:                     claimCondition.Price,
		currencyAddress:           claimCondition.CurrencyAddress,
		value:                     big.NewInt(0).Mul(big.NewInt(int64(quantity)), claimCondition.Price),
	}

	return claimVerification, nil
}

func (encoder *NFTDropEncoder) checkErc20Allowance(
	signerAddress string,
	value *big.Int,
	currencyAddress string,
) error {
	if !isNativeToken(currencyAddress) {
		provider := encoder.helper.GetProvider()
		erc20, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
		if err != nil {
			return err
		}

		owner := common.HexToAddress(signerAddress)
		spender := encoder.helper.getAddress()
		allowance, err := erc20.Allowance(&bind.CallOpts{}, owner, spender)
		if err != nil {
			return err
		}

		if allowance.Cmp(value) < 0 {
			return fmt.Errorf(
				"NFT Drop contract '%s' has insufficient allowance to spend ERC20 token '%s' on behalf of the user wallet '%s' "+
					"Please approve the contract to spend tokens with the "+
					"'marketplace.Encoder.ApproveClaimTo(signerAddress, destinationAddress, quantity)' method",
				encoder.helper.getAddress().Hex(),
				currencyAddress,
				signerAddress,
			)
		}
	}

	return nil
}
