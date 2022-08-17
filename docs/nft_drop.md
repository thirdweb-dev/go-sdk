
## NFT Drop

You can access the NFT Drop interface from the SDK as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetNFTDrop("{{contract_address}}")
```

```go
type NFTDrop struct {
    *ERC721
    ClaimConditions *NFTDropClaimConditions
    Encoder         *ContractEncoder
}
```

### func \(\*NFTDrop\) [Claim](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L238>)

```go
func (drop *NFTDrop) Claim(quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*NFTDrop\) [ClaimTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L257>)

```go
func (drop *NFTDrop) ClaimTo(destinationAddress string, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet\.

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

#### Example

```
address := "{{wallet_address}}"
quantity = 1

tx, err := contract.ClaimTo(address, quantity)
```

### func \(\*NFTDrop\) [CreateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L195>)

```go
func (drop *NFTDrop) CreateBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
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

metadatas := []*thirdweb.NFTMetadataInput{
	&thirdweb.NFTMetadataInput{
		Name: "Cool NFT",
		Description: "This is a cool NFT",
		Image: image1
	}
	&thirdweb.NFTMetadataInput{
		Name: "Cool NFT 2",
		Description: "This is also a cool NFT",
		Image: image2
	}
}

tx, err := contract.CreateBatch(metadatas)
```

### func \(\*NFTDrop\) [GetAllClaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L123>)

```go
func (drop *NFTDrop) GetAllClaimed() ([]*NFTMetadataOwner, error)
```

Get a list of all the NFTs that have been claimed from this contract\.

returns: a list of the metadatas of the claimed NFTs

#### Example

```
claimedNfts, err := contract.GetAllClaimed()
firstOwner := claimedNfts[0].Owner
```

### func \(\*NFTDrop\) [GetAllUnclaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L147>)

```go
func (drop *NFTDrop) GetAllUnclaimed() ([]*NFTMetadata, error)
```

Get a list of all the NFTs on this contract that have not yet been claimed\.

returns: a list of the metadatas of the unclaimed NFTs

#### Example

```
unclaimedNfts, err := contract.GetAllUnclaimed()
firstNftName := unclaimedNfts[0].Name
```

### func \(\*NFTDrop\) [GetOwned](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L78>)

```go
func (nft *NFTDrop) GetOwned(address string) ([]*NFTMetadataOwner, error)
```

Get the metadatas of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the metadata of all the NFTs owned by the address

#### Example

```
owner := "{{wallet_address}}"
nfts, err := contract.GetOwned(owner)
name := nfts[0].Metadata.Name
```

### func \(\*NFTDrop\) [GetOwnedTokenIDs](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L95>)

```go
func (nft *NFTDrop) GetOwnedTokenIDs(address string) ([]*big.Int, error)
```

Get the tokenIds of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the tokenIds of all the NFTs owned by the address
