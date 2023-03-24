
## ERC721 Signature Minting

You can access this interface from the NFT Collection contract under the signature interface.

```go
type ERC721SignatureMinting struct {
    Helper *contractHelper
}
```

### func \(\*ERC721SignatureMinting\) [Generate](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L315>)

```go
func (signature *ERC721SignatureMinting) Generate(ctx context.Context, payloadToSign *Signature721PayloadInput) (*SignedPayload721, error)
```

#### Generate a new payload from the given data

payloadToSign: the payload containing the data for the signature mint

returns: the payload signed by the minter's private key

#### Example

```
payload := &thirdweb.Signature721PayloadInput{
	To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6", // address to mint to
	Price:                0,                                            // cost of minting
	CurrencyAddress:      "0x0000000000000000000000000000000000000000", // currency to pay in order to mint
	MintStartTime:        0,                                            // time where minting is allowed to start (epoch seconds)
	MintEndTime:          100000000000000,                              // time when this signature expires (epoch seconds)
	PrimarySaleRecipient: "0x0000000000000000000000000000000000000000", // address to receive the primary sales of this mint
	Metadata: &thirdweb.NFTMetadataInput{																// metadata of the NFT to mint
 		Name:  "ERC721 Sigmint!",
	},
	RoyaltyRecipient: "0x0000000000000000000000000000000000000000",     // address to receive royalties of this mint
	RoyaltyBps:       0,                                                // royalty cut of this mint in basis points
}

signedPayload, err := contract.Signature.Generate(payload)
```

### func \(\*ERC721SignatureMinting\) [GenerateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L364>)

```go
func (signature *ERC721SignatureMinting) GenerateBatch(ctx context.Context, payloadsToSign []*Signature721PayloadInput) ([]*SignedPayload721, error)
```

#### Generate a batch of new payload from the given data

payloadToSign: the payloads containing the data for the signature mint

returns: the payloads signed by the minter's private key

#### Example

```
payload := []*thirdweb.Signature721PayloadInput{
	&thirdweb.Signature721PayloadInput{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &thirdweb.NFTMetadataInput{
 			Name:  "ERC721 Sigmint!",
 			Image: imageFile,
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
	},
	&thirdweb.Signature721PayloadInput{
		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
		Price:                0,
		CurrencyAddress:      "0x0000000000000000000000000000000000000000",
		MintStartTime:        0,
		MintEndTime:          100000000000000,
		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
		Metadata: &thirdweb.NFTMetadataInput{
 			Name:  "ERC721 Sigmint!",
 			Image: imageFile,
		},
		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
		RoyaltyBps:       0,
	},
}

signedPayload, err := contract.Signature.GenerateBatch(context.Background(), payload)
```

### func \(\*ERC721SignatureMinting\) [GenerateBatchWithUris](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L446>)

```go
func (signature *ERC721SignatureMinting) GenerateBatchWithUris(ctx context.Context, payloadsToSign []*Signature721PayloadInputWithUri) ([]*SignedPayload721, error)
```

### func \(\*ERC721SignatureMinting\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L69>)

```go
func (signature *ERC721SignatureMinting) Mint(ctx context.Context, signedPayload *SignedPayload721) (*types.Transaction, error)
```

Deprecated: use MintAndAwait

### func \(\*ERC721SignatureMinting\) [MintAndAwait](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L73>)

```go
func (signature *ERC721SignatureMinting) MintAndAwait(ctx context.Context, signedPayload *SignedPayload721) (*types.Transaction, error)
```

### func \(\*ERC721SignatureMinting\) [MintBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L147>)

```go
func (signature *ERC721SignatureMinting) MintBatch(ctx context.Context, signedPayloads []*SignedPayload721) (*types.Transaction, error)
```

Mint a batch of token with the data in given payload.

signedPayload: the list of payloads signed by the minters private key being used to mint

returns: the transaction receipt of the batch mint

#### Example

```
// Learn more about how to craft multiple payloads in the GenerateBatch() function
signedPayloads, err := contract.Signature.GenerateBatch(payloads)
tx, err := contract.Signature.MintBatch(context.Background(), signedPayloads)
```

### func \(\*ERC721SignatureMinting\) [MintWithOpts](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L88>)

```go
func (signature *ERC721SignatureMinting) MintWithOpts(ctx context.Context, signedPayload *SignedPayload721, txOpts *bind.TransactOpts) (*types.Transaction, error)
```

### func \(\*ERC721SignatureMinting\) [Verify](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc721_signature_minting.go#L264>)

```go
func (signature *ERC721SignatureMinting) Verify(ctx context.Context, signedPayload *SignedPayload721) (bool, error)
```

#### Verify that a signed payload is valid

signedPayload: the payload to verify

returns: true if the payload is valid, otherwise false.

#### Example

```
// Learn more about how to craft a payload in the Generate() function
signedPayload, err := contract.Signature.Generate(payload)
isValid, err := contract.Signature.Verify(signedPayload)
```
