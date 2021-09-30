# nftlabs
--
    import "."


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


#### func (*CloudflareGateway) Get

```go
func (gw *CloudflareGateway) Get(uri string) ([]byte, error)
```

#### func (*CloudflareGateway) Upload

```go
func (gw *CloudflareGateway) Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
```

#### type CollectionMetadata

```go
type CollectionMetadata struct {
	NftMetadata
	Creator string   `json:"creator"`
	Supply  *big.Int `json:"supply"`
}
```


#### type CommonModule

```go
type CommonModule interface {
	SetPrivateKey(privateKey string) error
	// contains filtered or unexported methods
}
```


#### type CreateCollectionArgs

```go
type CreateCollectionArgs struct {
	Supply   *big.Int    `json:"supply"`
	Metadata interface{} `json:"metadata"`
}
```


#### type CreatePackArgs

```go
type CreatePackArgs struct {
	AssetContractAddress  string
	Assets                []PackAssetAddition
	SecondsUntilOpenStart *big.Int
	SecondsUntilOpenEnd   *big.Int
	RewardsPerOpen        *big.Int
	Metadata              interface{}
}
```


#### type Currency

```go
type Currency interface {
	Get() (CurrencyMetadata, error)
	GetValue(value *big.Int) (*CurrencyValue, error)
	Balance() (CurrencyValue, error)
	BalanceOf(address string, tokenId string) (CurrencyValue, error)
	Transfer(to string, amount *big.Int) error
	Allowance(spender string) (*big.Int, error)
	SetAllowance(spender string, amount *big.Int) error
	Mint(amount *big.Int) error
	Burn(amount *big.Int) error
	BurnFrom(from string, amount *big.Int) error
	TransferFrom(from string, to string, amount *big.Int) error
	GrantRole(role Role, address string) error
	RevokeRole(role Role, address string) error
	TotalSupply() (*big.Int, error)
}
```


#### type CurrencyMetadata

```go
type CurrencyMetadata struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8  `json:"decimals"`
}
```


#### type CurrencyModule

```go
type CurrencyModule struct {
	Client  *ethclient.Client
	Address string
}
```


#### func (*CurrencyModule) Allowance

```go
func (sdk *CurrencyModule) Allowance(spender string) (*big.Int, error)
```

#### func (*CurrencyModule) Balance

```go
func (sdk *CurrencyModule) Balance() (CurrencyValue, error)
```

#### func (*CurrencyModule) BalanceOf

```go
func (sdk *CurrencyModule) BalanceOf(address string, tokenId string) (CurrencyValue, error)
```

#### func (*CurrencyModule) Burn

```go
func (sdk *CurrencyModule) Burn(amount *big.Int) error
```

#### func (*CurrencyModule) BurnFrom

```go
func (sdk *CurrencyModule) BurnFrom(from string, amount *big.Int) error
```

#### func (*CurrencyModule) Get

```go
func (sdk *CurrencyModule) Get() (CurrencyMetadata, error)
```

#### func (*CurrencyModule) GetValue

```go
func (sdk *CurrencyModule) GetValue(value *big.Int) (*CurrencyValue, error)
```

#### func (*CurrencyModule) GrantRole

```go
func (sdk *CurrencyModule) GrantRole(role Role, address string) error
```

#### func (*CurrencyModule) Mint

```go
func (sdk *CurrencyModule) Mint(amount *big.Int) error
```

#### func (*CurrencyModule) RevokeRole

```go
func (sdk *CurrencyModule) RevokeRole(role Role, address string) error
```

#### func (*CurrencyModule) SetAllowance

```go
func (sdk *CurrencyModule) SetAllowance(spender string, amount *big.Int) error
```

#### func (*CurrencyModule) TotalSupply

```go
func (sdk *CurrencyModule) TotalSupply() (*big.Int, error)
```

#### func (*CurrencyModule) Transfer

```go
func (sdk *CurrencyModule) Transfer(to string, amount *big.Int) error
```

#### func (*CurrencyModule) TransferFrom

```go
func (sdk *CurrencyModule) TransferFrom(from string, to string, amount *big.Int) error
```

#### type CurrencyValue

```go
type CurrencyValue struct {
	CurrencyMetadata
	Value        *big.Int `json:"value"`
	DisplayValue string   `json:"displayValue"`
}
```


#### type Gateway

```go
type Gateway interface {
	Get(uri string) ([]byte, error)
	Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
}
```


