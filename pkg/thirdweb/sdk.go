package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThirdwebSDK struct {
	*ProviderHandler
	storage Storage
}

func NewThirdwebSDK(provider *ethclient.Client, privateKey string, gatewayUrl string) (*ThirdwebSDK, error) {
	if gatewayUrl == "" {
		gatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	if handler, err := NewProviderHandler(provider, privateKey); err != nil {
		return nil, err
	} else {
		storage := NewIpfsStorage(gatewayUrl)
		sdk := &ThirdwebSDK{
			handler,
			storage,
		}
		return sdk, nil
	}
}

func (sdk *ThirdwebSDK) GetNFTCollection(address string) (*NFTCollection, error) {
	if collection, err := NewNFTCollection(sdk.GetProvider(), common.HexToAddress(address), sdk.GetRawPrivateKey(), sdk.storage); err != nil {
		return nil, err
	} else {
		return collection, nil
	}
}
