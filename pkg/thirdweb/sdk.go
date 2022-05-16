package thirdweb

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type ThirdwebSDK struct {
	provider *ethclient.Client
}

func NewThirdwebSDK(provider *ethclient.Client, privateKey string, options string, storage string) (*ThirdwebSDK, error) {
	sdk := &ThirdwebSDK{
		provider: provider,
	}

	return sdk, nil
}
