
## ThirdwebSDK

```go
type ThirdwebSDK struct {
    *ProviderHandler
    Storage  IpfsStorage
    Deployer ContractDeployer
}
```

### func [NewThirdwebSDK](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L25>)

```go
func NewThirdwebSDK(rpcUrlOrChainName string, options *SDKOptions) (*ThirdwebSDK, error)
```

#### NewThirdwebSDK

Create a new instance of the Thirdweb SDK

rpcUrlOrName: the name of the chain to connection to \(e\.g\. "rinkeby", "mumbai", "polygon", "mainnet", "fantom", "avalanche"\) or the RPC URL to connect to

options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL

### func \(\*ThirdwebSDK\) [GetContract](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L166>)

```go
func (sdk *ThirdwebSDK) GetContract(address string) (*SmartContract, error)
```

#### GetContract

Get an instance of a custom contract deployed with thirdweb deploy

address: the address of the contract

### func \(\*ThirdwebSDK\) [GetContractFromAbi](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L182>)

```go
func (sdk *ThirdwebSDK) GetContractFromAbi(address string, abi string) (*SmartContract, error)
```

#### GetContractFromABI

Get an instance of ant custom contract from its ABI

address: the address of the contract

abi: the ABI of the contract

### func \(\*ThirdwebSDK\) [GetEdition](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L94>)

```go
func (sdk *ThirdwebSDK) GetEdition(address string) (*Edition, error)
```

#### GetEdition

Get an Edition contract SDK instance

address: the address of the Edition contract

### func \(\*ThirdwebSDK\) [GetEditionDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L138>)

```go
func (sdk *ThirdwebSDK) GetEditionDrop(address string) (*EditionDrop, error)
```

#### GetEditionDrop

Get an Edition Drop contract SDK instance

address: the address of the Edition Drop contract

### func \(\*ThirdwebSDK\) [GetMultiwrap](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L152>)

```go
func (sdk *ThirdwebSDK) GetMultiwrap(address string) (*Multiwrap, error)
```

#### GetMultiwrap

Get a Multiwrap contract SDK instance

address: the address of the Multiwrap contract

### func \(\*ThirdwebSDK\) [GetNFTCollection](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L80>)

```go
func (sdk *ThirdwebSDK) GetNFTCollection(address string) (*NFTCollection, error)
```

#### GetNFTCollection

Get an NFT Collection contract SDK instance

address: the address of the NFT Collection contract

### func \(\*ThirdwebSDK\) [GetNFTDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L124>)

```go
func (sdk *ThirdwebSDK) GetNFTDrop(address string) (*NFTDrop, error)
```

#### GetNFTDrop

Get an NFT Drop contract SDK instance

address: the address of the NFT Drop contract

### func \(\*ThirdwebSDK\) [GetToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/sdk.go#L110>)

```go
func (sdk *ThirdwebSDK) GetToken(address string) (*Token, error)
```

#### GetToken

Returns a Token contract SDK instance

address: address of the token contract

Returns a Token contract SDK instance
