package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	ethAbi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type Pack interface {
	CommonModule
	Open(packId *big.Int) (PackNft, error)
	Get(tokenId *big.Int) (PackMetadata, error)
	GetAll() ([]PackMetadata, error)
	GetNfts(packId *big.Int) ([]PackNft, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Transfer(to string, tokenId *big.Int, quantity *big.Int) error
	Create(args CreatePackArgs) (PackMetadata, error)
}

type PackModule struct {
	Client        *ethclient.Client
	Address       string
	Options       *SdkOptions
	gateway       Gateway
	privateKey    *ecdsa.PrivateKey
	signerAddress common.Address
	module        *abi.Pack
}

func NewPackSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*PackModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	module, err := abi.NewPack(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	var gw Gateway
	gw = NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &PackModule{
		Client:  client,
		Address: address,
		Options: opt,
		gateway: gw,
		module:  module,
	}, nil
}

func (sdk *PackModule) DeployContract(name string) error {
	//chainID, err := sdk.Client.ChainID(context.Background())
	//if err != nil {
	//	return err
	//}
	//
	//parsedAbi, err := ethAbi.JSON(strings.NewReader(abi.ERC1155MetaData.ABI))
	//if err != nil {
	//	return err
	//}
	//
	//auth, err := bind.NewKeyedTransactorWithChainID(sdk.privateKey, chainID)
	//if err != nil {
	//	return err
	//}

	//result, err := bind.DeployContract(auth, parsedAbi, common.FromHex(abi.ERC1155MetaData.Bin), sdk.Client, )
	return nil
}

func (sdk *PackModule) Create(args CreatePackArgs) (PackMetadata, error) {
	if sdk.signerAddress == common.HexToAddress("0") {
		return PackMetadata{}, &NoSignerError{typeName: "pack"}
	}

	log.Printf("Wallet used = %v\n", sdk.signerAddress)

	ids := make([]*big.Int, 0)
	counts := make([]*big.Int, 0)

	for _, addition := range args.Assets {
		ids = append(ids, addition.NftId)
		counts = append(counts, addition.Supply)
	}

	nftSdkModule, err := newErc1155SdkModule(sdk.Client, args.AssetContractAddress, sdk.Options)
	if err != nil {
		return PackMetadata{}, err
	}

	stringsTy, _ := ethAbi.NewType("string", "string", nil)
	uint256Ty, _ := ethAbi.NewType("uint256", "uint256", nil)

	arguments := ethAbi.Arguments{
		{
			Type: stringsTy,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
		{
			Type: uint256Ty,
		},
	}

	uri, err := sdk.gateway.Upload(args.Metadata, sdk.Address, sdk.signerAddress.String())
	if err != nil {
		return PackMetadata{}, err
	}

	bytes, err := arguments.Pack(
		uri,
		args.SecondsUntilOpenStart,
		args.SecondsUntilOpenEnd,
		args.RewardsPerOpen,
	)
	if err != nil {
		log.Print("Failed to pack args")
		return PackMetadata{}, err
	}

	// TODO: check if whats added to pack is erc721 or erc1155, will do later when we support erc721
	tx, err := nftSdkModule.module.ERC1155Transactor.SafeBatchTransferFrom(&bind.TransactOpts{
		From:     sdk.signerAddress,
		Signer:   sdk.getSigner(),
		NoSend:   false,
		GasLimit: 100000,
	}, sdk.signerAddress, sdk.signerAddress, ids, counts, bytes)
	if err != nil {
		return PackMetadata{}, err
	}

	if err := waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts); err != nil {
		return PackMetadata{}, err
	}

	receipt, err := sdk.Client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Printf("Failed to lookup transaction receipt with hash %v\n", tx.Hash().String())
		return PackMetadata{}, err
	}

	log.Printf("Got receipt %v for tx %v\n", receipt.TxHash, tx.Hash())

	if newListing, err := sdk.getNewPack(receipt.Logs); err != nil {
		return PackMetadata{}, err
	} else {
		log.Printf("Getting pack with id %v\n", newListing)
		return sdk.Get(newListing)
	}
}

