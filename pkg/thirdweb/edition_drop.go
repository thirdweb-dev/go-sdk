package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type EditionDrop struct {
	abi    *abi.DropERC1155
	helper *contractHelper
	*ERC1155
	claimConditions *editionDropClaimConditions
}

func newEditionDrop(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*EditionDrop, error) {
	if contractAbi, err := abi.NewDropERC1155(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc1155, err := newERC1155(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				claimConditions, err := newEditionDropClaimConditions(address, provider, helper, storage)
				if err != nil {
					return nil, err
				}

				edition := &EditionDrop{
					contractAbi,
					helper,
					erc1155,
					claimConditions,
				}
				return edition, nil
			}
		}
	}
}

// CreateBatch
//
// Create a batch of NFTs on this contract
//
// metadatas: a list of the metadatas of the NFTs to create
//
// returns: the transaction receipt of the batch creation
func (drop *EditionDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
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
// tokenId: the token ID of the NFT to claim
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *EditionDrop) Claim(tokenId int, quantity int) (*types.Transaction, error) {
	address := drop.helper.GetSignerAddress().String()
	return drop.ClaimTo(address, tokenId, quantity)
}

// ClaimTo
//
// Claim NFTs from this contract to the connect wallet
//
// tokenId: the token ID of the NFT to claim
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *EditionDrop) ClaimTo(destinationAddress string, tokenId int, quantity int) (*types.Transaction, error) {
	claimVerification, err := drop.prepareClaim(tokenId, quantity)
	if err != nil {
		return nil, err
	}

	tx, err := drop.abi.Claim(
		drop.helper.getTxOptions(),
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(tokenId)),
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

func (drop *EditionDrop) prepareClaim(tokenId int, quantity int) (*ClaimVerification, error) {
	claimCondition, err := drop.claimConditions.GetActive(tokenId)
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
