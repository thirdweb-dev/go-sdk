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

// ISignatureMintERC721MintRequest is an auto generated low-level Go binding around an user-defined struct.
type ISignatureMintERC721MintRequest struct {
	To                     common.Address
	RoyaltyRecipient       common.Address
	RoyaltyBps             *big.Int
	PrimarySaleRecipient   common.Address
	Uri                    string
	Quantity               *big.Int
	PricePerToken          *big.Int
	Currency               common.Address
	ValidityStartTimestamp *big.Int
	ValidityEndTimestamp   *big.Int
	Uid                    [32]byte
}

// SignatureMintERC721MetaData contains all meta data concerning the SignatureMintERC721 contract.
var SignatureMintERC721MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintedTo\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIdMinted\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"primarySaleRecipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validityStartTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validityEndTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structISignatureMintERC721.MintRequest\",\"name\":\"mintRequest\",\"type\":\"tuple\"}],\"name\":\"TokensMintedWithSignature\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"primarySaleRecipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validityStartTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validityEndTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"internalType\":\"structISignatureMintERC721.MintRequest\",\"name\":\"req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mintWithSignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"primarySaleRecipient\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"validityStartTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"validityEndTimestamp\",\"type\":\"uint128\"},{\"internalType\":\"bytes32\",\"name\":\"uid\",\"type\":\"bytes32\"}],\"internalType\":\"structISignatureMintERC721.MintRequest\",\"name\":\"_req\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"verify\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SignatureMintERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use SignatureMintERC721MetaData.ABI instead.
var SignatureMintERC721ABI = SignatureMintERC721MetaData.ABI

// SignatureMintERC721 is an auto generated Go binding around an Ethereum contract.
type SignatureMintERC721 struct {
	SignatureMintERC721Caller     // Read-only binding to the contract
	SignatureMintERC721Transactor // Write-only binding to the contract
	SignatureMintERC721Filterer   // Log filterer for contract events
}

// SignatureMintERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type SignatureMintERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureMintERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SignatureMintERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureMintERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SignatureMintERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SignatureMintERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SignatureMintERC721Session struct {
	Contract     *SignatureMintERC721 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SignatureMintERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SignatureMintERC721CallerSession struct {
	Contract *SignatureMintERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SignatureMintERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SignatureMintERC721TransactorSession struct {
	Contract     *SignatureMintERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SignatureMintERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type SignatureMintERC721Raw struct {
	Contract *SignatureMintERC721 // Generic contract binding to access the raw methods on
}

// SignatureMintERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SignatureMintERC721CallerRaw struct {
	Contract *SignatureMintERC721Caller // Generic read-only contract binding to access the raw methods on
}

// SignatureMintERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SignatureMintERC721TransactorRaw struct {
	Contract *SignatureMintERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSignatureMintERC721 creates a new instance of SignatureMintERC721, bound to a specific deployed contract.
func NewSignatureMintERC721(address common.Address, backend bind.ContractBackend) (*SignatureMintERC721, error) {
	contract, err := bindSignatureMintERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC721{SignatureMintERC721Caller: SignatureMintERC721Caller{contract: contract}, SignatureMintERC721Transactor: SignatureMintERC721Transactor{contract: contract}, SignatureMintERC721Filterer: SignatureMintERC721Filterer{contract: contract}}, nil
}

// NewSignatureMintERC721Caller creates a new read-only instance of SignatureMintERC721, bound to a specific deployed contract.
func NewSignatureMintERC721Caller(address common.Address, caller bind.ContractCaller) (*SignatureMintERC721Caller, error) {
	contract, err := bindSignatureMintERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC721Caller{contract: contract}, nil
}

// NewSignatureMintERC721Transactor creates a new write-only instance of SignatureMintERC721, bound to a specific deployed contract.
func NewSignatureMintERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*SignatureMintERC721Transactor, error) {
	contract, err := bindSignatureMintERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC721Transactor{contract: contract}, nil
}

// NewSignatureMintERC721Filterer creates a new log filterer instance of SignatureMintERC721, bound to a specific deployed contract.
func NewSignatureMintERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*SignatureMintERC721Filterer, error) {
	contract, err := bindSignatureMintERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC721Filterer{contract: contract}, nil
}

// bindSignatureMintERC721 binds a generic wrapper to an already deployed contract.
func bindSignatureMintERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SignatureMintERC721MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureMintERC721 *SignatureMintERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureMintERC721.Contract.SignatureMintERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureMintERC721 *SignatureMintERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureMintERC721.Contract.SignatureMintERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureMintERC721 *SignatureMintERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureMintERC721.Contract.SignatureMintERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SignatureMintERC721 *SignatureMintERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SignatureMintERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SignatureMintERC721 *SignatureMintERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SignatureMintERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SignatureMintERC721 *SignatureMintERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SignatureMintERC721.Contract.contract.Transact(opts, method, params...)
}

