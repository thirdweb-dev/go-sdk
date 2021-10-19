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

// PackPackState is an auto generated low-level Go binding around an user-defined struct.
type PackPackState struct {
	Uri       string
	Creator   common.Address
	OpenStart *big.Int
}

// PackRewards is an auto generated low-level Go binding around an user-defined struct.
type PackRewards struct {
	Source         common.Address
	TokenIds       []*big.Int
	AmountsPacked  []*big.Int
	RewardsPerOpen *big.Int
}

// PackMetaData contains all meta data concerning the Pack contract.
var PackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_controlCenter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_linkToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_fees\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"packTotalSupply\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"openStart\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPack.PackState\",\"name\":\"packState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsPacked\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"rewardsPerOpen\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structPack.Rewards\",\"name\":\"rewards\",\"type\":\"tuple\"}],\"name\":\"PackCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"opener\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"rewardIds\",\"type\":\"uint256[]\"}],\"name\":\"PackOpenFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"opener\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"PackOpenRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"transferable\",\"type\":\"bool\"}],\"name\":\"RestrictedTransferUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_packId\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"currentRequestId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_packId\",\"type\":\"uint256\"}],\"name\":\"getPack\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"openStart\",\"type\":\"uint256\"}],\"internalType\":\"structPack.PackState\",\"name\":\"pack\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_packId\",\"type\":\"uint256\"}],\"name\":\"getPackWithRewards\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"openStart\",\"type\":\"uint256\"}],\"internalType\":\"structPack.PackState\",\"name\":\"pack\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"packTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsPacked\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"_ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_packId\",\"type\":\"uint256\"}],\"name\":\"openPack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"packs\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"openStart\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"randomnessRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"packId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"opener\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"rawFulfillRandomness\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"source\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rewardsPerOpen\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newFees\",\"type\":\"uint256\"}],\"name\":\"setChainlinkFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_restrictedTransfer\",\"type\":\"bool\"}],\"name\":\"setRestrictedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferLink\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transfersRestricted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vrfKeyHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PackABI is the input ABI used to generate the binding from.
// Deprecated: Use PackMetaData.ABI instead.
var PackABI = PackMetaData.ABI

// Pack is an auto generated Go binding around an Ethereum contract.
type Pack struct {
	PackCaller     // Read-only binding to the contract
	PackTransactor // Write-only binding to the contract
	PackFilterer   // Log filterer for contract events
}