#### type ISdk

```go
type ISdk interface {
	GetNftModule(address string) (Nft, error)
	GetMarketModule(address string) (Market, error)
	GetCurrencyModule(address string) (Currency, error)
	GetPackModule(address string) (Pack, error)
	GetNftCollectionModule(address string) (NftCollection, error)
	// contains filtered or unexported methods
}
```


#### type Listing

```go
type Listing struct {
	Id               *big.Int       `json:"id"`
	Seller           common.Address `json:"seller"`
	TokenContract    common.Address `json:"tokenContract"`
	TokenId          *big.Int       `json:"tokenId"`
	TokenMetadata    *NftMetadata   `json:"tokenMetadata"`
	Quantity         *big.Int       `json:"quantity"`
	CurrentContract  common.Address `json:"currentContract"`
	CurrencyMetadata *CurrencyValue `json:"currencyMetadata"`
	Price            *big.Int       `json:"price"`
	SaleStart        *time.Time     `json:"saleStart"`
	SaleEnd          *time.Time     `json:"saleEnd"`
}
```


#### type ListingFilter

```go
type ListingFilter struct {
	Seller        string   `json:"seller"`
	TokenContract string   `json:"tokenContract"`
	TokenId       *big.Int `json:"tokenId"`
}
```


#### type Market

```go
type Market interface {
	GetListing(listingId *big.Int) (Listing, error)
	GetAll(filter ListingFilter) ([]Listing, error)
	List(args NewListingArgs) (Listing, error)
	UnlistAll(listingId *big.Int) error
	Unlist(listingId *big.Int, quantity *big.Int) error
	Buy(listingId *big.Int, quantity *big.Int) error
	GetMarketFeeBps() (*big.Int, error)
	SetMarketFeeBps(fee *big.Int) error
}
```


#### type MarketModule

```go
type MarketModule struct {
	Client  *ethclient.Client
	Address string
}
```


#### func (*MarketModule) Buy

```go
func (sdk *MarketModule) Buy(listingId *big.Int, quantity *big.Int) error
```

#### func (*MarketModule) GetAll

```go
func (sdk *MarketModule) GetAll(filter ListingFilter) ([]Listing, error)
```

#### func (*MarketModule) GetListing

```go
func (sdk *MarketModule) GetListing(listingId *big.Int) (Listing, error)
```

#### func (*MarketModule) GetMarketFeeBps

```go
func (sdk *MarketModule) GetMarketFeeBps() (*big.Int, error)
```

#### func (*MarketModule) List

```go
func (sdk *MarketModule) List(args NewListingArgs) (Listing, error)
```
TODO: change args to struct

#### func (*MarketModule) SetMarketFeeBps

```go
func (sdk *MarketModule) SetMarketFeeBps(fee *big.Int) error
```

#### func (*MarketModule) Unlist

```go
func (sdk *MarketModule) Unlist(listingId *big.Int, quantity *big.Int) error
```

#### func (*MarketModule) UnlistAll

```go
func (sdk *MarketModule) UnlistAll(listingId *big.Int) error
```

#### type MintCollectionArgs

```go
type MintCollectionArgs struct {
	TokenId *big.Int `json:"tokenId"`
	Amount  *big.Int `json:"amount"`
}
```


#### type MintNftMetadata

```go
type MintNftMetadata struct {
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
	Image                string   `json:"image"`
	ExternalUrl          string   `json:"external_url"`
	SellerFeeBasisPoints *big.Int `json:"seller_fee_basis_points"`
	FeeRecipient         string   `json:"fee_recipient"`
	BackgroundColor      string   `json:"background_color"`
}
```


#### type NewListingArgs

```go
type NewListingArgs struct {
	AssetContractAddress    string   `json:"assetContractAddress"`
	TokenId                 *big.Int `json:"tokenId"`
	CurrencyContractAddress string   `json:"currencyContractAddress"`
	Price                   *big.Int `json:"price"`
	Quantity                *big.Int `json:"quantity"`
	SecondsUntilOpenStart   *big.Int `json:"secondsUntilOpenStart"`
	SecondsUntilOpenEnd     *big.Int `json:"secondsUntilOpenEnd"`
	RewardsPerOpen          *big.Int `json:"rewardsPerOpen"`
}
```


#### type Nft

