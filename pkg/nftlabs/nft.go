package nftlabs

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/internal/abi"
)

type Role string

const (
	AdminRole  Role = "admin"
	MinterRole      = "minter"
)

type Nft interface {
	Get(tokenId *big.Int) (NftMetadata, error)
	GetAll() ([]NftMetadata, error)
	GetOwned(address string) ([]NftMetadata, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string) (*big.Int, error)
	Transfer(to string, tokenId *big.Int) error
	TotalSupply() (*big.Int, error)
	SetApproval(operator string, approved bool) error
	Mint(metadata MintNftMetadata) (NftMetadata, error)
	MintBatch(meta []interface{}) ([]NftMetadata, error)
	Burn(tokenId *big.Int) error
	TransferFrom(from string, to string, tokenId *big.Int) error
	SetRoyaltyBps(amount *big.Int) error
	GrantRole(role Role, address string) error
	RevokeRole(role Role, address string) error

	MintTo(meta MintNftMetadata) (NftMetadata, error)
}

type NftModule struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
	module  *abi.NFT

	main ISdk
}

func (sdk *NftModule) setSigner(signer func(address common.Address, transaction *types.Transaction) (*types.Transaction, error)) {
}

func (sdk *NftModule) RevokeRole(role Role, address string) error {
	panic("implement me")
}

func (sdk *NftModule) MintBatch(meta []interface{}) ([]NftMetadata, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return nil, &NoSignerError{typeName: "nft"}
	}

	var wg sync.WaitGroup
	ch := make(chan string)

	for _, asset := range meta {
		wg.Add(1)
		go func(meta interface{}) {
			log.Printf("Uploading collection meta %v\n", meta)
			defer wg.Done()

			uri, err := sdk.main.getGateway().Upload(meta, sdk.Address, sdk.main.getSignerAddress().String())
			if err != nil {
				// TODO: need better handling, ts sdk does nothing if this fails
				log.Printf("Failed to upload one of the nft metadata in collection creation")
				ch <- ""
			} else {
				ch <- uri
			}
		}(asset)
	}

	results := make([]string, len(meta))
	for i := range results {
		results[i] = <-ch
	}

	wg.Wait()
	close(ch)

	tx, err := sdk.module.MintNFTBatch(&bind.TransactOpts{
		From:   sdk.main.getSignerAddress(),
		NoSend: false,
		Signer: sdk.main.getSigner(),
	}, sdk.main.getSignerAddress(), results)
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

	_, err = sdk.getNewMintedBatch(receipt.Logs)
	if err != nil {
		return nil, err
	}

	// TODO: process batch and fetch new nfts

	return nil, nil
}

func (sdk *NftModule) Burn(tokenId *big.Int) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{
			typeName: "nft",
		}
	}

	_, err := sdk.module.Burn(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, tokenId)

	return err
}

func (sdk *NftModule) TransferFrom(from string, to string, tokenId *big.Int) error {
	panic("implement me")
}

func (sdk *NftModule) SetRoyaltyBps(amount *big.Int) error {
	panic("implement me")
}

func (sdk *NftModule) GrantRole(role Role, address string) error {
	panic("implement me")
}

func (sdk *NftModule) MintTo(metadata MintNftMetadata) (NftMetadata, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return NftMetadata{}, &NoSignerError{
			typeName: "Nft",
		}
	}
	uri, err := sdk.main.getGateway().Upload(metadata, "", "")
	if err != nil {
		return NftMetadata{}, err
	}
	log.Printf("Got back uri = %v\n", uri)

	tx, err := sdk.module.NFTTransactor.MintNFT(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.main.getSigner(),
		From:   sdk.main.getSignerAddress(),
	}, sdk.main.getSignerAddress(), uri)

	if err := waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts); err != nil {
		// TODO: return clearer error
		return NftMetadata{}, err
	}

	receipt, err := sdk.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("Failed to lookup transaction receipt with hash %v\n", tx.Hash().String())
		return NftMetadata{}, err
	}

	tokenId, err := sdk.getNewMintedNft(receipt.Logs)
	if err != nil {
		return NftMetadata{}, err
	}

	return NftMetadata{
		Id:          tokenId,
		Image:       metadata.Image,
		Description: metadata.Description,
		Name:        metadata.Name,
	}, err
}

func (sdk *NftModule) Mint(metadata MintNftMetadata) (NftMetadata, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return NftMetadata{}, &NoSignerError{typeName: "nft"}
	}
	return sdk.MintTo(metadata)
}

func (sdk *NftModule) SetApproval(operator string, approved bool) error {
	panic("implement me")
}

func (sdk *NftModule) TotalSupply() (*big.Int, error) {
	return sdk.module.TotalSupply(&bind.CallOpts{})
}

