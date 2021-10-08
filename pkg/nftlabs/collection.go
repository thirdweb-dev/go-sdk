package nftlabs

import (
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/internal/abi"
	"log"
	"math/big"
	"sync"
)

type NftCollection interface {
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Burn(args NftCollectionBatchArgs) error
	BurnBatch(args []NftCollectionBatchArgs) error
	BurnBatchFrom(account string, args []NftCollectionBatchArgs) error
	BurnFrom(account string, args NftCollectionBatchArgs) error
	Create(metadata Metadata) ([]CollectionMetadata, error)
	CreateAndMint(metadataWithSupply CreateCollectionArgs) (CollectionMetadata, error)
	CreateAndMintBatch(metadataWithSupply CreateCollectionArgs) ([]CollectionMetadata, error)
	CreateBatch(metadata []Metadata) ([]CollectionMetadata, error)
	CreateWithErc20(tokenContract string, tokenAmount *big.Int, args CreateCollectionArgs) error
	CreateWithErc721(tokenContract string, tokenAmount *big.Int, args CreateCollectionArgs) error
	Get(tokenId *big.Int) (CollectionMetadata, error)
	GetAll() ([]CollectionMetadata, error)

	defaultModule
	IsApproved(address string, operator string) (bool, error)

	Mint(args MintCollectionArgs) error
	MintBatch(args []MintCollectionArgs) error
	MintBatchTo(toAddress string, args []MintCollectionArgs) error
	MintTo(toAddress string, args MintCollectionArgs) error

	SetApproval(operator string, approved bool) error
	SetRoyaltyBps(amount *big.Int) error

	Transfer(to string, tokenId *big.Int, amount *big.Int) error
	TransferBatchFrom(from string, to string, args NftCollectionBatchArgs, amount *big.Int) error
	TransferFrom(from string, to string, args NftCollectionBatchArgs) error

	getModule() *abi.NFTCollection
}

type NftCollectionModule struct {
	Client  *ethclient.Client
	Address string
	module  *abi.NFTCollection

	main ISdk
}

func (sdk *NftCollectionModule) Burn(args NftCollectionBatchArgs) error {
	panic("implement me")
}

func newNftCollectionModule(client *ethclient.Client, address string, main ISdk) (*NftCollectionModule, error) {
	module, err := abi.NewNFTCollection(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	return &NftCollectionModule{
		Client:  client,
		Address: address,
		module:  module,
		main: main,
	}, nil
}

func (sdk *NftCollectionModule) getModule() *abi.NFTCollection {
	return sdk.module
}

func (sdk *NftCollectionModule) Get(tokenId *big.Int) (CollectionMetadata, error) {
	creator, err := sdk.module.Creator(&bind.CallOpts{}, tokenId)
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

	supply, err := sdk.module.TotalSupply(&bind.CallOpts{}, tokenId)
	if err != nil{
		return CollectionMetadata{}, err
	}

	return CollectionMetadata{
		NftMetadata: metadata,
		Creator:     creator.String(),
		Supply:      supply,
	}, nil
}

func (sdk *NftCollectionModule) getMetadata(tokenId *big.Int) ([]byte, error) {
	uri, err := sdk.module.Uri(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}

	metadata, err := sdk.main.getGateway().Get(uri)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func (sdk *NftCollectionModule) GetAsync(tokenId *big.Int, ch chan<- CollectionMetadata, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := sdk.Get(tokenId)
	if err != nil {
		log.Printf("Failed to fetch nft with id %d\n err=%v", tokenId, err)
		ch <- CollectionMetadata{}
		return
	}
	ch <- result
}

func (sdk *NftCollectionModule) GetAll() ([]CollectionMetadata, error) {
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

func (sdk *NftCollectionModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error) {
	return sdk.module.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), tokenId)
}

func (sdk *NftCollectionModule) Balance(tokenId *big.Int) (*big.Int, error) {
	return sdk.module.BalanceOf(&bind.CallOpts{}, sdk.main.getSignerAddress(), tokenId)
}

func (sdk *NftCollectionModule) IsApproved(address string, operator string) (bool, error) {
	return sdk.module.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(address), common.HexToAddress(operator))
}

