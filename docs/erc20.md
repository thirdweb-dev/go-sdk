
## ERC20

This interface is currently support by the Token contract\. You can access all of its functions through a Token contract instance\.

```go
type ERC20 struct {}
```

### func \(\*ERC20\) [Allowance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L104>)

```go
func (erc20 *ERC20) Allowance(spender string) (*CurrencyValue, error)
```

Get a specified spenders allowance for the connected wallets tokens\.

spender: wallet address to check the allowance of

returns: allowance of the spender for the connected wallets tokens

#### Example

```
spender := "0x..."

allowance, err := contract.Allowance(spender)
allowanceValue := allowance.DisplayValue
```

### func \(\*ERC20\) [AllowanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L123>)

```go
func (erc20 *ERC20) AllowanceOf(owner string, spender string) (*CurrencyValue, error)
```

Get a specified spenders allowance for the a specific wallets tokens\.

owner: wallet address who owns the assets

spender: wallet address to check the allowance of

returns: allowance of the spender for the connected wallets tokens

#### Example

```
address := "{{wallet_address}}"
spender := "0x..."

allowance, err := contract.AllowanceOf(address, spender)
allowanceValue := allowance.DisplayValue
```

### func \(\*ERC20\) [Balance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L56>)

```go
func (erc20 *ERC20) Balance() (*CurrencyValue, error)
```

Get the token balance of the connected wallet\.

returns: balance of the connected wallet

#### Example

```
balance, err := contract.Balance()
 balanceValue := balance.DisplayValue
```

### func \(\*ERC20\) [BalanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L71>)

```go
func (erc20 *ERC20) BalanceOf(address string) (*CurrencyValue, error)
```

Get the balance of the specified wallet\.

address: wallet address to check the balance of

returns: balance of the specified wallet

#### Example

```
address := "{{wallet_address}}"
	balance, err := contract.BalanceOf()
 balanceValue := balance.DisplayValue
```

### func \(\*ERC20\) [Burn](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L294>)

```go
func (erc20 *ERC20) Burn(ctx context.Context, amount float64) (*types.Transaction, error)
```

Burn a specified amount of tokens from the connected wallet\.

amount: amount of tokens to burn

returns: transaction receipt of the burn

#### Example

```
amount := 1
tx, err := contract.Burn(context.Background(), amount)
```

### func \(\*ERC20\) [BurnFrom](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L326>)

```go
func (erc20 *ERC20) BurnFrom(ctx context.Context, holder string, amount float64) (*types.Transaction, error)
```

Burn a specified amount of tokens from a specific wallet\.

holder: wallet address to burn the tokens from

amount: amount of tokens to burn

returns: transaction receipt of the burn

#### Example

```
holder := "0x..."
amount := 1

tx, err := contract.BurnFrom(context.Background(), holder, amount)
```

### func \(\*ERC20\) [Get](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L44>)

```go
func (erc20 *ERC20) Get() (*Currency, error)
```

Get token metadata including name, symbol, decimals, etc\.

returns: the metadata for the token

#### Example

```
currency, err := contract.Get()
symbol := currency.Symbol
```

### func \(\*ERC20\) [SetAllowance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L213>)

```go
func (erc20 *ERC20) SetAllowance(ctx context.Context, spender string, amount float64) (*types.Transaction, error)
```

Sets the allowance of a wallet to spend the connected wallets funds\.

spender: wallet address to set the allowance of

amount: amount of tokens to grant the spender allowance of

returns: transaction receipt of the allowance set

#### Example

```
spender := "0x..."
amount := 1

tx, err := contract.SetAllowance(context.Background(), spender, amount)
```

### func \(\*ERC20\) [TotalSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L83>)

```go
func (erc20 *ERC20) TotalSupply() (*CurrencyValue, error)
```

Get the total minted supply of the token\.

returns: total minted supply of the token

### func \(\*ERC20\) [Transfer](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L146>)

```go
func (erc20 *ERC20) Transfer(ctx context.Context, to string, amount float64) (*types.Transaction, error)
```

Transfer a specified amount of tokens from the connected wallet to a specified address\.

to: address to transfer the tokens to

amount: amount of tokens to transfer

returns: transaction receipt of the transfer

#### Example

```
to := "0x..."
amount := 1

tx, err := contract.Transfer(context.Background(), to, amount)
```

### func \(\*ERC20\) [TransferBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L251>)

```go
func (erc20 *ERC20) TransferBatch(ctx context.Context, args []*TokenAmount) (*types.Transaction, error)
```

Transfer tokens from the connected wallet to many wallets\.

args: list of token amounts with amounts and addresses to transfer to

returns: transaction receipt of the transfers

#### Example

```
args = []*thirdweb.TokenAmount{
	&thirdweb.TokenAmount{
		ToAddress: "0x...",
		Amount:    1
	}
	&thirdweb.TokenAmount{
		ToAddress: "0x...",
		Amount:    2
	}
}

tx, err := contract.TransferBatch(context.Background(), args)
```

### func \(\*ERC20\) [TransferFrom](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L181>)

```go
func (erc20 *ERC20) TransferFrom(ctx context.Context, from string, to string, amount float64) (*types.Transaction, error)
```

Transfer a specified amount of tokens from one specified address to another\.

from: address to transfer the tokens from

to: address to transfer the tokens to

amount: amount of tokens to transfer

returns: transaction receipt of the transfer

#### Example

```
from := "{{wallet_address}}"
to := "0x..."
amount := 1

tx, err := contract.TransferFrom(context.Background(), from, to, amount)
```
