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

// This interface is currently support by the Edition and Edition Drop contracts.
// You can access all of its functions through an Edition or Edition Drop contract instance.
type ERC1155 struct {
	abi     *abi.TokenERC1155
	helper  *contractHelper
	storage storage
}

type EditionResult struct {
	nft *EditionMetadata
	err error
}

func newERC1155(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC1155, error) {
	if contractAbi, err := abi.NewTokenERC1155(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		return &ERC1155{
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
// returns: the metadata for the NFT and its supply
//
// Example
//
//		nft, err := contract.Get(context.Background(), 0)
//	 supply := nft.Supply
//		name := nft.Metadata.Name
func (erc1155 *ERC1155) Get(ctx context.Context, tokenId int) (*EditionMetadata, error) {
	supply := 0
	if totalSupply, err := erc1155.abi.TotalSupply(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId))); err == nil {
		supply = int(totalSupply.Int64())
	}

	if metadata, err := erc1155.getTokenMetadata(ctx, tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &EditionMetadata{
			Metadata: metadata,
			Supply:   supply,
		}
		return metadataOwner, nil
	}
}

// Get the metadata of all the NFTs on this contract.
//
// returns: the metadatas and supplies of all the NFTs on this contract
//
// Example
//
//	nfts, err := contract.GetAll(context.Background())
//	supplyOne := nfts[0].Supply
//	nameOne := nfts[0].Metadata.Name
func (erc1155 *ERC1155) GetAll(ctx context.Context) ([]*EditionMetadata, error) {
	if totalCount, err := erc1155.GetTotalCount(ctx); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}
		for i := 0; i < totalCount; i++ {
			tokenIds = append(tokenIds, big.NewInt(int64(i)))
		}
		return fetchEditionsByTokenId(ctx, erc1155, tokenIds)
	}
}

// Get the total number of NFTs on this contract.
//
// returns: the total number of NFTs on this contract
func (erc1155 *ERC1155) GetTotalCount(ctx context.Context) (int, error) {
	count, err := erc1155.abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	return int(count.Int64()), nil
}

// Get the metadatas of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the metadatas and supplies of all the NFTs owned by the address
//
// Example
//
//	owner := "{{wallet_address}}"
//	nfts, err := contract.GetOwned(context.Background(), owner)
//	name := nfts[0].Metadata.Name
func (erc1155 *ERC1155) GetOwned(ctx context.Context, address string) ([]*EditionMetadataOwner, error) {
	if address == "" {
		address = erc1155.helper.GetSignerAddress().String()
	}

	maxId, err := erc1155.abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	owners := []common.Address{}
	ids := []*big.Int{}
	for i := 0; i < int(maxId.Int64()); i++ {
		owners = append(owners, common.HexToAddress(address))
		ids = append(ids, big.NewInt(int64(i)))
	}

	balances, err := erc1155.abi.BalanceOfBatch(&bind.CallOpts{Context: ctx}, owners, ids)
	if err != nil {
		return nil, err
	}

	metadataOwners := []*EditionMetadataOwner{}
	metadatas, err := fetchEditionsByTokenId(ctx, erc1155, ids)
	if err != nil {
		return nil, err
	}
	for index, balance := range balances {
		metadata := metadatas[index]
		if err == nil {
			metadataOwner := &EditionMetadataOwner{
				Metadata:      metadata.Metadata,
				Supply:        metadata.Supply,
				Owner:         address,
				QuantityOwned: int(balance.Int64()),
			}
			metadataOwners = append(metadataOwners, metadataOwner)
		}
	}

	return metadataOwners, nil
}

// Get the total number of NFTs of a specific token ID.
//
// tokenId: the token ID to check the total supply of
//
// returns: the supply of NFTs on the specified token ID
func (erc1155 *ERC1155) TotalSupply(ctx context.Context, tokenId int) (int, error) {
	supply, err := erc1155.abi.TotalSupply(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId)))
	if err != nil {
		return 0, err
	}

	return int(supply.Int64()), nil
}

