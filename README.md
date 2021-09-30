# NFTLabs Go SDK

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/github.com/nftlabs/nftlabs-sdk-go/pkg/nftlabs)

This repository allows you to interact with NFTLabs Ethereum protocols.

The Go SDK allows you to query and transact with any NFTLabs contract
deployed through the [NFTLabs Console](https://console.nftlabs.co)

## Project Structure

The SDK is up of modules:

- **NFT**: Mint and manage NFTs that you can list directly to a marketplace,
or transfer out to a user/managing wallet
- **Currency**: Create your own ERC20 coin that you can
use throughout the NFTLabs protocols equip with all the standard methods (mint, transfer, burn, etc)
- **Marketplace**: List NFTs, Collections, or even Packs of any set of digital assets
for sale with automated royalty collection
- **Collection**: A collection is a form of ERC1155 token that can contain a group
of digital assets. Collections are listable directly into a Marketplace (or
transfer out to another wallet)
- **Pack**: Packs are a special form of ERC1155 asset that can contain many
digital assets with a fixed opening window

You can use the NFTLabs SDK to interact with all these modules in a straightforward
and easy to use interface.

Interacting with the modules is easy:
```go
package examples

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/pkg/nftlabs"
	"log"
	"math/big"
)

func main() {
	nftContractAddress := ""
	chainRpcUrl := "https://rpc-mumbai.maticvigil.com" // change this

	client, err := ethclient.Dial(chainRpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	sdk, err := nftlabs.NewSdk(client, &nftlabs.SdkOptions{PrivateKey: "// TODO"})
	if err != nil {
		log.Fatal(err)
	}

	// You can get Pack/Marketplace/Collection/Currency contracts the same way
	nftModule, err := sdk.GetNftModule(nftContractAddress)
	if err != nil {
		log.Fatal(err)
	}
}
```


## Examples

Checkout the sdk usage examples [here](https://github.com/nftlabs/nftlabs-sdk-go/tree/master/examples) and our CLI example [here](https://github.com/nftlabs/nftlabs-sdk-go/tree/master/cmd/nftlabs)

## Getting the CLI

The easiest way to play with the SDK is through the CLI

You can pull the CLI like this:

```bash
$ go install github.com/nftlabs/nftlabs-sdk-go/cmd/nftlabs 
$ nftlabs
 CLI for the nftlabs-protocols go SDK

Usage:
  nftlabs [command]

Available Commands:
  collection  Interact with a collection contract
  currency    Interact with a currency contract
  help        Help about any command
  marketplace Interact with a marketplace contract
  nft         Interact with an nft contract
  pack        Interact with a pack contract

Flags:
  -u, --chainRpcUrl string   chain url where all rpc requests will be sent (default "https://rpc-mumbai.maticvigil.com")
  -h, --help                 help for nftlabs
  -k, --privateKey string    private key used to sign transactions

Use "nftlabs [command] --help" for more information about a command.
```

> Change your chain by setting the `-u` or `--chainRpcUrl` flag

Some examples (you can run right now):

### Getting all NFTs in a contract
```bash
$ SAMPLE_NFT_CONTRACT=0xf25C389016B7Ddb1191D4de02e495633Fe34d453
$ nftlabs nft -a $NFT_ADDRESS getAll 
```

### Minting a currency (uses gas)

Deploy a Currency through
the [NFTLabs console dashboard](https://console.nftlabs.co) in order to
be able to mint your tokens. When you have a
currency address, pass it to the cli with the `-a` (shorthand for `--address`)

> You need to set a signing key when executing this
> method in order to sign the transaction

Mint 42 of the currency:

```bash
// -k sets the signing key
$ nftlabs -k $PKEY currency -a $CURRENCY_ADDRESS mint 42
```
