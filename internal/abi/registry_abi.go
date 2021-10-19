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
)

// RegistryMetaData contains all meta data concerning the Registry contract.
var RegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_forwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"defaultFeeBps\",\"type\":\"uint256\"}],\"name\":\"DefaultFeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newDeployer\",\"type\":\"address\"}],\"name\":\"DeployerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"controlAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"controlDeployer\",\"type\":\"address\"}],\"name\":\"NewProtocolControl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"control\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeBps\",\"type\":\"uint256\"}],\"name\":\"ProtocolControlFeeBpsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"TreasuryUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_PROVIDER_FEE_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"deployProtocol\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"contractIControlDeployer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forwarder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolControl\",\"type\":\"address\"}],\"name\":\"getFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getProtocolControl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_deployer\",\"type\":\"address\"}],\"name\":\"getProtocolControlCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFeeBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newDeployer\",\"type\":\"address\"}],\"name\":\"setDeployer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocolControl\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_newFeeBps\",\"type\":\"uint256\"}],\"name\":\"setProtocolControlFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryMetaData.ABI instead.
var RegistryABI = RegistryMetaData.ABI

// Registry is an auto generated Go binding around an Ethereum contract.
type Registry struct {
	RegistryCaller     // Read-only binding to the contract
	RegistryTransactor // Write-only binding to the contract
	RegistryFilterer   // Log filterer for contract events
}

// RegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrySession struct {
	Contract     *Registry         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistryCallerSession struct {
	Contract *RegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistryTransactorSession struct {
	Contract     *RegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistryRaw struct {
	Contract *Registry // Generic contract binding to access the raw methods on
}

// RegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistryCallerRaw struct {
	Contract *RegistryCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistryTransactorRaw struct {
	Contract *RegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistry creates a new instance of Registry, bound to a specific deployed contract.
func NewRegistry(address common.Address, backend bind.ContractBackend) (*Registry, error) {
	contract, err := bindRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registry{RegistryCaller: RegistryCaller{contract: contract}, RegistryTransactor: RegistryTransactor{contract: contract}, RegistryFilterer: RegistryFilterer{contract: contract}}, nil
}

// NewRegistryCaller creates a new read-only instance of Registry, bound to a specific deployed contract.
func NewRegistryCaller(address common.Address, caller bind.ContractCaller) (*RegistryCaller, error) {
	contract, err := bindRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryCaller{contract: contract}, nil
}

// NewRegistryTransactor creates a new write-only instance of Registry, bound to a specific deployed contract.
func NewRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryTransactor, error) {
	contract, err := bindRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryTransactor{contract: contract}, nil
}

// NewRegistryFilterer creates a new log filterer instance of Registry, bound to a specific deployed contract.
func NewRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryFilterer, error) {
	contract, err := bindRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryFilterer{contract: contract}, nil
}

// bindRegistry binds a generic wrapper to an already deployed contract.
func bindRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.RegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.RegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registry *RegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registry *RegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registry *RegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registry.Contract.contract.Transact(opts, method, params...)
}

// MAXPROVIDERFEEBPS is a free data retrieval call binding the contract method 0x5cd42315.
//
// Solidity: function MAX_PROVIDER_FEE_BPS() view returns(uint256)
func (_Registry *RegistryCaller) MAXPROVIDERFEEBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "MAX_PROVIDER_FEE_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPROVIDERFEEBPS is a free data retrieval call binding the contract method 0x5cd42315.
//
// Solidity: function MAX_PROVIDER_FEE_BPS() view returns(uint256)
func (_Registry *RegistrySession) MAXPROVIDERFEEBPS() (*big.Int, error) {
	return _Registry.Contract.MAXPROVIDERFEEBPS(&_Registry.CallOpts)
}

// MAXPROVIDERFEEBPS is a free data retrieval call binding the contract method 0x5cd42315.
//
// Solidity: function MAX_PROVIDER_FEE_BPS() view returns(uint256)
func (_Registry *RegistryCallerSession) MAXPROVIDERFEEBPS() (*big.Int, error) {
	return _Registry.Contract.MAXPROVIDERFEEBPS(&_Registry.CallOpts)
}