// Get the NFT balance of the connected wallet for a specific token ID.
//
// tokenId: the token ID of a specific token to check the balance of
//
// returns: the number of NFTs of the specified token ID owned by the connected wallet
func (erc1155 *ERC1155) Balance(ctx context.Context, tokenId int) (int, error) {
	address := erc1155.helper.GetSignerAddress().String()
	return erc1155.BalanceOf(ctx, address, tokenId)
}

// Get the NFT balance of a specific wallet.
//
// address: the address of the wallet to get the NFT balance of
//
// returns: the number of NFTs of the specified token ID owned by the specified wallet
//
// Example
//
//	address := "{{wallet_address}}"
//	tokenId := 0
//	balance, err := contract.BalanceOf(context.Background(), address, tokenId)
func (erc1155 *ERC1155) BalanceOf(ctx context.Context, address string, tokenId int) (int, error) {
	balance, err := erc1155.abi.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(address), big.NewInt(int64(tokenId)))
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
func (erc1155 *ERC1155) IsApproved(ctx context.Context, address string, operator string) (bool, error) {
	return erc1155.abi.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(address), common.HexToAddress(operator))
}

// Transfer a specific quantity of a token ID from the connected wallet to a specified address.
//
// to: wallet address to transfer the tokens to
//
// tokenId: the token ID of the NFT to transfer
//
// amount: number of NFTs of the token ID to transfer
//
// returns: the transaction of the NFT transfer
//
// Example
//
//	to := "0x..."
//	tokenId := 0
//	amount := 1
//
//	tx, err := contract.Transfer(context.Background(), to, tokenId, amount)
func (erc1155 *ERC1155) Transfer(ctx context.Context, to string, tokenId int, amount int) (*types.Transaction, error) {
	txOpts, err := erc1155.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc1155.abi.SafeTransferFrom(
		txOpts,
		erc1155.helper.GetSignerAddress(),
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(amount)),
		[]byte{},
	); err != nil {
		return nil, err
	} else {
		return erc1155.helper.awaitTx(tx.Hash())
	}
}

// Burn an amount of a specified NFT from the connected wallet.
//
// tokenId: tokenID of the token to burn
//
// amount: number of NFTs of the token ID to burn
//
// returns: the transaction receipt of the burn
//
// Example
//
//	tokenId := 0
//	amount := 1
//	tx, err := contract.Burn(context.Background(), tokenId, amount)
func (erc1155 *ERC1155) Burn(ctx context.Context, tokenId int, amount int) (*types.Transaction, error) {
	address := erc1155.helper.GetSignerAddress()
	txOpts, err := erc1155.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc1155.abi.Burn(
		txOpts,
		address,
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(amount)),
	); err != nil {
		return nil, err
	} else {
		return erc1155.helper.awaitTx(tx.Hash())
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
func (erc1155 *ERC1155) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error) {
	txOpts, err := erc1155.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc1155.abi.SetApprovalForAll(
		txOpts,
		common.HexToAddress(operator),
		approved,
	); err != nil {
		return nil, err
	} else {
		return erc1155.helper.awaitTx(tx.Hash())
	}
}

func (erc1155 *ERC1155) getTokenMetadata(ctx context.Context, tokenId int) (*NFTMetadata, error) {
	if uri, err := erc1155.abi.Uri(
		&bind.CallOpts{Context: ctx},
		big.NewInt(int64(tokenId)),
	); err != nil {
		return nil, &notFoundError{
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

func fetchEditionsByTokenId(ctx context.Context, erc1155 *ERC1155, tokenIds []*big.Int) ([]*EditionMetadata, error) {
	total := len(tokenIds)

	ch := make(chan *EditionResult)
	// fetch all nfts in parallel
	for i := 0; i < total; i++ {
		go func(id int) {
			if nft, err := erc1155.Get(ctx, id); err == nil {
				ch <- &EditionResult{nft, nil}
			} else {
				fmt.Println(err)
				ch <- &EditionResult{nil, err}
			}
		}(i)
	}
	// wait for all goroutines to emit
	results := make([]*EditionResult, total)
	for i := range results {
		results[i] = <-ch
	}
	// filter out errors
	nfts := []*EditionMetadata{}
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
