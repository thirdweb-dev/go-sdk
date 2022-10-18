
## Marketplace Encoder

The marketplace encoder class is used to get the unsigned transaction data for marketplace contract contract calls that can be signed at a later time after generation\.

It can be accessed from the SDK through the \`Encoder\` namespace of the marketplace contract:

You can access the Marketplace interface from the SDK as follows:

```
import (
		"github.com/thirdweb-dev/go-sdk/thirdweb"
	)

	privateKey = "..."

	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
		PrivateKey: privateKey,
	})

	marketplace, err := sdk.GetMarketplace("{{contract_address}}")

	// Now the encoder can be accessed from the contract
 marketplace.Encoder.CreateListing(...)
```

```go
type MarketplaceEncoder struct {
    *ContractEncoder
}
```

### func \(\*MarketplaceEncoder\) [ApproveBuyoutListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace_encoder.go#L119-L125>)

```go
func (encoder *MarketplaceEncoder) ApproveBuyoutListing(ctx context.Context, signerAddress string, listingId int, quantityDesired int, receiver string) (*types.Transaction, error)
```

Get the transaction data to approve the tokens needed for a buyout listing transaction\. If the transaction wouldn't require any additional token approvals, this method will return nil\.

signerAddress: the address intended to sign the transaction

listingId: the ID of the listing to buyout

quantityDesired: the quantity of the listed tokens to purchase

receiver: the address to receive the purchased tokens

returns: the transaction data for the token approval if an approval is needed, or nil

#### Example

```
// Address of the wallet we expect to sign this message
signerAddress := "0x..."
// ID of the listing to buyout
listingId := 1
// Quantity of the listed tokens to purchase
quantityDesired := 1
// receiver address to receive the purchased tokens
receiver := "0x..."

// Transaction data required for this request
tx, err := marketplace.Encoder.ApproveBuyoutListing(context.Background(), signerAddress, listingId, quantityDesired, receiver)

// Now you can get transaction all the standard data as needed
fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
fmt.Println(tx.Nonce())
```

### func \(\*MarketplaceEncoder\) [ApproveCreateListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace_encoder.go#L257>)

```go
func (encoder *MarketplaceEncoder) ApproveCreateListing(ctx context.Context, signerAddress string, listing *NewDirectListing) (*types.Transaction, error)
```

Get the transaction data to approve the tokens needed for a create liting transaction\. If the transaction wouldn't require any additional token approvals, this method will return nil\.

signerAddress: the address intended to sign the transaction

listing: the parameters for the new direct listing to create

returns: the transaction data for the create listing transaction

#### Example

```
// Address of the wallet we expect to sign this message
signerAddress := "0x..."

listing := &NewDirectListing{
	AssetContractAddress: "0x...", // Address of the asset contract
	TokenId: 0, // Token ID of the asset to list
	StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
	ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
	Quantity: 1, // Quantity of the asset to list
	CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
	BuyoutPricePerToken: 1, // Price per token of the asset to list
}

// Transaction data required for this request
tx, err := marketplace.Encoder.ApproveCreateListing(context.Background(), signerAddress, listing)

// Now you can get transaction data as needed
fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
fmt.Println(tx.Nonce())
```

### func \(\*MarketplaceEncoder\) [BuyoutListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace_encoder.go#L183-L189>)

```go
func (encoder *MarketplaceEncoder) BuyoutListing(ctx context.Context, signerAddress string, listingId int, quantityDesired int, receiver string) (*types.Transaction, error)
```

Get the transaction data for the buyout listing transaction\. This method will throw an error if the listing requires payment in ERC20 tokens and the ERC20 tokens haven't yet been approved by the spender\. You can get the transaction data of this required approval transaction from the \`ApproveBuyoutListing\` method\.

signerAddress: the address intended to sign the transaction

listingId: the ID of the listing to buyout

quantityDesired: the quantity of the listed tokens to purchase

receiver: the address to receive the purchased tokens

returns: the transaction data for this purchase

#### Example

```
// Address of the wallet we expect to sign this message
signerAddress := "0x..."
// ID of the listing to buyout
listingId := 1
// Quantity of the listed tokens to purchase
quantityDesired := 1
// receiver address to receive the purchased tokens
receiver := "0x..."

// Transaction data required for this request
tx, err := marketplace.Encoder.BuyoutListing(context.Background(), signerAddress, listingId, quantityDesired, receiver)

// Now you can get transaction all the standard data as needed
fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
fmt.Println(tx.Nonce())
```

### func \(\*MarketplaceEncoder\) [CancelListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace_encoder.go#L80>)

```go
func (encoder *MarketplaceEncoder) CancelListing(ctx context.Context, signerAddress string, listingId int) (*types.Transaction, error)
```

#### Get the data for the transaction required to cancel a listing on the marketplace

signerAddress: the address intended to sign the transaction

listingId: the ID of the listing to cancel

returns: the transaction data for the cancellation

#### Example

```
// Address of the wallet we expect to sign this message
signerAddress := "0x..."
// ID of the listing to cancel
listingId := 1

// Transaction data required for this request
tx, err := marketplace.Encoder.CancelListing(context.Background(), signerAddress, listingId)

// Now you can get transaction all the standard data as needed
fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
fmt.Println(tx.Nonce())
```

### func \(\*MarketplaceEncoder\) [CreateListing](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/marketplace_encoder.go#L302>)

```go
func (encoder *MarketplaceEncoder) CreateListing(ctx context.Context, signerAddress string, listing *NewDirectListing) (*types.Transaction, error)
```

Get the data for the transaction required to create a direct listing\. This method will throw an error if the tokens needed for the listing have not yet been approved by the asset owner for the marketplace to spend\. You can get the transaction data of this required approval transaction from the \`ApproveCreateListing\` method\.

signerAddress: the address intended to sign the transaction

listing: the parameters for the new direct listing to create

returns: the transaction data for the create listing transaction

#### Example

```
// Address of the wallet we expect to sign this message
signerAddress := "0x..."

listing := &NewDirectListing{
	AssetContractAddress: "0x...", // Address of the asset contract
	TokenId: 0, // Token ID of the asset to list
	StartTimeInEpochSeconds: int(time.Now().Unix()), // Defaults to current time
	ListingDurationInSeconds: int(time.Now().Unix() + 3600), // Defaults to current time + 1 hour
	Quantity: 1, // Quantity of the asset to list
	CurrencyContractAddress: "0x...", // Contract address of currency to sell for, defaults to native token
	BuyoutPricePerToken: 1, // Price per token of the asset to list
}

// Transaction data required for this request
tx, err := marketplace.Encoder.CreateListing(context.Background(), signerAddress, listing)

// Now you can get transaction data as needed
fmt.Println(tx.Data()) // Ex: get the data field or the nonce field (others are available)
fmt.Println(tx.Nonce())
```

## type [MarketplaceFilter](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L509-L514>)

```go
type MarketplaceFilter struct {
    Start         int
    Count         int
    Seller        string
    TokenContract string
}
```

## type [Metadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L13-L16>)

```go
type Metadata struct {
    MetadataUri    string
    MetadataObject interface{}
}
```
