package thirdweb

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThirdwebSDK struct {
	*ProviderHandler
	Storage IpfsStorage
}

// NewThirdwebSDK
//
// Create a new instance of the Thirdweb SDK
//
// rpcUrlOrName: the name of the chain to connection to (e.g. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche") or the RPC URL to connect to
//
// options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL
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

		storage := newIpfsStorage(defaultIpfsGatewayUrl)
		sdk := &ThirdwebSDK{
			ProviderHandler: handler,
			Storage:         *storage,
		}
		return sdk, nil
	} else {
		gatewayUrl := options.GatewayUrl
		if gatewayUrl == "" {
			gatewayUrl = defaultIpfsGatewayUrl
		}

		handler, err := NewProviderHandler(provider, options.PrivateKey)
		if err != nil {
			return nil, err
		}

		storage := newIpfsStorage(gatewayUrl)

		sdk := &ThirdwebSDK{
			ProviderHandler: handler,
			Storage:         *storage,
		}
		return sdk, nil
	}
}

// GetNFTCollection
//
// Get an NFT Collection contract SDK instance
//
// address: the address of the NFT Collection contract
func (sdk *ThirdwebSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	if contract, err := newNFTCollection(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), &sdk.Storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

// GetEdition
//
// Get an Edition contract SDK instance
//
// address: the address of the Edition contract
func (sdk *ThirdwebSDK) GetEdition(address string) (*Edition, error) {
	if contract, err := newEdition(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), &sdk.Storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

// GetToken
//
// Returns a Token contract SDK instance
//
// address: address of the token contract
//
// Returns a Token contract SDK instance
func (sdk *ThirdwebSDK) GetToken(address string) (*Token, error) {
	if contract, err := newToken(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), &sdk.Storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

// GetNFTDrop
//
// Get an NFT Drop contract SDK instance
//
// address: the address of the NFT Drop contract
func (sdk *ThirdwebSDK) GetNFTDrop(address string) (*NFTDrop, error) {
	if contract, err := newNFTDrop(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), &sdk.Storage); err != nil {
		return nil, err
	} else {
		return contract, nil
	}
}

// GetEditionDrop
//
// Get an Edition Drop contract SDK instance
//
// address: the address of the Edition Drop contract
func (sdk *ThirdwebSDK) GetEditionDrop(address string) (*EditionDrop, error) {
	if contract, err := newEditionDrop(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), &sdk.Storage); err != nil {
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
	case "mainnet", "ethereum":
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