// PackCaller is an auto generated read-only Go binding around an Ethereum contract.
type PackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PackSession struct {
	Contract     *Pack             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PackCallerSession struct {
	Contract *PackCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PackTransactorSession struct {
	Contract     *PackTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PackRaw is an auto generated low-level Go binding around an Ethereum contract.
type PackRaw struct {
	Contract *Pack // Generic contract binding to access the raw methods on
}

// PackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PackCallerRaw struct {
	Contract *PackCaller // Generic read-only contract binding to access the raw methods on
}

// PackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PackTransactorRaw struct {
	Contract *PackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPack creates a new instance of Pack, bound to a specific deployed contract.
func NewPack(address common.Address, backend bind.ContractBackend) (*Pack, error) {
	contract, err := bindPack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pack{PackCaller: PackCaller{contract: contract}, PackTransactor: PackTransactor{contract: contract}, PackFilterer: PackFilterer{contract: contract}}, nil
}

// NewPackCaller creates a new read-only instance of Pack, bound to a specific deployed contract.
func NewPackCaller(address common.Address, caller bind.ContractCaller) (*PackCaller, error) {
	contract, err := bindPack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PackCaller{contract: contract}, nil
}

// NewPackTransactor creates a new write-only instance of Pack, bound to a specific deployed contract.
func NewPackTransactor(address common.Address, transactor bind.ContractTransactor) (*PackTransactor, error) {
	contract, err := bindPack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PackTransactor{contract: contract}, nil
}

// NewPackFilterer creates a new log filterer instance of Pack, bound to a specific deployed contract.
func NewPackFilterer(address common.Address, filterer bind.ContractFilterer) (*PackFilterer, error) {
	contract, err := bindPack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PackFilterer{contract: contract}, nil
}

// bindPack binds a generic wrapper to an already deployed contract.
func bindPack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PackABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pack *PackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pack.Contract.PackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pack *PackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pack.Contract.PackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pack *PackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pack.Contract.PackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pack *PackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pack *PackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pack *PackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pack.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pack *PackCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pack *PackSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pack.Contract.DEFAULTADMINROLE(&_Pack.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Pack *PackCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Pack.Contract.DEFAULTADMINROLE(&_Pack.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Pack *PackCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Pack *PackSession) MINTERROLE() ([32]byte, error) {
	return _Pack.Contract.MINTERROLE(&_Pack.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Pack *PackCallerSession) MINTERROLE() ([32]byte, error) {
	return _Pack.Contract.MINTERROLE(&_Pack.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Pack *PackCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Pack *PackSession) PAUSERROLE() ([32]byte, error) {
	return _Pack.Contract.PAUSERROLE(&_Pack.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Pack *PackCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Pack.Contract.PAUSERROLE(&_Pack.CallOpts)
}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_Pack *PackCaller) TRANSFERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "TRANSFER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_Pack *PackSession) TRANSFERROLE() ([32]byte, error) {
	return _Pack.Contract.TRANSFERROLE(&_Pack.CallOpts)
}

// TRANSFERROLE is a free data retrieval call binding the contract method 0x206b60f9.
//
// Solidity: function TRANSFER_ROLE() view returns(bytes32)
func (_Pack *PackCallerSession) TRANSFERROLE() ([32]byte, error) {
	return _Pack.Contract.TRANSFERROLE(&_Pack.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Pack *PackCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "_contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Pack *PackSession) ContractURI() (string, error) {
	return _Pack.Contract.ContractURI(&_Pack.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Pack *PackCallerSession) ContractURI() (string, error) {
	return _Pack.Contract.ContractURI(&_Pack.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Pack *PackCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Pack *PackSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Pack.Contract.BalanceOf(&_Pack.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_Pack *PackCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _Pack.Contract.BalanceOf(&_Pack.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Pack *PackCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Pack *PackSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Pack.Contract.BalanceOfBatch(&_Pack.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_Pack *PackCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _Pack.Contract.BalanceOfBatch(&_Pack.CallOpts, accounts, ids)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Pack *PackCaller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Pack *PackSession) InternalContractURI() (string, error) {
	return _Pack.Contract.InternalContractURI(&_Pack.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Pack *PackCallerSession) InternalContractURI() (string, error) {
	return _Pack.Contract.InternalContractURI(&_Pack.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _packId) view returns(address)
func (_Pack *PackCaller) Creator(opts *bind.CallOpts, _packId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "creator", _packId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _packId) view returns(address)
func (_Pack *PackSession) Creator(_packId *big.Int) (common.Address, error) {
	return _Pack.Contract.Creator(&_Pack.CallOpts, _packId)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 _packId) view returns(address)
func (_Pack *PackCallerSession) Creator(_packId *big.Int) (common.Address, error) {
	return _Pack.Contract.Creator(&_Pack.CallOpts, _packId)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5318c48f.
//
// Solidity: function currentRequestId(uint256 , address ) view returns(bytes32)
func (_Pack *PackCaller) CurrentRequestId(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "currentRequestId", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5318c48f.
//
// Solidity: function currentRequestId(uint256 , address ) view returns(bytes32)
func (_Pack *PackSession) CurrentRequestId(arg0 *big.Int, arg1 common.Address) ([32]byte, error) {
	return _Pack.Contract.CurrentRequestId(&_Pack.CallOpts, arg0, arg1)
}

// CurrentRequestId is a free data retrieval call binding the contract method 0x5318c48f.
//
// Solidity: function currentRequestId(uint256 , address ) view returns(bytes32)
func (_Pack *PackCallerSession) CurrentRequestId(arg0 *big.Int, arg1 common.Address) ([32]byte, error) {
	return _Pack.Contract.CurrentRequestId(&_Pack.CallOpts, arg0, arg1)
}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 _packId) view returns((string,address,uint256) pack)
func (_Pack *PackCaller) GetPack(opts *bind.CallOpts, _packId *big.Int) (PackPackState, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "getPack", _packId)

	if err != nil {
		return *new(PackPackState), err
	}

	out0 := *abi.ConvertType(out[0], new(PackPackState)).(*PackPackState)

	return out0, err

}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 _packId) view returns((string,address,uint256) pack)
func (_Pack *PackSession) GetPack(_packId *big.Int) (PackPackState, error) {
	return _Pack.Contract.GetPack(&_Pack.CallOpts, _packId)
}

// GetPack is a free data retrieval call binding the contract method 0x895ec54c.
//
// Solidity: function getPack(uint256 _packId) view returns((string,address,uint256) pack)
func (_Pack *PackCallerSession) GetPack(_packId *big.Int) (PackPackState, error) {
	return _Pack.Contract.GetPack(&_Pack.CallOpts, _packId)
}

// GetPackWithRewards is a free data retrieval call binding the contract method 0x10e34e67.
//
// Solidity: function getPackWithRewards(uint256 _packId) view returns((string,address,uint256) pack, uint256 packTotalSupply, address source, uint256[] tokenIds, uint256[] amountsPacked)
func (_Pack *PackCaller) GetPackWithRewards(opts *bind.CallOpts, _packId *big.Int) (struct {
	Pack            PackPackState
	PackTotalSupply *big.Int
	Source          common.Address
	TokenIds        []*big.Int
	AmountsPacked   []*big.Int
}, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "getPackWithRewards", _packId)

	outstruct := new(struct {
		Pack            PackPackState
		PackTotalSupply *big.Int
		Source          common.Address
		TokenIds        []*big.Int
		AmountsPacked   []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pack = *abi.ConvertType(out[0], new(PackPackState)).(*PackPackState)
	outstruct.PackTotalSupply = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Source = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenIds = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)
	outstruct.AmountsPacked = *abi.ConvertType(out[4], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPackWithRewards is a free data retrieval call binding the contract method 0x10e34e67.
//
// Solidity: function getPackWithRewards(uint256 _packId) view returns((string,address,uint256) pack, uint256 packTotalSupply, address source, uint256[] tokenIds, uint256[] amountsPacked)
func (_Pack *PackSession) GetPackWithRewards(_packId *big.Int) (struct {
	Pack            PackPackState
	PackTotalSupply *big.Int
	Source          common.Address
	TokenIds        []*big.Int
	AmountsPacked   []*big.Int
}, error) {
	return _Pack.Contract.GetPackWithRewards(&_Pack.CallOpts, _packId)
}

// GetPackWithRewards is a free data retrieval call binding the contract method 0x10e34e67.
//
// Solidity: function getPackWithRewards(uint256 _packId) view returns((string,address,uint256) pack, uint256 packTotalSupply, address source, uint256[] tokenIds, uint256[] amountsPacked)
func (_Pack *PackCallerSession) GetPackWithRewards(_packId *big.Int) (struct {
	Pack            PackPackState
	PackTotalSupply *big.Int
	Source          common.Address
	TokenIds        []*big.Int
	AmountsPacked   []*big.Int
}, error) {
	return _Pack.Contract.GetPackWithRewards(&_Pack.CallOpts, _packId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pack *PackCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pack *PackSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pack.Contract.GetRoleAdmin(&_Pack.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Pack *PackCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Pack.Contract.GetRoleAdmin(&_Pack.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Pack *PackCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Pack *PackSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Pack.Contract.GetRoleMember(&_Pack.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Pack *PackCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Pack.Contract.GetRoleMember(&_Pack.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Pack *PackCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Pack *PackSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Pack.Contract.GetRoleMemberCount(&_Pack.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Pack *PackCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Pack.Contract.GetRoleMemberCount(&_Pack.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pack *PackCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pack *PackSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pack.Contract.HasRole(&_Pack.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Pack *PackCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Pack.Contract.HasRole(&_Pack.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Pack *PackCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Pack *PackSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Pack.Contract.IsApprovedForAll(&_Pack.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_Pack *PackCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _Pack.Contract.IsApprovedForAll(&_Pack.CallOpts, account, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Pack *PackCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Pack *PackSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Pack.Contract.IsTrustedForwarder(&_Pack.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Pack *PackCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Pack.Contract.IsTrustedForwarder(&_Pack.CallOpts, forwarder)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Pack *PackCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Pack *PackSession) NextTokenId() (*big.Int, error) {
	return _Pack.Contract.NextTokenId(&_Pack.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Pack *PackCallerSession) NextTokenId() (*big.Int, error) {
	return _Pack.Contract.NextTokenId(&_Pack.CallOpts)
}

// Packs is a free data retrieval call binding the contract method 0xb84c1392.
//
// Solidity: function packs(uint256 ) view returns(string uri, address creator, uint256 openStart)
func (_Pack *PackCaller) Packs(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Uri       string
	Creator   common.Address
	OpenStart *big.Int
}, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "packs", arg0)

	outstruct := new(struct {
		Uri       string
		Creator   common.Address
		OpenStart *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Uri = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Creator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.OpenStart = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Packs is a free data retrieval call binding the contract method 0xb84c1392.
//
// Solidity: function packs(uint256 ) view returns(string uri, address creator, uint256 openStart)
func (_Pack *PackSession) Packs(arg0 *big.Int) (struct {
	Uri       string
	Creator   common.Address
	OpenStart *big.Int
}, error) {
	return _Pack.Contract.Packs(&_Pack.CallOpts, arg0)
}

// Packs is a free data retrieval call binding the contract method 0xb84c1392.
//
// Solidity: function packs(uint256 ) view returns(string uri, address creator, uint256 openStart)
func (_Pack *PackCallerSession) Packs(arg0 *big.Int) (struct {
	Uri       string
	Creator   common.Address
	OpenStart *big.Int
}, error) {
	return _Pack.Contract.Packs(&_Pack.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pack *PackCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pack *PackSession) Paused() (bool, error) {
	return _Pack.Contract.Paused(&_Pack.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pack *PackCallerSession) Paused() (bool, error) {
	return _Pack.Contract.Paused(&_Pack.CallOpts)
}

// RandomnessRequests is a free data retrieval call binding the contract method 0x5e619b78.
//
// Solidity: function randomnessRequests(bytes32 ) view returns(uint256 packId, address opener)
func (_Pack *PackCaller) RandomnessRequests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	PackId *big.Int
	Opener common.Address
}, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "randomnessRequests", arg0)

	outstruct := new(struct {
		PackId *big.Int
		Opener common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PackId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Opener = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RandomnessRequests is a free data retrieval call binding the contract method 0x5e619b78.
//
// Solidity: function randomnessRequests(bytes32 ) view returns(uint256 packId, address opener)
func (_Pack *PackSession) RandomnessRequests(arg0 [32]byte) (struct {
	PackId *big.Int
	Opener common.Address
}, error) {
	return _Pack.Contract.RandomnessRequests(&_Pack.CallOpts, arg0)
}

// RandomnessRequests is a free data retrieval call binding the contract method 0x5e619b78.
//
// Solidity: function randomnessRequests(bytes32 ) view returns(uint256 packId, address opener)
func (_Pack *PackCallerSession) RandomnessRequests(arg0 [32]byte) (struct {
	PackId *big.Int
	Opener common.Address
}, error) {
	return _Pack.Contract.RandomnessRequests(&_Pack.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address source, uint256 rewardsPerOpen)
func (_Pack *PackCaller) Rewards(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Source         common.Address
	RewardsPerOpen *big.Int
}, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "rewards", arg0)

	outstruct := new(struct {
		Source         common.Address
		RewardsPerOpen *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Source = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RewardsPerOpen = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address source, uint256 rewardsPerOpen)
func (_Pack *PackSession) Rewards(arg0 *big.Int) (struct {
	Source         common.Address
	RewardsPerOpen *big.Int
}, error) {
	return _Pack.Contract.Rewards(&_Pack.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0xf301af42.
//
// Solidity: function rewards(uint256 ) view returns(address source, uint256 rewardsPerOpen)
func (_Pack *PackCallerSession) Rewards(arg0 *big.Int) (struct {
	Source         common.Address
	RewardsPerOpen *big.Int
}, error) {
	return _Pack.Contract.Rewards(&_Pack.CallOpts, arg0)
}

// RoyaltyBps is a free data retrieval call binding the contract method 0xc63adb2b.
//
// Solidity: function royaltyBps() view returns(uint256)
func (_Pack *PackCaller) RoyaltyBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "royaltyBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyBps is a free data retrieval call binding the contract method 0xc63adb2b.
//
// Solidity: function royaltyBps() view returns(uint256)
func (_Pack *PackSession) RoyaltyBps() (*big.Int, error) {
	return _Pack.Contract.RoyaltyBps(&_Pack.CallOpts)
}

// RoyaltyBps is a free data retrieval call binding the contract method 0xc63adb2b.
//
// Solidity: function royaltyBps() view returns(uint256)
func (_Pack *PackCallerSession) RoyaltyBps() (*big.Int, error) {
	return _Pack.Contract.RoyaltyBps(&_Pack.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Pack *PackCaller) RoyaltyInfo(opts *bind.CallOpts, arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "royaltyInfo", arg0, salePrice)

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
func (_Pack *PackSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Pack.Contract.RoyaltyInfo(&_Pack.CallOpts, arg0, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 , uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_Pack *PackCallerSession) RoyaltyInfo(arg0 *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _Pack.Contract.RoyaltyInfo(&_Pack.CallOpts, arg0, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pack *PackCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pack *PackSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pack.Contract.SupportsInterface(&_Pack.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Pack *PackCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Pack.Contract.SupportsInterface(&_Pack.CallOpts, interfaceId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _id) view returns(string)
func (_Pack *PackCaller) TokenURI(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "tokenURI", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _id) view returns(string)
func (_Pack *PackSession) TokenURI(_id *big.Int) (string, error) {
	return _Pack.Contract.TokenURI(&_Pack.CallOpts, _id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _id) view returns(string)
func (_Pack *PackCallerSession) TokenURI(_id *big.Int) (string, error) {
	return _Pack.Contract.TokenURI(&_Pack.CallOpts, _id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Pack *PackCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Pack *PackSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Pack.Contract.TotalSupply(&_Pack.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_Pack *PackCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _Pack.Contract.TotalSupply(&_Pack.CallOpts, id)
}

// TransfersRestricted is a free data retrieval call binding the contract method 0x8423df79.
//
// Solidity: function transfersRestricted() view returns(bool)
func (_Pack *PackCaller) TransfersRestricted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "transfersRestricted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransfersRestricted is a free data retrieval call binding the contract method 0x8423df79.
//
// Solidity: function transfersRestricted() view returns(bool)
func (_Pack *PackSession) TransfersRestricted() (bool, error) {
	return _Pack.Contract.TransfersRestricted(&_Pack.CallOpts)
}

// TransfersRestricted is a free data retrieval call binding the contract method 0x8423df79.
//
// Solidity: function transfersRestricted() view returns(bool)
func (_Pack *PackCallerSession) TransfersRestricted() (bool, error) {
	return _Pack.Contract.TransfersRestricted(&_Pack.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_Pack *PackCaller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_Pack *PackSession) Uri(_id *big.Int) (string, error) {
	return _Pack.Contract.Uri(&_Pack.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_Pack *PackCallerSession) Uri(_id *big.Int) (string, error) {
	return _Pack.Contract.Uri(&_Pack.CallOpts, _id)
}

// VrfFees is a free data retrieval call binding the contract method 0x10983e57.
//
// Solidity: function vrfFees() view returns(uint256)
func (_Pack *PackCaller) VrfFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "vrfFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VrfFees is a free data retrieval call binding the contract method 0x10983e57.
//
// Solidity: function vrfFees() view returns(uint256)
func (_Pack *PackSession) VrfFees() (*big.Int, error) {
	return _Pack.Contract.VrfFees(&_Pack.CallOpts)
}

// VrfFees is a free data retrieval call binding the contract method 0x10983e57.
//
// Solidity: function vrfFees() view returns(uint256)
func (_Pack *PackCallerSession) VrfFees() (*big.Int, error) {
	return _Pack.Contract.VrfFees(&_Pack.CallOpts)
}

// VrfKeyHash is a free data retrieval call binding the contract method 0x041d443e.
//
// Solidity: function vrfKeyHash() view returns(bytes32)
func (_Pack *PackCaller) VrfKeyHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pack.contract.Call(opts, &out, "vrfKeyHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VrfKeyHash is a free data retrieval call binding the contract method 0x041d443e.
//
// Solidity: function vrfKeyHash() view returns(bytes32)
func (_Pack *PackSession) VrfKeyHash() ([32]byte, error) {
	return _Pack.Contract.VrfKeyHash(&_Pack.CallOpts)
}

// VrfKeyHash is a free data retrieval call binding the contract method 0x041d443e.
//
// Solidity: function vrfKeyHash() view returns(bytes32)
func (_Pack *PackCallerSession) VrfKeyHash() ([32]byte, error) {
	return _Pack.Contract.VrfKeyHash(&_Pack.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Pack *PackTransactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Pack *PackSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.Burn(&_Pack.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_Pack *PackTransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.Burn(&_Pack.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Pack *PackTransactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Pack *PackSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Pack.Contract.BurnBatch(&_Pack.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_Pack *PackTransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _Pack.Contract.BurnBatch(&_Pack.TransactOpts, account, ids, values)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pack *PackTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pack *PackSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.Contract.GrantRole(&_Pack.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Pack *PackTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.Contract.GrantRole(&_Pack.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address , uint256 , uint256 , bytes ) returns()
func (_Pack *PackTransactor) Mint(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "mint", arg0, arg1, arg2, arg3)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address , uint256 , uint256 , bytes ) returns()
func (_Pack *PackSession) Mint(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.Contract.Mint(&_Pack.TransactOpts, arg0, arg1, arg2, arg3)
}

// Mint is a paid mutator transaction binding the contract method 0x731133e9.
//
// Solidity: function mint(address , uint256 , uint256 , bytes ) returns()
func (_Pack *PackTransactorSession) Mint(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.Contract.Mint(&_Pack.TransactOpts, arg0, arg1, arg2, arg3)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address , uint256[] , uint256[] , bytes ) returns()
func (_Pack *PackTransactor) MintBatch(opts *bind.TransactOpts, arg0 common.Address, arg1 []*big.Int, arg2 []*big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "mintBatch", arg0, arg1, arg2, arg3)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address , uint256[] , uint256[] , bytes ) returns()
func (_Pack *PackSession) MintBatch(arg0 common.Address, arg1 []*big.Int, arg2 []*big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.Contract.MintBatch(&_Pack.TransactOpts, arg0, arg1, arg2, arg3)
}

// MintBatch is a paid mutator transaction binding the contract method 0x1f7fdffa.
//
// Solidity: function mintBatch(address , uint256[] , uint256[] , bytes ) returns()
func (_Pack *PackTransactorSession) MintBatch(arg0 common.Address, arg1 []*big.Int, arg2 []*big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.Contract.MintBatch(&_Pack.TransactOpts, arg0, arg1, arg2, arg3)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pack *PackTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pack *PackSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pack.Contract.Multicall(&_Pack.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Pack *PackTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Pack.Contract.Multicall(&_Pack.TransactOpts, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address _operator, address , uint256[] _ids, uint256[] _values, bytes _data) returns(bytes4)
func (_Pack *PackTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, _operator common.Address, arg1 common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "onERC1155BatchReceived", _operator, arg1, _ids, _values, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address _operator, address , uint256[] _ids, uint256[] _values, bytes _data) returns(bytes4)
func (_Pack *PackSession) OnERC1155BatchReceived(_operator common.Address, arg1 common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _Pack.Contract.OnERC1155BatchReceived(&_Pack.TransactOpts, _operator, arg1, _ids, _values, _data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address _operator, address , uint256[] _ids, uint256[] _values, bytes _data) returns(bytes4)
func (_Pack *PackTransactorSession) OnERC1155BatchReceived(_operator common.Address, arg1 common.Address, _ids []*big.Int, _values []*big.Int, _data []byte) (*types.Transaction, error) {
	return _Pack.Contract.OnERC1155BatchReceived(&_Pack.TransactOpts, _operator, arg1, _ids, _values, _data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Pack *PackTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Pack *PackSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Pack.Contract.OnERC1155Received(&_Pack.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Pack *PackTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Pack.Contract.OnERC1155Received(&_Pack.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Pack *PackTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Pack *PackSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.Contract.OnERC721Received(&_Pack.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Pack *PackTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Pack.Contract.OnERC721Received(&_Pack.TransactOpts, arg0, arg1, arg2, arg3)
}

// OpenPack is a paid mutator transaction binding the contract method 0x50a88c7e.
//
// Solidity: function openPack(uint256 _packId) returns()
func (_Pack *PackTransactor) OpenPack(opts *bind.TransactOpts, _packId *big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "openPack", _packId)
}

// OpenPack is a paid mutator transaction binding the contract method 0x50a88c7e.
//
// Solidity: function openPack(uint256 _packId) returns()
func (_Pack *PackSession) OpenPack(_packId *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.OpenPack(&_Pack.TransactOpts, _packId)
}

// OpenPack is a paid mutator transaction binding the contract method 0x50a88c7e.
//
// Solidity: function openPack(uint256 _packId) returns()
func (_Pack *PackTransactorSession) OpenPack(_packId *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.OpenPack(&_Pack.TransactOpts, _packId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pack *PackTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pack *PackSession) Pause() (*types.Transaction, error) {
	return _Pack.Contract.Pause(&_Pack.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pack *PackTransactorSession) Pause() (*types.Transaction, error) {
	return _Pack.Contract.Pause(&_Pack.TransactOpts)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Pack *PackTransactor) RawFulfillRandomness(opts *bind.TransactOpts, requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "rawFulfillRandomness", requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Pack *PackSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.RawFulfillRandomness(&_Pack.TransactOpts, requestId, randomness)
}

// RawFulfillRandomness is a paid mutator transaction binding the contract method 0x94985ddd.
//
// Solidity: function rawFulfillRandomness(bytes32 requestId, uint256 randomness) returns()
func (_Pack *PackTransactorSession) RawFulfillRandomness(requestId [32]byte, randomness *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.RawFulfillRandomness(&_Pack.TransactOpts, requestId, randomness)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pack *PackTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pack *PackSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.Contract.RenounceRole(&_Pack.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Pack *PackTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.Contract.RenounceRole(&_Pack.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pack *PackTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pack *PackSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.Contract.RevokeRole(&_Pack.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Pack *PackTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Pack.Contract.RevokeRole(&_Pack.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Pack *PackTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Pack *PackSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Pack.Contract.SafeBatchTransferFrom(&_Pack.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_Pack *PackTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _Pack.Contract.SafeBatchTransferFrom(&_Pack.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Pack *PackTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Pack *PackSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Pack.Contract.SafeTransferFrom(&_Pack.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_Pack *PackTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _Pack.Contract.SafeTransferFrom(&_Pack.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Pack *PackTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Pack *PackSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Pack.Contract.SetApprovalForAll(&_Pack.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Pack *PackTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Pack.Contract.SetApprovalForAll(&_Pack.TransactOpts, operator, approved)
}

// SetChainlinkFees is a paid mutator transaction binding the contract method 0xb277ca45.
//
// Solidity: function setChainlinkFees(uint256 _newFees) returns()
func (_Pack *PackTransactor) SetChainlinkFees(opts *bind.TransactOpts, _newFees *big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "setChainlinkFees", _newFees)
}

// SetChainlinkFees is a paid mutator transaction binding the contract method 0xb277ca45.
//
// Solidity: function setChainlinkFees(uint256 _newFees) returns()
func (_Pack *PackSession) SetChainlinkFees(_newFees *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.SetChainlinkFees(&_Pack.TransactOpts, _newFees)
}

// SetChainlinkFees is a paid mutator transaction binding the contract method 0xb277ca45.
//
// Solidity: function setChainlinkFees(uint256 _newFees) returns()
func (_Pack *PackTransactorSession) SetChainlinkFees(_newFees *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.SetChainlinkFees(&_Pack.TransactOpts, _newFees)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Pack *PackTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Pack *PackSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Pack.Contract.SetContractURI(&_Pack.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Pack *PackTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Pack.Contract.SetContractURI(&_Pack.TransactOpts, _uri)
}

// SetRestrictedTransfer is a paid mutator transaction binding the contract method 0x8ba448c2.
//
// Solidity: function setRestrictedTransfer(bool _restrictedTransfer) returns()
func (_Pack *PackTransactor) SetRestrictedTransfer(opts *bind.TransactOpts, _restrictedTransfer bool) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "setRestrictedTransfer", _restrictedTransfer)
}

// SetRestrictedTransfer is a paid mutator transaction binding the contract method 0x8ba448c2.
//
// Solidity: function setRestrictedTransfer(bool _restrictedTransfer) returns()
func (_Pack *PackSession) SetRestrictedTransfer(_restrictedTransfer bool) (*types.Transaction, error) {
	return _Pack.Contract.SetRestrictedTransfer(&_Pack.TransactOpts, _restrictedTransfer)
}

// SetRestrictedTransfer is a paid mutator transaction binding the contract method 0x8ba448c2.
//
// Solidity: function setRestrictedTransfer(bool _restrictedTransfer) returns()
func (_Pack *PackTransactorSession) SetRestrictedTransfer(_restrictedTransfer bool) (*types.Transaction, error) {
	return _Pack.Contract.SetRestrictedTransfer(&_Pack.TransactOpts, _restrictedTransfer)
}

// SetRoyaltyBps is a paid mutator transaction binding the contract method 0x1f72d831.
//
// Solidity: function setRoyaltyBps(uint256 _royaltyBps) returns()
func (_Pack *PackTransactor) SetRoyaltyBps(opts *bind.TransactOpts, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "setRoyaltyBps", _royaltyBps)
}

// SetRoyaltyBps is a paid mutator transaction binding the contract method 0x1f72d831.
//
// Solidity: function setRoyaltyBps(uint256 _royaltyBps) returns()
func (_Pack *PackSession) SetRoyaltyBps(_royaltyBps *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.SetRoyaltyBps(&_Pack.TransactOpts, _royaltyBps)
}

// SetRoyaltyBps is a paid mutator transaction binding the contract method 0x1f72d831.
//
// Solidity: function setRoyaltyBps(uint256 _royaltyBps) returns()
func (_Pack *PackTransactorSession) SetRoyaltyBps(_royaltyBps *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.SetRoyaltyBps(&_Pack.TransactOpts, _royaltyBps)
}

// TransferLink is a paid mutator transaction binding the contract method 0xf99b1d68.
//
// Solidity: function transferLink(address _to, uint256 _amount) returns()
func (_Pack *PackTransactor) TransferLink(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "transferLink", _to, _amount)
}

// TransferLink is a paid mutator transaction binding the contract method 0xf99b1d68.
//
// Solidity: function transferLink(address _to, uint256 _amount) returns()
func (_Pack *PackSession) TransferLink(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.TransferLink(&_Pack.TransactOpts, _to, _amount)
}

// TransferLink is a paid mutator transaction binding the contract method 0xf99b1d68.
//
// Solidity: function transferLink(address _to, uint256 _amount) returns()
func (_Pack *PackTransactorSession) TransferLink(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Pack.Contract.TransferLink(&_Pack.TransactOpts, _to, _amount)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pack *PackTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pack.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pack *PackSession) Unpause() (*types.Transaction, error) {
	return _Pack.Contract.Unpause(&_Pack.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pack *PackTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pack.Contract.Unpause(&_Pack.TransactOpts)
}

// PackApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Pack contract.
type PackApprovalForAllIterator struct {
	Event *PackApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PackApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackApprovalForAll)
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
		it.Event = new(PackApprovalForAll)
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
func (it *PackApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackApprovalForAll represents a ApprovalForAll event raised by the Pack contract.
type PackApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Pack *PackFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*PackApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Pack.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PackApprovalForAllIterator{contract: _Pack.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_Pack *PackFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PackApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Pack.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackApprovalForAll)
				if err := _Pack.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Pack *PackFilterer) ParseApprovalForAll(log types.Log) (*PackApprovalForAll, error) {
	event := new(PackApprovalForAll)
	if err := _Pack.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackPackCreatedIterator is returned from FilterPackCreated and is used to iterate over the raw logs and unpacked data for PackCreated events raised by the Pack contract.
type PackPackCreatedIterator struct {
	Event *PackPackCreated // Event containing the contract specifics and raw log

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
func (it *PackPackCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackPackCreated)
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
		it.Event = new(PackPackCreated)
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
func (it *PackPackCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackPackCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackPackCreated represents a PackCreated event raised by the Pack contract.
type PackPackCreated struct {
	PackId          *big.Int
	RewardContract  common.Address
	Creator         common.Address
	PackTotalSupply *big.Int
	PackState       PackPackState
	Rewards         PackRewards
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPackCreated is a free log retrieval operation binding the contract event 0xd12628277f61a6acc1f051914551d13cdd8d117e3a5ce322869300fda84be712.
//
// Solidity: event PackCreated(uint256 indexed packId, address indexed rewardContract, address indexed creator, uint256 packTotalSupply, (string,address,uint256) packState, (address,uint256[],uint256[],uint256) rewards)
func (_Pack *PackFilterer) FilterPackCreated(opts *bind.FilterOpts, packId []*big.Int, rewardContract []common.Address, creator []common.Address) (*PackPackCreatedIterator, error) {

	var packIdRule []interface{}
	for _, packIdItem := range packId {
		packIdRule = append(packIdRule, packIdItem)
	}
	var rewardContractRule []interface{}
	for _, rewardContractItem := range rewardContract {
		rewardContractRule = append(rewardContractRule, rewardContractItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Pack.contract.FilterLogs(opts, "PackCreated", packIdRule, rewardContractRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &PackPackCreatedIterator{contract: _Pack.contract, event: "PackCreated", logs: logs, sub: sub}, nil
}

// WatchPackCreated is a free log subscription operation binding the contract event 0xd12628277f61a6acc1f051914551d13cdd8d117e3a5ce322869300fda84be712.
//
// Solidity: event PackCreated(uint256 indexed packId, address indexed rewardContract, address indexed creator, uint256 packTotalSupply, (string,address,uint256) packState, (address,uint256[],uint256[],uint256) rewards)
func (_Pack *PackFilterer) WatchPackCreated(opts *bind.WatchOpts, sink chan<- *PackPackCreated, packId []*big.Int, rewardContract []common.Address, creator []common.Address) (event.Subscription, error) {

	var packIdRule []interface{}
	for _, packIdItem := range packId {
		packIdRule = append(packIdRule, packIdItem)
	}
	var rewardContractRule []interface{}
	for _, rewardContractItem := range rewardContract {
		rewardContractRule = append(rewardContractRule, rewardContractItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Pack.contract.WatchLogs(opts, "PackCreated", packIdRule, rewardContractRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackPackCreated)
				if err := _Pack.contract.UnpackLog(event, "PackCreated", log); err != nil {
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

// ParsePackCreated is a log parse operation binding the contract event 0xd12628277f61a6acc1f051914551d13cdd8d117e3a5ce322869300fda84be712.
//
// Solidity: event PackCreated(uint256 indexed packId, address indexed rewardContract, address indexed creator, uint256 packTotalSupply, (string,address,uint256) packState, (address,uint256[],uint256[],uint256) rewards)
func (_Pack *PackFilterer) ParsePackCreated(log types.Log) (*PackPackCreated, error) {
	event := new(PackPackCreated)
	if err := _Pack.contract.UnpackLog(event, "PackCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackPackOpenFulfilledIterator is returned from FilterPackOpenFulfilled and is used to iterate over the raw logs and unpacked data for PackOpenFulfilled events raised by the Pack contract.
type PackPackOpenFulfilledIterator struct {
	Event *PackPackOpenFulfilled // Event containing the contract specifics and raw log

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
func (it *PackPackOpenFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackPackOpenFulfilled)
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
		it.Event = new(PackPackOpenFulfilled)
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
func (it *PackPackOpenFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackPackOpenFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackPackOpenFulfilled represents a PackOpenFulfilled event raised by the Pack contract.
type PackPackOpenFulfilled struct {
	PackId         *big.Int
	Opener         common.Address
	RequestId      [32]byte
	RewardContract common.Address
	RewardIds      []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPackOpenFulfilled is a free log retrieval operation binding the contract event 0x9cae8e743935573869e2fe7223782448d71211235c93076166106782de1c234f.
//
// Solidity: event PackOpenFulfilled(uint256 indexed packId, address indexed opener, bytes32 requestId, address indexed rewardContract, uint256[] rewardIds)
func (_Pack *PackFilterer) FilterPackOpenFulfilled(opts *bind.FilterOpts, packId []*big.Int, opener []common.Address, rewardContract []common.Address) (*PackPackOpenFulfilledIterator, error) {

	var packIdRule []interface{}
	for _, packIdItem := range packId {
		packIdRule = append(packIdRule, packIdItem)
	}
	var openerRule []interface{}
	for _, openerItem := range opener {
		openerRule = append(openerRule, openerItem)
	}

	var rewardContractRule []interface{}
	for _, rewardContractItem := range rewardContract {
		rewardContractRule = append(rewardContractRule, rewardContractItem)
	}

	logs, sub, err := _Pack.contract.FilterLogs(opts, "PackOpenFulfilled", packIdRule, openerRule, rewardContractRule)
	if err != nil {
		return nil, err
	}
	return &PackPackOpenFulfilledIterator{contract: _Pack.contract, event: "PackOpenFulfilled", logs: logs, sub: sub}, nil
}

// WatchPackOpenFulfilled is a free log subscription operation binding the contract event 0x9cae8e743935573869e2fe7223782448d71211235c93076166106782de1c234f.
//
// Solidity: event PackOpenFulfilled(uint256 indexed packId, address indexed opener, bytes32 requestId, address indexed rewardContract, uint256[] rewardIds)
func (_Pack *PackFilterer) WatchPackOpenFulfilled(opts *bind.WatchOpts, sink chan<- *PackPackOpenFulfilled, packId []*big.Int, opener []common.Address, rewardContract []common.Address) (event.Subscription, error) {

	var packIdRule []interface{}
	for _, packIdItem := range packId {
		packIdRule = append(packIdRule, packIdItem)
	}
	var openerRule []interface{}
	for _, openerItem := range opener {
		openerRule = append(openerRule, openerItem)
	}

	var rewardContractRule []interface{}
	for _, rewardContractItem := range rewardContract {
		rewardContractRule = append(rewardContractRule, rewardContractItem)
	}

	logs, sub, err := _Pack.contract.WatchLogs(opts, "PackOpenFulfilled", packIdRule, openerRule, rewardContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackPackOpenFulfilled)
				if err := _Pack.contract.UnpackLog(event, "PackOpenFulfilled", log); err != nil {
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

// ParsePackOpenFulfilled is a log parse operation binding the contract event 0x9cae8e743935573869e2fe7223782448d71211235c93076166106782de1c234f.
//
// Solidity: event PackOpenFulfilled(uint256 indexed packId, address indexed opener, bytes32 requestId, address indexed rewardContract, uint256[] rewardIds)
func (_Pack *PackFilterer) ParsePackOpenFulfilled(log types.Log) (*PackPackOpenFulfilled, error) {
	event := new(PackPackOpenFulfilled)
	if err := _Pack.contract.UnpackLog(event, "PackOpenFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackPackOpenRequestIterator is returned from FilterPackOpenRequest and is used to iterate over the raw logs and unpacked data for PackOpenRequest events raised by the Pack contract.
type PackPackOpenRequestIterator struct {
	Event *PackPackOpenRequest // Event containing the contract specifics and raw log

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
func (it *PackPackOpenRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackPackOpenRequest)
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
		it.Event = new(PackPackOpenRequest)
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
func (it *PackPackOpenRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackPackOpenRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackPackOpenRequest represents a PackOpenRequest event raised by the Pack contract.
type PackPackOpenRequest struct {
	PackId    *big.Int
	Opener    common.Address
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPackOpenRequest is a free log retrieval operation binding the contract event 0xb5c27afec62496e403fae78e07136ae7c58425e8128c7b680ea24dee8c8d9401.
//
// Solidity: event PackOpenRequest(uint256 indexed packId, address indexed opener, bytes32 requestId)
func (_Pack *PackFilterer) FilterPackOpenRequest(opts *bind.FilterOpts, packId []*big.Int, opener []common.Address) (*PackPackOpenRequestIterator, error) {

	var packIdRule []interface{}
	for _, packIdItem := range packId {
		packIdRule = append(packIdRule, packIdItem)
	}
	var openerRule []interface{}
	for _, openerItem := range opener {
		openerRule = append(openerRule, openerItem)
	}

	logs, sub, err := _Pack.contract.FilterLogs(opts, "PackOpenRequest", packIdRule, openerRule)
	if err != nil {
		return nil, err
	}
	return &PackPackOpenRequestIterator{contract: _Pack.contract, event: "PackOpenRequest", logs: logs, sub: sub}, nil
}

// WatchPackOpenRequest is a free log subscription operation binding the contract event 0xb5c27afec62496e403fae78e07136ae7c58425e8128c7b680ea24dee8c8d9401.
//
// Solidity: event PackOpenRequest(uint256 indexed packId, address indexed opener, bytes32 requestId)
func (_Pack *PackFilterer) WatchPackOpenRequest(opts *bind.WatchOpts, sink chan<- *PackPackOpenRequest, packId []*big.Int, opener []common.Address) (event.Subscription, error) {

	var packIdRule []interface{}
	for _, packIdItem := range packId {
		packIdRule = append(packIdRule, packIdItem)
	}
	var openerRule []interface{}
	for _, openerItem := range opener {
		openerRule = append(openerRule, openerItem)
	}

	logs, sub, err := _Pack.contract.WatchLogs(opts, "PackOpenRequest", packIdRule, openerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackPackOpenRequest)
				if err := _Pack.contract.UnpackLog(event, "PackOpenRequest", log); err != nil {
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

// ParsePackOpenRequest is a log parse operation binding the contract event 0xb5c27afec62496e403fae78e07136ae7c58425e8128c7b680ea24dee8c8d9401.
//
// Solidity: event PackOpenRequest(uint256 indexed packId, address indexed opener, bytes32 requestId)
func (_Pack *PackFilterer) ParsePackOpenRequest(log types.Log) (*PackPackOpenRequest, error) {
	event := new(PackPackOpenRequest)
	if err := _Pack.contract.UnpackLog(event, "PackOpenRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pack contract.
type PackPausedIterator struct {
	Event *PackPaused // Event containing the contract specifics and raw log

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
func (it *PackPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackPaused)
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
		it.Event = new(PackPaused)
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
func (it *PackPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackPaused represents a Paused event raised by the Pack contract.
type PackPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pack *PackFilterer) FilterPaused(opts *bind.FilterOpts) (*PackPausedIterator, error) {

	logs, sub, err := _Pack.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PackPausedIterator{contract: _Pack.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pack *PackFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PackPaused) (event.Subscription, error) {

	logs, sub, err := _Pack.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackPaused)
				if err := _Pack.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Pack *PackFilterer) ParsePaused(log types.Log) (*PackPaused, error) {
	event := new(PackPaused)
	if err := _Pack.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRestrictedTransferUpdatedIterator is returned from FilterRestrictedTransferUpdated and is used to iterate over the raw logs and unpacked data for RestrictedTransferUpdated events raised by the Pack contract.
type PackRestrictedTransferUpdatedIterator struct {
	Event *PackRestrictedTransferUpdated // Event containing the contract specifics and raw log

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
func (it *PackRestrictedTransferUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRestrictedTransferUpdated)
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
		it.Event = new(PackRestrictedTransferUpdated)
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
func (it *PackRestrictedTransferUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRestrictedTransferUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRestrictedTransferUpdated represents a RestrictedTransferUpdated event raised by the Pack contract.
type PackRestrictedTransferUpdated struct {
	Transferable bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRestrictedTransferUpdated is a free log retrieval operation binding the contract event 0xfb4ba02cee22486df888d7ffb97c6100ec3193781e025cb9a4bc6fc358d626cc.
//
// Solidity: event RestrictedTransferUpdated(bool transferable)
func (_Pack *PackFilterer) FilterRestrictedTransferUpdated(opts *bind.FilterOpts) (*PackRestrictedTransferUpdatedIterator, error) {

	logs, sub, err := _Pack.contract.FilterLogs(opts, "RestrictedTransferUpdated")
	if err != nil {
		return nil, err
	}
	return &PackRestrictedTransferUpdatedIterator{contract: _Pack.contract, event: "RestrictedTransferUpdated", logs: logs, sub: sub}, nil
}

// WatchRestrictedTransferUpdated is a free log subscription operation binding the contract event 0xfb4ba02cee22486df888d7ffb97c6100ec3193781e025cb9a4bc6fc358d626cc.
//
// Solidity: event RestrictedTransferUpdated(bool transferable)
func (_Pack *PackFilterer) WatchRestrictedTransferUpdated(opts *bind.WatchOpts, sink chan<- *PackRestrictedTransferUpdated) (event.Subscription, error) {

	logs, sub, err := _Pack.contract.WatchLogs(opts, "RestrictedTransferUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRestrictedTransferUpdated)
				if err := _Pack.contract.UnpackLog(event, "RestrictedTransferUpdated", log); err != nil {
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

// ParseRestrictedTransferUpdated is a log parse operation binding the contract event 0xfb4ba02cee22486df888d7ffb97c6100ec3193781e025cb9a4bc6fc358d626cc.
//
// Solidity: event RestrictedTransferUpdated(bool transferable)
func (_Pack *PackFilterer) ParseRestrictedTransferUpdated(log types.Log) (*PackRestrictedTransferUpdated, error) {
	event := new(PackRestrictedTransferUpdated)
	if err := _Pack.contract.UnpackLog(event, "RestrictedTransferUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Pack contract.
type PackRoleAdminChangedIterator struct {
	Event *PackRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PackRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRoleAdminChanged)
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
		it.Event = new(PackRoleAdminChanged)
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
func (it *PackRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRoleAdminChanged represents a RoleAdminChanged event raised by the Pack contract.
type PackRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pack *PackFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PackRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Pack.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PackRoleAdminChangedIterator{contract: _Pack.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Pack *PackFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PackRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Pack.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRoleAdminChanged)
				if err := _Pack.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Pack *PackFilterer) ParseRoleAdminChanged(log types.Log) (*PackRoleAdminChanged, error) {
	event := new(PackRoleAdminChanged)
	if err := _Pack.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Pack contract.
type PackRoleGrantedIterator struct {
	Event *PackRoleGranted // Event containing the contract specifics and raw log

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
func (it *PackRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRoleGranted)
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
		it.Event = new(PackRoleGranted)
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
func (it *PackRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRoleGranted represents a RoleGranted event raised by the Pack contract.
type PackRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pack *PackFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PackRoleGrantedIterator, error) {

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

	logs, sub, err := _Pack.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PackRoleGrantedIterator{contract: _Pack.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pack *PackFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PackRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pack.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRoleGranted)
				if err := _Pack.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Pack *PackFilterer) ParseRoleGranted(log types.Log) (*PackRoleGranted, error) {
	event := new(PackRoleGranted)
	if err := _Pack.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Pack contract.
type PackRoleRevokedIterator struct {
	Event *PackRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PackRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRoleRevoked)
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
		it.Event = new(PackRoleRevoked)
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
func (it *PackRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRoleRevoked represents a RoleRevoked event raised by the Pack contract.
type PackRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pack *PackFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PackRoleRevokedIterator, error) {

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

	logs, sub, err := _Pack.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PackRoleRevokedIterator{contract: _Pack.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Pack *PackFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PackRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pack.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRoleRevoked)
				if err := _Pack.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Pack *PackFilterer) ParseRoleRevoked(log types.Log) (*PackRoleRevoked, error) {
	event := new(PackRoleRevoked)
	if err := _Pack.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackRoyaltyUpdatedIterator is returned from FilterRoyaltyUpdated and is used to iterate over the raw logs and unpacked data for RoyaltyUpdated events raised by the Pack contract.
type PackRoyaltyUpdatedIterator struct {
	Event *PackRoyaltyUpdated // Event containing the contract specifics and raw log

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
func (it *PackRoyaltyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackRoyaltyUpdated)
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
		it.Event = new(PackRoyaltyUpdated)
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
func (it *PackRoyaltyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackRoyaltyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackRoyaltyUpdated represents a RoyaltyUpdated event raised by the Pack contract.
type PackRoyaltyUpdated struct {
	RoyaltyBps *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyUpdated is a free log retrieval operation binding the contract event 0x244ea8d7627f5a08f4299862bd5a45752842c183aee5b0fb0d1e4887bfa605b3.
//
// Solidity: event RoyaltyUpdated(uint256 royaltyBps)
func (_Pack *PackFilterer) FilterRoyaltyUpdated(opts *bind.FilterOpts) (*PackRoyaltyUpdatedIterator, error) {

	logs, sub, err := _Pack.contract.FilterLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return &PackRoyaltyUpdatedIterator{contract: _Pack.contract, event: "RoyaltyUpdated", logs: logs, sub: sub}, nil
}

// WatchRoyaltyUpdated is a free log subscription operation binding the contract event 0x244ea8d7627f5a08f4299862bd5a45752842c183aee5b0fb0d1e4887bfa605b3.
//
// Solidity: event RoyaltyUpdated(uint256 royaltyBps)
func (_Pack *PackFilterer) WatchRoyaltyUpdated(opts *bind.WatchOpts, sink chan<- *PackRoyaltyUpdated) (event.Subscription, error) {

	logs, sub, err := _Pack.contract.WatchLogs(opts, "RoyaltyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackRoyaltyUpdated)
				if err := _Pack.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
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
func (_Pack *PackFilterer) ParseRoyaltyUpdated(log types.Log) (*PackRoyaltyUpdated, error) {
	event := new(PackRoyaltyUpdated)
	if err := _Pack.contract.UnpackLog(event, "RoyaltyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the Pack contract.
type PackTransferBatchIterator struct {
	Event *PackTransferBatch // Event containing the contract specifics and raw log

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
func (it *PackTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackTransferBatch)
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
		it.Event = new(PackTransferBatch)
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
func (it *PackTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackTransferBatch represents a TransferBatch event raised by the Pack contract.
type PackTransferBatch struct {
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
func (_Pack *PackFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PackTransferBatchIterator, error) {

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

	logs, sub, err := _Pack.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PackTransferBatchIterator{contract: _Pack.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_Pack *PackFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *PackTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pack.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackTransferBatch)
				if err := _Pack.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_Pack *PackFilterer) ParseTransferBatch(log types.Log) (*PackTransferBatch, error) {
	event := new(PackTransferBatch)
	if err := _Pack.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the Pack contract.
type PackTransferSingleIterator struct {
	Event *PackTransferSingle // Event containing the contract specifics and raw log

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
func (it *PackTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackTransferSingle)
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
		it.Event = new(PackTransferSingle)
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
func (it *PackTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackTransferSingle represents a TransferSingle event raised by the Pack contract.
type PackTransferSingle struct {
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
func (_Pack *PackFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PackTransferSingleIterator, error) {

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

	logs, sub, err := _Pack.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PackTransferSingleIterator{contract: _Pack.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_Pack *PackFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *PackTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Pack.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackTransferSingle)
				if err := _Pack.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_Pack *PackFilterer) ParseTransferSingle(log types.Log) (*PackTransferSingle, error) {
	event := new(PackTransferSingle)
	if err := _Pack.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the Pack contract.
type PackURIIterator struct {
	Event *PackURI // Event containing the contract specifics and raw log

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
func (it *PackURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackURI)
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
		it.Event = new(PackURI)
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
func (it *PackURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackURI represents a URI event raised by the Pack contract.
type PackURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Pack *PackFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*PackURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pack.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &PackURIIterator{contract: _Pack.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_Pack *PackFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *PackURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pack.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackURI)
				if err := _Pack.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_Pack *PackFilterer) ParseURI(log types.Log) (*PackURI, error) {
	event := new(PackURI)
	if err := _Pack.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PackUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pack contract.
type PackUnpausedIterator struct {
	Event *PackUnpaused // Event containing the contract specifics and raw log

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
func (it *PackUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PackUnpaused)
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
		it.Event = new(PackUnpaused)
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
func (it *PackUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PackUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PackUnpaused represents a Unpaused event raised by the Pack contract.
type PackUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pack *PackFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PackUnpausedIterator, error) {

	logs, sub, err := _Pack.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PackUnpausedIterator{contract: _Pack.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pack *PackFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PackUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pack.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PackUnpaused)
				if err := _Pack.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Pack *PackFilterer) ParseUnpaused(log types.Log) (*PackUnpaused, error) {
	event := new(PackUnpaused)
	if err := _Pack.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