```go
type Nft interface {
	Get(tokenId *big.Int) (NftMetadata, error)
	GetAll() ([]NftMetadata, error)
	GetOwned(address string) ([]NftMetadata, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string) (*big.Int, error)
	Transfer(to string, tokenId *big.Int) error
	TotalSupply() (*big.Int, error)
	SetApproval(operator string, approved bool) error
	Mint(metadata MintNftMetadata) (NftMetadata, error)
	MintBatch(meta []interface{}) ([]NftMetadata, error)
	Burn(tokenId *big.Int) error
	TransferFrom(from string, to string, tokenId *big.Int) error
	SetRoyaltyBps(amount *big.Int) error
	GrantRole(role Role, address string) error
	RevokeRole(role Role, address string) error
	// contains filtered or unexported methods
}
```


#### type NftCollection

```go
type NftCollection interface {
	Get(tokenId *big.Int) (CollectionMetadata, error)
	GetAll() ([]CollectionMetadata, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	IsApproved(address string, operator string) (bool, error)
	SetApproved(operator string, approved bool) error
	Transfer(to string, tokenId *big.Int, amount *big.Int) error
	Create(args []CreateCollectionArgs) ([]CollectionMetadata, error)
	Mint(args MintCollectionArgs) error
}
```


#### type NftCollectionModule

```go
type NftCollectionModule struct {
	Client  *ethclient.Client
	Address string
}
```


#### func (*NftCollectionModule) Balance

```go
func (sdk *NftCollectionModule) Balance(tokenId *big.Int) (*big.Int, error)
```

#### func (*NftCollectionModule) BalanceOf

```go
func (sdk *NftCollectionModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
```

#### func (*NftCollectionModule) Create

```go
func (sdk *NftCollectionModule) Create(args []CreateCollectionArgs) ([]CollectionMetadata, error)
```

#### func (*NftCollectionModule) Get

```go
func (sdk *NftCollectionModule) Get(tokenId *big.Int) (CollectionMetadata, error)
```

#### func (*NftCollectionModule) GetAll

```go
func (sdk *NftCollectionModule) GetAll() ([]CollectionMetadata, error)
```

#### func (*NftCollectionModule) GetAsync

```go
func (sdk *NftCollectionModule) GetAsync(tokenId *big.Int, ch chan<- CollectionMetadata, wg *sync.WaitGroup)
```

#### func (*NftCollectionModule) IsApproved

```go
func (sdk *NftCollectionModule) IsApproved(address string, operator string) (bool, error)
```

#### func (*NftCollectionModule) Mint

```go
func (sdk *NftCollectionModule) Mint(args MintCollectionArgs) error
```

#### func (*NftCollectionModule) SetApproved

```go
func (sdk *NftCollectionModule) SetApproved(operator string, approved bool) error
```

#### func (*NftCollectionModule) Transfer

```go
func (sdk *NftCollectionModule) Transfer(to string, tokenId *big.Int, amount *big.Int) error
```

#### type NftMetadata

