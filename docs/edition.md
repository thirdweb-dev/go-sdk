
## Edition

You can access the Edition interface from the SDK as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetEdition("{{contract_address}}")
```

```go
type Edition struct {
    *ERC1155
    Signature *ERC1155SignatureMinting
}
```

### func \(\*Edition\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition.go#L65>)

```go
func (edition *Edition) Mint(metadataWithSupply *EditionMetadataInput) (*types.Transaction, error)
```

Mint an NFT to the connected wallet\.

metadataWithSupply: nft metadata with supply of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintAdditionalSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition.go#L121>)

```go
func (edition *Edition) MintAdditionalSupply(tokenId int, additionalSupply int) (*types.Transaction, error)
```

Mint additionaly supply of a token to the connected wallet\.

tokenId: token ID to mint additional supply of

additionalSupply: additional supply to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintAdditionalSupplyTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition.go#L135>)

```go
func (edition *Edition) MintAdditionalSupplyTo(to string, tokenId int, additionalSupply int) (*types.Transaction, error)
```

Mint additional supply of a token to the specified wallet\.

to: address of the wallet to mint NFTs to

tokenId: token Id to mint additional supply of

additionalySupply: additional supply to mint

returns: the transaction receipt of the mint

### func \(\*Edition\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition.go#L183>)

```go
func (edition *Edition) MintBatchTo(to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error)
```

Mint a batch of tokens to various wallets\.

to: address of the wallet to mint NFTs to

metadatasWithSupply: list of NFT metadatas with supplies to mint

returns: the transaction receipt of the mint

#### Example

```
metadatasWithSupply := []*thirdweb.EditionMetadataInput{
	&thirdweb.EditionMetadataInput{
		Metadata: &thirdweb.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
		},
		Supply: 100,
	},
	&thirdweb.EditionMetadataInput{
		Metadata: &thirdweb.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
		},
		Supply: 100,
	},
}

tx, err := contract.MintBatchTo("{{wallet_address}}", metadatasWithSupply)
```

### func \(\*Edition\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition.go#L93>)

```go
func (edition *Edition) MintTo(address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the specified wallet\.

address: the wallet address to mint the NFT to

metadataWithSupply: nft metadata with supply of the NFT to mint

returns: the transaction receipt of the mint

#### Example

```
image, err := os.Open("path/to/image.jpg")
defer image.Close()

metadataWithSupply := &thirdweb.EditionMetadataInput{
	Metadata: &thirdweb.NFTMetadataInput{
		Name: "Cool NFT",
		Description: "This is a cool NFT",
		Image: image,
	},
	Supply: 100,
}

tx, err := contract.MintTo("{{wallet_address}}", metadataWithSupply)
```
