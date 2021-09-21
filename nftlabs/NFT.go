package nftlabs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type NftSdk interface {
	Get(tokenId big.Int) (NftMetadata, error)
	GetAll() string
	Balance(tokenId string) uint64
	BalanceOf(address string, tokenId string) uint64
	Transfer(tokenId string, to string)
}

type NftSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	caller *abi.NFTCaller
}

func NewNftSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*NftSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
	opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewNFTCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	return &NftSdkModule{
		Client: client,
		Address: address,
		Options: opt,
		caller: caller,
	}, nil
}

func (sdk *NftSdkModule) Get(tokenId *big.Int) (NftMetadata, error) {
	tokenUri, err := sdk.caller.TokenURI(&bind.CallOpts{}, tokenId)
	if err != nil {
		return NftMetadata{}, err
	}

	gatewayUrl := replaceIpfsWithGateway(tokenUri, sdk.Options.IpfsGatewayUrl)
	resp, err := http.Get(gatewayUrl)
	if err != nil {
		return NftMetadata{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	metadata := NftMetadata{}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return NftMetadata{}, err
	}

	return metadata, nil
}

func (sdk *NftSdkModule) GetAll() (string, error) {
	maxId, err := sdk.caller.NextTokenId(&bind.CallOpts{})
	if err != nil {
		return "", err
	}

	fmt.Println("MaxID =", maxId)
	// TODO: complete fetch by calling Get on all ids
	return "", nil
}
