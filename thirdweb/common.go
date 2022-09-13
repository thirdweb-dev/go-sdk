package thirdweb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fxamacker/cbor"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// NFT

func fetchTokenMetadata(tokenId int, uri string, storage storage) (*NFTMetadata, error) {
	if body, err := storage.Get(uri); err != nil {
		return nil, err
	} else {
		metadata := &NFTMetadata{
			Id: big.NewInt(int64(tokenId)),
		}
		if err := json.Unmarshal(body, &metadata); err != nil {
			return nil, &unmarshalError{body: string(body), typeName: "nft", UnderlyingError: err}
		}

		return metadata, nil
	}
}

func uploadOrExtractUri(metadata *NFTMetadataInput, storage storage) (string, error) {
	metadataToUpload := map[string]interface{}{}
	mapstructure.Decode(metadata, &metadataToUpload)
	return storage.Upload(metadataToUpload, "", "")
}

func uploadOrExtractUris(metadatas []*NFTMetadataInput, storage storage) ([]string, error) {
	// Why is this necessary?
	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}

	dataToUpload := []map[string]interface{}{}
	mapstructure.Decode(data, &dataToUpload)

	baseUriWithUris, err := storage.UploadBatch(dataToUpload, 0, "", "")
	if err != nil {
		return nil, err
	}

	return baseUriWithUris.uris, nil
}

// TOKEN

func isNativeToken(tokenAddress string) bool {
	isZero := tokenAddress == zeroAddress
	isNative := strings.ToLower(tokenAddress) == nativeTokenAddress
	return isZero || isNative
}

func parseUnits(value float64, decimals int) (*big.Int, error) {
	floatValue := value * math.Pow(10, float64(decimals))
	bigNumber, ok := new(big.Int).SetString(fmt.Sprintf("%.0f", floatValue), 10)
	if !ok {
		return nil, fmt.Errorf("Failed to parse %f to big.Int", floatValue)
	}

	return bigNumber, nil
}

