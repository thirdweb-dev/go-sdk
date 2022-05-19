package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTCollection struct {
	contractWrapper *contractWrapper[*abi.TokenERC721]
	*ERC721
}

func newNFTCollection(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTCollection, error) {
	if erc721, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := newContractWrapper(erc721, address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				nftCollection := &NFTCollection{
					contractWrapper,
					erc721,
				}
				return nftCollection, nil
			}
		}
	}
}

// Mint
//
// Mint a new NFT to the connected wallet
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error) {
	address := nft.contractWrapper.GetSignerAddress().String()
	return nft.MintTo(address, metadata)
}

// Mint
//
// Mint a new NFT to the specified wallet
//
// address: the wallet address to mint to
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(metadata, nft.storage)
	if err != nil {
		return nil, err
	}

	tx, err := nft.contractWrapper.abi.MintTo(
		nft.contractWrapper.getTxOptions(),
		common.HexToAddress(address),
		uri,
	)
	if err != nil {
		return nil, err
	}

	return nft.contractWrapper.awaitTx(tx.Hash())
}

// MintBatch
//
// Mint a batch of new NFTs to the connected wallet
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	address := nft.contractWrapper.GetSignerAddress().String()
	return nft.MintBatchTo(address, metadatas)
}

// MintBatchTo
//
// Mint a batch of new NFTs to the specified wallet
//
// to: the wallet address to mint to
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	uris, err := uploadOrExtractUris(metadatas, nft.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	for _, uri := range uris {
		tx, err := nft.contractWrapper.abi.MintTo(
			nft.contractWrapper.getTxOptions(),
			common.HexToAddress(address), uri,
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := nft.contractWrapper.abi.Multicall(nft.contractWrapper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return nft.contractWrapper.awaitTx(tx.Hash())
}
