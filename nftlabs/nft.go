package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type NftSdk interface {
	CommonModule
	Get(tokenId *big.Int) (NftMetadata, error)
	GetAll() ([]NftMetadata, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Transfer(to string, tokenId *big.Int, amount *big.Int) error
}

type NftSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	caller *abi.NFTCaller
	transactor *abi.NFTTransactor

	privateKey *ecdsa.PrivateKey
	signerAddress common.Address
}

func NewNftSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*NftSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewNFTCaller(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	transactor, err := abi.NewNFTTransactor(common.HexToAddress(address), client)
	if err != nil {
		// TODO: return better error
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	var gw Gateway
	gw = NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &NftSdkModule{
		Client: client,
		Address: address,
		Options: opt,
		gateway: gw,
		caller: caller,
		transactor: transactor,
	}, nil
}

func (sdk *NftSdkModule) Get(tokenId *big.Int) (NftMetadata, error) {
	tokenUri, err := sdk.caller.TokenURI(&bind.CallOpts{}, tokenId)
	if err != nil {
		return NftMetadata{}, err
	}

	body, err := sdk.gateway.Get(tokenUri)
	metadata := NftMetadata{
		Id: tokenId,
	}
	if err := json.Unmarshal(body, &metadata); err != nil {
		return NftMetadata{}, &UnmarshalError{body: string(body), typeName: "pack", underlyingError: err}
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

func (sdk *NftSdkModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error) {
	return sdk.caller.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), tokenId)
}

func (sdk *NftSdkModule) Balance(tokenId *big.Int) (*big.Int, error) {
	// TODO: use the authenticated wallet to make this call (use wallet address when checking balance)
	panic("implement me")
}

func (sdk *NftSdkModule) Transfer(to string, tokenId *big.Int, amount *big.Int) error {
	if sdk.signerAddress == common.HexToAddress("0") {
		return &NoSignerError{typeName: "nft"}
	}

	// TODO: allow you to pass transact opts
	_, err := sdk.transactor.SafeTransferFrom(&bind.TransactOpts{
		NoSend: false,
		From: sdk.signerAddress,
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			ctx := context.Background()
			chainId, _ := sdk.Client.ChainID(ctx)
			return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
		},
	}, sdk.signerAddress, common.HexToAddress(to), tokenId, amount, nil)

	return err
}

func (sdk *NftSdkModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}
