
## Custom Contracts

### Custom Contracts

With the thirdweb SDK, you can get a contract instance for any contract\. Additionally, if you deployed your contract using thirdweb deploy, you can get a more explicit and intuitive interface to interact with your contracts\.

\# Getting a Custom Contract Instance

Let's take a look at how you can get a custom contract instance for one of your contracts deployed using the thirdweb deploy flow:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

privateKey = "..."

sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
	PrivateKey: privateKey,
})

// You can replace your own contract address here
contractAddress := "{{contract_address}}"

// Now you have a contract instance ready to go
contract, err := sdk.GetContract(contractAddress)
```

Alternatively, if you didn't deploy your contract with thirdweb deploy, you can still get a contract instance for any contract using your contracts ABI:

```
import (
	"github.com/thirdweb-dev/go-sdk/thirdweb"
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
type SmartContract struct {}
```

### func \(\*SmartContract\) [Call](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/smart_contract.go#L118>)

```go
func (c *SmartContract) Call(ctx context.Context, method string, args ...interface{}) (interface{}, error)
```

Call any function on your contract\.

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
