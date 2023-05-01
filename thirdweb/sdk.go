package thirdweb

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThirdwebSDK struct {
	*ProviderHandler
	Storage  IpfsStorage
	Deployer ContractDeployer
	Auth     WalletAuthenticator
}

// NewThirdwebSDK
//
// # Create a new instance of the Thirdweb SDK
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

	return NewThirdwebSDKFromProvider(provider, options)
}

func NewThirdwebSDKFromProvider(provider *ethclient.Client, options *SDKOptions) (*ThirdwebSDK, error) {
	// Define defaults for all the options
	privateKey := ""
	gatewayUrl := defaultIpfsGatewayUrl
	httpClient := http.DefaultClient

	// Override defaults with the options that are defined
	if options != nil {
		if options.PrivateKey != "" {
			privateKey = options.PrivateKey
		}

		if options.GatewayUrl != "" {
			gatewayUrl = options.GatewayUrl
		}

		if options.HttpClient != nil {
			httpClient = options.HttpClient
		}
	}

	storage := newIpfsStorage(gatewayUrl, httpClient)

	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	deployer, err := newContractDeployer(provider, privateKey, storage)
	if err != nil {
		return nil, err
	}

	if deployer == nil {
		fmt.Println("Warning: Contract deployments are not supported on this network. SDK instantiated without a Deployer.")
	}

	auth, err := newWalletAuthenticator(provider, privateKey)
	if err != nil {
		return nil, err
	}

	sdk := &ThirdwebSDK{
		ProviderHandler: handler,
		Storage:         *storage,
		Deployer:        *deployer,
		Auth:            *auth,
	}

	return sdk, nil
}

// GetNFTCollection
//
// # Get an NFT Collection contract SDK instance
//
// address: the address of the NFT Collection contract
func (sdk *ThirdwebSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	return newNFTCollection(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEdition
//
// # Get an Edition contract SDK instance
//
// address: the address of the Edition contract
func (sdk *ThirdwebSDK) GetEdition(address string) (*Edition, error) {
	return newEdition(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetToken
//
// # Returns a Token contract SDK instance
//
// address: address of the token contract
//
// Returns a Token contract SDK instance
func (sdk *ThirdwebSDK) GetToken(address string) (*Token, error) {
	return newToken(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetNFTDrop
//
// # Get an NFT Drop contract SDK instance
//
// address: the address of the NFT Drop contract
func (sdk *ThirdwebSDK) GetNFTDrop(address string) (*NFTDrop, error) {
	return newNFTDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetEditionDrop
//
// # Get an Edition Drop contract SDK instance
//
// address: the address of the Edition Drop contract
func (sdk *ThirdwebSDK) GetEditionDrop(address string) (*EditionDrop, error) {
	return newEditionDrop(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMultiwrap
//
// # Get a Multiwrap contract SDK instance
//
// address: the address of the Multiwrap contract
func (sdk *ThirdwebSDK) GetMultiwrap(address string) (*Multiwrap, error) {
	return newMultiwrap(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetMarketplace
//
// # Get a Marketplace contract SDK instance
//
// address: the address of the Marketplace contract
func (sdk *ThirdwebSDK) GetMarketplace(address string) (*Marketplace, error) {
	return newMarketplace(
		sdk.GetProvider(),
		common.HexToAddress(address),
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

// GetContract
//
// # Get an instance of a custom contract deployed with thirdweb deploy
//
// address: the address of the contract
func (sdk *ThirdwebSDK) GetContract(ctx context.Context, address string) (*SmartContract, error) {
	abi, err := fetchContractMetadataFromAddress(ctx, address, sdk.GetProvider(), &sdk.Storage)
	if err != nil {
		return nil, err
	}

	return sdk.GetContractFromAbi(address, abi)
}

// GetContractFromABI
//
// # Get an instance of ant custom contract from its ABI
//
// address: the address of the contract
//
// abi: the ABI of the contract
func (sdk *ThirdwebSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error) {
	return newSmartContract(
		sdk.GetProvider(),
		common.HexToAddress(address),
		abi,
		sdk.GetRawPrivateKey(),
		&sdk.Storage,
	)
}

func defaultRpc(network string) (string, error) {
	defaultApiKey := "718c5c811c7f3224efb283e04faab56a8a5cbde78d92a6d4cb905b41985d3856"
	return fmt.Sprintf("https://%s.rpc.thirdweb.com/%s", network, defaultApiKey), nil
}

func getDefaultRpcUrl(rpcUrlorName string) (string, error) {
	switch rpcUrlorName {
	case "mumbai":
		return defaultRpc("mumbai")
	case "goerli":
		return defaultRpc("goerli")
	case "polygon":
		return defaultRpc("polygon")
	case "mainnet", "ethereum":
		return defaultRpc("ethereum")
	case "fantom":
		return defaultRpc("fantom")
	case "avalanche":
		return defaultRpc("avalanche")
	case "optimism":
		return defaultRpc("optimism")
	case "optimism-goerli":
		return defaultRpc("optimism-goerli")
	case "arbitrum":
		return defaultRpc("arbitrum")
	case "arbitrum-goerli":
		return defaultRpc("arbitrum-goerli")
	default:
		if strings.HasPrefix(rpcUrlorName, "http") || strings.HasPrefix(rpcUrlorName, "wss") {
			return rpcUrlorName, nil
		} else {
			return "", fmt.Errorf("invalid rpc url or chain name: %s", rpcUrlorName)
		}
	}
}
