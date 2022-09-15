package thirdweb

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Custom Contracts
//
// With the thirdweb SDK, you can get a contract instance for any contract. Additionally,
// if you deployed your contract using thirdweb deploy, you can get a more explicit and
// intuitive interface to interact with your contracts.
//
// # Getting a Custom Contract Instance
//
// Let's take a look at how you can get a custom contract instance for one of your contracts
// deployed using the thirdweb deploy flow:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	// You can replace your own contract address here
//	contractAddress := "{{contract_address}}"
//
//	// Now you have a contract instance ready to go
//	contract, err := sdk.GetContract(contractAddress)
//
// Alternatively, if you didn't deploy your contract with thirdweb deploy, you can still get a
// contract instance for any contract using your contracts ABI:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	// You can replace your own contract address here
//	contractAddress := "{{contract_address}}"
//
//	// Add your contract ABI here
//	abi := "[...]"
//
//	// Now you have a contract instance ready to go
//	contract, err := sdk.GetContractFromAbi(contractAddress, abi)
//
// # Calling Contract Functions
//
// Now that you have an SDK instance for your contract, you can easily call any function on your contract
// with the contract "call" method as follows:
//
//	// The first parameter to the call function is the method name
//	// All other parameters to the call function get passed as arguments to your contract
//	balance, err := contract.Call("balanceOf", "{{wallet_address}}")
//
//	// You can also make a transaction to your contract with the call method
//	tx, err := contract.Call("mintTo", "{{wallet_address}}", "ipfs://...")
type SmartContract struct {
	abi      *abi.ABI
	contract *bind.BoundContract
	helper   *contractHelper
}

func newSmartContract(provider *ethclient.Client, address common.Address, contractAbi string, privateKey string, storage storage) (*SmartContract, error) {
	helper, err := newContractHelper(address, provider, privateKey)
	fmt.Println(privateKey)
	if err != nil {
		return nil, err
	}

	parsedAbi, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	boundContract := bind.NewBoundContract(address, parsedAbi, provider, provider, provider)
	contract := &SmartContract{
		abi:      &parsedAbi,
		contract: boundContract,
		helper:   helper,
	}

	return contract, nil
}

// Call any function on your contract.
//
// method: the name of the method on your contract you want to call
//
// args: the arguments to pass to the method
//
// Example
//
//	// The first parameter to the call function is the method name
//	// All other parameters to the call function get passed as arguments to your contract
//	balance, err := contract.Call("balanceOf", "{{wallet_address}}")
//
//	// You can also make a transaction to your contract with the call method
//	tx, err := contract.Call(context.Background(), "mintTo", "{{wallet_address}}", "ipfs://...")
func (c *SmartContract) Call(ctx context.Context, method string, args ...interface{}) (interface{}, error) {
	abiMethod, exist := c.abi.Methods[method]
	if !exist {
		return nil, fmt.Errorf("function '%s' not found in contract '%s'", method, c.helper.getAddress().String())
	}

	if len(abiMethod.Inputs) != len(args) {
		return nil, fmt.Errorf(
			"function '%s' requires %d arguments, but %d arguments were provided.\nExpected function signature '%s'",
			method,
			len(abiMethod.Inputs),
			len(args),
			abiMethod.Sig,
		)
	}

	// Validate argument input types and convert to proper input types for contract
	// So we can allow users to pass in string intead of address, int instead of big, etc.
	typedArgs := []interface{}{}
	for i, arg := range args {
		input := abiMethod.Inputs[i]
		inputType := fmt.Sprint(input.Type)

		if inputType == "address" {
			parsedArg, ok := arg.(string)
			if !ok {
				return nil, fmt.Errorf(
					"argument %d (%v) should be of type 'string', but type '%v' was provided",
					i,
					input.Name,
					reflect.TypeOf(arg),
				)
			}

			arg = common.HexToAddress(parsedArg)
		} else if strings.Contains(inputType, "int") {
			parsedArg, ok := arg.(int)
			if !ok {
				return nil, fmt.Errorf(
					"argument %d (%v) should be of type 'int', but type '%v' was provided",
					i,
					input.Name,
					reflect.TypeOf(arg),
				)
			}

			arg = big.NewInt(int64(parsedArg))
		}

		typedArgs = append(typedArgs, arg)
	}

	if abiMethod.StateMutability == "view" {
		var out []interface{}
		err := c.contract.Call(&bind.CallOpts{}, &out, method, typedArgs...)
		if err != nil {
			return nil, err
		}

		// If theres only one return value, return it directly instead of the tuple
		if len(out) == 1 {
			return out[0], nil
		}

		return out, nil
	} else {
		txOpts, err := c.helper.getTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := c.contract.Transact(txOpts, method, typedArgs...)
		if err != nil {
			return nil, err
		}

		return c.helper.awaitTx(tx.Hash())
	}
}
