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
	"github.com/ethereum/go-ethereum/core/types"
)

// This interface is currently supported by all contract encoder classes and provides a generic
// method to encode write function calls.
type ContractEncoder struct {
	abi      *abi.ABI
	contract *bind.BoundContract
	helper   *contractHelper
}

func newContractEncoder(contractAbi string, helper *contractHelper) (*ContractEncoder, error) {
	parsedAbi, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(helper.getAddress(), parsedAbi, helper.GetProvider(), helper.GetProvider(), helper.GetProvider())

	return &ContractEncoder{
		abi:      &parsedAbi,
		contract: contract,
		helper:   helper,
	}, nil
}

// Get the unsigned transaction data for any contract call on a contract.
//
// signerAddress: the address expected to sign this transaction
//
// method: the name of the contract function to encode transaction data for
//
// args: the arguments to pass to the contract function.
//
// returns: the encoded transaction data for the transaction.
//
// Example
//
//	toAddress := "0x..."
//	amount := 1
//
//	// Now you can get the transaction data for the contract call.
//	tx, err := contract.Encoder.Encode(context.Background(), "transfer", toAddress, amount)
//	fmt.Println(tx.Data()) // Now you can access all transaction data, like the following fields
//	fmt.Println(tx.Nonce())
//	fmt.Println(tx.Value())
func (encoder *ContractEncoder) Encode(ctx context.Context, signerAddress string, method string, args ...interface{}) (*types.Transaction, error) {
	abiMethod, exist := encoder.abi.Methods[method]
	if !exist {
		return nil, fmt.Errorf("function '%s' not found in contract '%s'", method, encoder.helper.getAddress().String())
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

	txOpts, err := encoder.helper.getUnsignedTxOptions(ctx, signerAddress)
	if err != nil {
		return nil, err
	}
	return encoder.contract.Transact(txOpts, method, typedArgs...)
}
