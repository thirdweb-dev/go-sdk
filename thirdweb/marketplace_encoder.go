package thirdweb

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// The marketplace encoder class is used to get the unsigned transaction data for marketplace contract
// contract calls that can be signed at a later time after generation.
//
// It can be accessed from the SDK through the `Encoder` namespace of the marketplace contract:
//
// You can access the Marketplace interface from the SDK as follows:
//
//		import (
//			"github.com/thirdweb-dev/go-sdk/thirdweb"
//		)
//
//		privateKey = "..."
//
//		sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//			PrivateKey: privateKey,
//		})
//
//		marketplace, err := sdk.GetMarketplace("{{contract_address}}")
//
//		// Now the encoder can be accessed from the contract
//	 marketplace.Encoder.CreateListing(...)
type MarketplaceEncoder struct {
	abi     *abi.Marketplace
	helper  *contractHelper
	storage storage
	*ContractEncoder
}

func newMarketplaceEncoder(contractAbi *abi.Marketplace, helper *contractHelper, storage storage) (*MarketplaceEncoder, error) {
	encoder, err := newContractEncoder(abi.MarketplaceABI, helper)
	if err != nil {
		return nil, err
	}

	return &MarketplaceEncoder{
		abi:             contractAbi,
		helper:          helper,
		storage:         storage,
		ContractEncoder: encoder,
	}, nil
}

// Get the data for the transaction required to cancel a listing on the marketplace
//
// signerAddress: the address intended to sign the transaction
//
// listingId: the ID of the listing to cancel
//
// returns: the transaction data for the cancellation
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//	// ID of the listing to cancel
//	listingId := 1
//
//	// Transaction data required for this request
//	tx, err := marketplace.Encoder.CancelListing(context.Background(), signerAddress, listingId)
//
//	// Now you can get transaction all the standard data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *MarketplaceEncoder) CancelListing(ctx context.Context, signerAddress string, listingId int) (*types.Transaction, error) {
	txOpts, err := encoder.helper.getUnsignedTxOptions(ctx, signerAddress)
	if err != nil {
		return nil, err
	}

	return encoder.abi.CancelDirectListing(txOpts, big.NewInt(int64(listingId)))
}

// Get the transaction data to approve the tokens needed for a buyout listing transaction. If the transaction
// wouldn't require any additional token approvals, this method will return nil.
//
// signerAddress: the address intended to sign the transaction
//
// listingId: the ID of the listing to buyout
//
// quantityDesired: the quantity of the listed tokens to purchase
//
// receiver: the address to receive the purchased tokens
//
// returns: the transaction data for the token approval if an approval is needed, or nil
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//	// ID of the listing to buyout
//	listingId := 1
//	// Quantity of the listed tokens to purchase
//	quantityDesired := 1
//	// receiver address to receive the purchased tokens
//	receiver := "0x..."
//
//	// Transaction data required for this request
//	tx, err := marketplace.Encoder.ApproveBuyoutListing(context.Background(), signerAddress, listingId, quantityDesired, receiver)
//
//	// Now you can get transaction all the standard data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *MarketplaceEncoder) ApproveBuyoutListing(
	ctx context.Context,
	signerAddress string,
	listingId int,
	quantityDesired int,
	receiver string,
) (*types.Transaction, error) {
	listing, err := encoder.validateListing(listingId)
	if err != nil {
		return nil, err
	}

	valid, err := isStillValidListing(encoder.helper, listing, quantityDesired)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("The asset on this listing has been moved from the lister's wallet, this listing is now invalid")
	}

	quantity := big.NewInt(int64(quantityDesired))
	value := listing.BuyoutCurrencyValuePerToken.Value.Mul(listing.BuyoutCurrencyValuePerToken.Value, quantity)

	return setErc20AllowanceEncoder(
		ctx,
		encoder.helper,
		signerAddress,
		value,
		listing.CurrencyContractAddress,
	)
}

