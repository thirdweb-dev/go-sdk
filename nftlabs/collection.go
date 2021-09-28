package nftlabs

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
	"log"
	"math/big"
	"sync"
)

type NftCollectionSdk interface {
	CommonModule
	Get(tokenId *big.Int) (CollectionMetadata, error)
	GetAll() ([]CollectionMetadata, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	IsApproved(address string, operator string) (bool, error)
	SetApproved(operator string, approved bool) (error)
	Transfer(to string, tokenId *big.Int, amount *big.Int) (error)
	Create(args CreateCollectionArgs) (CollectionMetadata, error)
	Mint(args MintCollectionArgs) (error)
}

type NftCollectionSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	module *abi.NFTCollection

	privateKey *ecdsa.PrivateKey
	signerAddress common.Address
}

func NewNftCollectionModule(client *ethclient.Client, address string, opt *SdkOptions) (*NftCollectionSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	module, err := abi.NewNFTCollection(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	var gw Gateway
	gw = NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &NftCollectionSdkModule{
		Client: client,
		Address: address,
		Options: opt,
		gateway: gw,
		module: module,
	}, nil
}

func (sdk *NftCollectionSdkModule) Get(tokenId *big.Int) (CollectionMetadata, error) {
	info, err := sdk.module.NftInfo(&bind.CallOpts{}, tokenId)
	if err != nil {
		return CollectionMetadata{}, err
	}

	body, err := sdk.getMetadata(tokenId)
	if err != nil {
		return CollectionMetadata{}, err
	}
	metadata := NftMetadata{
		Id: tokenId,
	}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return CollectionMetadata{}, &UnmarshalError{body: string(body), typeName: "nft", underlyingError: err}
	}

	log.Printf("Got metadata %v for tokenId %v\n", metadata, tokenId)

	return CollectionMetadata{
		NftMetadata: metadata,
		Creator:     info.Creator.String(),
		Supply:      info.Supply,
	}, nil
}

func (sdk *NftCollectionSdkModule) getMetadata(tokenId *big.Int) ([]byte, error) {
	uri, err := sdk.module.Uri(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}

	metadata, err := sdk.gateway.Get(uri)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func (sdk *NftCollectionSdkModule) GetAsync(tokenId *big.Int, ch chan<-CollectionMetadata, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := sdk.Get(tokenId)
	if err != nil {
		log.Printf("Failed to fetch nft with id %d\n err=%v", tokenId, err)
		ch <- CollectionMetadata{}
		return
	}
	ch <- result
}

func (sdk *NftCollectionSdkModule) GetAll() ([]CollectionMetadata, error) {
	maxId, err := sdk.module.NFTCollectionCaller.NextTokenId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	ch := make(chan CollectionMetadata)
	defer close(ch)

	for i := int64(0); i < maxId.Int64(); i++ {
		id := new(big.Int)
		id.SetInt64(i)

		wg.Add(1)
		go sdk.GetAsync(id, ch, &wg)
	}

	results := make([]CollectionMetadata, maxId.Int64())
	for i := range results {
		results[i] = <-ch
	}

	wg.Wait()
	return results, nil
}

func (sdk *NftCollectionSdkModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error) {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) Balance(tokenId *big.Int) (*big.Int, error) {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) IsApproved(address string, operator string) (bool, error) {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) SetApproved(operator string, approved bool) error {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) Transfer(to string, tokenId *big.Int, amount *big.Int) error {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) Create(args CreateCollectionArgs) (CollectionMetadata, error) {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) Mint(args MintCollectionArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionSdkModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}

func (sdk *NftCollectionSdkModule) getSignerAddress() common.Address {
	if sdk.signerAddress == common.HexToAddress("0") {
		return common.HexToAddress(sdk.Address)
	} else {
		return sdk.signerAddress
	}
}
