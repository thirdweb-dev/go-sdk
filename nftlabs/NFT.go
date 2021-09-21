package nftlabs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type NftSdk interface {
	Get(tokenId big.Int) (NftMetadata, error)
	GetAll() ([]NftMetadata, error)
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

	if resp.StatusCode != http.StatusOK {
		return NftMetadata{}, errors.New(fmt.Sprintf("Bad status code, %d", resp.StatusCode))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	metadata := NftMetadata{
		Id: tokenId,
	}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return NftMetadata{}, err
	}

	return metadata, nil
}

func (sdk *NftSdkModule) GetAsync(tokenId *big.Int, ch chan<-NftMetadata, errCh chan<-error, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := sdk.Get(tokenId)
	if err != nil {
		log.Printf("Failed to fetch nft with id %d\n err=%v", tokenId, err)
		errCh <- err
		return
	}
	ch <- result
}

func (sdk *NftSdkModule) GetAll() ([]NftMetadata, error) {
	maxId, err := sdk.caller.NextTokenId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	nfts := make([]NftMetadata, 0)
	ch := make(chan NftMetadata)
	errCh := make(chan error)

	defer close(ch)
	defer close(errCh)

	count := maxId.Int64()
	for i := int64(0); i < count; i++ {
		id := new(big.Int)
		id.SetInt64(i)

		wg.Add(1)
		go sdk.GetAsync(id, ch, errCh, &wg)
	}

	go func() {
		for v := range ch {
			nfts = append(nfts, v)
		}
	}()

	wg.Wait()
	return nfts, nil
}
