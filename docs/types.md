
## Types

```go
type SDKOptions struct {
    PrivateKey string
    GatewayUrl string
}
```

## type [Signature1155PayloadInput](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L152-L163>)

```go
type Signature1155PayloadInput struct {
    To                   string
    Price                float64
    CurrencyAddress      string
    MintStartTime        int
    MintEndTime          int
    PrimarySaleRecipient string
    Metadata             *NFTMetadataInput
    RoyaltyRecipient     string
    RoyaltyBps           int
    Quantity             int
}
```

## type [Signature1155PayloadInputWithTokenId](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L165-L177>)

```go
type Signature1155PayloadInputWithTokenId struct {
    To                   string
    Price                float64
    CurrencyAddress      string
    MintStartTime        int
    MintEndTime          int
    PrimarySaleRecipient string
    Metadata             *NFTMetadataInput
    TokenId              int
    RoyaltyRecipient     string
    RoyaltyBps           int
    Quantity             int
}
```

## type [Signature1155PayloadOutput](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L179-L193>)

```go
type Signature1155PayloadOutput struct {
    To                   string
    Price                float64
    CurrencyAddress      string
    MintStartTime        int
    MintEndTime          int
    PrimarySaleRecipient string
    Metadata             *NFTMetadataInput
    RoyaltyRecipient     string
    RoyaltyBps           int
    TokenId              int
    Quantity             int
    Uri                  string
    Uid                  [32]byte
}
```

## type [Signature721PayloadInput](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L121-L131>)

```go
type Signature721PayloadInput struct {
    To                   string
    Price                float64
    CurrencyAddress      string
    MintStartTime        int
    MintEndTime          int
    PrimarySaleRecipient string
    Metadata             *NFTMetadataInput
    RoyaltyRecipient     string
    RoyaltyBps           int
}
```

## type [Signature721PayloadOutput](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L133-L145>)

```go
type Signature721PayloadOutput struct {
    To                   string
    Price                float64
    CurrencyAddress      string
    MintStartTime        int
    MintEndTime          int
    PrimarySaleRecipient string
    Metadata             *NFTMetadataInput
    RoyaltyRecipient     string
    RoyaltyBps           int
    Uri                  string
    Uid                  [32]byte
}
```

## type [SignedPayload1155](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L195-L198>)

```go
type SignedPayload1155 struct {
    Payload   *Signature1155PayloadOutput
    Signature []byte
}
```

## type [SignedPayload721](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L147-L150>)

```go
type SignedPayload721 struct {
    Payload   *Signature721PayloadOutput
    Signature []byte
}
```
