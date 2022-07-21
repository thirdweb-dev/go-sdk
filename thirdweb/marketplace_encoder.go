package thirdweb

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

// The marketplace encoder class is used to get the unsigned transaction data for marketplace contract
// contract calls that can be signed at a later time after generation.
//
// It can be accessed from the SDK through the `Encoder` namespace of the marketplace contract:
//
// You can access the Marketplace interface from the SDK as follows:
//
// 	import (
// 		"github.com/thirdweb-dev/go-sdk/thirdweb"
// 	)
//
// 	privateKey = "..."
//
// 	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
// 	})
//
//	marketplace, err := sdk.GetMarketplace("{{contract_address}}")
//
// 	// Now the encoder can be accessed from the contract
//  marketplace.Encoder.CreateListing(...)
type MarketplaceEncoder struct {
	abi     *abi.Marketplace
	helper  *contractHelper
	storage storage
}

func newMarketplaceEncoder(abi *abi.Marketplace, helper *contractHelper, storage storage) *MarketplaceEncoder {
	return &MarketplaceEncoder{
		abi:     abi,
		helper:  helper,
		storage: storage,
	}
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
//  // Address of the wallet we expect to sign this message
// 	signerAddress := "0x..."
//  // ID of the listing to cancel
//  listingId := 1
//
//  // Transaction data required for this request
// 	tx, err := marketplace.Encoder.CancelListing(signerAddress, listingId)
//
//  // Now you can get transaction all the standard data as needed
//  fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
//  fmt.Println(tx.Nonce())
func (encoder *MarketplaceEncoder) CancelListing(signerAddress string, listingId int) (*types.Transaction, error) {
	txOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
	if err != nil {
		return nil, err
	}

	return encoder.abi.CancelDirectListing(txOpts, big.NewInt(int64(listingId)))
}

// Get the data for the transactions required to buyout a direct listing from the marketplace, including any necessary token
// approvals to approve the spending of any tokens.
//
// signerAddress: the address intended to sign the transaction
//
// listingId: the ID of the listing to buyout
//
// quantityDesired: the quantity of the listed tokens to purchase
//
// receiver: the address to receive the purchased tokens
//
// returns: the list of transaction data for the transactions that need to be made to make this purchase
//
// Example
//
//  // Address of the wallet we expect to sign this message
// 	signerAddress := "0x..."
//  // ID of the listing to buyout
//  listingId := 1
//  // Quantity of the listed tokens to purchase
//  quantityDesired := 1
//  // receiver address to receive the purchased tokens
//  receiver := "0x..."
//
//  // Transaction data required for this request
// 	txs, err := marketplace.Encoder.BuyoutListing(signerAddress, listingId, quantityDesired, receiver)
//
//  // Now you can get transaction all the standard data as needed
// 	firstTx := txs[0]
//  fmt.Println(firstTx.Data()) // Ex: get the data field or the nonce field (others are available)
//  fmt.Println(firstTx.Nonce())
func (encoder *MarketplaceEncoder) BuyoutListing(
	signerAddress string,
	listingId int,
	quantityDesired int,
	receiver string,
) ([]*types.Transaction, error) {
	transactions := []*types.Transaction{}

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
	value := big.NewInt(int64(listing.BuyoutPrice)).Mul(big.NewInt(int64(listing.BuyoutPrice)), quantity)

	txOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
	if err != nil {
		return nil, err
	}

	tx, err := encoder.setErc20Allowance(
		signerAddress,
		value,
		listing.CurrencyContractAddress,
		txOpts,
	)
	if err != nil {
		return nil, err
	}
	if tx != nil {
		transactions = append(transactions, tx)
	}

	tx, err = encoder.abi.Buy(
		txOpts,
		big.NewInt(int64(listingId)),
		common.HexToAddress(receiver),
		big.NewInt(int64(quantityDesired)),
		common.HexToAddress(listing.CurrencyContractAddress),
		value,
	)
	if err != nil {
		return nil, err
	}

	transactions = append(transactions, tx)
	return transactions, nil
}

// Get the data for the transactions required to create a new direct listing, including any necessary token
// approvals to approve token transfers for the marketplace.
//
// signerAddress: the address intended to sign the transaction
//
// listing: the parameters for the new direct listing to create
//
// returns: the list of transaction data for the transactions that need to be made to create this listing
//
// Example
//
//  // Address of the wallet we expect to sign this message
// 	signerAddress := "0x..."
//
// 	listing := &NewDirectListing{
// 		AssetContractAddress: "0x...", // Address of the asset contract
// 		TokenId: 0, // Token ID of the asset to list
// 		StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
// 		ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
// 		Quantity: 1, // Quantity of the asset to list
// 		CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
// 		BuyoutPricePerToken: 1, // Price per token of the asset to list
// 	}
//
//  // Transaction data required for this request
// 	txs, err := marketplace.Encoder.CreateListing(signerAddress, listing)
//
//  // Now you can get transaction data as needed
// 	firstTx := txs[0]
//  fmt.Println(firstTx.Data()) // Ex: get the data field or the nonce field (others are available)
//  fmt.Println(firstTx.Nonce())
func (encoder *MarketplaceEncoder) CreateListing(signerAddress string, listing *NewDirectListing) ([]*types.Transaction, error) {
	listing.fillDefaults()
	transactions := []*types.Transaction{}

	tx, err := encoder.handleTokenApproval(
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
	if tx != nil {
		transactions = append(transactions, tx)
	}

	normalizedPricePerToken, err := normalizePriceValue(
		encoder.helper.GetProvider(),
		listing.BuyoutPricePerToken,
		listing.CurrencyContractAddress,
	)
	if err != nil {
		return nil, err
	}

	txOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
	if err != nil {
		return nil, err
	}

	tx, err = encoder.abi.CreateListing(txOpts, abi.IMarketplaceListingParameters{
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
	if err != nil {
		return nil, err
	}

	transactions = append(transactions, tx)
	return transactions, nil
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

func (encoder *MarketplaceEncoder) setErc20Allowance(
	signerAddress string,
	value *big.Int,
	currencyAddress string,
	txOpts *bind.TransactOpts,
) (*types.Transaction, error) {
	if isNativeToken(currencyAddress) {
		txOpts.Value = value
		return nil, nil
	} else {
		provider := encoder.helper.GetProvider()
		erc20, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
		if err != nil {
			return nil, err
		}

		owner := common.HexToAddress(signerAddress)
		spender := encoder.helper.getAddress()
		allowance, err := erc20.Allowance(&bind.CallOpts{}, owner, spender)
		if err != nil {
			return nil, err
		}

		if allowance.Cmp(value) < 0 {
			// We can get options from the contract instead of ERC20 because they will be the same
			approvalOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
			if err != nil {
				return nil, err
			}

			return erc20.Approve(approvalOpts, spender, value)
		}

		return nil, nil
	}
}

func (encoder *MarketplaceEncoder) handleTokenApproval(
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

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return nil, err
		}

		if !approved {
			tokenApproved, err := contract.GetApproved(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
			if err != nil {
				return nil, err
			}

			txOpts, err := helper.getUnsignedTxOptions(signerAddress)
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
			txOpts, err := helper.getUnsignedTxOptions(signerAddress)
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
