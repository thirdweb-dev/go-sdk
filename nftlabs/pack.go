package nftlabs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
	"log"
	"math/big"
	"time"
)

type PackSdk interface {
	Open(packId big.Int) (NftMetadata, error)
	Get(tokenId big.Int) (NftMetadata, error)
	GetAll() ([]Pack, error)
	GetNfts(packId big.Int) ([]PackNft, error)
	Balance(tokenId big.Int) (big.Int, error)
	BalanceOf(address string, tokenId string) uint64
	Transfer(to string, tokenId string, quantity big.Int) error
}

type PackSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	caller *abi.PackCaller
}

func NewPackSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*PackSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewPackCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	var gw Gateway
	gw = NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &PackSdkModule{
		Client: client,
		Address: address,
		Options: opt,
		gateway: gw,
		caller: caller,
	}, nil
}

func (sdk *PackSdkModule) Get(packId *big.Int) (Pack, error) {
	packMeta, err := sdk.caller.GetPack(&bind.CallOpts{}, packId)
	if err != nil {
		return Pack{}, err
	}

	if packMeta.Uri == "" {
		return Pack{}, errors.New(fmt.Sprintf("Could not find pack metadata with id %d", packId))
	}

	packUri, err := sdk.caller.TokenURI(&bind.CallOpts{}, packId)
	if err != nil {
		return Pack{}, err
	}

	if packUri == "" {
		return Pack{}, errors.New(fmt.Sprintf("Could not find pack with id %d", packId))
	}

	body, err := sdk.gateway.Get(packUri)
	if err != nil {
		return Pack{}, err
	}

	// TODO: breakdown this object and apply to Pack
	log.Printf("body = %v", string(body))
	metadata := NftMetadata{
		Id: packId,
	}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return Pack{}, err
	}

	return Pack{
		Creator: packMeta.Creator,
		CurrentSupply: *packMeta.CurrentSupply,
		OpenStart: time.Unix(packMeta.OpenStart.Int64(), 0),
		OpenEnd: time.Unix(packMeta.OpenEnd.Int64(), 0),
		NftMetadata: metadata,
	}, nil
}