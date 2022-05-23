
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
}
```

### func \(\*NFTCollection\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L55>)

```go
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error)
```

Mint a new NFT to the connected wallet\.

metadata: metadata of the NFT to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L103>)

```go
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error)
```

Mint a batch of new NFTs to the connected wallet\.

metadatas: list of metadata of the NFTs to mint

returns: the transaction receipt of the mint

### func \(\*NFTCollection\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L130>)

```go
func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) (*types.Transaction, error)
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

tx, err := contract.MintBatchTo("{{wallet_address}}", metadatas)
```

### func \(\*NFTCollection\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/nft_collection.go#L80>)

```go
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error)
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

tx, err := contract.MintTo("{{wallet_address}}", metadata)
```
