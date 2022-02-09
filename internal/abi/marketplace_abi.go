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

// IMarketplaceListing is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceListing struct {
	ListingId            *big.Int
	TokenOwner           common.Address
	AssetContract        common.Address
	TokenId              *big.Int
	StartTime            *big.Int
	EndTime              *big.Int
	Quantity             *big.Int
	Currency             common.Address
	ReservePricePerToken *big.Int
	BuyoutPricePerToken  *big.Int
	TokenType            uint8
	ListingType          uint8
}

// IMarketplaceListingParameters is an auto generated low-level Go binding around an user-defined struct.
type IMarketplaceListingParameters struct {
	AssetContract        common.Address
	TokenId              *big.Int
	StartTime            *big.Int
	SecondsUntilEndTime  *big.Int
	QuantityToList       *big.Int
	CurrencyToAccept     common.Address
	ReservePricePerToken *big.Int
	BuyoutPricePerToken  *big.Int
	ListingType          uint8
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_controlCenter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_trustedForwarder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nativeTokenWrapper\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_marketFeeBps\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timeBuffer\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidBufferBps\",\"type\":\"uint256\"}],\"name\":\"AuctionBuffersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"cancelled\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionCreator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winningBidder\",\"type\":\"address\"}],\"name\":\"AuctionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"restricted\",\"type\":\"bool\"}],\"name\":\"ListingRestricted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"listingCreator\",\"type\":\"address\"}],\"name\":\"ListingUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newFee\",\"type\":\"uint64\"}],\"name\":\"MarketFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lister\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reservePricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyoutPricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIMarketplace.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"NewListing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offeror\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityWanted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalOfferAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"NewOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"lister\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPricePaid\",\"type\":\"uint256\"}],\"name\":\"NewSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LISTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"offeror\",\"type\":\"address\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bidBufferBps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantityToBuy\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalPrice\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_closeFor\",\"type\":\"address\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondsUntilEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityToList\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currencyToAccept\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reservePricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyoutPricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"}],\"internalType\":\"structIMarketplace.ListingParameters\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"createListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"assetContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reservePricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyoutPricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"enumIMarketplace.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"enumIMarketplace.ListingType\",\"name\":\"listingType\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketFeeBps\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeTokenWrapper\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantityWanted\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"}],\"name\":\"offer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"offers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"offeror\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantityWanted\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restrictedListerRoleOnly\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timeBuffer\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidBufferBps\",\"type\":\"uint256\"}],\"name\":\"setAuctionBuffers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeBps\",\"type\":\"uint256\"}],\"name\":\"setMarketFeeBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"restricted\",\"type\":\"bool\"}],\"name\":\"setRestrictedListerRoleOnly\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeBuffer\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantityToList\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reservePricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_buyoutPricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currencyToAccept\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondsUntilEndTime\",\"type\":\"uint256\"}],\"name\":\"updateListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"winningBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"offeror\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantityWanted\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marketplace.Contract.DEFAULTADMINROLE(&_Marketplace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Marketplace.Contract.DEFAULTADMINROLE(&_Marketplace.CallOpts)
}

// LISTERROLE is a free data retrieval call binding the contract method 0xdeb26b94.
//
// Solidity: function LISTER_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) LISTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "LISTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LISTERROLE is a free data retrieval call binding the contract method 0xdeb26b94.
//
// Solidity: function LISTER_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceSession) LISTERROLE() ([32]byte, error) {
	return _Marketplace.Contract.LISTERROLE(&_Marketplace.CallOpts)
}

// LISTERROLE is a free data retrieval call binding the contract method 0xdeb26b94.
//
// Solidity: function LISTER_ROLE() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) LISTERROLE() ([32]byte, error) {
	return _Marketplace.Contract.LISTERROLE(&_Marketplace.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint64)
