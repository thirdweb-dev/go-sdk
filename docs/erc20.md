
## ERC20

This interface is currently support by the Token contract. You can access all of its functions through a Token contract instance.

```go
type ERC20 struct {}
```

### func \(\*ERC20\) [Allowance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L113>)

```go
func (erc20 *ERC20) Allowance(ctx context.Context, spender string) (*CurrencyValue, error)
```

#### Get token allowance for a specific spender

spender: wallet address to check the allowance of

returns: allowance of the spender for the connected wallets tokens

#### Example

```
spender := "0x..."

allowance, err := contract.ERC20.Allowance(spender)
allowanceValue := allowance.DisplayValue
```

### func \(\*ERC20\) [AllowanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L132>)

```go
func (erc20 *ERC20) AllowanceOf(ctx context.Context, owner string, spender string) (*CurrencyValue, error)
```

#### Get token allowance for a specific spender and owner

owner: wallet address who owns the assets

spender: wallet address to check the allowance of

returns: allowance of the spender for the connected wallets tokens

#### Example

```
address := "{{wallet_address}}"
spender := "0x..."

allowance, err := contract.ERC20.AllowanceOf(address, spender)
allowanceValue := allowance.DisplayValue
```

### func \(\*ERC20\) [Balance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L61>)

```go
func (erc20 *ERC20) Balance(ctx context.Context) (*CurrencyValue, error)
```

#### Get token balance

@extension: ERC20

returns: balance of the connected wallet

#### Example

```
balance, err := contract.ERC20.Balance()
balanceValue := balance.DisplayValue
```

### func \(\*ERC20\) [BalanceOf](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L76>)

```go
func (erc20 *ERC20) BalanceOf(ctx context.Context, address string) (*CurrencyValue, error)
```

#### Get token balance of a specific wallet

address: wallet address to check the balance of

returns: balance of the specified wallet

#### Example

```
address := "{{wallet_address}}"
balance, err := contract.ERC20.BalanceOf()
balanceValue := balance.DisplayValue
```

### func \(\*ERC20\) [Burn](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L303>)

```go
func (erc20 *ERC20) Burn(ctx context.Context, amount float64) (*types.Transaction, error)
```

#### Burn tokens

amount: amount of tokens to burn

returns: transaction receipt of the burn

#### Example

```
amount := 1
tx, err := contract.ERC20.Burn(context.Background(), amount)
```

### func \(\*ERC20\) [BurnFrom](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L335>)

```go
func (erc20 *ERC20) BurnFrom(ctx context.Context, holder string, amount float64) (*types.Transaction, error)
```

#### Burn tokens from a specific wallet

holder: wallet address to burn the tokens from

amount: amount of tokens to burn

returns: transaction receipt of the burn

#### Example

```
holder := "0x..."
amount := 1

tx, err := contract.ERC20.BurnFrom(context.Background(), holder, amount)
```

### func \(\*ERC20\) [Get](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L47>)

```go
func (erc20 *ERC20) Get(ctx context.Context) (*Currency, error)
```

#### Get token metadata

@extension: ERC20

returns: the metadata for the token

#### Example

```
currency, err := contract.ERC20.Get()
symbol := currency.Symbol
```

### func \(\*ERC20\) [Mint](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L363>)

```go
func (erc20 *ERC20) Mint(ctx context.Context, amount float64) (*types.Transaction, error)
```

#### Mint tokens

amount: amount of tokens to mint

returns: transaction receipt of the mint

#### Example

```
tx, err := contract.ERC20.Mint(context.Background(), 1)
```

### func \(\*ERC20\) [MintBatchTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L416>)

```go
func (erc20 *ERC20) MintBatchTo(ctx context.Context, args []*TokenAmount) (*types.Transaction, error)
```

#### Mint tokens to many wallets

args: list of wallet addresses and amounts to mint

returns: transaction receipt of the mint

#### Example

```
args = []*thirdweb.TokenAmount{
	&thirdweb.TokenAmount{
		ToAddress: "{{wallet_address}}",
		Amount:    1
	}
	&thirdweb.TokenAmount{
		ToAddress: "{{wallet_address}}",
		Amount:    2
	}
}

tx, err := contract.ERC20.MintBatchTo(context.Background(), args)
```

### func \(\*ERC20\) [MintTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L378>)

```go
func (erc20 *ERC20) MintTo(ctx context.Context, to string, amount float64) (*types.Transaction, error)
```

#### Mint tokens to a specific wallet

to: wallet address to mint tokens to

amount: amount of tokens to mint

returns: transaction receipt of the mint

#### Example

```
tx, err := contract.ERC20.MintTo(context.Background(), "{{wallet_address}}", 1)
```

### func \(\*ERC20\) [SetAllowance](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L222>)

```go
func (erc20 *ERC20) SetAllowance(ctx context.Context, spender string, amount float64) (*types.Transaction, error)
```

#### Set token allowance

spender: wallet address to set the allowance of

amount: amount of tokens to grant the spender allowance of

returns: transaction receipt of the allowance set

#### Example

```
spender := "0x..."
amount := 1

tx, err := contract.ERC20.SetAllowance(context.Background(), spender, amount)
```

### func \(\*ERC20\) [TotalSupply](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L92>)

```go
func (erc20 *ERC20) TotalSupply(ctx context.Context) (*CurrencyValue, error)
```

#### Get the total minted supply

returns: total minted supply of the token

#### Example

```
supply, err := contract.ERC20.TotalSupply(context.Background())
```

### func \(\*ERC20\) [Transfer](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L155>)

```go
func (erc20 *ERC20) Transfer(ctx context.Context, to string, amount float64) (*types.Transaction, error)
```

#### Transfer tokens

to: address to transfer the tokens to

amount: amount of tokens to transfer

returns: transaction receipt of the transfer

#### Example

```
to := "0x..."
amount := 1

tx, err := contract.ERC20.Transfer(context.Background(), to, amount)
```

### func \(\*ERC20\) [TransferBatch](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L260>)

```go
func (erc20 *ERC20) TransferBatch(ctx context.Context, args []*TokenAmount) (*types.Transaction, error)
```

#### Transfer many tokens

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

tx, err := contract.ERC20.TransferBatch(context.Background(), args)
```

### func \(\*ERC20\) [TransferFrom](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/erc20.go#L190>)

```go
func (erc20 *ERC20) TransferFrom(ctx context.Context, from string, to string, amount float64) (*types.Transaction, error)
```

#### Transfer tokens from a specific wallet

from: address to transfer the tokens from

to: address to transfer the tokens to

amount: amount of tokens to transfer

returns: transaction receipt of the transfer

#### Example

```
from := "{{wallet_address}}"
to := "0x..."
amount := 1

tx, err := contract.ERC20.TransferFrom(context.Background(), from, to, amount)
```
