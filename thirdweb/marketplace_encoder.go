package thirdweb

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

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

func (encoder *MarketplaceEncoder) CancelListing(signerAddress string, listingId int) ([]byte, error) {
	txOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
	if err != nil {
		return nil, err
	}

	tx, err := encoder.abi.CancelDirectListing(txOpts, big.NewInt(int64(listingId)))
	if err != nil {
		return nil, err
	}

	return tx.Data(), nil
}

func (encoder *MarketplaceEncoder) BuyoutListing(signerAddress string, listingId int, quantityDesired int, receiver string) ([]byte, error) {
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

	err = encoder.setErc20Allowance(
		value,
		listing.CurrencyContractAddress,
		txOpts,
	)
	if err != nil {
		return nil, err
	}
	tx, err := encoder.abi.Buy(
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

	return tx.Data(), nil
}

func (encoder *MarketplaceEncoder) HandleTokenApproval(signerAddress string, listing *NewDirectListing) ([]byte, error) {
	listing.fillDefaults()

	erc165, err := abi.NewIERC165(common.HexToAddress(listing.AssetContractAddress), encoder.helper.GetProvider())
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
		contract, err := abi.NewTokenERC721(common.HexToAddress(listing.AssetContractAddress), encoder.helper.GetProvider())
		if err != nil {
			return nil, err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(signerAddress), common.HexToAddress(encoder.helper.getAddress().Hex()))
		if err != nil {
			return nil, err
		}

		if !approved {
			tokenApproved, err := contract.GetApproved(&bind.CallOpts{}, big.NewInt(int64(listing.TokenId)))
			if err != nil {
				return nil, err
			}

			txOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
			if err != nil {
				return nil, err
			}

			if strings.ToLower(tokenApproved.String()) != strings.ToLower(encoder.helper.getAddress().Hex()) {
				tx, err := contract.SetApprovalForAll(txOpts, common.HexToAddress(encoder.helper.getAddress().Hex()), true)
				if err != nil {
					return nil, err
				}

				return tx.Data(), nil
			}
		}
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(common.HexToAddress(listing.AssetContractAddress), encoder.helper.GetProvider())
		if err != nil {
			return nil, err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(signerAddress), common.HexToAddress(encoder.helper.getAddress().Hex()))
		if err != nil {
			return nil, err
		}

		if !approved {
			txOpts, err := encoder.helper.getUnsignedTxOptions(signerAddress)
			if err != nil {
				return nil, err
			}

			tx, err := contract.SetApprovalForAll(txOpts, common.HexToAddress(encoder.helper.getAddress().Hex()), true)
			if err != nil {
				return nil, err
			}

			return tx.Data(), nil
		}
	} else {
		return nil, errors.New("Contract does not implement ERC721 or ERC1155")
	}

	return nil, nil
}

func (encoder *MarketplaceEncoder) CreateListing(listing *NewDirectListing) ([]byte, error) {
	return nil, nil
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
	value *big.Int,
	currencyAddress string,
	txOpts *bind.TransactOpts,
) error {
	if isNativeToken(currencyAddress) {
		txOpts.Value = value
		return nil
	} else {
		provider := encoder.helper.GetProvider()
		erc20, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
		if err != nil {
			return err
		}

		owner := encoder.helper.GetSignerAddress()
		spender := encoder.helper.getAddress()
		allowance, err := erc20.Allowance(&bind.CallOpts{}, owner, spender)
		if err != nil {
			return err
		}

		if allowance.Cmp(value) < 0 {
			return fmt.Errorf(
				"Contract does not have sufficient token allowance. Please approve the contract with address '%s' to spend tokens on behalf of the signer.",
				encoder.helper.getAddress().Hex(),
			)
		}

		return nil
	}
}
