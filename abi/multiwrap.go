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

// ITokenBundleToken is an auto generated low-level Go binding around an user-defined struct.
type ITokenBundleToken struct {
	AssetContract common.Address
	TokenType     uint8
	TokenId       *big.Int
	TotalAmount   *big.Int
}

// MultiwrapMetaData contains all meta data concerning the Multiwrap contract.
var MultiwrapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nativeTokenWrapper\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prevURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRoyaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoyaltyBps\",\"type\":\"uint256\"}],\"name\":\"DefaultRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"unwrapper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientOfWrappedContents\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIdOfWrappedToken\",\"type\":\"uint256\"}],\"name\":\"TokensUnwrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wrapper\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipientOfWrappedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenIdOfWrappedToken\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"enumITokenBundle.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structITokenBundle.Token[]\",\"name\":\"wrappedContents\",\"type\":\"tuple[]\"}],\"name\":\"TokensWrapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyInfoForToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bundleId\",\"type\":\"uint256\"}],\"name\":\"getTokenCountOfBundle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bundleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getTokenOfBundle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"enumITokenBundle.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"internalType\":\"structITokenBundle.Token\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bundleId\",\"type\":\"uint256\"}],\"name\":\"getUriOfBundle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getWrappedContents\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"enumITokenBundle.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"internalType\":\"structITokenBundle.Token[]\",\"name\":\"contents\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoleWithSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_trustedForwarders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenIdToMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyInfoForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"enumITokenBundle.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"internalType\":\"structITokenBundle.Token[]\",\"name\":\"_tokensToWrap\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"_uriForWrappedToken\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"wrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MultiwrapABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiwrapMetaData.ABI instead.
var MultiwrapABI = MultiwrapMetaData.ABI

// Multiwrap is an auto generated Go binding around an Ethereum contract.
type Multiwrap struct {
	MultiwrapCaller     // Read-only binding to the contract
	MultiwrapTransactor // Write-only binding to the contract
	MultiwrapFilterer   // Log filterer for contract events
}

// MultiwrapCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiwrapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiwrapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiwrapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiwrapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiwrapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiwrapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiwrapSession struct {
	Contract     *Multiwrap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiwrapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiwrapCallerSession struct {
	Contract *MultiwrapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultiwrapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiwrapTransactorSession struct {
	Contract     *MultiwrapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultiwrapRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiwrapRaw struct {
	Contract *Multiwrap // Generic contract binding to access the raw methods on
}

// MultiwrapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiwrapCallerRaw struct {
	Contract *MultiwrapCaller // Generic read-only contract binding to access the raw methods on
}

// MultiwrapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiwrapTransactorRaw struct {
	Contract *MultiwrapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiwrap creates a new instance of Multiwrap, bound to a specific deployed contract.
func NewMultiwrap(address common.Address, backend bind.ContractBackend) (*Multiwrap, error) {
	contract, err := bindMultiwrap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multiwrap{MultiwrapCaller: MultiwrapCaller{contract: contract}, MultiwrapTransactor: MultiwrapTransactor{contract: contract}, MultiwrapFilterer: MultiwrapFilterer{contract: contract}}, nil
}

// NewMultiwrapCaller creates a new read-only instance of Multiwrap, bound to a specific deployed contract.
func NewMultiwrapCaller(address common.Address, caller bind.ContractCaller) (*MultiwrapCaller, error) {
	contract, err := bindMultiwrap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiwrapCaller{contract: contract}, nil
}

// NewMultiwrapTransactor creates a new write-only instance of Multiwrap, bound to a specific deployed contract.
func NewMultiwrapTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiwrapTransactor, error) {
	contract, err := bindMultiwrap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiwrapTransactor{contract: contract}, nil
}

// NewMultiwrapFilterer creates a new log filterer instance of Multiwrap, bound to a specific deployed contract.
func NewMultiwrapFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiwrapFilterer, error) {
	contract, err := bindMultiwrap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiwrapFilterer{contract: contract}, nil
}

