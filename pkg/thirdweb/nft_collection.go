package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

// You can access the NFT Collection interface from the SDK as follows:
//
//	contract, err := sdk.GetNFTCollection("{{contract_address}}")
type NFTCollection struct {
	abi    *abi.TokenERC721
	helper *contractHelper
	*ERC721
}

func newNFTCollection(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTCollection, error) {
	if contractAbi, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				nftCollection := &NFTCollection{
					contractAbi,
					helper,
					erc721,
				}
				return nftCollection, nil
			}
		}
	}
}

// Mint a new NFT to the connected wallet.
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	nft.Mint(&thirdweb.NFTMetadataInput{ Name: "NFT" })
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error) {
	address := nft.helper.GetSignerAddress().String()
	return nft.MintTo(address, metadata)
}

// Mint a new NFT to the specified wallet.
//
// address: the wallet address to mint to
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//
// 	nft.MintTo("0x0dcd5d886577d5081b0c52e242ef29e70be3e7bc", &thirdweb.NFTMetadataInput{ Name: "NFT" })
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(metadata, nft.storage)
	if err != nil {
		return nil, err
	}

	tx, err := nft.abi.MintTo(
		nft.helper.getTxOptions(),
		common.HexToAddress(address),
		uri,
	)
	if err != nil {
		return nil, err
	}

	return nft.helper.awaitTx(tx.Hash())
}

// MintBatch
//
// Mint a batch of new NFTs to the connected wallet
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	address := nft.helper.GetSignerAddress().String()
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
		tx, err := nft.abi.MintTo(
			nft.helper.getTxOptions(),
			common.HexToAddress(address), uri,
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := nft.abi.Multicall(nft.helper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return nft.helper.awaitTx(tx.Hash())
}
