
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

### func \(\*ContractDeployer\) [DeployMarketplace](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_deployer.go#L200>)

```go
func (deployer *ContractDeployer) DeployMarketplace(metadata *DeployMarketplaceMetadata) (string, error)
```

Deploy a new Marketplace contract\.

metadata: the contract metadata

returns: the address of the deployed contract

#### Example

```
address, err := sdk.Deployer.DeployMarketplace(
	&thirdweb.DeployMarketplaceMetadata{
		Name: "Go Marketplace",
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
