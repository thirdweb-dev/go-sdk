package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

// You can access the Marketplace interface from the SDK as follows:
//
// 	import (
// 		"github.com/thirdweb-dev/go-sdk/thirdweb"
// 	)
//
// 	privateKey = "..."
//
// 	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
// 	})
//
//	contract, err := sdk.GetMarketplace("{{contract_address}}")
type Marketplace struct {
	abi    *abi.Marketplace
	helper *contractHelper
}

func newMarketplace(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Marketplace, error) {
	if contractAbi, err := abi.NewMarketplace(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		marketplace := &Marketplace{
			contractAbi,
			helper,
		}
		return marketplace, nil
	}
}
