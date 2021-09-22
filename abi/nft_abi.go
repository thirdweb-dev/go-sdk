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

// NFTMetaData contains all meta data concerning the NFT contract.
var NFTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_controlCenter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftAmountRedeemed\",\"type\":\"uint256\"}],\"name\":\"ERC20Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftsMinted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nftURI\",\"type\":\"string\"}],\"name\":\"ERC20WrappedNfts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nativeNftTokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nativeNftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nativeNftURI\",\"type\":\"string\"}],\"name\":\"ERC721WrappedNft\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"nftURIs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nftSupplies\",\"type\":\"uint256[]\"}],\"name\":\"NativeNfts\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"NftRoyaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_nftURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_nftSupplies\",\"type\":\"uint256[]\"}],\"name\":\"createNativeNfts\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pack\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_nftURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_nftSupplies\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_packURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilOpenStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilOpenEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nftsPerOpen\",\"type\":\"uint256\"}],\"name\":\"createPackAtomic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"erc20WrappedNfts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"erc721WrappedNfts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"enumNFTCollection.UnderlyingType\",\"name\":\"underlyingType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftRoyaltyBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"redeemERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_URI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setNftRoyaltyBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numOfNftsToMint\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_nftURI\",\"type\":\"string\"}],\"name\":\"wrapERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_nftURI\",\"type\":\"string\"}],\"name\":\"wrapERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTMetaData.ABI instead.
var NFTABI = NFTMetaData.ABI

// NFT is an auto generated Go binding around an Ethereum contract.
type NFT struct {
	NFTCaller     // Read-only binding to the contract
	NFTTransactor // Write-only binding to the contract
	NFTFilterer   // Log filterer for contract events
}

// NFTCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTSession struct {
	Contract     *NFT              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCallerSession struct {
	Contract *NFTCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NFTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTTransactorSession struct {
	Contract     *NFTTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTRaw struct {
	Contract *NFT // Generic contract binding to access the raw methods on
}

// NFTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCallerRaw struct {
	Contract *NFTCaller // Generic read-only contract binding to access the raw methods on
}

// NFTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTTransactorRaw struct {
	Contract *NFTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFT creates a new instance of NFT, bound to a specific deployed contract.
func NewNFT(address common.Address, backend bind.ContractBackend) (*NFT, error) {
	contract, err := bindNFT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFT{NFTCaller: NFTCaller{contract: contract}, NFTTransactor: NFTTransactor{contract: contract}, NFTFilterer: NFTFilterer{contract: contract}}, nil
}

// NewNFTCaller creates a new read-only instance of NFT, bound to a specific deployed contract.
func NewNFTCaller(address common.Address, caller bind.ContractCaller) (*NFTCaller, error) {
	contract, err := bindNFT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCaller{contract: contract}, nil
}

// NewNFTTransactor creates a new write-only instance of NFT, bound to a specific deployed contract.
func NewNFTTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTTransactor, error) {
	contract, err := bindNFT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTTransactor{contract: contract}, nil
}

// NewNFTFilterer creates a new log filterer instance of NFT, bound to a specific deployed contract.
func NewNFTFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTFilterer, error) {
	contract, err := bindNFT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTFilterer{contract: contract}, nil
}

