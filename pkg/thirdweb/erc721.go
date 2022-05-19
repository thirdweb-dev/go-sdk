package thirdweb

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type ERC721 struct {
	contractWrapper *contractWrapper[*abi.TokenERC721]
	storage         storage
}

type NFTResult struct {
	nft *NFTMetadataOwner
	err error
}

func newERC721(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC721, error) {
	if erc721, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else if contractWrapper, err := newContractWrapper(erc721, address, provider, privateKey); err != nil {
		return nil, err
	} else {
		return &ERC721{
			contractWrapper,
			storage,
		}, nil
	}
}

// Get
//
// Get metadata for a token
//
// tokenId: token ID of the token to get the metadata for
//
// returns: the metadata for the NFT and its owner
func (erc721 *ERC721) Get(tokenId int) (*NFTMetadataOwner, error) {
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

// GetAll
//
// Get the metadata of all the NFTs on this contract
//
// returns: the metadata of all the NFTs on this contract
func (erc721 *ERC721) GetAll() ([]*NFTMetadataOwner, error) {
	if totalCount, err := erc721.GetTotalCount(); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}
		for i := 0; i < int(totalCount.Int64()); i++ {
			tokenIds = append(tokenIds, big.NewInt(int64(i)))
		}
		return fetchNFTsByTokenId(erc721, tokenIds)
	}
}

// GetTotalCount
//
// Get the total number of NFTs on this contract
//
// returns: the total number of NFTs on this contract
func (erc721 *ERC721) GetTotalCount() (*big.Int, error) {
	return erc721.contractWrapper.abi.NextTokenIdToMint(&bind.CallOpts{})
}

// GetOwned
//
// Get the metadatas of all the NFTs owned by a specific address
//
// address: the address of the owner of the NFTs
//
// returns: the metadata of all the NFTs owned by the address
func (erc721 *ERC721) GetOwned(address string) ([]*NFTMetadataOwner, error) {
	if address == "" {
		address = erc721.contractWrapper.GetSignerAddress().String()
	}

	if tokenIds, err := erc721.GetOwnedTokenIDs(address); err != nil {
		return nil, err
	} else {
		return fetchNFTsByTokenId(erc721, tokenIds)
	}
}

// GetOwnedTokenIDs
//
// Get the tokenIds of all the NFTs owned by a specific address
//
// address: the address of the owner of the NFTs
//
// returns: the tokenIds of all the NFTs owned by the address
func (erc721 *ERC721) GetOwnedTokenIDs(address string) ([]*big.Int, error) {
	if address == "" {
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

// OwnerOf
//
// Get the owner of an NFT
//
// tokenId: the token ID of the NFT to get the owner of
//
// returns: the owner of the NFT
func (erc721 *ERC721) OwnerOf(tokenId int) (string, error) {
	if address, err := erc721.contractWrapper.abi.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return "", err
	} else {
		return address.String(), nil
	}
}

// TotalSupply
//
// Get the total number of NFTs on this contract
//
// returns: the supply of NFTs on this contract
func (erc721 *ERC721) TotalSupply() (*big.Int, error) {
	return erc721.contractWrapper.abi.TotalSupply(&bind.CallOpts{})
}

// Balance
//
// Get the NFT balance of the connected wallet
//
// returns: the number of NFTs on this contract owned by the connected wallet
func (erc721 *ERC721) Balance() (*big.Int, error) {
	return erc721.BalanceOf(erc721.contractWrapper.GetSignerAddress().String())
}

// BalanceOf
//
// Get the NFT balance of a specific wallet
//
// address: the address of the wallet to get the NFT balance of
//
// returns: the number of NFTs on this contract owned by the specified wallet
func (erc721 *ERC721) BalanceOf(address string) (*big.Int, error) {
	return erc721.contractWrapper.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address))
}

// IsApproved
//
// Check whether an operator address is approved for all operations of a specifc addresses assets
//
// address: the address whose assets are to be checked
//
// operator: the address of the operator to check
//
// returns: true if the operator is approved for all operations of the assets, otherwise false
func (erc721 *ERC721) IsApproved(address string, operator string) (bool, error) {
	return erc721.contractWrapper.abi.IsApprovedForAll(&bind.CallOpts{}, common.HexToAddress(address), common.HexToAddress(operator))
}

// Transfer
//
// Transfer a specified token from the connected wallet to a specified address
//
// to: wallet address to transfer the tokens to
//
// tokenId: the token ID of the NFT to transfer
//
// returns: the transaction of the NFT transfer
func (erc721 *ERC721) Transfer(to string, tokenId int) (*types.Transaction, error) {
	if tx, err := erc721.contractWrapper.abi.SafeTransferFrom(erc721.contractWrapper.getTxOptions(), erc721.contractWrapper.GetSignerAddress(), common.HexToAddress(to), big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.contractWrapper.awaitTx(tx.Hash())
	}
}

// Burn
//
// Burn a specified NFT from the connected wallet
//
// tokenId: tokenID of the token to burn
//
// returns: the transaction receipt of the burn
func (erc721 *ERC721) Burn(tokenId int) (*types.Transaction, error) {
	if tx, err := erc721.contractWrapper.abi.Burn(&bind.TransactOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.contractWrapper.awaitTx(tx.Hash())
	}
}

// SetApprovalForAll
//
// Set the approval for all operations of a specific address's assets
//
// address: the address whose assets are to be approved
//
// operator: the address of the operator to set the approval for
//
// approved: true if the operator is approved for all operations of the assets, otherwise false
//
// returns: the transaction receipt of the approval
func (erc721 *ERC721) SetApprovalForAll(operator string, approved bool) (*types.Transaction, error) {
	if tx, err := erc721.contractWrapper.abi.SetApprovalForAll(erc721.contractWrapper.getTxOptions(), common.HexToAddress(operator), approved); err != nil {
		return nil, err
	} else {
		return erc721.contractWrapper.awaitTx(tx.Hash())
	}
}

func (erc721 *ERC721) getTokenMetadata(tokenId int) (*NFTMetadata, error) {
	if uri, err := erc721.contractWrapper.abi.TokenURI(&bind.CallOpts{}, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		if nft, err := fetchTokenMetadata(tokenId, uri, erc721.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}

func fetchNFTsByTokenId(erc721 *ERC721, tokenIds []*big.Int) ([]*NFTMetadataOwner, error) {
	total := len(tokenIds)

	ch := make(chan *NFTResult)
	// fetch all nfts in parallel
	for i := 0; i < total; i++ {
		go func(id int) {
			if nft, err := erc721.Get(id); err == nil {
				ch <- &NFTResult{nft, nil}
			} else {
				fmt.Println(err)
				ch <- &NFTResult{nil, err}
			}
		}(i)
	}
	// wait for all goroutines to emit
	results := make([]*NFTResult, total)
	for i := range results {
		results[i] = <-ch
	}
	// filter out errors
	nfts := []*NFTMetadataOwner{}
	for _, res := range results {
		if res.nft != nil {
			nfts = append(nfts, res.nft)
		}
	}
	// Sort by ID
	sort.SliceStable(nfts, func(i, j int) bool {
		return nfts[i].Metadata.Id.Cmp(nfts[j].Metadata.Id) < 0
	})
	return nfts, nil
}
