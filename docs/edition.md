
## Edition
You can access this interface through the SDK with `sdk.GetEdition(address)`.


```go
type Edition struct {
    *ERC1155
}
```

### func \(\*Edition\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/edition.go#L39>)

```go
func (edition *Edition) Mint(metadataWithSupply *EditionMetadataInput) (*types.Transaction, error)
```

#### Mint

Mint an NFT to the connected wallet

metadataWithSupply: nft metadata with supply of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintAdditionalSupply](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/edition.go#L83>)

```go
func (edition *Edition) MintAdditionalSupply(tokenId int, additionalSupply int) (*types.Transaction, error)
```

#### MintAdditionalSupply

Mint additionaly supply of a token to the connected wallet

tokenId: token ID to mint additional supply of

additionalSupply: additional supply to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintAdditionalSupplyTo](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/edition.go#L97>)

```go
func (edition *Edition) MintAdditionalSupplyTo(to string, tokenId int, additionalSupply int) (*types.Transaction, error)
```

#### MintAdditionalSupplyTo

to: address of the wallet to mint NFTs to

tokenId: token Id to mint additional supply of

additionalySupply: additional supply to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/edition.go#L124>)

```go
func (edition *Edition) MintBatchTo(to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error)
```

#### MintBatchTo

to: address of the wallet to mint NFTs to

metadatasWithSupply: list of NFT metadatas with supplies to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/edition.go#L53>)

```go
func (edition *Edition) MintTo(address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error)
```

#### MintTo

Mint a new NFT to the specified wallet

address: the wallet address to mint the NFT to

metadataWithSupply: nft metadata with supply of the NFT to mint

returns: the transaction receipt of the mint

## type [EditionMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/types.go#L44-L47>)

```go
type EditionMetadata struct {
    Metadata *NFTMetadata
    Supply   int
}
```

## type [EditionMetadataInput](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/types.go#L56-L59>)

```go
type EditionMetadataInput struct {
    Metadata *NFTMetadataInput
    Supply   int
}
```

## type [EditionMetadataOwner](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/types.go#L49-L54>)

```go
type EditionMetadataOwner struct {
    Metadata      *NFTMetadata
    Supply        int
    Owner         string
    QuantityOwned int
}
```

## type [EditionResult](<https://github.com/thirdweb-dev/go-sdk/blob/master/pkg/thirdweb/erc1155.go#L19-L22>)

```go
type EditionResult struct {
    // contains filtered or unexported fields
}
```
