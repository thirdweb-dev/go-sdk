package thirdweb

import "errors"

// SERVER URLS

const defaultIpfsGatewayUrl = "https://gateway.ipfscdn.io/ipfs/"
const twIpfsServerUrl = "https://upload.nftlabs.co"
const pinataIpfsUrl = "https://api.pinata.cloud/pinning/pinFileToIPFS"

// CONSTANT VALUES

const zeroAddress = "0x0000000000000000000000000000000000000000"
const nativeTokenAddress = "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee"
const defaultMerkleRoot = "0x0000000000000000000000000000000000000000000000000000000000000000"

// NATIVE TOKEN BY CHAIN

type ChainID int
type ChainName string

const (
	MAINNET   ChainID = 1
	RINKEBY           = 4
	GOERLI            = 5
	POLYGON           = 137
	FANTOM            = 250
	AVALANCHE         = 43114
	MUMBAI            = 80001
)

func getNativeTokenByChainId(chainId ChainID) (*NativeToken, error) {
	switch chainId {
	case MAINNET:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case RINKEBY:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0xc778417E063141139Fce010982780140Aa0cD5Ab",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case GOERLI:
		return &NativeToken{
			"Ether",
			"ETH",
			18,
			&WrappedToken{
				"0x0bb7509324ce409f7bbc4b701f932eaca9736ab7",
				"Wrapped Ether",
				"WETH",
			},
		}, nil
	case POLYGON:
		return &NativeToken{
			"Matic",
			"MATIC",
			18,
			&WrappedToken{
				"0x0d500B1d8E8eF31E21C99d1Db9A6444d3ADf1270",
				"Wrapped Matic",
				"WMATIC",
			},
		}, nil
	case MUMBAI:
		return &NativeToken{
			"Matic",
			"MATIC",
			18,
			&WrappedToken{
				"0x9c3C9283D3e44854697Cd22D3Faa240Cfb032889",
				"Wrapped Matic",
				"WMATIC",
			},
		}, nil
	case AVALANCHE:
		return &NativeToken{
			"Avalanche",
			"AVAX",
			18,
			&WrappedToken{
				"0xB31f66AA3C1e785363F0875A1B74E27b85FD66c7",
				"Wrapped AVAX",
				"WAVAX",
			},
		}, nil
	case FANTOM:
		return &NativeToken{
			"Fantom",
			"FTM",
			18,
			&WrappedToken{
				"0x21be370D5312f44cB42ce377BC9b8a0cEF1A4C83",
				"Wrapped Fantom",
				"WFTM",
			},
		}, nil
	default:
		return nil, errors.New("Unsupported chain id")
	}
}

// CONTRACT ADDRESSES BY CHAIN ID

const twFactoryAddress = "0x7c487845f98938Bb955B1D5AD069d9a30e4131fd"
const twRegistryAddress = "0x5DBC7B840baa9daBcBe9D2492E45D7244B54A2A0"
const ozDefenderForwarderAddress = "0xc82BbE41f2cF04e3a8efA18F7032BDD7f6d98a81"

func getContractAddressByChainId(chainId ChainID, contractName string) (string, error) {
	var addresses map[string]string

	switch chainId {
	case MAINNET:
		addresses = map[string]string{
			"BiconomyForwarder": "0x84a0856b038eaAd1cC7E297cF34A7e72685A8693",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case RINKEBY:
		addresses = map[string]string{
			"BiconomyForwarder": "0xFD4973FeB2031D4409fB57afEE5dF2051b171104",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case GOERLI:
		addresses = map[string]string{
			"BiconomyForwarder": zeroAddress,
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case POLYGON:
		addresses = map[string]string{
			"BiconomyForwarder": "0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case MUMBAI:
		addresses = map[string]string{
			"BiconomyForwarder": "0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case AVALANCHE:
		addresses = map[string]string{
			"BiconomyForwarder": "0x64CD353384109423a966dCd3Aa30D884C9b2E057",
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	case FANTOM:
		addresses = map[string]string{
			"BiconomyForwarder": zeroAddress,
			"TWFactory":         twFactoryAddress,
			"TWRegistry":        twRegistryAddress,
		}
	default:
		return "", errors.New("Unsupported chain id")
	}

	if address, ok := addresses[contractName]; ok {
		return address, nil
	} else {
		return "", errors.New("Unsupported contract name")
	}
}