func (sdk *PackModule) Get(packId *big.Int) (PackMetadata, error) {
	packMeta, err := sdk.module.PackCaller.GetPack(&bind.CallOpts{}, packId)
	if err != nil {
		return PackMetadata{}, err
	}

	if packMeta.Uri == "" {
		return PackMetadata{}, &NotFoundError{identifier: packId, typeName: "pack metadata"}
	}

	packUri, err := sdk.module.PackCaller.TokenURI(&bind.CallOpts{}, packId)
	if err != nil {
		return PackMetadata{}, err
	}

	if packUri == "" {
		return PackMetadata{}, &NotFoundError{identifier: packId, typeName: "pack"}
	}

	body, err := sdk.gateway.Get(packUri)
	if err != nil {
		return PackMetadata{}, err
	}

	// TODO: breakdown this object and apply to PackMetadata
	metadata := NftMetadata{
		Id: packId,
	}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return PackMetadata{}, err
	}

	return PackMetadata{
		Creator:       packMeta.Creator,
		CurrentSupply: *packMeta.CurrentSupply,
		OpenStart:     time.Unix(packMeta.OpenStart.Int64(), 0),
		OpenEnd:       time.Unix(packMeta.OpenEnd.Int64(), 0),
		NftMetadata:   metadata,
	}, nil
}

func (sdk *PackModule) Open(packId *big.Int) (PackNft, error) {
	panic("implement me")
}

func (sdk *PackModule) GetAsync(tokenId *big.Int, ch chan<- PackMetadata, wg *sync.WaitGroup) {
	defer wg.Done()

	result, err := sdk.Get(tokenId)
	if err != nil {
		log.Printf("Failed to fetch nft with id %d err=%v\n", tokenId, err)
		ch <- PackMetadata{}
		return
	}
	ch <- result
}

func (sdk *PackModule) GetAll() ([]PackMetadata, error) {
	maxId, err := sdk.module.PackCaller.NextTokenId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	ch := make(chan PackMetadata)
	defer close(ch)

	count := maxId.Int64()
	for i := int64(0); i < count; i++ {
		id := new(big.Int)
		id.SetInt64(i)

		wg.Add(1)
		go sdk.GetAsync(id, ch, &wg)
	}

	packs := make([]PackMetadata, count)
	for i := range packs {
		packs[i] = <-ch
	}

	wg.Wait()
	return packs, nil
}

func (sdk *PackModule) GetNfts(packId *big.Int) ([]PackNft, error) {
	result, err := sdk.module.PackCaller.GetPackWithRewards(&bind.CallOpts{}, packId)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	ch := make(chan PackNft)
	defer close(ch)

	// TODO: I hate instantiating the module here, could move to New function because it shares the same address as the pack contract
	nftModule, err := NewNftSdkModule(sdk.Client, sdk.Address, sdk.Options)
	if err != nil {
		return nil, err
	}

	for _, i := range result.TokenIds {
		wg.Add(1)

		go func(id *big.Int) {
			defer wg.Done()

			metadata, err := nftModule.Get(id)
			if err != nil {
				// TODO (IMPORTANT): what to do in this case?? ts-sdk moves on I think...
				log.Printf("Failed to get metdata for nft %d in pack %d\n", id, packId)

				ch <- PackNft{}
				return
			}

			ch <- PackNft{
				NftMetadata: metadata,
				Supply:      result.Pack.CurrentSupply,
			}
		}(i)
	}

	packNfts := make([]PackNft, len(result.TokenIds))
	for i := range packNfts {
		packNfts[i] = <-ch
	}

	wg.Wait()
	return packNfts, nil
}

func (sdk *PackModule) Balance(tokenId *big.Int) (*big.Int, error) {
	if sdk.signerAddress == common.HexToAddress("0") {
		return nil, &NoSignerError{typeName: "pack"}
	}

	return sdk.module.PackCaller.BalanceOf(&bind.CallOpts{}, sdk.signerAddress, tokenId)
}

func (sdk *PackModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error) {
	return sdk.module.PackCaller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), tokenId)
}

func (sdk *PackModule) Transfer(to string, tokenId *big.Int, quantity *big.Int) error {
	panic("implement me")
}

func (sdk *PackModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}

func (sdk *PackModule) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}

func (sdk *PackModule) getSignerAddress() common.Address {
	if sdk.signerAddress == common.HexToAddress("0") {
		return common.HexToAddress(sdk.Address)
	} else {
		return sdk.signerAddress
	}
}

func (sdk *PackModule) getNewPack(logs []*types.Log) (*big.Int, error) {
	var packId *big.Int
	for _, l := range logs {
		iterator, err := sdk.module.ParsePackCreated(*l)
		if err != nil {
			continue
		}

		if iterator.PackId != nil {
			fmt.Printf("PackMetadata id = %v, listing = %v\n", iterator.PackId, iterator.PackState)
			packId = iterator.PackId
			break
		}
	}

	if packId == nil {
		return nil, errors.New("Could not find Minted pack event for transaction")
	}

	return packId, nil
}
