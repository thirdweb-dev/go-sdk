
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
    ClaimConditions *EditionDropClaimConditions
    Encoder         *ContractEncoder
}
```

### func \(\*EditionDrop\) [Claim](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition_drop.go#L150>)

```go
func (drop *EditionDrop) Claim(ctx context.Context, tokenId int, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

tokenId: the token ID of the NFT to claim

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*EditionDrop\) [ClaimTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition_drop.go#L172>)

```go
func (drop *EditionDrop) ClaimTo(ctx context.Context, destinationAddress string, tokenId int, quantity int) (*types.Transaction, error)
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

tx, err := contract.ClaimTo(context.Background(), address, tokenId, quantity)
```

### func \(\*EditionDrop\) [CreateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/edition_drop.go#L103>)

```go
func (drop *EditionDrop) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error)
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

tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatasWithSupply)
```
