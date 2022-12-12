package thirdweb

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// The nft drop encoder class is used to get the unsigned transaction data for nft drop contract
// contract calls that can be signed at a later time after generation.
//
// It can be accessed from the SDK through the `Encoder` namespace of the nft drop contract:
//
// You can access the NFTDrop interface from the SDK as follows:
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
//	// Number of NFTs to claim
//	quantity = 1
//
//	tx, err := contract.Encoder.ApproveClaimTo(context.Background(), signerAddress, quantity)
//
//	// Now you can get all the standard transaction data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *NFTDropEncoder) ApproveClaimTo(ctx context.Context, signerAddress string, quantity int) (*types.Transaction, error) {
	claimVerification, err := encoder.prepareClaim(ctx, quantity)
	if err != nil {
		return nil, err
	}

	return setErc20AllowanceEncoder(
		ctx,
		encoder.helper,
		signerAddress,
		claimVerification.Value,
		claimVerification.CurrencyAddress,
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
	active, err := encoder.claimConditions.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	claimVerification, err := encoder.prepareClaim(ctx, quantity)
	if err != nil {
		return nil, err
	}

	txOpts, err := encoder.helper.getUnsignedTxOptions(ctx, signerAddress)
	if err != nil {
		return nil, err
	}

	txOpts.Value = claimVerification.Value

	// Check for ERC20 Approval
	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	var pricePerToken *big.Int
	if claimVerification.Price.Cmp(MaxUint256) == 0 {
		pricePerToken = active.Price
	} else {
		pricePerToken = claimVerification.Price
	}

	var currencyAddress string
	if claimVerification.CurrencyAddress != zeroAddress {
		currencyAddress = claimVerification.CurrencyAddress
	} else {
		currencyAddress = active.CurrencyAddress
	}

	totalPrice := pricePerToken.Mul(big.NewInt(int64(quantity)), pricePerToken)
	if err := encoder.checkErc20Allowance(
		ctx,
		signerAddress,
		totalPrice,
		currencyAddress,
	); err != nil {
		return nil, err
	}

	proof := abi.IDropAllowlistProof{
		Proof:                  claimVerification.Proofs,
		QuantityLimitPerWallet: claimVerification.MaxClaimable,
		PricePerToken:          claimVerification.Price,
		Currency:               common.HexToAddress(claimVerification.CurrencyAddress),
	}

	return encoder.abi.Claim(
		txOpts,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(active.CurrencyAddress),
		active.Price,
		proof,
		[]byte{},
	)
}

func (encoder *NFTDropEncoder) prepareClaim(ctx context.Context, quantity int) (*ClaimVerification, error) {
	active, err := encoder.claimConditions.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	merkleMetadata, err := encoder.claimConditions.getMerkleMetadata(ctx)
	if err != nil {
		return nil, err
	}

	address := encoder.helper.getAddress().String()
	claimVerification, err := prepareClaim(
		ctx,
		address,
		quantity,
		active,
		merkleMetadata,
		encoder.helper,
		encoder.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}

func (encoder *NFTDropEncoder) checkErc20Allowance(
	ctx context.Context,
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
		allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
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
