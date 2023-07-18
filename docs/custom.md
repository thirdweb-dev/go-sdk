
## Custom Contracts

### Custom Contracts

With the thirdweb SDK, you can get a contract instance for any contract. Additionally, if you deployed your contract using thirdweb deploy, you can get a more explicit and intuitive interface to interact with your contracts.

\# Getting a Custom Contract Instance

Let's take a look at how you can get a custom contract instance for one of your contracts deployed using the thirdweb deploy flow:

```
import (
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

privateKey := "..."
secretKey := "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
	SecretKey: secretKey
})

// You can replace your own contract address here
contractAddress := "{{contract_address}}"

// Now you have a contract instance ready to go
contract, err := sdk.GetContract(contractAddress)
```

Alternatively, if you didn't deploy your contract with thirdweb deploy, you can still get a contract instance for any contract using your contracts ABI:

```
import (
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

// You can replace your own contract address here
contractAddress := "{{contract_address}}"

// Add your contract ABI here
abi := "[...]"

// Now you have a contract instance ready to go
contract, err := sdk.GetContractFromAbi(contractAddress, abi)
```

\# Calling Contract Functions

Now that you have an SDK instance for your contract, you can easily call any function on your contract with the contract "call" method as follows:

```
// The first parameter to the call function is the method name
// All other parameters to the call function get passed as arguments to your contract
balance, err := contract.Call("balanceOf", "{{wallet_address}}")

// You can also make a transaction to your contract with the call method
tx, err := contract.Call("mintTo", "{{wallet_address}}", "ipfs://...")
```

```go
type SmartContract struct {
    Helper  *contractHelper
    Encoder *ContractEncoder
    Events  *ContractEvents
    ERC20   *ERC20
    ERC721  *ERC721
    ERC1155 *ERC1155
}
```

### func \(\*SmartContract\) [Call](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/smart_contract.go#L155>)

```go
func (c *SmartContract) Call(ctx context.Context, method string, args ...interface{}) (interface{}, error)
```

Call any function on your contract.

method: the name of the method on your contract you want to call

args: the arguments to pass to the method

#### Example

```
// The first parameter to the call function is the method name
// All other parameters to the call function get passed as arguments to your contract
balance, err := contract.Call("balanceOf", "{{wallet_address}}")

// You can also make a transaction to your contract with the call method
tx, err := contract.Call(context.Background(), "mintTo", "{{wallet_address}}", "ipfs://...")
```

## type [SnapshotClaim](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/snapshots.go#L18-L22>)

```go
type SnapshotClaim struct {
    Address      string   `json:"address"`
    MaxClaimable int      `json:"maxClaimable"`
    Proof        []string `json:"proof"`
}
```

## type [SnapshotEntry](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L593-L598>)

```go
type SnapshotEntry struct {
    Address         string `json:"address"`
    MaxClaimable    string `json:"maxClaimable"`
    Price           string `json:"price"`
    CurrencyAddress string `json:"currencyAddress"`
}
```

## type [SnapshotEntryWithProof](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L571-L577>)

```go
type SnapshotEntryWithProof struct {
    Address         string
    MaxClaimable    string
    Price           string
    CurrencyAddress string
    Proof           [][32]byte
}
```

## type [SnapshotInfo](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/snapshots.go#L24-L27>)

```go
type SnapshotInfo struct {
    MerkleRoot string          `json:"merkleRoot"`
    Claims     []SnapshotClaim `json:"claims"`
}
```

## type [SnapshotInfos](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/snapshots.go#L29-L33>)

```go
type SnapshotInfos struct {
    Snapshot    SnapshotInfo
    MerkleRoot  string
    SnapshotUri string
}
```

## type [SnapshotInput](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/snapshots.go#L13-L16>)

```go
type SnapshotInput struct {
    Address      string
    MaxClaimable int
}
```
