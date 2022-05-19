package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTDrop struct {
	abi    *abi.DropERC721
	helper *contractHelper
	*ERC721
	claimConditions *nftDropClaimConditions
}

func newNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTDrop, error) {
	if contractAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				claimConditions, err := newNFTDropClaimConditions(address, provider, helper, storage)
				if err != nil {
					return nil, err
				}

				nftCollection := &NFTDrop{
					contractAbi,
					helper,
					erc721,
					claimConditions,
				}
				return nftCollection, nil
			}
		}
	}
}

// GetAllClaimed
//
// Get a list of all the NFTs that have been claimed from this contract
//
// returns: a list of the metadatas of the claimed NFTs
func (drop *NFTDrop) GetAllClaimed() ([]*NFTMetadataOwner, error) {
	if maxId, err := drop.abi.NextTokenIdToClaim(&bind.CallOpts{}); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for i := 0; i < int(maxId.Int64()); i++ {
			if nft, err := drop.Get(i); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

// GetAllUnclaimed
//
// Get a list of all the NFTs on this contract that have not yet been claimed
//
// returns: a list of the metadatas of the unclaimed NFTs
func (drop *NFTDrop) GetAllUnclaimed() ([]*NFTMetadata, error) {
	maxId, err := drop.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	unmintedId, err := drop.abi.NextTokenIdToClaim(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	nfts := []*NFTMetadata{}
	for i := int(unmintedId.Int64()); i < int(maxId.Int64()); i++ {
		if nft, err := drop.getTokenMetadata(int(unmintedId.Int64()) + i); err == nil {
			nfts = append(nfts, nft)
		}
	}

	return nfts, nil
}

// CreateBatch
//
// Create a batch of NFTs on this contract
//
// metadatas: a list of the metadatas of the NFTs to create
//
// returns: the transaction receipt of the batch creation
func (drop *NFTDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	startNumber, err := drop.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	fileStartNumber := int(startNumber.Int64())

	contractAddress := drop.helper.getAddress().String()
	signerAddress := drop.helper.GetSignerAddress().String()

	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}
	batch, err := drop.storage.UploadBatch(
		data,
		fileStartNumber,
		contractAddress,
		signerAddress,
	)

	tx, err := drop.abi.LazyMint(
		drop.helper.getTxOptions(),
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.helper.awaitTx(tx.Hash())
}

// Claim
//
// Claim NFTs from this contract to the connect wallet
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *NFTDrop) Claim(quantity int) (*types.Transaction, error) {
	address := drop.helper.GetSignerAddress().String()
	return drop.ClaimTo(address, quantity)
}

// ClaimTo
//
// Claim NFTs from this contract to the connect wallet
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *NFTDrop) ClaimTo(destinationAddress string, quantity int) (*types.Transaction, error) {
	claimVerification, err := drop.prepareClaim(quantity)
	if err != nil {
		return nil, err
	}

	tx, err := drop.abi.Claim(
		drop.helper.getTxOptions(),
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(claimVerification.currencyAddress),
		big.NewInt(int64(claimVerification.price)),
		claimVerification.proofs,
		big.NewInt(int64(claimVerification.maxQuantityPerTransaction)),
	)
	if err != nil {
		return nil, err
	}

	return drop.helper.awaitTx(tx.Hash())
}

func (drop *NFTDrop) prepareClaim(quantity int) (*ClaimVerification, error) {
	claimCondition, err := drop.claimConditions.GetActive()
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		quantity,
		claimCondition,
		drop.helper,
		drop.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}
