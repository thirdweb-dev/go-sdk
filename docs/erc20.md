
## ERC20
This interface is supported by the `Token` contracts.


```go
type ERC20 struct {
    // contains filtered or unexported fields
}
```

### func \(\*ERC20\) [Allowance](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L88>)

```go
func (erc20 *ERC20) Allowance(spender string) (*CurrencyValue, error)
```

#### Allowance

Get a specified spenders allowance for the connected wallets tokens

spender: wallet address to check the allowance of

returns: allowance of the spender for the connected wallets tokens

### func \(\*ERC20\) [AllowanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L101>)

```go
func (erc20 *ERC20) AllowanceOf(owner string, spender string) (*CurrencyValue, error)
```

#### AllowanceOf

Get a specified spenders allowance for the a specific wallets tokens

owner: wallet address who owns the assets

spender: wallet address to check the allowance of

returns: allowance of the spender for the connected wallets tokens

### func \(\*ERC20\) [Balance](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L47>)

```go
func (erc20 *ERC20) Balance() (*CurrencyValue, error)
```

#### Balance

Get the token balance of the connected wallet

returns: balance of the connected wallet

### func \(\*ERC20\) [BalanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L58>)

```go
func (erc20 *ERC20) BalanceOf(address string) (*CurrencyValue, error)
```

#### BalanceOf

Get the balance of the specified wallet

address: wallet address to check the balance of

returns: balance of the specified wallet

### func \(\*ERC20\) [Burn](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L220>)

```go
func (erc20 *ERC20) Burn(amount float64) (*types.Transaction, error)
```

#### Burn

Burn a specified amount of tokens from the connected wallet

amount: amount of tokens to burn

returns: transaction receipt of the burn

### func \(\*ERC20\) [BurnFrom](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L243>)

```go
func (erc20 *ERC20) BurnFrom(holder string, amount float64) (*types.Transaction, error)
```

#### BurnFrom

Burn a specified amount of tokens from a specific wallet

holder: wallet address to burn the tokens from

amount: amount of tokens to burn

returns: transaction receipt of the burn

### func \(\*ERC20\) [Get](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L38>)

```go
func (erc20 *ERC20) Get() (*Currency, error)
```

#### Get

Get token metadata including name\, symbol\, decimals\, etc\.

returns: the metadata for the token

### func \(\*ERC20\) [SetAllowance](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L167>)

```go
func (erc20 *ERC20) SetAllowance(spender string, amount float64) (*types.Transaction, error)
```

#### SetAllowance

Sets the allowance of a wallet to spend the connected wallets funds

spender: wallet address to set the allowance of

amount: amount of tokens to grant the spender allowance of

returns: transaction receipt of the allowance set

### func \(\*ERC20\) [TotalSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L72>)

```go
func (erc20 *ERC20) TotalSupply() (*CurrencyValue, error)
```

#### TotalSupply

Get the total minted supply of the token

returns: total minted supply of the token

### func \(\*ERC20\) [Transfer](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L119>)

```go
func (erc20 *ERC20) Transfer(to string, amount float64) (*types.Transaction, error)
```

#### Transfer

Transfer a specified amount of tokens from the connected wallet to a specified address\.

to: address to transfer the tokens to

amount: amount of tokens to transfer

returns: transaction receipt of the transfer

### func \(\*ERC20\) [TransferBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L188>)

```go
func (erc20 *ERC20) TransferBatch(args []*TokenAmount) (*types.Transaction, error)
```

#### TransferBatch

Transfer tokens from the connected wallet to many wallets

args: list of token amounts with amounts and addresses to transfer to

returns: transaction receipt of the transfers

### func \(\*ERC20\) [TransferFrom](<https://github.com/thirdweb-dev/go-sdk/blob/main/pkg/thirdweb/erc20.go#L144>)

```go
func (erc20 *ERC20) TransferFrom(from string, to string, amount float64) (*types.Transaction, error)
```

#### TransferFrom

Transfer a specified amount of tokens from one specified address to another\.

from: address to transfer the tokens from

to: address to transfer the tokens to

amount: amount of tokens to transfer

returns: transaction receipt of the transfer
