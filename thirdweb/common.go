package thirdweb

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fxamacker/cbor"
	"github.com/mitchellh/mapstructure"
	"github.com/shopspring/decimal"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// NFT

func fetchTokenMetadata(ctx context.Context, tokenId int, uri string, storage storage) (*NFTMetadata, error) {
	if body, err := storage.Get(ctx, uri); err != nil {
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

func uploadOrExtractUri(ctx context.Context, metadata *NFTMetadataInput, storage storage) (string, error) {
	metadataToUpload := map[string]interface{}{}
	if err := mapstructure.Decode(metadata, &metadataToUpload); err != nil {
		return "", err
	}
	return storage.Upload(ctx, metadataToUpload, "", "")
}

func uploadOrExtractUris(ctx context.Context, metadatas []*NFTMetadataInput, storage storage) ([]string, error) {
	// Why is this necessary?
	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}

	dataToUpload := []map[string]interface{}{}
	if err := mapstructure.Decode(data, &dataToUpload); err != nil {
		return nil, err
	}

	baseUriWithUris, err := storage.UploadBatch(ctx, dataToUpload, 0, "", "")
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

func convertToReadableQuantity(bn *big.Int, decimals int) string {
	if bn.Cmp(new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)) == 0 {
		return "unlimited"
	} else {
		return fmt.Sprintf("%.2f", formatUnits(bn, decimals))
	}
}

func convertQuantityToBigNumber(quantity string, decimals int) (*big.Int, error) {
	if quantity == "unlimited" {
		MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
		return MaxUint256, nil
	} else {
		flt, err := strconv.ParseFloat(quantity, 64)
		if err != nil {
			return nil, err
		}

		return parseUnits(flt, decimals)
	}
}

// TODO: Update value to be string
func parseUnits(value float64, decimals int) (*big.Int, error) {
	floatValue := value * math.Pow(10, float64(decimals))
	bigNumber, ok := new(big.Int).SetString(fmt.Sprintf("%.0f", floatValue), 10)
	if !ok {
		return nil, fmt.Errorf("Failed to parse %f to big.Int", floatValue)
	}

	return bigNumber, nil
}

func normalizePriceValue(ctx context.Context, provider *ethclient.Client, price float64, currencyAddress string) (*big.Int, error) {
	metadata, err := fetchCurrencyMetadata(ctx, provider, currencyAddress)
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

func fetchCurrencyMetadata(ctx context.Context, provider *ethclient.Client, asset string) (*Currency, error) {
	if isNativeToken(asset) {
		chainId, err := provider.ChainID(ctx)
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

		name, err := contractAbi.Name(&bind.CallOpts{Context: ctx})
		symbol, err := contractAbi.Symbol(&bind.CallOpts{Context: ctx})
		decimals, err := contractAbi.Decimals(&bind.CallOpts{Context: ctx})

		currency := &Currency{
			name,
			symbol,
			int(decimals),
		}

		return currency, nil
	}
}

func fetchCurrencyValue(ctx context.Context, provider *ethclient.Client, asset string, price *big.Int) (*CurrencyValue, error) {
	metadata, err := fetchCurrencyMetadata(ctx, provider, asset)
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
	ctx context.Context,
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
		allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
		if err != nil {
			return nil, err
		}

		if allowance.Cmp(value) < 0 {
			// We can get options from the contract instead of ERC20 because they will be the same
			approvalOpts, err := contractToApprove.getUnsignedTxOptions(ctx, signerAddress)
			if err != nil {
				return nil, err
			}

			return erc20.Approve(approvalOpts, spender, value)
		}
	}

	return nil, nil
}

func setErc20Allowance(
	ctx context.Context,
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
		allowance, err := erc20.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
		if err != nil {
			return err
		}

		if allowance.Cmp(value) < 0 {
			// We can get options from the contract instead of ERC20 because they will be the same
			approvalOpts, err := contractToApprove.GetTxOptions(txOpts.Context)
			if err != nil {
				return err
			}

			tx, err := erc20.Approve(approvalOpts, spender, value)
			if err != nil {
				return err
			}

			_, err = contractToApprove.AwaitTx(ctx, tx.Hash())
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func approveErc20Allowance(
	ctx context.Context,
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

	erc20, err := newContractHelper(common.HexToAddress(currencyAddress), provider, contractToApprove.GetRawPrivateKey())
	if err != nil {
		return err
	}

	owner := contractToApprove.GetSignerAddress()
	spender := contractToApprove.getAddress()
	allowance, err := contractAbi.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
	if err != nil {
		return err
	}

	totalPrice := price.Mul(big.NewInt(int64(quantity)), price)

	if allowance.Cmp(totalPrice) < 0 {
		txOpts, err := erc20.GetTxOptions(ctx)
		if err != nil {
			return err
		}
		tx, err := contractAbi.Approve(txOpts, spender, allowance.Add(allowance, totalPrice))
		if err != nil {
			return err
		}

		if _, err := contractToApprove.AwaitTx(ctx, tx.Hash()); err != nil {
			return err
		}
	}

	return nil
}

func hasErc20Allowance(ctx context.Context, contractToApprove *contractHelper, currencyAddress string, value *big.Int) (bool, error) {
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
		allowance, err := contractAbi.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
		if err != nil {
			return false, err
		}

		return allowance.Cmp(value) >= 0, nil
	}
}

// DROP

func prepareClaim(
	ctx context.Context,
	addressToClaim string,
	quantity int,
	activeClaimCondition *ClaimConditionOutput,
	merkleMetadata *map[string]string,
	contractHelper *contractHelper,
	storage storage,
) (*ClaimVerification, error) {
	maxClaimable := activeClaimCondition.MaxClaimablePerWallet
	proofs := [][32]byte{}
	priceInProof := activeClaimCondition.Price
	currencyAddressInProof := activeClaimCondition.CurrencyAddress
	value := big.NewInt(0)

	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)

	// Snapshot and merkle proof stuff
	// TODO: Check if this string conversion works
	if !strings.HasPrefix(hex.EncodeToString(activeClaimCondition.MerkleRootHash[:]), zeroAddress) {
		snapshotEntry, err := fetchSnapshotEntryForAddress(
			ctx,
			common.HexToAddress(addressToClaim),
			activeClaimCondition.MerkleRootHash,
			merkleMetadata,
			contractHelper.GetProvider(),
			storage,
		)
		if err != nil {
			return nil, err
		}

		if snapshotEntry != nil {
			proofs = snapshotEntry.Proof

			maxClaimable, err = convertQuantityToBigNumber(snapshotEntry.MaxClaimable, 0)
			if err != nil {
				return nil, err
			}

			if snapshotEntry.Price == "unlimited" || snapshotEntry.Price == "" {
				priceInProof = MaxUint256
			} else {
				flt, err := strconv.ParseFloat(snapshotEntry.Price, 64)
				if err != nil {
					return nil, err
				}

				priceInProof, err = normalizePriceValue(
					ctx,
					contractHelper.GetProvider(),
					flt,
					snapshotEntry.CurrencyAddress,
				)
			}

			if snapshotEntry.CurrencyAddress == "" {
				currencyAddressInProof = zeroAddress
			} else {
				currencyAddressInProof = snapshotEntry.CurrencyAddress
			}
		}
	}

	var pricePerToken *big.Int
	if priceInProof.Cmp(MaxUint256) == 0 {
		pricePerToken = activeClaimCondition.Price
	} else {
		pricePerToken = priceInProof
	}

	var currencyAddress string
	if currencyAddressInProof != zeroAddress {
		currencyAddress = currencyAddressInProof
	} else {
		currencyAddress = activeClaimCondition.CurrencyAddress
	}

	if pricePerToken.Cmp(big.NewInt(0)) > 0 {
		if isNativeToken(currencyAddress) {
			value = big.NewInt(0).Mul(big.NewInt(int64(quantity)), pricePerToken)
		}
	}

	claimVerification := &ClaimVerification{
		Value:                  value,
		Proofs:                 proofs,
		MaxClaimable:           maxClaimable,
		Price:                  pricePerToken,
		CurrencyAddress:        currencyAddress,
		PriceInProof:           priceInProof,
		CurrencyAddressInProof: currencyAddressInProof,
	}

	return claimVerification, nil
}

func fetchSnapshotEntryForAddress(
	ctx context.Context,
	addressToClaim common.Address,
	merkleRootHash [32]byte,
	merkleMetadata *map[string]string,
	provider *ethclient.Client,
	storage storage,
) (*SnapshotEntryWithProof, error) {
	// TODO: Test with no merkle metadata
	if merkleMetadata == nil {
		return nil, nil
	}

	merkleRoot := "0x" + hex.EncodeToString(merkleRootHash[:])

	snapshotUri, exists := (*merkleMetadata)[merkleRoot]
	if exists {
		body, err := storage.Get(ctx, snapshotUri)
		if err != nil {
			return nil, err
		}

		metadata := &ShardedMerkleTreeInfo{}
		if err := json.Unmarshal(body, &metadata); err != nil {
			return nil, err
		}

		if metadata.MerkleRoot == merkleRoot {
			merkleTree := shardedMerkleTreeFromInfo(metadata, storage)
			return merkleTree.GetProof(ctx, addressToClaim.String(), provider)
		}
	}

	return nil, nil
}

func transformResultToClaimCondition(
	ctx context.Context,
	pm *abi.IClaimConditionClaimCondition,
	provider *ethclient.Client,
) (*ClaimConditionOutput, error) {
	currencyValue, err := fetchCurrencyValue(ctx, provider, pm.Currency.String(), pm.PricePerToken)
	if err != nil {
		return nil, err
	}

	startTime := time.Unix(pm.StartTimestamp.Int64(), 0)

	return &ClaimConditionOutput{
		StartTime:             startTime,
		MaxClaimableSupply:    pm.MaxClaimableSupply,
		MaxClaimablePerWallet: pm.QuantityLimitPerWallet,
		CurrentMintSupply:     pm.SupplyClaimed,
		AvailableSupply:       big.NewInt(0).Sub(pm.MaxClaimableSupply, pm.SupplyClaimed),
		WaitInSeconds:         big.NewInt(0),
		Price:                 pm.PricePerToken,
		CurrencyAddress:       pm.Currency.String(),
		CurrencyMetadata:      currencyValue,
		MerkleRootHash:        pm.MerkleRoot,
	}, nil
}

/**
func processClaimConditionInputs(
	claimConditionInputs []*ClaimConditionInput,
	tokenDecimals int,
	provider *ethclient.Client,
	storage storage,
) (*SnapshotInfo, error) {
	snapshotInfos := []*SnapshotInfos{}
	inputsWithSnapshots := []*ClaimConditionInput{}
	copy(inputsWithSnapshots, claimConditionInputs)

	for _, claimCondition := range inputsWithSnapshots {
		if len(claimCondition.Snapshot) > 0 {
			snapshotInfo, err := createSnapshot(
				claimCondition.Snapshot,
				tokenDecimals,
				storage,
			)
			if err != nil {
				return nil, err
			}

			snapshotInfos = append(snapshotInfos, snapshotInfo)
			claimCondition.MerkleRootHash = snapshotInfo.MerkleRoot
		} else {
			claimCondition.MerkleRootHash = "0x0000000000000000000000000000000000000000000000000000000000000000"
		}
	}

	parsedConditions := []*abi.IClaimConditionClaimCondition{}
	for _, claimCondition := range inputsWithSnapshots {
		condition := convertToContractModel(claimCondition, tokenDecimals, provider)
		parsedConditions = append(parsedConditions, condition)
	}

	return nil, nil
}

func convertToContractModel(
	c *ClaimConditionInput,
	tokenDecimals int,
	provider *ethclient.Client,
) *abi.IClaimConditionClaimCondition {
	currency := nativeTokenAddress
	if !isNativeToken(c.CurrencyAddress) {
		currency = c.CurrencyAddress
	}

	fmt.Println(currency)

	// TODO: Finish

	return nil
}
**/

// MULTIWRAP

func isTokenApprovedForTransfer(
	ctx context.Context,
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

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return false, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return false, err
	}

	if isErc721 {
		ierc721, err := abi.NewIERC721(common.HexToAddress(assetContract), provider)
		if err != nil {
			return false, err
		}

		approved, err := ierc721.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(from), common.HexToAddress(transferrerContractAddress))
		if err != nil {
			return false, err
		}

		if approved {
			return true, nil
		}

		address, err := ierc721.GetApproved(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId)))
		if err != nil {
			return false, err
		}

		return strings.ToLower(address.String()) == strings.ToLower(transferrerContractAddress), nil
	} else if isErc1155 {
		ierc1155, err := abi.NewIERC1155(common.HexToAddress(assetContract), provider)
		if err != nil {
			return false, err
		}

		return ierc1155.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(from), common.HexToAddress(transferrerContractAddress))
	} else {
		fmt.Println("Contract does not implement ERC1155 or ERC721")
		return false, nil
	}
}

