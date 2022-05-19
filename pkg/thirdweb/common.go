package thirdweb

import (
	"context"
	"encoding/json"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
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
	return storage.Upload(metadata, "", "")
}

func uploadOrExtractUris(metadatas []*NFTMetadataInput, storage storage) ([]string, error) {
	// Why is this necessary?
	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}

	baseUriWithUris, err := storage.UploadBatch(data, "", "")
	if err != nil {
		return nil, err
	}

	return baseUriWithUris.uris, nil
}

// TOKEN

func isNativeToken(tokenAddress string) bool {
	isZero := tokenAddress == "0x0000000000000000000000000000000000000000"
	isNative := strings.ToLower(tokenAddress) == "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
	return isZero || isNative
}

func parseUnits(value float64, decimals int) *big.Int {
	return big.NewInt(int64(value * math.Pow10(decimals)))
}

func formatUnits(value *big.Int, decimals int) float64 {
	// Importantly copy value to a new variable so big.Div doesn't mutate it
	bigNumber := big.NewInt(value.Int64())

	divisor := big.NewInt(int64(math.Pow10(decimals)))

	// Cast big number values to floats for float division
	return float64(bigNumber.Int64()) / float64(divisor.Int64())
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
		abi, err := abi.NewTokenERC20(common.HexToAddress(asset), provider)
		if err != nil {
			return nil, err
		}

		name, err := abi.Name(&bind.CallOpts{})
		symbol, err := abi.Symbol(&bind.CallOpts{})
		decimals, err := abi.Decimals(&bind.CallOpts{})

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

	displayValue := formatUnits(price, metadata.decimals)

	currencyValue := &CurrencyValue{
		metadata.name,
		metadata.symbol,
		metadata.decimals,
		price,
		displayValue,
	}
	return currencyValue, nil
}

func approveErc20Allowance(
	contractToApprove *contractWrapper[*abi.DropERC721],
	currencyAddress string,
	price *big.Int,
	quantity int,
) error {
	provider := contractToApprove.GetProvider()
	abi, err := abi.NewIERC20(common.HexToAddress(currencyAddress), provider)
	if err != nil {
		return err
	}

	erc20, err := newContractWrapper(abi, common.HexToAddress(currencyAddress), provider, "")
	if err != nil {
		return err
	}

	owner := contractToApprove.GetSignerAddress()
	spender := contractToApprove.getAddress()
	allowance, err := erc20.abi.Allowance(&bind.CallOpts{}, owner, spender)
	if err != nil {
		return err
	}

	totalPrice := price.Mul(big.NewInt(int64(quantity)), price)

	if allowance.Cmp(totalPrice) < 0 {
		erc20.abi.Approve(erc20.getTxOptions(), spender, allowance.Add(allowance, totalPrice))
	}

	return nil
}

// DROP

func prepareClaim(
	quantity int,
	activeClaimCondition *ClaimConditionOutput,
	contractWrapper *contractWrapper[*abi.DropERC721],
	storage storage,
) (*ClaimVerification, error) {
	maxClaimable := 0
	price := activeClaimCondition.price
	currencyAddress := activeClaimCondition.currencyAddress
	proofs := [][32]byte{}

	if price > 0 && !isNativeToken(currencyAddress) {
		approveErc20Allowance(
			contractWrapper,
			currencyAddress,
			big.NewInt(int64(activeClaimCondition.price)),
			quantity,
		)
	}

	claimVerification := &ClaimVerification{
		proofs:                    proofs,
		maxQuantityPerTransaction: maxClaimable,
		price:                     price,
		currencyAddress:           currencyAddress,
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

	price := formatUnits(pm.PricePerToken, currencyValue.decimals)

	return &ClaimConditionOutput{
		startTime:                   int(pm.StartTimestamp.Int64()),
		maxQuantity:                 int(pm.MaxClaimableSupply.Int64()),
		availableSupply:             int(pm.MaxClaimableSupply.Int64()) - int(pm.SupplyClaimed.Int64()),
		quantityLimitPerTransaction: int(pm.QuantityLimitPerTransaction.Int64()),
		waitInSeconds:               int(pm.WaitTimeInSecondsBetweenClaims.Int64()),
		price:                       price,
		currencyAddress:             pm.Currency.String(),
		currencyMetadata:            currencyValue,
	}, nil
}
