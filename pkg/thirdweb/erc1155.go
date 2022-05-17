package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type ERC1155 struct {
	contractWrapper *ContractWrapper[*abi.TokenERC1155]
	storage         Storage
}

func NewERC1155(contractWrapper *ContractWrapper[*abi.TokenERC1155], storage Storage) *ERC1155 {
	return &ERC1155{
		contractWrapper,
		storage,
	}
}

func (erc1155 *ERC1155) Get(tokenId int) (*EditionMetadata, error) {
	supply := 0
	if totalSupply, err := erc1155.contractWrapper.abi.TotalSupply(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err == nil {
		supply = int(totalSupply.Int64())
	}

	if metadata, err := erc1155.getTokenMetadata(tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &EditionMetadata{
			Metadata: metadata,
			Supply:   supply,
		}
		return metadataOwner, nil
	}
}

func (erc1155 *ERC1155) GetAll() ([]*EditionMetadata, error) {
	if maxId, err := erc1155.GetTotalCount(); err != nil {
		return nil, err
	} else {
		metadatas := []*EditionMetadata{}
		for i := 0; i < int(maxId.Int64()); i++ {
			if metadata, err := erc1155.Get(i); err == nil {
				metadatas = append(metadatas, metadata)
			}
		}
		return metadatas, nil
	}
}

func (erc1155 *ERC1155) GetTotalCount() (*big.Int, error) {
	return erc1155.contractWrapper.abi.NextTokenIdToMint(&bind.CallOpts{})
}

func (erc1155 *ERC1155) GetOwned(address string) ([]*EditionMetadataOwner, error) {
	if address == "" {
		address = erc1155.contractWrapper.GetSignerAddress().String()
	}

	maxId, err := erc1155.contractWrapper.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	owners := []common.Address{}
	ids := []*big.Int{}
	for i := 0; i < int(maxId.Int64()); i++ {
		owners = append(owners, common.HexToAddress(address))
		ids = append(ids, big.NewInt(int64(i)))
	}

	balances, err := erc1155.contractWrapper.abi.BalanceOfBatch(&bind.CallOpts{}, owners, ids)

	metadatas := []*EditionMetadataOwner{}
	for index, balance := range balances {
		metadata, err := erc1155.Get(int(ids[index].Int64()))
		if err == nil {
			metadataOwner := &EditionMetadataOwner{
				Metadata:      metadata.Metadata,
				Supply:        metadata.Supply,
				Owner:         address,
				QuantityOwned: int(balance.Int64()),
			}
			metadatas = append(metadatas, metadataOwner)
		}
	}

	return metadatas, nil
}

func (erc1155 *ERC1155) GetTotalSupply(tokenId int) (*big.Int, error) {
	return erc1155.contractWrapper.abi.TotalSupply(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
}

func (erc1155 *ERC1155) Balance(tokenId int) (*big.Int, error) {
	address := erc1155.contractWrapper.GetSignerAddress().String()
	return erc1155.BalanceOf(address, tokenId)
}

func (erc1155 *ERC1155) BalanceOf(address string, tokenId int) (*big.Int, error) {
	return erc1155.contractWrapper.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address), big.NewInt(int64(tokenId)))
}

func (erc1155 *ERC1155) IsApproved(address string, operator string) (bool, error) {
	return erc1155.contractWrapper.abi.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(address), common.HexToAddress(operator))
}

func (erc1155 *ERC1155) Transfer(to string, tokenId int, amount int) (*types.Transaction, error) {
	if tx, err := erc1155.contractWrapper.abi.SafeTransferFrom(
		erc1155.contractWrapper.getTxOptions(),
		erc1155.contractWrapper.GetSignerAddress(),
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(amount)),
		[]byte{},
	); err != nil {
		return nil, err
	} else {
		return erc1155.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc1155 *ERC1155) Burn(tokenId int, amount int) (*types.Transaction, error) {
	address := erc1155.contractWrapper.GetSignerAddress()
	if tx, err := erc1155.contractWrapper.abi.Burn(
		erc1155.contractWrapper.getTxOptions(),
		address,
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(amount)),
	); err != nil {
		return nil, err
	} else {
		return erc1155.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc1155 *ERC1155) SetApprovalForAll(operator string, approved bool) (*types.Transaction, error) {
	if tx, err := erc1155.contractWrapper.abi.SetApprovalForAll(
		erc1155.contractWrapper.getTxOptions(),
		common.HexToAddress(operator),
		approved,
	); err != nil {
		return nil, err
	} else {
		return erc1155.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc1155 *ERC1155) getTokenMetadata(tokenId int) (*NFTMetadata, error) {
	if uri, err := erc1155.contractWrapper.abi.Uri(
		&bind.CallOpts{},
		big.NewInt(int64(tokenId)),
	); err != nil {
		return nil, &NotFoundError{
			tokenId,
		}
	} else {
		if nft, err := fetchTokenMetadata(tokenId, uri, erc1155.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}