func (_Marketplace *MarketplaceCaller) MAXBPS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint64)
func (_Marketplace *MarketplaceSession) MAXBPS() (uint64, error) {
	return _Marketplace.Contract.MAXBPS(&_Marketplace.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint64)
func (_Marketplace *MarketplaceCallerSession) MAXBPS() (uint64, error) {
	return _Marketplace.Contract.MAXBPS(&_Marketplace.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_Marketplace *MarketplaceCaller) NATIVETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "NATIVE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_Marketplace *MarketplaceSession) NATIVETOKEN() (common.Address, error) {
	return _Marketplace.Contract.NATIVETOKEN(&_Marketplace.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_Marketplace *MarketplaceCallerSession) NATIVETOKEN() (common.Address, error) {
	return _Marketplace.Contract.NATIVETOKEN(&_Marketplace.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Marketplace *MarketplaceCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Marketplace *MarketplaceSession) VERSION() (*big.Int, error) {
	return _Marketplace.Contract.VERSION(&_Marketplace.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) VERSION() (*big.Int, error) {
	return _Marketplace.Contract.VERSION(&_Marketplace.CallOpts)
}

// BidBufferBps is a free data retrieval call binding the contract method 0x4e03f28d.
//
// Solidity: function bidBufferBps() view returns(uint64)
func (_Marketplace *MarketplaceCaller) BidBufferBps(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "bidBufferBps")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BidBufferBps is a free data retrieval call binding the contract method 0x4e03f28d.
//
// Solidity: function bidBufferBps() view returns(uint64)
func (_Marketplace *MarketplaceSession) BidBufferBps() (uint64, error) {
	return _Marketplace.Contract.BidBufferBps(&_Marketplace.CallOpts)
}

// BidBufferBps is a free data retrieval call binding the contract method 0x4e03f28d.
//
// Solidity: function bidBufferBps() view returns(uint64)
func (_Marketplace *MarketplaceCallerSession) BidBufferBps() (uint64, error) {
	return _Marketplace.Contract.BidBufferBps(&_Marketplace.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Marketplace *MarketplaceCaller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Marketplace *MarketplaceSession) InternalContractURI() (string, error) {
	return _Marketplace.Contract.InternalContractURI(&_Marketplace.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_Marketplace *MarketplaceCallerSession) InternalContractURI() (string, error) {
	return _Marketplace.Contract.InternalContractURI(&_Marketplace.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplace *MarketplaceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplace *MarketplaceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marketplace.Contract.GetRoleAdmin(&_Marketplace.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Marketplace.Contract.GetRoleAdmin(&_Marketplace.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplace *MarketplaceCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplace *MarketplaceSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Marketplace.Contract.GetRoleMember(&_Marketplace.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Marketplace.Contract.GetRoleMember(&_Marketplace.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplace *MarketplaceSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Marketplace.Contract.GetRoleMemberCount(&_Marketplace.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Marketplace.Contract.GetRoleMemberCount(&_Marketplace.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplace *MarketplaceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplace *MarketplaceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marketplace.Contract.HasRole(&_Marketplace.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Marketplace.Contract.HasRole(&_Marketplace.CallOpts, role, account)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Marketplace *MarketplaceSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Marketplace.Contract.IsTrustedForwarder(&_Marketplace.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _Marketplace.Contract.IsTrustedForwarder(&_Marketplace.CallOpts, forwarder)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address tokenOwner, address assetContract, uint256 tokenId, uint256 startTime, uint256 endTime, uint256 quantity, address currency, uint256 reservePricePerToken, uint256 buyoutPricePerToken, uint8 tokenType, uint8 listingType)
func (_Marketplace *MarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId            *big.Int
	TokenOwner           common.Address
	AssetContract        common.Address
	TokenId              *big.Int
	StartTime            *big.Int
	EndTime              *big.Int
	Quantity             *big.Int
	Currency             common.Address
	ReservePricePerToken *big.Int
	BuyoutPricePerToken  *big.Int
	TokenType            uint8
	ListingType          uint8
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		ListingId            *big.Int
		TokenOwner           common.Address
		AssetContract        common.Address
		TokenId              *big.Int
		StartTime            *big.Int
		EndTime              *big.Int
		Quantity             *big.Int
		Currency             common.Address
		ReservePricePerToken *big.Int
		BuyoutPricePerToken  *big.Int
		TokenType            uint8
		ListingType          uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TokenOwner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.AssetContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Quantity = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.ReservePricePerToken = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.BuyoutPricePerToken = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.TokenType = *abi.ConvertType(out[10], new(uint8)).(*uint8)
	outstruct.ListingType = *abi.ConvertType(out[11], new(uint8)).(*uint8)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address tokenOwner, address assetContract, uint256 tokenId, uint256 startTime, uint256 endTime, uint256 quantity, address currency, uint256 reservePricePerToken, uint256 buyoutPricePerToken, uint8 tokenType, uint8 listingType)
func (_Marketplace *MarketplaceSession) Listings(arg0 *big.Int) (struct {
	ListingId            *big.Int
	TokenOwner           common.Address
	AssetContract        common.Address
	TokenId              *big.Int
	StartTime            *big.Int
	EndTime              *big.Int
	Quantity             *big.Int
	Currency             common.Address
	ReservePricePerToken *big.Int
	BuyoutPricePerToken  *big.Int
	TokenType            uint8
	ListingType          uint8
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 listingId, address tokenOwner, address assetContract, uint256 tokenId, uint256 startTime, uint256 endTime, uint256 quantity, address currency, uint256 reservePricePerToken, uint256 buyoutPricePerToken, uint8 tokenType, uint8 listingType)
func (_Marketplace *MarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	ListingId            *big.Int
	TokenOwner           common.Address
	AssetContract        common.Address
	TokenId              *big.Int
	StartTime            *big.Int
	EndTime              *big.Int
	Quantity             *big.Int
	Currency             common.Address
	ReservePricePerToken *big.Int
	BuyoutPricePerToken  *big.Int
	TokenType            uint8
	ListingType          uint8
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// MarketFeeBps is a free data retrieval call binding the contract method 0x3f5c3e87.
//
// Solidity: function marketFeeBps() view returns(uint64)
func (_Marketplace *MarketplaceCaller) MarketFeeBps(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "marketFeeBps")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MarketFeeBps is a free data retrieval call binding the contract method 0x3f5c3e87.
//
// Solidity: function marketFeeBps() view returns(uint64)
func (_Marketplace *MarketplaceSession) MarketFeeBps() (uint64, error) {
	return _Marketplace.Contract.MarketFeeBps(&_Marketplace.CallOpts)
}

// MarketFeeBps is a free data retrieval call binding the contract method 0x3f5c3e87.
//
// Solidity: function marketFeeBps() view returns(uint64)
func (_Marketplace *MarketplaceCallerSession) MarketFeeBps() (uint64, error) {
	return _Marketplace.Contract.MarketFeeBps(&_Marketplace.CallOpts)
}

// NativeTokenWrapper is a free data retrieval call binding the contract method 0xf9ea29cb.
//
// Solidity: function nativeTokenWrapper() view returns(address)
func (_Marketplace *MarketplaceCaller) NativeTokenWrapper(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nativeTokenWrapper")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeTokenWrapper is a free data retrieval call binding the contract method 0xf9ea29cb.
//
// Solidity: function nativeTokenWrapper() view returns(address)
func (_Marketplace *MarketplaceSession) NativeTokenWrapper() (common.Address, error) {
	return _Marketplace.Contract.NativeTokenWrapper(&_Marketplace.CallOpts)
}

// NativeTokenWrapper is a free data retrieval call binding the contract method 0xf9ea29cb.
//
// Solidity: function nativeTokenWrapper() view returns(address)
func (_Marketplace *MarketplaceCallerSession) NativeTokenWrapper() (common.Address, error) {
	return _Marketplace.Contract.NativeTokenWrapper(&_Marketplace.CallOpts)
}

// Offers is a free data retrieval call binding the contract method 0xebdfbce5.
//
// Solidity: function offers(uint256 , address ) view returns(uint256 listingId, address offeror, uint256 quantityWanted, address currency, uint256 pricePerToken)
func (_Marketplace *MarketplaceCaller) Offers(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	ListingId      *big.Int
	Offeror        common.Address
	QuantityWanted *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "offers", arg0, arg1)

	outstruct := new(struct {
		ListingId      *big.Int
		Offeror        common.Address
		QuantityWanted *big.Int
		Currency       common.Address
		PricePerToken  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Offeror = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.QuantityWanted = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.PricePerToken = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offers is a free data retrieval call binding the contract method 0xebdfbce5.
//
// Solidity: function offers(uint256 , address ) view returns(uint256 listingId, address offeror, uint256 quantityWanted, address currency, uint256 pricePerToken)
func (_Marketplace *MarketplaceSession) Offers(arg0 *big.Int, arg1 common.Address) (struct {
	ListingId      *big.Int
	Offeror        common.Address
	QuantityWanted *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
}, error) {
	return _Marketplace.Contract.Offers(&_Marketplace.CallOpts, arg0, arg1)
}

// Offers is a free data retrieval call binding the contract method 0xebdfbce5.
//
// Solidity: function offers(uint256 , address ) view returns(uint256 listingId, address offeror, uint256 quantityWanted, address currency, uint256 pricePerToken)
func (_Marketplace *MarketplaceCallerSession) Offers(arg0 *big.Int, arg1 common.Address) (struct {
	ListingId      *big.Int
	Offeror        common.Address
	QuantityWanted *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
}, error) {
	return _Marketplace.Contract.Offers(&_Marketplace.CallOpts, arg0, arg1)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marketplace *MarketplaceCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Marketplace *MarketplaceCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Marketplace.Contract.OnERC721Received(&_Marketplace.CallOpts, arg0, arg1, arg2, arg3)
}

// RestrictedListerRoleOnly is a free data retrieval call binding the contract method 0x61096ec6.
//
// Solidity: function restrictedListerRoleOnly() view returns(bool)
func (_Marketplace *MarketplaceCaller) RestrictedListerRoleOnly(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "restrictedListerRoleOnly")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RestrictedListerRoleOnly is a free data retrieval call binding the contract method 0x61096ec6.
//
// Solidity: function restrictedListerRoleOnly() view returns(bool)
func (_Marketplace *MarketplaceSession) RestrictedListerRoleOnly() (bool, error) {
	return _Marketplace.Contract.RestrictedListerRoleOnly(&_Marketplace.CallOpts)
}

// RestrictedListerRoleOnly is a free data retrieval call binding the contract method 0x61096ec6.
//
// Solidity: function restrictedListerRoleOnly() view returns(bool)
func (_Marketplace *MarketplaceCallerSession) RestrictedListerRoleOnly() (bool, error) {
	return _Marketplace.Contract.RestrictedListerRoleOnly(&_Marketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplace *MarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplace *MarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplace.Contract.SupportsInterface(&_Marketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Marketplace.Contract.SupportsInterface(&_Marketplace.CallOpts, interfaceId)
}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint64)
func (_Marketplace *MarketplaceCaller) TimeBuffer(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "timeBuffer")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint64)
func (_Marketplace *MarketplaceSession) TimeBuffer() (uint64, error) {
	return _Marketplace.Contract.TimeBuffer(&_Marketplace.CallOpts)
}

// TimeBuffer is a free data retrieval call binding the contract method 0xec91f2a4.
//
// Solidity: function timeBuffer() view returns(uint64)
func (_Marketplace *MarketplaceCallerSession) TimeBuffer() (uint64, error) {
	return _Marketplace.Contract.TimeBuffer(&_Marketplace.CallOpts)
}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_Marketplace *MarketplaceCaller) TotalListings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "totalListings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_Marketplace *MarketplaceSession) TotalListings() (*big.Int, error) {
	return _Marketplace.Contract.TotalListings(&_Marketplace.CallOpts)
}

// TotalListings is a free data retrieval call binding the contract method 0xc78b616c.
//
// Solidity: function totalListings() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) TotalListings() (*big.Int, error) {
	return _Marketplace.Contract.TotalListings(&_Marketplace.CallOpts)
}

// WinningBid is a free data retrieval call binding the contract method 0xd4ac9b8c.
//
// Solidity: function winningBid(uint256 ) view returns(uint256 listingId, address offeror, uint256 quantityWanted, address currency, uint256 pricePerToken)
func (_Marketplace *MarketplaceCaller) WinningBid(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ListingId      *big.Int
	Offeror        common.Address
	QuantityWanted *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "winningBid", arg0)

	outstruct := new(struct {
		ListingId      *big.Int
		Offeror        common.Address
		QuantityWanted *big.Int
		Currency       common.Address
		PricePerToken  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ListingId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Offeror = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.QuantityWanted = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Currency = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.PricePerToken = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// WinningBid is a free data retrieval call binding the contract method 0xd4ac9b8c.
//
// Solidity: function winningBid(uint256 ) view returns(uint256 listingId, address offeror, uint256 quantityWanted, address currency, uint256 pricePerToken)
func (_Marketplace *MarketplaceSession) WinningBid(arg0 *big.Int) (struct {
	ListingId      *big.Int
	Offeror        common.Address
	QuantityWanted *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
}, error) {
	return _Marketplace.Contract.WinningBid(&_Marketplace.CallOpts, arg0)
}

// WinningBid is a free data retrieval call binding the contract method 0xd4ac9b8c.
//
// Solidity: function winningBid(uint256 ) view returns(uint256 listingId, address offeror, uint256 quantityWanted, address currency, uint256 pricePerToken)
func (_Marketplace *MarketplaceCallerSession) WinningBid(arg0 *big.Int) (struct {
	ListingId      *big.Int
	Offeror        common.Address
	QuantityWanted *big.Int
	Currency       common.Address
	PricePerToken  *big.Int
}, error) {
	return _Marketplace.Contract.WinningBid(&_Marketplace.CallOpts, arg0)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x918d407d.
//
// Solidity: function acceptOffer(uint256 _listingId, address offeror) returns()
func (_Marketplace *MarketplaceTransactor) AcceptOffer(opts *bind.TransactOpts, _listingId *big.Int, offeror common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "acceptOffer", _listingId, offeror)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x918d407d.
//
// Solidity: function acceptOffer(uint256 _listingId, address offeror) returns()
func (_Marketplace *MarketplaceSession) AcceptOffer(_listingId *big.Int, offeror common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AcceptOffer(&_Marketplace.TransactOpts, _listingId, offeror)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x918d407d.
//
// Solidity: function acceptOffer(uint256 _listingId, address offeror) returns()
func (_Marketplace *MarketplaceTransactorSession) AcceptOffer(_listingId *big.Int, offeror common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.AcceptOffer(&_Marketplace.TransactOpts, _listingId, offeror)
}

// Buy is a paid mutator transaction binding the contract method 0x8945257c.
//
// Solidity: function buy(uint256 _listingId, uint256 _quantityToBuy, address _currency, uint256 _totalPrice) payable returns()
func (_Marketplace *MarketplaceTransactor) Buy(opts *bind.TransactOpts, _listingId *big.Int, _quantityToBuy *big.Int, _currency common.Address, _totalPrice *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "buy", _listingId, _quantityToBuy, _currency, _totalPrice)
}

// Buy is a paid mutator transaction binding the contract method 0x8945257c.
//
// Solidity: function buy(uint256 _listingId, uint256 _quantityToBuy, address _currency, uint256 _totalPrice) payable returns()
func (_Marketplace *MarketplaceSession) Buy(_listingId *big.Int, _quantityToBuy *big.Int, _currency common.Address, _totalPrice *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, _listingId, _quantityToBuy, _currency, _totalPrice)
}

// Buy is a paid mutator transaction binding the contract method 0x8945257c.
//
// Solidity: function buy(uint256 _listingId, uint256 _quantityToBuy, address _currency, uint256 _totalPrice) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Buy(_listingId *big.Int, _quantityToBuy *big.Int, _currency common.Address, _totalPrice *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Buy(&_Marketplace.TransactOpts, _listingId, _quantityToBuy, _currency, _totalPrice)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x6bab66ae.
//
// Solidity: function closeAuction(uint256 _listingId, address _closeFor) returns()
func (_Marketplace *MarketplaceTransactor) CloseAuction(opts *bind.TransactOpts, _listingId *big.Int, _closeFor common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "closeAuction", _listingId, _closeFor)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x6bab66ae.
//
// Solidity: function closeAuction(uint256 _listingId, address _closeFor) returns()
func (_Marketplace *MarketplaceSession) CloseAuction(_listingId *big.Int, _closeFor common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.CloseAuction(&_Marketplace.TransactOpts, _listingId, _closeFor)
}

// CloseAuction is a paid mutator transaction binding the contract method 0x6bab66ae.
//
// Solidity: function closeAuction(uint256 _listingId, address _closeFor) returns()
func (_Marketplace *MarketplaceTransactorSession) CloseAuction(_listingId *big.Int, _closeFor common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.CloseAuction(&_Marketplace.TransactOpts, _listingId, _closeFor)
}

// CreateListing is a paid mutator transaction binding the contract method 0x296f4e16.
//
// Solidity: function createListing((address,uint256,uint256,uint256,uint256,address,uint256,uint256,uint8) _params) returns()
func (_Marketplace *MarketplaceTransactor) CreateListing(opts *bind.TransactOpts, _params IMarketplaceListingParameters) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createListing", _params)
}

// CreateListing is a paid mutator transaction binding the contract method 0x296f4e16.
//
// Solidity: function createListing((address,uint256,uint256,uint256,uint256,address,uint256,uint256,uint8) _params) returns()
func (_Marketplace *MarketplaceSession) CreateListing(_params IMarketplaceListingParameters) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, _params)
}

// CreateListing is a paid mutator transaction binding the contract method 0x296f4e16.
//
// Solidity: function createListing((address,uint256,uint256,uint256,uint256,address,uint256,uint256,uint8) _params) returns()
func (_Marketplace *MarketplaceTransactorSession) CreateListing(_params IMarketplaceListingParameters) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, _params)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.GrantRole(&_Marketplace.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.GrantRole(&_Marketplace.TransactOpts, role, account)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Marketplace *MarketplaceTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Marketplace *MarketplaceSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.Multicall(&_Marketplace.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_Marketplace *MarketplaceTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.Multicall(&_Marketplace.TransactOpts, data)
}

// Offer is a paid mutator transaction binding the contract method 0xacb1ba67.
//
// Solidity: function offer(uint256 _listingId, uint256 _quantityWanted, address _currency, uint256 _pricePerToken) payable returns()
func (_Marketplace *MarketplaceTransactor) Offer(opts *bind.TransactOpts, _listingId *big.Int, _quantityWanted *big.Int, _currency common.Address, _pricePerToken *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "offer", _listingId, _quantityWanted, _currency, _pricePerToken)
}

// Offer is a paid mutator transaction binding the contract method 0xacb1ba67.
//
// Solidity: function offer(uint256 _listingId, uint256 _quantityWanted, address _currency, uint256 _pricePerToken) payable returns()
func (_Marketplace *MarketplaceSession) Offer(_listingId *big.Int, _quantityWanted *big.Int, _currency common.Address, _pricePerToken *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Offer(&_Marketplace.TransactOpts, _listingId, _quantityWanted, _currency, _pricePerToken)
}

// Offer is a paid mutator transaction binding the contract method 0xacb1ba67.
//
// Solidity: function offer(uint256 _listingId, uint256 _quantityWanted, address _currency, uint256 _pricePerToken) payable returns()
func (_Marketplace *MarketplaceTransactorSession) Offer(_listingId *big.Int, _quantityWanted *big.Int, _currency common.Address, _pricePerToken *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Offer(&_Marketplace.TransactOpts, _listingId, _quantityWanted, _currency, _pricePerToken)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC1155BatchReceived(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC1155BatchReceived(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC1155Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Marketplace *MarketplaceTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.OnERC1155Received(&_Marketplace.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceRole(&_Marketplace.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceRole(&_Marketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RevokeRole(&_Marketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Marketplace *MarketplaceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.RevokeRole(&_Marketplace.TransactOpts, role, account)
}

// SetAuctionBuffers is a paid mutator transaction binding the contract method 0xea0e0241.
//
// Solidity: function setAuctionBuffers(uint256 _timeBuffer, uint256 _bidBufferBps) returns()
func (_Marketplace *MarketplaceTransactor) SetAuctionBuffers(opts *bind.TransactOpts, _timeBuffer *big.Int, _bidBufferBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setAuctionBuffers", _timeBuffer, _bidBufferBps)
}

// SetAuctionBuffers is a paid mutator transaction binding the contract method 0xea0e0241.
//
// Solidity: function setAuctionBuffers(uint256 _timeBuffer, uint256 _bidBufferBps) returns()
func (_Marketplace *MarketplaceSession) SetAuctionBuffers(_timeBuffer *big.Int, _bidBufferBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetAuctionBuffers(&_Marketplace.TransactOpts, _timeBuffer, _bidBufferBps)
}

// SetAuctionBuffers is a paid mutator transaction binding the contract method 0xea0e0241.
//
// Solidity: function setAuctionBuffers(uint256 _timeBuffer, uint256 _bidBufferBps) returns()
func (_Marketplace *MarketplaceTransactorSession) SetAuctionBuffers(_timeBuffer *big.Int, _bidBufferBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetAuctionBuffers(&_Marketplace.TransactOpts, _timeBuffer, _bidBufferBps)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Marketplace *MarketplaceTransactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Marketplace *MarketplaceSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetContractURI(&_Marketplace.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_Marketplace *MarketplaceTransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _Marketplace.Contract.SetContractURI(&_Marketplace.TransactOpts, _uri)
}

// SetMarketFeeBps is a paid mutator transaction binding the contract method 0xe4104eaf.
//
// Solidity: function setMarketFeeBps(uint256 _feeBps) returns()
func (_Marketplace *MarketplaceTransactor) SetMarketFeeBps(opts *bind.TransactOpts, _feeBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setMarketFeeBps", _feeBps)
}

// SetMarketFeeBps is a paid mutator transaction binding the contract method 0xe4104eaf.
//
// Solidity: function setMarketFeeBps(uint256 _feeBps) returns()
func (_Marketplace *MarketplaceSession) SetMarketFeeBps(_feeBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMarketFeeBps(&_Marketplace.TransactOpts, _feeBps)
}

// SetMarketFeeBps is a paid mutator transaction binding the contract method 0xe4104eaf.
//
// Solidity: function setMarketFeeBps(uint256 _feeBps) returns()
func (_Marketplace *MarketplaceTransactorSession) SetMarketFeeBps(_feeBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMarketFeeBps(&_Marketplace.TransactOpts, _feeBps)
}

// SetRestrictedListerRoleOnly is a paid mutator transaction binding the contract method 0x354c7ab6.
//
// Solidity: function setRestrictedListerRoleOnly(bool restricted) returns()
func (_Marketplace *MarketplaceTransactor) SetRestrictedListerRoleOnly(opts *bind.TransactOpts, restricted bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRestrictedListerRoleOnly", restricted)
}

// SetRestrictedListerRoleOnly is a paid mutator transaction binding the contract method 0x354c7ab6.
//
// Solidity: function setRestrictedListerRoleOnly(bool restricted) returns()
func (_Marketplace *MarketplaceSession) SetRestrictedListerRoleOnly(restricted bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRestrictedListerRoleOnly(&_Marketplace.TransactOpts, restricted)
}

// SetRestrictedListerRoleOnly is a paid mutator transaction binding the contract method 0x354c7ab6.
//
// Solidity: function setRestrictedListerRoleOnly(bool restricted) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRestrictedListerRoleOnly(restricted bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRestrictedListerRoleOnly(&_Marketplace.TransactOpts, restricted)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xc4b5b15f.
//
// Solidity: function updateListing(uint256 _listingId, uint256 _quantityToList, uint256 _reservePricePerToken, uint256 _buyoutPricePerToken, address _currencyToAccept, uint256 _startTime, uint256 _secondsUntilEndTime) returns()
func (_Marketplace *MarketplaceTransactor) UpdateListing(opts *bind.TransactOpts, _listingId *big.Int, _quantityToList *big.Int, _reservePricePerToken *big.Int, _buyoutPricePerToken *big.Int, _currencyToAccept common.Address, _startTime *big.Int, _secondsUntilEndTime *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateListing", _listingId, _quantityToList, _reservePricePerToken, _buyoutPricePerToken, _currencyToAccept, _startTime, _secondsUntilEndTime)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xc4b5b15f.
//
// Solidity: function updateListing(uint256 _listingId, uint256 _quantityToList, uint256 _reservePricePerToken, uint256 _buyoutPricePerToken, address _currencyToAccept, uint256 _startTime, uint256 _secondsUntilEndTime) returns()
func (_Marketplace *MarketplaceSession) UpdateListing(_listingId *big.Int, _quantityToList *big.Int, _reservePricePerToken *big.Int, _buyoutPricePerToken *big.Int, _currencyToAccept common.Address, _startTime *big.Int, _secondsUntilEndTime *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateListing(&_Marketplace.TransactOpts, _listingId, _quantityToList, _reservePricePerToken, _buyoutPricePerToken, _currencyToAccept, _startTime, _secondsUntilEndTime)
}

// UpdateListing is a paid mutator transaction binding the contract method 0xc4b5b15f.
//
// Solidity: function updateListing(uint256 _listingId, uint256 _quantityToList, uint256 _reservePricePerToken, uint256 _buyoutPricePerToken, address _currencyToAccept, uint256 _startTime, uint256 _secondsUntilEndTime) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateListing(_listingId *big.Int, _quantityToList *big.Int, _reservePricePerToken *big.Int, _buyoutPricePerToken *big.Int, _currencyToAccept common.Address, _startTime *big.Int, _secondsUntilEndTime *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateListing(&_Marketplace.TransactOpts, _listingId, _quantityToList, _reservePricePerToken, _buyoutPricePerToken, _currencyToAccept, _startTime, _secondsUntilEndTime)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// MarketplaceAuctionBuffersUpdatedIterator is returned from FilterAuctionBuffersUpdated and is used to iterate over the raw logs and unpacked data for AuctionBuffersUpdated events raised by the Marketplace contract.
type MarketplaceAuctionBuffersUpdatedIterator struct {
	Event *MarketplaceAuctionBuffersUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceAuctionBuffersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAuctionBuffersUpdated)
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
		it.Event = new(MarketplaceAuctionBuffersUpdated)
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
func (it *MarketplaceAuctionBuffersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAuctionBuffersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAuctionBuffersUpdated represents a AuctionBuffersUpdated event raised by the Marketplace contract.
type MarketplaceAuctionBuffersUpdated struct {
	TimeBuffer   *big.Int
	BidBufferBps *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAuctionBuffersUpdated is a free log retrieval operation binding the contract event 0x441ed6470e96704c3f8c9e70c209107078aab3f17311385e886081b91aa75088.
//
// Solidity: event AuctionBuffersUpdated(uint256 timeBuffer, uint256 bidBufferBps)
func (_Marketplace *MarketplaceFilterer) FilterAuctionBuffersUpdated(opts *bind.FilterOpts) (*MarketplaceAuctionBuffersUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AuctionBuffersUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceAuctionBuffersUpdatedIterator{contract: _Marketplace.contract, event: "AuctionBuffersUpdated", logs: logs, sub: sub}, nil
}

// WatchAuctionBuffersUpdated is a free log subscription operation binding the contract event 0x441ed6470e96704c3f8c9e70c209107078aab3f17311385e886081b91aa75088.
//
// Solidity: event AuctionBuffersUpdated(uint256 timeBuffer, uint256 bidBufferBps)
func (_Marketplace *MarketplaceFilterer) WatchAuctionBuffersUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceAuctionBuffersUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AuctionBuffersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAuctionBuffersUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "AuctionBuffersUpdated", log); err != nil {
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

// ParseAuctionBuffersUpdated is a log parse operation binding the contract event 0x441ed6470e96704c3f8c9e70c209107078aab3f17311385e886081b91aa75088.
//
// Solidity: event AuctionBuffersUpdated(uint256 timeBuffer, uint256 bidBufferBps)
func (_Marketplace *MarketplaceFilterer) ParseAuctionBuffersUpdated(log types.Log) (*MarketplaceAuctionBuffersUpdated, error) {
	event := new(MarketplaceAuctionBuffersUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "AuctionBuffersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceAuctionClosedIterator is returned from FilterAuctionClosed and is used to iterate over the raw logs and unpacked data for AuctionClosed events raised by the Marketplace contract.
type MarketplaceAuctionClosedIterator struct {
	Event *MarketplaceAuctionClosed // Event containing the contract specifics and raw log

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
func (it *MarketplaceAuctionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceAuctionClosed)
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
		it.Event = new(MarketplaceAuctionClosed)
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
func (it *MarketplaceAuctionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceAuctionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceAuctionClosed represents a AuctionClosed event raised by the Marketplace contract.
type MarketplaceAuctionClosed struct {
	ListingId      *big.Int
	Closer         common.Address
	Cancelled      bool
	AuctionCreator common.Address
	WinningBidder  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAuctionClosed is a free log retrieval operation binding the contract event 0x572cdc5ca5e918473319d0f4737494e4709ac879a7d0bcd11ce1bef24b24e81d.
//
// Solidity: event AuctionClosed(uint256 indexed listingId, address indexed closer, bool indexed cancelled, address auctionCreator, address winningBidder)
func (_Marketplace *MarketplaceFilterer) FilterAuctionClosed(opts *bind.FilterOpts, listingId []*big.Int, closer []common.Address, cancelled []bool) (*MarketplaceAuctionClosedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var closerRule []interface{}
	for _, closerItem := range closer {
		closerRule = append(closerRule, closerItem)
	}
	var cancelledRule []interface{}
	for _, cancelledItem := range cancelled {
		cancelledRule = append(cancelledRule, cancelledItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "AuctionClosed", listingIdRule, closerRule, cancelledRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceAuctionClosedIterator{contract: _Marketplace.contract, event: "AuctionClosed", logs: logs, sub: sub}, nil
}

// WatchAuctionClosed is a free log subscription operation binding the contract event 0x572cdc5ca5e918473319d0f4737494e4709ac879a7d0bcd11ce1bef24b24e81d.
//
// Solidity: event AuctionClosed(uint256 indexed listingId, address indexed closer, bool indexed cancelled, address auctionCreator, address winningBidder)
func (_Marketplace *MarketplaceFilterer) WatchAuctionClosed(opts *bind.WatchOpts, sink chan<- *MarketplaceAuctionClosed, listingId []*big.Int, closer []common.Address, cancelled []bool) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var closerRule []interface{}
	for _, closerItem := range closer {
		closerRule = append(closerRule, closerItem)
	}
	var cancelledRule []interface{}
	for _, cancelledItem := range cancelled {
		cancelledRule = append(cancelledRule, cancelledItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "AuctionClosed", listingIdRule, closerRule, cancelledRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceAuctionClosed)
				if err := _Marketplace.contract.UnpackLog(event, "AuctionClosed", log); err != nil {
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

// ParseAuctionClosed is a log parse operation binding the contract event 0x572cdc5ca5e918473319d0f4737494e4709ac879a7d0bcd11ce1bef24b24e81d.
//
// Solidity: event AuctionClosed(uint256 indexed listingId, address indexed closer, bool indexed cancelled, address auctionCreator, address winningBidder)
func (_Marketplace *MarketplaceFilterer) ParseAuctionClosed(log types.Log) (*MarketplaceAuctionClosed, error) {
	event := new(MarketplaceAuctionClosed)
	if err := _Marketplace.contract.UnpackLog(event, "AuctionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingRestrictedIterator is returned from FilterListingRestricted and is used to iterate over the raw logs and unpacked data for ListingRestricted events raised by the Marketplace contract.
type MarketplaceListingRestrictedIterator struct {
	Event *MarketplaceListingRestricted // Event containing the contract specifics and raw log

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
func (it *MarketplaceListingRestrictedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingRestricted)
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
		it.Event = new(MarketplaceListingRestricted)
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
func (it *MarketplaceListingRestrictedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingRestrictedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingRestricted represents a ListingRestricted event raised by the Marketplace contract.
type MarketplaceListingRestricted struct {
	Restricted bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterListingRestricted is a free log retrieval operation binding the contract event 0x80b4303f755d7d3d4d483a1580281ef7aaeb82947826a1dc63a6366875765cb0.
//
// Solidity: event ListingRestricted(bool restricted)
func (_Marketplace *MarketplaceFilterer) FilterListingRestricted(opts *bind.FilterOpts) (*MarketplaceListingRestrictedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingRestricted")
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingRestrictedIterator{contract: _Marketplace.contract, event: "ListingRestricted", logs: logs, sub: sub}, nil
}

// WatchListingRestricted is a free log subscription operation binding the contract event 0x80b4303f755d7d3d4d483a1580281ef7aaeb82947826a1dc63a6366875765cb0.
//
// Solidity: event ListingRestricted(bool restricted)
func (_Marketplace *MarketplaceFilterer) WatchListingRestricted(opts *bind.WatchOpts, sink chan<- *MarketplaceListingRestricted) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingRestricted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingRestricted)
				if err := _Marketplace.contract.UnpackLog(event, "ListingRestricted", log); err != nil {
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

// ParseListingRestricted is a log parse operation binding the contract event 0x80b4303f755d7d3d4d483a1580281ef7aaeb82947826a1dc63a6366875765cb0.
//
// Solidity: event ListingRestricted(bool restricted)
func (_Marketplace *MarketplaceFilterer) ParseListingRestricted(log types.Log) (*MarketplaceListingRestricted, error) {
	event := new(MarketplaceListingRestricted)
	if err := _Marketplace.contract.UnpackLog(event, "ListingRestricted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingUpdateIterator is returned from FilterListingUpdate and is used to iterate over the raw logs and unpacked data for ListingUpdate events raised by the Marketplace contract.
type MarketplaceListingUpdateIterator struct {
	Event *MarketplaceListingUpdate // Event containing the contract specifics and raw log

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
func (it *MarketplaceListingUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingUpdate)
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
		it.Event = new(MarketplaceListingUpdate)
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
func (it *MarketplaceListingUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingUpdate represents a ListingUpdate event raised by the Marketplace contract.
type MarketplaceListingUpdate struct {
	ListingId      *big.Int
	ListingCreator common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterListingUpdate is a free log retrieval operation binding the contract event 0xa00227275ba75aea329d91406a2884d227dc386f939f1d18e15a7317152432ca.
//
// Solidity: event ListingUpdate(uint256 indexed listingId, address indexed listingCreator)
func (_Marketplace *MarketplaceFilterer) FilterListingUpdate(opts *bind.FilterOpts, listingId []*big.Int, listingCreator []common.Address) (*MarketplaceListingUpdateIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var listingCreatorRule []interface{}
	for _, listingCreatorItem := range listingCreator {
		listingCreatorRule = append(listingCreatorRule, listingCreatorItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingUpdate", listingIdRule, listingCreatorRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingUpdateIterator{contract: _Marketplace.contract, event: "ListingUpdate", logs: logs, sub: sub}, nil
}

// WatchListingUpdate is a free log subscription operation binding the contract event 0xa00227275ba75aea329d91406a2884d227dc386f939f1d18e15a7317152432ca.
//
// Solidity: event ListingUpdate(uint256 indexed listingId, address indexed listingCreator)
func (_Marketplace *MarketplaceFilterer) WatchListingUpdate(opts *bind.WatchOpts, sink chan<- *MarketplaceListingUpdate, listingId []*big.Int, listingCreator []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var listingCreatorRule []interface{}
	for _, listingCreatorItem := range listingCreator {
		listingCreatorRule = append(listingCreatorRule, listingCreatorItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingUpdate", listingIdRule, listingCreatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingUpdate)
				if err := _Marketplace.contract.UnpackLog(event, "ListingUpdate", log); err != nil {
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

// ParseListingUpdate is a log parse operation binding the contract event 0xa00227275ba75aea329d91406a2884d227dc386f939f1d18e15a7317152432ca.
//
// Solidity: event ListingUpdate(uint256 indexed listingId, address indexed listingCreator)
func (_Marketplace *MarketplaceFilterer) ParseListingUpdate(log types.Log) (*MarketplaceListingUpdate, error) {
	event := new(MarketplaceListingUpdate)
	if err := _Marketplace.contract.UnpackLog(event, "ListingUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMarketFeeUpdateIterator is returned from FilterMarketFeeUpdate and is used to iterate over the raw logs and unpacked data for MarketFeeUpdate events raised by the Marketplace contract.
type MarketplaceMarketFeeUpdateIterator struct {
	Event *MarketplaceMarketFeeUpdate // Event containing the contract specifics and raw log

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
func (it *MarketplaceMarketFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMarketFeeUpdate)
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
		it.Event = new(MarketplaceMarketFeeUpdate)
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
func (it *MarketplaceMarketFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMarketFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMarketFeeUpdate represents a MarketFeeUpdate event raised by the Marketplace contract.
type MarketplaceMarketFeeUpdate struct {
	NewFee uint64
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketFeeUpdate is a free log retrieval operation binding the contract event 0x1923ecef8dbc1cebea2768819f7df282b72fb6d62bf99da204590b9d5cac7a7b.
//
// Solidity: event MarketFeeUpdate(uint64 newFee)
func (_Marketplace *MarketplaceFilterer) FilterMarketFeeUpdate(opts *bind.FilterOpts) (*MarketplaceMarketFeeUpdateIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MarketFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMarketFeeUpdateIterator{contract: _Marketplace.contract, event: "MarketFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchMarketFeeUpdate is a free log subscription operation binding the contract event 0x1923ecef8dbc1cebea2768819f7df282b72fb6d62bf99da204590b9d5cac7a7b.
//
// Solidity: event MarketFeeUpdate(uint64 newFee)
func (_Marketplace *MarketplaceFilterer) WatchMarketFeeUpdate(opts *bind.WatchOpts, sink chan<- *MarketplaceMarketFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MarketFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMarketFeeUpdate)
				if err := _Marketplace.contract.UnpackLog(event, "MarketFeeUpdate", log); err != nil {
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

// ParseMarketFeeUpdate is a log parse operation binding the contract event 0x1923ecef8dbc1cebea2768819f7df282b72fb6d62bf99da204590b9d5cac7a7b.
//
// Solidity: event MarketFeeUpdate(uint64 newFee)
func (_Marketplace *MarketplaceFilterer) ParseMarketFeeUpdate(log types.Log) (*MarketplaceMarketFeeUpdate, error) {
	event := new(MarketplaceMarketFeeUpdate)
	if err := _Marketplace.contract.UnpackLog(event, "MarketFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewListingIterator is returned from FilterNewListing and is used to iterate over the raw logs and unpacked data for NewListing events raised by the Marketplace contract.
type MarketplaceNewListingIterator struct {
	Event *MarketplaceNewListing // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewListingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewListing)
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
		it.Event = new(MarketplaceNewListing)
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
func (it *MarketplaceNewListingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewListingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewListing represents a NewListing event raised by the Marketplace contract.
type MarketplaceNewListing struct {
	ListingId     *big.Int
	AssetContract common.Address
	Lister        common.Address
	Listing       IMarketplaceListing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewListing is a free log retrieval operation binding the contract event 0x9e578277632a71dd17ab11c1f584c51deafef022c94389ecb050eb92713725f6.
//
// Solidity: event NewListing(uint256 indexed listingId, address indexed assetContract, address indexed lister, (uint256,address,address,uint256,uint256,uint256,uint256,address,uint256,uint256,uint8,uint8) listing)
func (_Marketplace *MarketplaceFilterer) FilterNewListing(opts *bind.FilterOpts, listingId []*big.Int, assetContract []common.Address, lister []common.Address) (*MarketplaceNewListingIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var listerRule []interface{}
	for _, listerItem := range lister {
		listerRule = append(listerRule, listerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewListing", listingIdRule, assetContractRule, listerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewListingIterator{contract: _Marketplace.contract, event: "NewListing", logs: logs, sub: sub}, nil
}

// WatchNewListing is a free log subscription operation binding the contract event 0x9e578277632a71dd17ab11c1f584c51deafef022c94389ecb050eb92713725f6.
//
// Solidity: event NewListing(uint256 indexed listingId, address indexed assetContract, address indexed lister, (uint256,address,address,uint256,uint256,uint256,uint256,address,uint256,uint256,uint8,uint8) listing)
func (_Marketplace *MarketplaceFilterer) WatchNewListing(opts *bind.WatchOpts, sink chan<- *MarketplaceNewListing, listingId []*big.Int, assetContract []common.Address, lister []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var listerRule []interface{}
	for _, listerItem := range lister {
		listerRule = append(listerRule, listerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewListing", listingIdRule, assetContractRule, listerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewListing)
				if err := _Marketplace.contract.UnpackLog(event, "NewListing", log); err != nil {
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

// ParseNewListing is a log parse operation binding the contract event 0x9e578277632a71dd17ab11c1f584c51deafef022c94389ecb050eb92713725f6.
//
// Solidity: event NewListing(uint256 indexed listingId, address indexed assetContract, address indexed lister, (uint256,address,address,uint256,uint256,uint256,uint256,address,uint256,uint256,uint8,uint8) listing)
func (_Marketplace *MarketplaceFilterer) ParseNewListing(log types.Log) (*MarketplaceNewListing, error) {
	event := new(MarketplaceNewListing)
	if err := _Marketplace.contract.UnpackLog(event, "NewListing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewOfferIterator is returned from FilterNewOffer and is used to iterate over the raw logs and unpacked data for NewOffer events raised by the Marketplace contract.
type MarketplaceNewOfferIterator struct {
	Event *MarketplaceNewOffer // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewOffer)
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
		it.Event = new(MarketplaceNewOffer)
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
func (it *MarketplaceNewOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewOffer represents a NewOffer event raised by the Marketplace contract.
type MarketplaceNewOffer struct {
	ListingId        *big.Int
	Offeror          common.Address
	ListingType      uint8
	QuantityWanted   *big.Int
	TotalOfferAmount *big.Int
	Currency         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewOffer is a free log retrieval operation binding the contract event 0x8a412352601a288b3de40254a9de2ab14a497aa3638a7e558480680a56e2705d.
//
// Solidity: event NewOffer(uint256 indexed listingId, address indexed offeror, uint8 indexed listingType, uint256 quantityWanted, uint256 totalOfferAmount, address currency)
func (_Marketplace *MarketplaceFilterer) FilterNewOffer(opts *bind.FilterOpts, listingId []*big.Int, offeror []common.Address, listingType []uint8) (*MarketplaceNewOfferIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var offerorRule []interface{}
	for _, offerorItem := range offeror {
		offerorRule = append(offerorRule, offerorItem)
	}
	var listingTypeRule []interface{}
	for _, listingTypeItem := range listingType {
		listingTypeRule = append(listingTypeRule, listingTypeItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewOffer", listingIdRule, offerorRule, listingTypeRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewOfferIterator{contract: _Marketplace.contract, event: "NewOffer", logs: logs, sub: sub}, nil
}

// WatchNewOffer is a free log subscription operation binding the contract event 0x8a412352601a288b3de40254a9de2ab14a497aa3638a7e558480680a56e2705d.
//
// Solidity: event NewOffer(uint256 indexed listingId, address indexed offeror, uint8 indexed listingType, uint256 quantityWanted, uint256 totalOfferAmount, address currency)
func (_Marketplace *MarketplaceFilterer) WatchNewOffer(opts *bind.WatchOpts, sink chan<- *MarketplaceNewOffer, listingId []*big.Int, offeror []common.Address, listingType []uint8) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var offerorRule []interface{}
	for _, offerorItem := range offeror {
		offerorRule = append(offerorRule, offerorItem)
	}
	var listingTypeRule []interface{}
	for _, listingTypeItem := range listingType {
		listingTypeRule = append(listingTypeRule, listingTypeItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewOffer", listingIdRule, offerorRule, listingTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewOffer)
				if err := _Marketplace.contract.UnpackLog(event, "NewOffer", log); err != nil {
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

// ParseNewOffer is a log parse operation binding the contract event 0x8a412352601a288b3de40254a9de2ab14a497aa3638a7e558480680a56e2705d.
//
// Solidity: event NewOffer(uint256 indexed listingId, address indexed offeror, uint8 indexed listingType, uint256 quantityWanted, uint256 totalOfferAmount, address currency)
func (_Marketplace *MarketplaceFilterer) ParseNewOffer(log types.Log) (*MarketplaceNewOffer, error) {
	event := new(MarketplaceNewOffer)
	if err := _Marketplace.contract.UnpackLog(event, "NewOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNewSaleIterator is returned from FilterNewSale and is used to iterate over the raw logs and unpacked data for NewSale events raised by the Marketplace contract.
type MarketplaceNewSaleIterator struct {
	Event *MarketplaceNewSale // Event containing the contract specifics and raw log

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
func (it *MarketplaceNewSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNewSale)
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
		it.Event = new(MarketplaceNewSale)
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
func (it *MarketplaceNewSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNewSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNewSale represents a NewSale event raised by the Marketplace contract.
type MarketplaceNewSale struct {
	ListingId      *big.Int
	AssetContract  common.Address
	Lister         common.Address
	Buyer          common.Address
	QuantityBought *big.Int
	TotalPricePaid *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewSale is a free log retrieval operation binding the contract event 0x306e6cde5eb293794d557a3a6c844de939e6206b05e6910451c512852bf654a5.
//
// Solidity: event NewSale(uint256 indexed listingId, address indexed assetContract, address indexed lister, address buyer, uint256 quantityBought, uint256 totalPricePaid)
func (_Marketplace *MarketplaceFilterer) FilterNewSale(opts *bind.FilterOpts, listingId []*big.Int, assetContract []common.Address, lister []common.Address) (*MarketplaceNewSaleIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var listerRule []interface{}
	for _, listerItem := range lister {
		listerRule = append(listerRule, listerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NewSale", listingIdRule, assetContractRule, listerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNewSaleIterator{contract: _Marketplace.contract, event: "NewSale", logs: logs, sub: sub}, nil
}

// WatchNewSale is a free log subscription operation binding the contract event 0x306e6cde5eb293794d557a3a6c844de939e6206b05e6910451c512852bf654a5.
//
// Solidity: event NewSale(uint256 indexed listingId, address indexed assetContract, address indexed lister, address buyer, uint256 quantityBought, uint256 totalPricePaid)
func (_Marketplace *MarketplaceFilterer) WatchNewSale(opts *bind.WatchOpts, sink chan<- *MarketplaceNewSale, listingId []*big.Int, assetContract []common.Address, lister []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var assetContractRule []interface{}
	for _, assetContractItem := range assetContract {
		assetContractRule = append(assetContractRule, assetContractItem)
	}
	var listerRule []interface{}
	for _, listerItem := range lister {
		listerRule = append(listerRule, listerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NewSale", listingIdRule, assetContractRule, listerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNewSale)
				if err := _Marketplace.contract.UnpackLog(event, "NewSale", log); err != nil {
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

// ParseNewSale is a log parse operation binding the contract event 0x306e6cde5eb293794d557a3a6c844de939e6206b05e6910451c512852bf654a5.
//
// Solidity: event NewSale(uint256 indexed listingId, address indexed assetContract, address indexed lister, address buyer, uint256 quantityBought, uint256 totalPricePaid)
func (_Marketplace *MarketplaceFilterer) ParseNewSale(log types.Log) (*MarketplaceNewSale, error) {
	event := new(MarketplaceNewSale)
	if err := _Marketplace.contract.UnpackLog(event, "NewSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Marketplace contract.
type MarketplaceRoleAdminChangedIterator struct {
	Event *MarketplaceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarketplaceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRoleAdminChanged)
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
		it.Event = new(MarketplaceRoleAdminChanged)
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
func (it *MarketplaceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRoleAdminChanged represents a RoleAdminChanged event raised by the Marketplace contract.
type MarketplaceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplace *MarketplaceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MarketplaceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceRoleAdminChangedIterator{contract: _Marketplace.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Marketplace *MarketplaceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MarketplaceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRoleAdminChanged)
				if err := _Marketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRoleAdminChanged(log types.Log) (*MarketplaceRoleAdminChanged, error) {
	event := new(MarketplaceRoleAdminChanged)
	if err := _Marketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Marketplace contract.
type MarketplaceRoleGrantedIterator struct {
	Event *MarketplaceRoleGranted // Event containing the contract specifics and raw log

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
func (it *MarketplaceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRoleGranted)
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
		it.Event = new(MarketplaceRoleGranted)
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
func (it *MarketplaceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRoleGranted represents a RoleGranted event raised by the Marketplace contract.
type MarketplaceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketplaceRoleGrantedIterator, error) {

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

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceRoleGrantedIterator{contract: _Marketplace.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MarketplaceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRoleGranted)
				if err := _Marketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRoleGranted(log types.Log) (*MarketplaceRoleGranted, error) {
	event := new(MarketplaceRoleGranted)
	if err := _Marketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Marketplace contract.
type MarketplaceRoleRevokedIterator struct {
	Event *MarketplaceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MarketplaceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceRoleRevoked)
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
		it.Event = new(MarketplaceRoleRevoked)
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
func (it *MarketplaceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceRoleRevoked represents a RoleRevoked event raised by the Marketplace contract.
type MarketplaceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketplaceRoleRevokedIterator, error) {

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

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceRoleRevokedIterator{contract: _Marketplace.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Marketplace *MarketplaceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MarketplaceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceRoleRevoked)
				if err := _Marketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseRoleRevoked(log types.Log) (*MarketplaceRoleRevoked, error) {
	event := new(MarketplaceRoleRevoked)
	if err := _Marketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