// bindNFT binds a generic wrapper to an already deployed contract.
func bindNFT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.NFTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.NFTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFT *NFTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFT *NFTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFT *NFTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFT.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFT *NFTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFT *NFTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NFT.Contract.DEFAULTADMINROLE(&_NFT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFT *NFTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NFT.Contract.DEFAULTADMINROLE(&_NFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NFT *NFTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NFT *NFTSession) MINTERROLE() ([32]byte, error) {
	return _NFT.Contract.MINTERROLE(&_NFT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NFT *NFTCallerSession) MINTERROLE() ([32]byte, error) {
	return _NFT.Contract.MINTERROLE(&_NFT.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NFT *NFTCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NFT *NFTSession) PAUSERROLE() ([32]byte, error) {
	return _NFT.Contract.PAUSERROLE(&_NFT.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NFT *NFTCallerSession) PAUSERROLE() ([32]byte, error) {
	return _NFT.Contract.PAUSERROLE(&_NFT.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_NFT *NFTCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "_contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_NFT *NFTSession) ContractURI() (string, error) {
	return _NFT.Contract.ContractURI(&_NFT.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_NFT *NFTCallerSession) ContractURI() (string, error) {
	return _NFT.Contract.ContractURI(&_NFT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_NFT *NFTCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_NFT *NFTSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _NFT.Contract.BalanceOf(&_NFT.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_NFT *NFTCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _NFT.Contract.BalanceOf(&_NFT.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_NFT *NFTCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_NFT *NFTSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _NFT.Contract.BalanceOfBatch(&_NFT.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_NFT *NFTCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _NFT.Contract.BalanceOfBatch(&_NFT.CallOpts, accounts, ids)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFT *NFTCaller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFT *NFTSession) InternalContractURI() (string, error) {
	return _NFT.Contract.InternalContractURI(&_NFT.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFT *NFTCallerSession) InternalContractURI() (string, error) {
	return _NFT.Contract.InternalContractURI(&_NFT.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _nftId) view returns(address)
func (_NFT *NFTCaller) Creator(opts *bind.CallOpts, _nftId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "creator", _nftId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _nftId) view returns(address)
func (_NFT *NFTSession) Creator(_nftId *big.Int) (common.Address, error) {
	return _NFT.Contract.Creator(&_NFT.CallOpts, _nftId)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _nftId) view returns(address)
func (_NFT *NFTCallerSession) Creator(_nftId *big.Int) (common.Address, error) {
	return _NFT.Contract.Creator(&_NFT.CallOpts, _nftId)
}

// Erc20WrappedNfts is a free data retrieval call binding the contract method 0x1ceb2dc2.
//
// Solidity: function erc20WrappedNfts(uint256 ) view returns(address tokenContract, uint256 shares, uint256 underlyingTokenAmount)
func (_NFT *NFTCaller) Erc20WrappedNfts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenContract         common.Address
	Shares                *big.Int
	UnderlyingTokenAmount *big.Int
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "erc20WrappedNfts", arg0)

	outstruct := new(struct {
		TokenContract         common.Address
		Shares                *big.Int
		UnderlyingTokenAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Shares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnderlyingTokenAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Erc20WrappedNfts is a free data retrieval call binding the contract method 0x1ceb2dc2.
//
// Solidity: function erc20WrappedNfts(uint256 ) view returns(address tokenContract, uint256 shares, uint256 underlyingTokenAmount)
func (_NFT *NFTSession) Erc20WrappedNfts(arg0 *big.Int) (struct {
	TokenContract         common.Address
	Shares                *big.Int
	UnderlyingTokenAmount *big.Int
}, error) {
	return _NFT.Contract.Erc20WrappedNfts(&_NFT.CallOpts, arg0)
}

// Erc20WrappedNfts is a free data retrieval call binding the contract method 0x1ceb2dc2.
//
// Solidity: function erc20WrappedNfts(uint256 ) view returns(address tokenContract, uint256 shares, uint256 underlyingTokenAmount)
func (_NFT *NFTCallerSession) Erc20WrappedNfts(arg0 *big.Int) (struct {
	TokenContract         common.Address
	Shares                *big.Int
	UnderlyingTokenAmount *big.Int
}, error) {
	return _NFT.Contract.Erc20WrappedNfts(&_NFT.CallOpts, arg0)
}

// Erc721WrappedNfts is a free data retrieval call binding the contract method 0xddfa3e22.
//
// Solidity: function erc721WrappedNfts(uint256 ) view returns(address nftContract, uint256 nftTokenId)
func (_NFT *NFTCaller) Erc721WrappedNfts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	NftContract common.Address
	NftTokenId  *big.Int
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "erc721WrappedNfts", arg0)

	outstruct := new(struct {
		NftContract common.Address
		NftTokenId  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NftContract = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftTokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Erc721WrappedNfts is a free data retrieval call binding the contract method 0xddfa3e22.
//
// Solidity: function erc721WrappedNfts(uint256 ) view returns(address nftContract, uint256 nftTokenId)
func (_NFT *NFTSession) Erc721WrappedNfts(arg0 *big.Int) (struct {
	NftContract common.Address
	NftTokenId  *big.Int
}, error) {
	return _NFT.Contract.Erc721WrappedNfts(&_NFT.CallOpts, arg0)
}

// Erc721WrappedNfts is a free data retrieval call binding the contract method 0xddfa3e22.
//
// Solidity: function erc721WrappedNfts(uint256 ) view returns(address nftContract, uint256 nftTokenId)
func (_NFT *NFTCallerSession) Erc721WrappedNfts(arg0 *big.Int) (struct {
	NftContract common.Address
	NftTokenId  *big.Int
}, error) {
	return _NFT.Contract.Erc721WrappedNfts(&_NFT.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFT *NFTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFT *NFTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NFT.Contract.GetRoleAdmin(&_NFT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFT *NFTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NFT.Contract.GetRoleAdmin(&_NFT.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NFT *NFTCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NFT *NFTSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NFT.Contract.GetRoleMember(&_NFT.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NFT *NFTCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NFT.Contract.GetRoleMember(&_NFT.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NFT *NFTCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NFT *NFTSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NFT.Contract.GetRoleMemberCount(&_NFT.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NFT *NFTCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NFT.Contract.GetRoleMemberCount(&_NFT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFT *NFTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFT *NFTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NFT.Contract.HasRole(&_NFT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFT *NFTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NFT.Contract.HasRole(&_NFT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_NFT *NFTCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_NFT *NFTSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _NFT.Contract.IsApprovedForAll(&_NFT.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_NFT *NFTCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _NFT.Contract.IsApprovedForAll(&_NFT.CallOpts, account, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFT *NFTCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFT *NFTSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NFT.Contract.IsTrustedForwarder(&_NFT.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFT *NFTCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NFT.Contract.IsTrustedForwarder(&_NFT.CallOpts, forwarder)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_NFT *NFTCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_NFT *NFTSession) NextTokenId() (*big.Int, error) {
	return _NFT.Contract.NextTokenId(&_NFT.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_NFT *NFTCallerSession) NextTokenId() (*big.Int, error) {
	return _NFT.Contract.NextTokenId(&_NFT.CallOpts)
}

// NftInfo is a free data retrieval call binding the contract method 0x1f8bc790.
//
// Solidity: function nftInfo(uint256 ) view returns(address creator, string uri, uint256 supply, uint8 underlyingType)
func (_NFT *NFTCaller) NftInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator        common.Address
	Uri            string
	Supply         *big.Int
	UnderlyingType uint8
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "nftInfo", arg0)

	outstruct := new(struct {
		Creator        common.Address
		Uri            string
		Supply         *big.Int
		UnderlyingType uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Uri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Supply = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnderlyingType = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// NftInfo is a free data retrieval call binding the contract method 0x1f8bc790.
//
// Solidity: function nftInfo(uint256 ) view returns(address creator, string uri, uint256 supply, uint8 underlyingType)
func (_NFT *NFTSession) NftInfo(arg0 *big.Int) (struct {
	Creator        common.Address
	Uri            string
	Supply         *big.Int
	UnderlyingType uint8
}, error) {
	return _NFT.Contract.NftInfo(&_NFT.CallOpts, arg0)
}

// NftInfo is a free data retrieval call binding the contract method 0x1f8bc790.
//
// Solidity: function nftInfo(uint256 ) view returns(address creator, string uri, uint256 supply, uint8 underlyingType)
func (_NFT *NFTCallerSession) NftInfo(arg0 *big.Int) (struct {
	Creator        common.Address
	Uri            string
	Supply         *big.Int
	UnderlyingType uint8
}, error) {
	return _NFT.Contract.NftInfo(&_NFT.CallOpts, arg0)
}

// NftRoyaltyBps is a free data retrieval call binding the contract method 0x45d62dbc.
//
// Solidity: function nftRoyaltyBps() view returns(uint256)
func (_NFT *NFTCaller) NftRoyaltyBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "nftRoyaltyBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftRoyaltyBps is a free data retrieval call binding the contract method 0x45d62dbc.
//
// Solidity: function nftRoyaltyBps() view returns(uint256)
func (_NFT *NFTSession) NftRoyaltyBps() (*big.Int, error) {
	return _NFT.Contract.NftRoyaltyBps(&_NFT.CallOpts)
}

// NftRoyaltyBps is a free data retrieval call binding the contract method 0x45d62dbc.
//
// Solidity: function nftRoyaltyBps() view returns(uint256)
func (_NFT *NFTCallerSession) NftRoyaltyBps() (*big.Int, error) {
	return _NFT.Contract.NftRoyaltyBps(&_NFT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFT *NFTCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFT *NFTSession) Paused() (bool, error) {
	return _NFT.Contract.Paused(&_NFT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFT *NFTCallerSession) Paused() (bool, error) {
	return _NFT.Contract.Paused(&_NFT.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFT *NFTCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_NFT *NFTSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _NFT.Contract.RoyaltyInfo(&_NFT.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFT *NFTCallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _NFT.Contract.RoyaltyInfo(&_NFT.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFT.Contract.SupportsInterface(&_NFT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFT *NFTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFT.Contract.SupportsInterface(&_NFT.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _nftId) view returns(string)
func (_NFT *NFTCaller) TokenURI(opts *bind.CallOpts, _nftId *big.Int) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "tokenURI", _nftId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _nftId) view returns(string)
func (_NFT *NFTSession) TokenURI(_nftId *big.Int) (string, error) {
	return _NFT.Contract.TokenURI(&_NFT.CallOpts, _nftId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _nftId) view returns(string)
func (_NFT *NFTCallerSession) TokenURI(_nftId *big.Int) (string, error) {
	return _NFT.Contract.TokenURI(&_NFT.CallOpts, _nftId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _nftId) view returns(string)
func (_NFT *NFTCaller) Uri(opts *bind.CallOpts, _nftId *big.Int) (string, error) {
	var out []interface{}
	err := _NFT.contract.Call(opts, &out, "uri", _nftId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _nftId) view returns(string)
func (_NFT *NFTSession) Uri(_nftId *big.Int) (string, error) {
	return _NFT.Contract.Uri(&_NFT.CallOpts, _nftId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _nftId) view returns(string)
func (_NFT *NFTCallerSession) Uri(_nftId *big.Int) (string, error) {
	return _NFT.Contract.Uri(&_NFT.CallOpts, _nftId)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_NFT *NFTTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_NFT *NFTSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Burn(&_NFT.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_NFT *NFTTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.Burn(&_NFT.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_NFT *NFTTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_NFT *NFTSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BurnBatch(&_NFT.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_NFT *NFTTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.BurnBatch(&_NFT.TransactOpts, account, ids, values)
}

// CreateNativeNfts is a paid mutator transaction binding the contract method 0x41466a95.
//
// Solidity: function createNativeNfts(string[] _nftURIs, uint256[] _nftSupplies) returns(uint256[] nftIds)
func (_NFT *NFTTransactor) CreateNativeNfts(opts *bind.TransactOpts, _nftURIs []string, _nftSupplies []*big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "createNativeNfts", _nftURIs, _nftSupplies)
}

// CreateNativeNfts is a paid mutator transaction binding the contract method 0x41466a95.
//
// Solidity: function createNativeNfts(string[] _nftURIs, uint256[] _nftSupplies) returns(uint256[] nftIds)
func (_NFT *NFTSession) CreateNativeNfts(_nftURIs []string, _nftSupplies []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.CreateNativeNfts(&_NFT.TransactOpts, _nftURIs, _nftSupplies)
}

// CreateNativeNfts is a paid mutator transaction binding the contract method 0x41466a95.
//
// Solidity: function createNativeNfts(string[] _nftURIs, uint256[] _nftSupplies) returns(uint256[] nftIds)
func (_NFT *NFTTransactorSession) CreateNativeNfts(_nftURIs []string, _nftSupplies []*big.Int) (*types.Transaction, error) {
	return _NFT.Contract.CreateNativeNfts(&_NFT.TransactOpts, _nftURIs, _nftSupplies)
}

// CreatePackAtomic is a paid mutator transaction binding the contract method 0x788f1f5e.
//
// Solidity: function createPackAtomic(address _pack, string[] _nftURIs, uint256[] _nftSupplies, string _packURI, uint256 _secondsUntilOpenStart, uint256 _secondsUntilOpenEnd, uint256 _nftsPerOpen) returns()
func (_NFT *NFTTransactor) CreatePackAtomic(opts *bind.TransactOpts, _pack common.Address, _nftURIs []string, _nftSupplies []*big.Int, _packURI string, _secondsUntilOpenStart *big.Int, _secondsUntilOpenEnd *big.Int, _nftsPerOpen *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "createPackAtomic", _pack, _nftURIs, _nftSupplies, _packURI, _secondsUntilOpenStart, _secondsUntilOpenEnd, _nftsPerOpen)
}

// CreatePackAtomic is a paid mutator transaction binding the contract method 0x788f1f5e.
//
// Solidity: function createPackAtomic(address _pack, string[] _nftURIs, uint256[] _nftSupplies, string _packURI, uint256 _secondsUntilOpenStart, uint256 _secondsUntilOpenEnd, uint256 _nftsPerOpen) returns()
func (_NFT *NFTSession) CreatePackAtomic(_pack common.Address, _nftURIs []string, _nftSupplies []*big.Int, _packURI string, _secondsUntilOpenStart *big.Int, _secondsUntilOpenEnd *big.Int, _nftsPerOpen *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.CreatePackAtomic(&_NFT.TransactOpts, _pack, _nftURIs, _nftSupplies, _packURI, _secondsUntilOpenStart, _secondsUntilOpenEnd, _nftsPerOpen)
}

// CreatePackAtomic is a paid mutator transaction binding the contract method 0x788f1f5e.
//
// Solidity: function createPackAtomic(address _pack, string[] _nftURIs, uint256[] _nftSupplies, string _packURI, uint256 _secondsUntilOpenStart, uint256 _secondsUntilOpenEnd, uint256 _nftsPerOpen) returns()
func (_NFT *NFTTransactorSession) CreatePackAtomic(_pack common.Address, _nftURIs []string, _nftSupplies []*big.Int, _packURI string, _secondsUntilOpenStart *big.Int, _secondsUntilOpenEnd *big.Int, _nftsPerOpen *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.CreatePackAtomic(&_NFT.TransactOpts, _pack, _nftURIs, _nftSupplies, _packURI, _secondsUntilOpenStart, _secondsUntilOpenEnd, _nftsPerOpen)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFT *NFTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFT *NFTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.GrantRole(&_NFT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFT *NFTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.GrantRole(&_NFT.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFT *NFTTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "mint", to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFT *NFTSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.Mint(&_NFT.TransactOpts, to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFT *NFTTransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.Mint(&_NFT.TransactOpts, to, id, amount, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFT *NFTTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "mintBatch", to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFT *NFTSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.MintBatch(&_NFT.TransactOpts, to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFT *NFTTransactorSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.MintBatch(&_NFT.TransactOpts, to, ids, amounts, data)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFT *NFTTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFT *NFTSession) Pause() (*types.Transaction, error) {
	return _NFT.Contract.Pause(&_NFT.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFT *NFTTransactorSession) Pause() (*types.Transaction, error) {
	return _NFT.Contract.Pause(&_NFT.TransactOpts)
}

// RedeemERC20 is a paid mutator transaction binding the contract method 0x7a61080b.
//
// Solidity: function redeemERC20(uint256 _nftId, uint256 _amount) returns()
func (_NFT *NFTTransactor) RedeemERC20(opts *bind.TransactOpts, _nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "redeemERC20", _nftId, _amount)
}

// RedeemERC20 is a paid mutator transaction binding the contract method 0x7a61080b.
//
// Solidity: function redeemERC20(uint256 _nftId, uint256 _amount) returns()
func (_NFT *NFTSession) RedeemERC20(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.RedeemERC20(&_NFT.TransactOpts, _nftId, _amount)
}

// RedeemERC20 is a paid mutator transaction binding the contract method 0x7a61080b.
//
// Solidity: function redeemERC20(uint256 _nftId, uint256 _amount) returns()
func (_NFT *NFTTransactorSession) RedeemERC20(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.RedeemERC20(&_NFT.TransactOpts, _nftId, _amount)
}

// RedeemERC721 is a paid mutator transaction binding the contract method 0x93b778d4.
//
// Solidity: function redeemERC721(uint256 _nftId) returns()
func (_NFT *NFTTransactor) RedeemERC721(opts *bind.TransactOpts, _nftId *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "redeemERC721", _nftId)
}

// RedeemERC721 is a paid mutator transaction binding the contract method 0x93b778d4.
//
// Solidity: function redeemERC721(uint256 _nftId) returns()
func (_NFT *NFTSession) RedeemERC721(_nftId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.RedeemERC721(&_NFT.TransactOpts, _nftId)
}

// RedeemERC721 is a paid mutator transaction binding the contract method 0x93b778d4.
//
// Solidity: function redeemERC721(uint256 _nftId) returns()
func (_NFT *NFTTransactorSession) RedeemERC721(_nftId *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.RedeemERC721(&_NFT.TransactOpts, _nftId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFT *NFTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFT *NFTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.RenounceRole(&_NFT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFT *NFTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.RenounceRole(&_NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFT *NFTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFT *NFTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.RevokeRole(&_NFT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFT *NFTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFT.Contract.RevokeRole(&_NFT.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFT *NFTTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFT *NFTSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeBatchTransferFrom(&_NFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFT *NFTTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeBatchTransferFrom(&_NFT.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFT *NFTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFT *NFTSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom(&_NFT.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFT *NFTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFT.Contract.SafeTransferFrom(&_NFT.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFT *NFTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFT *NFTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.Contract.SetApprovalForAll(&_NFT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFT *NFTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFT.Contract.SetApprovalForAll(&_NFT.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_NFT *NFTTransactor) SetContractURI(opts *bind.TransactOpts, _URI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setContractURI", _URI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_NFT *NFTSession) SetContractURI(_URI string) (*types.Transaction, error) {
	return _NFT.Contract.SetContractURI(&_NFT.TransactOpts, _URI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_NFT *NFTTransactorSession) SetContractURI(_URI string) (*types.Transaction, error) {
	return _NFT.Contract.SetContractURI(&_NFT.TransactOpts, _URI)
}

// SetNftRoyaltyBps is a paid mutator transaction binding the contract method 0x2f9c131a.
//
// Solidity: function setNftRoyaltyBps(uint256 _royaltyBps) returns()
func (_NFT *NFTTransactor) SetNftRoyaltyBps(opts *bind.TransactOpts, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "setNftRoyaltyBps", _royaltyBps)
}

// SetNftRoyaltyBps is a paid mutator transaction binding the contract method 0x2f9c131a.
//
// Solidity: function setNftRoyaltyBps(uint256 _royaltyBps) returns()
func (_NFT *NFTSession) SetNftRoyaltyBps(_royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SetNftRoyaltyBps(&_NFT.TransactOpts, _royaltyBps)
}

// SetNftRoyaltyBps is a paid mutator transaction binding the contract method 0x2f9c131a.
//
// Solidity: function setNftRoyaltyBps(uint256 _royaltyBps) returns()
func (_NFT *NFTTransactorSession) SetNftRoyaltyBps(_royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFT.Contract.SetNftRoyaltyBps(&_NFT.TransactOpts, _royaltyBps)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFT *NFTTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFT *NFTSession) Unpause() (*types.Transaction, error) {
	return _NFT.Contract.Unpause(&_NFT.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFT *NFTTransactorSession) Unpause() (*types.Transaction, error) {
	return _NFT.Contract.Unpause(&_NFT.TransactOpts)
}

// WrapERC20 is a paid mutator transaction binding the contract method 0xdb884b0c.
//
// Solidity: function wrapERC20(address _tokenContract, uint256 _tokenAmount, uint256 _numOfNftsToMint, string _nftURI) returns()
func (_NFT *NFTTransactor) WrapERC20(opts *bind.TransactOpts, _tokenContract common.Address, _tokenAmount *big.Int, _numOfNftsToMint *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "wrapERC20", _tokenContract, _tokenAmount, _numOfNftsToMint, _nftURI)
}

// WrapERC20 is a paid mutator transaction binding the contract method 0xdb884b0c.
//
// Solidity: function wrapERC20(address _tokenContract, uint256 _tokenAmount, uint256 _numOfNftsToMint, string _nftURI) returns()
func (_NFT *NFTSession) WrapERC20(_tokenContract common.Address, _tokenAmount *big.Int, _numOfNftsToMint *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFT.Contract.WrapERC20(&_NFT.TransactOpts, _tokenContract, _tokenAmount, _numOfNftsToMint, _nftURI)
}

// WrapERC20 is a paid mutator transaction binding the contract method 0xdb884b0c.
//
// Solidity: function wrapERC20(address _tokenContract, uint256 _tokenAmount, uint256 _numOfNftsToMint, string _nftURI) returns()
func (_NFT *NFTTransactorSession) WrapERC20(_tokenContract common.Address, _tokenAmount *big.Int, _numOfNftsToMint *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFT.Contract.WrapERC20(&_NFT.TransactOpts, _tokenContract, _tokenAmount, _numOfNftsToMint, _nftURI)
}

// WrapERC721 is a paid mutator transaction binding the contract method 0x090a3282.
//
// Solidity: function wrapERC721(address _nftContract, uint256 _tokenId, string _nftURI) returns()
func (_NFT *NFTTransactor) WrapERC721(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFT.contract.Transact(opts, "wrapERC721", _nftContract, _tokenId, _nftURI)
}

// WrapERC721 is a paid mutator transaction binding the contract method 0x090a3282.
//
// Solidity: function wrapERC721(address _nftContract, uint256 _tokenId, string _nftURI) returns()
func (_NFT *NFTSession) WrapERC721(_nftContract common.Address, _tokenId *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFT.Contract.WrapERC721(&_NFT.TransactOpts, _nftContract, _tokenId, _nftURI)
}

// WrapERC721 is a paid mutator transaction binding the contract method 0x090a3282.
//
// Solidity: function wrapERC721(address _nftContract, uint256 _tokenId, string _nftURI) returns()
func (_NFT *NFTTransactorSession) WrapERC721(_nftContract common.Address, _tokenId *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFT.Contract.WrapERC721(&_NFT.TransactOpts, _nftContract, _tokenId, _nftURI)
}

// NFTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFT contract.
type NFTApprovalForAllIterator struct {
	Event *NFTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NFTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTApprovalForAll)
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
		it.Event = new(NFTApprovalForAll)
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
func (it *NFTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTApprovalForAll represents a ApprovalForAll event raised by the NFT contract.
type NFTApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_NFT *NFTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*NFTApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTApprovalForAllIterator{contract: _NFT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_NFT *NFTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTApprovalForAll)
				if err := _NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_NFT *NFTFilterer) ParseApprovalForAll(log types.Log) (*NFTApprovalForAll, error) {
	event := new(NFTApprovalForAll)
	if err := _NFT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTERC20RedeemedIterator is returned from FilterERC20Redeemed and is used to iterate over the raw logs and unpacked data for ERC20Redeemed events raised by the NFT contract.
type NFTERC20RedeemedIterator struct {
	Event *NFTERC20Redeemed // Event containing the contract specifics and raw log

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
func (it *NFTERC20RedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTERC20Redeemed)
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
		it.Event = new(NFTERC20Redeemed)
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
func (it *NFTERC20RedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTERC20RedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTERC20Redeemed represents a ERC20Redeemed event raised by the NFT contract.
type NFTERC20Redeemed struct {
	Redeemer            common.Address
	TokenContract       common.Address
	TokenAmountReceived *big.Int
	NftAmountRedeemed   *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterERC20Redeemed is a free log retrieval operation binding the contract event 0xb7ca9e3aa58248cbaa6a5f83ca826cc57bf8ecd128fd88da420e294353c652f9.
//
// Solidity: event ERC20Redeemed(address indexed redeemer, address indexed tokenContract, uint256 tokenAmountReceived, uint256 nftAmountRedeemed)
func (_NFT *NFTFilterer) FilterERC20Redeemed(opts *bind.FilterOpts, redeemer []common.Address, tokenContract []common.Address) (*NFTERC20RedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ERC20Redeemed", redeemerRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTERC20RedeemedIterator{contract: _NFT.contract, event: "ERC20Redeemed", logs: logs, sub: sub}, nil
}

// WatchERC20Redeemed is a free log subscription operation binding the contract event 0xb7ca9e3aa58248cbaa6a5f83ca826cc57bf8ecd128fd88da420e294353c652f9.
//
// Solidity: event ERC20Redeemed(address indexed redeemer, address indexed tokenContract, uint256 tokenAmountReceived, uint256 nftAmountRedeemed)
func (_NFT *NFTFilterer) WatchERC20Redeemed(opts *bind.WatchOpts, sink chan<- *NFTERC20Redeemed, redeemer []common.Address, tokenContract []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ERC20Redeemed", redeemerRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTERC20Redeemed)
				if err := _NFT.contract.UnpackLog(event, "ERC20Redeemed", log); err != nil {
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

// ParseERC20Redeemed is a log parse operation binding the contract event 0xb7ca9e3aa58248cbaa6a5f83ca826cc57bf8ecd128fd88da420e294353c652f9.
//
// Solidity: event ERC20Redeemed(address indexed redeemer, address indexed tokenContract, uint256 tokenAmountReceived, uint256 nftAmountRedeemed)
func (_NFT *NFTFilterer) ParseERC20Redeemed(log types.Log) (*NFTERC20Redeemed, error) {
	event := new(NFTERC20Redeemed)
	if err := _NFT.contract.UnpackLog(event, "ERC20Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTERC20WrappedNftsIterator is returned from FilterERC20WrappedNfts and is used to iterate over the raw logs and unpacked data for ERC20WrappedNfts events raised by the NFT contract.
type NFTERC20WrappedNftsIterator struct {
	Event *NFTERC20WrappedNfts // Event containing the contract specifics and raw log

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
func (it *NFTERC20WrappedNftsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTERC20WrappedNfts)
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
		it.Event = new(NFTERC20WrappedNfts)
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
func (it *NFTERC20WrappedNftsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTERC20WrappedNftsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTERC20WrappedNfts represents a ERC20WrappedNfts event raised by the NFT contract.
type NFTERC20WrappedNfts struct {
	Creator       common.Address
	TokenContract common.Address
	TokenAmount   *big.Int
	NftsMinted    *big.Int
	NftURI        string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterERC20WrappedNfts is a free log retrieval operation binding the contract event 0xd8bffec4446d425694f73630f7822315baa62da5311e89f0d83335ef6a3917ba.
//
// Solidity: event ERC20WrappedNfts(address indexed creator, address indexed tokenContract, uint256 tokenAmount, uint256 nftsMinted, string nftURI)
func (_NFT *NFTFilterer) FilterERC20WrappedNfts(opts *bind.FilterOpts, creator []common.Address, tokenContract []common.Address) (*NFTERC20WrappedNftsIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ERC20WrappedNfts", creatorRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTERC20WrappedNftsIterator{contract: _NFT.contract, event: "ERC20WrappedNfts", logs: logs, sub: sub}, nil
}

// WatchERC20WrappedNfts is a free log subscription operation binding the contract event 0xd8bffec4446d425694f73630f7822315baa62da5311e89f0d83335ef6a3917ba.
//
// Solidity: event ERC20WrappedNfts(address indexed creator, address indexed tokenContract, uint256 tokenAmount, uint256 nftsMinted, string nftURI)
func (_NFT *NFTFilterer) WatchERC20WrappedNfts(opts *bind.WatchOpts, sink chan<- *NFTERC20WrappedNfts, creator []common.Address, tokenContract []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var tokenContractRule []interface{}
	for _, tokenContractItem := range tokenContract {
		tokenContractRule = append(tokenContractRule, tokenContractItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ERC20WrappedNfts", creatorRule, tokenContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTERC20WrappedNfts)
				if err := _NFT.contract.UnpackLog(event, "ERC20WrappedNfts", log); err != nil {
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

// ParseERC20WrappedNfts is a log parse operation binding the contract event 0xd8bffec4446d425694f73630f7822315baa62da5311e89f0d83335ef6a3917ba.
//
// Solidity: event ERC20WrappedNfts(address indexed creator, address indexed tokenContract, uint256 tokenAmount, uint256 nftsMinted, string nftURI)
func (_NFT *NFTFilterer) ParseERC20WrappedNfts(log types.Log) (*NFTERC20WrappedNfts, error) {
	event := new(NFTERC20WrappedNfts)
	if err := _NFT.contract.UnpackLog(event, "ERC20WrappedNfts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTERC721RedeemedIterator is returned from FilterERC721Redeemed and is used to iterate over the raw logs and unpacked data for ERC721Redeemed events raised by the NFT contract.
type NFTERC721RedeemedIterator struct {
	Event *NFTERC721Redeemed // Event containing the contract specifics and raw log

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
func (it *NFTERC721RedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTERC721Redeemed)
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
		it.Event = new(NFTERC721Redeemed)
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
func (it *NFTERC721RedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTERC721RedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTERC721Redeemed represents a ERC721Redeemed event raised by the NFT contract.
type NFTERC721Redeemed struct {
	Redeemer         common.Address
	NftContract      common.Address
	NftTokenId       *big.Int
	NativeNftTokenId *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721Redeemed is a free log retrieval operation binding the contract event 0xea7b88a04220a2292e142652e8e1a59cc00e7d27d8ad6b11da6cf5870e6ba866.
//
// Solidity: event ERC721Redeemed(address indexed redeemer, address indexed nftContract, uint256 nftTokenId, uint256 nativeNftTokenId)
func (_NFT *NFTFilterer) FilterERC721Redeemed(opts *bind.FilterOpts, redeemer []common.Address, nftContract []common.Address) (*NFTERC721RedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ERC721Redeemed", redeemerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTERC721RedeemedIterator{contract: _NFT.contract, event: "ERC721Redeemed", logs: logs, sub: sub}, nil
}

// WatchERC721Redeemed is a free log subscription operation binding the contract event 0xea7b88a04220a2292e142652e8e1a59cc00e7d27d8ad6b11da6cf5870e6ba866.
//
// Solidity: event ERC721Redeemed(address indexed redeemer, address indexed nftContract, uint256 nftTokenId, uint256 nativeNftTokenId)
func (_NFT *NFTFilterer) WatchERC721Redeemed(opts *bind.WatchOpts, sink chan<- *NFTERC721Redeemed, redeemer []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ERC721Redeemed", redeemerRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTERC721Redeemed)
				if err := _NFT.contract.UnpackLog(event, "ERC721Redeemed", log); err != nil {
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

// ParseERC721Redeemed is a log parse operation binding the contract event 0xea7b88a04220a2292e142652e8e1a59cc00e7d27d8ad6b11da6cf5870e6ba866.
//
// Solidity: event ERC721Redeemed(address indexed redeemer, address indexed nftContract, uint256 nftTokenId, uint256 nativeNftTokenId)
func (_NFT *NFTFilterer) ParseERC721Redeemed(log types.Log) (*NFTERC721Redeemed, error) {
	event := new(NFTERC721Redeemed)
	if err := _NFT.contract.UnpackLog(event, "ERC721Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTERC721WrappedNftIterator is returned from FilterERC721WrappedNft and is used to iterate over the raw logs and unpacked data for ERC721WrappedNft events raised by the NFT contract.
type NFTERC721WrappedNftIterator struct {
	Event *NFTERC721WrappedNft // Event containing the contract specifics and raw log

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
func (it *NFTERC721WrappedNftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTERC721WrappedNft)
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
		it.Event = new(NFTERC721WrappedNft)
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
func (it *NFTERC721WrappedNftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTERC721WrappedNftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTERC721WrappedNft represents a ERC721WrappedNft event raised by the NFT contract.
type NFTERC721WrappedNft struct {
	Creator          common.Address
	NftContract      common.Address
	NftTokenId       *big.Int
	NativeNftTokenId *big.Int
	NativeNftURI     string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterERC721WrappedNft is a free log retrieval operation binding the contract event 0xdae1779be741094a20f8287fd2c7a4d2d745eeae29c438867017066ae27a1bc4.
//
// Solidity: event ERC721WrappedNft(address indexed creator, address indexed nftContract, uint256 nftTokenId, uint256 nativeNftTokenId, string nativeNftURI)
func (_NFT *NFTFilterer) FilterERC721WrappedNft(opts *bind.FilterOpts, creator []common.Address, nftContract []common.Address) (*NFTERC721WrappedNftIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "ERC721WrappedNft", creatorRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return &NFTERC721WrappedNftIterator{contract: _NFT.contract, event: "ERC721WrappedNft", logs: logs, sub: sub}, nil
}

// WatchERC721WrappedNft is a free log subscription operation binding the contract event 0xdae1779be741094a20f8287fd2c7a4d2d745eeae29c438867017066ae27a1bc4.
//
// Solidity: event ERC721WrappedNft(address indexed creator, address indexed nftContract, uint256 nftTokenId, uint256 nativeNftTokenId, string nativeNftURI)
func (_NFT *NFTFilterer) WatchERC721WrappedNft(opts *bind.WatchOpts, sink chan<- *NFTERC721WrappedNft, creator []common.Address, nftContract []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "ERC721WrappedNft", creatorRule, nftContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTERC721WrappedNft)
				if err := _NFT.contract.UnpackLog(event, "ERC721WrappedNft", log); err != nil {
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

// ParseERC721WrappedNft is a log parse operation binding the contract event 0xdae1779be741094a20f8287fd2c7a4d2d745eeae29c438867017066ae27a1bc4.
//
// Solidity: event ERC721WrappedNft(address indexed creator, address indexed nftContract, uint256 nftTokenId, uint256 nativeNftTokenId, string nativeNftURI)
func (_NFT *NFTFilterer) ParseERC721WrappedNft(log types.Log) (*NFTERC721WrappedNft, error) {
	event := new(NFTERC721WrappedNft)
	if err := _NFT.contract.UnpackLog(event, "ERC721WrappedNft", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTNativeNftsIterator is returned from FilterNativeNfts and is used to iterate over the raw logs and unpacked data for NativeNfts events raised by the NFT contract.
type NFTNativeNftsIterator struct {
	Event *NFTNativeNfts // Event containing the contract specifics and raw log

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
func (it *NFTNativeNftsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTNativeNfts)
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
		it.Event = new(NFTNativeNfts)
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
func (it *NFTNativeNftsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTNativeNftsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTNativeNfts represents a NativeNfts event raised by the NFT contract.
type NFTNativeNfts struct {
	Creator     common.Address
	NftIds      []*big.Int
	NftURIs     []string
	NftSupplies []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNativeNfts is a free log retrieval operation binding the contract event 0x9ead818be0b43d79a78daaabe6c787d254ac1279ac99d9ff62b6b227804f7b33.
//
// Solidity: event NativeNfts(address indexed creator, uint256[] nftIds, string[] nftURIs, uint256[] nftSupplies)
func (_NFT *NFTFilterer) FilterNativeNfts(opts *bind.FilterOpts, creator []common.Address) (*NFTNativeNftsIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "NativeNfts", creatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTNativeNftsIterator{contract: _NFT.contract, event: "NativeNfts", logs: logs, sub: sub}, nil
}

// WatchNativeNfts is a free log subscription operation binding the contract event 0x9ead818be0b43d79a78daaabe6c787d254ac1279ac99d9ff62b6b227804f7b33.
//
// Solidity: event NativeNfts(address indexed creator, uint256[] nftIds, string[] nftURIs, uint256[] nftSupplies)
func (_NFT *NFTFilterer) WatchNativeNfts(opts *bind.WatchOpts, sink chan<- *NFTNativeNfts, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "NativeNfts", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTNativeNfts)
				if err := _NFT.contract.UnpackLog(event, "NativeNfts", log); err != nil {
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

// ParseNativeNfts is a log parse operation binding the contract event 0x9ead818be0b43d79a78daaabe6c787d254ac1279ac99d9ff62b6b227804f7b33.
//
// Solidity: event NativeNfts(address indexed creator, uint256[] nftIds, string[] nftURIs, uint256[] nftSupplies)
func (_NFT *NFTFilterer) ParseNativeNfts(log types.Log) (*NFTNativeNfts, error) {
	event := new(NFTNativeNfts)
	if err := _NFT.contract.UnpackLog(event, "NativeNfts", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTNftRoyaltyUpdatedIterator is returned from FilterNftRoyaltyUpdated and is used to iterate over the raw logs and unpacked data for NftRoyaltyUpdated events raised by the NFT contract.
type NFTNftRoyaltyUpdatedIterator struct {
	Event *NFTNftRoyaltyUpdated // Event containing the contract specifics and raw log

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
func (it *NFTNftRoyaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTNftRoyaltyUpdated)
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
		it.Event = new(NFTNftRoyaltyUpdated)
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
func (it *NFTNftRoyaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTNftRoyaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTNftRoyaltyUpdated represents a NftRoyaltyUpdated event raised by the NFT contract.
type NFTNftRoyaltyUpdated struct {
	RoyaltyBps *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNftRoyaltyUpdated is a free log retrieval operation binding the contract event 0xfe0145b1017176fc94d0dfab7c8631a2f2c4b287fc395f4b3e8bfec35e150dac.
//
// Solidity: event NftRoyaltyUpdated(uint256 royaltyBps)
func (_NFT *NFTFilterer) FilterNftRoyaltyUpdated(opts *bind.FilterOpts) (*NFTNftRoyaltyUpdatedIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "NftRoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTNftRoyaltyUpdatedIterator{contract: _NFT.contract, event: "NftRoyaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchNftRoyaltyUpdated is a free log subscription operation binding the contract event 0xfe0145b1017176fc94d0dfab7c8631a2f2c4b287fc395f4b3e8bfec35e150dac.
//
// Solidity: event NftRoyaltyUpdated(uint256 royaltyBps)
func (_NFT *NFTFilterer) WatchNftRoyaltyUpdated(opts *bind.WatchOpts, sink chan<- *NFTNftRoyaltyUpdated) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "NftRoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTNftRoyaltyUpdated)
				if err := _NFT.contract.UnpackLog(event, "NftRoyaltyUpdated", log); err != nil {
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

// ParseNftRoyaltyUpdated is a log parse operation binding the contract event 0xfe0145b1017176fc94d0dfab7c8631a2f2c4b287fc395f4b3e8bfec35e150dac.
//
// Solidity: event NftRoyaltyUpdated(uint256 royaltyBps)
func (_NFT *NFTFilterer) ParseNftRoyaltyUpdated(log types.Log) (*NFTNftRoyaltyUpdated, error) {
	event := new(NFTNftRoyaltyUpdated)
	if err := _NFT.contract.UnpackLog(event, "NftRoyaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NFT contract.
type NFTPausedIterator struct {
	Event *NFTPaused // Event containing the contract specifics and raw log

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
func (it *NFTPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTPaused)
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
		it.Event = new(NFTPaused)
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
func (it *NFTPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTPaused represents a Paused event raised by the NFT contract.
type NFTPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFT *NFTFilterer) FilterPaused(opts *bind.FilterOpts) (*NFTPausedIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NFTPausedIterator{contract: _NFT.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFT *NFTFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NFTPaused) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTPaused)
				if err := _NFT.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFT *NFTFilterer) ParsePaused(log types.Log) (*NFTPaused, error) {
	event := new(NFTPaused)
	if err := _NFT.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NFT contract.
type NFTRoleAdminChangedIterator struct {
	Event *NFTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NFTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTRoleAdminChanged)
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
		it.Event = new(NFTRoleAdminChanged)
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
func (it *NFTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTRoleAdminChanged represents a RoleAdminChanged event raised by the NFT contract.
type NFTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFT *NFTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NFTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NFT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NFTRoleAdminChangedIterator{contract: _NFT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFT *NFTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NFTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NFT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTRoleAdminChanged)
				if err := _NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NFT *NFTFilterer) ParseRoleAdminChanged(log types.Log) (*NFTRoleAdminChanged, error) {
	event := new(NFTRoleAdminChanged)
	if err := _NFT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NFT contract.
type NFTRoleGrantedIterator struct {
	Event *NFTRoleGranted // Event containing the contract specifics and raw log

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
func (it *NFTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTRoleGranted)
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
		it.Event = new(NFTRoleGranted)
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
func (it *NFTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTRoleGranted represents a RoleGranted event raised by the NFT contract.
type NFTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFT *NFTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NFTRoleGrantedIterator, error) {

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

	logs, sub, err := _NFT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NFTRoleGrantedIterator{contract: _NFT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFT *NFTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NFTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NFT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTRoleGranted)
				if err := _NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NFT *NFTFilterer) ParseRoleGranted(log types.Log) (*NFTRoleGranted, error) {
	event := new(NFTRoleGranted)
	if err := _NFT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NFT contract.
type NFTRoleRevokedIterator struct {
	Event *NFTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NFTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTRoleRevoked)
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
		it.Event = new(NFTRoleRevoked)
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
func (it *NFTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTRoleRevoked represents a RoleRevoked event raised by the NFT contract.
type NFTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFT *NFTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NFTRoleRevokedIterator, error) {

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

	logs, sub, err := _NFT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NFTRoleRevokedIterator{contract: _NFT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFT *NFTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NFTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NFT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTRoleRevoked)
				if err := _NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NFT *NFTFilterer) ParseRoleRevoked(log types.Log) (*NFTRoleRevoked, error) {
	event := new(NFTRoleRevoked)
	if err := _NFT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the NFT contract.
type NFTTransferBatchIterator struct {
	Event *NFTTransferBatch // Event containing the contract specifics and raw log

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
func (it *NFTTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTTransferBatch)
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
		it.Event = new(NFTTransferBatch)
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
func (it *NFTTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTTransferBatch represents a TransferBatch event raised by the NFT contract.
type NFTTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_NFT *NFTFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*NFTTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NFTTransferBatchIterator{contract: _NFT.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_NFT *NFTFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *NFTTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTTransferBatch)
				if err := _NFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_NFT *NFTFilterer) ParseTransferBatch(log types.Log) (*NFTTransferBatch, error) {
	event := new(NFTTransferBatch)
	if err := _NFT.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the NFT contract.
type NFTTransferSingleIterator struct {
	Event *NFTTransferSingle // Event containing the contract specifics and raw log

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
func (it *NFTTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTTransferSingle)
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
		it.Event = new(NFTTransferSingle)
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
func (it *NFTTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTTransferSingle represents a TransferSingle event raised by the NFT contract.
type NFTTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_NFT *NFTFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*NFTTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NFTTransferSingleIterator{contract: _NFT.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_NFT *NFTFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *NFTTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTTransferSingle)
				if err := _NFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_NFT *NFTFilterer) ParseTransferSingle(log types.Log) (*NFTTransferSingle, error) {
	event := new(NFTTransferSingle)
	if err := _NFT.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the NFT contract.
type NFTURIIterator struct {
	Event *NFTURI // Event containing the contract specifics and raw log

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
func (it *NFTURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTURI)
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
		it.Event = new(NFTURI)
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
func (it *NFTURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTURI represents a URI event raised by the NFT contract.
type NFTURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_NFT *NFTFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*NFTURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NFT.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &NFTURIIterator{contract: _NFT.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_NFT *NFTFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *NFTURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NFT.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTURI)
				if err := _NFT.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_NFT *NFTFilterer) ParseURI(log types.Log) (*NFTURI, error) {
	event := new(NFTURI)
	if err := _NFT.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NFT contract.
type NFTUnpausedIterator struct {
	Event *NFTUnpaused // Event containing the contract specifics and raw log

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
func (it *NFTUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTUnpaused)
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
		it.Event = new(NFTUnpaused)
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
func (it *NFTUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTUnpaused represents a Unpaused event raised by the NFT contract.
type NFTUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFT *NFTFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NFTUnpausedIterator, error) {

	logs, sub, err := _NFT.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NFTUnpausedIterator{contract: _NFT.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFT *NFTFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NFTUnpaused) (event.Subscription, error) {

	logs, sub, err := _NFT.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTUnpaused)
				if err := _NFT.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFT *NFTFilterer) ParseUnpaused(log types.Log) (*NFTUnpaused, error) {
	event := new(NFTUnpaused)
	if err := _NFT.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
