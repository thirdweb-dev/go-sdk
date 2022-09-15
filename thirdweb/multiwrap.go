package thirdweb

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// You can access the Multiwrap interface from the SDK as follows:
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
//	contract, err := sdk.GetMultiwrap("{{contract_address}}")
type Multiwrap struct {
	abi    *abi.Multiwrap
	helper *contractHelper
	*ERC721
	Encoder *ContractEncoder
}

func newMultiwrap(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Multiwrap, error) {
	if contractAbi, err := abi.NewMultiwrap(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				encoder, err := newContractEncoder(abi.MultiwrapABI, helper)
				if err != nil {
					return nil, err
				}

				multiwrap := &Multiwrap{
					contractAbi,
					helper,
					erc721,
					encoder,
				}
				return multiwrap, nil
			}
		}
	}
}

// Get the contents of a wrapped token bundle.
//
// wrappedTokenId: the ID of the wrapped token bundle
//
// returns: the contents of the wrapped token bundle
//
// Example
//
//	tokenId := 0
//	contents, err := contract.GetWrappedContents(tokenId)
//	erc20Tokens := contents.Erc20Tokens
//	erc721Tokens := contents.Erc721Tokens
//	erc1155Tokens := contents.Erc1155Tokens
func (multiwrap *Multiwrap) GetWrappedContents(wrappedTokenId int) (*MultiwrapBundle, error) {
	wrappedTokens, err := multiwrap.abi.GetWrappedContents(&bind.CallOpts{}, big.NewInt(int64(wrappedTokenId)))
	if err != nil {
		return nil, err
	}

	log.Println("Tokens: ", len(wrappedTokens))

	erc20Tokens := []*MultiwrapERC20{}
	erc721Tokens := []*MultiwrapERC721{}
	erc1155Tokens := []*MultiwrapERC1155{}

	for _, wrappedToken := range wrappedTokens {
		switch wrappedToken.TokenType {
		case 0:
			tokenMetadata, err := fetchCurrencyMetadata(multiwrap.helper.GetProvider(), wrappedToken.AssetContract.String())
			if err != nil {
				return nil, err
			}

			erc20Tokens = append(erc20Tokens, &MultiwrapERC20{
				ContractAddress: wrappedToken.AssetContract.String(),
				Quantity:        formatUnits(wrappedToken.TotalAmount, tokenMetadata.Decimals),
			})
			continue
		case 1:
			erc721Tokens = append(erc721Tokens, &MultiwrapERC721{
				ContractAddress: wrappedToken.AssetContract.String(),
				TokenId:         int(wrappedToken.TokenId.Int64()),
			})
		case 2:
			erc1155Tokens = append(erc1155Tokens, &MultiwrapERC1155{
				ContractAddress: wrappedToken.AssetContract.String(),
				TokenId:         int(wrappedToken.TokenId.Int64()),
				Quantity:        int(wrappedToken.TotalAmount.Int64()),
			})
		}
	}

	tokens := &MultiwrapBundle{
		ERC20Tokens:   erc20Tokens,
		ERC721Tokens:  erc721Tokens,
		ERC1155Tokens: erc1155Tokens,
	}
	return tokens, nil
}

// Wrap any number of ERC20, ERC721, or ERC1155 tokens into a single wrapped token
//
// contents: the tokens to wrap into a single wrapped token
//
// wrappedTokenMetadata: the NFT Metadata or URI to as the metadata for the wrapped token
//
// recipientAddress: the optional address to send the wrapped token to
//
// returns: the transaction receipt of the wrapping
//
// Example
//
//	contents := &thirdweb.MultiwrapBundle{
//		ERC20Tokens: []*thirdweb.MultiwrapERC20{
//			&thirdweb.MultiwrapERC20{
//				ContractAddress: "0x...",
//				Quantity:        1,
//			},
//		},
//		ERC721Tokens: []*thirdweb.MultiwrapERC721{
//			&thirdweb.MultiwrapERC721{
//				ContractAddress: "0x...",
//				TokenId:         1,
//			},
//		},
//		ERC1155Tokens: []*thirdweb.MultiwrapERC1155{
//			&thirdweb.MultiwrapERC1155{
//				ContractAddress: "0x...",
//				TokenId:         1,
//				Quantity:        1,
//			},
//		},
//	}
//
//	wrappedTokenMetadata := &thirdweb.NFTMetadataInput{
//		Name: "Wrapped Token"
//	}
//
//	// This will mint the wrapped token to the connected wallet
//	tx, err := contract.Wrap(contents, wrappedTokenMetadata, "")
func (multiwrap *Multiwrap) Wrap(ctx context.Context, contents *MultiwrapBundle, wrappedTokenMetadata interface{}, recipientAddress string) (*types.Transaction, error) {
	uri, ok := wrappedTokenMetadata.(string)
	if !ok {
		tokenMetadata, ok := wrappedTokenMetadata.(*NFTMetadataInput)
		if ok {
			tokenUri, err := uploadOrExtractUri(tokenMetadata, multiwrap.storage)
			if err != nil {
				return nil, err
			}

			uri = tokenUri
		} else {
			return nil, errors.New("wrappedTokenMetadata must be a string or NFTMetadataInput")
		}
	}

	if recipientAddress == "" {
		recipientAddress = multiwrap.helper.GetSignerAddress().String()
	}

	tokens, err := multiwrap.toTokenStructList(contents)
	if err != nil {
		return nil, err
	}

	txOpts, err := multiwrap.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := multiwrap.abi.Wrap(txOpts, tokens, uri, common.HexToAddress(recipientAddress))
	if err != nil {
		return nil, err
	}

	return multiwrap.helper.awaitTx(tx.Hash())
}

