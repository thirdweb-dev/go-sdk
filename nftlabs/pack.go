package nftlabs

import (
	"context"
	"encoding/json"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"

	ethAbi "github.com/ethereum/go-ethereum/accounts/abi"
)

type PackSdk interface {
	Open(packId *big.Int) (PackNft, error)
	Get(tokenId *big.Int) (Pack, error)
	GetAll() ([]Pack, error)
	GetNfts(packId *big.Int) ([]PackNft, error)
	Balance(tokenId *big.Int) (big.Int, error)
	BalanceOf(address string, tokenId string) uint64
	Transfer(to string, tokenId *big.Int, quantity *big.Int) error
	Create(nftContractAddress string, assets []PackNftAddition) error
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

func (sdk *PackSdkModule) Create(nftContractAddress string, assets []PackNftAddition) error {
	// TODO: allow user to supply this to sdk
	privateKey, err := crypto.HexToECDSA("ommitted")
	if err != nil {
			return err
	}

	publicAddress, err := getPublicAddress(privateKey)
	if err != nil {
		// TODO: return better error
		return err
	}

	log.Printf("Wallet used = %v\n", publicAddress)

	ids := make([]*big.Int, 0)
	counts := make([]*big.Int, 0)

	for _, addition := range assets {
		ids = append(ids, addition.NftId)
		counts = append(counts, addition.Supply)
	}

	log.Printf("ids = %v counts = %v\n", ids, counts)

	nftSdkModule, err := NewNftSdkModule(sdk.Client, nftContractAddress, sdk.Options)
	if err != nil {
		return err
	}

	stringsTy, _ := ethAbi.NewType("string", "string", nil)
	uint256Ty, _ := ethAbi.NewType("uint", "uint", nil)

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

	bytes, _ := arguments.Pack(
		"",
        big.NewInt(0),
        big.NewInt(0),
        big.NewInt(1),
    )

	_, err = nftSdkModule.transactor.SafeTransferFrom(&bind.TransactOpts{
		NoSend: false,
		From: publicAddress,
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			ctx := context.Background()
			chainId, _ := sdk.Client.ChainID(ctx)
			return types.SignTx(transaction, types.NewEIP155Signer(chainId), privateKey)
		},
	}, publicAddress, common.HexToAddress("0xC7dB76a94EB034f575317Ed436DE5934eB8286bB"), ids[0], counts[0], bytes)

	if err != nil {
		return err
	}

	return nil
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
		log.Printf("Failed to fetch nft with id %d err=%v\n", tokenId, err)
		ch <- Pack{}
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

	ch := make(chan Pack)
	defer close(ch)

	count := maxId.Int64()
	log.Printf("Found %d packs\n", count)
	for i := int64(0); i < count; i++ {
		id := new(big.Int)
		id.SetInt64(i)

		wg.Add(1)
		go sdk.GetAsync(id, ch, &wg)
	}

	packs := make([]Pack, count)
	for i := range packs {
		packs[i] = <-ch
	}

	wg.Wait()
	return packs, nil
}

func (sdk *PackSdkModule) GetNfts(packId *big.Int) ([]PackNft, error) {
	result, err := sdk.caller.GetPackWithRewards(&bind.CallOpts{}, packId)
	if err != nil {
		return nil, err
	}

	var wg *sync.WaitGroup
	packNfts := make([]PackNft, 0)

	ch := make(chan PackNft)
	defer close(ch)

	// TODO: I hate instantiating the module here, could move to New function because it shares the same address as the pack contract
	nftModule, err := NewNftSdkModule(sdk.Client, sdk.Address, sdk.Options)
	if err != nil {
		return nil, err
	}

	for _, i := range result.TokenIds {
		wg.Add(1)

		go func (id *big.Int) {
			defer wg.Done()

			metadata, err := nftModule.Get(id)
			if err != nil {
				// TODO (IMPORTANT): what to do in this case?? ts-sdk moves on I think...
				log.Printf("Failed to get metdata for nft %d in pack %d\n", id, packId)
				return
			}

			ch <- PackNft{
				NftMetadata: metadata,
				Supply:      result.Pack.CurrentSupply,
			}
		}(i)
	}

	go func () {
		for packNft := range ch {
			packNfts = append(packNfts, packNft)
		}
	}()

	wg.Wait()
	return packNfts, nil
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