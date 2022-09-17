
## NFT Collection

You can access the NFT Collection interface from the SDK as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetNFTCollection("{{contract_address}}")
```

```go
type NFTCollection struct {
    *ERC721
    Signature *ERC721SignatureMinting
    Encoder   *ContractEncoder
}
```

### func \(\*NFTCollection\) [GetOwned](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L79>)

```go
func (nft *NFTCollection) GetOwned(ctx context.Context, address string) ([]*NFTMetadataOwner, error)
```

Get the metadatas of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the metadata of all the NFTs owned by the address

#### Example

```
owner := "{{wallet_address}}"
nfts, err := contract.GetOwned(context.Background(), owner)
name := nfts[0].Metadata.Name
```

### func \(\*NFTCollection\) [GetOwnedTokenIDs](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L96>)

```go
func (nft *NFTCollection) GetOwnedTokenIDs(ctx context.Context, address string) ([]*big.Int, error)
```

Get the tokenIds of all the NFTs owned by a specific address\.

address: the address of the owner of the NFTs

returns: the tokenIds of all the NFTs owned by the address

### func \(\*NFTCollection\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L121>)

```go
func (nft *NFTCollection) Mint(ctx context.Context, metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the connected wallet\.

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L173>)

```go
func (nft *NFTCollection) MintBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Mint a batch of new NFTs to the connected wallet\.

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L200>)

```go
func (nft *NFTCollection) MintBatchTo(ctx context.Context, address string, metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Mint a batch of new NFTs to the specified wallet\.

to: the wallet address to mint to

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

#### Example

```
metadatas := []*thirdweb.NFTMetadataInput{
	&thirdweb.NFTMetadataInput{
		Name: "Cool NFT",
		Description: "This is a cool NFT",
	}
	&thirdweb.NFTMetadataInput{
		Name: "Cool NFT 2",
		Description: "This is also a cool NFT",
	}
}

tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatas)
```

### func \(\*NFTCollection\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L146>)

```go
func (nft *NFTCollection) MintTo(ctx context.Context, address string, metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the specified wallet\.

address: the wallet address to mint to

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

#### Example

```
image, err := os.Open("path/to/image.jpg")
defer image.Close()

metadata := &thirdweb.NFTMetadataInput{
	Name: "Cool NFT",
	Description: "This is a cool NFT",
	Image: image,
}

tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", metadata)
```
