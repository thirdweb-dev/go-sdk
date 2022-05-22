
## NFT Collection

You can access the NFT Collection interface from the SDK as follows:

```
contract, err := sdk.GetNFTCollection("{{contract_address}}")
```

```go
type NFTCollection struct {
    *ERC721
}
```

### func \(\*NFTCollection\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_collection.go#L49>)

```go
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the connected wallet\.

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

#### Example

```
nft.Mint(&thirdweb.NFTMetadataInput{ Name: "NFT" })
```

### func \(\*NFTCollection\) [MintBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_collection.go#L91>)

```go
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

#### MintBatch

Mint a batch of new NFTs to the connected wallet

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_collection.go#L105>)

```go
func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

#### MintBatchTo

Mint a batch of new NFTs to the specified wallet

to: the wallet address to mint to

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_collection.go#L66>)

```go
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the specified wallet\.

address: the wallet address to mint to

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

#### Example

```
nft.MintTo("0x0dcd5d886577d5081b0c52e242ef29e70be3e7bc", &thirdweb.NFTMetadataInput{ Name: "NFT" })
```
