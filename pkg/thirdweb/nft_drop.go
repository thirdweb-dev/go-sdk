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
	contractWrapper *ContractWrapper[*abi.DropERC721]
	*ERC721
	claimConditions *NFTDropClaimConditions
}

func NewNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage Storage) (*NFTDrop, error) {
	if dropAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if contractWrapper, err := NewContractWrapper(dropAbi, address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := NewERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				claimConditions := NewNFTDropClaimConditions(contractWrapper, storage)
				nftCollection := &NFTDrop{
					contractWrapper,
					erc721,
					claimConditions,
				}
				return nftCollection, nil
			}
		}
	}
}

func (drop *NFTDrop) GetAllClaimed() ([]*NFTMetadataOwner, error) {
	if maxId, err := drop.contractWrapper.abi.NextTokenIdToClaim(&bind.CallOpts{}); err != nil {
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

func (drop *NFTDrop) GetAllUnclaimed() ([]*NFTMetadata, error) {
	maxId, err := drop.contractWrapper.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	unmintedId, err := drop.contractWrapper.abi.NextTokenIdToClaim(&bind.CallOpts{})
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

func (drop *NFTDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	contractAddress := drop.contractWrapper.getAddress().String()
	signerAddress := drop.contractWrapper.GetSignerAddress().String()

	data := []any{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}
	batch, err := drop.storage.UploadBatch(
		data,
		contractAddress,
		signerAddress,
	)

	tx, err := drop.contractWrapper.abi.LazyMint(
		drop.contractWrapper.getTxOptions(),
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.contractWrapper.awaitTx(tx.Hash())
}

func (drop *NFTDrop) Claim(quantity int) (*types.Transaction, error) {
	address := drop.contractWrapper.GetSignerAddress().String()
	return drop.ClaimTo(address, quantity)
}

func (drop *NFTDrop) ClaimTo(destinationAddress string, quantity int) (*types.Transaction, error) {
	claimVerification, err := drop.prepareClaim(quantity)
	if err != nil {
		return nil, err
	}

	tx, err := drop.contractWrapper.abi.Claim(
		drop.contractWrapper.getTxOptions(),
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

	return drop.contractWrapper.awaitTx(tx.Hash())
}

func (drop *NFTDrop) prepareClaim(quantity int) (*ClaimVerification, error) {
	claimCondition, err := drop.claimConditions.GetActive()
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		quantity,
		claimCondition,
		drop.contractWrapper,
		drop.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}
