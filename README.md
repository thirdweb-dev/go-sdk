# nftlabs
--
    import "."

- [# nftlabs](#-nftlabs)
- [Usage](#usage)
		- [type CloudflareGateway](#type-cloudflaregateway)
		- [func  NewCloudflareGateway](#func--newcloudflaregateway)
		- [func (*CloudflareGateway) Get](#func-cloudflaregateway-get)
		- [type CommonModule](#type-commonmodule)
		- [type CreatePackArgs](#type-createpackargs)
		- [type Currency](#type-currency)
		- [type CurrencySdk](#type-currencysdk)
		- [type CurrencySdkModule](#type-currencysdkmodule)
		- [func  NewCurrencySdkModule](#func--newcurrencysdkmodule)
		- [func (*CurrencySdkModule) Balance](#func-currencysdkmodule-balance)
		- [func (*CurrencySdkModule) BalanceOf](#func-currencysdkmodule-balanceof)
		- [func (*CurrencySdkModule) Get](#func-currencysdkmodule-get)
		- [func (*CurrencySdkModule) GetValue](#func-currencysdkmodule-getvalue)
		- [func (*CurrencySdkModule) Transfer](#func-currencysdkmodule-transfer)
		- [type CurrencyValue](#type-currencyvalue)
		- [type Gateway](#type-gateway)
		- [type Listing](#type-listing)
		- [type MarketSdk](#type-marketsdk)
		- [type MarketSdkModule](#type-marketsdkmodule)
		- [func  NewMarketSdkModule](#func--newmarketsdkmodule)
		- [func (*MarketSdkModule) Buy](#func-marketsdkmodule-buy)
		- [func (*MarketSdkModule) GetAll](#func-marketsdkmodule-getall)
		- [func (*MarketSdkModule) GetListing](#func-marketsdkmodule-getlisting)
		- [func (*MarketSdkModule) List](#func-marketsdkmodule-list)
		- [func (*MarketSdkModule) SetPrivateKey](#func-marketsdkmodule-setprivatekey)
		- [func (*MarketSdkModule) Unlist](#func-marketsdkmodule-unlist)
		- [func (*MarketSdkModule) UnlistAll](#func-marketsdkmodule-unlistall)
		- [type NftMetadata](#type-nftmetadata)
		- [type NftSdk](#type-nftsdk)
		- [type NftSdkModule](#type-nftsdkmodule)
		- [func  NewNftSdkModule](#func--newnftsdkmodule)
		- [func (*NftSdkModule) Balance](#func-nftsdkmodule-balance)
		- [func (*NftSdkModule) BalanceOf](#func-nftsdkmodule-balanceof)
		- [func (*NftSdkModule) Get](#func-nftsdkmodule-get)
		- [func (*NftSdkModule) GetAll](#func-nftsdkmodule-getall)
		- [func (*NftSdkModule) GetAsync](#func-nftsdkmodule-getasync)
		- [func (*NftSdkModule) SetPrivateKey](#func-nftsdkmodule-setprivatekey)
		- [func (*NftSdkModule) Transfer](#func-nftsdkmodule-transfer)
		- [type NoAddressError](#type-noaddresserror)
		- [func (*NoAddressError) Error](#func-noaddresserror-error)
		- [type NoSignerError](#type-nosignererror)
		- [func (*NoSignerError) Error](#func-nosignererror-error)
		- [type NotFoundError](#type-notfounderror)
		- [func (*NotFoundError) Error](#func-notfounderror-error)
		- [type Pack](#type-pack)
		- [type PackNft](#type-packnft)
		- [type PackNftAddition](#type-packnftaddition)
		- [type PackSdk](#type-packsdk)
		- [type PackSdkModule](#type-packsdkmodule)
		- [func  NewPackSdkModule](#func--newpacksdkmodule)
		- [func (*PackSdkModule) Balance](#func-packsdkmodule-balance)
		- [func (*PackSdkModule) BalanceOf](#func-packsdkmodule-balanceof)
		- [func (*PackSdkModule) Create](#func-packsdkmodule-create)
		- [func (*PackSdkModule) DeployContract](#func-packsdkmodule-deploycontract)
		- [func (*PackSdkModule) Get](#func-packsdkmodule-get)
		- [func (*PackSdkModule) GetAll](#func-packsdkmodule-getall)
		- [func (*PackSdkModule) GetAsync](#func-packsdkmodule-getasync)
		- [func (*PackSdkModule) GetNfts](#func-packsdkmodule-getnfts)
		- [func (*PackSdkModule) Open](#func-packsdkmodule-open)
		- [func (*PackSdkModule) SetPrivateKey](#func-packsdkmodule-setprivatekey)
		- [func (*PackSdkModule) Transfer](#func-packsdkmodule-transfer)
		- [type SdkModule](#type-sdkmodule)
		- [type SdkOptions](#type-sdkoptions)
		- [type SigningMethod](#type-signingmethod)
		- [type UnmarshalError](#type-unmarshalerror)
		- [func (*UnmarshalError) Error](#func-unmarshalerror-error)
		- [type UnsupportedFunctionError](#type-unsupportedfunctionerror)
		- [func (*UnsupportedFunctionError) Error](#func-unsupportedfunctionerror-error)


## Usage

```go
var (
	InterfaceIdErc165           = [4]byte{1, 255, 201, 167}  // 0x01ffc9a7
	InterfaceIdErc721           = [4]byte{128, 172, 88, 205} // 0x80ac58cd
	InterfaceIdErc721Metadata   = [4]byte{91, 94, 19, 159}   // 0x5b5e139f
	InterfaceIdErc721Enumerable = [4]byte{120, 14, 157, 99}  // 0x780e9d63
	InterfaceIdErc1155          = [4]byte{217, 182, 122, 38} // 0xd9b67a26
)
```

#### type CloudflareGateway

```go
type CloudflareGateway struct {
	Url string
}
```


#### func  NewCloudflareGateway

```go
func NewCloudflareGateway(uri string) *CloudflareGateway
```

#### func (*CloudflareGateway) Get

```go
func (gw *CloudflareGateway) Get(uri string) ([]byte, error)
```

#### type CommonModule

```go
type CommonModule interface {
	SetPrivateKey(privateKey string) error
	// contains filtered or unexported methods
}
```


#### type CreatePackArgs

```go
type CreatePackArgs struct {
	AssetContractAddress  string
	Assets                []PackNftAddition
	SecondsUntilOpenStart *big.Int
	SecondsUntilOpenEnd   *big.Int
	RewardsPerOpen        *big.Int
}
```


#### type Currency

```go
type Currency struct {
	Name     string
	Symbol   string
	Decimals uint8
}
```


#### type CurrencySdk

```go
type CurrencySdk interface {
	Get() (Currency, error)
	GetValue(value *big.Int) (*CurrencyValue, error)
	Balance() (CurrencyValue, error)
	BalanceOf(address string, tokenId string) (CurrencyValue, error)
	Transfer(to string, amount *big.Int) error
}
```


#### type CurrencySdkModule

```go
type CurrencySdkModule struct {
	Client  *ethclient.Client
	Address string
}
```


#### func  NewCurrencySdkModule

```go
func NewCurrencySdkModule(client *ethclient.Client, asset string) (*CurrencySdkModule, error)
```

#### func (*CurrencySdkModule) Balance

```go
func (c *CurrencySdkModule) Balance() (CurrencyValue, error)
```

#### func (*CurrencySdkModule) BalanceOf

```go
func (c *CurrencySdkModule) BalanceOf(address string, tokenId string) (CurrencyValue, error)
```

#### func (*CurrencySdkModule) Get

```go
func (c *CurrencySdkModule) Get() (Currency, error)
```

#### func (*CurrencySdkModule) GetValue

```go
func (c *CurrencySdkModule) GetValue(value *big.Int) (*CurrencyValue, error)
```

#### func (*CurrencySdkModule) Transfer

```go
func (c *CurrencySdkModule) Transfer(to string, amount *big.Int) error
```

#### type CurrencyValue

```go
type CurrencyValue struct {
	Currency
	Value        string
	DisplayValue uint64
}
```


#### type Gateway

```go
type Gateway interface {
	Get(uri string) ([]byte, error)
}
```


#### type Listing

```go
type Listing struct {
	Id               *big.Int
	Seller           common.Address
	TokenContract    common.Address
	TokenId          *big.Int
	TokenMetadata    *NftMetadata
	Quantity         *big.Int
	CurrentContract  common.Address
	CurrencyMetadata *CurrencyValue // TODO: use currency type here
	Price            *big.Int
	SaleStart        *time.Time
	SaleEnd          *time.Time
}
```


#### type MarketSdk

```go
type MarketSdk interface {
	CommonModule
	GetListing(listingId *big.Int) (Listing, error)
	GetAll() ([]Listing, error)
	List(
		assetContract string,
		tokenId *big.Int,
		currencyContractAddress string,
		price *big.Int,
		quantity *big.Int,
		secondsUntilStart *big.Int,
		secondsUntilEnd *big.Int) (Listing, error)
	UnlistAll(listingId *big.Int) error
	Unlist(listingId *big.Int, quantity *big.Int) error
	Buy(listingId *big.Int, quantity *big.Int) error
}
```


#### type MarketSdkModule

```go
type MarketSdkModule struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
}
```


#### func  NewMarketSdkModule

```go
func NewMarketSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*MarketSdkModule, error)
```

#### func (*MarketSdkModule) Buy

```go
func (sdk *MarketSdkModule) Buy(listingId *big.Int, quantity *big.Int) error
```

#### func (*MarketSdkModule) GetAll

```go
func (sdk *MarketSdkModule) GetAll() ([]Listing, error)
```

#### func (*MarketSdkModule) GetListing

```go
func (sdk *MarketSdkModule) GetListing(listingId *big.Int) (Listing, error)
```

#### func (*MarketSdkModule) List

```go
func (sdk *MarketSdkModule) List(
	assetContractAddress string,
	tokenId *big.Int,
	currencyContractAddress string,
	pricePerToken *big.Int,
	quantity *big.Int,
	secondsUntilStart *big.Int,
	secondsUntilEnd *big.Int) (Listing, error)
```

#### func (*MarketSdkModule) SetPrivateKey

```go
func (sdk *MarketSdkModule) SetPrivateKey(privateKey string) error
```

#### func (*MarketSdkModule) Unlist

```go
func (sdk *MarketSdkModule) Unlist(listingId *big.Int, quantity *big.Int) error
```

#### func (*MarketSdkModule) UnlistAll

```go
func (sdk *MarketSdkModule) UnlistAll(listingId *big.Int) error
```

#### type NftMetadata

```go
type NftMetadata struct {
	Id          *big.Int
	Uri         string      `json:"external_url"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
}
```


#### type NftSdk

```go
type NftSdk interface {
	CommonModule
	Get(tokenId *big.Int) (NftMetadata, error)
	GetAll() ([]NftMetadata, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string) (*big.Int, error)
	Transfer(to string, tokenId *big.Int) error
}
```


#### type NftSdkModule

```go
type NftSdkModule struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
}
```


#### func  NewNftSdkModule

```go
func NewNftSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*NftSdkModule, error)
```

#### func (*NftSdkModule) Balance

```go
func (sdk *NftSdkModule) Balance(tokenId *big.Int) (*big.Int, error)
```

#### func (*NftSdkModule) BalanceOf

```go
func (sdk *NftSdkModule) BalanceOf(address string) (*big.Int, error)
```

#### func (*NftSdkModule) Get

```go
func (sdk *NftSdkModule) Get(tokenId *big.Int) (NftMetadata, error)
```

#### func (*NftSdkModule) GetAll

```go
func (sdk *NftSdkModule) GetAll() ([]NftMetadata, error)
```

#### func (*NftSdkModule) GetAsync

```go
func (sdk *NftSdkModule) GetAsync(tokenId *big.Int, ch chan<- NftMetadata, errCh chan<- error, wg *sync.WaitGroup)
```

#### func (*NftSdkModule) SetPrivateKey

```go
func (sdk *NftSdkModule) SetPrivateKey(privateKey string) error
```

#### func (*NftSdkModule) Transfer

```go
func (sdk *NftSdkModule) Transfer(to string, tokenId *big.Int) error
```

#### type NoAddressError

```go
type NoAddressError struct {
}
```


#### func (*NoAddressError) Error

```go
func (m *NoAddressError) Error() string
```

#### type NoSignerError

```go
type NoSignerError struct {
	Err error
}
```


#### func (*NoSignerError) Error

```go
func (m *NoSignerError) Error() string
```

#### type NotFoundError

```go
type NotFoundError struct {
}
```


#### func (*NotFoundError) Error

```go
func (m *NotFoundError) Error() string
```

#### type Pack

```go
type Pack struct {
	NftMetadata
	Creator       common.Address
	CurrentSupply big.Int
	OpenStart     time.Time
	OpenEnd       time.Time
}
```


#### type PackNft

```go
type PackNft struct {
	NftMetadata
	Supply *big.Int
}
```


#### type PackNftAddition

```go
type PackNftAddition struct {
	NftId  *big.Int
	Supply *big.Int
}
```


#### type PackSdk

```go
type PackSdk interface {
	CommonModule
	Open(packId *big.Int) (PackNft, error)
	Get(tokenId *big.Int) (Pack, error)
	GetAll() ([]Pack, error)
	GetNfts(packId *big.Int) ([]PackNft, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Transfer(to string, tokenId *big.Int, quantity *big.Int) error
	Create(args CreatePackArgs) (Pack, error)
}
```


#### type PackSdkModule

```go
type PackSdkModule struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
}
```


#### func  NewPackSdkModule

```go
func NewPackSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*PackSdkModule, error)
```

#### func (*PackSdkModule) Balance

```go
func (sdk *PackSdkModule) Balance(tokenId *big.Int) (*big.Int, error)
```

#### func (*PackSdkModule) BalanceOf

```go
func (sdk *PackSdkModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
```

#### func (*PackSdkModule) Create

```go
func (sdk *PackSdkModule) Create(args CreatePackArgs) (Pack, error)
```

#### func (*PackSdkModule) DeployContract

```go
func (sdk *PackSdkModule) DeployContract(name string) error
```

#### func (*PackSdkModule) Get

```go
func (sdk *PackSdkModule) Get(packId *big.Int) (Pack, error)
```

#### func (*PackSdkModule) GetAll

```go
func (sdk *PackSdkModule) GetAll() ([]Pack, error)
```

#### func (*PackSdkModule) GetAsync

```go
func (sdk *PackSdkModule) GetAsync(tokenId *big.Int, ch chan<- Pack, wg *sync.WaitGroup)
```

#### func (*PackSdkModule) GetNfts

```go
func (sdk *PackSdkModule) GetNfts(packId *big.Int) ([]PackNft, error)
```

#### func (*PackSdkModule) Open

```go
func (sdk *PackSdkModule) Open(packId *big.Int) (PackNft, error)
```

#### func (*PackSdkModule) SetPrivateKey

```go
func (sdk *PackSdkModule) SetPrivateKey(privateKey string) error
```

#### func (*PackSdkModule) Transfer

```go
func (sdk *PackSdkModule) Transfer(to string, tokenId *big.Int, quantity *big.Int) error
```

#### type SdkModule

```go
type SdkModule interface {
	GetSignerAddress() string
}
```


#### type SdkOptions

```go
type SdkOptions struct {
	IpfsGatewayUrl string
}
```


#### type SigningMethod

```go
type SigningMethod = func(common.Address, *types.Transaction) (*types.Transaction, error)
```


#### type UnmarshalError

```go
type UnmarshalError struct {
}
```


#### func (*UnmarshalError) Error

```go
func (m *UnmarshalError) Error() string
```

#### type UnsupportedFunctionError

```go
type UnsupportedFunctionError struct {
}
```


#### func (*UnsupportedFunctionError) Error

```go
func (m *UnsupportedFunctionError) Error() string
```
