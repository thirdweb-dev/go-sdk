
## Marketplace

You can access the Marketplace interface from the SDK as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

contract, err := sdk.GetMarketplace("{{contract_address}}")
```

```go
type Marketplace struct {
    Encoder *MarketplaceEncoder
}
```

### func \(\*Marketplace\) [BuyoutListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L179>)

```go
func (marketplace *Marketplace) BuyoutListing(listingId int, quantityDesired int) (*types.Transaction, error)
```

Buy a specific listing from the marketplace\.

listingId: listing ID of the asset you want to buy

quantityDesired: the quantity of the asset to buy from the listing

returns: transaction receipt of the purchase

### func \(\*Marketplace\) [BuyoutListingTo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L199>)

```go
func (marketplace *Marketplace) BuyoutListingTo(listingId int, quantityDesired int, receiver string) (*types.Transaction, error)
```

Buy a specific listing from the marketplace to a specific address\.

listingId: listing ID of the asset you want to buy

quantityDesired: the quantity of the asset to buy from the listing

receiver: specific address to receive the assets from the listing

returns: transaction receipt of the purchase

#### Example

```
listingId := 0
quantityDesired := 1
receiver := "0x..."
receipt, err := marketplace.BuyoutListingTo(listingId, quantityDesired, receiver)
```

### func \(\*Marketplace\) [CancelListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L158>)

```go
func (marketplace *Marketplace) CancelListing(listingId int) (*types.Transaction, error)
```

Cancel a listing on the marketplace\.

listingId: listing ID to cancel

returns: transaction receipt of the cancellation

#### Example

```
listingId := 0
receipt, err := marketplace.CancelListing(listingId)
```

### func \(\*Marketplace\) [CreateListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L266>)

```go
func (marketplace *Marketplace) CreateListing(listing *NewDirectListing) (int, error)
```

Create a new listing on the marketplace where people can buy an asset directly\.

listing: the data for the listing to create

returns: the ID of the listing that was created

#### Example

```
listing := &NewDirectListing{
	AssetContractAddress: "0x...", // Address of the asset contract
	TokenId: 0, // Token ID of the asset to list
	StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
	ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
	Quantity: 1, // Quantity of the asset to list
	CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
	BuyoutPricePerToken: 1, // Price per token of the asset to list
}

listingId, err := marketplace.CreateListing(listing)
```

### func \(\*Marketplace\) [GetActiveListings](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L96>)

```go
func (marketplace *Marketplace) GetActiveListings(filter *MarketplaceFilter) ([]*DirectListing, error)
```

Get all active listings from the marketplace\.

filter: optional filter parameters

returns: all active listings in the marketplace

#### Example

```
listings, err := marketplace.GetActiveListings()
// Price per token of the first listing
listings[0].BuyoutCurrencyValuePerToken.DisplayValue
```

### func \(\*Marketplace\) [GetAllListings](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L127>)

```go
func (marketplace *Marketplace) GetAllListings(filter *MarketplaceFilter) ([]*DirectListing, error)
```

Get all the listings from the marketplace\.

filter: optional filter parameters

returns: all listings in the marketplace

Example listings, err := marketplace\.GetAllListings\(\) // Price per token of the first listing listings\[0\]\.BuyoutCurrencyValuePerToken\.DisplayValue

### func \(\*Marketplace\) [GetListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L68>)

```go
func (marketplace *Marketplace) GetListing(listingId int) (*DirectListing, error)
```

Get a single listing from the marketplace\.

listingId: listing ID to get

returns: listing at the given listing ID

#### Example

```
listingId := 0
listing, err := marketplace.GetListing(listingId)
```

### func \(\*Marketplace\) [GetTotalCount](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace.go#L139>)

```go
func (marketplace *Marketplace) GetTotalCount() (int, error)
```

Get the total number of listings in the marketplace\.

returns: total number of listings in the marketplace