// DefaultFeeBps is a free data retrieval call binding the contract method 0xbcae25a4.
//
// Solidity: function defaultFeeBps() view returns(uint256)
func (_Registry *RegistryCaller) DefaultFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "defaultFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultFeeBps is a free data retrieval call binding the contract method 0xbcae25a4.
//
// Solidity: function defaultFeeBps() view returns(uint256)
func (_Registry *RegistrySession) DefaultFeeBps() (*big.Int, error) {
	return _Registry.Contract.DefaultFeeBps(&_Registry.CallOpts)
}

// DefaultFeeBps is a free data retrieval call binding the contract method 0xbcae25a4.
//
// Solidity: function defaultFeeBps() view returns(uint256)
func (_Registry *RegistryCallerSession) DefaultFeeBps() (*big.Int, error) {
	return _Registry.Contract.DefaultFeeBps(&_Registry.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Registry *RegistryCaller) Deployer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "deployer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Registry *RegistrySession) Deployer() (common.Address, error) {
	return _Registry.Contract.Deployer(&_Registry.CallOpts)
}

// Deployer is a free data retrieval call binding the contract method 0xd5f39488.
//
// Solidity: function deployer() view returns(address)
func (_Registry *RegistryCallerSession) Deployer() (common.Address, error) {
	return _Registry.Contract.Deployer(&_Registry.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_Registry *RegistryCaller) Forwarder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "forwarder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_Registry *RegistrySession) Forwarder() (common.Address, error) {
	return _Registry.Contract.Forwarder(&_Registry.CallOpts)
}

// Forwarder is a free data retrieval call binding the contract method 0xf645d4f9.
//
// Solidity: function forwarder() view returns(address)
func (_Registry *RegistryCallerSession) Forwarder() (common.Address, error) {
	return _Registry.Contract.Forwarder(&_Registry.CallOpts)
}

