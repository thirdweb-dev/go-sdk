<p align="center">
<br />
<a href="https://thirdweb.com"><img src="https://github.com/thirdweb-dev/typescript-sdk/blob/main/logo.svg?raw=true" width="200" alt=""/></a>
<br />
</p>
<h1 align="center">Thirdweb Go SDK</h1>
<p align="center">
<a href="https://discord.gg/thirdweb"><img alt="Join our Discord!" src="https://img.shields.io/discord/834227967404146718.svg?color=7289da&label=discord&logo=discord&style=flat"/></a>

# Installation

```bash
go get github.com/thirdweb-dev/go-sdk
```

## Getting Started

To start using this SDK, you just need to pass in a provider configuration.

### Instantiating the SDK

Once you have all the necessary dependencies, you can follow the following setup steps to get started with SDK read-only functions:

```go
import (
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

func main() {
	// Add your own RPC URL here or use a public one
	rpcUrl := "https://rpc-mumbai.maticvigil.com"

	// Now you can create a new instance of the SDK to use
	sdk, _ := thirdweb.NewThirdwebSDK(rpcUrl, nil)
}

main()
```

### Working With Contracts

Once you instantiate the SDK, you can use it to access your thirdweb contracts. You can use SDK's contract getter functions like `GetNFTCollection`, `GetEdition`, and `GetNFTDrop`, to get the respective SDK contract instances. To use an NFT Collection contract for example, you can do the following.

```go
import (
	"fmt"

	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

func main() {
	rpcUrl := "https://rpc-mumbai.maticvigil.com"
	sdk, _ := thirdweb.NewThirdwebSDK(rpcUrl, nil)

	// Add your NFT Collection contract address here
	address := "0x..."
	nft, _ := sdk.GetNFTCollection(address)

	// Now you can use any of the read-only SDK contract functions
	nfts, _ = nft.GetAll()
	fmt.Println(len(nfts))
}

main()
```

### Signing Transactions

> :warning: Never commit private keys to file tracking history, or your account could be compromised.

Meanwhile, if you want to use write functions as well and connect a signer, you can use the following setup:

```go
import (
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

func main() {
	rpcUrl := "https://rpc-mumbai.maticvigil.com"

	// Add your private key (make sure it isnt directly written in this file)
	privateKey := "..."

	// Instantiate the SDK with your privateKey
	sdk, _ := thirdweb.NewThirdwebSDK(rpcUrl, &thirdweb.SDKOptions{
		PrivateKey: privateKey,
	})
	
	address := "0x..."
	nft, _ := sdk.GetNFTCollection(address)

	// Now you can use any of the read-only SDK contract functions
	tx, _ = nft.Mint(
		&thirdweb.NFTMetadataInput{
			Name:        "Test NFT",
			Description: "Minted with the thirdweb Go SDK",
		}
	)
}

main()
```


