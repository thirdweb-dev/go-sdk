package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThirdwebSDK struct {
	*ProviderHandler
	storage Storage
}

func NewThirdwebSDK(rpcUrl string, options *SDKOptions) (*ThirdwebSDK, error) {
	provider, err := ethclient.Dial(rpcUrl)
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