func (sdk *NftCollectionModule) SetApproved(operator string, approved bool) error {
	if tx, err := sdk.module.SetApprovalForAll(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.main.getSigner(),
		From: sdk.main.getSignerAddress(),
	}, common.HexToAddress(operator), approved); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) Transfer(to string, tokenId *big.Int, amount *big.Int) error {
	tx, err := sdk.module.SafeTransferFrom(&bind.TransactOpts{
		NoSend: false,
		From: sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, sdk.main.getSignerAddress(), common.HexToAddress(to), tokenId, amount, nil)
	if err != nil {
		return err
	}
	return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
}

//func (sdk *NftCollectionModule) Create(args []CreateCollectionArgs) ([]CollectionMetadata, error) {
//	assetMeta, err := sdk.uploadBatchMetadata(args)
//	if err != nil {
//		return nil, err
//	}
//
//	uris := make([]string, len(assetMeta))
//	supplies := make([]*big.Int, len(assetMeta))
//	for i, m := range assetMeta {
//		uris[i] = m.Uri
//		supplies[i] = m.Supply
//	}
//
//	tx, err := sdk.module.CreateNativeTokens(&bind.TransactOpts{
//		NoSend: false,
//		From:   sdk.main.getSignerAddress(),
//		Signer: sdk.main.getSigner(),
//	}, uris, supplies)
//
//	if err := waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts); err != nil {
//		// TODO: return clearer error
//		return nil, err
//	}
//
//	receipt, err := sdk.Client.TransactionReceipt(context.Background(), tx.Hash())
//	if err != nil {
//		log.Printf("Failed to lookup transaction receipt with hash %v\n", tx.Hash().String())
//		return nil, err
//	}
//
//	nftIds, err := sdk.getNewCollection(receipt.Logs)
//	if err != nil {
//		return nil, err
//	}
//
//	var wg sync.WaitGroup
//	ch := make(chan CollectionMetadata)
//	defer close(ch)
//	for _, nftId := range nftIds {
//		wg.Add(1)
//		go sdk.GetAsync(nftId, ch, &wg)
//	}
//
//	results := make([]CollectionMetadata, len(nftIds))
//	for i := range results {
//		results[i] = <-ch
//	}
//
//	wg.Wait()
//	return results, nil
//}

func (sdk *NftCollectionModule) getNewCollection(logs []*types.Log) ([]*big.Int, error) {
	var tokenIds []*big.Int
	for _, l := range logs {
		iterator, err := sdk.module.ParseNativeTokens(*l)
		if err != nil {
			continue
		}

		if iterator.TokenIds != nil {
			tokenIds = iterator.TokenIds
			break
		}
	}

	if tokenIds == nil {
		return nil, errors.New("Could not find Minted Nfts event for transaction")
	}

	return tokenIds, nil
}

// Could/should go into the Storage interface UploadBatch(args ...interface{})
func (sdk *NftCollectionModule) uploadBatchMetadata(args []CreateCollectionArgs) ([]collectionAssetMetadata, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
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

			uri, err := sdk.main.getGateway().Upload(meta.Metadata, sdk.Address, sdk.main.getSignerAddress().String())
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

func (sdk *NftCollectionModule) Mint(args MintCollectionArgs) error {
	if tx, err := sdk.module.Mint(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.main.getSigner(),
		From: sdk.main.getSignerAddress(),
	}, common.HexToAddress(sdk.Address), args.TokenId, args.Amount, nil); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) BurnBatch(args []NftCollectionBatchArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) BurnBatchFrom(account string, args []NftCollectionBatchArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) BurnFrom(account string, args NftCollectionBatchArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) Create(metadata Metadata) ([]CollectionMetadata, error) {
	panic("implement me")
}

func (sdk *NftCollectionModule) CreateAndMint(metadataWithSupply CreateCollectionArgs) (CollectionMetadata, error) {
	panic("implement me")
}

func (sdk *NftCollectionModule) CreateAndMintBatch(metadataWithSupply CreateCollectionArgs) ([]CollectionMetadata, error) {
	panic("implement me")
}

func (sdk *NftCollectionModule) CreateBatch(metadata []Metadata) ([]CollectionMetadata, error) {
	panic("implement me")
}

func (sdk *NftCollectionModule) CreateWithErc20(tokenContract string, tokenAmount *big.Int, args CreateCollectionArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) CreateWithErc721(tokenContract string, tokenAmount *big.Int, args CreateCollectionArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) GrantRole(role Role, address string) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) RevokeRole(role Role, address string) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) MintBatch(args []MintCollectionArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) MintBatchTo(toAddress string, args []MintCollectionArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) MintTo(toAddress string, args MintCollectionArgs) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) SetApproval(operator string, approved bool) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) SetRoyaltyBps(amount *big.Int) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) TransferBatchFrom(from string, to string, args NftCollectionBatchArgs, amount *big.Int) error {
	panic("implement me")
}

func (sdk *NftCollectionModule) TransferFrom(from string, to string, args NftCollectionBatchArgs) error {
	panic("implement me")
}
