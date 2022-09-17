
## Provider

```go
type ProviderHandler struct {}
```

### func [NewProviderHandler](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L22>)

```go
func NewProviderHandler(provider *ethclient.Client, privateKey string) (*ProviderHandler, error)
```

### func \(\*ProviderHandler\) [GetChainID](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L64>)

```go
func (handler *ProviderHandler) GetChainID() (*big.Int, error)
```

### func \(\*ProviderHandler\) [GetPrivateKey](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L60>)

```go
func (handler *ProviderHandler) GetPrivateKey() *ecdsa.PrivateKey
```

### func \(\*ProviderHandler\) [GetProvider](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L48>)

```go
func (handler *ProviderHandler) GetProvider() *ethclient.Client
```

### func \(\*ProviderHandler\) [GetRawPrivateKey](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L56>)

```go
func (handler *ProviderHandler) GetRawPrivateKey() string
```

### func \(\*ProviderHandler\) [GetSignerAddress](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L52>)

```go
func (handler *ProviderHandler) GetSignerAddress() common.Address
```

### func \(\*ProviderHandler\) [UpdatePrivateKey](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L40>)

```go
func (handler *ProviderHandler) UpdatePrivateKey(privateKey string) error
```

### func \(\*ProviderHandler\) [UpdateProvider](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/provider_handler.go#L36>)

```go
func (handler *ProviderHandler) UpdateProvider(provider *ethclient.Client)
```