// CUSTOM CONTRACTS

func fetchContractMetadataFromAddress(ctx context.Context, address string, provider *ethclient.Client, storage storage) (string, error) {
	metadataUri, err := resolveContractUriFromAddress(ctx, address, provider)
	if err != nil {
		return "", err
	}

	metadata, err := fetchContractMetadata(ctx, metadataUri, storage)
	if err != nil {
		return "", err
	}

	return metadata, nil
}

func extractMinimalProxyImplementationAddress(bytecode []byte) string {
	bytecodeString := "0x" + hex.EncodeToString(bytecode)

//  EIP-1167 clone minimal proxy - https://eips.ethereum.org/EIPS/eip-1167
if strings.HasPrefix(bytecodeString, "0x363d3d373d3d3d363d73") {
	implementationAddress := bytecodeString[22 : 22+40]
	return "0x" + implementationAddress
}

// Minimal Proxy with receive() from 0xSplits - https://github.com/0xSplits/splits-contracts/blob/c7b741926ec9746182d0d1e2c4c2046102e5d337/contracts/libraries/Clones.sol
if strings.HasPrefix(bytecodeString, "0x36603057343d5230") {
	implementationAddress := bytecodeString[122 : 122+40]
	return "0x" + implementationAddress
}

// 0age's minimal proxy - https://medium.com/coinmonks/the-more-minimal-proxy-5756ae08ee48
if strings.HasPrefix(bytecodeString, "0x3d3d3d3d363d3d37363d73") {
	implementationAddress := bytecodeString[24 : 24+40]
	return "0x" + implementationAddress
}

// vyper's minimal proxy (uniswap v1) - https://etherscan.io/address/0x09cabec1ead1c0ba254b09efb3ee13841712be14#code
if strings.HasPrefix(bytecodeString, "0x366000600037611000600036600073") {
	implementationAddress := bytecodeString[32 : 32+40]
	return "0x" + implementationAddress
}

return ""
}

