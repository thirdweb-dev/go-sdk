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

// TWFactoryMetaData contains all meta data concerning the TWFactory contract.
var TWFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"contractType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ImplementationAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"name\":\"ImplementationApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"deployer\",\"type\":\"address\"}],\"name\":\"ProxyDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FACTORY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"}],\"name\":\"addImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_toApprove\",\"type\":\"bool\"}],\"name\":\"approveImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"currentVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"deployProxy\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deployProxyByImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"deployedProxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_salt\",\"type\":\"bytes32\"}],\"name\":\"deployProxyDeterministic\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"deployer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_version\",\"type\":\"uint256\"}],\"name\":\"getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_type\",\"type\":\"bytes32\"}],\"name\":\"getLatestImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractTWRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TWFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use TWFactoryMetaData.ABI instead.
var TWFactoryABI = TWFactoryMetaData.ABI

// TWFactory is an auto generated Go binding around an Ethereum contract.
type TWFactory struct {
	TWFactoryCaller     // Read-only binding to the contract
	TWFactoryTransactor // Write-only binding to the contract
	TWFactoryFilterer   // Log filterer for contract events
}

// TWFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TWFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TWFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TWFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TWFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TWFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TWFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TWFactorySession struct {
	Contract     *TWFactory        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TWFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TWFactoryCallerSession struct {
	Contract *TWFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TWFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TWFactoryTransactorSession struct {
	Contract     *TWFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TWFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TWFactoryRaw struct {
	Contract *TWFactory // Generic contract binding to access the raw methods on
}

// TWFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TWFactoryCallerRaw struct {
	Contract *TWFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TWFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TWFactoryTransactorRaw struct {
	Contract *TWFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTWFactory creates a new instance of TWFactory, bound to a specific deployed contract.
func NewTWFactory(address common.Address, backend bind.ContractBackend) (*TWFactory, error) {
	contract, err := bindTWFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TWFactory{TWFactoryCaller: TWFactoryCaller{contract: contract}, TWFactoryTransactor: TWFactoryTransactor{contract: contract}, TWFactoryFilterer: TWFactoryFilterer{contract: contract}}, nil
}

// NewTWFactoryCaller creates a new read-only instance of TWFactory, bound to a specific deployed contract.
func NewTWFactoryCaller(address common.Address, caller bind.ContractCaller) (*TWFactoryCaller, error) {
	contract, err := bindTWFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TWFactoryCaller{contract: contract}, nil
}

// NewTWFactoryTransactor creates a new write-only instance of TWFactory, bound to a specific deployed contract.
func NewTWFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TWFactoryTransactor, error) {
	contract, err := bindTWFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TWFactoryTransactor{contract: contract}, nil
}

// NewTWFactoryFilterer creates a new log filterer instance of TWFactory, bound to a specific deployed contract.
func NewTWFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TWFactoryFilterer, error) {
	contract, err := bindTWFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TWFactoryFilterer{contract: contract}, nil
}

// bindTWFactory binds a generic wrapper to an already deployed contract.
func bindTWFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TWFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TWFactory *TWFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TWFactory.Contract.TWFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TWFactory *TWFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TWFactory.Contract.TWFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TWFactory *TWFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TWFactory.Contract.TWFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TWFactory *TWFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TWFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TWFactory *TWFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TWFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TWFactory *TWFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TWFactory.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TWFactory *TWFactoryCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TWFactory *TWFactorySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TWFactory.Contract.DEFAULTADMINROLE(&_TWFactory.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TWFactory *TWFactoryCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TWFactory.Contract.DEFAULTADMINROLE(&_TWFactory.CallOpts)
}

// FACTORYROLE is a free data retrieval call binding the contract method 0x04a0fb17.
//
// Solidity: function FACTORY_ROLE() view returns(bytes32)
func (_TWFactory *TWFactoryCaller) FACTORYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "FACTORY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FACTORYROLE is a free data retrieval call binding the contract method 0x04a0fb17.
//
// Solidity: function FACTORY_ROLE() view returns(bytes32)
func (_TWFactory *TWFactorySession) FACTORYROLE() ([32]byte, error) {
	return _TWFactory.Contract.FACTORYROLE(&_TWFactory.CallOpts)
}

// FACTORYROLE is a free data retrieval call binding the contract method 0x04a0fb17.
//
// Solidity: function FACTORY_ROLE() view returns(bytes32)
func (_TWFactory *TWFactoryCallerSession) FACTORYROLE() ([32]byte, error) {
	return _TWFactory.Contract.FACTORYROLE(&_TWFactory.CallOpts)
}

// Approval is a free data retrieval call binding the contract method 0x9430b496.
//
// Solidity: function approval(address ) view returns(bool)
func (_TWFactory *TWFactoryCaller) Approval(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "approval", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approval is a free data retrieval call binding the contract method 0x9430b496.
//
// Solidity: function approval(address ) view returns(bool)
func (_TWFactory *TWFactorySession) Approval(arg0 common.Address) (bool, error) {
	return _TWFactory.Contract.Approval(&_TWFactory.CallOpts, arg0)
}

// Approval is a free data retrieval call binding the contract method 0x9430b496.
//
// Solidity: function approval(address ) view returns(bool)
func (_TWFactory *TWFactoryCallerSession) Approval(arg0 common.Address) (bool, error) {
	return _TWFactory.Contract.Approval(&_TWFactory.CallOpts, arg0)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x3b426d3f.
//
// Solidity: function currentVersion(bytes32 ) view returns(uint256)
func (_TWFactory *TWFactoryCaller) CurrentVersion(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "currentVersion", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentVersion is a free data retrieval call binding the contract method 0x3b426d3f.
//
// Solidity: function currentVersion(bytes32 ) view returns(uint256)
func (_TWFactory *TWFactorySession) CurrentVersion(arg0 [32]byte) (*big.Int, error) {
	return _TWFactory.Contract.CurrentVersion(&_TWFactory.CallOpts, arg0)
}

// CurrentVersion is a free data retrieval call binding the contract method 0x3b426d3f.
//
// Solidity: function currentVersion(bytes32 ) view returns(uint256)
func (_TWFactory *TWFactoryCallerSession) CurrentVersion(arg0 [32]byte) (*big.Int, error) {
	return _TWFactory.Contract.CurrentVersion(&_TWFactory.CallOpts, arg0)
}

// Deployer is a free data retrieval call binding the contract method 0xb9caf9d9.
//
// Solidity: function deployer(address ) view returns(address)
func (_TWFactory *TWFactoryCaller) Deployer(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "deployer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deployer is a free data retrieval call binding the contract method 0xb9caf9d9.
//
// Solidity: function deployer(address ) view returns(address)
func (_TWFactory *TWFactorySession) Deployer(arg0 common.Address) (common.Address, error) {
	return _TWFactory.Contract.Deployer(&_TWFactory.CallOpts, arg0)
}

// Deployer is a free data retrieval call binding the contract method 0xb9caf9d9.
//
// Solidity: function deployer(address ) view returns(address)
func (_TWFactory *TWFactoryCallerSession) Deployer(arg0 common.Address) (common.Address, error) {
	return _TWFactory.Contract.Deployer(&_TWFactory.CallOpts, arg0)
}

// GetImplementation is a free data retrieval call binding the contract method 0xdd47595a.
//
// Solidity: function getImplementation(bytes32 _type, uint256 _version) view returns(address)
func (_TWFactory *TWFactoryCaller) GetImplementation(opts *bind.CallOpts, _type [32]byte, _version *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "getImplementation", _type, _version)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xdd47595a.
//
// Solidity: function getImplementation(bytes32 _type, uint256 _version) view returns(address)
func (_TWFactory *TWFactorySession) GetImplementation(_type [32]byte, _version *big.Int) (common.Address, error) {
	return _TWFactory.Contract.GetImplementation(&_TWFactory.CallOpts, _type, _version)
}

// GetImplementation is a free data retrieval call binding the contract method 0xdd47595a.
//
// Solidity: function getImplementation(bytes32 _type, uint256 _version) view returns(address)
func (_TWFactory *TWFactoryCallerSession) GetImplementation(_type [32]byte, _version *big.Int) (common.Address, error) {
	return _TWFactory.Contract.GetImplementation(&_TWFactory.CallOpts, _type, _version)
}

// GetLatestImplementation is a free data retrieval call binding the contract method 0x44ab6680.
//
// Solidity: function getLatestImplementation(bytes32 _type) view returns(address)
func (_TWFactory *TWFactoryCaller) GetLatestImplementation(opts *bind.CallOpts, _type [32]byte) (common.Address, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "getLatestImplementation", _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLatestImplementation is a free data retrieval call binding the contract method 0x44ab6680.
//
// Solidity: function getLatestImplementation(bytes32 _type) view returns(address)
func (_TWFactory *TWFactorySession) GetLatestImplementation(_type [32]byte) (common.Address, error) {
	return _TWFactory.Contract.GetLatestImplementation(&_TWFactory.CallOpts, _type)
}

// GetLatestImplementation is a free data retrieval call binding the contract method 0x44ab6680.
//
// Solidity: function getLatestImplementation(bytes32 _type) view returns(address)
func (_TWFactory *TWFactoryCallerSession) GetLatestImplementation(_type [32]byte) (common.Address, error) {
	return _TWFactory.Contract.GetLatestImplementation(&_TWFactory.CallOpts, _type)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TWFactory *TWFactoryCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TWFactory *TWFactorySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TWFactory.Contract.GetRoleAdmin(&_TWFactory.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TWFactory *TWFactoryCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TWFactory.Contract.GetRoleAdmin(&_TWFactory.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TWFactory *TWFactoryCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TWFactory *TWFactorySession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TWFactory.Contract.GetRoleMember(&_TWFactory.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_TWFactory *TWFactoryCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _TWFactory.Contract.GetRoleMember(&_TWFactory.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TWFactory *TWFactoryCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TWFactory *TWFactorySession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TWFactory.Contract.GetRoleMemberCount(&_TWFactory.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_TWFactory *TWFactoryCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _TWFactory.Contract.GetRoleMemberCount(&_TWFactory.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TWFactory *TWFactoryCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TWFactory *TWFactorySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TWFactory.Contract.HasRole(&_TWFactory.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TWFactory *TWFactoryCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TWFactory.Contract.HasRole(&_TWFactory.CallOpts, role, account)
}

// Implementation is a free data retrieval call binding the contract method 0xe92016a4.
//
// Solidity: function implementation(bytes32 , uint256 ) view returns(address)
func (_TWFactory *TWFactoryCaller) Implementation(opts *bind.CallOpts, arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "implementation", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0xe92016a4.
//
// Solidity: function implementation(bytes32 , uint256 ) view returns(address)
func (_TWFactory *TWFactorySession) Implementation(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _TWFactory.Contract.Implementation(&_TWFactory.CallOpts, arg0, arg1)
}

// Implementation is a free data retrieval call binding the contract method 0xe92016a4.
//
// Solidity: function implementation(bytes32 , uint256 ) view returns(address)
func (_TWFactory *TWFactoryCallerSession) Implementation(arg0 [32]byte, arg1 *big.Int) (common.Address, error) {
	return _TWFactory.Contract.Implementation(&_TWFactory.CallOpts, arg0, arg1)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TWFactory *TWFactoryCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TWFactory *TWFactorySession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TWFactory.Contract.IsTrustedForwarder(&_TWFactory.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_TWFactory *TWFactoryCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _TWFactory.Contract.IsTrustedForwarder(&_TWFactory.CallOpts, forwarder)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TWFactory *TWFactoryCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TWFactory *TWFactorySession) Registry() (common.Address, error) {
	return _TWFactory.Contract.Registry(&_TWFactory.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_TWFactory *TWFactoryCallerSession) Registry() (common.Address, error) {
	return _TWFactory.Contract.Registry(&_TWFactory.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TWFactory *TWFactoryCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TWFactory.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TWFactory *TWFactorySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TWFactory.Contract.SupportsInterface(&_TWFactory.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TWFactory *TWFactoryCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TWFactory.Contract.SupportsInterface(&_TWFactory.CallOpts, interfaceId)
}

// AddImplementation is a paid mutator transaction binding the contract method 0xc6e2a400.
//
// Solidity: function addImplementation(address _implementation) returns()
func (_TWFactory *TWFactoryTransactor) AddImplementation(opts *bind.TransactOpts, _implementation common.Address) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "addImplementation", _implementation)
}

// AddImplementation is a paid mutator transaction binding the contract method 0xc6e2a400.
//
// Solidity: function addImplementation(address _implementation) returns()
func (_TWFactory *TWFactorySession) AddImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.AddImplementation(&_TWFactory.TransactOpts, _implementation)
}

// AddImplementation is a paid mutator transaction binding the contract method 0xc6e2a400.
//
// Solidity: function addImplementation(address _implementation) returns()
func (_TWFactory *TWFactoryTransactorSession) AddImplementation(_implementation common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.AddImplementation(&_TWFactory.TransactOpts, _implementation)
}

// ApproveImplementation is a paid mutator transaction binding the contract method 0x56fb0958.
//
// Solidity: function approveImplementation(address _implementation, bool _toApprove) returns()
func (_TWFactory *TWFactoryTransactor) ApproveImplementation(opts *bind.TransactOpts, _implementation common.Address, _toApprove bool) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "approveImplementation", _implementation, _toApprove)
}

// ApproveImplementation is a paid mutator transaction binding the contract method 0x56fb0958.
//
// Solidity: function approveImplementation(address _implementation, bool _toApprove) returns()
func (_TWFactory *TWFactorySession) ApproveImplementation(_implementation common.Address, _toApprove bool) (*types.Transaction, error) {
	return _TWFactory.Contract.ApproveImplementation(&_TWFactory.TransactOpts, _implementation, _toApprove)
}

// ApproveImplementation is a paid mutator transaction binding the contract method 0x56fb0958.
//
// Solidity: function approveImplementation(address _implementation, bool _toApprove) returns()
func (_TWFactory *TWFactoryTransactorSession) ApproveImplementation(_implementation common.Address, _toApprove bool) (*types.Transaction, error) {
	return _TWFactory.Contract.ApproveImplementation(&_TWFactory.TransactOpts, _implementation, _toApprove)
}

// DeployProxy is a paid mutator transaction binding the contract method 0xec54d72f.
//
// Solidity: function deployProxy(bytes32 _type, bytes _data) returns(address)
func (_TWFactory *TWFactoryTransactor) DeployProxy(opts *bind.TransactOpts, _type [32]byte, _data []byte) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "deployProxy", _type, _data)
}

// DeployProxy is a paid mutator transaction binding the contract method 0xec54d72f.
//
// Solidity: function deployProxy(bytes32 _type, bytes _data) returns(address)
func (_TWFactory *TWFactorySession) DeployProxy(_type [32]byte, _data []byte) (*types.Transaction, error) {
	return _TWFactory.Contract.DeployProxy(&_TWFactory.TransactOpts, _type, _data)
}

// DeployProxy is a paid mutator transaction binding the contract method 0xec54d72f.
//
// Solidity: function deployProxy(bytes32 _type, bytes _data) returns(address)
func (_TWFactory *TWFactoryTransactorSession) DeployProxy(_type [32]byte, _data []byte) (*types.Transaction, error) {
	return _TWFactory.Contract.DeployProxy(&_TWFactory.TransactOpts, _type, _data)
}

// DeployProxyByImplementation is a paid mutator transaction binding the contract method 0x11b804ab.
//
// Solidity: function deployProxyByImplementation(address _implementation, bytes _data, bytes32 _salt) returns(address deployedProxy)
func (_TWFactory *TWFactoryTransactor) DeployProxyByImplementation(opts *bind.TransactOpts, _implementation common.Address, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "deployProxyByImplementation", _implementation, _data, _salt)
}

// DeployProxyByImplementation is a paid mutator transaction binding the contract method 0x11b804ab.
//
// Solidity: function deployProxyByImplementation(address _implementation, bytes _data, bytes32 _salt) returns(address deployedProxy)
func (_TWFactory *TWFactorySession) DeployProxyByImplementation(_implementation common.Address, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWFactory.Contract.DeployProxyByImplementation(&_TWFactory.TransactOpts, _implementation, _data, _salt)
}

// DeployProxyByImplementation is a paid mutator transaction binding the contract method 0x11b804ab.
//
// Solidity: function deployProxyByImplementation(address _implementation, bytes _data, bytes32 _salt) returns(address deployedProxy)
func (_TWFactory *TWFactoryTransactorSession) DeployProxyByImplementation(_implementation common.Address, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWFactory.Contract.DeployProxyByImplementation(&_TWFactory.TransactOpts, _implementation, _data, _salt)
}

// DeployProxyDeterministic is a paid mutator transaction binding the contract method 0x1e5e1e99.
//
// Solidity: function deployProxyDeterministic(bytes32 _type, bytes _data, bytes32 _salt) returns(address)
func (_TWFactory *TWFactoryTransactor) DeployProxyDeterministic(opts *bind.TransactOpts, _type [32]byte, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "deployProxyDeterministic", _type, _data, _salt)
}

// DeployProxyDeterministic is a paid mutator transaction binding the contract method 0x1e5e1e99.
//
// Solidity: function deployProxyDeterministic(bytes32 _type, bytes _data, bytes32 _salt) returns(address)
func (_TWFactory *TWFactorySession) DeployProxyDeterministic(_type [32]byte, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWFactory.Contract.DeployProxyDeterministic(&_TWFactory.TransactOpts, _type, _data, _salt)
}

// DeployProxyDeterministic is a paid mutator transaction binding the contract method 0x1e5e1e99.
//
// Solidity: function deployProxyDeterministic(bytes32 _type, bytes _data, bytes32 _salt) returns(address)
func (_TWFactory *TWFactoryTransactorSession) DeployProxyDeterministic(_type [32]byte, _data []byte, _salt [32]byte) (*types.Transaction, error) {
	return _TWFactory.Contract.DeployProxyDeterministic(&_TWFactory.TransactOpts, _type, _data, _salt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactoryTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactorySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.GrantRole(&_TWFactory.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactoryTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.GrantRole(&_TWFactory.TransactOpts, role, account)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TWFactory *TWFactoryTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TWFactory *TWFactorySession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TWFactory.Contract.Multicall(&_TWFactory.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_TWFactory *TWFactoryTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _TWFactory.Contract.Multicall(&_TWFactory.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactoryTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactorySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.RenounceRole(&_TWFactory.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactoryTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.RenounceRole(&_TWFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactoryTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactorySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.RevokeRole(&_TWFactory.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TWFactory *TWFactoryTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TWFactory.Contract.RevokeRole(&_TWFactory.TransactOpts, role, account)
}

// TWFactoryImplementationAddedIterator is returned from FilterImplementationAdded and is used to iterate over the raw logs and unpacked data for ImplementationAdded events raised by the TWFactory contract.
type TWFactoryImplementationAddedIterator struct {
	Event *TWFactoryImplementationAdded // Event containing the contract specifics and raw log

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
func (it *TWFactoryImplementationAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWFactoryImplementationAdded)
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
		it.Event = new(TWFactoryImplementationAdded)
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
func (it *TWFactoryImplementationAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWFactoryImplementationAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWFactoryImplementationAdded represents a ImplementationAdded event raised by the TWFactory contract.
type TWFactoryImplementationAdded struct {
	Implementation common.Address
	ContractType   [32]byte
	Version        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterImplementationAdded is a free log retrieval operation binding the contract event 0xc39db2d47bafbb20367a9c840abffa57a2bc243c1f1e67c939ea0e89e59ed01a.
//
// Solidity: event ImplementationAdded(address implementation, bytes32 indexed contractType, uint256 version)
func (_TWFactory *TWFactoryFilterer) FilterImplementationAdded(opts *bind.FilterOpts, contractType [][32]byte) (*TWFactoryImplementationAddedIterator, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _TWFactory.contract.FilterLogs(opts, "ImplementationAdded", contractTypeRule)
	if err != nil {
		return nil, err
	}
	return &TWFactoryImplementationAddedIterator{contract: _TWFactory.contract, event: "ImplementationAdded", logs: logs, sub: sub}, nil
}

// WatchImplementationAdded is a free log subscription operation binding the contract event 0xc39db2d47bafbb20367a9c840abffa57a2bc243c1f1e67c939ea0e89e59ed01a.
//
// Solidity: event ImplementationAdded(address implementation, bytes32 indexed contractType, uint256 version)
func (_TWFactory *TWFactoryFilterer) WatchImplementationAdded(opts *bind.WatchOpts, sink chan<- *TWFactoryImplementationAdded, contractType [][32]byte) (event.Subscription, error) {

	var contractTypeRule []interface{}
	for _, contractTypeItem := range contractType {
		contractTypeRule = append(contractTypeRule, contractTypeItem)
	}

	logs, sub, err := _TWFactory.contract.WatchLogs(opts, "ImplementationAdded", contractTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWFactoryImplementationAdded)
				if err := _TWFactory.contract.UnpackLog(event, "ImplementationAdded", log); err != nil {
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

// ParseImplementationAdded is a log parse operation binding the contract event 0xc39db2d47bafbb20367a9c840abffa57a2bc243c1f1e67c939ea0e89e59ed01a.
//
// Solidity: event ImplementationAdded(address implementation, bytes32 indexed contractType, uint256 version)
func (_TWFactory *TWFactoryFilterer) ParseImplementationAdded(log types.Log) (*TWFactoryImplementationAdded, error) {
	event := new(TWFactoryImplementationAdded)
	if err := _TWFactory.contract.UnpackLog(event, "ImplementationAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TWFactoryImplementationApprovedIterator is returned from FilterImplementationApproved and is used to iterate over the raw logs and unpacked data for ImplementationApproved events raised by the TWFactory contract.
type TWFactoryImplementationApprovedIterator struct {
	Event *TWFactoryImplementationApproved // Event containing the contract specifics and raw log

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
func (it *TWFactoryImplementationApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWFactoryImplementationApproved)
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
		it.Event = new(TWFactoryImplementationApproved)
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
func (it *TWFactoryImplementationApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWFactoryImplementationApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWFactoryImplementationApproved represents a ImplementationApproved event raised by the TWFactory contract.
type TWFactoryImplementationApproved struct {
	Implementation common.Address
	IsApproved     bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterImplementationApproved is a free log retrieval operation binding the contract event 0x46c2f0868ef35772e9324a42eb6fa484490cca8494538a909cf05c897d7d4108.
//
// Solidity: event ImplementationApproved(address implementation, bool isApproved)
func (_TWFactory *TWFactoryFilterer) FilterImplementationApproved(opts *bind.FilterOpts) (*TWFactoryImplementationApprovedIterator, error) {

	logs, sub, err := _TWFactory.contract.FilterLogs(opts, "ImplementationApproved")
	if err != nil {
		return nil, err
	}
	return &TWFactoryImplementationApprovedIterator{contract: _TWFactory.contract, event: "ImplementationApproved", logs: logs, sub: sub}, nil
}

// WatchImplementationApproved is a free log subscription operation binding the contract event 0x46c2f0868ef35772e9324a42eb6fa484490cca8494538a909cf05c897d7d4108.
//
// Solidity: event ImplementationApproved(address implementation, bool isApproved)
func (_TWFactory *TWFactoryFilterer) WatchImplementationApproved(opts *bind.WatchOpts, sink chan<- *TWFactoryImplementationApproved) (event.Subscription, error) {

	logs, sub, err := _TWFactory.contract.WatchLogs(opts, "ImplementationApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWFactoryImplementationApproved)
				if err := _TWFactory.contract.UnpackLog(event, "ImplementationApproved", log); err != nil {
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

// ParseImplementationApproved is a log parse operation binding the contract event 0x46c2f0868ef35772e9324a42eb6fa484490cca8494538a909cf05c897d7d4108.
//
// Solidity: event ImplementationApproved(address implementation, bool isApproved)
func (_TWFactory *TWFactoryFilterer) ParseImplementationApproved(log types.Log) (*TWFactoryImplementationApproved, error) {
	event := new(TWFactoryImplementationApproved)
	if err := _TWFactory.contract.UnpackLog(event, "ImplementationApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TWFactoryProxyDeployedIterator is returned from FilterProxyDeployed and is used to iterate over the raw logs and unpacked data for ProxyDeployed events raised by the TWFactory contract.
type TWFactoryProxyDeployedIterator struct {
	Event *TWFactoryProxyDeployed // Event containing the contract specifics and raw log

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
func (it *TWFactoryProxyDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWFactoryProxyDeployed)
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
		it.Event = new(TWFactoryProxyDeployed)
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
func (it *TWFactoryProxyDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWFactoryProxyDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWFactoryProxyDeployed represents a ProxyDeployed event raised by the TWFactory contract.
type TWFactoryProxyDeployed struct {
	Implementation common.Address
	Proxy          common.Address
	Deployer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProxyDeployed is a free log retrieval operation binding the contract event 0x9e0862c4ebff2150fbbfd3f8547483f55bdec0c34fd977d3fccaa55d6c4ce784.
//
// Solidity: event ProxyDeployed(address indexed implementation, address proxy, address indexed deployer)
func (_TWFactory *TWFactoryFilterer) FilterProxyDeployed(opts *bind.FilterOpts, implementation []common.Address, deployer []common.Address) (*TWFactoryProxyDeployedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _TWFactory.contract.FilterLogs(opts, "ProxyDeployed", implementationRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return &TWFactoryProxyDeployedIterator{contract: _TWFactory.contract, event: "ProxyDeployed", logs: logs, sub: sub}, nil
}

// WatchProxyDeployed is a free log subscription operation binding the contract event 0x9e0862c4ebff2150fbbfd3f8547483f55bdec0c34fd977d3fccaa55d6c4ce784.
//
// Solidity: event ProxyDeployed(address indexed implementation, address proxy, address indexed deployer)
func (_TWFactory *TWFactoryFilterer) WatchProxyDeployed(opts *bind.WatchOpts, sink chan<- *TWFactoryProxyDeployed, implementation []common.Address, deployer []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	var deployerRule []interface{}
	for _, deployerItem := range deployer {
		deployerRule = append(deployerRule, deployerItem)
	}

	logs, sub, err := _TWFactory.contract.WatchLogs(opts, "ProxyDeployed", implementationRule, deployerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWFactoryProxyDeployed)
				if err := _TWFactory.contract.UnpackLog(event, "ProxyDeployed", log); err != nil {
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

// ParseProxyDeployed is a log parse operation binding the contract event 0x9e0862c4ebff2150fbbfd3f8547483f55bdec0c34fd977d3fccaa55d6c4ce784.
//
// Solidity: event ProxyDeployed(address indexed implementation, address proxy, address indexed deployer)
func (_TWFactory *TWFactoryFilterer) ParseProxyDeployed(log types.Log) (*TWFactoryProxyDeployed, error) {
	event := new(TWFactoryProxyDeployed)
	if err := _TWFactory.contract.UnpackLog(event, "ProxyDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TWFactoryRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TWFactory contract.
type TWFactoryRoleAdminChangedIterator struct {
	Event *TWFactoryRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TWFactoryRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWFactoryRoleAdminChanged)
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
		it.Event = new(TWFactoryRoleAdminChanged)
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
func (it *TWFactoryRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWFactoryRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWFactoryRoleAdminChanged represents a RoleAdminChanged event raised by the TWFactory contract.
type TWFactoryRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TWFactory *TWFactoryFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TWFactoryRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TWFactory.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TWFactoryRoleAdminChangedIterator{contract: _TWFactory.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TWFactory *TWFactoryFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TWFactoryRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _TWFactory.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWFactoryRoleAdminChanged)
				if err := _TWFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TWFactory *TWFactoryFilterer) ParseRoleAdminChanged(log types.Log) (*TWFactoryRoleAdminChanged, error) {
	event := new(TWFactoryRoleAdminChanged)
	if err := _TWFactory.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TWFactoryRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TWFactory contract.
type TWFactoryRoleGrantedIterator struct {
	Event *TWFactoryRoleGranted // Event containing the contract specifics and raw log

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
func (it *TWFactoryRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWFactoryRoleGranted)
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
		it.Event = new(TWFactoryRoleGranted)
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
func (it *TWFactoryRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWFactoryRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWFactoryRoleGranted represents a RoleGranted event raised by the TWFactory contract.
type TWFactoryRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TWFactory *TWFactoryFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TWFactoryRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TWFactory.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TWFactoryRoleGrantedIterator{contract: _TWFactory.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TWFactory *TWFactoryFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TWFactoryRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TWFactory.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWFactoryRoleGranted)
				if err := _TWFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TWFactory *TWFactoryFilterer) ParseRoleGranted(log types.Log) (*TWFactoryRoleGranted, error) {
	event := new(TWFactoryRoleGranted)
	if err := _TWFactory.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TWFactoryRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TWFactory contract.
type TWFactoryRoleRevokedIterator struct {
	Event *TWFactoryRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TWFactoryRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TWFactoryRoleRevoked)
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
		it.Event = new(TWFactoryRoleRevoked)
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
func (it *TWFactoryRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TWFactoryRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TWFactoryRoleRevoked represents a RoleRevoked event raised by the TWFactory contract.
type TWFactoryRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TWFactory *TWFactoryFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TWFactoryRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TWFactory.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TWFactoryRoleRevokedIterator{contract: _TWFactory.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TWFactory *TWFactoryFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TWFactoryRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _TWFactory.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TWFactoryRoleRevoked)
				if err := _TWFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TWFactory *TWFactoryFilterer) ParseRoleRevoked(log types.Log) (*TWFactoryRoleRevoked, error) {
	event := new(TWFactoryRoleRevoked)
	if err := _TWFactory.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
