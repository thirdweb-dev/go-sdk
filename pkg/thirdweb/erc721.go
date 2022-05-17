package thirdweb

import (
	"encoding/json"
	"math/big"

	"../../internal/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ERC721 struct {
	contractWrapper *ContractWrapper[*abi.TokenERC721]
	storage         Storage
}

func NewERC721(contractWrapper *ContractWrapper[*abi.TokenERC721], storage Storage) *ERC721 {
	return &ERC721{
		contractWrapper,
		storage,
	}
}

func (self *ERC721) Get(tokenId int) (*NFTMetadataOwner, error) {
	owner := "0x0000000000000000000000000000000000000000"
	if address, err := self.OwnerOf(tokenId); err == nil {
		owner = address
	}

	if metadata, err := self.getTokenMetadata(tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &NFTMetadataOwner{
			Metadata: metadata,
			Owner:    owner,
		}
		return metadataOwner, nil
	}
}

func (self *ERC721) GetAll() ([]*NFTMetadataOwner, error) {
	if totalCount, err := self.GetTotalCount(); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for i := 0; i < int(totalCount.Int64()); i++ {
			if nft, err := self.Get(i); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

func (self *ERC721) GetTotalCount() (*big.Int, error) {
	return self.contractWrapper.abi.TokenERC721Caller.NextTokenIdToMint(&bind.CallOpts{})
}

func (self *ERC721) GetOwned(address string) ([]*NFTMetadataOwner, error) {
	if tokenIds, err := self.GetOwnedTokenIDs(address); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for _, tokenId := range tokenIds {
			if nft, err := self.Get(int(tokenId.Int64())); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

func (self *ERC721) GetOwnedTokenIDs(address string) ([]*big.Int, error) {
	if address != "" {
		address = self.contractWrapper.GetSignerAddress().String()
	}

	if balance, err := self.contractWrapper.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address)); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}

		for i := 0; i < int(balance.Int64()); i++ {
			if tokenId, err := self.contractWrapper.abi.TokenOfOwnerByIndex(&bind.CallOpts{}, common.HexToAddress(address), big.NewInt(int64(i))); err == nil {
				tokenIds = append(tokenIds, tokenId)
			}
		}

		return tokenIds, nil
	}
}

func (self *ERC721) OwnerOf(tokenId int) (string, error) {
	if address, err := self.contractWrapper.abi.TokenERC721Caller.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return "", err
	} else {
		return address.String(), nil
	}
}

func (self *ERC721) TotalSupply() (*big.Int, error) {
	return self.contractWrapper.abi.TokenERC721Caller.TotalSupply(&bind.CallOpts{})
}

func (self *ERC721) Balance() (*big.Int, error) {
	return self.BalanceOf(self.contractWrapper.GetSignerAddress().String())
}

func (self *ERC721) BalanceOf(address string) (*big.Int, error) {
	return self.contractWrapper.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

func (self *ERC721) IsApproved(address string, operator string) (bool, error) {
	return self.contractWrapper.abi.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(address), common.HexToAddress(operator))
}

func (self *ERC721) getTokenMetadata(tokenId int) (*NFTMetadata, error) {
	if uri, err := self.contractWrapper.abi.TokenURI(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return nil, &NotFoundError{
			tokenId,
		}
	} else {
		if nft, err := fetchTokenMetadata(tokenId, uri, self.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}

func fetchTokenMetadata(tokenId int, uri string, storage Storage) (*NFTMetadata, error) {
	if body, err := storage.Get(uri); err != nil {
		return nil, err
	} else {
		metadata := &NFTMetadata{
			Id: big.NewInt(int64(tokenId)),
		}
		if err := json.Unmarshal(body, &metadata); err != nil {
			return nil, &UnmarshalError{body: string(body), typeName: "nft", UnderlyingError: err}
		}

		return metadata, nil
	}
}
