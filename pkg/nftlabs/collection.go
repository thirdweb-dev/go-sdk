package nftlabs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/internal/abi"
	"golang.org/x/sync/errgroup"
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
	Create(metadata Metadata) (CollectionMetadata, error)
	CreateAndMint(metadataWithSupply CreateCollectionArgs) (CollectionMetadata, error)
	CreateAndMintBatch(metadataWithSupply []CreateCollectionArgs) ([]CollectionMetadata, error)
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
	SetRestrictedTransfer(restricted bool) error

	Transfer(to string, tokenId *big.Int, amount *big.Int) error
	TransferBatchFrom(from string, to string, args []NftCollectionBatchArgs, amount *big.Int) error
	TransferFrom(from string, to string, args NftCollectionBatchArgs) error

	getModule() *abi.NFTCollection
}

type NftCollectionModule struct {
	Client  *ethclient.Client
	Address string
	module  *abi.NFTCollection

	main ISdk

	defaultModuleImpl
}

func (sdk *NftCollectionModule) Burn(args NftCollectionBatchArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}

	return sdk.BurnFrom(sdk.main.getSignerAddress().String(), args)
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

	for i := big.NewInt(0); i.Cmp(maxId) == -1; i.Add(i, big.NewInt(1)) {
		wg.Add(1)
		go sdk.GetAsync(i, ch, &wg)
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
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}
	if tx, err := sdk.module.SetApprovalForAll(sdk.main.getTransactOpts(true), common.HexToAddress(operator), approved); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) Transfer(to string, tokenId *big.Int, amount *big.Int) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}
	tx, err := sdk.module.SafeTransferFrom(sdk.main.getTransactOpts(true), sdk.main.getSignerAddress(), common.HexToAddress(to), tokenId, amount, nil)
	if err != nil {
		return err
	}
	return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
}

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

func (sdk *NftCollectionModule) Mint(args MintCollectionArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}

	return sdk.MintTo(sdk.main.getSignerAddress().String(), args)
}

func (sdk *NftCollectionModule) BurnBatch(args []NftCollectionBatchArgs) error {
	return sdk.BurnBatchFrom(sdk.main.getSignerAddress().String(), args)
}

func (sdk *NftCollectionModule) BurnBatchFrom(account string, args []NftCollectionBatchArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}
	ids := make([]*big.Int, len(args))
	amounts := make([]*big.Int, len(args))
	for i, arg := range args {
		ids[i] = arg.TokenId
		amounts[i] = arg.Amount
	}

	if tx, err := sdk.module.BurnBatch(sdk.main.getTransactOpts(true), common.HexToAddress(account), ids, amounts); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) BurnFrom(account string, args NftCollectionBatchArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}
	if tx, err := sdk.module.Burn(sdk.main.getTransactOpts(true), common.HexToAddress(account), args.TokenId, args.Amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) Create(metadata Metadata) (CollectionMetadata, error) {
	if results, err := sdk.CreateBatch([]Metadata{metadata}); err != nil {
		return CollectionMetadata{}, err
	} else {
		return results[0], nil
	}
}

func (sdk *NftCollectionModule) CreateAndMint(metadataWithSupply CreateCollectionArgs) (CollectionMetadata, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return CollectionMetadata{}, &NoSignerError{typeName: "NFTCollection"}
	}

	if result, err := sdk.CreateAndMintBatch([]CreateCollectionArgs{metadataWithSupply}); err != nil {
		return CollectionMetadata{}, err
	} else {
		return result[0], nil
	}
}

func (sdk *NftCollectionModule) CreateAndMintBatch(metadataWithSupply []CreateCollectionArgs) ([]CollectionMetadata, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return nil, &NoSignerError{typeName: "NFTCollection"}
	}

	uris := make([]string, len(metadataWithSupply))
	supplies := make([]*big.Int, len(metadataWithSupply))
	for i, arg := range metadataWithSupply {
		if arg.Metadata.MetadataUri != "" {
			uris[i] = arg.Metadata.MetadataUri
			continue
		}
		if uri, err := sdk.main.getGateway().Upload(arg.Metadata.MetadataObject, sdk.Address, sdk.main.getSignerAddress().String()); err != nil {
			return nil, err
		} else {
			uris[i] = uri
		}

		supplies[i] = arg.Supply
	}

	log.Printf("uris=%v sup=%v\n", uris, supplies)
	tx, err := sdk.module.CreateNativeTokens(sdk.main.getTransactOpts(true), sdk.main.getSignerAddress(), uris, supplies, nil)
	if err != nil {
		return nil, err
	}
	if err := waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts); err != nil {
		return nil, err
	}

	receipt, err := sdk.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		return nil, err
	}

	ids, err := sdk.getMintedNativeTokens(receipt.Logs)
	wg := new(errgroup.Group)
	results := make([]CollectionMetadata, len(ids))
	for i, asset := range ids {
		func(meta interface{}, index int) {
			wg.Go(func() error {
				uri, err := sdk.Get(asset)
				if err != nil {
					return err
				} else {
					results[index] = uri
					return nil
				}
			})
		}(asset, i)
	}

	if err := wg.Wait(); err != nil {
		log.Println("Failed to get the newly minted batch")
		return nil, err
	}

	return results, nil
}

