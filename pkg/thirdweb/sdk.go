package thirdweb

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThirdwebSDK struct {
	*ProviderHandler
	storage Storage
}

func NewThirdwebSDK(rpcUrlOrChainName string, options *SDKOptions) (*ThirdwebSDK, error) {
	rpc, err := getDefaultRpcUrl(rpcUrlOrChainName)
	if err != nil {
		return nil, err
	}

	provider, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}

	if options == nil {
		handler, err := NewProviderHandler(provider, "")
		if err != nil {
			return nil, err
		}

		storage := NewIpfsStorage(DEFAULT_IPFS_GATEWAY_URL)
		sdk := &ThirdwebSDK{
			handler, storage,
		}
		return sdk, nil
	} else {
		gatewayUrl := options.GatewayUrl
		if gatewayUrl == "" {
			gatewayUrl = DEFAULT_IPFS_GATEWAY_URL
		}

		handler, err := NewProviderHandler(provider, options.PrivateKey)
		if err != nil {
			return nil, err
		}

		storage := NewIpfsStorage(gatewayUrl)

		sdk := &ThirdwebSDK{
			handler, storage,
		}
		return sdk, nil
	}
}

func (sdk *ThirdwebSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	if contract, err := NewNFTCollection(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), sdk.storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

func (sdk *ThirdwebSDK) GetEdition(address string) (*Edition, error) {
	if contract, err := NewEdition(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), sdk.storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

func (sdk *ThirdwebSDK) GetNFTDrop(address string) (*NFTDrop, error) {
	if contract, err := NewNFTDrop(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), sdk.storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

func getDefaultRpcUrl(rpcUrlorName string) (string, error) {
	switch rpcUrlorName {
    case "mumbai":
		return "https://polygon-mumbai.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
    case "rinkeby":
        return "https://eth-rinkeby.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
    case "goerli":
        return "https://eth-goerli.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "polygon":
        return "https://polygon-mainnet.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "mainnet":
        return "https://eth-mainnet.g.alchemy.com/v2/_gg7wSSi0KMBsdKnGVfHDueq6xMB9EkC", nil
	case "fantom":
        return "https://rpc.ftm.tools", nil
	case "avalanche":
        return "https://rpc.ankr.com/avalanche", nil
	default:
		if strings.HasPrefix(rpcUrlorName, "http") {
			return rpcUrlorName, nil
		} else {
			return "", fmt.Errorf("invalid rpc url or chain name: %s", rpcUrlorName)
		}
    }
}