// bindMultiwrap binds a generic wrapper to an already deployed contract.
func bindMultiwrap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiwrapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multiwrap *MultiwrapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multiwrap.Contract.MultiwrapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multiwrap *MultiwrapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multiwrap.Contract.MultiwrapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multiwrap *MultiwrapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multiwrap.Contract.MultiwrapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multiwrap *MultiwrapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multiwrap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multiwrap *MultiwrapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multiwrap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multiwrap *MultiwrapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multiwrap.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Multiwrap *MultiwrapCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Multiwrap *MultiwrapSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Multiwrap.Contract.DEFAULTADMINROLE(&_Multiwrap.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Multiwrap *MultiwrapCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Multiwrap.Contract.DEFAULTADMINROLE(&_Multiwrap.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Multiwrap *MultiwrapCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Multiwrap *MultiwrapSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Multiwrap.Contract.BalanceOf(&_Multiwrap.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Multiwrap *MultiwrapCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Multiwrap.Contract.BalanceOf(&_Multiwrap.CallOpts, owner)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_Multiwrap *MultiwrapCaller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_Multiwrap *MultiwrapSession) ContractType() ([32]byte, error) {
	return _Multiwrap.Contract.ContractType(&_Multiwrap.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_Multiwrap *MultiwrapCallerSession) ContractType() ([32]byte, error) {
	return _Multiwrap.Contract.ContractType(&_Multiwrap.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Multiwrap *MultiwrapCaller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Multiwrap *MultiwrapSession) InternalContractURI() (string, error) {
	return _Multiwrap.Contract.InternalContractURI(&_Multiwrap.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Multiwrap *MultiwrapCallerSession) InternalContractURI() (string, error) {
	return _Multiwrap.Contract.InternalContractURI(&_Multiwrap.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_Multiwrap *MultiwrapCaller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_Multiwrap *MultiwrapSession) ContractVersion() (uint8, error) {
	return _Multiwrap.Contract.ContractVersion(&_Multiwrap.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_Multiwrap *MultiwrapCallerSession) ContractVersion() (uint8, error) {
	return _Multiwrap.Contract.ContractVersion(&_Multiwrap.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Multiwrap *MultiwrapCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Multiwrap *MultiwrapSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Multiwrap.Contract.GetApproved(&_Multiwrap.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Multiwrap *MultiwrapCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Multiwrap.Contract.GetApproved(&_Multiwrap.CallOpts, tokenId)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_Multiwrap *MultiwrapCaller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

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
func (_Multiwrap *MultiwrapSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _Multiwrap.Contract.GetDefaultRoyaltyInfo(&_Multiwrap.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_Multiwrap *MultiwrapCallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _Multiwrap.Contract.GetDefaultRoyaltyInfo(&_Multiwrap.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Multiwrap *MultiwrapCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Multiwrap *MultiwrapSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Multiwrap.Contract.GetRoleAdmin(&_Multiwrap.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Multiwrap *MultiwrapCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Multiwrap.Contract.GetRoleAdmin(&_Multiwrap.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_Multiwrap *MultiwrapCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_Multiwrap *MultiwrapSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Multiwrap.Contract.GetRoleMember(&_Multiwrap.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_Multiwrap *MultiwrapCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Multiwrap.Contract.GetRoleMember(&_Multiwrap.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_Multiwrap *MultiwrapCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_Multiwrap *MultiwrapSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Multiwrap.Contract.GetRoleMemberCount(&_Multiwrap.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_Multiwrap *MultiwrapCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Multiwrap.Contract.GetRoleMemberCount(&_Multiwrap.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_Multiwrap *MultiwrapCaller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_Multiwrap *MultiwrapSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _Multiwrap.Contract.GetRoyaltyInfoForToken(&_Multiwrap.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_Multiwrap *MultiwrapCallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _Multiwrap.Contract.GetRoyaltyInfoForToken(&_Multiwrap.CallOpts, _tokenId)
}

// GetTokenCountOfBundle is a free data retrieval call binding the contract method 0xd0d2fe25.
//
// Solidity: function getTokenCountOfBundle(uint256 _bundleId) view returns(uint256)
func (_Multiwrap *MultiwrapCaller) GetTokenCountOfBundle(opts *bind.CallOpts, _bundleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getTokenCountOfBundle", _bundleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCountOfBundle is a free data retrieval call binding the contract method 0xd0d2fe25.
//
// Solidity: function getTokenCountOfBundle(uint256 _bundleId) view returns(uint256)
func (_Multiwrap *MultiwrapSession) GetTokenCountOfBundle(_bundleId *big.Int) (*big.Int, error) {
	return _Multiwrap.Contract.GetTokenCountOfBundle(&_Multiwrap.CallOpts, _bundleId)
}

// GetTokenCountOfBundle is a free data retrieval call binding the contract method 0xd0d2fe25.
//
// Solidity: function getTokenCountOfBundle(uint256 _bundleId) view returns(uint256)
func (_Multiwrap *MultiwrapCallerSession) GetTokenCountOfBundle(_bundleId *big.Int) (*big.Int, error) {
	return _Multiwrap.Contract.GetTokenCountOfBundle(&_Multiwrap.CallOpts, _bundleId)
}

// GetTokenOfBundle is a free data retrieval call binding the contract method 0x1da799c9.
//
// Solidity: function getTokenOfBundle(uint256 _bundleId, uint256 index) view returns((address,uint8,uint256,uint256))
func (_Multiwrap *MultiwrapCaller) GetTokenOfBundle(opts *bind.CallOpts, _bundleId *big.Int, index *big.Int) (ITokenBundleToken, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getTokenOfBundle", _bundleId, index)

	if err != nil {
		return *new(ITokenBundleToken), err
	}

	out0 := *abi.ConvertType(out[0], new(ITokenBundleToken)).(*ITokenBundleToken)

	return out0, err

}

// GetTokenOfBundle is a free data retrieval call binding the contract method 0x1da799c9.
//
// Solidity: function getTokenOfBundle(uint256 _bundleId, uint256 index) view returns((address,uint8,uint256,uint256))
func (_Multiwrap *MultiwrapSession) GetTokenOfBundle(_bundleId *big.Int, index *big.Int) (ITokenBundleToken, error) {
	return _Multiwrap.Contract.GetTokenOfBundle(&_Multiwrap.CallOpts, _bundleId, index)
}

// GetTokenOfBundle is a free data retrieval call binding the contract method 0x1da799c9.
//
// Solidity: function getTokenOfBundle(uint256 _bundleId, uint256 index) view returns((address,uint8,uint256,uint256))
func (_Multiwrap *MultiwrapCallerSession) GetTokenOfBundle(_bundleId *big.Int, index *big.Int) (ITokenBundleToken, error) {
	return _Multiwrap.Contract.GetTokenOfBundle(&_Multiwrap.CallOpts, _bundleId, index)
}

// GetUriOfBundle is a free data retrieval call binding the contract method 0x61195e94.
//
// Solidity: function getUriOfBundle(uint256 _bundleId) view returns(string)
func (_Multiwrap *MultiwrapCaller) GetUriOfBundle(opts *bind.CallOpts, _bundleId *big.Int) (string, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getUriOfBundle", _bundleId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUriOfBundle is a free data retrieval call binding the contract method 0x61195e94.
//
// Solidity: function getUriOfBundle(uint256 _bundleId) view returns(string)
func (_Multiwrap *MultiwrapSession) GetUriOfBundle(_bundleId *big.Int) (string, error) {
	return _Multiwrap.Contract.GetUriOfBundle(&_Multiwrap.CallOpts, _bundleId)
}

// GetUriOfBundle is a free data retrieval call binding the contract method 0x61195e94.
//
// Solidity: function getUriOfBundle(uint256 _bundleId) view returns(string)
func (_Multiwrap *MultiwrapCallerSession) GetUriOfBundle(_bundleId *big.Int) (string, error) {
	return _Multiwrap.Contract.GetUriOfBundle(&_Multiwrap.CallOpts, _bundleId)
}

// GetWrappedContents is a free data retrieval call binding the contract method 0xd5576d26.
//
// Solidity: function getWrappedContents(uint256 _tokenId) view returns((address,uint8,uint256,uint256)[] contents)
func (_Multiwrap *MultiwrapCaller) GetWrappedContents(opts *bind.CallOpts, _tokenId *big.Int) ([]ITokenBundleToken, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "getWrappedContents", _tokenId)

	if err != nil {
		return *new([]ITokenBundleToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]ITokenBundleToken)).(*[]ITokenBundleToken)

	return out0, err

}

// GetWrappedContents is a free data retrieval call binding the contract method 0xd5576d26.
//
// Solidity: function getWrappedContents(uint256 _tokenId) view returns((address,uint8,uint256,uint256)[] contents)
func (_Multiwrap *MultiwrapSession) GetWrappedContents(_tokenId *big.Int) ([]ITokenBundleToken, error) {
	return _Multiwrap.Contract.GetWrappedContents(&_Multiwrap.CallOpts, _tokenId)
}

// GetWrappedContents is a free data retrieval call binding the contract method 0xd5576d26.
//
// Solidity: function getWrappedContents(uint256 _tokenId) view returns((address,uint8,uint256,uint256)[] contents)
func (_Multiwrap *MultiwrapCallerSession) GetWrappedContents(_tokenId *big.Int) ([]ITokenBundleToken, error) {
	return _Multiwrap.Contract.GetWrappedContents(&_Multiwrap.CallOpts, _tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Multiwrap *MultiwrapCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Multiwrap *MultiwrapSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Multiwrap.Contract.HasRole(&_Multiwrap.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Multiwrap *MultiwrapCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Multiwrap.Contract.HasRole(&_Multiwrap.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_Multiwrap *MultiwrapCaller) HasRoleWithSwitch(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "hasRoleWithSwitch", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_Multiwrap *MultiwrapSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _Multiwrap.Contract.HasRoleWithSwitch(&_Multiwrap.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_Multiwrap *MultiwrapCallerSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _Multiwrap.Contract.HasRoleWithSwitch(&_Multiwrap.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Multiwrap *MultiwrapCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Multiwrap *MultiwrapSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Multiwrap.Contract.IsApprovedForAll(&_Multiwrap.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Multiwrap *MultiwrapCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Multiwrap.Contract.IsApprovedForAll(&_Multiwrap.CallOpts, owner, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Multiwrap *MultiwrapCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Multiwrap *MultiwrapSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Multiwrap.Contract.IsTrustedForwarder(&_Multiwrap.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Multiwrap *MultiwrapCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Multiwrap.Contract.IsTrustedForwarder(&_Multiwrap.CallOpts, forwarder)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Multiwrap *MultiwrapCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Multiwrap *MultiwrapSession) Name() (string, error) {
	return _Multiwrap.Contract.Name(&_Multiwrap.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Multiwrap *MultiwrapCallerSession) Name() (string, error) {
	return _Multiwrap.Contract.Name(&_Multiwrap.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_Multiwrap *MultiwrapCaller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_Multiwrap *MultiwrapSession) NextTokenIdToMint() (*big.Int, error) {
	return _Multiwrap.Contract.NextTokenIdToMint(&_Multiwrap.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_Multiwrap *MultiwrapCallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _Multiwrap.Contract.NextTokenIdToMint(&_Multiwrap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Multiwrap *MultiwrapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Multiwrap *MultiwrapSession) Owner() (common.Address, error) {
	return _Multiwrap.Contract.Owner(&_Multiwrap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Multiwrap *MultiwrapCallerSession) Owner() (common.Address, error) {
	return _Multiwrap.Contract.Owner(&_Multiwrap.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Multiwrap *MultiwrapCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Multiwrap *MultiwrapSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Multiwrap.Contract.OwnerOf(&_Multiwrap.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Multiwrap *MultiwrapCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Multiwrap.Contract.OwnerOf(&_Multiwrap.CallOpts, tokenId)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Multiwrap *MultiwrapCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_Multiwrap *MultiwrapSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Multiwrap.Contract.RoyaltyInfo(&_Multiwrap.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Multiwrap *MultiwrapCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Multiwrap.Contract.RoyaltyInfo(&_Multiwrap.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Multiwrap *MultiwrapCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Multiwrap *MultiwrapSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Multiwrap.Contract.SupportsInterface(&_Multiwrap.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Multiwrap *MultiwrapCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Multiwrap.Contract.SupportsInterface(&_Multiwrap.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Multiwrap *MultiwrapCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Multiwrap *MultiwrapSession) Symbol() (string, error) {
	return _Multiwrap.Contract.Symbol(&_Multiwrap.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Multiwrap *MultiwrapCallerSession) Symbol() (string, error) {
	return _Multiwrap.Contract.Symbol(&_Multiwrap.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Multiwrap *MultiwrapCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Multiwrap.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Multiwrap *MultiwrapSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Multiwrap.Contract.TokenURI(&_Multiwrap.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_Multiwrap *MultiwrapCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Multiwrap.Contract.TokenURI(&_Multiwrap.CallOpts, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.Approve(&_Multiwrap.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.Approve(&_Multiwrap.TransactOpts, to, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.GrantRole(&_Multiwrap.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.GrantRole(&_Multiwrap.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x754b8fe7.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_Multiwrap *MultiwrapTransactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _royaltyRecipient, _royaltyBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x754b8fe7.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_Multiwrap *MultiwrapSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.Initialize(&_Multiwrap.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _royaltyRecipient, _royaltyBps)
}

// Initialize is a paid mutator transaction binding the contract method 0x754b8fe7.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_Multiwrap *MultiwrapTransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.Initialize(&_Multiwrap.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _royaltyRecipient, _royaltyBps)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multiwrap *MultiwrapTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multiwrap *MultiwrapSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.Multicall(&_Multiwrap.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Multiwrap *MultiwrapTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.Multicall(&_Multiwrap.TransactOpts, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.OnERC1155BatchReceived(&_Multiwrap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.OnERC1155BatchReceived(&_Multiwrap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.OnERC1155Received(&_Multiwrap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.OnERC1155Received(&_Multiwrap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.OnERC721Received(&_Multiwrap.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Multiwrap *MultiwrapTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.OnERC721Received(&_Multiwrap.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.RenounceRole(&_Multiwrap.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.RenounceRole(&_Multiwrap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.RevokeRole(&_Multiwrap.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Multiwrap *MultiwrapTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.RevokeRole(&_Multiwrap.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.SafeTransferFrom(&_Multiwrap.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.SafeTransferFrom(&_Multiwrap.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Multiwrap *MultiwrapTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Multiwrap *MultiwrapSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.SafeTransferFrom0(&_Multiwrap.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Multiwrap *MultiwrapTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Multiwrap.Contract.SafeTransferFrom0(&_Multiwrap.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Multiwrap *MultiwrapTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Multiwrap *MultiwrapSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetApprovalForAll(&_Multiwrap.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Multiwrap *MultiwrapTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetApprovalForAll(&_Multiwrap.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Multiwrap *MultiwrapTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Multiwrap *MultiwrapSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetContractURI(&_Multiwrap.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Multiwrap *MultiwrapTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetContractURI(&_Multiwrap.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_Multiwrap *MultiwrapTransactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_Multiwrap *MultiwrapSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetDefaultRoyaltyInfo(&_Multiwrap.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_Multiwrap *MultiwrapTransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetDefaultRoyaltyInfo(&_Multiwrap.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_Multiwrap *MultiwrapTransactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_Multiwrap *MultiwrapSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetOwner(&_Multiwrap.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_Multiwrap *MultiwrapTransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetOwner(&_Multiwrap.TransactOpts, _newOwner)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_Multiwrap *MultiwrapTransactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_Multiwrap *MultiwrapSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetRoyaltyInfoForToken(&_Multiwrap.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_Multiwrap *MultiwrapTransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.SetRoyaltyInfoForToken(&_Multiwrap.TransactOpts, _tokenId, _recipient, _bps)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.TransferFrom(&_Multiwrap.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Multiwrap *MultiwrapTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Multiwrap.Contract.TransferFrom(&_Multiwrap.TransactOpts, from, to, tokenId)
}

// Unwrap is a paid mutator transaction binding the contract method 0x7647691d.
//
// Solidity: function unwrap(uint256 _tokenId, address _recipient) returns()
func (_Multiwrap *MultiwrapTransactor) Unwrap(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "unwrap", _tokenId, _recipient)
}

// Unwrap is a paid mutator transaction binding the contract method 0x7647691d.
//
// Solidity: function unwrap(uint256 _tokenId, address _recipient) returns()
func (_Multiwrap *MultiwrapSession) Unwrap(_tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.Unwrap(&_Multiwrap.TransactOpts, _tokenId, _recipient)
}

// Unwrap is a paid mutator transaction binding the contract method 0x7647691d.
//
// Solidity: function unwrap(uint256 _tokenId, address _recipient) returns()
func (_Multiwrap *MultiwrapTransactorSession) Unwrap(_tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.Unwrap(&_Multiwrap.TransactOpts, _tokenId, _recipient)
}

// Wrap is a paid mutator transaction binding the contract method 0x29e471dd.
//
// Solidity: function wrap((address,uint8,uint256,uint256)[] _tokensToWrap, string _uriForWrappedToken, address _recipient) payable returns(uint256 tokenId)
func (_Multiwrap *MultiwrapTransactor) Wrap(opts *bind.TransactOpts, _tokensToWrap []ITokenBundleToken, _uriForWrappedToken string, _recipient common.Address) (*types.Transaction, error) {
	return _Multiwrap.contract.Transact(opts, "wrap", _tokensToWrap, _uriForWrappedToken, _recipient)
}

// Wrap is a paid mutator transaction binding the contract method 0x29e471dd.
//
// Solidity: function wrap((address,uint8,uint256,uint256)[] _tokensToWrap, string _uriForWrappedToken, address _recipient) payable returns(uint256 tokenId)
func (_Multiwrap *MultiwrapSession) Wrap(_tokensToWrap []ITokenBundleToken, _uriForWrappedToken string, _recipient common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.Wrap(&_Multiwrap.TransactOpts, _tokensToWrap, _uriForWrappedToken, _recipient)
}

// Wrap is a paid mutator transaction binding the contract method 0x29e471dd.
//
// Solidity: function wrap((address,uint8,uint256,uint256)[] _tokensToWrap, string _uriForWrappedToken, address _recipient) payable returns(uint256 tokenId)
func (_Multiwrap *MultiwrapTransactorSession) Wrap(_tokensToWrap []ITokenBundleToken, _uriForWrappedToken string, _recipient common.Address) (*types.Transaction, error) {
	return _Multiwrap.Contract.Wrap(&_Multiwrap.TransactOpts, _tokensToWrap, _uriForWrappedToken, _recipient)
}

// MultiwrapApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Multiwrap contract.
type MultiwrapApprovalIterator struct {
	Event *MultiwrapApproval // Event containing the contract specifics and raw log

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
func (it *MultiwrapApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapApproval)
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
		it.Event = new(MultiwrapApproval)
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
func (it *MultiwrapApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapApproval represents a Approval event raised by the Multiwrap contract.
type MultiwrapApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Multiwrap *MultiwrapFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MultiwrapApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapApprovalIterator{contract: _Multiwrap.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Multiwrap *MultiwrapFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MultiwrapApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapApproval)
				if err := _Multiwrap.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Multiwrap *MultiwrapFilterer) ParseApproval(log types.Log) (*MultiwrapApproval, error) {
	event := new(MultiwrapApproval)
	if err := _Multiwrap.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Multiwrap contract.
type MultiwrapApprovalForAllIterator struct {
	Event *MultiwrapApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MultiwrapApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapApprovalForAll)
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
		it.Event = new(MultiwrapApprovalForAll)
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
func (it *MultiwrapApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapApprovalForAll represents a ApprovalForAll event raised by the Multiwrap contract.
type MultiwrapApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Multiwrap *MultiwrapFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MultiwrapApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapApprovalForAllIterator{contract: _Multiwrap.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Multiwrap *MultiwrapFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MultiwrapApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapApprovalForAll)
				if err := _Multiwrap.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Multiwrap *MultiwrapFilterer) ParseApprovalForAll(log types.Log) (*MultiwrapApprovalForAll, error) {
	event := new(MultiwrapApprovalForAll)
	if err := _Multiwrap.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the Multiwrap contract.
type MultiwrapContractURIUpdatedIterator struct {
	Event *MultiwrapContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *MultiwrapContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapContractURIUpdated)
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
		it.Event = new(MultiwrapContractURIUpdated)
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
func (it *MultiwrapContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapContractURIUpdated represents a ContractURIUpdated event raised by the Multiwrap contract.
type MultiwrapContractURIUpdated struct {
	PrevURI string
	NewURI  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_Multiwrap *MultiwrapFilterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*MultiwrapContractURIUpdatedIterator, error) {

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiwrapContractURIUpdatedIterator{contract: _Multiwrap.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_Multiwrap *MultiwrapFilterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *MultiwrapContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapContractURIUpdated)
				if err := _Multiwrap.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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

// ParseContractURIUpdated is a log parse operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_Multiwrap *MultiwrapFilterer) ParseContractURIUpdated(log types.Log) (*MultiwrapContractURIUpdated, error) {
	event := new(MultiwrapContractURIUpdated)
	if err := _Multiwrap.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapDefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the Multiwrap contract.
type MultiwrapDefaultRoyaltyIterator struct {
	Event *MultiwrapDefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *MultiwrapDefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapDefaultRoyalty)
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
		it.Event = new(MultiwrapDefaultRoyalty)
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
func (it *MultiwrapDefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapDefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapDefaultRoyalty represents a DefaultRoyalty event raised by the Multiwrap contract.
type MultiwrapDefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_Multiwrap *MultiwrapFilterer) FilterDefaultRoyalty(opts *bind.FilterOpts) (*MultiwrapDefaultRoyaltyIterator, error) {

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "DefaultRoyalty")
	if err != nil {
		return nil, err
	}
	return &MultiwrapDefaultRoyaltyIterator{contract: _Multiwrap.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_Multiwrap *MultiwrapFilterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *MultiwrapDefaultRoyalty) (event.Subscription, error) {

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "DefaultRoyalty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapDefaultRoyalty)
				if err := _Multiwrap.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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
// Solidity: event DefaultRoyalty(address newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_Multiwrap *MultiwrapFilterer) ParseDefaultRoyalty(log types.Log) (*MultiwrapDefaultRoyalty, error) {
	event := new(MultiwrapDefaultRoyalty)
	if err := _Multiwrap.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapOwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the Multiwrap contract.
type MultiwrapOwnerUpdatedIterator struct {
	Event *MultiwrapOwnerUpdated // Event containing the contract specifics and raw log

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
func (it *MultiwrapOwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapOwnerUpdated)
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
		it.Event = new(MultiwrapOwnerUpdated)
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
func (it *MultiwrapOwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapOwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapOwnerUpdated represents a OwnerUpdated event raised by the Multiwrap contract.
type MultiwrapOwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address prevOwner, address newOwner)
func (_Multiwrap *MultiwrapFilterer) FilterOwnerUpdated(opts *bind.FilterOpts) (*MultiwrapOwnerUpdatedIterator, error) {

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &MultiwrapOwnerUpdatedIterator{contract: _Multiwrap.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address prevOwner, address newOwner)
func (_Multiwrap *MultiwrapFilterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *MultiwrapOwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapOwnerUpdated)
				if err := _Multiwrap.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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

// ParseOwnerUpdated is a log parse operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address prevOwner, address newOwner)
func (_Multiwrap *MultiwrapFilterer) ParseOwnerUpdated(log types.Log) (*MultiwrapOwnerUpdated, error) {
	event := new(MultiwrapOwnerUpdated)
	if err := _Multiwrap.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Multiwrap contract.
type MultiwrapRoleAdminChangedIterator struct {
	Event *MultiwrapRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MultiwrapRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapRoleAdminChanged)
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
		it.Event = new(MultiwrapRoleAdminChanged)
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
func (it *MultiwrapRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapRoleAdminChanged represents a RoleAdminChanged event raised by the Multiwrap contract.
type MultiwrapRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Multiwrap *MultiwrapFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MultiwrapRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapRoleAdminChangedIterator{contract: _Multiwrap.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Multiwrap *MultiwrapFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MultiwrapRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapRoleAdminChanged)
				if err := _Multiwrap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Multiwrap *MultiwrapFilterer) ParseRoleAdminChanged(log types.Log) (*MultiwrapRoleAdminChanged, error) {
	event := new(MultiwrapRoleAdminChanged)
	if err := _Multiwrap.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Multiwrap contract.
type MultiwrapRoleGrantedIterator struct {
	Event *MultiwrapRoleGranted // Event containing the contract specifics and raw log

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
func (it *MultiwrapRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapRoleGranted)
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
		it.Event = new(MultiwrapRoleGranted)
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
func (it *MultiwrapRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapRoleGranted represents a RoleGranted event raised by the Multiwrap contract.
type MultiwrapRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Multiwrap *MultiwrapFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MultiwrapRoleGrantedIterator, error) {

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

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapRoleGrantedIterator{contract: _Multiwrap.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Multiwrap *MultiwrapFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MultiwrapRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapRoleGranted)
				if err := _Multiwrap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Multiwrap *MultiwrapFilterer) ParseRoleGranted(log types.Log) (*MultiwrapRoleGranted, error) {
	event := new(MultiwrapRoleGranted)
	if err := _Multiwrap.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Multiwrap contract.
type MultiwrapRoleRevokedIterator struct {
	Event *MultiwrapRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MultiwrapRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapRoleRevoked)
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
		it.Event = new(MultiwrapRoleRevoked)
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
func (it *MultiwrapRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapRoleRevoked represents a RoleRevoked event raised by the Multiwrap contract.
type MultiwrapRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Multiwrap *MultiwrapFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MultiwrapRoleRevokedIterator, error) {

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

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapRoleRevokedIterator{contract: _Multiwrap.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Multiwrap *MultiwrapFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MultiwrapRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapRoleRevoked)
				if err := _Multiwrap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Multiwrap *MultiwrapFilterer) ParseRoleRevoked(log types.Log) (*MultiwrapRoleRevoked, error) {
	event := new(MultiwrapRoleRevoked)
	if err := _Multiwrap.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapRoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the Multiwrap contract.
type MultiwrapRoyaltyForTokenIterator struct {
	Event *MultiwrapRoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *MultiwrapRoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapRoyaltyForToken)
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
		it.Event = new(MultiwrapRoyaltyForToken)
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
func (it *MultiwrapRoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapRoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapRoyaltyForToken represents a RoyaltyForToken event raised by the Multiwrap contract.
type MultiwrapRoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address royaltyRecipient, uint256 royaltyBps)
func (_Multiwrap *MultiwrapFilterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int) (*MultiwrapRoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapRoyaltyForTokenIterator{contract: _Multiwrap.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address royaltyRecipient, uint256 royaltyBps)
func (_Multiwrap *MultiwrapFilterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *MultiwrapRoyaltyForToken, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapRoyaltyForToken)
				if err := _Multiwrap.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address royaltyRecipient, uint256 royaltyBps)
func (_Multiwrap *MultiwrapFilterer) ParseRoyaltyForToken(log types.Log) (*MultiwrapRoyaltyForToken, error) {
	event := new(MultiwrapRoyaltyForToken)
	if err := _Multiwrap.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapTokensUnwrappedIterator is returned from FilterTokensUnwrapped and is used to iterate over the raw logs and unpacked data for TokensUnwrapped events raised by the Multiwrap contract.
type MultiwrapTokensUnwrappedIterator struct {
	Event *MultiwrapTokensUnwrapped // Event containing the contract specifics and raw log

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
func (it *MultiwrapTokensUnwrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapTokensUnwrapped)
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
		it.Event = new(MultiwrapTokensUnwrapped)
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
func (it *MultiwrapTokensUnwrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapTokensUnwrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapTokensUnwrapped represents a TokensUnwrapped event raised by the Multiwrap contract.
type MultiwrapTokensUnwrapped struct {
	Unwrapper                  common.Address
	RecipientOfWrappedContents common.Address
	TokenIdOfWrappedToken      *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterTokensUnwrapped is a free log retrieval operation binding the contract event 0xe9a1b206a63887b7a73ef37983c4105047ae75c9ac0920a1a256eea52f264a73.
//
// Solidity: event TokensUnwrapped(address indexed unwrapper, address indexed recipientOfWrappedContents, uint256 indexed tokenIdOfWrappedToken)
func (_Multiwrap *MultiwrapFilterer) FilterTokensUnwrapped(opts *bind.FilterOpts, unwrapper []common.Address, recipientOfWrappedContents []common.Address, tokenIdOfWrappedToken []*big.Int) (*MultiwrapTokensUnwrappedIterator, error) {

	var unwrapperRule []interface{}
	for _, unwrapperItem := range unwrapper {
		unwrapperRule = append(unwrapperRule, unwrapperItem)
	}
	var recipientOfWrappedContentsRule []interface{}
	for _, recipientOfWrappedContentsItem := range recipientOfWrappedContents {
		recipientOfWrappedContentsRule = append(recipientOfWrappedContentsRule, recipientOfWrappedContentsItem)
	}
	var tokenIdOfWrappedTokenRule []interface{}
	for _, tokenIdOfWrappedTokenItem := range tokenIdOfWrappedToken {
		tokenIdOfWrappedTokenRule = append(tokenIdOfWrappedTokenRule, tokenIdOfWrappedTokenItem)
	}

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "TokensUnwrapped", unwrapperRule, recipientOfWrappedContentsRule, tokenIdOfWrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapTokensUnwrappedIterator{contract: _Multiwrap.contract, event: "TokensUnwrapped", logs: logs, sub: sub}, nil
}

// WatchTokensUnwrapped is a free log subscription operation binding the contract event 0xe9a1b206a63887b7a73ef37983c4105047ae75c9ac0920a1a256eea52f264a73.
//
// Solidity: event TokensUnwrapped(address indexed unwrapper, address indexed recipientOfWrappedContents, uint256 indexed tokenIdOfWrappedToken)
func (_Multiwrap *MultiwrapFilterer) WatchTokensUnwrapped(opts *bind.WatchOpts, sink chan<- *MultiwrapTokensUnwrapped, unwrapper []common.Address, recipientOfWrappedContents []common.Address, tokenIdOfWrappedToken []*big.Int) (event.Subscription, error) {

	var unwrapperRule []interface{}
	for _, unwrapperItem := range unwrapper {
		unwrapperRule = append(unwrapperRule, unwrapperItem)
	}
	var recipientOfWrappedContentsRule []interface{}
	for _, recipientOfWrappedContentsItem := range recipientOfWrappedContents {
		recipientOfWrappedContentsRule = append(recipientOfWrappedContentsRule, recipientOfWrappedContentsItem)
	}
	var tokenIdOfWrappedTokenRule []interface{}
	for _, tokenIdOfWrappedTokenItem := range tokenIdOfWrappedToken {
		tokenIdOfWrappedTokenRule = append(tokenIdOfWrappedTokenRule, tokenIdOfWrappedTokenItem)
	}

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "TokensUnwrapped", unwrapperRule, recipientOfWrappedContentsRule, tokenIdOfWrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapTokensUnwrapped)
				if err := _Multiwrap.contract.UnpackLog(event, "TokensUnwrapped", log); err != nil {
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

// ParseTokensUnwrapped is a log parse operation binding the contract event 0xe9a1b206a63887b7a73ef37983c4105047ae75c9ac0920a1a256eea52f264a73.
//
// Solidity: event TokensUnwrapped(address indexed unwrapper, address indexed recipientOfWrappedContents, uint256 indexed tokenIdOfWrappedToken)
func (_Multiwrap *MultiwrapFilterer) ParseTokensUnwrapped(log types.Log) (*MultiwrapTokensUnwrapped, error) {
	event := new(MultiwrapTokensUnwrapped)
	if err := _Multiwrap.contract.UnpackLog(event, "TokensUnwrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapTokensWrappedIterator is returned from FilterTokensWrapped and is used to iterate over the raw logs and unpacked data for TokensWrapped events raised by the Multiwrap contract.
type MultiwrapTokensWrappedIterator struct {
	Event *MultiwrapTokensWrapped // Event containing the contract specifics and raw log

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
func (it *MultiwrapTokensWrappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapTokensWrapped)
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
		it.Event = new(MultiwrapTokensWrapped)
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
func (it *MultiwrapTokensWrappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapTokensWrappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapTokensWrapped represents a TokensWrapped event raised by the Multiwrap contract.
type MultiwrapTokensWrapped struct {
	Wrapper                 common.Address
	RecipientOfWrappedToken common.Address
	TokenIdOfWrappedToken   *big.Int
	WrappedContents         []ITokenBundleToken
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterTokensWrapped is a free log retrieval operation binding the contract event 0xd37c0c0e6fd9c0d30d0b9fba1aa4f1123dedc7e681bb5b2c2d96298650032d4c.
//
// Solidity: event TokensWrapped(address indexed wrapper, address indexed recipientOfWrappedToken, uint256 indexed tokenIdOfWrappedToken, (address,uint8,uint256,uint256)[] wrappedContents)
func (_Multiwrap *MultiwrapFilterer) FilterTokensWrapped(opts *bind.FilterOpts, wrapper []common.Address, recipientOfWrappedToken []common.Address, tokenIdOfWrappedToken []*big.Int) (*MultiwrapTokensWrappedIterator, error) {

	var wrapperRule []interface{}
	for _, wrapperItem := range wrapper {
		wrapperRule = append(wrapperRule, wrapperItem)
	}
	var recipientOfWrappedTokenRule []interface{}
	for _, recipientOfWrappedTokenItem := range recipientOfWrappedToken {
		recipientOfWrappedTokenRule = append(recipientOfWrappedTokenRule, recipientOfWrappedTokenItem)
	}
	var tokenIdOfWrappedTokenRule []interface{}
	for _, tokenIdOfWrappedTokenItem := range tokenIdOfWrappedToken {
		tokenIdOfWrappedTokenRule = append(tokenIdOfWrappedTokenRule, tokenIdOfWrappedTokenItem)
	}

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "TokensWrapped", wrapperRule, recipientOfWrappedTokenRule, tokenIdOfWrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapTokensWrappedIterator{contract: _Multiwrap.contract, event: "TokensWrapped", logs: logs, sub: sub}, nil
}

// WatchTokensWrapped is a free log subscription operation binding the contract event 0xd37c0c0e6fd9c0d30d0b9fba1aa4f1123dedc7e681bb5b2c2d96298650032d4c.
//
// Solidity: event TokensWrapped(address indexed wrapper, address indexed recipientOfWrappedToken, uint256 indexed tokenIdOfWrappedToken, (address,uint8,uint256,uint256)[] wrappedContents)
func (_Multiwrap *MultiwrapFilterer) WatchTokensWrapped(opts *bind.WatchOpts, sink chan<- *MultiwrapTokensWrapped, wrapper []common.Address, recipientOfWrappedToken []common.Address, tokenIdOfWrappedToken []*big.Int) (event.Subscription, error) {

	var wrapperRule []interface{}
	for _, wrapperItem := range wrapper {
		wrapperRule = append(wrapperRule, wrapperItem)
	}
	var recipientOfWrappedTokenRule []interface{}
	for _, recipientOfWrappedTokenItem := range recipientOfWrappedToken {
		recipientOfWrappedTokenRule = append(recipientOfWrappedTokenRule, recipientOfWrappedTokenItem)
	}
	var tokenIdOfWrappedTokenRule []interface{}
	for _, tokenIdOfWrappedTokenItem := range tokenIdOfWrappedToken {
		tokenIdOfWrappedTokenRule = append(tokenIdOfWrappedTokenRule, tokenIdOfWrappedTokenItem)
	}

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "TokensWrapped", wrapperRule, recipientOfWrappedTokenRule, tokenIdOfWrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapTokensWrapped)
				if err := _Multiwrap.contract.UnpackLog(event, "TokensWrapped", log); err != nil {
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

// ParseTokensWrapped is a log parse operation binding the contract event 0xd37c0c0e6fd9c0d30d0b9fba1aa4f1123dedc7e681bb5b2c2d96298650032d4c.
//
// Solidity: event TokensWrapped(address indexed wrapper, address indexed recipientOfWrappedToken, uint256 indexed tokenIdOfWrappedToken, (address,uint8,uint256,uint256)[] wrappedContents)
func (_Multiwrap *MultiwrapFilterer) ParseTokensWrapped(log types.Log) (*MultiwrapTokensWrapped, error) {
	event := new(MultiwrapTokensWrapped)
	if err := _Multiwrap.contract.UnpackLog(event, "TokensWrapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiwrapTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Multiwrap contract.
type MultiwrapTransferIterator struct {
	Event *MultiwrapTransfer // Event containing the contract specifics and raw log

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
func (it *MultiwrapTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiwrapTransfer)
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
		it.Event = new(MultiwrapTransfer)
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
func (it *MultiwrapTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiwrapTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiwrapTransfer represents a Transfer event raised by the Multiwrap contract.
type MultiwrapTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Multiwrap *MultiwrapFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MultiwrapTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Multiwrap.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiwrapTransferIterator{contract: _Multiwrap.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Multiwrap *MultiwrapFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MultiwrapTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Multiwrap.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiwrapTransfer)
				if err := _Multiwrap.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Multiwrap *MultiwrapFilterer) ParseTransfer(log types.Log) (*MultiwrapTransfer, error) {
	event := new(MultiwrapTransfer)
	if err := _Multiwrap.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