// Get the transaction data for the buyout listing transaction. This method will throw an error if the listing requires
// payment in ERC20 tokens and the ERC20 tokens haven't yet been approved by the spender. You can get the
// transaction data of this required approval transaction from the `ApproveBuyoutListing` method.
//
// signerAddress: the address intended to sign the transaction
//
// listingId: the ID of the listing to buyout
//
// quantityDesired: the quantity of the listed tokens to purchase
//
// receiver: the address to receive the purchased tokens
//
// returns: the transaction data for this purchase
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//	// ID of the listing to buyout
//	listingId := 1
//	// Quantity of the listed tokens to purchase
//	quantityDesired := 1
//	// receiver address to receive the purchased tokens
//	receiver := "0x..."
//
//	// Transaction data required for this request
//	tx, err := marketplace.Encoder.BuyoutListing(context.Background(), signerAddress, listingId, quantityDesired, receiver)
//
//	// Now you can get transaction all the standard data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *MarketplaceEncoder) BuyoutListing(
	ctx context.Context,
	signerAddress string,
	listingId int,
	quantityDesired int,
	receiver string,
) (*types.Transaction, error) {
	listing, err := encoder.validateListing(listingId)
	if err != nil {
		return nil, err
	}

	valid, err := isStillValidListing(encoder.helper, listing, quantityDesired)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, errors.New("The asset on this listing has been moved from the lister's wallet, this listing is now invalid")
	}

	quantity := big.NewInt(int64(quantityDesired))
	value := listing.BuyoutCurrencyValuePerToken.Value.Mul(listing.BuyoutCurrencyValuePerToken.Value, quantity)

	err = encoder.checkErc20Allowance(
		signerAddress,
		value,
		listing.CurrencyContractAddress,
	)
	if err != nil {
		return nil, err
	}

	txOpts, err := encoder.helper.getUnsignedTxOptions(ctx, signerAddress)
	return encoder.abi.Buy(
		txOpts,
		big.NewInt(int64(listingId)),
		common.HexToAddress(receiver),
		big.NewInt(int64(quantityDesired)),
		common.HexToAddress(listing.CurrencyContractAddress),
		value,
	)
}

// Get the transaction data to approve the tokens needed for a create liting transaction. If the transaction
// wouldn't require any additional token approvals, this method will return nil.
//
// signerAddress: the address intended to sign the transaction
//
// listing: the parameters for the new direct listing to create
//
// returns: the transaction data for the create listing transaction
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//
//	listing := &NewDirectListing{
//		AssetContractAddress: "0x...", // Address of the asset contract
//		TokenId: 0, // Token ID of the asset to list
//		StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
//		ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
//		Quantity: 1, // Quantity of the asset to list
//		CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
//		BuyoutPricePerToken: 1, // Price per token of the asset to list
//	}
//
//	// Transaction data required for this request
//	tx, err := marketplace.Encoder.ApproveCreateListing(context.Background(), signerAddress, listing)
//
//	// Now you can get transaction data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *MarketplaceEncoder) ApproveCreateListing(ctx context.Context, signerAddress string, listing *NewDirectListing) (*types.Transaction, error) {
	listing.fillDefaults()
	return encoder.handleTokenApproval(
		ctx,
		signerAddress,
		encoder.helper.GetProvider(),
		encoder.helper,
		encoder.helper.getAddress().Hex(),
		listing.AssetContractAddress,
		listing.TokenId,
		encoder.helper.GetSignerAddress().Hex(),
	)
}

// Get the data for the transaction required to create a direct listing. This method will throw an error if
// the tokens needed for the listing have not yet been approved by the asset owner for the marketplace to spend.
// You can get the transaction data of this required approval transaction from the `ApproveCreateListing` method.
//
// signerAddress: the address intended to sign the transaction
//
// listing: the parameters for the new direct listing to create
//
// returns: the transaction data for the create listing transaction
//
// Example
//
//	// Address of the wallet we expect to sign this message
//	signerAddress := "0x..."
//
//	listing := &NewDirectListing{
//		AssetContractAddress: "0x...", // Address of the asset contract
//		TokenId: 0, // Token ID of the asset to list
//		StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
//		ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
//		Quantity: 1, // Quantity of the asset to list
//		CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
//		BuyoutPricePerToken: 1, // Price per token of the asset to list
//	}
//
//	// Transaction data required for this request
//	tx, err := marketplace.Encoder.CreateListing(context.Background(), signerAddress, listing)
//
//	// Now you can get transaction data as needed
//	fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//	fmt.Println(tx.Nonce())
func (encoder *MarketplaceEncoder) CreateListing(ctx context.Context, signerAddress string, listing *NewDirectListing) (*types.Transaction, error) {
	listing.fillDefaults()

	err := encoder.checkTokenApproval(
		signerAddress,
		encoder.helper.GetProvider(),
		encoder.helper,
		encoder.helper.getAddress().Hex(),
		listing.AssetContractAddress,
		listing.TokenId,
		encoder.helper.GetSignerAddress().Hex(),
	)
	if err != nil {
		return nil, err
	}

	normalizedPricePerToken, err := normalizePriceValue(
		encoder.helper.GetProvider(),
		listing.BuyoutPricePerToken,
		listing.CurrencyContractAddress,
	)
	if err != nil {
		return nil, err
	}

	txOpts, err := encoder.helper.getUnsignedTxOptions(ctx, signerAddress)
	if err != nil {
		return nil, err
	}

	return encoder.abi.CreateListing(txOpts, abi.IMarketplaceListingParameters{
		AssetContract:        common.HexToAddress(listing.AssetContractAddress),
		TokenId:              big.NewInt(int64(listing.TokenId)),
		StartTime:            big.NewInt(int64(listing.StartTimeInEpochSeconds)),
		SecondsUntilEndTime:  big.NewInt(int64(listing.ListingDurationInSeconds)),
		QuantityToList:       big.NewInt(int64(listing.Quantity)),
		CurrencyToAccept:     common.HexToAddress(listing.CurrencyContractAddress),
		ReservePricePerToken: normalizedPricePerToken,
		BuyoutPricePerToken:  normalizedPricePerToken,
		ListingType:          0,
	})
}

