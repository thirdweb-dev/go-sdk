
## Contract Deployments

The contract deployer lets you deploy new contracts to the blockchain using just the thirdweb SDK\. You can access the contract deployer interface as follows:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

// Now you can deploy a contract
address, err := sdk.Deployer.DeployNFTCollection(
	&thirdweb.DeployNFTCollectionMetadata{
		Name: "Go NFT",
	}
})
```

```go
type ContractDeployer struct {
    *ProviderHandler
}
```

### func \(\*ContractDeployer\) [DeployEdition](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L110>)

```go
func (deployer *ContractDeployer) DeployEdition(metadata *DeployEditionMetadata) (string, error)
```

Deploy a new Edition contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEdition(
	&thirdweb.DeployEditionMetadata{
		Name: "Go Edition",
	}
})
```

### func \(\*ContractDeployer\) [DeployEditionDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L164>)

```go
func (deployer *ContractDeployer) DeployEditionDrop(metadata *DeployEditionDropMetadata) (string, error)
```

Deploy a new Edition Drop contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployEditionDrop(
	&thirdweb.DeployEditionDropMetadata{
		Name: "Go Edition Drop",
	}
})
```

### func \(\*ContractDeployer\) [DeployMultiwrap](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L182>)

```go
func (deployer *ContractDeployer) DeployMultiwrap(metadata *DeployMultiwrapMetadata) (string, error)
```

Deploy a new Multiwrap contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMultiwrap(
	&thirdweb.DeployMultiwrapMetadata{
		Name: "Go Multiwrap",
	}
})
```

### func \(\*ContractDeployer\) [DeployNFTCollection](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L92>)

```go
func (deployer *ContractDeployer) DeployNFTCollection(metadata *DeployNFTCollectionMetadata) (string, error)
```

Deploy a new NFT Collection contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTCollection(
	&thirdweb.DeployNFTCollectionMetadata{
		Name: "Go NFT",
	}
})
```

### func \(\*ContractDeployer\) [DeployNFTDrop](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L146>)

```go
func (deployer *ContractDeployer) DeployNFTDrop(metadata *DeployNFTDropMetadata) (string, error)
```

Deploy a new NFT Drop contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployNFTDrop(
	&thirdweb.DeployNFTDropMetadata{
		Name: "Go NFT Drop",
	}
})
```

### func \(\*ContractDeployer\) [DeployToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L128>)

```go
func (deployer *ContractDeployer) DeployToken(metadata *DeployTokenMetadata) (string, error)
```

Deploy a new Token contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployToken(
	&thirdweb.DeployTokenMetadata{
		Name: "Go Token",
	}
})
```

## type [Currency](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L89-L93>)

```go
type Currency struct {
    Name     string
    Symbol   string
    Decimals int
}
```

## type [CurrencyValue](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L95-L101>)

```go
type CurrencyValue struct {
    Name         string
    Symbol       string
    Decimals     int
    Value        *big.Int
    DisplayValue float64
}
```

## type [DeployEditionDropMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L392-L405>)

```go
type DeployEditionDropMetadata struct {
    Name                   string      `mapstructure:"name" json:"name"`
    Description            string      `mapstructure:"description" json:"description"`
    Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
    ExternalLink           string      `mapstructure:"external_link" json:"external_link"`
    SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
    FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
    Merkle                 interface{} `mapstructure:"merkle" json:"merkle"`
    Symbol                 string      `mapstructure:"symbol" json:"symbol"`
    PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
    PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
    PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
    TrustedForwarders      []string    `mapstructure:"trusted_forwarders,omitempty" json:"trusted_forwarders"`
}
```

## type [DeployEditionMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L301-L313>)

```go
type DeployEditionMetadata struct {
    Name                   string      `mapstructure:"name" json:"name"`
    Description            string      `mapstructure:"description" json:"description"`
    Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
    ExternalLink           string      `mapstructure:"external_link" json:"external_link"`
    SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
    FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
    Symbol                 string      `mapstructure:"symbol" json:"symbol"`
    PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
    PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
    PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
    TrustedForwarders      []string    `mapstructure:"trusted_forwarders,omitempty" json:"trusted_forwarders"`
}
```

## type [DeployMultiwrapMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L425-L434>)

```go
type DeployMultiwrapMetadata struct {
    Name                 string      `mapstructure:"name" json:"name"`
    Description          string      `mapstructure:"description" json:"description"`
    Image                interface{} `mapstructure:"image,omitempty" json:"image"`
    ExternalLink         string      `mapstructure:"external_link" json:"external_link"`
    SellerFeeBasisPoints int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
    FeeRecipient         string      `mapstructure:"fee_recipient" json:"fee_recipient"`
    Symbol               string      `mapstructure:"symbol" json:"symbol"`
    TrustedForwarders    []string    `mapstructure:"trusted_forwarders,omitempty" json:"trusted_forwarders"`
}
```

## type [DeployNFTCollectionMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L269-L281>)

```go
type DeployNFTCollectionMetadata struct {
    Name                   string      `mapstructure:"name" json:"name"`
    Description            string      `mapstructure:"description" json:"description"`
    Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
    ExternalLink           string      `mapstructure:"external_link" json:"external_link"`
    SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
    FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
    Symbol                 string      `mapstructure:"symbol" json:"symbol"`
    PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
    PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
    PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
    TrustedForwarders      []string    `mapstructure:"trusted_forwarders,omitempty" json:"trusted_forwarders"`
}
```

## type [DeployNFTDropMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L359-L372>)

```go
type DeployNFTDropMetadata struct {
    Name                   string      `mapstructure:"name" json:"name"`
    Description            string      `mapstructure:"description" json:"description"`
    Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
    ExternalLink           string      `mapstructure:"external_link" json:"external_link"`
    SellerFeeBasisPoints   int         `mapstructure:"seller_fee_basis_points" json:"seller_fee_basis_points"`
    FeeRecipient           string      `mapstructure:"fee_recipient" json:"fee_recipient"`
    Merkle                 interface{} `mapstructure:"merkle" json:"merkle"`
    Symbol                 string      `mapstructure:"symbol" json:"symbol"`
    PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
    PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
    PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
    TrustedForwarders      []string    `mapstructure:"trusted_forwarders,omitempty" json:"trusted_forwarders"`
}
```

## type [DeployTokenMetadata](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L333-L343>)

```go
type DeployTokenMetadata struct {
    Name                   string      `mapstructure:"name" json:"name"`
    Description            string      `mapstructure:"description" json:"description"`
    Image                  interface{} `mapstructure:"image,omitempty" json:"image"`
    ExternalLink           string      `mapstructure:"external_link" json:"external_link"`
    Symbol                 string      `mapstructure:"symbol" json:"symbol"`
    PrimarySaleRecipient   string      `mapstructure:"primary_sale_recipient" json:"primary_sale_recipient"`
    PlatformFeeBasisPoints int         `mapstructure:"platform_fee_basis_points" json:"platform_fee_basis_points"`
    PlatformFeeRecipient   string      `mapstructure:"platform_fee_recipient" json:"platform_fee_recipient"`
    TrustedForwarders      []string    `mapstructure:"trusted_forwarders,omitempty" json:"trusted_forwarders"`
}
```
