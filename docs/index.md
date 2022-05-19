<p align="center">
<br />
<a href="https://thirdweb.com"><img src="https://github.com/thirdweb-dev/typescript-sdk/blob/main/logo.svg?raw=true" width="200" alt=""/></a>
<br />
</p>
<h1 align="center">Thirdweb Go SDK</h1>
<p align="center">
<a href="https://discord.gg/thirdweb"><img alt="Join our Discord!" src="https://img.shields.io/discord/834227967404146718.svg?color=7289da&label=discord&logo=discord&style=flat"/></a>
</p>

# Installation

```bash
go get github.com/thirdweb-dev/go-sdk
```

## Getting Started

To start using this SDK, you just need to pass in a provider configuration.

### Instantiating the SDK

Once you have all the necessary dependencies, you can follow the following setup steps to get started with SDK read-only functions:

```go
package main

import (
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

func main() {
	// Creates a new read only SDK instance to read data from your contracts
	// you can pass either
	// - a chain name (rinkeby, mumbai, mainnet, polygon, etc)
	// - your own rpc URL
	sdk, _ := thirdweb.NewThirdwebSDK("polygon", nil)
}
```

### Working With Contracts

Once you instantiate the SDK, you can use it to access your thirdweb contracts. You can use SDK's contract getter functions like `GetNFTCollection`, `GetEdition`, and `GetNFTDrop`, to get the respective SDK contract instances. To use an NFT Collection contract for example, you can do the following.

```go
package main

import (
	"fmt"

	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

func main() {
	sdk, _ := thirdweb.NewThirdwebSDK("mumbai", nil)

	// Add your NFT Collection contract address here
	address := "0x..."
	nft, _ := sdk.GetNFTCollection(address)

	// Now you can use any of the read-only SDK contract functions
	nfts, _ := nft.GetAll()
	fmt.Println(len(nfts))
}
```

### Signing Transactions

In order to execute transactions on your contract, the SDK needs to know the wallet that is performing those transactions. For that it needs the private key of the wallet you want to execute transactions from.

> :warning: Never commit private keys to file tracking history, or your account could be compromised.

To connect your wallet to the SDK, you can use the following setup:

```go
package main

import (
	"github.com/thirdweb-dev/go-sdk/pkg/thirdweb"
)

func main() {
	// Instantiate the SDK with your privateKey (ideally from an env variable)
	sdk, _ := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		PrivateKey: privateKey,
	})

	address := "0x..."
	nft, _ := sdk.GetNFTCollection(address)

	// Now you can execute transactions using the SDK contract functions
	tx, _ = nft.Mint(
		&thirdweb.NFTMetadataInput{
			Name:        "Test NFT",
			Description: "Minted with the thirdweb Go SDK",
		}
	)
}
```
