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

// NFTCollectionMetaData contains all meta data concerning the NFTCollection contract.
var NFTCollectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_controlCenter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceOfUnderlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountReceived\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sharesRedeemed\",\"type\":\"uint256\"}],\"name\":\"ERC20Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceOfUnderlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmountOfUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"ERC20WrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceOfUnderlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenIdOfUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sourceOfUnderlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenIdOfUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"ERC721WrappedToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"tokenURIs\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenSupplies\",\"type\":\"uint256[]\"}],\"name\":\"NativeTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"_nftURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_nftSupplies\",\"type\":\"uint256[]\"}],\"name\":\"createNativeTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_pack\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"_nftURIs\",\"type\":\"string[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_nftSupplies\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"_packURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilOpenStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilOpenEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nftsPerOpen\",\"type\":\"uint256\"}],\"name\":\"createPackAtomic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"erc20WrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underlyingTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"erc721WrappedTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"redeemERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"redeemERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_URI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_restrictedTransfer\",\"type\":\"bool\"}],\"name\":\"setRestrictedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"enumNFTCollection.UnderlyingType\",\"name\":\"underlyingType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfersRestricted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_numOfNftsToMint\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_nftURI\",\"type\":\"string\"}],\"name\":\"wrapERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_nftURI\",\"type\":\"string\"}],\"name\":\"wrapERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTCollectionABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTCollectionMetaData.ABI instead.
var NFTCollectionABI = NFTCollectionMetaData.ABI

// NFTCollection is an auto generated Go binding around an Ethereum contract.
type NFTCollection struct {
	NFTCollectionCaller     // Read-only binding to the contract
	NFTCollectionTransactor // Write-only binding to the contract
	NFTCollectionFilterer   // Log filterer for contract events
}

// NFTCollectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTCollectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTCollectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTCollectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTCollectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTCollectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTCollectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTCollectionSession struct {
	Contract     *NFTCollection    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTCollectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTCollectionCallerSession struct {
	Contract *NFTCollectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// NFTCollectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTCollectionTransactorSession struct {
	Contract     *NFTCollectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// NFTCollectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTCollectionRaw struct {
	Contract *NFTCollection // Generic contract binding to access the raw methods on
}

// NFTCollectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTCollectionCallerRaw struct {
	Contract *NFTCollectionCaller // Generic read-only contract binding to access the raw methods on
}

// NFTCollectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTCollectionTransactorRaw struct {
	Contract *NFTCollectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTCollection creates a new instance of NFTCollection, bound to a specific deployed contract.
func NewNFTCollection(address common.Address, backend bind.ContractBackend) (*NFTCollection, error) {
	contract, err := bindNFTCollection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTCollection{NFTCollectionCaller: NFTCollectionCaller{contract: contract}, NFTCollectionTransactor: NFTCollectionTransactor{contract: contract}, NFTCollectionFilterer: NFTCollectionFilterer{contract: contract}}, nil
}

// NewNFTCollectionCaller creates a new read-only instance of NFTCollection, bound to a specific deployed contract.
func NewNFTCollectionCaller(address common.Address, caller bind.ContractCaller) (*NFTCollectionCaller, error) {
	contract, err := bindNFTCollection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionCaller{contract: contract}, nil
}

// NewNFTCollectionTransactor creates a new write-only instance of NFTCollection, bound to a specific deployed contract.
func NewNFTCollectionTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTCollectionTransactor, error) {
	contract, err := bindNFTCollection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionTransactor{contract: contract}, nil
}

// NewNFTCollectionFilterer creates a new log filterer instance of NFTCollection, bound to a specific deployed contract.
func NewNFTCollectionFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTCollectionFilterer, error) {
	contract, err := bindNFTCollection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionFilterer{contract: contract}, nil
}

