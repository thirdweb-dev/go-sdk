package thirdweb

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// This interface is currently support by the Edition and Edition Drop contracts.
// You can access all of its functions through an Edition or Edition Drop contract instance.
type ERC1155Standard struct {
	erc1155 *ERC1155
}

func newERC1155Standard(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC1155Standard, error) {
	erc1155, err := newERC1155(provider, address, privateKey, storage)
	if err != nil {
		return nil, err
	}
	
	return &ERC1155Standard{
		erc1155,
	}, nil
	
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
func (erc1155 *ERC1155Standard) Get(ctx context.Context, tokenId int) (*EditionMetadata, error) {
	return erc1155.erc1155.Get(ctx, tokenId)
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
func (erc1155 *ERC1155Standard) GetAll(ctx context.Context) ([]*EditionMetadata, error) {
	return erc1155.erc1155.GetAll(ctx)
}

// Get the total number of NFTs on this contract.
//
// returns: the total number of NFTs on this contract
func (erc1155 *ERC1155Standard) GetTotalCount(ctx context.Context) (int, error) {
	return erc1155.erc1155.GetTotalCount(ctx)
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
func (erc1155 *ERC1155Standard) GetOwned(ctx context.Context, address string) ([]*EditionMetadataOwner, error) {
	return erc1155.erc1155.GetOwned(ctx, address)
}

// Get the total number of NFTs of a specific token ID.
//
// tokenId: the token ID to check the total supply of
//
// returns: the supply of NFTs on the specified token ID
func (erc1155 *ERC1155Standard) TotalSupply(ctx context.Context, tokenId int) (int, error) {
	return erc1155.erc1155.TotalSupply(ctx, tokenId)
}

// Get the NFT balance of the connected wallet for a specific token ID.
//
// tokenId: the token ID of a specific token to check the balance of
//
// returns: the number of NFTs of the specified token ID owned by the connected wallet
func (erc1155 *ERC1155Standard) Balance(ctx context.Context, tokenId int) (int, error) {
	return erc1155.erc1155.Balance(ctx, tokenId)
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
func (erc1155 *ERC1155Standard) BalanceOf(ctx context.Context, address string, tokenId int) (int, error) {
	return erc1155.erc1155.BalanceOf(ctx, address, tokenId)
}

// Check whether an operator address is approved for all operations of a specifc addresses assets.
//
// address: the address whose assets are to be checked
//
// operator: the address of the operator to check
//
// returns: true if the operator is approved for all operations of the assets, otherwise false
func (erc1155 *ERC1155Standard) IsApproved(ctx context.Context, address string, operator string) (bool, error) {
	return erc1155.erc1155.IsApproved(ctx, address, operator)
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
func (erc1155 *ERC1155Standard) Transfer(ctx context.Context, to string, tokenId int, amount int) (*types.Transaction, error) {
	return erc1155.erc1155.Transfer(ctx, to, tokenId, amount)
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
func (erc1155 *ERC1155Standard) Burn(ctx context.Context, tokenId int, amount int) (*types.Transaction, error) {
	return erc1155.erc1155.Burn(ctx, tokenId, amount)
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
func (erc1155 *ERC1155Standard) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error) {
	return erc1155.erc1155.SetApprovalForAll(ctx, operator, approved)
}
