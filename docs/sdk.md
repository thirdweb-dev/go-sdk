
## ThirdwebSDK


```go
type SDKOptions struct {
    PrivateKey string
    GatewayUrl string
}
```

## type [ThirdwebSDK](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L11-L14>)

```go
type ThirdwebSDK struct {
    *ProviderHandler
    Storage IpfsStorage
}
```

### func [NewThirdwebSDK](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L23>)

```go
func NewThirdwebSDK(rpcUrlOrChainName string, options *SDKOptions) (*ThirdwebSDK, error)
```

#### NewThirdwebSDK

Create a new instance of the Thirdweb SDK

rpcUrlOrName: the name of the chain to connection to \(e\.g\. "rinkeby"\, "mumbai"\, "polygon"\, "mainnet"\, "fantom"\, "avalanche"\) or the RPC URL to connect to

options: an SDKOptions instance to specify a private key and/or an IPFS gateway URL

### func \(\*ThirdwebSDK\) [GetEdition](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L85>)

```go
func (sdk *ThirdwebSDK) GetEdition(address string) (*Edition, error)
```

#### GetEdition

Get an Edition contract SDK instance

address: the address of the Edition contract

### func \(\*ThirdwebSDK\) [GetEditionDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L126>)

```go
func (sdk *ThirdwebSDK) GetEditionDrop(address string) (*EditionDrop, error)
```

#### GetEditionDrop

Get an Edition Drop contract SDK instance

address: the address of the Edition Drop contract

### func \(\*ThirdwebSDK\) [GetNFTCollection](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L72>)

```go
func (sdk *ThirdwebSDK) GetNFTCollection(address string) (*NFTCollection, error)
```

#### GetNFTCollection

Get an NFT Collection contract SDK instance

address: the address of the NFT Collection contract

### func \(\*ThirdwebSDK\) [GetNFTDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L113>)

```go
func (sdk *ThirdwebSDK) GetNFTDrop(address string) (*NFTDrop, error)
```

#### GetNFTDrop

Get an NFT Drop contract SDK instance

address: the address of the NFT Drop contract

### func \(\*ThirdwebSDK\) [GetToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/sdk.go#L100>)

```go
func (sdk *ThirdwebSDK) GetToken(address string) (*Token, error)
```

#### GetToken

Returns a Token contract SDK instance

address: address of the token contract

Returns a Token contract SDK instance

## type [Token](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L11-L15>)

```go
type Token struct {
    *ERC20
    // contains filtered or unexported fields
}
```

### func \(\*Token\) [DelegateTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L157>)

```go
func (token *Token) DelegateTo(delegatreeAddress string) (*types.Transaction, error)
```

#### DelegateTo

Delegate the connected wallets tokens to a specified wallet

delegatreeAddress: wallet address to delegate tokens to

returns: transaction receipt of the delegation

### func \(\*Token\) [GetDelegation](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L66>)

```go
func (token *Token) GetDelegation() (string, error)
```

#### GetDelegation

Get the connected wallets delegatee address for this token\.

returns: delegation address of the connected wallet

### func \(\*Token\) [GetDelegationOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L75>)

```go
func (token *Token) GetDelegationOf(address string) (string, error)
```

#### GetDelegationOf

Get a specified wallets delegatee for this token

returns: delegation address of the connected wallet

### func \(\*Token\) [GetVoteBalance](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L41>)

```go
func (token *Token) GetVoteBalance() (*CurrencyValue, error)
```

#### GetVoteBalance

Get the connected wallets voting power in this token

returns: vote balance of the connected wallet

### func \(\*Token\) [GetVoteBalanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L52>)

```go
func (token *Token) GetVoteBalanceOf(address string) (*CurrencyValue, error)
```

#### GetVoteBalanceOf

Get the voting power of the specified wallet in this token

address: wallet address to check the vote balance of

returns: vote balance of the specified wallet

### func \(\*Token\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L91>)

```go
func (token *Token) Mint(amount float64) (*types.Transaction, error)
```

#### Mint

Mint tokens to the connected wallet

amount: amount of tokens to mint

returns: transaction receipt of the mint

### func \(\*Token\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L125>)

```go
func (token *Token) MintBatchTo(args []*TokenAmount) (*types.Transaction, error)
```

#### MintBatchTo

Mint tokens to a list of wallets

args: list of wallet addresses and amounts to mint

returns: transaction receipt of the mint

### func \(\*Token\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/token.go#L104>)

```go
func (token *Token) MintTo(to string, amount float64) (*types.Transaction, error)
```

#### MintTo

Mint tokens to a specified wallet

to: wallet address to mint tokens to

amount: amount of tokens to mint

returns: transaction receipt of the mint

## type [TokenAmount](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/types.go#L102-L105>)

```go
type TokenAmount struct {
    ToAddress string
    Amount    float64
}
```
