
## NFT Drop
You can access this interface through the SDK with `sdk.GetNFTDrop(address)`.


```go
type NFTDrop struct {
    *ERC721
    // contains filtered or unexported fields
}
```

### func \(\*NFTDrop\) [Claim](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_drop.go#L141>)

```go
func (drop *NFTDrop) Claim(quantity int) (*types.Transaction, error)
```

#### Claim

Claim NFTs from this contract to the connect wallet

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*NFTDrop\) [ClaimTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_drop.go#L155>)

```go
func (drop *NFTDrop) ClaimTo(destinationAddress string, quantity int) (*types.Transaction, error)
```

#### ClaimTo

Claim NFTs from this contract to the connect wallet

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*NFTDrop\) [CreateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_drop.go#L100>)

```go
func (drop *NFTDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

#### CreateBatch

Create a batch of NFTs on this contract

metadatas: a list of the metadatas of the NFTs to create

returns: the transaction receipt of the batch creation

### func \(\*NFTDrop\) [GetAllClaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_drop.go#L52>)

```go
func (drop *NFTDrop) GetAllClaimed() ([]*NFTMetadataOwner, error)
```

#### GetAllClaimed

Get a list of all the NFTs that have been claimed from this contract

returns: a list of the metadatas of the claimed NFTs

### func \(\*NFTDrop\) [GetAllUnclaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/nft_drop.go#L73>)

```go
func (drop *NFTDrop) GetAllUnclaimed() ([]*NFTMetadata, error)
```

#### GetAllUnclaimed

Get a list of all the NFTs on this contract that have not yet been claimed

returns: a list of the metadatas of the unclaimed NFTs

## type [NFTMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L17-L27>)

```go
type NFTMetadata struct {
    Id              *big.Int    `json:"id"`
    Uri             string      `json:"uri"`
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Image           interface{} `json:"image"`
    ExternalUrl     string      `json:"external_url"`
    AnimationUrl    string      `json:"animation_url"`
    BackgroundColor string      `json:"background_color"`
    Properties      interface{} `json:"properties,omitempty"`
}
```

## type [NFTMetadataInput](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L29-L37>)

```go
type NFTMetadataInput struct {
    Name            string      `json:"name"`
    Description     string      `json:"description"`
    Image           interface{} `json:"image"`
    ExternalUrl     string      `json:"external_url"`
    AnimationUrl    string      `json:"animation_url"`
    BackgroundColor string      `json:"background_color"`
    Properties      interface{} `json:"properties,omitempty"`
}
```

## type [NFTMetadataOwner](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L39-L42>)

```go
type NFTMetadataOwner struct {
    Metadata *NFTMetadata
    Owner    string
}
```

## type [NFTResult](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc721.go#L21-L24>)

```go
type NFTResult struct {
    // contains filtered or unexported fields
}
```

## type [NativeToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L113-L118>)

```go
type NativeToken struct {
    // contains filtered or unexported fields
}
```
