
## Edition Drop
You can access this interface through the SDK with `sdk.GetEditionDrop(address)`.


```go
type EditionDrop struct {
    *ERC1155
    // contains filtered or unexported fields
}
```

### func \(\*EditionDrop\) [Claim](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/edition_drop.go#L96>)

```go
func (drop *EditionDrop) Claim(tokenId int, quantity int) (*types.Transaction, error)
```

#### Claim

Claim NFTs from this contract to the connect wallet

tokenId: the token ID of the NFT to claim

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*EditionDrop\) [ClaimTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/edition_drop.go#L112>)

```go
func (drop *EditionDrop) ClaimTo(destinationAddress string, tokenId int, quantity int) (*types.Transaction, error)
```

#### ClaimTo

Claim NFTs from this contract to the connect wallet

tokenId: the token ID of the NFT to claim

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*EditionDrop\) [CreateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/edition_drop.go#L54>)

```go
func (drop *EditionDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

#### CreateBatch

Create a batch of NFTs on this contract

metadatas: a list of the metadatas of the NFTs to create

returns: the transaction receipt of the batch creation

## type [EditionMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L44-L47>)

```go
type EditionMetadata struct {
    Metadata *NFTMetadata
    Supply   int
}
```

## type [EditionMetadataInput](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L56-L59>)

```go
type EditionMetadataInput struct {
    Metadata *NFTMetadataInput
    Supply   int
}
```

## type [EditionMetadataOwner](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L49-L54>)

```go
type EditionMetadataOwner struct {
    Metadata      *NFTMetadata
    Supply        int
    Owner         string
    QuantityOwned int
}
```

## type [EditionResult](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc1155.go#L21-L24>)

```go
type EditionResult struct {
    // contains filtered or unexported fields
}
```
