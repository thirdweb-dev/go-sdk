package thirdweb

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProviderHandler struct {
	provider      *ethclient.Client
	privateKey    *ecdsa.PrivateKey
	rawPrivateKey string
	signerAddress common.Address
}

func NewProviderHandler(provider *ethclient.Client, privateKey string) (*ProviderHandler, error) {
	handler := &ProviderHandler{
		provider: provider,
	}

	if privateKey != "" {
		if err := handler.updateAccount(privateKey); err != nil {
			return nil, err
		}
	}

	return handler, nil
}

func (handler *ProviderHandler) UpdateProvider(provider *ethclient.Client) {
	handler.provider = provider
}

func (handler *ProviderHandler) UpdatePrivateKey(privateKey string) error {
	if err := handler.updateAccount(privateKey); err != nil {
		return err
	} else {
		return nil
	}
}

func (handler *ProviderHandler) GetProvider() *ethclient.Client {
	return handler.provider
}

func (handler *ProviderHandler) GetSignerAddress() common.Address {
	return handler.signerAddress
}

func (handler *ProviderHandler) GetRawPrivateKey() string {
	return handler.rawPrivateKey
}

func (handler *ProviderHandler) GetPrivateKey() *ecdsa.PrivateKey {
	return handler.privateKey
}

func (handler *ProviderHandler) GetChainID(ctx context.Context) (*big.Int, error) {
	return handler.provider.ChainID(ctx)
}

func (handler *ProviderHandler) getSigner(ctx context.Context) (bind.SignerFn, error) {
	chainId, err := handler.GetChainID(ctx)
	if err != nil {
		return nil, err
	}
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(transaction, types.LatestSignerForChainID(chainId), handler.privateKey)
	}, nil
}

func (handler *ProviderHandler) updateAccount(privateKey string) error {
	if key, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		handler.privateKey = key
		handler.signerAddress = publicAddress
		handler.rawPrivateKey = privateKey
		return nil
	}
}

func getPublicAddress(key *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := key.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// TODO: return better error
		return [20]byte{}, errors.New("Failed to decode public key from private key, maybe wrong format")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func processPrivateKey(privateKey string) (*ecdsa.PrivateKey, common.Address, error) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		// TODO: return better error
		return nil, [20]byte{}, err
	}

	publicAddress, err := getPublicAddress(key)
	if err != nil {
		// TODO: return better error
		return nil, [20]byte{}, err
	}

	return key, publicAddress, nil
}
