package nftlabs

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
	"log"
	"math/big"
	"sync"
	"time"
)

type PackSdk interface {
	Open(packId *big.Int) (PackNft, error)
	Get(tokenId *big.Int) (Pack, error)
	GetAll() ([]Pack, error)
	GetNfts(packId *big.Int) ([]PackNft, error)
	Balance(tokenId *big.Int) (big.Int, error)
	BalanceOf(address string, tokenId string) uint64
	Transfer(to string, tokenId *big.Int, quantity *big.Int) error
}

type PackSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	caller *abi.PackCaller
	transactor *abi.PackTransactor
}

func NewPackSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*PackSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewPackCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	transactor, err := abi.NewPackTransactor(common.HexToAddress(address), client)
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
		transactor: transactor,
		caller: caller,
	}, nil
}

func (sdk *PackSdkModule) Get(packId *big.Int) (Pack, error) {
	packMeta, err := sdk.caller.GetPack(&bind.CallOpts{}, packId)
	if err != nil {
		return Pack{}, err
	}

	if packMeta.Uri == "" {
		return Pack{}, &NotFoundError{identifier: packId, typeName: "pack metadata"}
	}

	packUri, err := sdk.caller.TokenURI(&bind.CallOpts{}, packId)
	if err != nil {
		return Pack{}, err
	}

	if packUri == "" {
		return Pack{}, &NotFoundError{identifier: packId, typeName: "pack"}
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

func (sdk *PackSdkModule) CreatePack(packId *big.Int) (PackNft, error) {
	panic("implement me")
}

func (sdk *PackSdkModule) Open(packId *big.Int) (PackNft, error) {
	panic("implement me")
}

func (sdk *PackSdkModule) GetAsync(tokenId *big.Int, ch chan<-Pack, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := sdk.Get(tokenId)
	if err != nil {
		log.Printf("Failed to fetch nft with id %d\n err=%v", tokenId, err)
		return
	}
	ch <- result
}

func (sdk *PackSdkModule) GetAll() ([]Pack, error) {
	maxId, err := sdk.caller.NextTokenId(&bind.CallOpts{});
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	packs := make([]Pack, 0)
	ch := make(chan Pack)
	defer close(ch)

	count := maxId.Int64()
	for i := int64(0); i < count; i++ {
		id := new(big.Int)
		id.SetInt64(i)

		wg.Add(1)
		go sdk.GetAsync(id, ch, &wg)
	}

	go func() {
		for v := range ch {
			packs = append(packs, v)
		}
	}()

	wg.Wait()
	return packs, nil
}

func (sdk *PackSdkModule) GetNfts(packId *big.Int) ([]PackNft, error) {
	panic("implement me")
}

func (sdk *PackSdkModule) Balance(tokenId *big.Int) (big.Int, error) {
	panic("implement me")
}

func (sdk *PackSdkModule) BalanceOf(address string, tokenId string) uint64 {
	panic("implement me")
}

func (sdk *PackSdkModule) Transfer(to string, tokenId *big.Int, quantity *big.Int) error {
	panic("implement me")
}