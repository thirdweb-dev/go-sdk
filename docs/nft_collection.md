

```go
type NFTCollection struct {
    *ERC721
    // contains filtered or unexported fields
}
```

### func \(\*NFTCollection\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/nft_collection.go#L42>)

```go
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error)
```

#### Mint

Mint a new NFT to the connected wallet

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatch](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/nft_collection.go#L81>)

```go
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

#### MintBatch

Mint a batch of new NFTs to the connected wallet

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/nft_collection.go#L95>)

```go
func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

#### MintBatchTo

Mint a batch of new NFTs to the specified wallet

to: the wallet address to mint to

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/nft_collection.go#L56>)

```go
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error)
```

#### Mint

Mint a new NFT to the specified wallet

address: the wallet address to mint to

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint
