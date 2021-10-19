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

// MarketListing is an auto generated low-level Go binding around an user-defined struct.
type MarketListing struct {
	ListingId      *big.Int
	Seller         common.Address
	AssetContract  common.Address
	TokenId        *big.Int
	Quantity       *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
	SaleStart      *big.Int
	SaleEnd        *big.Int
	TokensPerBuyer *big.Int
	TokenType      uint8
}

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_controlCenter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structMarket.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"ListingUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"newFee\",\"type\":\"uint128\"}],\"name\":\"MarketFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structMarket.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"NewListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structMarket.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"NewSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restricted\",\"type\":\"bool\"}],\"name\":\"RestrictedListerRoleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LISTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"addToListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"boughtFromListing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Listing[]\",\"name\":\"allListings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"getListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getListingsByAsset\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Listing[]\",\"name\":\"tokenListings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetContract\",\"type\":\"address\"}],\"name\":\"getListingsByAssetContract\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Listing[]\",\"name\":\"tokenListings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getListingsBySeller\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Listing[]\",\"name\":\"sellerListings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilEnd\",\"type\":\"uint256\"}],\"name\":\"list\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"enumMarket.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketFeeBps\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrictedListerRoleOnly\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_URI\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"feeBps\",\"type\":\"uint128\"}],\"name\":\"setMarketFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"restricted\",\"type\":\"bool\"}],\"name\":\"setRestrictedListerRoleOnly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"}],\"name\":\"unlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokensPerBuyer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilEnd\",\"type\":\"uint256\"}],\"name\":\"updateListingParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Market *MarketCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Market *MarketSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Market.Contract.DEFAULTADMINROLE(&_Market.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Market *MarketCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Market.Contract.DEFAULTADMINROLE(&_Market.CallOpts)
}

// LISTERROLE is a free data retrieval call binding the contract method 0xdeb26b94.
//
// Solidity: function LISTER_ROLE() view returns(bytes32)
func (_Market *MarketCaller) LISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "LISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LISTERROLE is a free data retrieval call binding the contract method 0xdeb26b94.
//
// Solidity: function LISTER_ROLE() view returns(bytes32)
func (_Market *MarketSession) LISTERROLE() ([32]byte, error) {
	return _Market.Contract.LISTERROLE(&_Market.CallOpts)
}

