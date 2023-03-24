
## NFT Drop

You can access the NFT Drop interface from the SDK as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetNFTDrop("{{contract_address}}")
```

```go
type NFTDrop struct {
    *ERC721Standard
    Abi             *abi.DropERC721
    Helper          *contractHelper
    ClaimConditions *NFTDropClaimConditions
    Encoder         *NFTDropEncoder
    Events          *ContractEvents
}
```

### func \(\*NFTDrop\) [Claim](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L227>)

```go
func (drop *NFTDrop) Claim(ctx context.Context, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet.

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

### func \(\*NFTDrop\) [ClaimTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L245>)

```go
func (drop *NFTDrop) ClaimTo(ctx context.Context, destinationAddress string, quantity int) (*types.Transaction, error)
```

Claim NFTs from this contract to the connect wallet.

destinationAddress: the address of the wallet to claim the NFTs to

quantity: the number of NFTs to claim

returns: the transaction receipt of the claim

#### Example

```
address := "{{wallet_address}}"
quantity = 1

tx, err := contract.ClaimTo(context.Background(), address, quantity)
```

### func \(\*NFTDrop\) [CreateBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L218>)

```go
func (drop *NFTDrop) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Create a batch of NFTs on this contract.

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

tx, err := contract.CreateBatch(context.Background(), metadatas)
```

### func \(\*NFTDrop\) [GetAllClaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L138>)

```go
func (drop *NFTDrop) GetAllClaimed(ctx context.Context) ([]*NFTMetadataOwner, error)
```

Get a list of all the NFTs that have been claimed from this contract.

returns: a list of the metadatas of the claimed NFTs

#### Example

```
claimedNfts, err := contract.GetAllClaimed(context.Background())
firstOwner := claimedNfts[0].Owner
```

### func \(\*NFTDrop\) [GetAllUnclaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L150>)

```go
func (drop *NFTDrop) GetAllUnclaimed(ctx context.Context) ([]*NFTMetadata, error)
```

Get a list of all the NFTs on this contract that have not yet been claimed.

returns: a list of the metadatas of the unclaimed NFTs

#### Example

```
unclaimedNfts, err := contract.GetAllUnclaimed(context.Background())
firstNftName := unclaimedNfts[0].Name
```

### func \(\*NFTDrop\) [GetClaimArguments](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L249-L256>)

```go
func (drop *NFTDrop) GetClaimArguments(ctx context.Context, destinationAddress string, quantity int) (*ClaimArguments, error)
```

### func \(\*NFTDrop\) [GetClaimIneligibilityReasons](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L186>)

```go
func (drop *NFTDrop) GetClaimIneligibilityReasons(ctx context.Context, quantity int, addressToCheck string) ([]ClaimEligibility, error)
```

### func \(\*NFTDrop\) [GetClaimInfo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L182>)

```go
func (drop *NFTDrop) GetClaimInfo(ctx context.Context, address string) (*ClaimInfo, error)
```

### func \(\*NFTDrop\) [GetOwned](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L88>)

```go
func (nft *NFTDrop) GetOwned(ctx context.Context, address string) ([]*NFTMetadataOwner, error)
```

Get the metadatas of all the NFTs owned by a specific address.

address: the address of the owner of the NFTs

returns: the metadata of all the NFTs owned by the address

#### Example

```
owner := "{{wallet_address}}"
nfts, err := contract.GetOwned(context.Background(), owner)
name := nfts[0].Metadata.Name
```

### func \(\*NFTDrop\) [GetOwnedTokenIDs](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L105>)

```go
func (nft *NFTDrop) GetOwnedTokenIDs(ctx context.Context, address string) ([]*big.Int, error)
```

Get the tokenIds of all the NFTs owned by a specific address.

address: the address of the owner of the NFTs

returns: the tokenIds of all the NFTs owned by the address

### func \(\*NFTDrop\) [GetTotalClaimed](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L164>)

```go
func (drop *NFTDrop) GetTotalClaimed(ctx context.Context, address string) (*big.Int, error)
```

### func \(\*NFTDrop\) [TotalClaimedSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L155>)

```go
func (drop *NFTDrop) TotalClaimedSupply(ctx context.Context) (int, error)
```

Get the total number of NFTs that have been claimed.

### func \(\*NFTDrop\) [TotalUnclaimedSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_drop.go#L160>)

```go
func (drop *NFTDrop) TotalUnclaimedSupply(ctx context.Context) (int, error)
```

Get the total number of NFTs that have not yet been claimed.