func (sdk *NftModule) GetOwned(address string) ([]NftMetadata, error) {
	var addressToCheck common.Address
	if address == "" && sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return nil, &NoSignerError{typeName: "nft"}
	} else if address != "" {
		addressToCheck = common.HexToAddress(address)
	} else {
		addressToCheck = sdk.main.getSignerAddress()
	}

	balance, err := sdk.module.BalanceOf(&bind.CallOpts{}, addressToCheck)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	ch := make(chan NftMetadata)
	errCh := make(chan error)

	defer close(errCh)

	for i := big.NewInt(0); i.Cmp(balance) == -1; i.Add(big.NewInt(1), i) {
		wg.Add(1)
		go func(idStr string) {
			defer wg.Done()

			id := big.NewInt(0)
			id.SetString(idStr, 10)

			result, err := sdk.module.TokenOfOwnerByIndex(&bind.CallOpts{}, addressToCheck, id)
			if err != nil {
				log.Printf("Failed to get token of owner by index for user %v, err = %v, id=%d\n", addressToCheck.String(), err, id)
				ch <- NftMetadata{}
				return
			}

			nft, err := sdk.Get(result)
			if err != nil {
				log.Printf("Failed to get token for user %v, err = %v\n", addressToCheck.String(), err)
				ch <- NftMetadata{}
			}

			ch <- nft
		}(i.String())
	}

	results := make([]NftMetadata, balance.Int64())
	for i := range results {
		results[i] = <-ch
	}

	wg.Wait()
	close(ch)

	return results, nil
}

func newNftModule(client *ethclient.Client, address string, main ISdk) (Nft, error) {
	module, err := abi.NewNFT(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	return &NftModule{
		Client:  client,
		Address: address,
		module:  module,
		main: main,
	}, nil
}

func (sdk *NftModule) Get(tokenId *big.Int) (NftMetadata, error) {
	tokenUri, err := sdk.module.NFTCaller.TokenURI(&bind.CallOpts{}, tokenId)
	if err != nil {
		return NftMetadata{}, err
	}

	body, err := sdk.main.getGateway().Get(tokenUri)
	metadata := NftMetadata{
		Id: tokenId,
	}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return NftMetadata{}, &UnmarshalError{body: string(body), typeName: "nft", underlyingError: err}
	}

	return metadata, nil
}

func (sdk *NftModule) GetAsync(tokenId *big.Int, ch chan<- NftMetadata, errCh chan<- error, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := sdk.Get(tokenId)
	if err != nil {
		log.Printf("Failed to fetch nft with id %d\n err=%v", tokenId, err)
		errCh <- err
		return
	}
	ch <- result
}

func (sdk *NftModule) GetAll() ([]NftMetadata, error) {
	maxId, err := sdk.module.NFTCaller.NextTokenId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	// nfts := make([]NftMetadata, 0)
	ch := make(chan NftMetadata)
	errCh := make(chan error)

	defer close(errCh)

	count := maxId.Int64()
	for i := int64(0); i < count; i++ {
		id := new(big.Int)
		id.SetInt64(i)

		wg.Add(1)
		go sdk.GetAsync(id, ch, errCh, &wg)
	}

	results := make([]NftMetadata, count)
	for i := range results {
		results[i] = <-ch
	}

	wg.Wait()
	close(ch)

	return results, nil
}

func (sdk *NftModule) BalanceOf(address string) (*big.Int, error) {
	return sdk.module.NFTCaller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

func (sdk *NftModule) Balance(tokenId *big.Int) (*big.Int, error) {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return nil, &NoSignerError{typeName: "nft"}
	}

	return sdk.module.NFTCaller.BalanceOf(&bind.CallOpts{}, sdk.main.getSignerAddress())
}

func (sdk *NftModule) Transfer(to string, tokenId *big.Int) error {
	if sdk.main.getSignerAddress() == common.HexToAddress("0") {
		return &NoSignerError{typeName: "nft"}
	}

	// TODO: allow you to pass transact opts
	_, err := sdk.module.NFTTransactor.SafeTransferFrom(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.main.getSignerAddress(),
		Signer: sdk.main.getSigner(),
	}, sdk.main.getSignerAddress(), common.HexToAddress(to), tokenId)

	return err
}

func (sdk *NftModule) getNewMintedNft(logs []*types.Log) (*big.Int, error) {
	var tokenId *big.Int
	for _, l := range logs {
		iterator, err := sdk.module.ParseMinted(*l)
		if err != nil {
			continue
		}

		if iterator.TokenId != nil {
			tokenId = iterator.TokenId
			break
		}
	}

	if tokenId == nil {
		return nil, errors.New("Could not find Minted event for transaction")
	}

	return tokenId, nil
}

func (sdk *NftModule) getNewMintedBatch(logs []*types.Log) (*abi.NFTMintedBatch, error) {
	var batch *abi.NFTMintedBatch
	for _, l := range logs {
		iterator, err := sdk.module.ParseMintedBatch(*l)
		if err != nil {
			continue
		}

		if iterator.TokenIds != nil {
			batch = iterator
			break
		}
	}

	if batch == nil {
		return nil, errors.New("Could not find Minted batch event for transaction")
	}

	return batch, nil
}
