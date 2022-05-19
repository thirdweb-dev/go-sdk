package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type Edition struct {
	*ERC1155
}

func newEdition(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Edition, error) {
	if erc1155, err := abi.NewTokenERC1155(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := newContractWrapper(erc1155, address, provider, privateKey); err != nil {
			return nil, err
		} else {
			erc1155 := newERC1155(contractWrapper, storage)
			edition := &Edition{
				erc1155,
			}
			return edition, nil
		}
	}
}

// Mint
//
// Mint an NFT to the connected wallet
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) Mint(metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	address := edition.contractWrapper.GetSignerAddress().String()
	return edition.MintTo(address, metadataWithSupply)
}

// MintTo
//
// Mint a new NFT to the specified wallet
//
// address: the wallet address to mint the NFT to
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintTo(address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(metadataWithSupply.Metadata, edition.storage)
	if err != nil {
		return nil, err
	}

	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	tx, err := edition.contractWrapper.abi.MintTo(
		edition.contractWrapper.getTxOptions(),
		common.HexToAddress(address),
		MaxUint256,
		uri,
		big.NewInt(int64(metadataWithSupply.Supply)),
	)
	if err != nil {
		return nil, err
	}

	return edition.contractWrapper.awaitTx((tx.Hash()))
}

// MintAdditionalSupply
//
// Mint additionaly supply of a token to the connected wallet
//
// tokenId: token ID to mint additional supply of
//
// additionalSupply: additional supply to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintAdditionalSupply(tokenId int, additionalSupply int) (*types.Transaction, error) {
	address := edition.contractWrapper.GetSignerAddress().String()
	return edition.MintAdditionalSupplyTo(address, tokenId, additionalSupply)
}

// MintAdditionalSupplyTo
//
// to: address of the wallet to mint NFTs to
//
// tokenId: token Id to mint additional supply of
//
// additionalySupply: additional supply to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintAdditionalSupplyTo(to string, tokenId int, additionalSupply int) (*types.Transaction, error) {
	metadata, err := edition.getTokenMetadata(tokenId)
	if err != nil {
		return nil, err
	}

	tx, err := edition.contractWrapper.abi.MintTo(
		edition.contractWrapper.getTxOptions(),
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
		metadata.Uri,
		big.NewInt(int64(additionalSupply)),
	)
	if err != nil {
		return nil, err
	}

	return edition.contractWrapper.awaitTx(tx.Hash())
}

// MintBatchTo
//
// to: address of the wallet to mint NFTs to
//
// metadatasWithSupply: list of NFT metadatas with supplies to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintBatchTo(to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	metadatas := []*NFTMetadataInput{}
	for _, metadataWithSupply := range metadatasWithSupply {
		metadatas = append(metadatas, metadataWithSupply.Metadata)
	}

	supplies := []int{}
	for _, metadataWithSupply := range metadatasWithSupply {
		supplies = append(supplies, metadataWithSupply.Supply)
	}

	uris, err := uploadOrExtractUris(metadatas, edition.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	for index, uri := range uris {
		tx, err := edition.contractWrapper.abi.MintTo(
			edition.contractWrapper.getTxOptions(),
			common.HexToAddress(to),
			MaxUint256,
			uri,
			big.NewInt(int64(supplies[index])),
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := edition.contractWrapper.abi.Multicall(edition.contractWrapper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return edition.contractWrapper.awaitTx(tx.Hash())
}
