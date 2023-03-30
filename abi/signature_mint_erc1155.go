// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ISignatureMintERC1155MintRequest is an auto generated low-level Go binding around an user-defined struct.
type ISignatureMintERC1155MintRequest struct {
	To                     common.Address
	RoyaltyRecipient       common.Address
	RoyaltyBps             *big.Int
	PrimarySaleRecipient   common.Address
	TokenId                *big.Int
	Uri                    string
	Quantity               *big.Int
	PricePerToken          *big.Int
	Currency               common.Address
	ValidityStartTimestamp *big.Int
	ValidityEndTimestamp   *big.Int
	Uid                    [32]byte
}

// SignatureMintERC1155MetaData contains all meta data concerning the SignatureMintERC1155 contract.
var SignatureMintERC1155MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintedTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIdMinted\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"primarySaleRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validityStartTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validityEndTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structISignatureMintERC1155.MintRequest\",\"name\":\"mintRequest\",\"type\":\"tuple\"}],\"name\":\"TokensMintedWithSignature\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"primarySaleRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validityStartTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validityEndTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"internalType\":\"structISignatureMintERC1155.MintRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintWithSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"primarySaleRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validityStartTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validityEndTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"internalType\":\"structISignatureMintERC1155.MintRequest\",\"name\":\"_req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SignatureMintERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureMintERC1155MetaData.ABI instead.
var SignatureMintERC1155ABI = SignatureMintERC1155MetaData.ABI

// SignatureMintERC1155 is an auto generated Go binding around an Ethereum contract.
type SignatureMintERC1155 struct {
	SignatureMintERC1155Caller     // Read-only binding to the contract
	SignatureMintERC1155Transactor // Write-only binding to the contract
	SignatureMintERC1155Filterer   // Log filterer for contract events
}

// SignatureMintERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureMintERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureMintERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureMintERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureMintERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureMintERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureMintERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureMintERC1155Session struct {
	Contract     *SignatureMintERC1155 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SignatureMintERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureMintERC1155CallerSession struct {
	Contract *SignatureMintERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SignatureMintERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureMintERC1155TransactorSession struct {
	Contract     *SignatureMintERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SignatureMintERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureMintERC1155Raw struct {
	Contract *SignatureMintERC1155 // Generic contract binding to access the raw methods on
}

// SignatureMintERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureMintERC1155CallerRaw struct {
	Contract *SignatureMintERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// SignatureMintERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureMintERC1155TransactorRaw struct {
	Contract *SignatureMintERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureMintERC1155 creates a new instance of SignatureMintERC1155, bound to a specific deployed contract.
func NewSignatureMintERC1155(address common.Address, backend bind.ContractBackend) (*SignatureMintERC1155, error) {
	contract, err := bindSignatureMintERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC1155{SignatureMintERC1155Caller: SignatureMintERC1155Caller{contract: contract}, SignatureMintERC1155Transactor: SignatureMintERC1155Transactor{contract: contract}, SignatureMintERC1155Filterer: SignatureMintERC1155Filterer{contract: contract}}, nil
}

// NewSignatureMintERC1155Caller creates a new read-only instance of SignatureMintERC1155, bound to a specific deployed contract.
func NewSignatureMintERC1155Caller(address common.Address, caller bind.ContractCaller) (*SignatureMintERC1155Caller, error) {
	contract, err := bindSignatureMintERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC1155Caller{contract: contract}, nil
}

// NewSignatureMintERC1155Transactor creates a new write-only instance of SignatureMintERC1155, bound to a specific deployed contract.
func NewSignatureMintERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*SignatureMintERC1155Transactor, error) {
	contract, err := bindSignatureMintERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC1155Transactor{contract: contract}, nil
}

// NewSignatureMintERC1155Filterer creates a new log filterer instance of SignatureMintERC1155, bound to a specific deployed contract.
func NewSignatureMintERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*SignatureMintERC1155Filterer, error) {
	contract, err := bindSignatureMintERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC1155Filterer{contract: contract}, nil
}

