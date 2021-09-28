package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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
	Create(args []CreateCollectionArgs) ([]CollectionMetadata, error)
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

func (sdk *NftCollectionSdkModule) Create(args []CreateCollectionArgs) ([]CollectionMetadata, error) {
	assetMeta, err := sdk.uploadBatchMetadata(args)
	if err != nil {
		return nil, err
	}

	uris := make([]string, len(assetMeta))
	supplies := make([]*big.Int, len(assetMeta))
	for i, m := range assetMeta {
		uris[i] = m.Uri
		supplies[i] = m.Supply
	}

	tx, err := sdk.module.NFTCollectionTransactor.CreateNativeNfts(&bind.TransactOpts{
		NoSend: false,
		From: sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, uris, supplies)

	if err := waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts); err != nil {
		// TODO: return clearer error
		return nil, err
	}

	receipt, err := sdk.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("Failed to lookup transaction receipt with hash %v\n", tx.Hash().String())
		return nil, err
	}

	nftIds, err := sdk.getNewCollection(receipt.Logs)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	ch := make(chan CollectionMetadata)
	defer close(ch)
	for _, nftId := range nftIds {
		wg.Add(1)
		go sdk.GetAsync(nftId, ch, &wg)
	}

	results := make([]CollectionMetadata, len(nftIds))
	for i := range results {
		results[i] = <-ch
	}

	wg.Wait()
	return results, nil
}

func (sdk *NftCollectionSdkModule) getNewCollection(logs []*types.Log) ([]*big.Int, error) {
	var tokenIds []*big.Int
	for _, l := range logs {
		iterator, err := sdk.module.ParseNativeNfts(*l)
		if err != nil {
			continue
		}

		if iterator.NftIds != nil {
			tokenIds = iterator.NftIds
			break
		}
	}

	if tokenIds == nil {
		return nil, errors.New("Could not find Minted Nfts event for transaction")
	}

	return tokenIds, nil
}

func (sdk *NftCollectionSdkModule) uploadBatchMetadata(args []CreateCollectionArgs) ([]collectionAssetMetadata, error) {
	if sdk.signerAddress == common.HexToAddress("0") {
		return nil, &NoSignerError{typeName: "collection"}
	}

	var wg sync.WaitGroup
	ch := make(chan collectionAssetMetadata)
	defer close(ch)

	for _, asset := range args {
		wg.Add(1)
		go func(meta CreateCollectionArgs) {
			log.Printf("Uploading collection meta %v\n", meta.Metadata)
			defer wg.Done()

			uri, err := sdk.gateway.Upload(meta.Metadata, sdk.Address, sdk.signerAddress.String())
			if err != nil {
				// TODO: need better handling, ts sdk does nothing if this fails
				log.Printf("Failed to upload one of the nft metadata in collection creation")
				ch <- collectionAssetMetadata{}
			} else {
				ch <- collectionAssetMetadata{
					Uri:    uri,
					Supply: meta.Supply,
				}
			}
		}(asset)
	}

	results := make([]collectionAssetMetadata, len(args))
	for i := range results {
		results[i] = <-ch
	}

	wg.Wait()
	return results, nil
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

func (sdk *NftCollectionSdkModule) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}