// Unwrap a wrapped token bundle into its contents
//
// wrappedTokenId: the ID of the wrapped token bundle
//
// recipientAddress: the optional address to send the wrapped token to
//
// returns: the contents of the wrapped token bundle
//
// Example
//
//	tokenId := 0
//	tx, err := contract.Unwrap(context.Background(), tokenId, "")
func (multiwrap *Multiwrap) Unwrap(ctx context.Context, wrappedTokenId int, recipientAddress string) (*types.Transaction, error) {
	if recipientAddress == "" {
		recipientAddress = multiwrap.helper.GetSignerAddress().String()
	}

	txOpts, err := multiwrap.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := multiwrap.abi.Unwrap(txOpts, big.NewInt(int64(wrappedTokenId)), common.HexToAddress(recipientAddress))
	if err != nil {
		return nil, err
	}

	return multiwrap.helper.awaitTx(tx.Hash())
}

func (multiwrap *Multiwrap) toTokenStructList(contents *MultiwrapBundle) ([]abi.ITokenBundleToken, error) {
	tokens := []abi.ITokenBundleToken{}
	provider := multiwrap.helper.GetProvider()
	owner := multiwrap.helper.GetSignerAddress()

	for _, erc20 := range contents.ERC20Tokens {
		normalizedQuantity, err := normalizePriceValue(
			provider,
			erc20.Quantity,
			erc20.ContractAddress,
		)
		if err != nil {
			return nil, err
		}

		hasAllowance, err := hasErc20Allowance(
			multiwrap.helper,
			erc20.ContractAddress,
			normalizedQuantity,
		)

		if !hasAllowance {
			return nil, fmt.Errorf(
				fmt.Sprintf("ERC20 with contract address %v does not have enough allowance to transfer.", erc20.ContractAddress) +
					"You can set allowance to the multiwrap contract to transfer these tokens by running:\n" +
					fmt.Sprintf("token, _ := sdk.GetToken(\"%v\")\ntoken.SetAllowance(\"%v\", %f)", erc20.ContractAddress, owner.String(), erc20.Quantity),
			)
		}

		token := abi.ITokenBundleToken{
			TokenType:     0,
			AssetContract: common.HexToAddress(erc20.ContractAddress),
			TotalAmount:   normalizedQuantity,
			TokenId:       big.NewInt(0),
		}
		tokens = append(tokens, token)
	}

	for _, erc721 := range contents.ERC721Tokens {
		isApproved, err := isTokenApprovedForTransfer(
			provider,
			multiwrap.helper.getAddress().String(),
			erc721.ContractAddress,
			erc721.TokenId,
			owner.String(),
		)
		if err != nil {
			return nil, err
		}

		if !isApproved {
			return nil, fmt.Errorf(
				fmt.Sprintf("ERC721 with contract address %v does not have enough allowance to transfer.", erc721.ContractAddress) +
					"You can set allowance to the multiwrap contract to transfer this token by running:\n" +
					fmt.Sprintf("nft, _ := sdk.GetNFTCollection(\"%v\")\nnft.SetApprovalForToken(\"%v\", %d)", erc721.ContractAddress, multiwrap.helper.getAddress().String(), erc721.TokenId),
			)
		}

		token := abi.ITokenBundleToken{
			TokenType:     1,
			AssetContract: common.HexToAddress(erc721.ContractAddress),
			TotalAmount:   big.NewInt(0),
			TokenId:       big.NewInt(int64(erc721.TokenId)),
		}
		tokens = append(tokens, token)
	}

	for _, erc1155 := range contents.ERC1155Tokens {
		isApproved, err := isTokenApprovedForTransfer(
			provider,
			multiwrap.helper.getAddress().String(),
			erc1155.ContractAddress,
			erc1155.TokenId,
			owner.String(),
		)
		if err != nil {
			return nil, err
		}

		if !isApproved {
			return nil, fmt.Errorf(
				fmt.Sprintf("ERC1155 with contract address %v does not have enough allowance to transfer.", erc1155.ContractAddress) +
					"You can set allowance to the multiwrap contract to transfer this token by running:\n" +
					fmt.Sprintf("edition, _ := sdk.GetEdition(\"%v\")\nedition.SetApprovalForAll(\"%v\", true)", erc1155.ContractAddress, multiwrap.helper.getAddress().String()),
			)
		}

		token := abi.ITokenBundleToken{
			TokenType:     2,
			AssetContract: common.HexToAddress(erc1155.ContractAddress),
			TotalAmount:   big.NewInt(int64(erc1155.Quantity)),
			TokenId:       big.NewInt(int64(erc1155.TokenId)),
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}
