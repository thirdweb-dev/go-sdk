
## Multiwrap

```go
type Multiwrap struct {
    *ERC721
}
```

### func \(\*Multiwrap\) [GetWrappedContents](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/multiwrap.go#L55>)

```go
func (multiwrap *Multiwrap) GetWrappedContents(wrappedTokenId int) (*MultiwrapBundle, error)
```

Get the contents of a wrapped token bundle\.

wrappedTokenId: the ID of the wrapped token bundle

returns: the contents of the wrapped token bundle

#### Example

```
tokenId := 0
contents, err := contract.GetWrappedContents(tokenId)
erc20Tokens := contents.Erc20Tokens
erc721Tokens := contents.Erc721Tokens
erc1155Tokens := contents.Erc1155Tokens
```

### func \(\*Multiwrap\) [Unwrap](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/multiwrap.go#L186>)

```go
func (multiwrap *Multiwrap) Unwrap(wrappedTokenId int, recipientAddress string) (*types.Transaction, error)
```

#### Unwrap a wrapped token bundle into its contents

wrappedTokenId: the ID of the wrapped token bundle

recipientAddress: the optional address to send the wrapped token to

returns: the contents of the wrapped token bundle

#### Example

```
tokenId := 0
tx, err := contract.Unwrap(tokenId, "")
```

### func \(\*Multiwrap\) [Wrap](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/multiwrap.go#L141>)

```go
func (multiwrap *Multiwrap) Wrap(contents *MultiwrapBundle, wrappedTokenMetadata interface{}, recipientAddress string) (*types.Transaction, error)
```

## type [MultiwrapBundle](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L215-L219>)

```go
type MultiwrapBundle struct {
    ERC20Tokens   []*MultiwrapERC20
    ERC721Tokens  []*MultiwrapERC721
    ERC1155Tokens []*MultiwrapERC1155
}
```

## type [MultiwrapERC1155](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L209-L213>)

```go
type MultiwrapERC1155 struct {
    ContractAddress string
    TokenId         int
    Quantity        int
}
```

## type [MultiwrapERC20](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L199-L202>)

```go
type MultiwrapERC20 struct {
    ContractAddress string
    Quantity        float64
}
```

## type [MultiwrapERC721](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L204-L207>)

```go
type MultiwrapERC721 struct {
    ContractAddress string
    TokenId         int
}
```
