
## Tokenn
You can access this interface through the SDK with `sdk.GetToken(address)`.


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