func (sdk *NftCollectionModule) CreateBatch(metadata []Metadata) ([]CollectionMetadata, error) {
	metadataWithSupply := make([]CreateCollectionArgs, len(metadata))
	for i, meta := range metadata {
		metadataWithSupply[i] = CreateCollectionArgs{
			Metadata: meta,
			Supply:   big.NewInt(0),
		}
	}
	return sdk.CreateAndMintBatch(metadataWithSupply)
}

func (sdk *NftCollectionModule) CreateWithErc20(tokenContract string, tokenAmount *big.Int, args CreateCollectionArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}

	uri, err := sdk.main.getGateway().Upload(args.Metadata, sdk.Address, sdk.main.getSignerAddress().String())
	if err != nil {
		return err
	}

	if tx, err := sdk.module.WrapERC20(sdk.main.getTransactOpts(true), common.HexToAddress(tokenContract), tokenAmount, args.Supply, uri); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) CreateWithErc721(tokenContract string, tokenAmount *big.Int, args CreateCollectionArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}

	uri, err := sdk.main.getGateway().Upload(args.Metadata, sdk.Address, sdk.main.getSignerAddress().String())
	if err != nil {
		return err
	}

	if tx, err := sdk.module.WrapERC721(sdk.main.getTransactOpts(true), common.HexToAddress(tokenContract), tokenAmount, uri); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) MintBatch(args []MintCollectionArgs) error {
	return sdk.MintBatchTo(sdk.main.getSignerAddress().String(), args)
}

func (sdk *NftCollectionModule) MintBatchTo(toAddress string, args []MintCollectionArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}

	ids := make([]*big.Int, len(args))
	amounts := make([]*big.Int, len(args))
	for i, arg := range args {
		ids[i] = arg.TokenId
		amounts[i] = arg.Amount
	}

	if tx, err := sdk.module.MintBatch(sdk.main.getTransactOpts(true), common.HexToAddress(toAddress), ids, amounts, nil); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) MintTo(toAddress string, args MintCollectionArgs) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}

	if tx, err := sdk.module.Mint(sdk.main.getTransactOpts(true), common.HexToAddress(toAddress), args.TokenId, args.Amount, nil); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) SetApproval(operator string, approved bool) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}
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

func (sdk *NftCollectionModule) SetRoyaltyBps(amount *big.Int) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "NFTCollection"}
	}
	if tx, err := sdk.module.SetRoyaltyBps(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.main.getSigner(),
		From: sdk.main.getSignerAddress(),
	}, amount); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) TransferBatchFrom(from string, to string, args []NftCollectionBatchArgs, amount *big.Int) error {
	ids := make([]*big.Int, len(args))
	amounts := make([]*big.Int, len(args))
	for i, arg := range args {
		ids[i] = arg.TokenId
		amounts[i] = arg.Amount
	}
	if tx, err := sdk.module.SafeBatchTransferFrom(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.main.getSigner(),
		From: sdk.main.getSignerAddress(),
	}, common.HexToAddress(from), common.HexToAddress(to), ids, amounts, nil); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) TransferFrom(from string, to string, args NftCollectionBatchArgs) error {
	if tx, err := sdk.module.SafeTransferFrom(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.main.getSigner(),
		From: sdk.main.getSignerAddress(),
	}, common.HexToAddress(from), common.HexToAddress(to), args.TokenId, args.Amount, nil); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *NftCollectionModule) getMintedNativeTokens(logs []*types.Log) ([]*big.Int, error) {
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
		return nil, errors.New("Could not find Minted batch event for transaction")
	}

	return tokenIds, nil

}

// SetRestrictedTransfer will disable all transfers if set to true
func (sdk *NftCollectionModule) SetRestrictedTransfer(restricted bool) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "collection"}
	}
	if tx, err := sdk.module.SetRestrictedTransfer(sdk.main.getTransactOpts(true), restricted); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}