// bindNFTCollection binds a generic wrapper to an already deployed contract.
func bindNFTCollection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTCollectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTCollection *NFTCollectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTCollection.Contract.NFTCollectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTCollection *NFTCollectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTCollection.Contract.NFTCollectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTCollection *NFTCollectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTCollection.Contract.NFTCollectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTCollection *NFTCollectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTCollection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTCollection *NFTCollectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTCollection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTCollection *NFTCollectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTCollection.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NFTCollection.Contract.DEFAULTADMINROLE(&_NFTCollection.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NFTCollection.Contract.DEFAULTADMINROLE(&_NFTCollection.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionSession) MINTERROLE() ([32]byte, error) {
	return _NFTCollection.Contract.MINTERROLE(&_NFTCollection.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCallerSession) MINTERROLE() ([32]byte, error) {
	return _NFTCollection.Contract.MINTERROLE(&_NFTCollection.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionSession) PAUSERROLE() ([32]byte, error) {
	return _NFTCollection.Contract.PAUSERROLE(&_NFTCollection.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCallerSession) PAUSERROLE() ([32]byte, error) {
	return _NFTCollection.Contract.PAUSERROLE(&_NFTCollection.CallOpts)
}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCaller) TRANSFERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "TRANSFER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionSession) TRANSFERROLE() ([32]byte, error) {
	return _NFTCollection.Contract.TRANSFERROLE(&_NFTCollection.CallOpts)
}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_NFTCollection *NFTCollectionCallerSession) TRANSFERROLE() ([32]byte, error) {
	return _NFTCollection.Contract.TRANSFERROLE(&_NFTCollection.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_NFTCollection *NFTCollectionCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "_contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_NFTCollection *NFTCollectionSession) ContractURI() (string, error) {
	return _NFTCollection.Contract.ContractURI(&_NFTCollection.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_NFTCollection *NFTCollectionCallerSession) ContractURI() (string, error) {
	return _NFTCollection.Contract.ContractURI(&_NFTCollection.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_NFTCollection *NFTCollectionCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_NFTCollection *NFTCollectionSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _NFTCollection.Contract.BalanceOf(&_NFTCollection.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_NFTCollection *NFTCollectionCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _NFTCollection.Contract.BalanceOf(&_NFTCollection.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_NFTCollection *NFTCollectionCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_NFTCollection *NFTCollectionSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _NFTCollection.Contract.BalanceOfBatch(&_NFTCollection.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_NFTCollection *NFTCollectionCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _NFTCollection.Contract.BalanceOfBatch(&_NFTCollection.CallOpts, accounts, ids)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFTCollection *NFTCollectionCaller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFTCollection *NFTCollectionSession) InternalContractURI() (string, error) {
	return _NFTCollection.Contract.InternalContractURI(&_NFTCollection.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_NFTCollection *NFTCollectionCallerSession) InternalContractURI() (string, error) {
	return _NFTCollection.Contract.InternalContractURI(&_NFTCollection.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _nftId) view returns(address)
func (_NFTCollection *NFTCollectionCaller) Creator(opts *bind.CallOpts, _nftId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "creator", _nftId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _nftId) view returns(address)
func (_NFTCollection *NFTCollectionSession) Creator(_nftId *big.Int) (common.Address, error) {
	return _NFTCollection.Contract.Creator(&_NFTCollection.CallOpts, _nftId)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _nftId) view returns(address)
func (_NFTCollection *NFTCollectionCallerSession) Creator(_nftId *big.Int) (common.Address, error) {
	return _NFTCollection.Contract.Creator(&_NFTCollection.CallOpts, _nftId)
}

// Erc20WrappedTokens is a free data retrieval call binding the contract method 0x5cd9913d.
//
// Solidity: function erc20WrappedTokens(uint256 ) view returns(address source, uint256 shares, uint256 underlyingTokenAmount)
func (_NFTCollection *NFTCollectionCaller) Erc20WrappedTokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Source                common.Address
	Shares                *big.Int
	UnderlyingTokenAmount *big.Int
}, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "erc20WrappedTokens", arg0)

	outstruct := new(struct {
		Source                common.Address
		Shares                *big.Int
		UnderlyingTokenAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Source = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Shares = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UnderlyingTokenAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Erc20WrappedTokens is a free data retrieval call binding the contract method 0x5cd9913d.
//
// Solidity: function erc20WrappedTokens(uint256 ) view returns(address source, uint256 shares, uint256 underlyingTokenAmount)
func (_NFTCollection *NFTCollectionSession) Erc20WrappedTokens(arg0 *big.Int) (struct {
	Source                common.Address
	Shares                *big.Int
	UnderlyingTokenAmount *big.Int
}, error) {
	return _NFTCollection.Contract.Erc20WrappedTokens(&_NFTCollection.CallOpts, arg0)
}

// Erc20WrappedTokens is a free data retrieval call binding the contract method 0x5cd9913d.
//
// Solidity: function erc20WrappedTokens(uint256 ) view returns(address source, uint256 shares, uint256 underlyingTokenAmount)
func (_NFTCollection *NFTCollectionCallerSession) Erc20WrappedTokens(arg0 *big.Int) (struct {
	Source                common.Address
	Shares                *big.Int
	UnderlyingTokenAmount *big.Int
}, error) {
	return _NFTCollection.Contract.Erc20WrappedTokens(&_NFTCollection.CallOpts, arg0)
}

// Erc721WrappedTokens is a free data retrieval call binding the contract method 0x367a182b.
//
// Solidity: function erc721WrappedTokens(uint256 ) view returns(address source, uint256 tokenId)
func (_NFTCollection *NFTCollectionCaller) Erc721WrappedTokens(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Source  common.Address
	TokenId *big.Int
}, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "erc721WrappedTokens", arg0)

	outstruct := new(struct {
		Source  common.Address
		TokenId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Source = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Erc721WrappedTokens is a free data retrieval call binding the contract method 0x367a182b.
//
// Solidity: function erc721WrappedTokens(uint256 ) view returns(address source, uint256 tokenId)
func (_NFTCollection *NFTCollectionSession) Erc721WrappedTokens(arg0 *big.Int) (struct {
	Source  common.Address
	TokenId *big.Int
}, error) {
	return _NFTCollection.Contract.Erc721WrappedTokens(&_NFTCollection.CallOpts, arg0)
}

// Erc721WrappedTokens is a free data retrieval call binding the contract method 0x367a182b.
//
// Solidity: function erc721WrappedTokens(uint256 ) view returns(address source, uint256 tokenId)
func (_NFTCollection *NFTCollectionCallerSession) Erc721WrappedTokens(arg0 *big.Int) (struct {
	Source  common.Address
	TokenId *big.Int
}, error) {
	return _NFTCollection.Contract.Erc721WrappedTokens(&_NFTCollection.CallOpts, arg0)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFTCollection *NFTCollectionCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFTCollection *NFTCollectionSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NFTCollection.Contract.GetRoleAdmin(&_NFTCollection.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NFTCollection *NFTCollectionCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NFTCollection.Contract.GetRoleAdmin(&_NFTCollection.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NFTCollection *NFTCollectionCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NFTCollection *NFTCollectionSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NFTCollection.Contract.GetRoleMember(&_NFTCollection.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NFTCollection *NFTCollectionCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NFTCollection.Contract.GetRoleMember(&_NFTCollection.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NFTCollection *NFTCollectionCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NFTCollection *NFTCollectionSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NFTCollection.Contract.GetRoleMemberCount(&_NFTCollection.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NFTCollection *NFTCollectionCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NFTCollection.Contract.GetRoleMemberCount(&_NFTCollection.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFTCollection *NFTCollectionCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFTCollection *NFTCollectionSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NFTCollection.Contract.HasRole(&_NFTCollection.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NFTCollection *NFTCollectionCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NFTCollection.Contract.HasRole(&_NFTCollection.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_NFTCollection *NFTCollectionCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_NFTCollection *NFTCollectionSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _NFTCollection.Contract.IsApprovedForAll(&_NFTCollection.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_NFTCollection *NFTCollectionCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _NFTCollection.Contract.IsApprovedForAll(&_NFTCollection.CallOpts, account, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFTCollection *NFTCollectionCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFTCollection *NFTCollectionSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NFTCollection.Contract.IsTrustedForwarder(&_NFTCollection.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_NFTCollection *NFTCollectionCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _NFTCollection.Contract.IsTrustedForwarder(&_NFTCollection.CallOpts, forwarder)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_NFTCollection *NFTCollectionCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_NFTCollection *NFTCollectionSession) NextTokenId() (*big.Int, error) {
	return _NFTCollection.Contract.NextTokenId(&_NFTCollection.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_NFTCollection *NFTCollectionCallerSession) NextTokenId() (*big.Int, error) {
	return _NFTCollection.Contract.NextTokenId(&_NFTCollection.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFTCollection *NFTCollectionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFTCollection *NFTCollectionSession) Paused() (bool, error) {
	return _NFTCollection.Contract.Paused(&_NFTCollection.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NFTCollection *NFTCollectionCallerSession) Paused() (bool, error) {
	return _NFTCollection.Contract.Paused(&_NFTCollection.CallOpts)
}

// RoyaltyBps is a free data retrieval call binding the contract method 0xc63adb2b.
//
// Solidity: function royaltyBps() view returns(uint256)
func (_NFTCollection *NFTCollectionCaller) RoyaltyBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "royaltyBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBps is a free data retrieval call binding the contract method 0xc63adb2b.
//
// Solidity: function royaltyBps() view returns(uint256)
func (_NFTCollection *NFTCollectionSession) RoyaltyBps() (*big.Int, error) {
	return _NFTCollection.Contract.RoyaltyBps(&_NFTCollection.CallOpts)
}

// RoyaltyBps is a free data retrieval call binding the contract method 0xc63adb2b.
//
// Solidity: function royaltyBps() view returns(uint256)
func (_NFTCollection *NFTCollectionCallerSession) RoyaltyBps() (*big.Int, error) {
	return _NFTCollection.Contract.RoyaltyBps(&_NFTCollection.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFTCollection *NFTCollectionCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "royaltyInfo", arg0, salePrice)

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
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFTCollection *NFTCollectionSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _NFTCollection.Contract.RoyaltyInfo(&_NFTCollection.CallOpts, arg0, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_NFTCollection *NFTCollectionCallerSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _NFTCollection.Contract.RoyaltyInfo(&_NFTCollection.CallOpts, arg0, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTCollection *NFTCollectionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTCollection *NFTCollectionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTCollection.Contract.SupportsInterface(&_NFTCollection.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NFTCollection *NFTCollectionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NFTCollection.Contract.SupportsInterface(&_NFTCollection.CallOpts, interfaceId)
}

// TokenState is a free data retrieval call binding the contract method 0x9745cc3d.
//
// Solidity: function tokenState(uint256 ) view returns(address creator, string uri, uint8 underlyingType)
func (_NFTCollection *NFTCollectionCaller) TokenState(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Creator        common.Address
	Uri            string
	UnderlyingType uint8
}, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "tokenState", arg0)

	outstruct := new(struct {
		Creator        common.Address
		Uri            string
		UnderlyingType uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Creator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Uri = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.UnderlyingType = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// TokenState is a free data retrieval call binding the contract method 0x9745cc3d.
//
// Solidity: function tokenState(uint256 ) view returns(address creator, string uri, uint8 underlyingType)
func (_NFTCollection *NFTCollectionSession) TokenState(arg0 *big.Int) (struct {
	Creator        common.Address
	Uri            string
	UnderlyingType uint8
}, error) {
	return _NFTCollection.Contract.TokenState(&_NFTCollection.CallOpts, arg0)
}

// TokenState is a free data retrieval call binding the contract method 0x9745cc3d.
//
// Solidity: function tokenState(uint256 ) view returns(address creator, string uri, uint8 underlyingType)
func (_NFTCollection *NFTCollectionCallerSession) TokenState(arg0 *big.Int) (struct {
	Creator        common.Address
	Uri            string
	UnderlyingType uint8
}, error) {
	return _NFTCollection.Contract.TokenState(&_NFTCollection.CallOpts, arg0)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _nftId) view returns(string)
func (_NFTCollection *NFTCollectionCaller) TokenURI(opts *bind.CallOpts, _nftId *big.Int) (string, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "tokenURI", _nftId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _nftId) view returns(string)
func (_NFTCollection *NFTCollectionSession) TokenURI(_nftId *big.Int) (string, error) {
	return _NFTCollection.Contract.TokenURI(&_NFTCollection.CallOpts, _nftId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _nftId) view returns(string)
func (_NFTCollection *NFTCollectionCallerSession) TokenURI(_nftId *big.Int) (string, error) {
	return _NFTCollection.Contract.TokenURI(&_NFTCollection.CallOpts, _nftId)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_NFTCollection *NFTCollectionCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_NFTCollection *NFTCollectionSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _NFTCollection.Contract.TotalSupply(&_NFTCollection.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_NFTCollection *NFTCollectionCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _NFTCollection.Contract.TotalSupply(&_NFTCollection.CallOpts, id)
}

// TransfersRestricted is a free data retrieval call binding the contract method 0x8423df79.
//
// Solidity: function transfersRestricted() view returns(bool)
func (_NFTCollection *NFTCollectionCaller) TransfersRestricted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "transfersRestricted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersRestricted is a free data retrieval call binding the contract method 0x8423df79.
//
// Solidity: function transfersRestricted() view returns(bool)
func (_NFTCollection *NFTCollectionSession) TransfersRestricted() (bool, error) {
	return _NFTCollection.Contract.TransfersRestricted(&_NFTCollection.CallOpts)
}

// TransfersRestricted is a free data retrieval call binding the contract method 0x8423df79.
//
// Solidity: function transfersRestricted() view returns(bool)
func (_NFTCollection *NFTCollectionCallerSession) TransfersRestricted() (bool, error) {
	return _NFTCollection.Contract.TransfersRestricted(&_NFTCollection.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _nftId) view returns(string)
func (_NFTCollection *NFTCollectionCaller) Uri(opts *bind.CallOpts, _nftId *big.Int) (string, error) {
	var out []interface{}
	err := _NFTCollection.contract.Call(opts, &out, "uri", _nftId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _nftId) view returns(string)
func (_NFTCollection *NFTCollectionSession) Uri(_nftId *big.Int) (string, error) {
	return _NFTCollection.Contract.Uri(&_NFTCollection.CallOpts, _nftId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _nftId) view returns(string)
func (_NFTCollection *NFTCollectionCallerSession) Uri(_nftId *big.Int) (string, error) {
	return _NFTCollection.Contract.Uri(&_NFTCollection.CallOpts, _nftId)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_NFTCollection *NFTCollectionTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_NFTCollection *NFTCollectionSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.Burn(&_NFTCollection.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_NFTCollection *NFTCollectionTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.Burn(&_NFTCollection.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_NFTCollection *NFTCollectionTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_NFTCollection *NFTCollectionSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.BurnBatch(&_NFTCollection.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_NFTCollection *NFTCollectionTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.BurnBatch(&_NFTCollection.TransactOpts, account, ids, values)
}

// CreateNativeTokens is a paid mutator transaction binding the contract method 0x49a13fdf.
//
// Solidity: function createNativeTokens(string[] _nftURIs, uint256[] _nftSupplies) returns(uint256[] nftIds)
func (_NFTCollection *NFTCollectionTransactor) CreateNativeTokens(opts *bind.TransactOpts, _nftURIs []string, _nftSupplies []*big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "createNativeTokens", _nftURIs, _nftSupplies)
}

// CreateNativeTokens is a paid mutator transaction binding the contract method 0x49a13fdf.
//
// Solidity: function createNativeTokens(string[] _nftURIs, uint256[] _nftSupplies) returns(uint256[] nftIds)
func (_NFTCollection *NFTCollectionSession) CreateNativeTokens(_nftURIs []string, _nftSupplies []*big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.CreateNativeTokens(&_NFTCollection.TransactOpts, _nftURIs, _nftSupplies)
}

// CreateNativeTokens is a paid mutator transaction binding the contract method 0x49a13fdf.
//
// Solidity: function createNativeTokens(string[] _nftURIs, uint256[] _nftSupplies) returns(uint256[] nftIds)
func (_NFTCollection *NFTCollectionTransactorSession) CreateNativeTokens(_nftURIs []string, _nftSupplies []*big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.CreateNativeTokens(&_NFTCollection.TransactOpts, _nftURIs, _nftSupplies)
}

// CreatePackAtomic is a paid mutator transaction binding the contract method 0x788f1f5e.
//
// Solidity: function createPackAtomic(address _pack, string[] _nftURIs, uint256[] _nftSupplies, string _packURI, uint256 _secondsUntilOpenStart, uint256 _secondsUntilOpenEnd, uint256 _nftsPerOpen) returns()
func (_NFTCollection *NFTCollectionTransactor) CreatePackAtomic(opts *bind.TransactOpts, _pack common.Address, _nftURIs []string, _nftSupplies []*big.Int, _packURI string, _secondsUntilOpenStart *big.Int, _secondsUntilOpenEnd *big.Int, _nftsPerOpen *big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "createPackAtomic", _pack, _nftURIs, _nftSupplies, _packURI, _secondsUntilOpenStart, _secondsUntilOpenEnd, _nftsPerOpen)
}

// CreatePackAtomic is a paid mutator transaction binding the contract method 0x788f1f5e.
//
// Solidity: function createPackAtomic(address _pack, string[] _nftURIs, uint256[] _nftSupplies, string _packURI, uint256 _secondsUntilOpenStart, uint256 _secondsUntilOpenEnd, uint256 _nftsPerOpen) returns()
func (_NFTCollection *NFTCollectionSession) CreatePackAtomic(_pack common.Address, _nftURIs []string, _nftSupplies []*big.Int, _packURI string, _secondsUntilOpenStart *big.Int, _secondsUntilOpenEnd *big.Int, _nftsPerOpen *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.CreatePackAtomic(&_NFTCollection.TransactOpts, _pack, _nftURIs, _nftSupplies, _packURI, _secondsUntilOpenStart, _secondsUntilOpenEnd, _nftsPerOpen)
}

// CreatePackAtomic is a paid mutator transaction binding the contract method 0x788f1f5e.
//
// Solidity: function createPackAtomic(address _pack, string[] _nftURIs, uint256[] _nftSupplies, string _packURI, uint256 _secondsUntilOpenStart, uint256 _secondsUntilOpenEnd, uint256 _nftsPerOpen) returns()
func (_NFTCollection *NFTCollectionTransactorSession) CreatePackAtomic(_pack common.Address, _nftURIs []string, _nftSupplies []*big.Int, _packURI string, _secondsUntilOpenStart *big.Int, _secondsUntilOpenEnd *big.Int, _nftsPerOpen *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.CreatePackAtomic(&_NFTCollection.TransactOpts, _pack, _nftURIs, _nftSupplies, _packURI, _secondsUntilOpenStart, _secondsUntilOpenEnd, _nftsPerOpen)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.Contract.GrantRole(&_NFTCollection.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.Contract.GrantRole(&_NFTCollection.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "mint", to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFTCollection *NFTCollectionSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.Mint(&_NFTCollection.TransactOpts, to, id, amount, data)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.Mint(&_NFTCollection.TransactOpts, to, id, amount, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactor) MintBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "mintBatch", to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFTCollection *NFTCollectionSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.MintBatch(&_NFTCollection.TransactOpts, to, ids, amounts, data)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactorSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.MintBatch(&_NFTCollection.TransactOpts, to, ids, amounts, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.OnERC1155BatchReceived(&_NFTCollection.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.OnERC1155BatchReceived(&_NFTCollection.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.OnERC1155Received(&_NFTCollection.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.OnERC1155Received(&_NFTCollection.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.OnERC721Received(&_NFTCollection.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_NFTCollection *NFTCollectionTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.OnERC721Received(&_NFTCollection.TransactOpts, arg0, arg1, arg2, arg3)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFTCollection *NFTCollectionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFTCollection *NFTCollectionSession) Pause() (*types.Transaction, error) {
	return _NFTCollection.Contract.Pause(&_NFTCollection.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NFTCollection *NFTCollectionTransactorSession) Pause() (*types.Transaction, error) {
	return _NFTCollection.Contract.Pause(&_NFTCollection.TransactOpts)
}

// RedeemERC20 is a paid mutator transaction binding the contract method 0x7a61080b.
//
// Solidity: function redeemERC20(uint256 _nftId, uint256 _amount) returns()
func (_NFTCollection *NFTCollectionTransactor) RedeemERC20(opts *bind.TransactOpts, _nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "redeemERC20", _nftId, _amount)
}

// RedeemERC20 is a paid mutator transaction binding the contract method 0x7a61080b.
//
// Solidity: function redeemERC20(uint256 _nftId, uint256 _amount) returns()
func (_NFTCollection *NFTCollectionSession) RedeemERC20(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.RedeemERC20(&_NFTCollection.TransactOpts, _nftId, _amount)
}

// RedeemERC20 is a paid mutator transaction binding the contract method 0x7a61080b.
//
// Solidity: function redeemERC20(uint256 _nftId, uint256 _amount) returns()
func (_NFTCollection *NFTCollectionTransactorSession) RedeemERC20(_nftId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.RedeemERC20(&_NFTCollection.TransactOpts, _nftId, _amount)
}

// RedeemERC721 is a paid mutator transaction binding the contract method 0x93b778d4.
//
// Solidity: function redeemERC721(uint256 _nftId) returns()
func (_NFTCollection *NFTCollectionTransactor) RedeemERC721(opts *bind.TransactOpts, _nftId *big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "redeemERC721", _nftId)
}

// RedeemERC721 is a paid mutator transaction binding the contract method 0x93b778d4.
//
// Solidity: function redeemERC721(uint256 _nftId) returns()
func (_NFTCollection *NFTCollectionSession) RedeemERC721(_nftId *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.RedeemERC721(&_NFTCollection.TransactOpts, _nftId)
}

// RedeemERC721 is a paid mutator transaction binding the contract method 0x93b778d4.
//
// Solidity: function redeemERC721(uint256 _nftId) returns()
func (_NFTCollection *NFTCollectionTransactorSession) RedeemERC721(_nftId *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.RedeemERC721(&_NFTCollection.TransactOpts, _nftId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.Contract.RenounceRole(&_NFTCollection.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.Contract.RenounceRole(&_NFTCollection.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.Contract.RevokeRole(&_NFTCollection.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NFTCollection *NFTCollectionTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NFTCollection.Contract.RevokeRole(&_NFTCollection.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFTCollection *NFTCollectionSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.SafeBatchTransferFrom(&_NFTCollection.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.SafeBatchTransferFrom(&_NFTCollection.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFTCollection *NFTCollectionSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.SafeTransferFrom(&_NFTCollection.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_NFTCollection *NFTCollectionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _NFTCollection.Contract.SafeTransferFrom(&_NFTCollection.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTCollection *NFTCollectionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTCollection *NFTCollectionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetApprovalForAll(&_NFTCollection.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NFTCollection *NFTCollectionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetApprovalForAll(&_NFTCollection.TransactOpts, operator, approved)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_NFTCollection *NFTCollectionTransactor) SetContractURI(opts *bind.TransactOpts, _URI string) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "setContractURI", _URI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_NFTCollection *NFTCollectionSession) SetContractURI(_URI string) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetContractURI(&_NFTCollection.TransactOpts, _URI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_NFTCollection *NFTCollectionTransactorSession) SetContractURI(_URI string) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetContractURI(&_NFTCollection.TransactOpts, _URI)
}

// SetRestrictedTransfer is a paid mutator transaction binding the contract method 0x8ba448c2.
//
// Solidity: function setRestrictedTransfer(bool _restrictedTransfer) returns()
func (_NFTCollection *NFTCollectionTransactor) SetRestrictedTransfer(opts *bind.TransactOpts, _restrictedTransfer bool) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "setRestrictedTransfer", _restrictedTransfer)
}

// SetRestrictedTransfer is a paid mutator transaction binding the contract method 0x8ba448c2.
//
// Solidity: function setRestrictedTransfer(bool _restrictedTransfer) returns()
func (_NFTCollection *NFTCollectionSession) SetRestrictedTransfer(_restrictedTransfer bool) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetRestrictedTransfer(&_NFTCollection.TransactOpts, _restrictedTransfer)
}

// SetRestrictedTransfer is a paid mutator transaction binding the contract method 0x8ba448c2.
//
// Solidity: function setRestrictedTransfer(bool _restrictedTransfer) returns()
func (_NFTCollection *NFTCollectionTransactorSession) SetRestrictedTransfer(_restrictedTransfer bool) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetRestrictedTransfer(&_NFTCollection.TransactOpts, _restrictedTransfer)
}

// SetRoyaltyBps is a paid mutator transaction binding the contract method 0x1f72d831.
//
// Solidity: function setRoyaltyBps(uint256 _royaltyBps) returns()
func (_NFTCollection *NFTCollectionTransactor) SetRoyaltyBps(opts *bind.TransactOpts, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "setRoyaltyBps", _royaltyBps)
}

// SetRoyaltyBps is a paid mutator transaction binding the contract method 0x1f72d831.
//
// Solidity: function setRoyaltyBps(uint256 _royaltyBps) returns()
func (_NFTCollection *NFTCollectionSession) SetRoyaltyBps(_royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetRoyaltyBps(&_NFTCollection.TransactOpts, _royaltyBps)
}

// SetRoyaltyBps is a paid mutator transaction binding the contract method 0x1f72d831.
//
// Solidity: function setRoyaltyBps(uint256 _royaltyBps) returns()
func (_NFTCollection *NFTCollectionTransactorSession) SetRoyaltyBps(_royaltyBps *big.Int) (*types.Transaction, error) {
	return _NFTCollection.Contract.SetRoyaltyBps(&_NFTCollection.TransactOpts, _royaltyBps)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFTCollection *NFTCollectionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFTCollection *NFTCollectionSession) Unpause() (*types.Transaction, error) {
	return _NFTCollection.Contract.Unpause(&_NFTCollection.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NFTCollection *NFTCollectionTransactorSession) Unpause() (*types.Transaction, error) {
	return _NFTCollection.Contract.Unpause(&_NFTCollection.TransactOpts)
}

// WrapERC20 is a paid mutator transaction binding the contract method 0xdb884b0c.
//
// Solidity: function wrapERC20(address _tokenContract, uint256 _tokenAmount, uint256 _numOfNftsToMint, string _nftURI) returns()
func (_NFTCollection *NFTCollectionTransactor) WrapERC20(opts *bind.TransactOpts, _tokenContract common.Address, _tokenAmount *big.Int, _numOfNftsToMint *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "wrapERC20", _tokenContract, _tokenAmount, _numOfNftsToMint, _nftURI)
}

// WrapERC20 is a paid mutator transaction binding the contract method 0xdb884b0c.
//
// Solidity: function wrapERC20(address _tokenContract, uint256 _tokenAmount, uint256 _numOfNftsToMint, string _nftURI) returns()
func (_NFTCollection *NFTCollectionSession) WrapERC20(_tokenContract common.Address, _tokenAmount *big.Int, _numOfNftsToMint *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFTCollection.Contract.WrapERC20(&_NFTCollection.TransactOpts, _tokenContract, _tokenAmount, _numOfNftsToMint, _nftURI)
}

// WrapERC20 is a paid mutator transaction binding the contract method 0xdb884b0c.
//
// Solidity: function wrapERC20(address _tokenContract, uint256 _tokenAmount, uint256 _numOfNftsToMint, string _nftURI) returns()
func (_NFTCollection *NFTCollectionTransactorSession) WrapERC20(_tokenContract common.Address, _tokenAmount *big.Int, _numOfNftsToMint *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFTCollection.Contract.WrapERC20(&_NFTCollection.TransactOpts, _tokenContract, _tokenAmount, _numOfNftsToMint, _nftURI)
}

// WrapERC721 is a paid mutator transaction binding the contract method 0x090a3282.
//
// Solidity: function wrapERC721(address _nftContract, uint256 _tokenId, string _nftURI) returns()
func (_NFTCollection *NFTCollectionTransactor) WrapERC721(opts *bind.TransactOpts, _nftContract common.Address, _tokenId *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFTCollection.contract.Transact(opts, "wrapERC721", _nftContract, _tokenId, _nftURI)
}

// WrapERC721 is a paid mutator transaction binding the contract method 0x090a3282.
//
// Solidity: function wrapERC721(address _nftContract, uint256 _tokenId, string _nftURI) returns()
func (_NFTCollection *NFTCollectionSession) WrapERC721(_nftContract common.Address, _tokenId *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFTCollection.Contract.WrapERC721(&_NFTCollection.TransactOpts, _nftContract, _tokenId, _nftURI)
}

// WrapERC721 is a paid mutator transaction binding the contract method 0x090a3282.
//
// Solidity: function wrapERC721(address _nftContract, uint256 _tokenId, string _nftURI) returns()
func (_NFTCollection *NFTCollectionTransactorSession) WrapERC721(_nftContract common.Address, _tokenId *big.Int, _nftURI string) (*types.Transaction, error) {
	return _NFTCollection.Contract.WrapERC721(&_NFTCollection.TransactOpts, _nftContract, _tokenId, _nftURI)
}

// NFTCollectionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NFTCollection contract.
type NFTCollectionApprovalForAllIterator struct {
	Event *NFTCollectionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NFTCollectionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionApprovalForAll)
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
		it.Event = new(NFTCollectionApprovalForAll)
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
func (it *NFTCollectionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionApprovalForAll represents a ApprovalForAll event raised by the NFTCollection contract.
type NFTCollectionApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_NFTCollection *NFTCollectionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*NFTCollectionApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionApprovalForAllIterator{contract: _NFTCollection.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_NFTCollection *NFTCollectionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NFTCollectionApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionApprovalForAll)
				if err := _NFTCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseApprovalForAll(log types.Log) (*NFTCollectionApprovalForAll, error) {
	event := new(NFTCollectionApprovalForAll)
	if err := _NFTCollection.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionERC20RedeemedIterator is returned from FilterERC20Redeemed and is used to iterate over the raw logs and unpacked data for ERC20Redeemed events raised by the NFTCollection contract.
type NFTCollectionERC20RedeemedIterator struct {
	Event *NFTCollectionERC20Redeemed // Event containing the contract specifics and raw log

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
func (it *NFTCollectionERC20RedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionERC20Redeemed)
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
		it.Event = new(NFTCollectionERC20Redeemed)
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
func (it *NFTCollectionERC20RedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionERC20RedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionERC20Redeemed represents a ERC20Redeemed event raised by the NFTCollection contract.
type NFTCollectionERC20Redeemed struct {
	Redeemer            common.Address
	TokenId             *big.Int
	SourceOfUnderlying  common.Address
	TokenAmountReceived *big.Int
	SharesRedeemed      *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterERC20Redeemed is a free log retrieval operation binding the contract event 0xb80333a30b48beb3745280a5e4bb1b8e53f79569e454ec99ba13678dfbe728eb.
//
// Solidity: event ERC20Redeemed(address indexed redeemer, uint256 indexed tokenId, address indexed sourceOfUnderlying, uint256 tokenAmountReceived, uint256 sharesRedeemed)
func (_NFTCollection *NFTCollectionFilterer) FilterERC20Redeemed(opts *bind.FilterOpts, redeemer []common.Address, tokenId []*big.Int, sourceOfUnderlying []common.Address) (*NFTCollectionERC20RedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "ERC20Redeemed", redeemerRule, tokenIdRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionERC20RedeemedIterator{contract: _NFTCollection.contract, event: "ERC20Redeemed", logs: logs, sub: sub}, nil
}

// WatchERC20Redeemed is a free log subscription operation binding the contract event 0xb80333a30b48beb3745280a5e4bb1b8e53f79569e454ec99ba13678dfbe728eb.
//
// Solidity: event ERC20Redeemed(address indexed redeemer, uint256 indexed tokenId, address indexed sourceOfUnderlying, uint256 tokenAmountReceived, uint256 sharesRedeemed)
func (_NFTCollection *NFTCollectionFilterer) WatchERC20Redeemed(opts *bind.WatchOpts, sink chan<- *NFTCollectionERC20Redeemed, redeemer []common.Address, tokenId []*big.Int, sourceOfUnderlying []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "ERC20Redeemed", redeemerRule, tokenIdRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionERC20Redeemed)
				if err := _NFTCollection.contract.UnpackLog(event, "ERC20Redeemed", log); err != nil {
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

// ParseERC20Redeemed is a log parse operation binding the contract event 0xb80333a30b48beb3745280a5e4bb1b8e53f79569e454ec99ba13678dfbe728eb.
//
// Solidity: event ERC20Redeemed(address indexed redeemer, uint256 indexed tokenId, address indexed sourceOfUnderlying, uint256 tokenAmountReceived, uint256 sharesRedeemed)
func (_NFTCollection *NFTCollectionFilterer) ParseERC20Redeemed(log types.Log) (*NFTCollectionERC20Redeemed, error) {
	event := new(NFTCollectionERC20Redeemed)
	if err := _NFTCollection.contract.UnpackLog(event, "ERC20Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionERC20WrappedTokenIterator is returned from FilterERC20WrappedToken and is used to iterate over the raw logs and unpacked data for ERC20WrappedToken events raised by the NFTCollection contract.
type NFTCollectionERC20WrappedTokenIterator struct {
	Event *NFTCollectionERC20WrappedToken // Event containing the contract specifics and raw log

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
func (it *NFTCollectionERC20WrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionERC20WrappedToken)
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
		it.Event = new(NFTCollectionERC20WrappedToken)
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
func (it *NFTCollectionERC20WrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionERC20WrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionERC20WrappedToken represents a ERC20WrappedToken event raised by the NFTCollection contract.
type NFTCollectionERC20WrappedToken struct {
	Creator                 common.Address
	SourceOfUnderlying      common.Address
	TotalAmountOfUnderlying *big.Int
	Shares                  *big.Int
	TokenId                 *big.Int
	TokenURI                string
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterERC20WrappedToken is a free log retrieval operation binding the contract event 0x53af9a53d4a11cb382035aa13d6d8e24575569953176d87464ba37e7013a6aa4.
//
// Solidity: event ERC20WrappedToken(address indexed creator, address indexed sourceOfUnderlying, uint256 totalAmountOfUnderlying, uint256 shares, uint256 tokenId, string tokenURI)
func (_NFTCollection *NFTCollectionFilterer) FilterERC20WrappedToken(opts *bind.FilterOpts, creator []common.Address, sourceOfUnderlying []common.Address) (*NFTCollectionERC20WrappedTokenIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "ERC20WrappedToken", creatorRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionERC20WrappedTokenIterator{contract: _NFTCollection.contract, event: "ERC20WrappedToken", logs: logs, sub: sub}, nil
}

// WatchERC20WrappedToken is a free log subscription operation binding the contract event 0x53af9a53d4a11cb382035aa13d6d8e24575569953176d87464ba37e7013a6aa4.
//
// Solidity: event ERC20WrappedToken(address indexed creator, address indexed sourceOfUnderlying, uint256 totalAmountOfUnderlying, uint256 shares, uint256 tokenId, string tokenURI)
func (_NFTCollection *NFTCollectionFilterer) WatchERC20WrappedToken(opts *bind.WatchOpts, sink chan<- *NFTCollectionERC20WrappedToken, creator []common.Address, sourceOfUnderlying []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "ERC20WrappedToken", creatorRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionERC20WrappedToken)
				if err := _NFTCollection.contract.UnpackLog(event, "ERC20WrappedToken", log); err != nil {
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

// ParseERC20WrappedToken is a log parse operation binding the contract event 0x53af9a53d4a11cb382035aa13d6d8e24575569953176d87464ba37e7013a6aa4.
//
// Solidity: event ERC20WrappedToken(address indexed creator, address indexed sourceOfUnderlying, uint256 totalAmountOfUnderlying, uint256 shares, uint256 tokenId, string tokenURI)
func (_NFTCollection *NFTCollectionFilterer) ParseERC20WrappedToken(log types.Log) (*NFTCollectionERC20WrappedToken, error) {
	event := new(NFTCollectionERC20WrappedToken)
	if err := _NFTCollection.contract.UnpackLog(event, "ERC20WrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionERC721RedeemedIterator is returned from FilterERC721Redeemed and is used to iterate over the raw logs and unpacked data for ERC721Redeemed events raised by the NFTCollection contract.
type NFTCollectionERC721RedeemedIterator struct {
	Event *NFTCollectionERC721Redeemed // Event containing the contract specifics and raw log

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
func (it *NFTCollectionERC721RedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionERC721Redeemed)
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
		it.Event = new(NFTCollectionERC721Redeemed)
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
func (it *NFTCollectionERC721RedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionERC721RedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionERC721Redeemed represents a ERC721Redeemed event raised by the NFTCollection contract.
type NFTCollectionERC721Redeemed struct {
	Redeemer            common.Address
	SourceOfUnderlying  common.Address
	TokenIdOfUnderlying *big.Int
	TokenId             *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterERC721Redeemed is a free log retrieval operation binding the contract event 0xea7b88a04220a2292e142652e8e1a59cc00e7d27d8ad6b11da6cf5870e6ba866.
//
// Solidity: event ERC721Redeemed(address indexed redeemer, address indexed sourceOfUnderlying, uint256 tokenIdOfUnderlying, uint256 tokenId)
func (_NFTCollection *NFTCollectionFilterer) FilterERC721Redeemed(opts *bind.FilterOpts, redeemer []common.Address, sourceOfUnderlying []common.Address) (*NFTCollectionERC721RedeemedIterator, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "ERC721Redeemed", redeemerRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionERC721RedeemedIterator{contract: _NFTCollection.contract, event: "ERC721Redeemed", logs: logs, sub: sub}, nil
}

// WatchERC721Redeemed is a free log subscription operation binding the contract event 0xea7b88a04220a2292e142652e8e1a59cc00e7d27d8ad6b11da6cf5870e6ba866.
//
// Solidity: event ERC721Redeemed(address indexed redeemer, address indexed sourceOfUnderlying, uint256 tokenIdOfUnderlying, uint256 tokenId)
func (_NFTCollection *NFTCollectionFilterer) WatchERC721Redeemed(opts *bind.WatchOpts, sink chan<- *NFTCollectionERC721Redeemed, redeemer []common.Address, sourceOfUnderlying []common.Address) (event.Subscription, error) {

	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "ERC721Redeemed", redeemerRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionERC721Redeemed)
				if err := _NFTCollection.contract.UnpackLog(event, "ERC721Redeemed", log); err != nil {
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
// Solidity: event ERC721Redeemed(address indexed redeemer, address indexed sourceOfUnderlying, uint256 tokenIdOfUnderlying, uint256 tokenId)
func (_NFTCollection *NFTCollectionFilterer) ParseERC721Redeemed(log types.Log) (*NFTCollectionERC721Redeemed, error) {
	event := new(NFTCollectionERC721Redeemed)
	if err := _NFTCollection.contract.UnpackLog(event, "ERC721Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionERC721WrappedTokenIterator is returned from FilterERC721WrappedToken and is used to iterate over the raw logs and unpacked data for ERC721WrappedToken events raised by the NFTCollection contract.
type NFTCollectionERC721WrappedTokenIterator struct {
	Event *NFTCollectionERC721WrappedToken // Event containing the contract specifics and raw log

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
func (it *NFTCollectionERC721WrappedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionERC721WrappedToken)
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
		it.Event = new(NFTCollectionERC721WrappedToken)
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
func (it *NFTCollectionERC721WrappedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionERC721WrappedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionERC721WrappedToken represents a ERC721WrappedToken event raised by the NFTCollection contract.
type NFTCollectionERC721WrappedToken struct {
	Creator             common.Address
	SourceOfUnderlying  common.Address
	TokenIdOfUnderlying *big.Int
	TokenId             *big.Int
	TokenURI            string
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterERC721WrappedToken is a free log retrieval operation binding the contract event 0x797091bed1dcc535a592c3da1115fdb7467204ef9fb70943f0fc855ebe0a30f5.
//
// Solidity: event ERC721WrappedToken(address indexed creator, address indexed sourceOfUnderlying, uint256 tokenIdOfUnderlying, uint256 tokenId, string tokenURI)
func (_NFTCollection *NFTCollectionFilterer) FilterERC721WrappedToken(opts *bind.FilterOpts, creator []common.Address, sourceOfUnderlying []common.Address) (*NFTCollectionERC721WrappedTokenIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "ERC721WrappedToken", creatorRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionERC721WrappedTokenIterator{contract: _NFTCollection.contract, event: "ERC721WrappedToken", logs: logs, sub: sub}, nil
}

// WatchERC721WrappedToken is a free log subscription operation binding the contract event 0x797091bed1dcc535a592c3da1115fdb7467204ef9fb70943f0fc855ebe0a30f5.
//
// Solidity: event ERC721WrappedToken(address indexed creator, address indexed sourceOfUnderlying, uint256 tokenIdOfUnderlying, uint256 tokenId, string tokenURI)
func (_NFTCollection *NFTCollectionFilterer) WatchERC721WrappedToken(opts *bind.WatchOpts, sink chan<- *NFTCollectionERC721WrappedToken, creator []common.Address, sourceOfUnderlying []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var sourceOfUnderlyingRule []interface{}
	for _, sourceOfUnderlyingItem := range sourceOfUnderlying {
		sourceOfUnderlyingRule = append(sourceOfUnderlyingRule, sourceOfUnderlyingItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "ERC721WrappedToken", creatorRule, sourceOfUnderlyingRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionERC721WrappedToken)
				if err := _NFTCollection.contract.UnpackLog(event, "ERC721WrappedToken", log); err != nil {
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

// ParseERC721WrappedToken is a log parse operation binding the contract event 0x797091bed1dcc535a592c3da1115fdb7467204ef9fb70943f0fc855ebe0a30f5.
//
// Solidity: event ERC721WrappedToken(address indexed creator, address indexed sourceOfUnderlying, uint256 tokenIdOfUnderlying, uint256 tokenId, string tokenURI)
func (_NFTCollection *NFTCollectionFilterer) ParseERC721WrappedToken(log types.Log) (*NFTCollectionERC721WrappedToken, error) {
	event := new(NFTCollectionERC721WrappedToken)
	if err := _NFTCollection.contract.UnpackLog(event, "ERC721WrappedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionNativeTokensIterator is returned from FilterNativeTokens and is used to iterate over the raw logs and unpacked data for NativeTokens events raised by the NFTCollection contract.
type NFTCollectionNativeTokensIterator struct {
	Event *NFTCollectionNativeTokens // Event containing the contract specifics and raw log

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
func (it *NFTCollectionNativeTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionNativeTokens)
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
		it.Event = new(NFTCollectionNativeTokens)
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
func (it *NFTCollectionNativeTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionNativeTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionNativeTokens represents a NativeTokens event raised by the NFTCollection contract.
type NFTCollectionNativeTokens struct {
	Creator       common.Address
	TokenIds      []*big.Int
	TokenURIs     []string
	TokenSupplies []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNativeTokens is a free log retrieval operation binding the contract event 0xbea5a4b7cf9a3f7d8bf7c987614db5dcfdcd917a2cd5788a210c0ebdccfc5da4.
//
// Solidity: event NativeTokens(address indexed creator, uint256[] tokenIds, string[] tokenURIs, uint256[] tokenSupplies)
func (_NFTCollection *NFTCollectionFilterer) FilterNativeTokens(opts *bind.FilterOpts, creator []common.Address) (*NFTCollectionNativeTokensIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "NativeTokens", creatorRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionNativeTokensIterator{contract: _NFTCollection.contract, event: "NativeTokens", logs: logs, sub: sub}, nil
}

// WatchNativeTokens is a free log subscription operation binding the contract event 0xbea5a4b7cf9a3f7d8bf7c987614db5dcfdcd917a2cd5788a210c0ebdccfc5da4.
//
// Solidity: event NativeTokens(address indexed creator, uint256[] tokenIds, string[] tokenURIs, uint256[] tokenSupplies)
func (_NFTCollection *NFTCollectionFilterer) WatchNativeTokens(opts *bind.WatchOpts, sink chan<- *NFTCollectionNativeTokens, creator []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "NativeTokens", creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionNativeTokens)
				if err := _NFTCollection.contract.UnpackLog(event, "NativeTokens", log); err != nil {
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

// ParseNativeTokens is a log parse operation binding the contract event 0xbea5a4b7cf9a3f7d8bf7c987614db5dcfdcd917a2cd5788a210c0ebdccfc5da4.
//
// Solidity: event NativeTokens(address indexed creator, uint256[] tokenIds, string[] tokenURIs, uint256[] tokenSupplies)
func (_NFTCollection *NFTCollectionFilterer) ParseNativeTokens(log types.Log) (*NFTCollectionNativeTokens, error) {
	event := new(NFTCollectionNativeTokens)
	if err := _NFTCollection.contract.UnpackLog(event, "NativeTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NFTCollection contract.
type NFTCollectionPausedIterator struct {
	Event *NFTCollectionPaused // Event containing the contract specifics and raw log

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
func (it *NFTCollectionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionPaused)
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
		it.Event = new(NFTCollectionPaused)
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
func (it *NFTCollectionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionPaused represents a Paused event raised by the NFTCollection contract.
type NFTCollectionPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFTCollection *NFTCollectionFilterer) FilterPaused(opts *bind.FilterOpts) (*NFTCollectionPausedIterator, error) {

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NFTCollectionPausedIterator{contract: _NFTCollection.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NFTCollection *NFTCollectionFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NFTCollectionPaused) (event.Subscription, error) {

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionPaused)
				if err := _NFTCollection.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParsePaused(log types.Log) (*NFTCollectionPaused, error) {
	event := new(NFTCollectionPaused)
	if err := _NFTCollection.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NFTCollection contract.
type NFTCollectionRoleAdminChangedIterator struct {
	Event *NFTCollectionRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NFTCollectionRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionRoleAdminChanged)
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
		it.Event = new(NFTCollectionRoleAdminChanged)
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
func (it *NFTCollectionRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionRoleAdminChanged represents a RoleAdminChanged event raised by the NFTCollection contract.
type NFTCollectionRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFTCollection *NFTCollectionFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NFTCollectionRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionRoleAdminChangedIterator{contract: _NFTCollection.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NFTCollection *NFTCollectionFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NFTCollectionRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionRoleAdminChanged)
				if err := _NFTCollection.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseRoleAdminChanged(log types.Log) (*NFTCollectionRoleAdminChanged, error) {
	event := new(NFTCollectionRoleAdminChanged)
	if err := _NFTCollection.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NFTCollection contract.
type NFTCollectionRoleGrantedIterator struct {
	Event *NFTCollectionRoleGranted // Event containing the contract specifics and raw log

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
func (it *NFTCollectionRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionRoleGranted)
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
		it.Event = new(NFTCollectionRoleGranted)
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
func (it *NFTCollectionRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionRoleGranted represents a RoleGranted event raised by the NFTCollection contract.
type NFTCollectionRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTCollection *NFTCollectionFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NFTCollectionRoleGrantedIterator, error) {

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

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionRoleGrantedIterator{contract: _NFTCollection.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTCollection *NFTCollectionFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NFTCollectionRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionRoleGranted)
				if err := _NFTCollection.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseRoleGranted(log types.Log) (*NFTCollectionRoleGranted, error) {
	event := new(NFTCollectionRoleGranted)
	if err := _NFTCollection.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NFTCollection contract.
type NFTCollectionRoleRevokedIterator struct {
	Event *NFTCollectionRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NFTCollectionRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionRoleRevoked)
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
		it.Event = new(NFTCollectionRoleRevoked)
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
func (it *NFTCollectionRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionRoleRevoked represents a RoleRevoked event raised by the NFTCollection contract.
type NFTCollectionRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTCollection *NFTCollectionFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NFTCollectionRoleRevokedIterator, error) {

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

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionRoleRevokedIterator{contract: _NFTCollection.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NFTCollection *NFTCollectionFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NFTCollectionRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionRoleRevoked)
				if err := _NFTCollection.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseRoleRevoked(log types.Log) (*NFTCollectionRoleRevoked, error) {
	event := new(NFTCollectionRoleRevoked)
	if err := _NFTCollection.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionRoyaltyUpdatedIterator is returned from FilterRoyaltyUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyUpdated events raised by the NFTCollection contract.
type NFTCollectionRoyaltyUpdatedIterator struct {
	Event *NFTCollectionRoyaltyUpdated // Event containing the contract specifics and raw log

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
func (it *NFTCollectionRoyaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionRoyaltyUpdated)
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
		it.Event = new(NFTCollectionRoyaltyUpdated)
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
func (it *NFTCollectionRoyaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionRoyaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionRoyaltyUpdated represents a RoyaltyUpdated event raised by the NFTCollection contract.
type NFTCollectionRoyaltyUpdated struct {
	RoyaltyBps *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyUpdated is a free log retrieval operation binding the contract event 0x244ea8d7627f5a08f4299862bd5a45752842c183aee5b0fb0d1e4887bfa605b3.
//
// Solidity: event RoyaltyUpdated(uint256 royaltyBps)
func (_NFTCollection *NFTCollectionFilterer) FilterRoyaltyUpdated(opts *bind.FilterOpts) (*NFTCollectionRoyaltyUpdatedIterator, error) {

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &NFTCollectionRoyaltyUpdatedIterator{contract: _NFTCollection.contract, event: "RoyaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyUpdated is a free log subscription operation binding the contract event 0x244ea8d7627f5a08f4299862bd5a45752842c183aee5b0fb0d1e4887bfa605b3.
//
// Solidity: event RoyaltyUpdated(uint256 royaltyBps)
func (_NFTCollection *NFTCollectionFilterer) WatchRoyaltyUpdated(opts *bind.WatchOpts, sink chan<- *NFTCollectionRoyaltyUpdated) (event.Subscription, error) {

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionRoyaltyUpdated)
				if err := _NFTCollection.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
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

// ParseRoyaltyUpdated is a log parse operation binding the contract event 0x244ea8d7627f5a08f4299862bd5a45752842c183aee5b0fb0d1e4887bfa605b3.
//
// Solidity: event RoyaltyUpdated(uint256 royaltyBps)
func (_NFTCollection *NFTCollectionFilterer) ParseRoyaltyUpdated(log types.Log) (*NFTCollectionRoyaltyUpdated, error) {
	event := new(NFTCollectionRoyaltyUpdated)
	if err := _NFTCollection.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the NFTCollection contract.
type NFTCollectionTransferBatchIterator struct {
	Event *NFTCollectionTransferBatch // Event containing the contract specifics and raw log

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
func (it *NFTCollectionTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionTransferBatch)
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
		it.Event = new(NFTCollectionTransferBatch)
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
func (it *NFTCollectionTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionTransferBatch represents a TransferBatch event raised by the NFTCollection contract.
type NFTCollectionTransferBatch struct {
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
func (_NFTCollection *NFTCollectionFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*NFTCollectionTransferBatchIterator, error) {

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

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionTransferBatchIterator{contract: _NFTCollection.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_NFTCollection *NFTCollectionFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *NFTCollectionTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionTransferBatch)
				if err := _NFTCollection.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseTransferBatch(log types.Log) (*NFTCollectionTransferBatch, error) {
	event := new(NFTCollectionTransferBatch)
	if err := _NFTCollection.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the NFTCollection contract.
type NFTCollectionTransferSingleIterator struct {
	Event *NFTCollectionTransferSingle // Event containing the contract specifics and raw log

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
func (it *NFTCollectionTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionTransferSingle)
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
		it.Event = new(NFTCollectionTransferSingle)
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
func (it *NFTCollectionTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionTransferSingle represents a TransferSingle event raised by the NFTCollection contract.
type NFTCollectionTransferSingle struct {
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
func (_NFTCollection *NFTCollectionFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*NFTCollectionTransferSingleIterator, error) {

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

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionTransferSingleIterator{contract: _NFTCollection.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_NFTCollection *NFTCollectionFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *NFTCollectionTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionTransferSingle)
				if err := _NFTCollection.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseTransferSingle(log types.Log) (*NFTCollectionTransferSingle, error) {
	event := new(NFTCollectionTransferSingle)
	if err := _NFTCollection.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the NFTCollection contract.
type NFTCollectionURIIterator struct {
	Event *NFTCollectionURI // Event containing the contract specifics and raw log

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
func (it *NFTCollectionURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionURI)
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
		it.Event = new(NFTCollectionURI)
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
func (it *NFTCollectionURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionURI represents a URI event raised by the NFTCollection contract.
type NFTCollectionURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_NFTCollection *NFTCollectionFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*NFTCollectionURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &NFTCollectionURIIterator{contract: _NFTCollection.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_NFTCollection *NFTCollectionFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *NFTCollectionURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionURI)
				if err := _NFTCollection.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseURI(log types.Log) (*NFTCollectionURI, error) {
	event := new(NFTCollectionURI)
	if err := _NFTCollection.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTCollectionUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NFTCollection contract.
type NFTCollectionUnpausedIterator struct {
	Event *NFTCollectionUnpaused // Event containing the contract specifics and raw log

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
func (it *NFTCollectionUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTCollectionUnpaused)
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
		it.Event = new(NFTCollectionUnpaused)
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
func (it *NFTCollectionUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTCollectionUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTCollectionUnpaused represents a Unpaused event raised by the NFTCollection contract.
type NFTCollectionUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFTCollection *NFTCollectionFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NFTCollectionUnpausedIterator, error) {

	logs, sub, err := _NFTCollection.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NFTCollectionUnpausedIterator{contract: _NFTCollection.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NFTCollection *NFTCollectionFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NFTCollectionUnpaused) (event.Subscription, error) {

	logs, sub, err := _NFTCollection.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTCollectionUnpaused)
				if err := _NFTCollection.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NFTCollection *NFTCollectionFilterer) ParseUnpaused(log types.Log) (*NFTCollectionUnpaused, error) {
	event := new(NFTCollectionUnpaused)
	if err := _NFTCollection.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