// bindSignatureMintERC1155 binds a generic wrapper to an already deployed contract.
func bindSignatureMintERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignatureMintERC1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureMintERC1155 *SignatureMintERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureMintERC1155.Contract.SignatureMintERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureMintERC1155 *SignatureMintERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureMintERC1155.Contract.SignatureMintERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureMintERC1155 *SignatureMintERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureMintERC1155.Contract.SignatureMintERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureMintERC1155 *SignatureMintERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureMintERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureMintERC1155 *SignatureMintERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureMintERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureMintERC1155 *SignatureMintERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureMintERC1155.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0xb17cd86f.
//
// Solidity: function verify((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool success, address signer)
func (_SignatureMintERC1155 *SignatureMintERC1155Caller) Verify(opts *bind.CallOpts, _req ISignatureMintERC1155MintRequest, _signature []byte) (struct {
	Success bool
	Signer  common.Address
}, error) {
	var out []interface{}
	err := _SignatureMintERC1155.contract.Call(opts, &out, "verify", _req, _signature)

	outstruct := new(struct {
		Success bool
		Signer  common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Success = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Signer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Verify is a free data retrieval call binding the contract method 0xb17cd86f.
//
// Solidity: function verify((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool success, address signer)
func (_SignatureMintERC1155 *SignatureMintERC1155Session) Verify(_req ISignatureMintERC1155MintRequest, _signature []byte) (struct {
	Success bool
	Signer  common.Address
}, error) {
	return _SignatureMintERC1155.Contract.Verify(&_SignatureMintERC1155.CallOpts, _req, _signature)
}

// Verify is a free data retrieval call binding the contract method 0xb17cd86f.
//
// Solidity: function verify((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool success, address signer)
func (_SignatureMintERC1155 *SignatureMintERC1155CallerSession) Verify(_req ISignatureMintERC1155MintRequest, _signature []byte) (struct {
	Success bool
	Signer  common.Address
}, error) {
	return _SignatureMintERC1155.Contract.Verify(&_SignatureMintERC1155.CallOpts, _req, _signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x98a6e993.
//
// Solidity: function mintWithSignature((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) req, bytes signature) payable returns(address signer)
func (_SignatureMintERC1155 *SignatureMintERC1155Transactor) MintWithSignature(opts *bind.TransactOpts, req ISignatureMintERC1155MintRequest, signature []byte) (*types.Transaction, error) {
	return _SignatureMintERC1155.contract.Transact(opts, "mintWithSignature", req, signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x98a6e993.
//
// Solidity: function mintWithSignature((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) req, bytes signature) payable returns(address signer)
func (_SignatureMintERC1155 *SignatureMintERC1155Session) MintWithSignature(req ISignatureMintERC1155MintRequest, signature []byte) (*types.Transaction, error) {
	return _SignatureMintERC1155.Contract.MintWithSignature(&_SignatureMintERC1155.TransactOpts, req, signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x98a6e993.
//
// Solidity: function mintWithSignature((address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) req, bytes signature) payable returns(address signer)
func (_SignatureMintERC1155 *SignatureMintERC1155TransactorSession) MintWithSignature(req ISignatureMintERC1155MintRequest, signature []byte) (*types.Transaction, error) {
	return _SignatureMintERC1155.Contract.MintWithSignature(&_SignatureMintERC1155.TransactOpts, req, signature)
}

// SignatureMintERC1155TokensMintedWithSignatureIterator is returned from FilterTokensMintedWithSignature and is used to iterate over the raw logs and unpacked data for TokensMintedWithSignature events raised by the SignatureMintERC1155 contract.
type SignatureMintERC1155TokensMintedWithSignatureIterator struct {
	Event *SignatureMintERC1155TokensMintedWithSignature // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SignatureMintERC1155TokensMintedWithSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureMintERC1155TokensMintedWithSignature)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SignatureMintERC1155TokensMintedWithSignature)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SignatureMintERC1155TokensMintedWithSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureMintERC1155TokensMintedWithSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureMintERC1155TokensMintedWithSignature represents a TokensMintedWithSignature event raised by the SignatureMintERC1155 contract.
type SignatureMintERC1155TokensMintedWithSignature struct {
	Signer        common.Address
	MintedTo      common.Address
	TokenIdMinted *big.Int
	MintRequest   ISignatureMintERC1155MintRequest
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensMintedWithSignature is a free log retrieval operation binding the contract event 0x0b35afaf155daeef41cc46df86f058df2855c57d30ab134647a6b587e7cc8c39.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_SignatureMintERC1155 *SignatureMintERC1155Filterer) FilterTokensMintedWithSignature(opts *bind.FilterOpts, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (*SignatureMintERC1155TokensMintedWithSignatureIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _SignatureMintERC1155.contract.FilterLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC1155TokensMintedWithSignatureIterator{contract: _SignatureMintERC1155.contract, event: "TokensMintedWithSignature", logs: logs, sub: sub}, nil
}

// WatchTokensMintedWithSignature is a free log subscription operation binding the contract event 0x0b35afaf155daeef41cc46df86f058df2855c57d30ab134647a6b587e7cc8c39.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_SignatureMintERC1155 *SignatureMintERC1155Filterer) WatchTokensMintedWithSignature(opts *bind.WatchOpts, sink chan<- *SignatureMintERC1155TokensMintedWithSignature, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}
	var mintedToRule []interface{}
	for _, mintedToItem := range mintedTo {
		mintedToRule = append(mintedToRule, mintedToItem)
	}
	var tokenIdMintedRule []interface{}
	for _, tokenIdMintedItem := range tokenIdMinted {
		tokenIdMintedRule = append(tokenIdMintedRule, tokenIdMintedItem)
	}

	logs, sub, err := _SignatureMintERC1155.contract.WatchLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureMintERC1155TokensMintedWithSignature)
				if err := _SignatureMintERC1155.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensMintedWithSignature is a log parse operation binding the contract event 0x0b35afaf155daeef41cc46df86f058df2855c57d30ab134647a6b587e7cc8c39.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,uint256,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_SignatureMintERC1155 *SignatureMintERC1155Filterer) ParseTokensMintedWithSignature(log types.Log) (*SignatureMintERC1155TokensMintedWithSignature, error) {
	event := new(SignatureMintERC1155TokensMintedWithSignature)
	if err := _SignatureMintERC1155.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