// LISTERROLE is a free data retrieval call binding the contract method 0xdeb26b94.
//
// Solidity: function LISTER_ROLE() view returns(bytes32)
func (_Market *MarketCallerSession) LISTERROLE() ([32]byte, error) {
	return _Market.Contract.LISTERROLE(&_Market.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Market *MarketCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Market *MarketSession) PAUSERROLE() ([32]byte, error) {
	return _Market.Contract.PAUSERROLE(&_Market.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Market *MarketCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Market.Contract.PAUSERROLE(&_Market.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Market *MarketCaller) ContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "_contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Market *MarketSession) ContractURI() (string, error) {
	return _Market.Contract.ContractURI(&_Market.CallOpts)
}

// ContractURI is a free data retrieval call binding the contract method 0xc0e72740.
//
// Solidity: function _contractURI() view returns(string)
func (_Market *MarketCallerSession) ContractURI() (string, error) {
	return _Market.Contract.ContractURI(&_Market.CallOpts)
}

// BoughtFromListing is a free data retrieval call binding the contract method 0x20c7852c.
//
// Solidity: function boughtFromListing(uint256 , address ) view returns(uint256)
func (_Market *MarketCaller) BoughtFromListing(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "boughtFromListing", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BoughtFromListing is a free data retrieval call binding the contract method 0x20c7852c.
//
// Solidity: function boughtFromListing(uint256 , address ) view returns(uint256)
func (_Market *MarketSession) BoughtFromListing(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.BoughtFromListing(&_Market.CallOpts, arg0, arg1)
}

// BoughtFromListing is a free data retrieval call binding the contract method 0x20c7852c.
//
// Solidity: function boughtFromListing(uint256 , address ) view returns(uint256)
func (_Market *MarketCallerSession) BoughtFromListing(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.BoughtFromListing(&_Market.CallOpts, arg0, arg1)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Market *MarketCaller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Market *MarketSession) InternalContractURI() (string, error) {
	return _Market.Contract.InternalContractURI(&_Market.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Market *MarketCallerSession) InternalContractURI() (string, error) {
	return _Market.Contract.InternalContractURI(&_Market.CallOpts)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] allListings)
func (_Market *MarketCaller) GetAllListings(opts *bind.CallOpts) ([]MarketListing, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getAllListings")

	if err != nil {
		return *new([]MarketListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketListing)).(*[]MarketListing)

	return out0, err

}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] allListings)
func (_Market *MarketSession) GetAllListings() ([]MarketListing, error) {
	return _Market.Contract.GetAllListings(&_Market.CallOpts)
}

// GetAllListings is a free data retrieval call binding the contract method 0xae73ccec.
//
// Solidity: function getAllListings() view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] allListings)
func (_Market *MarketCallerSession) GetAllListings() ([]MarketListing, error) {
	return _Market.Contract.GetAllListings(&_Market.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketCaller) GetListing(opts *bind.CallOpts, _listingId *big.Int) (MarketListing, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getListing", _listingId)

	if err != nil {
		return *new(MarketListing), err
	}

	out0 := *abi.ConvertType(out[0], new(MarketListing)).(*MarketListing)

	return out0, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketSession) GetListing(_listingId *big.Int) (MarketListing, error) {
	return _Market.Contract.GetListing(&_Market.CallOpts, _listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 _listingId) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketCallerSession) GetListing(_listingId *big.Int) (MarketListing, error) {
	return _Market.Contract.GetListing(&_Market.CallOpts, _listingId)
}

// GetListingsByAsset is a free data retrieval call binding the contract method 0x1d0cd75e.
//
// Solidity: function getListingsByAsset(address _assetContract, uint256 _tokenId) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] tokenListings)
func (_Market *MarketCaller) GetListingsByAsset(opts *bind.CallOpts, _assetContract common.Address, _tokenId *big.Int) ([]MarketListing, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getListingsByAsset", _assetContract, _tokenId)

	if err != nil {
		return *new([]MarketListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketListing)).(*[]MarketListing)

	return out0, err

}

// GetListingsByAsset is a free data retrieval call binding the contract method 0x1d0cd75e.
//
// Solidity: function getListingsByAsset(address _assetContract, uint256 _tokenId) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] tokenListings)
func (_Market *MarketSession) GetListingsByAsset(_assetContract common.Address, _tokenId *big.Int) ([]MarketListing, error) {
	return _Market.Contract.GetListingsByAsset(&_Market.CallOpts, _assetContract, _tokenId)
}

// GetListingsByAsset is a free data retrieval call binding the contract method 0x1d0cd75e.
//
// Solidity: function getListingsByAsset(address _assetContract, uint256 _tokenId) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] tokenListings)
func (_Market *MarketCallerSession) GetListingsByAsset(_assetContract common.Address, _tokenId *big.Int) ([]MarketListing, error) {
	return _Market.Contract.GetListingsByAsset(&_Market.CallOpts, _assetContract, _tokenId)
}

// GetListingsByAssetContract is a free data retrieval call binding the contract method 0x4edeaa81.
//
// Solidity: function getListingsByAssetContract(address _assetContract) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] tokenListings)
func (_Market *MarketCaller) GetListingsByAssetContract(opts *bind.CallOpts, _assetContract common.Address) ([]MarketListing, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getListingsByAssetContract", _assetContract)

	if err != nil {
		return *new([]MarketListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketListing)).(*[]MarketListing)

	return out0, err

}

// GetListingsByAssetContract is a free data retrieval call binding the contract method 0x4edeaa81.
//
// Solidity: function getListingsByAssetContract(address _assetContract) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] tokenListings)
func (_Market *MarketSession) GetListingsByAssetContract(_assetContract common.Address) ([]MarketListing, error) {
	return _Market.Contract.GetListingsByAssetContract(&_Market.CallOpts, _assetContract)
}

// GetListingsByAssetContract is a free data retrieval call binding the contract method 0x4edeaa81.
//
// Solidity: function getListingsByAssetContract(address _assetContract) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] tokenListings)
func (_Market *MarketCallerSession) GetListingsByAssetContract(_assetContract common.Address) ([]MarketListing, error) {
	return _Market.Contract.GetListingsByAssetContract(&_Market.CallOpts, _assetContract)
}

// GetListingsBySeller is a free data retrieval call binding the contract method 0xd8cba251.
//
// Solidity: function getListingsBySeller(address _seller) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] sellerListings)
func (_Market *MarketCaller) GetListingsBySeller(opts *bind.CallOpts, _seller common.Address) ([]MarketListing, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getListingsBySeller", _seller)

	if err != nil {
		return *new([]MarketListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketListing)).(*[]MarketListing)

	return out0, err

}

// GetListingsBySeller is a free data retrieval call binding the contract method 0xd8cba251.
//
// Solidity: function getListingsBySeller(address _seller) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] sellerListings)
func (_Market *MarketSession) GetListingsBySeller(_seller common.Address) ([]MarketListing, error) {
	return _Market.Contract.GetListingsBySeller(&_Market.CallOpts, _seller)
}

// GetListingsBySeller is a free data retrieval call binding the contract method 0xd8cba251.
//
// Solidity: function getListingsBySeller(address _seller) view returns((uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8)[] sellerListings)
func (_Market *MarketCallerSession) GetListingsBySeller(_seller common.Address) ([]MarketListing, error) {
	return _Market.Contract.GetListingsBySeller(&_Market.CallOpts, _seller)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Market *MarketCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Market *MarketSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Market.Contract.GetRoleAdmin(&_Market.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Market *MarketCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Market.Contract.GetRoleAdmin(&_Market.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Market *MarketCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Market *MarketSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Market.Contract.GetRoleMember(&_Market.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Market *MarketCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Market.Contract.GetRoleMember(&_Market.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Market *MarketCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Market *MarketSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Market.Contract.GetRoleMemberCount(&_Market.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Market *MarketCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Market.Contract.GetRoleMemberCount(&_Market.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Market *MarketCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Market *MarketSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Market.Contract.HasRole(&_Market.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Market *MarketCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Market.Contract.HasRole(&_Market.CallOpts, role, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Market *MarketCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Market *MarketSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Market.Contract.IsTrustedForwarder(&_Market.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Market *MarketCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Market.Contract.IsTrustedForwarder(&_Market.CallOpts, forwarder)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address seller, address assetContract, uint256 tokenId, uint256 quantity, address currency, uint256 pricePerToken, uint256 saleStart, uint256 saleEnd, uint256 tokensPerBuyer, uint8 tokenType)
func (_Market *MarketCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId      *big.Int
	Seller         common.Address
	AssetContract  common.Address
	TokenId        *big.Int
	Quantity       *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
	SaleStart      *big.Int
	SaleEnd        *big.Int
	TokensPerBuyer *big.Int
	TokenType      uint8
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		ListingId      *big.Int
		Seller         common.Address
		AssetContract  common.Address
		TokenId        *big.Int
		Quantity       *big.Int
		Currency       common.Address
		PricePerToken  *big.Int
		SaleStart      *big.Int
		SaleEnd        *big.Int
		TokensPerBuyer *big.Int
		TokenType      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AssetContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Quantity = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.PricePerToken = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.SaleStart = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.SaleEnd = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.TokensPerBuyer = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokenType = *abi.ConvertType(out[10], new(uint8)).(*uint8)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address seller, address assetContract, uint256 tokenId, uint256 quantity, address currency, uint256 pricePerToken, uint256 saleStart, uint256 saleEnd, uint256 tokensPerBuyer, uint8 tokenType)
func (_Market *MarketSession) Listings(arg0 *big.Int) (struct {
	ListingId      *big.Int
	Seller         common.Address
	AssetContract  common.Address
	TokenId        *big.Int
	Quantity       *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
	SaleStart      *big.Int
	SaleEnd        *big.Int
	TokensPerBuyer *big.Int
	TokenType      uint8
}, error) {
	return _Market.Contract.Listings(&_Market.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address seller, address assetContract, uint256 tokenId, uint256 quantity, address currency, uint256 pricePerToken, uint256 saleStart, uint256 saleEnd, uint256 tokensPerBuyer, uint8 tokenType)
func (_Market *MarketCallerSession) Listings(arg0 *big.Int) (struct {
	ListingId      *big.Int
	Seller         common.Address
	AssetContract  common.Address
	TokenId        *big.Int
	Quantity       *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
	SaleStart      *big.Int
	SaleEnd        *big.Int
	TokensPerBuyer *big.Int
	TokenType      uint8
}, error) {
	return _Market.Contract.Listings(&_Market.CallOpts, arg0)
}

// MarketFeeBps is a free data retrieval call binding the contract method 0x3f5c3e87.
//
// Solidity: function marketFeeBps() view returns(uint128)
func (_Market *MarketCaller) MarketFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "marketFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MarketFeeBps is a free data retrieval call binding the contract method 0x3f5c3e87.
//
// Solidity: function marketFeeBps() view returns(uint128)
func (_Market *MarketSession) MarketFeeBps() (*big.Int, error) {
	return _Market.Contract.MarketFeeBps(&_Market.CallOpts)
}

// MarketFeeBps is a free data retrieval call binding the contract method 0x3f5c3e87.
//
// Solidity: function marketFeeBps() view returns(uint128)
func (_Market *MarketCallerSession) MarketFeeBps() (*big.Int, error) {
	return _Market.Contract.MarketFeeBps(&_Market.CallOpts)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Market *MarketCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Market *MarketSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Market.Contract.OnERC721Received(&_Market.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Market *MarketCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Market.Contract.OnERC721Received(&_Market.CallOpts, arg0, arg1, arg2, arg3)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketSession) Paused() (bool, error) {
	return _Market.Contract.Paused(&_Market.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Market *MarketCallerSession) Paused() (bool, error) {
	return _Market.Contract.Paused(&_Market.CallOpts)
}

// RestrictedListerRoleOnly is a free data retrieval call binding the contract method 0x61096ec6.
//
// Solidity: function restrictedListerRoleOnly() view returns(bool)
func (_Market *MarketCaller) RestrictedListerRoleOnly(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "restrictedListerRoleOnly")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RestrictedListerRoleOnly is a free data retrieval call binding the contract method 0x61096ec6.
//
// Solidity: function restrictedListerRoleOnly() view returns(bool)
func (_Market *MarketSession) RestrictedListerRoleOnly() (bool, error) {
	return _Market.Contract.RestrictedListerRoleOnly(&_Market.CallOpts)
}

// RestrictedListerRoleOnly is a free data retrieval call binding the contract method 0x61096ec6.
//
// Solidity: function restrictedListerRoleOnly() view returns(bool)
func (_Market *MarketCallerSession) RestrictedListerRoleOnly() (bool, error) {
	return _Market.Contract.RestrictedListerRoleOnly(&_Market.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Market *MarketCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Market *MarketSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Market.Contract.SupportsInterface(&_Market.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Market *MarketCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Market.Contract.SupportsInterface(&_Market.CallOpts, interfaceId)
}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_Market *MarketCaller) TotalListings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "totalListings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_Market *MarketSession) TotalListings() (*big.Int, error) {
	return _Market.Contract.TotalListings(&_Market.CallOpts)
}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_Market *MarketCallerSession) TotalListings() (*big.Int, error) {
	return _Market.Contract.TotalListings(&_Market.CallOpts)
}

// AddToListing is a paid mutator transaction binding the contract method 0x583e12b6.
//
// Solidity: function addToListing(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketTransactor) AddToListing(opts *bind.TransactOpts, _listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "addToListing", _listingId, _quantity)
}

// AddToListing is a paid mutator transaction binding the contract method 0x583e12b6.
//
// Solidity: function addToListing(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketSession) AddToListing(_listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AddToListing(&_Market.TransactOpts, _listingId, _quantity)
}

// AddToListing is a paid mutator transaction binding the contract method 0x583e12b6.
//
// Solidity: function addToListing(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketTransactorSession) AddToListing(_listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.Contract.AddToListing(&_Market.TransactOpts, _listingId, _quantity)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketTransactor) Buy(opts *bind.TransactOpts, _listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "buy", _listingId, _quantity)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketSession) Buy(_listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Buy(&_Market.TransactOpts, _listingId, _quantity)
}

// Buy is a paid mutator transaction binding the contract method 0xd6febde8.
//
// Solidity: function buy(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketTransactorSession) Buy(_listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Buy(&_Market.TransactOpts, _listingId, _quantity)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Market *MarketTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Market *MarketSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.Contract.GrantRole(&_Market.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Market *MarketTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.Contract.GrantRole(&_Market.TransactOpts, role, account)
}

// List is a paid mutator transaction binding the contract method 0x9b782ff2.
//
// Solidity: function list(address _assetContract, uint256 _tokenId, address _currency, uint256 _pricePerToken, uint256 _quantity, uint256 _tokensPerBuyer, uint256 _secondsUntilStart, uint256 _secondsUntilEnd) returns()
func (_Market *MarketTransactor) List(opts *bind.TransactOpts, _assetContract common.Address, _tokenId *big.Int, _currency common.Address, _pricePerToken *big.Int, _quantity *big.Int, _tokensPerBuyer *big.Int, _secondsUntilStart *big.Int, _secondsUntilEnd *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "list", _assetContract, _tokenId, _currency, _pricePerToken, _quantity, _tokensPerBuyer, _secondsUntilStart, _secondsUntilEnd)
}

// List is a paid mutator transaction binding the contract method 0x9b782ff2.
//
// Solidity: function list(address _assetContract, uint256 _tokenId, address _currency, uint256 _pricePerToken, uint256 _quantity, uint256 _tokensPerBuyer, uint256 _secondsUntilStart, uint256 _secondsUntilEnd) returns()
func (_Market *MarketSession) List(_assetContract common.Address, _tokenId *big.Int, _currency common.Address, _pricePerToken *big.Int, _quantity *big.Int, _tokensPerBuyer *big.Int, _secondsUntilStart *big.Int, _secondsUntilEnd *big.Int) (*types.Transaction, error) {
	return _Market.Contract.List(&_Market.TransactOpts, _assetContract, _tokenId, _currency, _pricePerToken, _quantity, _tokensPerBuyer, _secondsUntilStart, _secondsUntilEnd)
}

// List is a paid mutator transaction binding the contract method 0x9b782ff2.
//
// Solidity: function list(address _assetContract, uint256 _tokenId, address _currency, uint256 _pricePerToken, uint256 _quantity, uint256 _tokensPerBuyer, uint256 _secondsUntilStart, uint256 _secondsUntilEnd) returns()
func (_Market *MarketTransactorSession) List(_assetContract common.Address, _tokenId *big.Int, _currency common.Address, _pricePerToken *big.Int, _quantity *big.Int, _tokensPerBuyer *big.Int, _secondsUntilStart *big.Int, _secondsUntilEnd *big.Int) (*types.Transaction, error) {
	return _Market.Contract.List(&_Market.TransactOpts, _assetContract, _tokenId, _currency, _pricePerToken, _quantity, _tokensPerBuyer, _secondsUntilStart, _secondsUntilEnd)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Market *MarketTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Market *MarketSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Market.Contract.Multicall(&_Market.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Market *MarketTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Market.Contract.Multicall(&_Market.TransactOpts, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Market *MarketTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Market *MarketSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Market.Contract.OnERC1155BatchReceived(&_Market.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Market *MarketTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Market.Contract.OnERC1155BatchReceived(&_Market.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Market *MarketTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Market *MarketSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Market.Contract.OnERC1155Received(&_Market.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Market *MarketTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Market.Contract.OnERC1155Received(&_Market.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Market *MarketTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Market *MarketSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.Contract.RenounceRole(&_Market.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Market *MarketTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.Contract.RenounceRole(&_Market.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Market *MarketTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Market *MarketSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.Contract.RevokeRole(&_Market.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Market *MarketTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Market.Contract.RevokeRole(&_Market.TransactOpts, role, account)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_Market *MarketTransactor) SetContractURI(opts *bind.TransactOpts, _URI string) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setContractURI", _URI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_Market *MarketSession) SetContractURI(_URI string) (*types.Transaction, error) {
	return _Market.Contract.SetContractURI(&_Market.TransactOpts, _URI)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _URI) returns()
func (_Market *MarketTransactorSession) SetContractURI(_URI string) (*types.Transaction, error) {
	return _Market.Contract.SetContractURI(&_Market.TransactOpts, _URI)
}

// SetMarketFeeBps is a paid mutator transaction binding the contract method 0x09679b39.
//
// Solidity: function setMarketFeeBps(uint128 feeBps) returns()
func (_Market *MarketTransactor) SetMarketFeeBps(opts *bind.TransactOpts, feeBps *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setMarketFeeBps", feeBps)
}

// SetMarketFeeBps is a paid mutator transaction binding the contract method 0x09679b39.
//
// Solidity: function setMarketFeeBps(uint128 feeBps) returns()
func (_Market *MarketSession) SetMarketFeeBps(feeBps *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetMarketFeeBps(&_Market.TransactOpts, feeBps)
}

// SetMarketFeeBps is a paid mutator transaction binding the contract method 0x09679b39.
//
// Solidity: function setMarketFeeBps(uint128 feeBps) returns()
func (_Market *MarketTransactorSession) SetMarketFeeBps(feeBps *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetMarketFeeBps(&_Market.TransactOpts, feeBps)
}

// SetRestrictedListerRoleOnly is a paid mutator transaction binding the contract method 0x354c7ab6.
//
// Solidity: function setRestrictedListerRoleOnly(bool restricted) returns()
func (_Market *MarketTransactor) SetRestrictedListerRoleOnly(opts *bind.TransactOpts, restricted bool) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setRestrictedListerRoleOnly", restricted)
}

// SetRestrictedListerRoleOnly is a paid mutator transaction binding the contract method 0x354c7ab6.
//
// Solidity: function setRestrictedListerRoleOnly(bool restricted) returns()
func (_Market *MarketSession) SetRestrictedListerRoleOnly(restricted bool) (*types.Transaction, error) {
	return _Market.Contract.SetRestrictedListerRoleOnly(&_Market.TransactOpts, restricted)
}

// SetRestrictedListerRoleOnly is a paid mutator transaction binding the contract method 0x354c7ab6.
//
// Solidity: function setRestrictedListerRoleOnly(bool restricted) returns()
func (_Market *MarketTransactorSession) SetRestrictedListerRoleOnly(restricted bool) (*types.Transaction, error) {
	return _Market.Contract.SetRestrictedListerRoleOnly(&_Market.TransactOpts, restricted)
}

// Unlist is a paid mutator transaction binding the contract method 0x4fd69f3c.
//
// Solidity: function unlist(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketTransactor) Unlist(opts *bind.TransactOpts, _listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "unlist", _listingId, _quantity)
}

// Unlist is a paid mutator transaction binding the contract method 0x4fd69f3c.
//
// Solidity: function unlist(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketSession) Unlist(_listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Unlist(&_Market.TransactOpts, _listingId, _quantity)
}

// Unlist is a paid mutator transaction binding the contract method 0x4fd69f3c.
//
// Solidity: function unlist(uint256 _listingId, uint256 _quantity) returns()
func (_Market *MarketTransactorSession) Unlist(_listingId *big.Int, _quantity *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Unlist(&_Market.TransactOpts, _listingId, _quantity)
}

// UpdateListingParams is a paid mutator transaction binding the contract method 0x31851012.
//
// Solidity: function updateListingParams(uint256 _listingId, uint256 _pricePerToken, address _currency, uint256 _tokensPerBuyer, uint256 _secondsUntilStart, uint256 _secondsUntilEnd) returns()
func (_Market *MarketTransactor) UpdateListingParams(opts *bind.TransactOpts, _listingId *big.Int, _pricePerToken *big.Int, _currency common.Address, _tokensPerBuyer *big.Int, _secondsUntilStart *big.Int, _secondsUntilEnd *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "updateListingParams", _listingId, _pricePerToken, _currency, _tokensPerBuyer, _secondsUntilStart, _secondsUntilEnd)
}

// UpdateListingParams is a paid mutator transaction binding the contract method 0x31851012.
//
// Solidity: function updateListingParams(uint256 _listingId, uint256 _pricePerToken, address _currency, uint256 _tokensPerBuyer, uint256 _secondsUntilStart, uint256 _secondsUntilEnd) returns()
func (_Market *MarketSession) UpdateListingParams(_listingId *big.Int, _pricePerToken *big.Int, _currency common.Address, _tokensPerBuyer *big.Int, _secondsUntilStart *big.Int, _secondsUntilEnd *big.Int) (*types.Transaction, error) {
	return _Market.Contract.UpdateListingParams(&_Market.TransactOpts, _listingId, _pricePerToken, _currency, _tokensPerBuyer, _secondsUntilStart, _secondsUntilEnd)
}

// UpdateListingParams is a paid mutator transaction binding the contract method 0x31851012.
//
// Solidity: function updateListingParams(uint256 _listingId, uint256 _pricePerToken, address _currency, uint256 _tokensPerBuyer, uint256 _secondsUntilStart, uint256 _secondsUntilEnd) returns()
func (_Market *MarketTransactorSession) UpdateListingParams(_listingId *big.Int, _pricePerToken *big.Int, _currency common.Address, _tokensPerBuyer *big.Int, _secondsUntilStart *big.Int, _secondsUntilEnd *big.Int) (*types.Transaction, error) {
	return _Market.Contract.UpdateListingParams(&_Market.TransactOpts, _listingId, _pricePerToken, _currency, _tokensPerBuyer, _secondsUntilStart, _secondsUntilEnd)
}

// MarketListingUpdateIterator is returned from FilterListingUpdate and is used to iterate over the raw logs and unpacked data for ListingUpdate events raised by the Market contract.
type MarketListingUpdateIterator struct {
	Event *MarketListingUpdate // Event containing the contract specifics and raw log

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
func (it *MarketListingUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketListingUpdate)
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
		it.Event = new(MarketListingUpdate)
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
func (it *MarketListingUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketListingUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketListingUpdate represents a ListingUpdate event raised by the Market contract.
type MarketListingUpdate struct {
	Seller    common.Address
	ListingId *big.Int
	Listing   MarketListing
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingUpdate is a free log retrieval operation binding the contract event 0xb4c8fd604ad229ddf17ffe34f0851a2abf5113b048f8235b12edbabcc6dd9193.
//
// Solidity: event ListingUpdate(address indexed seller, uint256 indexed listingId, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) FilterListingUpdate(opts *bind.FilterOpts, seller []common.Address, listingId []*big.Int) (*MarketListingUpdateIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "ListingUpdate", sellerRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketListingUpdateIterator{contract: _Market.contract, event: "ListingUpdate", logs: logs, sub: sub}, nil
}

// WatchListingUpdate is a free log subscription operation binding the contract event 0xb4c8fd604ad229ddf17ffe34f0851a2abf5113b048f8235b12edbabcc6dd9193.
//
// Solidity: event ListingUpdate(address indexed seller, uint256 indexed listingId, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) WatchListingUpdate(opts *bind.WatchOpts, sink chan<- *MarketListingUpdate, seller []common.Address, listingId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "ListingUpdate", sellerRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketListingUpdate)
				if err := _Market.contract.UnpackLog(event, "ListingUpdate", log); err != nil {
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

// ParseListingUpdate is a log parse operation binding the contract event 0xb4c8fd604ad229ddf17ffe34f0851a2abf5113b048f8235b12edbabcc6dd9193.
//
// Solidity: event ListingUpdate(address indexed seller, uint256 indexed listingId, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) ParseListingUpdate(log types.Log) (*MarketListingUpdate, error) {
	event := new(MarketListingUpdate)
	if err := _Market.contract.UnpackLog(event, "ListingUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketMarketFeeUpdateIterator is returned from FilterMarketFeeUpdate and is used to iterate over the raw logs and unpacked data for MarketFeeUpdate events raised by the Market contract.
type MarketMarketFeeUpdateIterator struct {
	Event *MarketMarketFeeUpdate // Event containing the contract specifics and raw log

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
func (it *MarketMarketFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketMarketFeeUpdate)
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
		it.Event = new(MarketMarketFeeUpdate)
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
func (it *MarketMarketFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketMarketFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketMarketFeeUpdate represents a MarketFeeUpdate event raised by the Market contract.
type MarketMarketFeeUpdate struct {
	NewFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketFeeUpdate is a free log retrieval operation binding the contract event 0xd50e64e6eb05cd7ceafe1a28b1a7ad949edb90451106259c7117252d605178ef.
//
// Solidity: event MarketFeeUpdate(uint128 newFee)
func (_Market *MarketFilterer) FilterMarketFeeUpdate(opts *bind.FilterOpts) (*MarketMarketFeeUpdateIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "MarketFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &MarketMarketFeeUpdateIterator{contract: _Market.contract, event: "MarketFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchMarketFeeUpdate is a free log subscription operation binding the contract event 0xd50e64e6eb05cd7ceafe1a28b1a7ad949edb90451106259c7117252d605178ef.
//
// Solidity: event MarketFeeUpdate(uint128 newFee)
func (_Market *MarketFilterer) WatchMarketFeeUpdate(opts *bind.WatchOpts, sink chan<- *MarketMarketFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "MarketFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketMarketFeeUpdate)
				if err := _Market.contract.UnpackLog(event, "MarketFeeUpdate", log); err != nil {
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

// ParseMarketFeeUpdate is a log parse operation binding the contract event 0xd50e64e6eb05cd7ceafe1a28b1a7ad949edb90451106259c7117252d605178ef.
//
// Solidity: event MarketFeeUpdate(uint128 newFee)
func (_Market *MarketFilterer) ParseMarketFeeUpdate(log types.Log) (*MarketMarketFeeUpdate, error) {
	event := new(MarketMarketFeeUpdate)
	if err := _Market.contract.UnpackLog(event, "MarketFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketNewListingIterator is returned from FilterNewListing and is used to iterate over the raw logs and unpacked data for NewListing events raised by the Market contract.
type MarketNewListingIterator struct {
	Event *MarketNewListing // Event containing the contract specifics and raw log

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
func (it *MarketNewListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketNewListing)
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
		it.Event = new(MarketNewListing)
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
func (it *MarketNewListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketNewListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketNewListing represents a NewListing event raised by the Market contract.
type MarketNewListing struct {
	AssetContract common.Address
	Seller        common.Address
	ListingId     *big.Int
	Listing       MarketListing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewListing is a free log retrieval operation binding the contract event 0x70c5741a020504dbda58d308c3efe5275326c456d3b00c541925222c40f7c62e.
//
// Solidity: event NewListing(address indexed assetContract, address indexed seller, uint256 indexed listingId, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) FilterNewListing(opts *bind.FilterOpts, assetContract []common.Address, seller []common.Address, listingId []*big.Int) (*MarketNewListingIterator, error) {

	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "NewListing", assetContractRule, sellerRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketNewListingIterator{contract: _Market.contract, event: "NewListing", logs: logs, sub: sub}, nil
}

// WatchNewListing is a free log subscription operation binding the contract event 0x70c5741a020504dbda58d308c3efe5275326c456d3b00c541925222c40f7c62e.
//
// Solidity: event NewListing(address indexed assetContract, address indexed seller, uint256 indexed listingId, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) WatchNewListing(opts *bind.WatchOpts, sink chan<- *MarketNewListing, assetContract []common.Address, seller []common.Address, listingId []*big.Int) (event.Subscription, error) {

	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "NewListing", assetContractRule, sellerRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketNewListing)
				if err := _Market.contract.UnpackLog(event, "NewListing", log); err != nil {
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

// ParseNewListing is a log parse operation binding the contract event 0x70c5741a020504dbda58d308c3efe5275326c456d3b00c541925222c40f7c62e.
//
// Solidity: event NewListing(address indexed assetContract, address indexed seller, uint256 indexed listingId, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) ParseNewListing(log types.Log) (*MarketNewListing, error) {
	event := new(MarketNewListing)
	if err := _Market.contract.UnpackLog(event, "NewListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketNewSaleIterator is returned from FilterNewSale and is used to iterate over the raw logs and unpacked data for NewSale events raised by the Market contract.
type MarketNewSaleIterator struct {
	Event *MarketNewSale // Event containing the contract specifics and raw log

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
func (it *MarketNewSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketNewSale)
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
		it.Event = new(MarketNewSale)
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
func (it *MarketNewSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketNewSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketNewSale represents a NewSale event raised by the Market contract.
type MarketNewSale struct {
	AssetContract common.Address
	Seller        common.Address
	ListingId     *big.Int
	Buyer         common.Address
	Quantity      *big.Int
	Listing       MarketListing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewSale is a free log retrieval operation binding the contract event 0xc848190182320d1cb2ea6d8a80041c6780e56643bc41fe9060bb1f1349902cba.
//
// Solidity: event NewSale(address indexed assetContract, address indexed seller, uint256 indexed listingId, address buyer, uint256 quantity, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) FilterNewSale(opts *bind.FilterOpts, assetContract []common.Address, seller []common.Address, listingId []*big.Int) (*MarketNewSaleIterator, error) {

	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "NewSale", assetContractRule, sellerRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketNewSaleIterator{contract: _Market.contract, event: "NewSale", logs: logs, sub: sub}, nil
}

// WatchNewSale is a free log subscription operation binding the contract event 0xc848190182320d1cb2ea6d8a80041c6780e56643bc41fe9060bb1f1349902cba.
//
// Solidity: event NewSale(address indexed assetContract, address indexed seller, uint256 indexed listingId, address buyer, uint256 quantity, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) WatchNewSale(opts *bind.WatchOpts, sink chan<- *MarketNewSale, assetContract []common.Address, seller []common.Address, listingId []*big.Int) (event.Subscription, error) {

	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "NewSale", assetContractRule, sellerRule, listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketNewSale)
				if err := _Market.contract.UnpackLog(event, "NewSale", log); err != nil {
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

// ParseNewSale is a log parse operation binding the contract event 0xc848190182320d1cb2ea6d8a80041c6780e56643bc41fe9060bb1f1349902cba.
//
// Solidity: event NewSale(address indexed assetContract, address indexed seller, uint256 indexed listingId, address buyer, uint256 quantity, (uint256,address,address,uint256,uint256,address,uint256,uint256,uint256,uint256,uint8) listing)
func (_Market *MarketFilterer) ParseNewSale(log types.Log) (*MarketNewSale, error) {
	event := new(MarketNewSale)
	if err := _Market.contract.UnpackLog(event, "NewSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Market contract.
type MarketPausedIterator struct {
	Event *MarketPaused // Event containing the contract specifics and raw log

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
func (it *MarketPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketPaused)
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
		it.Event = new(MarketPaused)
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
func (it *MarketPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketPaused represents a Paused event raised by the Market contract.
type MarketPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Market *MarketFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketPausedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketPausedIterator{contract: _Market.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Market *MarketFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketPaused) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketPaused)
				if err := _Market.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Market *MarketFilterer) ParsePaused(log types.Log) (*MarketPaused, error) {
	event := new(MarketPaused)
	if err := _Market.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketRestrictedListerRoleUpdatedIterator is returned from FilterRestrictedListerRoleUpdated and is used to iterate over the raw logs and unpacked data for RestrictedListerRoleUpdated events raised by the Market contract.
type MarketRestrictedListerRoleUpdatedIterator struct {
	Event *MarketRestrictedListerRoleUpdated // Event containing the contract specifics and raw log

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
func (it *MarketRestrictedListerRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketRestrictedListerRoleUpdated)
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
		it.Event = new(MarketRestrictedListerRoleUpdated)
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
func (it *MarketRestrictedListerRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketRestrictedListerRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketRestrictedListerRoleUpdated represents a RestrictedListerRoleUpdated event raised by the Market contract.
type MarketRestrictedListerRoleUpdated struct {
	Restricted bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRestrictedListerRoleUpdated is a free log retrieval operation binding the contract event 0x56668abce4c06e28f52176150c07856e31ee073b9b62df62a5afbf1a22aec24c.
//
// Solidity: event RestrictedListerRoleUpdated(bool restricted)
func (_Market *MarketFilterer) FilterRestrictedListerRoleUpdated(opts *bind.FilterOpts) (*MarketRestrictedListerRoleUpdatedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "RestrictedListerRoleUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketRestrictedListerRoleUpdatedIterator{contract: _Market.contract, event: "RestrictedListerRoleUpdated", logs: logs, sub: sub}, nil
}

// WatchRestrictedListerRoleUpdated is a free log subscription operation binding the contract event 0x56668abce4c06e28f52176150c07856e31ee073b9b62df62a5afbf1a22aec24c.
//
// Solidity: event RestrictedListerRoleUpdated(bool restricted)
func (_Market *MarketFilterer) WatchRestrictedListerRoleUpdated(opts *bind.WatchOpts, sink chan<- *MarketRestrictedListerRoleUpdated) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "RestrictedListerRoleUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketRestrictedListerRoleUpdated)
				if err := _Market.contract.UnpackLog(event, "RestrictedListerRoleUpdated", log); err != nil {
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

// ParseRestrictedListerRoleUpdated is a log parse operation binding the contract event 0x56668abce4c06e28f52176150c07856e31ee073b9b62df62a5afbf1a22aec24c.
//
// Solidity: event RestrictedListerRoleUpdated(bool restricted)
func (_Market *MarketFilterer) ParseRestrictedListerRoleUpdated(log types.Log) (*MarketRestrictedListerRoleUpdated, error) {
	event := new(MarketRestrictedListerRoleUpdated)
	if err := _Market.contract.UnpackLog(event, "RestrictedListerRoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Market contract.
type MarketRoleAdminChangedIterator struct {
	Event *MarketRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarketRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketRoleAdminChanged)
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
		it.Event = new(MarketRoleAdminChanged)
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
func (it *MarketRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketRoleAdminChanged represents a RoleAdminChanged event raised by the Market contract.
type MarketRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Market *MarketFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MarketRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Market.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MarketRoleAdminChangedIterator{contract: _Market.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Market *MarketFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MarketRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Market.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketRoleAdminChanged)
				if err := _Market.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Market *MarketFilterer) ParseRoleAdminChanged(log types.Log) (*MarketRoleAdminChanged, error) {
	event := new(MarketRoleAdminChanged)
	if err := _Market.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Market contract.
type MarketRoleGrantedIterator struct {
	Event *MarketRoleGranted // Event containing the contract specifics and raw log

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
func (it *MarketRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketRoleGranted)
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
		it.Event = new(MarketRoleGranted)
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
func (it *MarketRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketRoleGranted represents a RoleGranted event raised by the Market contract.
type MarketRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Market *MarketFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketRoleGrantedIterator, error) {

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

	logs, sub, err := _Market.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketRoleGrantedIterator{contract: _Market.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Market *MarketFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MarketRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Market.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketRoleGranted)
				if err := _Market.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Market *MarketFilterer) ParseRoleGranted(log types.Log) (*MarketRoleGranted, error) {
	event := new(MarketRoleGranted)
	if err := _Market.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Market contract.
type MarketRoleRevokedIterator struct {
	Event *MarketRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MarketRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketRoleRevoked)
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
		it.Event = new(MarketRoleRevoked)
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
func (it *MarketRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketRoleRevoked represents a RoleRevoked event raised by the Market contract.
type MarketRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Market *MarketFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketRoleRevokedIterator, error) {

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

	logs, sub, err := _Market.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketRoleRevokedIterator{contract: _Market.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Market *MarketFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MarketRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Market.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketRoleRevoked)
				if err := _Market.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Market *MarketFilterer) ParseRoleRevoked(log types.Log) (*MarketRoleRevoked, error) {
	event := new(MarketRoleRevoked)
	if err := _Market.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Market contract.
type MarketUnpausedIterator struct {
	Event *MarketUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketUnpaused)
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
		it.Event = new(MarketUnpaused)
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
func (it *MarketUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketUnpaused represents a Unpaused event raised by the Market contract.
type MarketUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Market *MarketFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketUnpausedIterator, error) {

	logs, sub, err := _Market.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketUnpausedIterator{contract: _Market.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Market *MarketFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketUnpaused) (event.Subscription, error) {

	logs, sub, err := _Market.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketUnpaused)
				if err := _Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Market *MarketFilterer) ParseUnpaused(log types.Log) (*MarketUnpaused, error) {
	event := new(MarketUnpaused)
	if err := _Market.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