// GetFeeBps is a free data retrieval call binding the contract method 0x73c6f378.
//
// Solidity: function getFeeBps(address protocolControl) view returns(uint256)
func (_Registry *RegistryCaller) GetFeeBps(opts *bind.CallOpts, protocolControl common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getFeeBps", protocolControl)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBps is a free data retrieval call binding the contract method 0x73c6f378.
//
// Solidity: function getFeeBps(address protocolControl) view returns(uint256)
func (_Registry *RegistrySession) GetFeeBps(protocolControl common.Address) (*big.Int, error) {
	return _Registry.Contract.GetFeeBps(&_Registry.CallOpts, protocolControl)
}

// GetFeeBps is a free data retrieval call binding the contract method 0x73c6f378.
//
// Solidity: function getFeeBps(address protocolControl) view returns(uint256)
func (_Registry *RegistryCallerSession) GetFeeBps(protocolControl common.Address) (*big.Int, error) {
	return _Registry.Contract.GetFeeBps(&_Registry.CallOpts, protocolControl)
}

// GetProtocolControl is a free data retrieval call binding the contract method 0xdc62163e.
//
// Solidity: function getProtocolControl(address _deployer, uint256 index) view returns(address)
func (_Registry *RegistryCaller) GetProtocolControl(opts *bind.CallOpts, _deployer common.Address, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getProtocolControl", _deployer, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolControl is a free data retrieval call binding the contract method 0xdc62163e.
//
// Solidity: function getProtocolControl(address _deployer, uint256 index) view returns(address)
func (_Registry *RegistrySession) GetProtocolControl(_deployer common.Address, index *big.Int) (common.Address, error) {
	return _Registry.Contract.GetProtocolControl(&_Registry.CallOpts, _deployer, index)
}

// GetProtocolControl is a free data retrieval call binding the contract method 0xdc62163e.
//
// Solidity: function getProtocolControl(address _deployer, uint256 index) view returns(address)
func (_Registry *RegistryCallerSession) GetProtocolControl(_deployer common.Address, index *big.Int) (common.Address, error) {
	return _Registry.Contract.GetProtocolControl(&_Registry.CallOpts, _deployer, index)
}

// GetProtocolControlCount is a free data retrieval call binding the contract method 0xcd7e89b1.
//
// Solidity: function getProtocolControlCount(address _deployer) view returns(uint256)
func (_Registry *RegistryCaller) GetProtocolControlCount(opts *bind.CallOpts, _deployer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "getProtocolControlCount", _deployer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolControlCount is a free data retrieval call binding the contract method 0xcd7e89b1.
//
// Solidity: function getProtocolControlCount(address _deployer) view returns(uint256)
func (_Registry *RegistrySession) GetProtocolControlCount(_deployer common.Address) (*big.Int, error) {
	return _Registry.Contract.GetProtocolControlCount(&_Registry.CallOpts, _deployer)
}

// GetProtocolControlCount is a free data retrieval call binding the contract method 0xcd7e89b1.
//
// Solidity: function getProtocolControlCount(address _deployer) view returns(uint256)
func (_Registry *RegistryCallerSession) GetProtocolControlCount(_deployer common.Address) (*big.Int, error) {
	return _Registry.Contract.GetProtocolControlCount(&_Registry.CallOpts, _deployer)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistrySession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registry *RegistryCallerSession) Owner() (common.Address, error) {
	return _Registry.Contract.Owner(&_Registry.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Registry *RegistryCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registry.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Registry *RegistrySession) Treasury() (common.Address, error) {
	return _Registry.Contract.Treasury(&_Registry.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Registry *RegistryCallerSession) Treasury() (common.Address, error) {
	return _Registry.Contract.Treasury(&_Registry.CallOpts)
}

// DeployProtocol is a paid mutator transaction binding the contract method 0xd566464b.
//
// Solidity: function deployProtocol(string uri) returns()
func (_Registry *RegistryTransactor) DeployProtocol(opts *bind.TransactOpts, uri string) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "deployProtocol", uri)
}

// DeployProtocol is a paid mutator transaction binding the contract method 0xd566464b.
//
// Solidity: function deployProtocol(string uri) returns()
func (_Registry *RegistrySession) DeployProtocol(uri string) (*types.Transaction, error) {
	return _Registry.Contract.DeployProtocol(&_Registry.TransactOpts, uri)
}

// DeployProtocol is a paid mutator transaction binding the contract method 0xd566464b.
//
// Solidity: function deployProtocol(string uri) returns()
func (_Registry *RegistryTransactorSession) DeployProtocol(uri string) (*types.Transaction, error) {
	return _Registry.Contract.DeployProtocol(&_Registry.TransactOpts, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registry *RegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registry.Contract.RenounceOwnership(&_Registry.TransactOpts)
}

// SetDefaultFeeBps is a paid mutator transaction binding the contract method 0x4373a286.
//
// Solidity: function setDefaultFeeBps(uint256 _newFeeBps) returns()
func (_Registry *RegistryTransactor) SetDefaultFeeBps(opts *bind.TransactOpts, _newFeeBps *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setDefaultFeeBps", _newFeeBps)
}

// SetDefaultFeeBps is a paid mutator transaction binding the contract method 0x4373a286.
//
// Solidity: function setDefaultFeeBps(uint256 _newFeeBps) returns()
func (_Registry *RegistrySession) SetDefaultFeeBps(_newFeeBps *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetDefaultFeeBps(&_Registry.TransactOpts, _newFeeBps)
}

// SetDefaultFeeBps is a paid mutator transaction binding the contract method 0x4373a286.
//
// Solidity: function setDefaultFeeBps(uint256 _newFeeBps) returns()
func (_Registry *RegistryTransactorSession) SetDefaultFeeBps(_newFeeBps *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetDefaultFeeBps(&_Registry.TransactOpts, _newFeeBps)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _newDeployer) returns()
func (_Registry *RegistryTransactor) SetDeployer(opts *bind.TransactOpts, _newDeployer common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setDeployer", _newDeployer)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _newDeployer) returns()
func (_Registry *RegistrySession) SetDeployer(_newDeployer common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetDeployer(&_Registry.TransactOpts, _newDeployer)
}

// SetDeployer is a paid mutator transaction binding the contract method 0x96214735.
//
// Solidity: function setDeployer(address _newDeployer) returns()
func (_Registry *RegistryTransactorSession) SetDeployer(_newDeployer common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetDeployer(&_Registry.TransactOpts, _newDeployer)
}

// SetProtocolControlFeeBps is a paid mutator transaction binding the contract method 0x7c160011.
//
// Solidity: function setProtocolControlFeeBps(address protocolControl, uint256 _newFeeBps) returns()
func (_Registry *RegistryTransactor) SetProtocolControlFeeBps(opts *bind.TransactOpts, protocolControl common.Address, _newFeeBps *big.Int) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setProtocolControlFeeBps", protocolControl, _newFeeBps)
}

// SetProtocolControlFeeBps is a paid mutator transaction binding the contract method 0x7c160011.
//
// Solidity: function setProtocolControlFeeBps(address protocolControl, uint256 _newFeeBps) returns()
func (_Registry *RegistrySession) SetProtocolControlFeeBps(protocolControl common.Address, _newFeeBps *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetProtocolControlFeeBps(&_Registry.TransactOpts, protocolControl, _newFeeBps)
}

// SetProtocolControlFeeBps is a paid mutator transaction binding the contract method 0x7c160011.
//
// Solidity: function setProtocolControlFeeBps(address protocolControl, uint256 _newFeeBps) returns()
func (_Registry *RegistryTransactorSession) SetProtocolControlFeeBps(protocolControl common.Address, _newFeeBps *big.Int) (*types.Transaction, error) {
	return _Registry.Contract.SetProtocolControlFeeBps(&_Registry.TransactOpts, protocolControl, _newFeeBps)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Registry *RegistryTransactor) SetTreasury(opts *bind.TransactOpts, _newTreasury common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "setTreasury", _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Registry *RegistrySession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetTreasury(&_Registry.TransactOpts, _newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address _newTreasury) returns()
func (_Registry *RegistryTransactorSession) SetTreasury(_newTreasury common.Address) (*types.Transaction, error) {
	return _Registry.Contract.SetTreasury(&_Registry.TransactOpts, _newTreasury)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registry *RegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registry.Contract.TransferOwnership(&_Registry.TransactOpts, newOwner)
}

// RegistryDefaultFeeBpsUpdatedIterator is returned from FilterDefaultFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for DefaultFeeBpsUpdated events raised by the Registry contract.
type RegistryDefaultFeeBpsUpdatedIterator struct {
	Event *RegistryDefaultFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryDefaultFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDefaultFeeBpsUpdated)
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
		it.Event = new(RegistryDefaultFeeBpsUpdated)
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
func (it *RegistryDefaultFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDefaultFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDefaultFeeBpsUpdated represents a DefaultFeeBpsUpdated event raised by the Registry contract.
type RegistryDefaultFeeBpsUpdated struct {
	DefaultFeeBps *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDefaultFeeBpsUpdated is a free log retrieval operation binding the contract event 0xaf07358220f107097f5c9bc5cacec9de23c6bd67d61633e5abaa3aa698d94fcf.
//
// Solidity: event DefaultFeeBpsUpdated(uint256 defaultFeeBps)
func (_Registry *RegistryFilterer) FilterDefaultFeeBpsUpdated(opts *bind.FilterOpts) (*RegistryDefaultFeeBpsUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "DefaultFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryDefaultFeeBpsUpdatedIterator{contract: _Registry.contract, event: "DefaultFeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchDefaultFeeBpsUpdated is a free log subscription operation binding the contract event 0xaf07358220f107097f5c9bc5cacec9de23c6bd67d61633e5abaa3aa698d94fcf.
//
// Solidity: event DefaultFeeBpsUpdated(uint256 defaultFeeBps)
func (_Registry *RegistryFilterer) WatchDefaultFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *RegistryDefaultFeeBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "DefaultFeeBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDefaultFeeBpsUpdated)
				if err := _Registry.contract.UnpackLog(event, "DefaultFeeBpsUpdated", log); err != nil {
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

// ParseDefaultFeeBpsUpdated is a log parse operation binding the contract event 0xaf07358220f107097f5c9bc5cacec9de23c6bd67d61633e5abaa3aa698d94fcf.
//
// Solidity: event DefaultFeeBpsUpdated(uint256 defaultFeeBps)
func (_Registry *RegistryFilterer) ParseDefaultFeeBpsUpdated(log types.Log) (*RegistryDefaultFeeBpsUpdated, error) {
	event := new(RegistryDefaultFeeBpsUpdated)
	if err := _Registry.contract.UnpackLog(event, "DefaultFeeBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryDeployerUpdatedIterator is returned from FilterDeployerUpdated and is used to iterate over the raw logs and unpacked data for DeployerUpdated events raised by the Registry contract.
type RegistryDeployerUpdatedIterator struct {
	Event *RegistryDeployerUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryDeployerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryDeployerUpdated)
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
		it.Event = new(RegistryDeployerUpdated)
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
func (it *RegistryDeployerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryDeployerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryDeployerUpdated represents a DeployerUpdated event raised by the Registry contract.
type RegistryDeployerUpdated struct {
	NewDeployer common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeployerUpdated is a free log retrieval operation binding the contract event 0x6db6dcdd05f1728263d8f644adcb07da9d18505aa9b2e33360b2715a878a711e.
//
// Solidity: event DeployerUpdated(address newDeployer)
func (_Registry *RegistryFilterer) FilterDeployerUpdated(opts *bind.FilterOpts) (*RegistryDeployerUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "DeployerUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryDeployerUpdatedIterator{contract: _Registry.contract, event: "DeployerUpdated", logs: logs, sub: sub}, nil
}

// WatchDeployerUpdated is a free log subscription operation binding the contract event 0x6db6dcdd05f1728263d8f644adcb07da9d18505aa9b2e33360b2715a878a711e.
//
// Solidity: event DeployerUpdated(address newDeployer)
func (_Registry *RegistryFilterer) WatchDeployerUpdated(opts *bind.WatchOpts, sink chan<- *RegistryDeployerUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "DeployerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryDeployerUpdated)
				if err := _Registry.contract.UnpackLog(event, "DeployerUpdated", log); err != nil {
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

// ParseDeployerUpdated is a log parse operation binding the contract event 0x6db6dcdd05f1728263d8f644adcb07da9d18505aa9b2e33360b2715a878a711e.
//
// Solidity: event DeployerUpdated(address newDeployer)
func (_Registry *RegistryFilterer) ParseDeployerUpdated(log types.Log) (*RegistryDeployerUpdated, error) {
	event := new(RegistryDeployerUpdated)
	if err := _Registry.contract.UnpackLog(event, "DeployerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryNewProtocolControlIterator is returned from FilterNewProtocolControl and is used to iterate over the raw logs and unpacked data for NewProtocolControl events raised by the Registry contract.
type RegistryNewProtocolControlIterator struct {
	Event *RegistryNewProtocolControl // Event containing the contract specifics and raw log

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
func (it *RegistryNewProtocolControlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryNewProtocolControl)
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
		it.Event = new(RegistryNewProtocolControl)
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
func (it *RegistryNewProtocolControlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryNewProtocolControlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryNewProtocolControl represents a NewProtocolControl event raised by the Registry contract.
type RegistryNewProtocolControl struct {
	Deployer        common.Address
	Version         *big.Int
	ControlAddress  common.Address
	ControlDeployer common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewProtocolControl is a free log retrieval operation binding the contract event 0x87f4df5222ed82bbc7706c066c5d45f3f5682016cb20d642c5491b7b411340b4.
//
// Solidity: event NewProtocolControl(address indexed deployer, uint256 indexed version, address indexed controlAddress, address controlDeployer)
func (_Registry *RegistryFilterer) FilterNewProtocolControl(opts *bind.FilterOpts, deployer []common.Address, version []*big.Int, controlAddress []common.Address) (*RegistryNewProtocolControlIterator, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var controlAddressRule []interface{}
	for _, controlAddressItem := range controlAddress {
		controlAddressRule = append(controlAddressRule, controlAddressItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "NewProtocolControl", deployerRule, versionRule, controlAddressRule)
	if err != nil {
		return nil, err
	}
	return &RegistryNewProtocolControlIterator{contract: _Registry.contract, event: "NewProtocolControl", logs: logs, sub: sub}, nil
}

// WatchNewProtocolControl is a free log subscription operation binding the contract event 0x87f4df5222ed82bbc7706c066c5d45f3f5682016cb20d642c5491b7b411340b4.
//
// Solidity: event NewProtocolControl(address indexed deployer, uint256 indexed version, address indexed controlAddress, address controlDeployer)
func (_Registry *RegistryFilterer) WatchNewProtocolControl(opts *bind.WatchOpts, sink chan<- *RegistryNewProtocolControl, deployer []common.Address, version []*big.Int, controlAddress []common.Address) (event.Subscription, error) {

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var controlAddressRule []interface{}
	for _, controlAddressItem := range controlAddress {
		controlAddressRule = append(controlAddressRule, controlAddressItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "NewProtocolControl", deployerRule, versionRule, controlAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryNewProtocolControl)
				if err := _Registry.contract.UnpackLog(event, "NewProtocolControl", log); err != nil {
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

// ParseNewProtocolControl is a log parse operation binding the contract event 0x87f4df5222ed82bbc7706c066c5d45f3f5682016cb20d642c5491b7b411340b4.
//
// Solidity: event NewProtocolControl(address indexed deployer, uint256 indexed version, address indexed controlAddress, address controlDeployer)
func (_Registry *RegistryFilterer) lControl(log types.Log) (*RegistryNewProtocolControl, error) {
	event := new(RegistryNewProtocolControl)
	if err := _Registry.contract.UnpackLog(event, "NewProtocolControl", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registry contract.
type RegistryOwnershipTransferredIterator struct {
	Event *RegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryOwnershipTransferred)
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
		it.Event = new(RegistryOwnershipTransferred)
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
func (it *RegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryOwnershipTransferred represents a OwnershipTransferred event raised by the Registry contract.
type RegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryOwnershipTransferredIterator{contract: _Registry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryOwnershipTransferred)
				if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registry *RegistryFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryOwnershipTransferred, error) {
	event := new(RegistryOwnershipTransferred)
	if err := _Registry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryProtocolControlFeeBpsUpdatedIterator is returned from FilterProtocolControlFeeBpsUpdated and is used to iterate over the raw logs and unpacked data for ProtocolControlFeeBpsUpdated events raised by the Registry contract.
type RegistryProtocolControlFeeBpsUpdatedIterator struct {
	Event *RegistryProtocolControlFeeBpsUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryProtocolControlFeeBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryProtocolControlFeeBpsUpdated)
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
		it.Event = new(RegistryProtocolControlFeeBpsUpdated)
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
func (it *RegistryProtocolControlFeeBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryProtocolControlFeeBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryProtocolControlFeeBpsUpdated represents a ProtocolControlFeeBpsUpdated event raised by the Registry contract.
type RegistryProtocolControlFeeBpsUpdated struct {
	Control common.Address
	FeeBps  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProtocolControlFeeBpsUpdated is a free log retrieval operation binding the contract event 0x3f3f5381919d082b3f6431a55b5affd99c52c10d6b16a850a6c75377c17285a7.
//
// Solidity: event ProtocolControlFeeBpsUpdated(address indexed control, uint256 feeBps)
func (_Registry *RegistryFilterer) FilterProtocolControlFeeBpsUpdated(opts *bind.FilterOpts, control []common.Address) (*RegistryProtocolControlFeeBpsUpdatedIterator, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}

	logs, sub, err := _Registry.contract.FilterLogs(opts, "ProtocolControlFeeBpsUpdated", controlRule)
	if err != nil {
		return nil, err
	}
	return &RegistryProtocolControlFeeBpsUpdatedIterator{contract: _Registry.contract, event: "ProtocolControlFeeBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolControlFeeBpsUpdated is a free log subscription operation binding the contract event 0x3f3f5381919d082b3f6431a55b5affd99c52c10d6b16a850a6c75377c17285a7.
//
// Solidity: event ProtocolControlFeeBpsUpdated(address indexed control, uint256 feeBps)
func (_Registry *RegistryFilterer) WatchProtocolControlFeeBpsUpdated(opts *bind.WatchOpts, sink chan<- *RegistryProtocolControlFeeBpsUpdated, control []common.Address) (event.Subscription, error) {

	var controlRule []interface{}
	for _, controlItem := range control {
		controlRule = append(controlRule, controlItem)
	}

	logs, sub, err := _Registry.contract.WatchLogs(opts, "ProtocolControlFeeBpsUpdated", controlRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryProtocolControlFeeBpsUpdated)
				if err := _Registry.contract.UnpackLog(event, "ProtocolControlFeeBpsUpdated", log); err != nil {
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

// ParseProtocolControlFeeBpsUpdated is a log parse operation binding the contract event 0x3f3f5381919d082b3f6431a55b5affd99c52c10d6b16a850a6c75377c17285a7.
//
// Solidity: event ProtocolControlFeeBpsUpdated(address indexed control, uint256 feeBps)
func (_Registry *RegistryFilterer) ParseProtocolControlFeeBpsUpdated(log types.Log) (*RegistryProtocolControlFeeBpsUpdated, error) {
	event := new(RegistryProtocolControlFeeBpsUpdated)
	if err := _Registry.contract.UnpackLog(event, "ProtocolControlFeeBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryTreasuryUpdatedIterator is returned from FilterTreasuryUpdated and is used to iterate over the raw logs and unpacked data for TreasuryUpdated events raised by the Registry contract.
type RegistryTreasuryUpdatedIterator struct {
	Event *RegistryTreasuryUpdated // Event containing the contract specifics and raw log

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
func (it *RegistryTreasuryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryTreasuryUpdated)
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
		it.Event = new(RegistryTreasuryUpdated)
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
func (it *RegistryTreasuryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryTreasuryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryTreasuryUpdated represents a TreasuryUpdated event raised by the Registry contract.
type RegistryTreasuryUpdated struct {
	NewTreasury common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTreasuryUpdated is a free log retrieval operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_Registry *RegistryFilterer) FilterTreasuryUpdated(opts *bind.FilterOpts) (*RegistryTreasuryUpdatedIterator, error) {

	logs, sub, err := _Registry.contract.FilterLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistryTreasuryUpdatedIterator{contract: _Registry.contract, event: "TreasuryUpdated", logs: logs, sub: sub}, nil
}

// WatchTreasuryUpdated is a free log subscription operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_Registry *RegistryFilterer) WatchTreasuryUpdated(opts *bind.WatchOpts, sink chan<- *RegistryTreasuryUpdated) (event.Subscription, error) {

	logs, sub, err := _Registry.contract.WatchLogs(opts, "TreasuryUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryTreasuryUpdated)
				if err := _Registry.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
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

// ParseTreasuryUpdated is a log parse operation binding the contract event 0x7dae230f18360d76a040c81f050aa14eb9d6dc7901b20fc5d855e2a20fe814d1.
//
// Solidity: event TreasuryUpdated(address newTreasury)
func (_Registry *RegistryFilterer) ParseTreasuryUpdated(log types.Log) (*RegistryTreasuryUpdated, error) {
	event := new(RegistryTreasuryUpdated)
	if err := _Registry.contract.UnpackLog(event, "TreasuryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
