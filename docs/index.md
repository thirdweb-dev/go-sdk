<p align="center">
<br />
<a href="https://thirdweb.com"><img src="https://github.com/thirdweb-dev/typescript-sdk/blob/main/logo.svg?raw=true" width="200" alt=""/></a>
<br />
</p>
<h1 align="center">thirdweb Go SDK</h1>
<p align="center">
<a href="https://discord.gg/thirdweb"><img alt="Join our Discord!" src="https://img.shields.io/discord/834227967404146718.svg?color=7289da&label=discord&logo=discord&style=flat"/></a>
</p>

# Installation

To install the SDK with the `go get` command, run the following:

```bash
go get github.com/thirdweb-dev/go-sdk/v2/thirdweb
```

## Getting Started

To start using this SDK, you just need to pass in a provider configuration. It's also strongly recommended that you use your thirdweb API keys with the SDK in order to get the best infrastructure performance (across RPCs, IPFS, etc.) - you can learn more about creating and using API keys [here](https://portal.thirdweb.com/api-keys).

### Instantiating the SDK

Once you have all the necessary dependencies, you can follow the following setup steps to get started with SDK read-only functions:

```go
package main

import (
	"fmt"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func main() {
	// Your secret key from the thirdweb api keys dashboard
	secretKey := "..."

	// Creates a new SDK instance to get read-only data for your contracts, you can pass:
	// - a chain name (mainnet, rinkeby, goerli, polygon, mumbai, avalanche, fantom)
	// - a custom RPC URL
	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		SecretKey: secretKey,
	})
	if err != nil {
		panic(err)
	}

	// Now we can interact with the SDK, like displaying the connected chain ID
	chainId, err := sdk.GetChainID()
	if err != nil {
		panic(err)
	}

	fmt.Println("New SDK instance create on chain", chainId)
}
```

### Working With Contracts

Once you instantiate the SDK, you can use it to access your thirdweb contracts. You can use SDK's contract getter functions like `GetNFTCollection`, `GetEdition`, and `GetNFTDrop`, to get the respective SDK contract instances. To use an NFT Collection contract for example, you can do the following.

```go
package main

import (
	"fmt"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func main() {
	// Get your secret key that you created from the thirdweb dashboard
	secretKey := "..."

	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		SecretKey: secretKey,
	})
	if err != nil {
		panic(err)
	}

	// Add your NFT Collection contract address here
	address := "0x..."
	nft, err := sdk.GetNFTCollection(address)
	if err != nil {
		panic(err)
	}

	// Now you can use any of the read-only SDK contract functions
	nfts, err := nft.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d NFTs found on this contract\n", len(nfts))
}
```

### Signing Transactions

In order to execute transactions on your contract, the SDK needs to know the wallet that is performing those transactions. For that it needs the private key of the wallet you want to execute transactions from.

> :warning: Never commit private keys to file tracking history, or your account could be compromised.

To connect your wallet to the SDK, you can use the following setup:

```go
package main

import (
	"fmt"
	"encoding/json"

	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

func main() {
	// Get your secret key that you created from the thirdweb dashboard
	secretKey := "..."

	// Get your private key securely (preferably from an environment variable)
	privateKey := "..."

	// Instantiate the SDK with your privateKey
	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		SecretKey: secretKey,
		PrivateKey: privateKey,
	})
	if err != nil {
		panic(err)
	}

	// Replace your contract address here
	address := "0x..."
	nft, err := sdk.GetNFTCollection(address)
	if err != nil {
		panic(err)
	}

	// Now you can execute transactions using the SDK contract functions
	tx, err := nft.Mint(
		&thirdweb.NFTMetadataInput{
			Name:        "Test NFT",
			Description: "Minted with the thirdweb Go SDK",
			Image: "ipfs://QmcCJC4T37rykDjR6oorM8hpB9GQWHKWbAi2YR1uTabUZu/0",
		},
	)
	if err != nil {
		panic(err)
	}

	result, _ := json.Marshal(&tx)
	fmt.Println(string(result))
}
```
