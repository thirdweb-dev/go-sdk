package thirdweb

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTCollection struct {
	*ERC721
}

func NewNFTCollection(provider *ethclient.Client, address common.Address, privateKey string, storage Storage) (*NFTCollection, error) {
	if erc721, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := NewContractWrapper(erc721, provider, privateKey); err != nil {
			return nil, err
		} else {
			erc721 := NewERC721(contractWrapper, storage)
			nftCollection := &NFTCollection{
				erc721,
			}
			return nftCollection, nil
		}
	}
}

func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) error {
	address := nft.contractWrapper.GetSignerAddress().String()
	return nft.MintTo(address, metadata)
}

func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) error {
	uri, err := uploadOrExtractUri(metadata, nft.storage)
	if err != nil {
		return err
	}

	tx, err := nft.contractWrapper.abi.MintTo(
		nft.contractWrapper.getTxOptions(),
		common.HexToAddress(address),
		uri,
	)
	if err != nil {
		return err
	}

	return nft.contractWrapper.awaitTx(tx.Hash())
}

func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) error {
	address := nft.contractWrapper.GetSignerAddress().String()
	return nft.MintBatchTo(address, metadatas)
}

func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) error {
	uris, err := uploadOrExtractUris(metadatas, nft.storage)
	if err != nil {
		return err
	}

	encoded := [][]byte{}
	for _, uri := range uris {
		tx, err := nft.contractWrapper.abi.MintTo(
			nft.contractWrapper.getTxOptions(),
			common.HexToAddress(address), uri,
		)
		if err != nil {
			return err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := nft.contractWrapper.abi.Multicall(nft.contractWrapper.getTxOptions(), encoded)
	if err != nil {
		return err
	}

	return nft.contractWrapper.awaitTx(tx.Hash())
}
