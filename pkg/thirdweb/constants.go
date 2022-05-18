package thirdweb

import "errors"

const DEFAULT_MERKLE_ROOT = "0x0000000000000000000000000000000000000000000000000000000000000000"

type ChainID int

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
