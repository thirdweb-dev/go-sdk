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

// IRoyaltyMetaData contains all meta data concerning the IRoyalty contract.
var IRoyaltyMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRoyaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoyaltyBps\",\"type\":\"uint256\"}],\"name\":\"DefaultRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyForToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getDefaultRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyInfoForToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyInfoForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IRoyaltyABI is the input ABI used to generate the binding from.
// Deprecated: Use IRoyaltyMetaData.ABI instead.
var IRoyaltyABI = IRoyaltyMetaData.ABI

// IRoyalty is an auto generated Go binding around an Ethereum contract.
type IRoyalty struct {
	IRoyaltyCaller     // Read-only binding to the contract
	IRoyaltyTransactor // Write-only binding to the contract
	IRoyaltyFilterer   // Log filterer for contract events
}

// IRoyaltyCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRoyaltyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoyaltyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRoyaltyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoyaltyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRoyaltyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRoyaltySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRoyaltySession struct {
	Contract     *IRoyalty         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRoyaltyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRoyaltyCallerSession struct {
	Contract *IRoyaltyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IRoyaltyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRoyaltyTransactorSession struct {
	Contract     *IRoyaltyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRoyaltyRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRoyaltyRaw struct {
	Contract *IRoyalty // Generic contract binding to access the raw methods on
}

// IRoyaltyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRoyaltyCallerRaw struct {
	Contract *IRoyaltyCaller // Generic read-only contract binding to access the raw methods on
}

// IRoyaltyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRoyaltyTransactorRaw struct {
	Contract *IRoyaltyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRoyalty creates a new instance of IRoyalty, bound to a specific deployed contract.
func NewIRoyalty(address common.Address, backend bind.ContractBackend) (*IRoyalty, error) {
	contract, err := bindIRoyalty(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRoyalty{IRoyaltyCaller: IRoyaltyCaller{contract: contract}, IRoyaltyTransactor: IRoyaltyTransactor{contract: contract}, IRoyaltyFilterer: IRoyaltyFilterer{contract: contract}}, nil
}

// NewIRoyaltyCaller creates a new read-only instance of IRoyalty, bound to a specific deployed contract.
func NewIRoyaltyCaller(address common.Address, caller bind.ContractCaller) (*IRoyaltyCaller, error) {
	contract, err := bindIRoyalty(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRoyaltyCaller{contract: contract}, nil
}

// NewIRoyaltyTransactor creates a new write-only instance of IRoyalty, bound to a specific deployed contract.
func NewIRoyaltyTransactor(address common.Address, transactor bind.ContractTransactor) (*IRoyaltyTransactor, error) {
	contract, err := bindIRoyalty(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRoyaltyTransactor{contract: contract}, nil
}

// NewIRoyaltyFilterer creates a new log filterer instance of IRoyalty, bound to a specific deployed contract.
func NewIRoyaltyFilterer(address common.Address, filterer bind.ContractFilterer) (*IRoyaltyFilterer, error) {
	contract, err := bindIRoyalty(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRoyaltyFilterer{contract: contract}, nil
}

// bindIRoyalty binds a generic wrapper to an already deployed contract.
func bindIRoyalty(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRoyaltyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoyalty *IRoyaltyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoyalty.Contract.IRoyaltyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoyalty *IRoyaltyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoyalty.Contract.IRoyaltyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoyalty *IRoyaltyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoyalty.Contract.IRoyaltyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRoyalty *IRoyaltyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRoyalty.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRoyalty *IRoyaltyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRoyalty.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRoyalty *IRoyaltyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRoyalty.Contract.contract.Transact(opts, method, params...)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_IRoyalty *IRoyaltyCaller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _IRoyalty.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_IRoyalty *IRoyaltySession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _IRoyalty.Contract.GetDefaultRoyaltyInfo(&_IRoyalty.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_IRoyalty *IRoyaltyCallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _IRoyalty.Contract.GetDefaultRoyaltyInfo(&_IRoyalty.CallOpts)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 tokenId) view returns(address, uint16)
func (_IRoyalty *IRoyaltyCaller) GetRoyaltyInfoForToken(opts *bind.CallOpts, tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _IRoyalty.contract.Call(opts, &out, "getRoyaltyInfoForToken", tokenId)

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 tokenId) view returns(address, uint16)
func (_IRoyalty *IRoyaltySession) GetRoyaltyInfoForToken(tokenId *big.Int) (common.Address, uint16, error) {
	return _IRoyalty.Contract.GetRoyaltyInfoForToken(&_IRoyalty.CallOpts, tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 tokenId) view returns(address, uint16)
func (_IRoyalty *IRoyaltyCallerSession) GetRoyaltyInfoForToken(tokenId *big.Int) (common.Address, uint16, error) {
	return _IRoyalty.Contract.GetRoyaltyInfoForToken(&_IRoyalty.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_IRoyalty *IRoyaltyCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _IRoyalty.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_IRoyalty *IRoyaltySession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _IRoyalty.Contract.RoyaltyInfo(&_IRoyalty.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_IRoyalty *IRoyaltyCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _IRoyalty.Contract.RoyaltyInfo(&_IRoyalty.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRoyalty *IRoyaltyCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IRoyalty.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRoyalty *IRoyaltySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRoyalty.Contract.SupportsInterface(&_IRoyalty.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRoyalty *IRoyaltyCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRoyalty.Contract.SupportsInterface(&_IRoyalty.CallOpts, interfaceId)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_IRoyalty *IRoyaltyTransactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _IRoyalty.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_IRoyalty *IRoyaltySession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _IRoyalty.Contract.SetDefaultRoyaltyInfo(&_IRoyalty.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_IRoyalty *IRoyaltyTransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _IRoyalty.Contract.SetDefaultRoyaltyInfo(&_IRoyalty.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 tokenId, address recipient, uint256 bps) returns()
func (_IRoyalty *IRoyaltyTransactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, tokenId *big.Int, recipient common.Address, bps *big.Int) (*types.Transaction, error) {
	return _IRoyalty.contract.Transact(opts, "setRoyaltyInfoForToken", tokenId, recipient, bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 tokenId, address recipient, uint256 bps) returns()
func (_IRoyalty *IRoyaltySession) SetRoyaltyInfoForToken(tokenId *big.Int, recipient common.Address, bps *big.Int) (*types.Transaction, error) {
	return _IRoyalty.Contract.SetRoyaltyInfoForToken(&_IRoyalty.TransactOpts, tokenId, recipient, bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 tokenId, address recipient, uint256 bps) returns()
func (_IRoyalty *IRoyaltyTransactorSession) SetRoyaltyInfoForToken(tokenId *big.Int, recipient common.Address, bps *big.Int) (*types.Transaction, error) {
	return _IRoyalty.Contract.SetRoyaltyInfoForToken(&_IRoyalty.TransactOpts, tokenId, recipient, bps)
}

// IRoyaltyDefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the IRoyalty contract.
type IRoyaltyDefaultRoyaltyIterator struct {
	Event *IRoyaltyDefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *IRoyaltyDefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoyaltyDefaultRoyalty)
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
		it.Event = new(IRoyaltyDefaultRoyalty)
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
func (it *IRoyaltyDefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoyaltyDefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoyaltyDefaultRoyalty represents a DefaultRoyalty event raised by the IRoyalty contract.
type IRoyaltyDefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_IRoyalty *IRoyaltyFilterer) FilterDefaultRoyalty(opts *bind.FilterOpts, newRoyaltyRecipient []common.Address) (*IRoyaltyDefaultRoyaltyIterator, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _IRoyalty.contract.FilterLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &IRoyaltyDefaultRoyaltyIterator{contract: _IRoyalty.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_IRoyalty *IRoyaltyFilterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *IRoyaltyDefaultRoyalty, newRoyaltyRecipient []common.Address) (event.Subscription, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _IRoyalty.contract.WatchLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoyaltyDefaultRoyalty)
				if err := _IRoyalty.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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

// ParseDefaultRoyalty is a log parse operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_IRoyalty *IRoyaltyFilterer) ParseDefaultRoyalty(log types.Log) (*IRoyaltyDefaultRoyalty, error) {
	event := new(IRoyaltyDefaultRoyalty)
	if err := _IRoyalty.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRoyaltyRoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the IRoyalty contract.
type IRoyaltyRoyaltyForTokenIterator struct {
	Event *IRoyaltyRoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *IRoyaltyRoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRoyaltyRoyaltyForToken)
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
		it.Event = new(IRoyaltyRoyaltyForToken)
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
func (it *IRoyaltyRoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRoyaltyRoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRoyaltyRoyaltyForToken represents a RoyaltyForToken event raised by the IRoyalty contract.
type IRoyaltyRoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_IRoyalty *IRoyaltyFilterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int, royaltyRecipient []common.Address) (*IRoyaltyRoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _IRoyalty.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &IRoyaltyRoyaltyForTokenIterator{contract: _IRoyalty.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_IRoyalty *IRoyaltyFilterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *IRoyaltyRoyaltyForToken, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _IRoyalty.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRoyaltyRoyaltyForToken)
				if err := _IRoyalty.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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

// ParseRoyaltyForToken is a log parse operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_IRoyalty *IRoyaltyFilterer) ParseRoyaltyForToken(log types.Log) (*IRoyaltyRoyaltyForToken, error) {
	event := new(IRoyaltyRoyaltyForToken)
	if err := _IRoyalty.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