func resolveContractUriFromAddress(ctx context.Context, address string, provider *ethclient.Client) (string, error) {
	bytecode, err := provider.CodeAt(ctx, common.HexToAddress(address), nil)
	if err != nil {
		return "", err
	}

	if hex.EncodeToString(bytecode) == "0x" {
		return "", fmt.Errorf("Contract at '%s' does not exist", address)
	}

	implemntationAddress := extractMinimalProxyImplementationAddress(bytecode)
	if implemntationAddress != "" {
		return resolveContractUriFromAddress(ctx, implemntationAddress, provider)
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

func fetchContractMetadata(ctx context.Context, uri string, storage storage) (string, error) {
	body, err := storage.Get(ctx, uri)
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
	ctx context.Context,
	contractAddress string,
	provider *ethclient.Client,
	tokenId int,
	storage storage,
) (*NFTMetadata, error) {
	erc165, err := abi.NewIERC165(common.HexToAddress(contractAddress), provider)
	if err != nil {
		return nil, err
	}

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return nil, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return nil, err
	}

	uri := ""
	if isErc721 {
		contract, err := abi.NewTokenERC721(common.HexToAddress(contractAddress), provider)
		if err != nil {
			return nil, err
		}

		uri, err = contract.TokenURI(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId)))
	} else if isErc1155 {
		contract, err := abi.NewTokenERC1155(common.HexToAddress(contractAddress), provider)
		if err != nil {
			return nil, err
		}

		uri, err = contract.Uri(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId)))
	}

	return fetchTokenMetadata(ctx, tokenId, uri, storage)
}

