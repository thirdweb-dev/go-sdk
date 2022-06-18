
## Edition Drop

You can access the Edition Drop interface from the SDK as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetEditionDrop("{{contract_address}}")
```

```go
type EditionDrop struct {
    *ERC1155
}
```

### func \(\*EditionDrop\) [Claim](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition_drop.go#L138>)

```go
func (drop *EditionDrop) Claim(tokenId int, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

tokenId: the token ID of the NFT to claim

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*EditionDrop\) [ClaimTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition_drop.go#L160>)

```go
func (drop *EditionDrop) ClaimTo(destinationAddress string, tokenId int, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

tokenId: the token ID of the NFT to claim

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

#### Example

```
address = "{{wallet_address}}"
tokenId = 0
quantity = 1

tx, err := contract.ClaimTo(address, tokenId, quantity)
```

### func \(\*EditionDrop\) [CreateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition_drop.go#L94>)

```go
func (drop *EditionDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Create a batch of NFTs on this contract\.

metadatas: a list of the metadatas of the NFTs to create

returns: the transaction receipt of the batch creation

#### Example

```
image0, err := os.Open("path/to/image/0.jpg")
defer image0.Close()

image1, err := os.Open("path/to/image/1.jpg")
defer image1.Close()

metadatasWithSupply := []*thirdweb.EditionMetadataInput{
	&thirdweb.EditionMetadataInput{
		Metadata: &thirdweb.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
			Image: image0,
		},
		Supply: 100,
	},
	&thirdweb.EditionMetadataInput{
		Metadata: &thirdweb.NFTMetadataInput{
			Name: "Cool NFT",
			Description: "This is a cool NFT",
			Image: image1,
		},
		Supply: 100,
	},
}

tx, err := contract.MintBatchTo("{{wallet_address}}", metadatasWithSupply)
```

## type [EditionMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L45-L48>)

```go
type EditionMetadata struct {
    Metadata *NFTMetadata
    Supply   int
}
```

## type [EditionMetadataInput](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L57-L60>)

```go
type EditionMetadataInput struct {
    Metadata *NFTMetadataInput
    Supply   int
}
```

## type [EditionMetadataOwner](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L50-L55>)

```go
type EditionMetadataOwner struct {
    Metadata      *NFTMetadata
    Supply        int
    Owner         string
    QuantityOwned int
}
```

## type [EditionResult](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc1155.go#L23-L26>)

```go
type EditionResult struct {}
```
