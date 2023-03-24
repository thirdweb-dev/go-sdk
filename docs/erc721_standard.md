
## ERC721 Standard

This interface is currently support by the NFT Collection and NFT Drop contracts. You can access all of its functions through an NFT Collection or NFT Drop contract instance.

```go
type ERC721Standard struct {}
```

### func \(\*ERC721Standard\) [Balance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L83>)

```go
func (erc721 *ERC721Standard) Balance(ctx context.Context) (int, error)
```

Get the NFT balance of the connected wallet.

returns: the number of NFTs on this contract owned by the connected wallet

### func \(\*ERC721Standard\) [BalanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L97>)

```go
func (erc721 *ERC721Standard) BalanceOf(ctx context.Context, address string) (int, error)
```

Get the NFT balance of a specific wallet.

address: the address of the wallet to get the NFT balance of

returns: the number of NFTs on this contract owned by the specified wallet

#### Example

```
address := "{{wallet_address}}"
balance, err := contract.BalanceOf(context.Background(), address)
```

### func \(\*ERC721Standard\) [Burn](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L140>)

```go
func (erc721 *ERC721Standard) Burn(ctx context.Context, tokenId int) (*types.Transaction, error)
```

Burn a specified NFT from the connected wallet.

tokenId: tokenID of the token to burn

returns: the transaction receipt of the burn

#### Example

```
tokenId := 0
tx, err := contract.Burn(context.Background(), tokenId)
```

### func \(\*ERC721Standard\) [Get](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L40>)

```go
func (erc721 *ERC721Standard) Get(ctx context.Context, tokenId int) (*NFTMetadataOwner, error)
```

Get metadata for a token.

tokenId: token ID of the token to get the metadata for

returns: the metadata for the NFT and its owner

#### Example

```
nft, err := contract.Get(context.Background(), 0)
 owner := nft.Owner
	name := nft.Metadata.Name
```

### func \(\*ERC721Standard\) [GetAll](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L53>)

```go
func (erc721 *ERC721Standard) GetAll(ctx context.Context) ([]*NFTMetadataOwner, error)
```

Get the metadata of all the NFTs on this contract.

returns: the metadata of all the NFTs on this contract

#### Example

```
nfts, err := contract.GetAll(context.Background())
ownerOne := nfts[0].Owner
nameOne := nfts[0].Metadata.Name
```

### func \(\*ERC721Standard\) [GetTotalCount](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L60>)

```go
func (erc721 *ERC721Standard) GetTotalCount(ctx context.Context) (int, error)
```

Get the total number of NFTs on this contract.

returns: the total number of NFTs on this contract

### func \(\*ERC721Standard\) [IsApproved](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L108>)

```go
func (erc721 *ERC721Standard) IsApproved(ctx context.Context, address string, operator string) (bool, error)
```

Check whether an operator address is approved for all operations of a specifc addresses assets.

address: the address whose assets are to be checked

operator: the address of the operator to check

returns: true if the operator is approved for all operations of the assets, otherwise false

### func \(\*ERC721Standard\) [OwnerOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L69>)

```go
func (erc721 *ERC721Standard) OwnerOf(ctx context.Context, tokenId int) (string, error)
```

Get the owner of an NFT.

tokenId: the token ID of the NFT to get the owner of

returns: the owner of the NFT

### func \(\*ERC721Standard\) [SetApprovalForAll](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L153>)

```go
func (erc721 *ERC721Standard) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error)
```

Set the approval for all operations of a specific address's assets.

address: the address whose assets are to be approved

operator: the address of the operator to set the approval for

approved: true if the operator is approved for all operations of the assets, otherwise false

returns: the transaction receipt of the approval

### func \(\*ERC721Standard\) [SetApprovalForToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L165>)

```go
func (erc721 *ERC721Standard) SetApprovalForToken(ctx context.Context, operator string, tokenId int) (*types.Transaction, error)
```

Approve an operator for the NFT owner, which allows the operator to call transferFrom or safeTransferFrom for the specified token.

operator: the address of the operator to approve

tokenId: the token ID of the NFT to approve

returns: the transaction receipt of the approval

### func \(\*ERC721Standard\) [TotalSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L76>)

```go
func (erc721 *ERC721Standard) TotalSupply(ctx context.Context) (int, error)
```

Get the total number of NFTs on this contract.

returns: the supply of NFTs on this contract

### func \(\*ERC721Standard\) [Transfer](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_standard.go#L126>)

```go
func (erc721 *ERC721Standard) Transfer(ctx context.Context, to string, tokenId int) (*types.Transaction, error)
```

Transfer a specified token from the connected wallet to a specified address.

to: wallet address to transfer the tokens to

tokenId: the token ID of the NFT to transfer

returns: the transaction of the NFT transfer

#### Example

```
to := "0x..."
tokenId := 0

tx, err := contract.Transfer(context.Background(), to, tokenId)
```