func normalizePriceValue(provider *ethclient.Client, price float64, currencyAddress string) (*big.Int, error) {
	metadata, err := fetchCurrencyMetadata(provider, currencyAddress)
	if err != nil {
		return nil, err
	}

	value, err := parseUnits(price, metadata.Decimals)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func formatUnits(value *big.Int, decimals int) float64 {
	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	num, _ := decimal.NewFromString(value.String())
	formatted := num.Div(mul)

	return formatted.InexactFloat64()
}

func fetchCurrencyMetadata(provider *ethclient.Client, asset string) (*Currency, error) {
	if isNativeToken(asset) {
		chainId, err := provider.ChainID(context.Background())
		if err != nil {
			return nil, err
		}
		nativeToken, err := getNativeTokenByChainId(ChainID(chainId.Int64()))
		currency := &Currency{
			nativeToken.name,
			nativeToken.symbol,
			nativeToken.decimals,
		}
		return currency, nil
	} else {
		contractAbi, err := abi.NewTokenERC20(common.HexToAddress(asset), provider)
		if err != nil {
			return nil, err
		}

		name, err := contractAbi.Name(&bind.CallOpts{})
		symbol, err := contractAbi.Symbol(&bind.CallOpts{})
		decimals, err := contractAbi.Decimals(&bind.CallOpts{})

		currency := &Currency{
			name,
			symbol,
			int(decimals),
		}

		return currency, nil
	}
}

func fetchCurrencyValue(provider *ethclient.Client, asset string, price *big.Int) (*CurrencyValue, error) {
	metadata, err := fetchCurrencyMetadata(provider, asset)
	if err != nil {
		return nil, err
	}

	displayValue := formatUnits(price, metadata.Decimals)

	currencyValue := &CurrencyValue{
		metadata.Name,
		metadata.Symbol,
		metadata.Decimals,
		price,
		displayValue,
	}
	return currencyValue, nil
}

func setErc20AllowanceEncoder(
	contractToApprove *contractHelper,
	signerAddress string,
	value *big.Int,
	currencyAddress string,
) (*types.Transaction, error) {
	if !isNativeToken(currencyAddress) {
		provider := contractToApprove.GetProvider()
		erc20, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
		if err != nil {
			return nil, err
		}

		owner := common.HexToAddress(signerAddress)
		spender := contractToApprove.getAddress()
		allowance, err := erc20.Allowance(&bind.CallOpts{}, owner, spender)
		if err != nil {
			return nil, err
		}

		if allowance.Cmp(value) < 0 {
			// We can get options from the contract instead of ERC20 because they will be the same
			approvalOpts, err := contractToApprove.getUnsignedTxOptions(signerAddress)
			if err != nil {
				return nil, err
			}

			return erc20.Approve(approvalOpts, spender, value)
		}
	}

	return nil, nil
}

func setErc20Allowance(
	contractToApprove *contractHelper,
	value *big.Int,
	currencyAddress string,
	txOpts *bind.TransactOpts,
) error {
	if isNativeToken(currencyAddress) {
		txOpts.Value = value
		return nil
	} else {
		provider := contractToApprove.GetProvider()
		erc20, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
		if err != nil {
			return err
		}

		owner := contractToApprove.GetSignerAddress()
		spender := contractToApprove.getAddress()
		allowance, err := erc20.Allowance(&bind.CallOpts{}, owner, spender)
		if err != nil {
			return err
		}

		if allowance.Cmp(value) < 0 {
			// We can get options from the contract instead of ERC20 because they will be the same
			approvalOpts, err := contractToApprove.getTxOptions()
			if err != nil {
				return err
			}

			tx, err := erc20.Approve(approvalOpts, spender, value)
			if err != nil {
				return err
			}

			_, err = contractToApprove.awaitTx(tx.Hash())
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func approveErc20Allowance(
	contractToApprove *contractHelper,
	currencyAddress string,
	price *big.Int,
	quantity int,
) error {
	provider := contractToApprove.GetProvider()
	contractAbi, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
	if err != nil {
		return err
	}

	erc20, err := newContractHelper(common.HexToAddress(currencyAddress), provider, "")
	if err != nil {
		return err
	}

	owner := contractToApprove.GetSignerAddress()
	spender := contractToApprove.getAddress()
	allowance, err := contractAbi.Allowance(&bind.CallOpts{}, owner, spender)
	if err != nil {
		return err
	}

	totalPrice := price.Mul(big.NewInt(int64(quantity)), price)

	if allowance.Cmp(totalPrice) < 0 {
		txOpts, err := erc20.getTxOptions()
		if err != nil {
			return err
		}
		contractAbi.Approve(txOpts, spender, allowance.Add(allowance, totalPrice))
	}

	return nil
}

func hasErc20Allowance(contractToApprove *contractHelper, currencyAddress string, value *big.Int) (bool, error) {
	if isNativeToken(currencyAddress) {
		return true, nil
	} else {
		provider := contractToApprove.GetProvider()
		contractAbi, err := abi.NewTokenERC20(common.HexToAddress(currencyAddress), provider)
		if err != nil {
			return false, err
		}

		owner := contractToApprove.GetSignerAddress()
		spender := contractToApprove.getAddress()
		allowance, err := contractAbi.Allowance(&bind.CallOpts{}, owner, spender)
		if err != nil {
			return false, err
		}

		return allowance.Cmp(value) >= 0, nil
	}
}

// DROP

func prepareClaim(
	quantity int,
	activeClaimCondition *ClaimConditionOutput,
	contractHelper *contractHelper,
	storage storage,
) (*ClaimVerification, error) {
	maxClaimable := 0
	price := activeClaimCondition.Price
	currencyAddress := activeClaimCondition.CurrencyAddress
	proofs := [][32]byte{}
	value := big.NewInt(0)

	if price.Cmp(big.NewInt(0)) > 0 {
		if isNativeToken(currencyAddress) {
			value = big.NewInt(0).Mul(big.NewInt(int64(quantity)), price)
		} else {
			approveErc20Allowance(
				contractHelper,
				currencyAddress,
				price,
				quantity,
			)
		}
	}

	claimVerification := &ClaimVerification{
		proofs:                    proofs,
		maxQuantityPerTransaction: maxClaimable,
		price:                     price,
		currencyAddress:           currencyAddress,
		value:                     value,
	}
	return claimVerification, nil
}

func transformResultToClaimCondition(
	pm *abi.IDropClaimConditionClaimCondition,
	provider *ethclient.Client,
	storage storage,
) (*ClaimConditionOutput, error) {
	currencyValue, err := fetchCurrencyValue(provider, pm.Currency.String(), pm.PricePerToken)
	if err != nil {
		return nil, err
	}

	return &ClaimConditionOutput{
		StartTime:                   pm.StartTimestamp,
		MaxQuantity:                 pm.MaxClaimableSupply,
		AvailableSupply:             big.NewInt(0).Sub(pm.MaxClaimableSupply, pm.SupplyClaimed),
		QuantityLimitPerTransaction: pm.QuantityLimitPerTransaction,
		WaitInSeconds:               pm.WaitTimeInSecondsBetweenClaims,
		Price:                       pm.PricePerToken,
		CurrencyAddress:             pm.Currency.String(),
		CurrencyMetadata:            currencyValue,
	}, nil
}

// MULTIWRAP

func isTokenApprovedForTransfer(
	provider *ethclient.Client,
	transferrerContractAddress string,
	assetContract string,
	tokenId int,
	from string,
) (bool, error) {
	erc165, err := abi.NewIERC165(common.HexToAddress(assetContract), provider)
	if err != nil {
		return false, err
	}

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return false, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return false, err
	}

	if isErc721 {
		ierc721, err := abi.NewIERC721(common.HexToAddress(assetContract), provider)
		if err != nil {
			return false, err
		}

		approved, err := ierc721.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(from), common.HexToAddress(transferrerContractAddress))
		if err != nil {
			return false, err
		}

		if approved {
			return true, nil
		}

		address, err := ierc721.GetApproved(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
		if err != nil {
			return false, err
		}

		return strings.ToLower(address.String()) == strings.ToLower(transferrerContractAddress), nil
	} else if isErc1155 {
		ierc1155, err := abi.NewIERC1155(common.HexToAddress(assetContract), provider)
		if err != nil {
			return false, err
		}

		return ierc1155.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(from), common.HexToAddress(transferrerContractAddress))
	} else {
		fmt.Println("Contract does not implement ERC1155 or ERC721")
		return false, nil
	}
}

// CUSTOM CONTRACTS

func fetchContractMetadataFromAddress(address string, provider *ethclient.Client, storage storage) (string, error) {
	metadataUri, err := resolveContractUriFromAddress(address, provider)
	if err != nil {
		return "", err
	}

	metadata, err := fetchContractMetadata(metadataUri, storage)
	if err != nil {
		return "", err
	}

	return metadata, nil
}

func resolveContractUriFromAddress(address string, provider *ethclient.Client) (string, error) {
	bytecode, err := provider.CodeAt(context.Background(), common.HexToAddress(address), nil)
	if err != nil {
		return "", err
	}

	return extractIPFSHashFromBytecode(bytecode)
}

func extractIPFSHashFromBytecode(bytecode []byte) (string, error) {
	cborLength := int(bytecode[len(bytecode)-2])*0x100 + int(bytecode[len(bytecode)-1])
	cborBytecode := bytecode[len(bytecode)-cborLength-2 : len(bytecode)-2]

	cborData := map[string][]byte{}
	if err := cbor.Unmarshal(cborBytecode, &cborData); err != nil {
		return "", err
	}

	ipfsHash := base58.Encode(cborData["ipfs"])
	return fmt.Sprintf("ipfs://%v", ipfsHash), nil
}

func fetchContractMetadata(uri string, storage storage) (string, error) {
	body, err := storage.Get(uri)
	if err != nil {
		return "", err
	}

	metadata := map[string]interface{}{}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return "", err
	}

	abiBytes, err := json.Marshal(metadata["output"].(map[string]interface{})["abi"])
	if err != nil {
		return "", err
	}

	return string(abiBytes), nil
}

// MARKETPLACE

func fetchTokenMetadataForContract(
	contractAddress string,
	provider *ethclient.Client,
	tokenId int,
	storage storage,
) (*NFTMetadata, error) {
	erc165, err := abi.NewIERC165(common.HexToAddress(contractAddress), provider)
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

	uri := ""
	if isErc721 {
		contract, err := abi.NewTokenERC721(common.HexToAddress(contractAddress), provider)
		if err != nil {
			return nil, err
		}

		uri, err = contract.TokenURI(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(common.HexToAddress(contractAddress), provider)
		if err != nil {
			return nil, err
		}

		uri, err = contract.Uri(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	}

	return fetchTokenMetadata(tokenId, uri, storage)
}

func handleTokenApproval(
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

			txOpts, err := helper.getTxOptions()
			if err != nil {
				return err
			}

			if strings.ToLower(tokenApproved.String()) != strings.ToLower(marketplaceAddress) {
				tx, err := contract.SetApprovalForAll(txOpts, common.HexToAddress(marketplaceAddress), true)
				if err != nil {
					return err
				}

				_, err = helper.awaitTx(tx.Hash())
				if err != nil {
					return err
				}
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
			txOpts, err := helper.getTxOptions()
			if err != nil {
				return err
			}

			tx, err := contract.SetApprovalForAll(txOpts, common.HexToAddress(marketplaceAddress), true)
			if err != nil {
				return err
			}

			_, err = helper.awaitTx(tx.Hash())
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("Contract does not implement ERC721 or ERC1155")
	}

	return nil
}

func isStillValidListing(
	helper *contractHelper,
	listing *DirectListing,
	quantity int,
) (bool, error) {
	approved, err := isTokenApprovedForTransfer(
		helper.GetProvider(),
		helper.getAddress().Hex(),
		listing.AssetContractAddress,
		listing.TokenId,
		listing.SellerAddress,
	)
	if err != nil {
		return false, err
	}

	if !approved {
		return false, nil
	}

	erc165, err := abi.NewIERC165(common.HexToAddress(listing.AssetContractAddress), helper.GetProvider())
	if err != nil {
		return false, err
	}

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return false, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return false, err
	}

	if isErc721 {
		contract, err := abi.NewTokenERC721(
			common.HexToAddress(listing.AssetContractAddress),
			helper.GetProvider(),
		)
		if err != nil {
			return false, err
		}

		ownerOf, err := contract.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(listing.TokenId)))
		if err != nil {
			return false, err
		}

		return strings.ToLower(ownerOf.Hex()) == strings.ToLower(listing.SellerAddress), nil
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(
			common.HexToAddress(listing.AssetContractAddress),
			helper.GetProvider(),
		)
		if err != nil {
			return false, err
		}

		balance, err := contract.BalanceOf(
			&bind.CallOpts{},
			common.HexToAddress(listing.SellerAddress),
			big.NewInt(int64(listing.TokenId)),
		)

		return balance.Int64() >= int64(quantity), nil
	} else {
		return false, errors.New("Contract does not implement ERC721 or ERC1155")
	}
}