func handleTokenApproval(
	ctx context.Context,
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

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
	if err != nil {
		return err
	}

	if isErc721 {
		contract, err := abi.NewTokenERC721(common.HexToAddress(assetContract), provider)
		if err != nil {
			return err
		}

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{
			Context: ctx,
		}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return err
		}

		if !approved {
			tokenApproved, err := contract.GetApproved(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId)))
			if err != nil {
				return err
			}

			txOpts, err := helper.GetTxOptions(ctx)
			if err != nil {
				return err
			}

			if strings.ToLower(tokenApproved.String()) != strings.ToLower(marketplaceAddress) {
				tx, err := contract.SetApprovalForAll(txOpts, common.HexToAddress(marketplaceAddress), true)
				if err != nil {
					return err
				}

				_, err = helper.AwaitTx(ctx, tx.Hash())
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

		approved, err := contract.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(from), common.HexToAddress(marketplaceAddress))
		if err != nil {
			return err
		}

		if !approved {
			txOpts, err := helper.GetTxOptions(ctx)
			if err != nil {
				return err
			}

			tx, err := contract.SetApprovalForAll(txOpts, common.HexToAddress(marketplaceAddress), true)
			if err != nil {
				return err
			}

			_, err = helper.AwaitTx(ctx, tx.Hash())
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
	ctx context.Context,
	helper *contractHelper,
	listing *DirectListing,
	quantity int,
) (bool, error) {
	approved, err := isTokenApprovedForTransfer(
		ctx,
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

	isErc721, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0x80, 0xAC, 0x58, 0xCD})
	if err != nil {
		return false, err
	}

	isErc1155, err := erc165.SupportsInterface(&bind.CallOpts{Context: ctx}, [4]byte{0xD9, 0xB6, 0x7A, 0x26})
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

		ownerOf, err := contract.OwnerOf(&bind.CallOpts{Context: ctx}, big.NewInt(int64(listing.TokenId)))
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
			&bind.CallOpts{Context: ctx},
			common.HexToAddress(listing.SellerAddress),
			big.NewInt(int64(listing.TokenId)),
		)

		return balance.Int64() >= int64(quantity), nil
	} else {
		return false, errors.New("Contract does not implement ERC721 or ERC1155")
	}
}

func mapListing(
	ctx context.Context,
	helper *contractHelper,
	storage storage,
	listing abi.IMarketplaceListing,
) (*DirectListing, error) {
	currencyValue, err := fetchCurrencyValue(
		ctx,
		helper.GetProvider(),
		listing.Currency.String(),
		listing.BuyoutPricePerToken,
	)
	if err != nil {
		return nil, err
	}

	asset, err := fetchTokenMetadataForContract(
		ctx,
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