```go
type NftMetadata struct {
	Id          *big.Int    `json:"id"`
	Uri         string      `json:"external_url"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Image       string      `json:"image"`
	Attributes  interface{} `json:"attributes"`
}
```


#### type NftModule

```go
type NftModule struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
}
```


#### func (*NftModule) Balance

```go
func (sdk *NftModule) Balance(tokenId *big.Int) (*big.Int, error)
```

#### func (*NftModule) BalanceOf

```go
func (sdk *NftModule) BalanceOf(address string) (*big.Int, error)
```

#### func (*NftModule) Burn

```go
func (sdk *NftModule) Burn(tokenId *big.Int) error
```

#### func (*NftModule) Get

```go
func (sdk *NftModule) Get(tokenId *big.Int) (NftMetadata, error)
```

#### func (*NftModule) GetAll

```go
func (sdk *NftModule) GetAll() ([]NftMetadata, error)
```

#### func (*NftModule) GetAsync

```go
func (sdk *NftModule) GetAsync(tokenId *big.Int, ch chan<- NftMetadata, errCh chan<- error, wg *sync.WaitGroup)
```

#### func (*NftModule) GetOwned

```go
func (sdk *NftModule) GetOwned(address string) ([]NftMetadata, error)
```

#### func (*NftModule) GrantRole

```go
func (sdk *NftModule) GrantRole(role Role, address string) error
```

#### func (*NftModule) Mint

```go
func (sdk *NftModule) Mint(metadata MintNftMetadata) (NftMetadata, error)
```

#### func (*NftModule) MintBatch

```go
func (sdk *NftModule) MintBatch(meta []interface{}) ([]NftMetadata, error)
```

#### func (*NftModule) RevokeRole

```go
func (sdk *NftModule) RevokeRole(role Role, address string) error
```

#### func (*NftModule) SetApproval

```go
func (sdk *NftModule) SetApproval(operator string, approved bool) error
```

#### func (*NftModule) SetRoyaltyBps

```go
func (sdk *NftModule) SetRoyaltyBps(amount *big.Int) error
```

#### func (*NftModule) TotalSupply

```go
func (sdk *NftModule) TotalSupply() (*big.Int, error)
```

#### func (*NftModule) Transfer

```go
func (sdk *NftModule) Transfer(to string, tokenId *big.Int) error
```

#### func (*NftModule) TransferFrom

```go
func (sdk *NftModule) TransferFrom(from string, to string, tokenId *big.Int) error
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
type Pack interface {
	Open(packId *big.Int) (PackNft, error)
	Get(tokenId *big.Int) (PackMetadata, error)
	GetAll() ([]PackMetadata, error)
	GetNfts(packId *big.Int) ([]PackNft, error)
	Balance(tokenId *big.Int) (*big.Int, error)
	BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
	Transfer(to string, tokenId *big.Int, quantity *big.Int) error
	Create(args CreatePackArgs) (PackMetadata, error)
}
```


#### type PackAssetAddition

```go
type PackAssetAddition struct {
	NftId  *big.Int
	Supply *big.Int
}
```


#### type PackMetadata

```go
type PackMetadata struct {
	NftMetadata
	Creator       common.Address
	CurrentSupply big.Int
	OpenStart     time.Time
	OpenEnd       time.Time
}
```


#### type PackModule

```go
type PackModule struct {
	Client  *ethclient.Client
	Address string
}
```


#### func (*PackModule) Balance

```go
func (sdk *PackModule) Balance(tokenId *big.Int) (*big.Int, error)
```

#### func (*PackModule) BalanceOf

```go
func (sdk *PackModule) BalanceOf(address string, tokenId *big.Int) (*big.Int, error)
```

#### func (*PackModule) Create

```go
func (sdk *PackModule) Create(args CreatePackArgs) (PackMetadata, error)
```

#### func (*PackModule) Get

```go
func (sdk *PackModule) Get(packId *big.Int) (PackMetadata, error)
```

#### func (*PackModule) GetAll

```go
func (sdk *PackModule) GetAll() ([]PackMetadata, error)
```

#### func (*PackModule) GetAsync

```go
func (sdk *PackModule) GetAsync(tokenId *big.Int, ch chan<- PackMetadata, wg *sync.WaitGroup)
```

#### func (*PackModule) GetNfts

```go
func (sdk *PackModule) GetNfts(packId *big.Int) ([]PackNft, error)
```

#### func (*PackModule) Open

```go
func (sdk *PackModule) Open(packId *big.Int) (PackNft, error)
```

#### func (*PackModule) Transfer

```go
func (sdk *PackModule) Transfer(to string, tokenId *big.Int, quantity *big.Int) error
```

#### type PackNft

```go
type PackNft struct {
	NftMetadata
	Supply *big.Int
}
```


#### type Role

```go
type Role string
```


```go
const (
	AdminRole  Role = "admin"
	MinterRole      = "minter"
)
```

#### type Sdk

```go
type Sdk struct {
}
```


#### func  NewSdk

```go
func NewSdk(client *ethclient.Client, opt *SdkOptions) (*Sdk, error)
```

#### func (*Sdk) GetCurrencyModule

```go
func (sdk *Sdk) GetCurrencyModule(address string) (Currency, error)
```

#### func (*Sdk) GetGateway

```go
func (sdk *Sdk) GetGateway(address string) Gateway
```

#### func (*Sdk) GetMarketModule

```go
func (sdk *Sdk) GetMarketModule(address string) (Market, error)
```

#### func (*Sdk) GetNftCollectionModule

```go
func (sdk *Sdk) GetNftCollectionModule(address string) (NftCollection, error)
```

#### func (*Sdk) GetNftModule

```go
func (sdk *Sdk) GetNftModule(address string) (Nft, error)
```

#### func (*Sdk) GetPackModule

```go
func (sdk *Sdk) GetPackModule(address string) (Pack, error)
```

#### type SdkOptions

```go
type SdkOptions struct {
	IpfsGatewayUrl string
	PrivateKey     string
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