func mapListing(
	helper *contractHelper,
	storage storage,
	listing abi.IMarketplaceListing,
) (*DirectListing, error) {
	currencyValue, err := fetchCurrencyValue(
		helper.GetProvider(),
		listing.Currency.String(),
		listing.BuyoutPricePerToken,
	)
	if err != nil {
		return nil, err
	}

	asset, err := fetchTokenMetadataForContract(
		listing.AssetContract.String(),
		helper.GetProvider(),
		int(listing.TokenId.Int64()),
		storage,
	)
	if err != nil {
		return nil, err
	}

	return &DirectListing{
		AssetContractAddress:        listing.AssetContract.String(),
		BuyoutPrice:                 listing.BuyoutPricePerToken.String(),
		CurrencyContractAddress:     listing.Currency.String(),
		BuyoutCurrencyValuePerToken: currencyValue,
		Id:                          listing.ListingId.String(),
		TokenId:                     int(listing.TokenId.Int64()),
		Quantity:                    int(listing.Quantity.Int64()),
		StartTimeInEpochSeconds:     int(listing.StartTime.Int64()),
		EndTimeInEpochSeconds:       int(listing.EndTime.Int64()),
		SellerAddress:               listing.TokenOwner.String(),
		Asset:                       asset,
	}, nil
}
