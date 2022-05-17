package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type DropERC721 struct {
	contractWrapper *ContractWrapper[*abi.DropERC721]
	storage         Storage
}

func NewDropERC721(contractWrapper *ContractWrapper[*abi.DropERC721], storage Storage) *DropERC721 {
	return &DropERC721{
		contractWrapper,
		storage,
	}
}

func (erc721 *DropERC721) Get(tokenId int) (*NFTMetadataOwner, error) {
	owner := "0x0000000000000000000000000000000000000000"
	if address, err := erc721.OwnerOf(tokenId); err == nil {
		owner = address
	}

	if metadata, err := erc721.getTokenMetadata(tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &NFTMetadataOwner{
			Metadata: metadata,
			Owner:    owner,
		}
		return metadataOwner, nil
	}
}

func (erc721 *DropERC721) GetAll() ([]*NFTMetadataOwner, error) {
	if totalCount, err := erc721.GetTotalCount(); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for i := 0; i < int(totalCount.Int64()); i++ {
			if nft, err := erc721.Get(i); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

func (erc721 *DropERC721) GetTotalCount() (*big.Int, error) {
	return erc721.contractWrapper.abi.NextTokenIdToMint(&bind.CallOpts{})
}

func (erc721 *DropERC721) GetOwned(address string) ([]*NFTMetadataOwner, error) {
	if tokenIds, err := erc721.GetOwnedTokenIDs(address); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for _, tokenId := range tokenIds {
			if nft, err := erc721.Get(int(tokenId.Int64())); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

func (erc721 *DropERC721) GetOwnedTokenIDs(address string) ([]*big.Int, error) {
	if address != "" {
		address = erc721.contractWrapper.GetSignerAddress().String()
	}

	if balance, err := erc721.contractWrapper.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address)); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}

		for i := 0; i < int(balance.Int64()); i++ {
			if tokenId, err := erc721.contractWrapper.abi.TokenOfOwnerByIndex(&bind.CallOpts{}, common.HexToAddress(address), big.NewInt(int64(i))); err == nil {
				tokenIds = append(tokenIds, tokenId)
			}
		}

		return tokenIds, nil
	}
}

func (erc721 *DropERC721) OwnerOf(tokenId int) (string, error) {
	if address, err := erc721.contractWrapper.abi.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return "", err
	} else {
		return address.String(), nil
	}
}

func (erc721 *DropERC721) TotalSupply() (*big.Int, error) {
	return erc721.contractWrapper.abi.TotalSupply(&bind.CallOpts{})
}

func (erc721 *DropERC721) Balance() (*big.Int, error) {
	return erc721.BalanceOf(erc721.contractWrapper.GetSignerAddress().String())
}

func (erc721 *DropERC721) BalanceOf(address string) (*big.Int, error) {
	return erc721.contractWrapper.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

func (erc721 *DropERC721) IsApproved(address string, operator string) (bool, error) {
	return erc721.contractWrapper.abi.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(address), common.HexToAddress(operator))
}

func (erc721 *DropERC721) Transfer(to string, tokenId int) (*types.Transaction, error) {
	if tx, err := erc721.contractWrapper.abi.SafeTransferFrom(erc721.contractWrapper.getTxOptions(), erc721.contractWrapper.GetSignerAddress(), common.HexToAddress(to), big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc721 *DropERC721) Burn(tokenId int) (*types.Transaction, error) {
	if tx, err := erc721.contractWrapper.abi.Burn(&bind.TransactOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc721 *DropERC721) SetApprovalForAll(operator string, approved bool) (*types.Transaction, error) {
	if tx, err := erc721.contractWrapper.abi.SetApprovalForAll(erc721.contractWrapper.getTxOptions(), common.HexToAddress(operator), approved); err != nil {
		return nil, err
	} else {
		return erc721.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc721 *DropERC721) getTokenMetadata(tokenId int) (*NFTMetadata, error) {
	if uri, err := erc721.contractWrapper.abi.TokenURI(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return nil, &NotFoundError{
			tokenId,
		}
	} else {
		if nft, err := fetchTokenMetadata(tokenId, uri, erc721.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}
