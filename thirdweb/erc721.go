package thirdweb

import (
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// This interface is currently support by the NFT Collection and NFT Drop contracts.
// You can access all of its functions through an NFT Collection or NFT Drop contract instance.
type ERC721 struct {
	abi     *abi.TokenERC721
	helper  *contractHelper
	storage storage
}

type NFTResult struct {
	nft *NFTMetadataOwner
	err error
}

func newERC721(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC721, error) {
	if contractAbi, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		return &ERC721{
			contractAbi,
			helper,
			storage,
		}, nil
	}
}

// Get metadata for a token.
//
// tokenId: token ID of the token to get the metadata for
//
// returns: the metadata for the NFT and its owner
//
// Example
//
//		nft, err := contract.Get(context.Background(), 0)
//	 owner := nft.Owner
//		name := nft.Metadata.Name
func (erc721 *ERC721) Get(ctx context.Context, tokenId int) (*NFTMetadataOwner, error) {
	owner := "0x0000000000000000000000000000000000000000"
	if address, err := erc721.OwnerOf(ctx, tokenId); err == nil {
		owner = address
	}

	if metadata, err := erc721.getTokenMetadata(ctx, tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &NFTMetadataOwner{
			Metadata: metadata,
			Owner:    owner,
		}
		return metadataOwner, nil
	}
}

// Get the metadata of all the NFTs on this contract.
//
// returns: the metadata of all the NFTs on this contract
//
// Example
//
//	nfts, err := contract.GetAll(context.Background())
//	ownerOne := nfts[0].Owner
//	nameOne := nfts[0].Metadata.Name
func (erc721 *ERC721) GetAll(ctx context.Context) ([]*NFTMetadataOwner, error) {
	if totalCount, err := erc721.GetTotalCount(ctx); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}
		for i := 0; i < totalCount; i++ {
			tokenIds = append(tokenIds, big.NewInt(int64(i)))
		}
		return erc721.fetchNFTsByTokenId(ctx, tokenIds)
	}
}

// Get the total number of NFTs on this contract.
//
// returns: the total number of NFTs on this contract
func (erc721 *ERC721) GetTotalCount(ctx context.Context) (int, error) {
	count, err := erc721.abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	return int(count.Int64()), nil
}

// Get the owner of an NFT.
//
// tokenId: the token ID of the NFT to get the owner of
//
// returns: the owner of the NFT
func (erc721 *ERC721) OwnerOf(ctx context.Context, tokenId int) (string, error) {
	if address, err := erc721.abi.OwnerOf(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(tokenId))); err != nil {
		return "", err
	} else {
		return address.String(), nil
	}
}

// Get the total number of NFTs on this contract.
//
// returns: the supply of NFTs on this contract
func (erc721 *ERC721) TotalSupply(ctx context.Context) (int, error) {
	supply, err := erc721.abi.TotalSupply(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return 0, err
	}

	return int(supply.Int64()), nil
}

// Get the NFT balance of the connected wallet.
//
// returns: the number of NFTs on this contract owned by the connected wallet
func (erc721 *ERC721) Balance(ctx context.Context) (int, error) {
	return erc721.BalanceOf(ctx, erc721.helper.GetSignerAddress().String())
}

// Get the NFT balance of a specific wallet.
//
// address: the address of the wallet to get the NFT balance of
//
// returns: the number of NFTs on this contract owned by the specified wallet
//
// Example
//
//	address := "{{wallet_address}}"
//	balance, err := contract.BalanceOf(context.Background(), address)
func (erc721 *ERC721) BalanceOf(ctx context.Context, address string) (int, error) {
	balance, err := erc721.abi.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, common.HexToAddress(address))
	if err != nil {
		return 0, err
	}

	return int(balance.Int64()), nil
}

// Check whether an operator address is approved for all operations of a specifc addresses assets.
//
// address: the address whose assets are to be checked
//
// operator: the address of the operator to check
//
// returns: true if the operator is approved for all operations of the assets, otherwise false
func (erc721 *ERC721) IsApproved(ctx context.Context, address string, operator string) (bool, error) {
	return erc721.abi.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(address), common.HexToAddress(operator))
}

// Transfer a specified token from the connected wallet to a specified address.
//
// to: wallet address to transfer the tokens to
//
// tokenId: the token ID of the NFT to transfer
//
// returns: the transaction of the NFT transfer
//
// Example
//
//	to := "0x..."
//	tokenId := 0
//
//	tx, err := contract.Transfer(context.Background(), to, tokenId)
func (erc721 *ERC721) Transfer(ctx context.Context, to string, tokenId int) (*types.Transaction, error) {
	txOpts, err := erc721.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.abi.SafeTransferFrom(txOpts, erc721.helper.GetSignerAddress(), common.HexToAddress(to), big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.helper.awaitTx(tx.Hash())
	}
}

// Burn a specified NFT from the connected wallet.
//
// tokenId: tokenID of the token to burn
//
// returns: the transaction receipt of the burn
//
// Example
//
//	tokenId := 0
//	tx, err := contract.Burn(context.Background(), tokenId)
func (erc721 *ERC721) Burn(ctx context.Context, tokenId int) (*types.Transaction, error) {
	txOpts, err := erc721.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.abi.Burn(txOpts, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.helper.awaitTx(tx.Hash())
	}
}

// Set the approval for all operations of a specific address's assets.
//
// address: the address whose assets are to be approved
//
// operator: the address of the operator to set the approval for
//
// approved: true if the operator is approved for all operations of the assets, otherwise false
//
// returns: the transaction receipt of the approval
func (erc721 *ERC721) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error) {
	txOpts, err := erc721.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.abi.SetApprovalForAll(txOpts, common.HexToAddress(operator), approved); err != nil {
		return nil, err
	} else {
		return erc721.helper.awaitTx(tx.Hash())
	}
}

// Approve an operator for the NFT owner, which allows the operator to call transferFrom or
// safeTransferFrom for the specified token.
//
// operator: the address of the operator to approve
//
// tokenId: the token ID of the NFT to approve
//
// returns: the transaction receipt of the approval
func (erc721 *ERC721) SetApprovalForToken(ctx context.Context, operator string, tokenId int) (*types.Transaction, error) {
	txOpts, err := erc721.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.abi.Approve(txOpts, common.HexToAddress(operator), big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.helper.awaitTx(tx.Hash())
	}
}

func (erc721 *ERC721) getTokenMetadata(ctx context.Context, tokenId int) (*NFTMetadata, error) {
	if uri, err := erc721.abi.TokenURI(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		if nft, err := fetchTokenMetadata(tokenId, uri, erc721.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}

func (erc721 *ERC721) fetchNFTsByTokenId(ctx context.Context, tokenIds []*big.Int) ([]*NFTMetadataOwner, error) {
	total := len(tokenIds)

	ch := make(chan *NFTResult)
	// fetch all nfts in parallel
	for i := 0; i < total; i++ {
		go func(id int) {
			if nft, err := erc721.Get(ctx, id); err == nil {
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