func (encoder *MarketplaceEncoder) validateListing(listingId int) (*DirectListing, error) {
	listing, err := encoder.abi.Listings(&bind.CallOpts{}, big.NewInt(int64(listingId)))
	if err != nil {
		return nil, err
	}

	if listing.AssetContract.String() == zeroAddress {
		return nil, fmt.Errorf("Failed to find listing with ID %d", listingId)
	}

	if listing.ListingType == 0 {
		return mapListing(encoder.helper, encoder.storage, listing)
	} else {
		return nil, fmt.Errorf("Unknown listing type: %d", listingId)
	}
}

func (encoder *MarketplaceEncoder) checkErc20Allowance(
	signerAddress string,
	value *big.Int,
	currencyAddress string,
) error {
	if isNativeToken(currencyAddress) {
		return nil
	} else {
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
				"Marketplace contract '%s' has insufficient allowance to spend ERC20 token '%s' on behalf of the user wallet '%s' "+
					"Please approve the contract to spend tokens with the "+
					"'marketplace.Encoder.ApproveBuyoutListing(signerAddress, listingId, quantityDesired, receiver)' method",
				encoder.helper.getAddress().Hex(),
				currencyAddress,
				signerAddress,
			)
		}

		return nil
	}
}

func (encoder *MarketplaceEncoder) checkTokenApproval(
	signerAddress string,
	provider *ethclient.Client,
	helper *contractHelper,
	marketplaceAddress string,
	assetContract string,
	tokenId int,
	from string,
) error {
	erc165, err := abi.NewIERC165(common.HexToAddress(assetContract), provider)
	if err != nil {
		return err
	}

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return err
	}

	if isErc721 {
		contract, err := abi.NewTokenERC721(common.HexToAddress(assetContract), provider)
		if err != nil {
			return err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return err
		}

		if !approved {
			tokenApproved, err := contract.GetApproved(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
			if err != nil {
				return err
			}

			if strings.ToLower(tokenApproved.String()) != strings.ToLower(marketplaceAddress) {
				return fmt.Errorf(
					"Marketplace contract '%s' doesn't have approval to transfer ERC721 token '%s' with ID %d from the user wallet '%s' "+
						"Please approve the contract to transfer tokens with the "+
						"'marketplace.Encoder.ApproveCreateListing(signerAddress, listing)' method",
					helper.getAddress().Hex(),
					assetContract,
					tokenId,
					signerAddress,
				)
			}
		}
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(common.HexToAddress(assetContract), provider)
		if err != nil {
			return err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return err
		}

		if !approved {
			return fmt.Errorf(
				"Marketplace contract '%s' doesn't have approval to transfer ERC1155 token '%s' with ID %d from the user wallet '%s' "+
					"Please approve the contract to transfer tokens with the "+
					"'marketplace.Encoder.ApproveCreateListing(signerAddress, listing)' method",
				helper.getAddress().Hex(),
				assetContract,
				tokenId,
				signerAddress,
			)
		}
	} else {
		return errors.New("Contract does not implement ERC721 or ERC1155.")
	}

	return nil
}

func (encoder *MarketplaceEncoder) handleTokenApproval(
	ctx context.Context,
	signerAddress string,
	provider *ethclient.Client,
	helper *contractHelper,
	marketplaceAddress string,
	assetContract string,
	tokenId int,
	from string,
) (*types.Transaction, error) {
	erc165, err := abi.NewIERC165(common.HexToAddress(assetContract), provider)
	if err != nil {
		return nil, err
	}

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return nil, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return nil, err
	}

	if isErc721 {
		contract, err := abi.NewTokenERC721(common.HexToAddress(assetContract), provider)
		if err != nil {
			return nil, err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{
			Context: ctx,
		}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return nil, err
		}

		if !approved {
			tokenApproved, err := contract.GetApproved(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
			if err != nil {
				return nil, err
			}

			txOpts, err := helper.getUnsignedTxOptions(ctx, signerAddress)
			if err != nil {
				return nil, err
			}

			if strings.ToLower(tokenApproved.String()) != strings.ToLower(marketplaceAddress) {
				return contract.SetApprovalForAll(txOpts, common.HexToAddress(marketplaceAddress), true)
			}
		}
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(common.HexToAddress(assetContract), provider)
		if err != nil {
			return nil, err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return nil, err
		}

		if !approved {
			txOpts, err := helper.getUnsignedTxOptions(ctx, signerAddress)
			if err != nil {
				return nil, err
			}

			return contract.SetApprovalForAll(txOpts, common.HexToAddress(marketplaceAddress), true)
		}
	} else {
		return nil, errors.New("Contract does not implement ERC721 or ERC1155")
	}

	return nil, nil
}