// Verify is a free data retrieval call binding the contract method 0x252e82e8.
//
// Solidity: function verify((address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool success, address signer)
func (_SignatureMintERC721 *SignatureMintERC721Caller) Verify(opts *bind.CallOpts, _req ISignatureMintERC721MintRequest, _signature []byte) (struct {
	Success bool
	Signer  common.Address
}, error) {
	var out []interface{}
	err := _SignatureMintERC721.contract.Call(opts, &out, "verify", _req, _signature)

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

// Verify is a free data retrieval call binding the contract method 0x252e82e8.
//
// Solidity: function verify((address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool success, address signer)
func (_SignatureMintERC721 *SignatureMintERC721Session) Verify(_req ISignatureMintERC721MintRequest, _signature []byte) (struct {
	Success bool
	Signer  common.Address
}, error) {
	return _SignatureMintERC721.Contract.Verify(&_SignatureMintERC721.CallOpts, _req, _signature)
}

// Verify is a free data retrieval call binding the contract method 0x252e82e8.
//
// Solidity: function verify((address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) _req, bytes _signature) view returns(bool success, address signer)
func (_SignatureMintERC721 *SignatureMintERC721CallerSession) Verify(_req ISignatureMintERC721MintRequest, _signature []byte) (struct {
	Success bool
	Signer  common.Address
}, error) {
	return _SignatureMintERC721.Contract.Verify(&_SignatureMintERC721.CallOpts, _req, _signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x439c7be5.
//
// Solidity: function mintWithSignature((address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) req, bytes signature) payable returns(address signer)
func (_SignatureMintERC721 *SignatureMintERC721Transactor) MintWithSignature(opts *bind.TransactOpts, req ISignatureMintERC721MintRequest, signature []byte) (*types.Transaction, error) {
	return _SignatureMintERC721.contract.Transact(opts, "mintWithSignature", req, signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x439c7be5.
//
// Solidity: function mintWithSignature((address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) req, bytes signature) payable returns(address signer)
func (_SignatureMintERC721 *SignatureMintERC721Session) MintWithSignature(req ISignatureMintERC721MintRequest, signature []byte) (*types.Transaction, error) {
	return _SignatureMintERC721.Contract.MintWithSignature(&_SignatureMintERC721.TransactOpts, req, signature)
}

// MintWithSignature is a paid mutator transaction binding the contract method 0x439c7be5.
//
// Solidity: function mintWithSignature((address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) req, bytes signature) payable returns(address signer)
func (_SignatureMintERC721 *SignatureMintERC721TransactorSession) MintWithSignature(req ISignatureMintERC721MintRequest, signature []byte) (*types.Transaction, error) {
	return _SignatureMintERC721.Contract.MintWithSignature(&_SignatureMintERC721.TransactOpts, req, signature)
}

// SignatureMintERC721TokensMintedWithSignatureIterator is returned from FilterTokensMintedWithSignature and is used to iterate over the raw logs and unpacked data for TokensMintedWithSignature events raised by the SignatureMintERC721 contract.
type SignatureMintERC721TokensMintedWithSignatureIterator struct {
	Event *SignatureMintERC721TokensMintedWithSignature // Event containing the contract specifics and raw log

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
func (it *SignatureMintERC721TokensMintedWithSignatureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SignatureMintERC721TokensMintedWithSignature)
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
		it.Event = new(SignatureMintERC721TokensMintedWithSignature)
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
func (it *SignatureMintERC721TokensMintedWithSignatureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SignatureMintERC721TokensMintedWithSignatureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SignatureMintERC721TokensMintedWithSignature represents a TokensMintedWithSignature event raised by the SignatureMintERC721 contract.
type SignatureMintERC721TokensMintedWithSignature struct {
	Signer        common.Address
	MintedTo      common.Address
	TokenIdMinted *big.Int
	MintRequest   ISignatureMintERC721MintRequest
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokensMintedWithSignature is a free log retrieval operation binding the contract event 0xee0cf9c3e87795b1932d13f80f892f620f567b4465e768ced5d64aa44ca1d64c.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_SignatureMintERC721 *SignatureMintERC721Filterer) FilterTokensMintedWithSignature(opts *bind.FilterOpts, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (*SignatureMintERC721TokensMintedWithSignatureIterator, error) {

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

	logs, sub, err := _SignatureMintERC721.contract.FilterLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return &SignatureMintERC721TokensMintedWithSignatureIterator{contract: _SignatureMintERC721.contract, event: "TokensMintedWithSignature", logs: logs, sub: sub}, nil
}

// WatchTokensMintedWithSignature is a free log subscription operation binding the contract event 0xee0cf9c3e87795b1932d13f80f892f620f567b4465e768ced5d64aa44ca1d64c.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_SignatureMintERC721 *SignatureMintERC721Filterer) WatchTokensMintedWithSignature(opts *bind.WatchOpts, sink chan<- *SignatureMintERC721TokensMintedWithSignature, signer []common.Address, mintedTo []common.Address, tokenIdMinted []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _SignatureMintERC721.contract.WatchLogs(opts, "TokensMintedWithSignature", signerRule, mintedToRule, tokenIdMintedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SignatureMintERC721TokensMintedWithSignature)
				if err := _SignatureMintERC721.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
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

// ParseTokensMintedWithSignature is a log parse operation binding the contract event 0xee0cf9c3e87795b1932d13f80f892f620f567b4465e768ced5d64aa44ca1d64c.
//
// Solidity: event TokensMintedWithSignature(address indexed signer, address indexed mintedTo, uint256 indexed tokenIdMinted, (address,address,uint256,address,string,uint256,uint256,address,uint128,uint128,bytes32) mintRequest)
func (_SignatureMintERC721 *SignatureMintERC721Filterer) ParseTokensMintedWithSignature(log types.Log) (*SignatureMintERC721TokensMintedWithSignature, error) {
	event := new(SignatureMintERC721TokensMintedWithSignature)
	if err := _SignatureMintERC721.contract.UnpackLog(event, "TokensMintedWithSignature", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
