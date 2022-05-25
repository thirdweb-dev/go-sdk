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

// IDropClaimConditionClaimCondition is an auto generated low-level Go binding around an user-defined struct.
type IDropClaimConditionClaimCondition struct {
	StartTimestamp                 *big.Int
	MaxClaimableSupply             *big.Int
	SupplyClaimed                  *big.Int
	QuantityLimitPerTransaction    *big.Int
	WaitTimeInSecondsBetweenClaims *big.Int
	MerkleRoot                     [32]byte
	PricePerToken                  *big.Int
	Currency                       common.Address
}

// DropERC1155MetaData contains all meta data concerning the DropERC1155 contract.
var DropERC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_thirdwebFee\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerTransaction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"waitTimeInSecondsBetweenClaims\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structIDropClaimCondition.ClaimCondition[]\",\"name\":\"claimConditions\",\"type\":\"tuple[]\"}],\"name\":\"ClaimConditionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRoyaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoyaltyBps\",\"type\":\"uint256\"}],\"name\":\"DefaultRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"MaxTotalSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"MaxWalletClaimCountUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"}],\"name\":\"PlatformFeeInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"PrimarySaleRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"claimConditionIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityClaimed\",\"type\":\"uint256\"}],\"name\":\"TokensClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"TokensLazyMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"WalletClaimCountUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"burnBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_proofMaxQuantityPerTransaction\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"claimCondition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentStartId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getActiveClaimConditionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"}],\"name\":\"getClaimConditionById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerTransaction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"waitTimeInSecondsBetweenClaims\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIDropClaimCondition.ClaimCondition\",\"name\":\"condition\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"}],\"name\":\"getClaimTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lastClaimTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextValidClaimTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyInfoForToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_trustedForwarders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_saleRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_royaltyBps\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_platformFeeBps\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_baseURIForTokens\",\"type\":\"string\"}],\"name\":\"lazyMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maxWalletClaimCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenIdToMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"primarySaleRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"saleRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerTransaction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"waitTimeInSecondsBetweenClaims\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIDropClaimCondition.ClaimCondition[]\",\"name\":\"_phases\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"_resetClaimEligibility\",\"type\":\"bool\"}],\"name\":\"setClaimConditions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"setMaxWalletClaimCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFeeBps\",\"type\":\"uint256\"}],\"name\":\"setPlatformFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_saleRecipient\",\"type\":\"address\"}],\"name\":\"setPrimarySaleRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyInfoForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"setWalletClaimCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"verifyMaxQuantityPerTransaction\",\"type\":\"bool\"}],\"name\":\"verifyClaim\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_proofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"_proofMaxQuantityPerTransaction\",\"type\":\"uint256\"}],\"name\":\"verifyClaimMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"validMerkleProof\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"merkleProofIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletClaimCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DropERC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use DropERC1155MetaData.ABI instead.
var DropERC1155ABI = DropERC1155MetaData.ABI

// DropERC1155 is an auto generated Go binding around an Ethereum contract.
type DropERC1155 struct {
	DropERC1155Caller     // Read-only binding to the contract
	DropERC1155Transactor // Write-only binding to the contract
	DropERC1155Filterer   // Log filterer for contract events
}

// DropERC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type DropERC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DropERC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DropERC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DropERC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DropERC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DropERC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DropERC1155Session struct {
	Contract     *DropERC1155      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DropERC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DropERC1155CallerSession struct {
	Contract *DropERC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DropERC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DropERC1155TransactorSession struct {
	Contract     *DropERC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DropERC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type DropERC1155Raw struct {
	Contract *DropERC1155 // Generic contract binding to access the raw methods on
}

// DropERC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DropERC1155CallerRaw struct {
	Contract *DropERC1155Caller // Generic read-only contract binding to access the raw methods on
}

// DropERC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DropERC1155TransactorRaw struct {
	Contract *DropERC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDropERC1155 creates a new instance of DropERC1155, bound to a specific deployed contract.
func NewDropERC1155(address common.Address, backend bind.ContractBackend) (*DropERC1155, error) {
	contract, err := bindDropERC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DropERC1155{DropERC1155Caller: DropERC1155Caller{contract: contract}, DropERC1155Transactor: DropERC1155Transactor{contract: contract}, DropERC1155Filterer: DropERC1155Filterer{contract: contract}}, nil
}

// NewDropERC1155Caller creates a new read-only instance of DropERC1155, bound to a specific deployed contract.
func NewDropERC1155Caller(address common.Address, caller bind.ContractCaller) (*DropERC1155Caller, error) {
	contract, err := bindDropERC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DropERC1155Caller{contract: contract}, nil
}

// NewDropERC1155Transactor creates a new write-only instance of DropERC1155, bound to a specific deployed contract.
func NewDropERC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*DropERC1155Transactor, error) {
	contract, err := bindDropERC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DropERC1155Transactor{contract: contract}, nil
}

// NewDropERC1155Filterer creates a new log filterer instance of DropERC1155, bound to a specific deployed contract.
func NewDropERC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*DropERC1155Filterer, error) {
	contract, err := bindDropERC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DropERC1155Filterer{contract: contract}, nil
}

// bindDropERC1155 binds a generic wrapper to an already deployed contract.
func bindDropERC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DropERC1155ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DropERC1155 *DropERC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DropERC1155.Contract.DropERC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DropERC1155 *DropERC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DropERC1155.Contract.DropERC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DropERC1155 *DropERC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DropERC1155.Contract.DropERC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DropERC1155 *DropERC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DropERC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DropERC1155 *DropERC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DropERC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DropERC1155 *DropERC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DropERC1155.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DropERC1155 *DropERC1155Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DropERC1155 *DropERC1155Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _DropERC1155.Contract.DEFAULTADMINROLE(&_DropERC1155.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DropERC1155 *DropERC1155CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DropERC1155.Contract.DEFAULTADMINROLE(&_DropERC1155.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.BalanceOf(&_DropERC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.BalanceOf(&_DropERC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_DropERC1155 *DropERC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_DropERC1155 *DropERC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _DropERC1155.Contract.BalanceOfBatch(&_DropERC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_DropERC1155 *DropERC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _DropERC1155.Contract.BalanceOfBatch(&_DropERC1155.CallOpts, accounts, ids)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xe9703d25.
//
// Solidity: function claimCondition(uint256 ) view returns(uint256 currentStartId, uint256 count)
func (_DropERC1155 *DropERC1155Caller) ClaimCondition(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "claimCondition", arg0)

	outstruct := new(struct {
		CurrentStartId *big.Int
		Count          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentStartId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ClaimCondition is a free data retrieval call binding the contract method 0xe9703d25.
//
// Solidity: function claimCondition(uint256 ) view returns(uint256 currentStartId, uint256 count)
func (_DropERC1155 *DropERC1155Session) ClaimCondition(arg0 *big.Int) (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _DropERC1155.Contract.ClaimCondition(&_DropERC1155.CallOpts, arg0)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xe9703d25.
//
// Solidity: function claimCondition(uint256 ) view returns(uint256 currentStartId, uint256 count)
func (_DropERC1155 *DropERC1155CallerSession) ClaimCondition(arg0 *big.Int) (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _DropERC1155.Contract.ClaimCondition(&_DropERC1155.CallOpts, arg0)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_DropERC1155 *DropERC1155Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_DropERC1155 *DropERC1155Session) ContractType() ([32]byte, error) {
	return _DropERC1155.Contract.ContractType(&_DropERC1155.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_DropERC1155 *DropERC1155CallerSession) ContractType() ([32]byte, error) {
	return _DropERC1155.Contract.ContractType(&_DropERC1155.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_DropERC1155 *DropERC1155Caller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_DropERC1155 *DropERC1155Session) InternalContractURI() (string, error) {
	return _DropERC1155.Contract.InternalContractURI(&_DropERC1155.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_DropERC1155 *DropERC1155CallerSession) InternalContractURI() (string, error) {
	return _DropERC1155.Contract.InternalContractURI(&_DropERC1155.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_DropERC1155 *DropERC1155Caller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_DropERC1155 *DropERC1155Session) ContractVersion() (uint8, error) {
	return _DropERC1155.Contract.ContractVersion(&_DropERC1155.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_DropERC1155 *DropERC1155CallerSession) ContractVersion() (uint8, error) {
	return _DropERC1155.Contract.ContractVersion(&_DropERC1155.CallOpts)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0x5ab063e8.
//
// Solidity: function getActiveClaimConditionId(uint256 _tokenId) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) GetActiveClaimConditionId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getActiveClaimConditionId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0x5ab063e8.
//
// Solidity: function getActiveClaimConditionId(uint256 _tokenId) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) GetActiveClaimConditionId(_tokenId *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.GetActiveClaimConditionId(&_DropERC1155.CallOpts, _tokenId)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0x5ab063e8.
//
// Solidity: function getActiveClaimConditionId(uint256 _tokenId) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) GetActiveClaimConditionId(_tokenId *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.GetActiveClaimConditionId(&_DropERC1155.CallOpts, _tokenId)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0xd45b28d7.
//
// Solidity: function getClaimConditionById(uint256 _tokenId, uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address) condition)
func (_DropERC1155 *DropERC1155Caller) GetClaimConditionById(opts *bind.CallOpts, _tokenId *big.Int, _conditionId *big.Int) (IDropClaimConditionClaimCondition, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getClaimConditionById", _tokenId, _conditionId)

	if err != nil {
		return *new(IDropClaimConditionClaimCondition), err
	}

	out0 := *abi.ConvertType(out[0], new(IDropClaimConditionClaimCondition)).(*IDropClaimConditionClaimCondition)

	return out0, err

}

// GetClaimConditionById is a free data retrieval call binding the contract method 0xd45b28d7.
//
// Solidity: function getClaimConditionById(uint256 _tokenId, uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address) condition)
func (_DropERC1155 *DropERC1155Session) GetClaimConditionById(_tokenId *big.Int, _conditionId *big.Int) (IDropClaimConditionClaimCondition, error) {
	return _DropERC1155.Contract.GetClaimConditionById(&_DropERC1155.CallOpts, _tokenId, _conditionId)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0xd45b28d7.
//
// Solidity: function getClaimConditionById(uint256 _tokenId, uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address) condition)
func (_DropERC1155 *DropERC1155CallerSession) GetClaimConditionById(_tokenId *big.Int, _conditionId *big.Int) (IDropClaimConditionClaimCondition, error) {
	return _DropERC1155.Contract.GetClaimConditionById(&_DropERC1155.CallOpts, _tokenId, _conditionId)
}

// GetClaimTimestamp is a free data retrieval call binding the contract method 0x622a6c31.
//
// Solidity: function getClaimTimestamp(uint256 _tokenId, uint256 _conditionId, address _claimer) view returns(uint256 lastClaimTimestamp, uint256 nextValidClaimTimestamp)
func (_DropERC1155 *DropERC1155Caller) GetClaimTimestamp(opts *bind.CallOpts, _tokenId *big.Int, _conditionId *big.Int, _claimer common.Address) (struct {
	LastClaimTimestamp      *big.Int
	NextValidClaimTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getClaimTimestamp", _tokenId, _conditionId, _claimer)

	outstruct := new(struct {
		LastClaimTimestamp      *big.Int
		NextValidClaimTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LastClaimTimestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextValidClaimTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetClaimTimestamp is a free data retrieval call binding the contract method 0x622a6c31.
//
// Solidity: function getClaimTimestamp(uint256 _tokenId, uint256 _conditionId, address _claimer) view returns(uint256 lastClaimTimestamp, uint256 nextValidClaimTimestamp)
func (_DropERC1155 *DropERC1155Session) GetClaimTimestamp(_tokenId *big.Int, _conditionId *big.Int, _claimer common.Address) (struct {
	LastClaimTimestamp      *big.Int
	NextValidClaimTimestamp *big.Int
}, error) {
	return _DropERC1155.Contract.GetClaimTimestamp(&_DropERC1155.CallOpts, _tokenId, _conditionId, _claimer)
}

// GetClaimTimestamp is a free data retrieval call binding the contract method 0x622a6c31.
//
// Solidity: function getClaimTimestamp(uint256 _tokenId, uint256 _conditionId, address _claimer) view returns(uint256 lastClaimTimestamp, uint256 nextValidClaimTimestamp)
func (_DropERC1155 *DropERC1155CallerSession) GetClaimTimestamp(_tokenId *big.Int, _conditionId *big.Int, _claimer common.Address) (struct {
	LastClaimTimestamp      *big.Int
	NextValidClaimTimestamp *big.Int
}, error) {
	return _DropERC1155.Contract.GetClaimTimestamp(&_DropERC1155.CallOpts, _tokenId, _conditionId, _claimer)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_DropERC1155 *DropERC1155Caller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

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
func (_DropERC1155 *DropERC1155Session) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _DropERC1155.Contract.GetDefaultRoyaltyInfo(&_DropERC1155.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_DropERC1155 *DropERC1155CallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _DropERC1155.Contract.GetDefaultRoyaltyInfo(&_DropERC1155.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_DropERC1155 *DropERC1155Caller) GetPlatformFeeInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getPlatformFeeInfo")

	if err != nil {
		return *new(common.Address), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return out0, out1, err

}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_DropERC1155 *DropERC1155Session) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _DropERC1155.Contract.GetPlatformFeeInfo(&_DropERC1155.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_DropERC1155 *DropERC1155CallerSession) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _DropERC1155.Contract.GetPlatformFeeInfo(&_DropERC1155.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DropERC1155 *DropERC1155Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DropERC1155 *DropERC1155Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DropERC1155.Contract.GetRoleAdmin(&_DropERC1155.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DropERC1155 *DropERC1155CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DropERC1155.Contract.GetRoleAdmin(&_DropERC1155.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DropERC1155 *DropERC1155Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DropERC1155 *DropERC1155Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DropERC1155.Contract.GetRoleMember(&_DropERC1155.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_DropERC1155 *DropERC1155CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DropERC1155.Contract.GetRoleMember(&_DropERC1155.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DropERC1155.Contract.GetRoleMemberCount(&_DropERC1155.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DropERC1155.Contract.GetRoleMemberCount(&_DropERC1155.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_DropERC1155 *DropERC1155Caller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

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
func (_DropERC1155 *DropERC1155Session) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _DropERC1155.Contract.GetRoyaltyInfoForToken(&_DropERC1155.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_DropERC1155 *DropERC1155CallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _DropERC1155.Contract.GetRoyaltyInfoForToken(&_DropERC1155.CallOpts, _tokenId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DropERC1155 *DropERC1155Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DropERC1155 *DropERC1155Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DropERC1155.Contract.HasRole(&_DropERC1155.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DropERC1155 *DropERC1155CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DropERC1155.Contract.HasRole(&_DropERC1155.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_DropERC1155 *DropERC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_DropERC1155 *DropERC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _DropERC1155.Contract.IsApprovedForAll(&_DropERC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_DropERC1155 *DropERC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _DropERC1155.Contract.IsApprovedForAll(&_DropERC1155.CallOpts, account, operator)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_DropERC1155 *DropERC1155Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_DropERC1155 *DropERC1155Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _DropERC1155.Contract.IsTrustedForwarder(&_DropERC1155.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_DropERC1155 *DropERC1155CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _DropERC1155.Contract.IsTrustedForwarder(&_DropERC1155.CallOpts, forwarder)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x24aaffaa.
//
// Solidity: function maxTotalSupply(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) MaxTotalSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "maxTotalSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x24aaffaa.
//
// Solidity: function maxTotalSupply(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) MaxTotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.MaxTotalSupply(&_DropERC1155.CallOpts, arg0)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x24aaffaa.
//
// Solidity: function maxTotalSupply(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) MaxTotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.MaxTotalSupply(&_DropERC1155.CallOpts, arg0)
}

// MaxWalletClaimCount is a free data retrieval call binding the contract method 0xb79cade4.
//
// Solidity: function maxWalletClaimCount(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) MaxWalletClaimCount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "maxWalletClaimCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxWalletClaimCount is a free data retrieval call binding the contract method 0xb79cade4.
//
// Solidity: function maxWalletClaimCount(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) MaxWalletClaimCount(arg0 *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.MaxWalletClaimCount(&_DropERC1155.CallOpts, arg0)
}

// MaxWalletClaimCount is a free data retrieval call binding the contract method 0xb79cade4.
//
// Solidity: function maxWalletClaimCount(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) MaxWalletClaimCount(arg0 *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.MaxWalletClaimCount(&_DropERC1155.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DropERC1155 *DropERC1155Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DropERC1155 *DropERC1155Session) Name() (string, error) {
	return _DropERC1155.Contract.Name(&_DropERC1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DropERC1155 *DropERC1155CallerSession) Name() (string, error) {
	return _DropERC1155.Contract.Name(&_DropERC1155.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_DropERC1155 *DropERC1155Session) NextTokenIdToMint() (*big.Int, error) {
	return _DropERC1155.Contract.NextTokenIdToMint(&_DropERC1155.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _DropERC1155.Contract.NextTokenIdToMint(&_DropERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DropERC1155 *DropERC1155Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DropERC1155 *DropERC1155Session) Owner() (common.Address, error) {
	return _DropERC1155.Contract.Owner(&_DropERC1155.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DropERC1155 *DropERC1155CallerSession) Owner() (common.Address, error) {
	return _DropERC1155.Contract.Owner(&_DropERC1155.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_DropERC1155 *DropERC1155Caller) PrimarySaleRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "primarySaleRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_DropERC1155 *DropERC1155Session) PrimarySaleRecipient() (common.Address, error) {
	return _DropERC1155.Contract.PrimarySaleRecipient(&_DropERC1155.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_DropERC1155 *DropERC1155CallerSession) PrimarySaleRecipient() (common.Address, error) {
	return _DropERC1155.Contract.PrimarySaleRecipient(&_DropERC1155.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_DropERC1155 *DropERC1155Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_DropERC1155 *DropERC1155Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _DropERC1155.Contract.RoyaltyInfo(&_DropERC1155.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_DropERC1155 *DropERC1155CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _DropERC1155.Contract.RoyaltyInfo(&_DropERC1155.CallOpts, tokenId, salePrice)
}

// SaleRecipient is a free data retrieval call binding the contract method 0xc7337d6b.
//
// Solidity: function saleRecipient(uint256 ) view returns(address)
func (_DropERC1155 *DropERC1155Caller) SaleRecipient(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "saleRecipient", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleRecipient is a free data retrieval call binding the contract method 0xc7337d6b.
//
// Solidity: function saleRecipient(uint256 ) view returns(address)
func (_DropERC1155 *DropERC1155Session) SaleRecipient(arg0 *big.Int) (common.Address, error) {
	return _DropERC1155.Contract.SaleRecipient(&_DropERC1155.CallOpts, arg0)
}

// SaleRecipient is a free data retrieval call binding the contract method 0xc7337d6b.
//
// Solidity: function saleRecipient(uint256 ) view returns(address)
func (_DropERC1155 *DropERC1155CallerSession) SaleRecipient(arg0 *big.Int) (common.Address, error) {
	return _DropERC1155.Contract.SaleRecipient(&_DropERC1155.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DropERC1155 *DropERC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DropERC1155 *DropERC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DropERC1155.Contract.SupportsInterface(&_DropERC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DropERC1155 *DropERC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DropERC1155.Contract.SupportsInterface(&_DropERC1155.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DropERC1155 *DropERC1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DropERC1155 *DropERC1155Session) Symbol() (string, error) {
	return _DropERC1155.Contract.Symbol(&_DropERC1155.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DropERC1155 *DropERC1155CallerSession) Symbol() (string, error) {
	return _DropERC1155.Contract.Symbol(&_DropERC1155.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) TotalSupply(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "totalSupply", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) TotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.TotalSupply(&_DropERC1155.CallOpts, arg0)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 ) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) TotalSupply(arg0 *big.Int) (*big.Int, error) {
	return _DropERC1155.Contract.TotalSupply(&_DropERC1155.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string _tokenURI)
func (_DropERC1155 *DropERC1155Caller) Uri(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "uri", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string _tokenURI)
func (_DropERC1155 *DropERC1155Session) Uri(_tokenId *big.Int) (string, error) {
	return _DropERC1155.Contract.Uri(&_DropERC1155.CallOpts, _tokenId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _tokenId) view returns(string _tokenURI)
func (_DropERC1155 *DropERC1155CallerSession) Uri(_tokenId *big.Int) (string, error) {
	return _DropERC1155.Contract.Uri(&_DropERC1155.CallOpts, _tokenId)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xa157f71c.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _tokenId, uint256 _quantity, address _currency, uint256 _pricePerToken, bool verifyMaxQuantityPerTransaction) view returns()
func (_DropERC1155 *DropERC1155Caller) VerifyClaim(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address, _tokenId *big.Int, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, verifyMaxQuantityPerTransaction bool) error {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "verifyClaim", _conditionId, _claimer, _tokenId, _quantity, _currency, _pricePerToken, verifyMaxQuantityPerTransaction)

	if err != nil {
		return err
	}

	return err

}

// VerifyClaim is a free data retrieval call binding the contract method 0xa157f71c.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _tokenId, uint256 _quantity, address _currency, uint256 _pricePerToken, bool verifyMaxQuantityPerTransaction) view returns()
func (_DropERC1155 *DropERC1155Session) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _tokenId *big.Int, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, verifyMaxQuantityPerTransaction bool) error {
	return _DropERC1155.Contract.VerifyClaim(&_DropERC1155.CallOpts, _conditionId, _claimer, _tokenId, _quantity, _currency, _pricePerToken, verifyMaxQuantityPerTransaction)
}

// VerifyClaim is a free data retrieval call binding the contract method 0xa157f71c.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _tokenId, uint256 _quantity, address _currency, uint256 _pricePerToken, bool verifyMaxQuantityPerTransaction) view returns()
func (_DropERC1155 *DropERC1155CallerSession) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _tokenId *big.Int, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, verifyMaxQuantityPerTransaction bool) error {
	return _DropERC1155.Contract.VerifyClaim(&_DropERC1155.CallOpts, _conditionId, _claimer, _tokenId, _quantity, _currency, _pricePerToken, verifyMaxQuantityPerTransaction)
}

// VerifyClaimMerkleProof is a free data retrieval call binding the contract method 0x71d53a5b.
//
// Solidity: function verifyClaimMerkleProof(uint256 _conditionId, address _claimer, uint256 _tokenId, uint256 _quantity, bytes32[] _proofs, uint256 _proofMaxQuantityPerTransaction) view returns(bool validMerkleProof, uint256 merkleProofIndex)
func (_DropERC1155 *DropERC1155Caller) VerifyClaimMerkleProof(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address, _tokenId *big.Int, _quantity *big.Int, _proofs [][32]byte, _proofMaxQuantityPerTransaction *big.Int) (struct {
	ValidMerkleProof bool
	MerkleProofIndex *big.Int
}, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "verifyClaimMerkleProof", _conditionId, _claimer, _tokenId, _quantity, _proofs, _proofMaxQuantityPerTransaction)

	outstruct := new(struct {
		ValidMerkleProof bool
		MerkleProofIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ValidMerkleProof = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.MerkleProofIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// VerifyClaimMerkleProof is a free data retrieval call binding the contract method 0x71d53a5b.
//
// Solidity: function verifyClaimMerkleProof(uint256 _conditionId, address _claimer, uint256 _tokenId, uint256 _quantity, bytes32[] _proofs, uint256 _proofMaxQuantityPerTransaction) view returns(bool validMerkleProof, uint256 merkleProofIndex)
func (_DropERC1155 *DropERC1155Session) VerifyClaimMerkleProof(_conditionId *big.Int, _claimer common.Address, _tokenId *big.Int, _quantity *big.Int, _proofs [][32]byte, _proofMaxQuantityPerTransaction *big.Int) (struct {
	ValidMerkleProof bool
	MerkleProofIndex *big.Int
}, error) {
	return _DropERC1155.Contract.VerifyClaimMerkleProof(&_DropERC1155.CallOpts, _conditionId, _claimer, _tokenId, _quantity, _proofs, _proofMaxQuantityPerTransaction)
}

// VerifyClaimMerkleProof is a free data retrieval call binding the contract method 0x71d53a5b.
//
// Solidity: function verifyClaimMerkleProof(uint256 _conditionId, address _claimer, uint256 _tokenId, uint256 _quantity, bytes32[] _proofs, uint256 _proofMaxQuantityPerTransaction) view returns(bool validMerkleProof, uint256 merkleProofIndex)
func (_DropERC1155 *DropERC1155CallerSession) VerifyClaimMerkleProof(_conditionId *big.Int, _claimer common.Address, _tokenId *big.Int, _quantity *big.Int, _proofs [][32]byte, _proofMaxQuantityPerTransaction *big.Int) (struct {
	ValidMerkleProof bool
	MerkleProofIndex *big.Int
}, error) {
	return _DropERC1155.Contract.VerifyClaimMerkleProof(&_DropERC1155.CallOpts, _conditionId, _claimer, _tokenId, _quantity, _proofs, _proofMaxQuantityPerTransaction)
}

// WalletClaimCount is a free data retrieval call binding the contract method 0xc16ce64e.
//
// Solidity: function walletClaimCount(uint256 , address ) view returns(uint256)
func (_DropERC1155 *DropERC1155Caller) WalletClaimCount(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DropERC1155.contract.Call(opts, &out, "walletClaimCount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WalletClaimCount is a free data retrieval call binding the contract method 0xc16ce64e.
//
// Solidity: function walletClaimCount(uint256 , address ) view returns(uint256)
func (_DropERC1155 *DropERC1155Session) WalletClaimCount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _DropERC1155.Contract.WalletClaimCount(&_DropERC1155.CallOpts, arg0, arg1)
}

// WalletClaimCount is a free data retrieval call binding the contract method 0xc16ce64e.
//
// Solidity: function walletClaimCount(uint256 , address ) view returns(uint256)
func (_DropERC1155 *DropERC1155CallerSession) WalletClaimCount(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _DropERC1155.Contract.WalletClaimCount(&_DropERC1155.CallOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_DropERC1155 *DropERC1155Transactor) Burn(opts *bind.TransactOpts, account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "burn", account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_DropERC1155 *DropERC1155Session) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.Burn(&_DropERC1155.TransactOpts, account, id, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address account, uint256 id, uint256 value) returns()
func (_DropERC1155 *DropERC1155TransactorSession) Burn(account common.Address, id *big.Int, value *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.Burn(&_DropERC1155.TransactOpts, account, id, value)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_DropERC1155 *DropERC1155Transactor) BurnBatch(opts *bind.TransactOpts, account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "burnBatch", account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_DropERC1155 *DropERC1155Session) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.BurnBatch(&_DropERC1155.TransactOpts, account, ids, values)
}

// BurnBatch is a paid mutator transaction binding the contract method 0x6b20c454.
//
// Solidity: function burnBatch(address account, uint256[] ids, uint256[] values) returns()
func (_DropERC1155 *DropERC1155TransactorSession) BurnBatch(account common.Address, ids []*big.Int, values []*big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.BurnBatch(&_DropERC1155.TransactOpts, account, ids, values)
}

// Claim is a paid mutator transaction binding the contract method 0xb4c5faa1.
//
// Solidity: function claim(address _receiver, uint256 _tokenId, uint256 _quantity, address _currency, uint256 _pricePerToken, bytes32[] _proofs, uint256 _proofMaxQuantityPerTransaction) payable returns()
func (_DropERC1155 *DropERC1155Transactor) Claim(opts *bind.TransactOpts, _receiver common.Address, _tokenId *big.Int, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _proofs [][32]byte, _proofMaxQuantityPerTransaction *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "claim", _receiver, _tokenId, _quantity, _currency, _pricePerToken, _proofs, _proofMaxQuantityPerTransaction)
}

// Claim is a paid mutator transaction binding the contract method 0xb4c5faa1.
//
// Solidity: function claim(address _receiver, uint256 _tokenId, uint256 _quantity, address _currency, uint256 _pricePerToken, bytes32[] _proofs, uint256 _proofMaxQuantityPerTransaction) payable returns()
func (_DropERC1155 *DropERC1155Session) Claim(_receiver common.Address, _tokenId *big.Int, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _proofs [][32]byte, _proofMaxQuantityPerTransaction *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.Claim(&_DropERC1155.TransactOpts, _receiver, _tokenId, _quantity, _currency, _pricePerToken, _proofs, _proofMaxQuantityPerTransaction)
}

// Claim is a paid mutator transaction binding the contract method 0xb4c5faa1.
//
// Solidity: function claim(address _receiver, uint256 _tokenId, uint256 _quantity, address _currency, uint256 _pricePerToken, bytes32[] _proofs, uint256 _proofMaxQuantityPerTransaction) payable returns()
func (_DropERC1155 *DropERC1155TransactorSession) Claim(_receiver common.Address, _tokenId *big.Int, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _proofs [][32]byte, _proofMaxQuantityPerTransaction *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.Claim(&_DropERC1155.TransactOpts, _receiver, _tokenId, _quantity, _currency, _pricePerToken, _proofs, _proofMaxQuantityPerTransaction)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.GrantRole(&_DropERC1155.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.GrantRole(&_DropERC1155.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_DropERC1155 *DropERC1155Transactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_DropERC1155 *DropERC1155Session) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.Initialize(&_DropERC1155.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_DropERC1155 *DropERC1155TransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.Initialize(&_DropERC1155.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// LazyMint is a paid mutator transaction binding the contract method 0x47158264.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens) returns()
func (_DropERC1155 *DropERC1155Transactor) LazyMint(opts *bind.TransactOpts, _amount *big.Int, _baseURIForTokens string) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "lazyMint", _amount, _baseURIForTokens)
}

// LazyMint is a paid mutator transaction binding the contract method 0x47158264.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens) returns()
func (_DropERC1155 *DropERC1155Session) LazyMint(_amount *big.Int, _baseURIForTokens string) (*types.Transaction, error) {
	return _DropERC1155.Contract.LazyMint(&_DropERC1155.TransactOpts, _amount, _baseURIForTokens)
}

// LazyMint is a paid mutator transaction binding the contract method 0x47158264.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens) returns()
func (_DropERC1155 *DropERC1155TransactorSession) LazyMint(_amount *big.Int, _baseURIForTokens string) (*types.Transaction, error) {
	return _DropERC1155.Contract.LazyMint(&_DropERC1155.TransactOpts, _amount, _baseURIForTokens)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_DropERC1155 *DropERC1155Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_DropERC1155 *DropERC1155Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DropERC1155.Contract.Multicall(&_DropERC1155.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_DropERC1155 *DropERC1155TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DropERC1155.Contract.Multicall(&_DropERC1155.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.RenounceRole(&_DropERC1155.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.RenounceRole(&_DropERC1155.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.RevokeRole(&_DropERC1155.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DropERC1155 *DropERC1155TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.RevokeRole(&_DropERC1155.TransactOpts, role, account)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_DropERC1155 *DropERC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_DropERC1155 *DropERC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC1155.Contract.SafeBatchTransferFrom(&_DropERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC1155.Contract.SafeBatchTransferFrom(&_DropERC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_DropERC1155 *DropERC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_DropERC1155 *DropERC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC1155.Contract.SafeTransferFrom(&_DropERC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC1155.Contract.SafeTransferFrom(&_DropERC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DropERC1155 *DropERC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DropERC1155 *DropERC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetApprovalForAll(&_DropERC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetApprovalForAll(&_DropERC1155.TransactOpts, operator, approved)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0xab073c22.
//
// Solidity: function setClaimConditions(uint256 _tokenId, (uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address)[] _phases, bool _resetClaimEligibility) returns()
func (_DropERC1155 *DropERC1155Transactor) SetClaimConditions(opts *bind.TransactOpts, _tokenId *big.Int, _phases []IDropClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setClaimConditions", _tokenId, _phases, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0xab073c22.
//
// Solidity: function setClaimConditions(uint256 _tokenId, (uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address)[] _phases, bool _resetClaimEligibility) returns()
func (_DropERC1155 *DropERC1155Session) SetClaimConditions(_tokenId *big.Int, _phases []IDropClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetClaimConditions(&_DropERC1155.TransactOpts, _tokenId, _phases, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0xab073c22.
//
// Solidity: function setClaimConditions(uint256 _tokenId, (uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address)[] _phases, bool _resetClaimEligibility) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetClaimConditions(_tokenId *big.Int, _phases []IDropClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetClaimConditions(&_DropERC1155.TransactOpts, _tokenId, _phases, _resetClaimEligibility)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_DropERC1155 *DropERC1155Transactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_DropERC1155 *DropERC1155Session) SetContractURI(_uri string) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetContractURI(&_DropERC1155.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetContractURI(&_DropERC1155.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_DropERC1155 *DropERC1155Transactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_DropERC1155 *DropERC1155Session) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetDefaultRoyaltyInfo(&_DropERC1155.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetDefaultRoyaltyInfo(&_DropERC1155.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x87198cf2.
//
// Solidity: function setMaxTotalSupply(uint256 _tokenId, uint256 _maxTotalSupply) returns()
func (_DropERC1155 *DropERC1155Transactor) SetMaxTotalSupply(opts *bind.TransactOpts, _tokenId *big.Int, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setMaxTotalSupply", _tokenId, _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x87198cf2.
//
// Solidity: function setMaxTotalSupply(uint256 _tokenId, uint256 _maxTotalSupply) returns()
func (_DropERC1155 *DropERC1155Session) SetMaxTotalSupply(_tokenId *big.Int, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetMaxTotalSupply(&_DropERC1155.TransactOpts, _tokenId, _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x87198cf2.
//
// Solidity: function setMaxTotalSupply(uint256 _tokenId, uint256 _maxTotalSupply) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetMaxTotalSupply(_tokenId *big.Int, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetMaxTotalSupply(&_DropERC1155.TransactOpts, _tokenId, _maxTotalSupply)
}

// SetMaxWalletClaimCount is a paid mutator transaction binding the contract method 0x832c3a58.
//
// Solidity: function setMaxWalletClaimCount(uint256 _tokenId, uint256 _count) returns()
func (_DropERC1155 *DropERC1155Transactor) SetMaxWalletClaimCount(opts *bind.TransactOpts, _tokenId *big.Int, _count *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setMaxWalletClaimCount", _tokenId, _count)
}

// SetMaxWalletClaimCount is a paid mutator transaction binding the contract method 0x832c3a58.
//
// Solidity: function setMaxWalletClaimCount(uint256 _tokenId, uint256 _count) returns()
func (_DropERC1155 *DropERC1155Session) SetMaxWalletClaimCount(_tokenId *big.Int, _count *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetMaxWalletClaimCount(&_DropERC1155.TransactOpts, _tokenId, _count)
}

// SetMaxWalletClaimCount is a paid mutator transaction binding the contract method 0x832c3a58.
//
// Solidity: function setMaxWalletClaimCount(uint256 _tokenId, uint256 _count) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetMaxWalletClaimCount(_tokenId *big.Int, _count *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetMaxWalletClaimCount(&_DropERC1155.TransactOpts, _tokenId, _count)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_DropERC1155 *DropERC1155Transactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_DropERC1155 *DropERC1155Session) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetOwner(&_DropERC1155.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetOwner(&_DropERC1155.TransactOpts, _newOwner)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_DropERC1155 *DropERC1155Transactor) SetPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setPlatformFeeInfo", _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_DropERC1155 *DropERC1155Session) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetPlatformFeeInfo(&_DropERC1155.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetPlatformFeeInfo(&_DropERC1155.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_DropERC1155 *DropERC1155Transactor) SetPrimarySaleRecipient(opts *bind.TransactOpts, _saleRecipient common.Address) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setPrimarySaleRecipient", _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_DropERC1155 *DropERC1155Session) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetPrimarySaleRecipient(&_DropERC1155.TransactOpts, _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetPrimarySaleRecipient(&_DropERC1155.TransactOpts, _saleRecipient)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_DropERC1155 *DropERC1155Transactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_DropERC1155 *DropERC1155Session) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetRoyaltyInfoForToken(&_DropERC1155.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetRoyaltyInfoForToken(&_DropERC1155.TransactOpts, _tokenId, _recipient, _bps)
}

// SetWalletClaimCount is a paid mutator transaction binding the contract method 0xb1014400.
//
// Solidity: function setWalletClaimCount(uint256 _tokenId, address _claimer, uint256 _count) returns()
func (_DropERC1155 *DropERC1155Transactor) SetWalletClaimCount(opts *bind.TransactOpts, _tokenId *big.Int, _claimer common.Address, _count *big.Int) (*types.Transaction, error) {
	return _DropERC1155.contract.Transact(opts, "setWalletClaimCount", _tokenId, _claimer, _count)
}

// SetWalletClaimCount is a paid mutator transaction binding the contract method 0xb1014400.
//
// Solidity: function setWalletClaimCount(uint256 _tokenId, address _claimer, uint256 _count) returns()
func (_DropERC1155 *DropERC1155Session) SetWalletClaimCount(_tokenId *big.Int, _claimer common.Address, _count *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetWalletClaimCount(&_DropERC1155.TransactOpts, _tokenId, _claimer, _count)
}

// SetWalletClaimCount is a paid mutator transaction binding the contract method 0xb1014400.
//
// Solidity: function setWalletClaimCount(uint256 _tokenId, address _claimer, uint256 _count) returns()
func (_DropERC1155 *DropERC1155TransactorSession) SetWalletClaimCount(_tokenId *big.Int, _claimer common.Address, _count *big.Int) (*types.Transaction, error) {
	return _DropERC1155.Contract.SetWalletClaimCount(&_DropERC1155.TransactOpts, _tokenId, _claimer, _count)
}

// DropERC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DropERC1155 contract.
type DropERC1155ApprovalForAllIterator struct {
	Event *DropERC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DropERC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155ApprovalForAll)
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
		it.Event = new(DropERC1155ApprovalForAll)
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
func (it *DropERC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155ApprovalForAll represents a ApprovalForAll event raised by the DropERC1155 contract.
type DropERC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_DropERC1155 *DropERC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*DropERC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155ApprovalForAllIterator{contract: _DropERC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_DropERC1155 *DropERC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DropERC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155ApprovalForAll)
				if err := _DropERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseApprovalForAll(log types.Log) (*DropERC1155ApprovalForAll, error) {
	event := new(DropERC1155ApprovalForAll)
	if err := _DropERC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155ClaimConditionsUpdatedIterator is returned from FilterClaimConditionsUpdated and is used to iterate over the raw logs and unpacked data for ClaimConditionsUpdated events raised by the DropERC1155 contract.
type DropERC1155ClaimConditionsUpdatedIterator struct {
	Event *DropERC1155ClaimConditionsUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155ClaimConditionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155ClaimConditionsUpdated)
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
		it.Event = new(DropERC1155ClaimConditionsUpdated)
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
func (it *DropERC1155ClaimConditionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155ClaimConditionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155ClaimConditionsUpdated represents a ClaimConditionsUpdated event raised by the DropERC1155 contract.
type DropERC1155ClaimConditionsUpdated struct {
	TokenId         *big.Int
	ClaimConditions []IDropClaimConditionClaimCondition
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterClaimConditionsUpdated is a free log retrieval operation binding the contract event 0x7822655b74d50f461cbd7ca5dfc8b5e48b21fa2157bd7d277888fccce85af2ae.
//
// Solidity: event ClaimConditionsUpdated(uint256 indexed tokenId, (uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address)[] claimConditions)
func (_DropERC1155 *DropERC1155Filterer) FilterClaimConditionsUpdated(opts *bind.FilterOpts, tokenId []*big.Int) (*DropERC1155ClaimConditionsUpdatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "ClaimConditionsUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155ClaimConditionsUpdatedIterator{contract: _DropERC1155.contract, event: "ClaimConditionsUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimConditionsUpdated is a free log subscription operation binding the contract event 0x7822655b74d50f461cbd7ca5dfc8b5e48b21fa2157bd7d277888fccce85af2ae.
//
// Solidity: event ClaimConditionsUpdated(uint256 indexed tokenId, (uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address)[] claimConditions)
func (_DropERC1155 *DropERC1155Filterer) WatchClaimConditionsUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155ClaimConditionsUpdated, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "ClaimConditionsUpdated", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155ClaimConditionsUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
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

// ParseClaimConditionsUpdated is a log parse operation binding the contract event 0x7822655b74d50f461cbd7ca5dfc8b5e48b21fa2157bd7d277888fccce85af2ae.
//
// Solidity: event ClaimConditionsUpdated(uint256 indexed tokenId, (uint256,uint256,uint256,uint256,uint256,bytes32,uint256,address)[] claimConditions)
func (_DropERC1155 *DropERC1155Filterer) ParseClaimConditionsUpdated(log types.Log) (*DropERC1155ClaimConditionsUpdated, error) {
	event := new(DropERC1155ClaimConditionsUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155DefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the DropERC1155 contract.
type DropERC1155DefaultRoyaltyIterator struct {
	Event *DropERC1155DefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *DropERC1155DefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155DefaultRoyalty)
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
		it.Event = new(DropERC1155DefaultRoyalty)
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
func (it *DropERC1155DefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155DefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155DefaultRoyalty represents a DefaultRoyalty event raised by the DropERC1155 contract.
type DropERC1155DefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_DropERC1155 *DropERC1155Filterer) FilterDefaultRoyalty(opts *bind.FilterOpts) (*DropERC1155DefaultRoyaltyIterator, error) {

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "DefaultRoyalty")
	if err != nil {
		return nil, err
	}
	return &DropERC1155DefaultRoyaltyIterator{contract: _DropERC1155.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_DropERC1155 *DropERC1155Filterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *DropERC1155DefaultRoyalty) (event.Subscription, error) {

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "DefaultRoyalty")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155DefaultRoyalty)
				if err := _DropERC1155.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseDefaultRoyalty(log types.Log) (*DropERC1155DefaultRoyalty, error) {
	event := new(DropERC1155DefaultRoyalty)
	if err := _DropERC1155.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155MaxTotalSupplyUpdatedIterator is returned from FilterMaxTotalSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalSupplyUpdated events raised by the DropERC1155 contract.
type DropERC1155MaxTotalSupplyUpdatedIterator struct {
	Event *DropERC1155MaxTotalSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155MaxTotalSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155MaxTotalSupplyUpdated)
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
		it.Event = new(DropERC1155MaxTotalSupplyUpdated)
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
func (it *DropERC1155MaxTotalSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155MaxTotalSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155MaxTotalSupplyUpdated represents a MaxTotalSupplyUpdated event raised by the DropERC1155 contract.
type DropERC1155MaxTotalSupplyUpdated struct {
	TokenId        *big.Int
	MaxTotalSupply *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalSupplyUpdated is a free log retrieval operation binding the contract event 0xc58cd6132bb46df23d468939c03dd023b74b509aaa6b04c39d5a6461c65963bd.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 tokenId, uint256 maxTotalSupply)
func (_DropERC1155 *DropERC1155Filterer) FilterMaxTotalSupplyUpdated(opts *bind.FilterOpts) (*DropERC1155MaxTotalSupplyUpdatedIterator, error) {

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "MaxTotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC1155MaxTotalSupplyUpdatedIterator{contract: _DropERC1155.contract, event: "MaxTotalSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalSupplyUpdated is a free log subscription operation binding the contract event 0xc58cd6132bb46df23d468939c03dd023b74b509aaa6b04c39d5a6461c65963bd.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 tokenId, uint256 maxTotalSupply)
func (_DropERC1155 *DropERC1155Filterer) WatchMaxTotalSupplyUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155MaxTotalSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "MaxTotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155MaxTotalSupplyUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "MaxTotalSupplyUpdated", log); err != nil {
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

// ParseMaxTotalSupplyUpdated is a log parse operation binding the contract event 0xc58cd6132bb46df23d468939c03dd023b74b509aaa6b04c39d5a6461c65963bd.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 tokenId, uint256 maxTotalSupply)
func (_DropERC1155 *DropERC1155Filterer) ParseMaxTotalSupplyUpdated(log types.Log) (*DropERC1155MaxTotalSupplyUpdated, error) {
	event := new(DropERC1155MaxTotalSupplyUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "MaxTotalSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155MaxWalletClaimCountUpdatedIterator is returned from FilterMaxWalletClaimCountUpdated and is used to iterate over the raw logs and unpacked data for MaxWalletClaimCountUpdated events raised by the DropERC1155 contract.
type DropERC1155MaxWalletClaimCountUpdatedIterator struct {
	Event *DropERC1155MaxWalletClaimCountUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155MaxWalletClaimCountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155MaxWalletClaimCountUpdated)
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
		it.Event = new(DropERC1155MaxWalletClaimCountUpdated)
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
func (it *DropERC1155MaxWalletClaimCountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155MaxWalletClaimCountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155MaxWalletClaimCountUpdated represents a MaxWalletClaimCountUpdated event raised by the DropERC1155 contract.
type DropERC1155MaxWalletClaimCountUpdated struct {
	TokenId *big.Int
	Count   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMaxWalletClaimCountUpdated is a free log retrieval operation binding the contract event 0x07fa2d0eb2fe8b8e6fbee6073cf9d84659d6db054d221579a0373ae29bc9d73d.
//
// Solidity: event MaxWalletClaimCountUpdated(uint256 tokenId, uint256 count)
func (_DropERC1155 *DropERC1155Filterer) FilterMaxWalletClaimCountUpdated(opts *bind.FilterOpts) (*DropERC1155MaxWalletClaimCountUpdatedIterator, error) {

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "MaxWalletClaimCountUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC1155MaxWalletClaimCountUpdatedIterator{contract: _DropERC1155.contract, event: "MaxWalletClaimCountUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxWalletClaimCountUpdated is a free log subscription operation binding the contract event 0x07fa2d0eb2fe8b8e6fbee6073cf9d84659d6db054d221579a0373ae29bc9d73d.
//
// Solidity: event MaxWalletClaimCountUpdated(uint256 tokenId, uint256 count)
func (_DropERC1155 *DropERC1155Filterer) WatchMaxWalletClaimCountUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155MaxWalletClaimCountUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "MaxWalletClaimCountUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155MaxWalletClaimCountUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "MaxWalletClaimCountUpdated", log); err != nil {
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

// ParseMaxWalletClaimCountUpdated is a log parse operation binding the contract event 0x07fa2d0eb2fe8b8e6fbee6073cf9d84659d6db054d221579a0373ae29bc9d73d.
//
// Solidity: event MaxWalletClaimCountUpdated(uint256 tokenId, uint256 count)
func (_DropERC1155 *DropERC1155Filterer) ParseMaxWalletClaimCountUpdated(log types.Log) (*DropERC1155MaxWalletClaimCountUpdated, error) {
	event := new(DropERC1155MaxWalletClaimCountUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "MaxWalletClaimCountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155OwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the DropERC1155 contract.
type DropERC1155OwnerUpdatedIterator struct {
	Event *DropERC1155OwnerUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155OwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155OwnerUpdated)
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
		it.Event = new(DropERC1155OwnerUpdated)
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
func (it *DropERC1155OwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155OwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155OwnerUpdated represents a OwnerUpdated event raised by the DropERC1155 contract.
type DropERC1155OwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address prevOwner, address newOwner)
func (_DropERC1155 *DropERC1155Filterer) FilterOwnerUpdated(opts *bind.FilterOpts) (*DropERC1155OwnerUpdatedIterator, error) {

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC1155OwnerUpdatedIterator{contract: _DropERC1155.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address prevOwner, address newOwner)
func (_DropERC1155 *DropERC1155Filterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155OwnerUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "OwnerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155OwnerUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseOwnerUpdated(log types.Log) (*DropERC1155OwnerUpdated, error) {
	event := new(DropERC1155OwnerUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155PlatformFeeInfoUpdatedIterator is returned from FilterPlatformFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeInfoUpdated events raised by the DropERC1155 contract.
type DropERC1155PlatformFeeInfoUpdatedIterator struct {
	Event *DropERC1155PlatformFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155PlatformFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155PlatformFeeInfoUpdated)
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
		it.Event = new(DropERC1155PlatformFeeInfoUpdated)
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
func (it *DropERC1155PlatformFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155PlatformFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155PlatformFeeInfoUpdated represents a PlatformFeeInfoUpdated event raised by the DropERC1155 contract.
type DropERC1155PlatformFeeInfoUpdated struct {
	PlatformFeeRecipient common.Address
	PlatformFeeBps       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeInfoUpdated is a free log retrieval operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address platformFeeRecipient, uint256 platformFeeBps)
func (_DropERC1155 *DropERC1155Filterer) FilterPlatformFeeInfoUpdated(opts *bind.FilterOpts) (*DropERC1155PlatformFeeInfoUpdatedIterator, error) {

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "PlatformFeeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC1155PlatformFeeInfoUpdatedIterator{contract: _DropERC1155.contract, event: "PlatformFeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeInfoUpdated is a free log subscription operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address platformFeeRecipient, uint256 platformFeeBps)
func (_DropERC1155 *DropERC1155Filterer) WatchPlatformFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155PlatformFeeInfoUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "PlatformFeeInfoUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155PlatformFeeInfoUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
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

// ParsePlatformFeeInfoUpdated is a log parse operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address platformFeeRecipient, uint256 platformFeeBps)
func (_DropERC1155 *DropERC1155Filterer) ParsePlatformFeeInfoUpdated(log types.Log) (*DropERC1155PlatformFeeInfoUpdated, error) {
	event := new(DropERC1155PlatformFeeInfoUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155PrimarySaleRecipientUpdatedIterator is returned from FilterPrimarySaleRecipientUpdated and is used to iterate over the raw logs and unpacked data for PrimarySaleRecipientUpdated events raised by the DropERC1155 contract.
type DropERC1155PrimarySaleRecipientUpdatedIterator struct {
	Event *DropERC1155PrimarySaleRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155PrimarySaleRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155PrimarySaleRecipientUpdated)
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
		it.Event = new(DropERC1155PrimarySaleRecipientUpdated)
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
func (it *DropERC1155PrimarySaleRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155PrimarySaleRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155PrimarySaleRecipientUpdated represents a PrimarySaleRecipientUpdated event raised by the DropERC1155 contract.
type DropERC1155PrimarySaleRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimarySaleRecipientUpdated is a free log retrieval operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_DropERC1155 *DropERC1155Filterer) FilterPrimarySaleRecipientUpdated(opts *bind.FilterOpts, recipient []common.Address) (*DropERC1155PrimarySaleRecipientUpdatedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155PrimarySaleRecipientUpdatedIterator{contract: _DropERC1155.contract, event: "PrimarySaleRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchPrimarySaleRecipientUpdated is a free log subscription operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_DropERC1155 *DropERC1155Filterer) WatchPrimarySaleRecipientUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155PrimarySaleRecipientUpdated, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155PrimarySaleRecipientUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
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

// ParsePrimarySaleRecipientUpdated is a log parse operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_DropERC1155 *DropERC1155Filterer) ParsePrimarySaleRecipientUpdated(log types.Log) (*DropERC1155PrimarySaleRecipientUpdated, error) {
	event := new(DropERC1155PrimarySaleRecipientUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DropERC1155 contract.
type DropERC1155RoleAdminChangedIterator struct {
	Event *DropERC1155RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DropERC1155RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155RoleAdminChanged)
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
		it.Event = new(DropERC1155RoleAdminChanged)
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
func (it *DropERC1155RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155RoleAdminChanged represents a RoleAdminChanged event raised by the DropERC1155 contract.
type DropERC1155RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DropERC1155 *DropERC1155Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DropERC1155RoleAdminChangedIterator, error) {

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

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155RoleAdminChangedIterator{contract: _DropERC1155.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DropERC1155 *DropERC1155Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DropERC1155RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155RoleAdminChanged)
				if err := _DropERC1155.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseRoleAdminChanged(log types.Log) (*DropERC1155RoleAdminChanged, error) {
	event := new(DropERC1155RoleAdminChanged)
	if err := _DropERC1155.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DropERC1155 contract.
type DropERC1155RoleGrantedIterator struct {
	Event *DropERC1155RoleGranted // Event containing the contract specifics and raw log

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
func (it *DropERC1155RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155RoleGranted)
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
		it.Event = new(DropERC1155RoleGranted)
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
func (it *DropERC1155RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155RoleGranted represents a RoleGranted event raised by the DropERC1155 contract.
type DropERC1155RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC1155 *DropERC1155Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DropERC1155RoleGrantedIterator, error) {

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

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155RoleGrantedIterator{contract: _DropERC1155.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC1155 *DropERC1155Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DropERC1155RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155RoleGranted)
				if err := _DropERC1155.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseRoleGranted(log types.Log) (*DropERC1155RoleGranted, error) {
	event := new(DropERC1155RoleGranted)
	if err := _DropERC1155.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DropERC1155 contract.
type DropERC1155RoleRevokedIterator struct {
	Event *DropERC1155RoleRevoked // Event containing the contract specifics and raw log

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
func (it *DropERC1155RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155RoleRevoked)
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
		it.Event = new(DropERC1155RoleRevoked)
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
func (it *DropERC1155RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155RoleRevoked represents a RoleRevoked event raised by the DropERC1155 contract.
type DropERC1155RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC1155 *DropERC1155Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DropERC1155RoleRevokedIterator, error) {

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

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155RoleRevokedIterator{contract: _DropERC1155.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC1155 *DropERC1155Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DropERC1155RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155RoleRevoked)
				if err := _DropERC1155.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseRoleRevoked(log types.Log) (*DropERC1155RoleRevoked, error) {
	event := new(DropERC1155RoleRevoked)
	if err := _DropERC1155.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155RoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the DropERC1155 contract.
type DropERC1155RoyaltyForTokenIterator struct {
	Event *DropERC1155RoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *DropERC1155RoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155RoyaltyForToken)
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
		it.Event = new(DropERC1155RoyaltyForToken)
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
func (it *DropERC1155RoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155RoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155RoyaltyForToken represents a RoyaltyForToken event raised by the DropERC1155 contract.
type DropERC1155RoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address royaltyRecipient, uint256 royaltyBps)
func (_DropERC1155 *DropERC1155Filterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int) (*DropERC1155RoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155RoyaltyForTokenIterator{contract: _DropERC1155.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address royaltyRecipient, uint256 royaltyBps)
func (_DropERC1155 *DropERC1155Filterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *DropERC1155RoyaltyForToken, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155RoyaltyForToken)
				if err := _DropERC1155.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseRoyaltyForToken(log types.Log) (*DropERC1155RoyaltyForToken, error) {
	event := new(DropERC1155RoyaltyForToken)
	if err := _DropERC1155.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155TokensClaimedIterator is returned from FilterTokensClaimed and is used to iterate over the raw logs and unpacked data for TokensClaimed events raised by the DropERC1155 contract.
type DropERC1155TokensClaimedIterator struct {
	Event *DropERC1155TokensClaimed // Event containing the contract specifics and raw log

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
func (it *DropERC1155TokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155TokensClaimed)
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
		it.Event = new(DropERC1155TokensClaimed)
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
func (it *DropERC1155TokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155TokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155TokensClaimed represents a TokensClaimed event raised by the DropERC1155 contract.
type DropERC1155TokensClaimed struct {
	ClaimConditionIndex *big.Int
	TokenId             *big.Int
	Claimer             common.Address
	Receiver            common.Address
	QuantityClaimed     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensClaimed is a free log retrieval operation binding the contract event 0x4f72e6585331094d368e469f11198272039d08cbddfcda1577e192687a83afb6.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, uint256 indexed tokenId, address indexed claimer, address receiver, uint256 quantityClaimed)
func (_DropERC1155 *DropERC1155Filterer) FilterTokensClaimed(opts *bind.FilterOpts, claimConditionIndex []*big.Int, tokenId []*big.Int, claimer []common.Address) (*DropERC1155TokensClaimedIterator, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "TokensClaimed", claimConditionIndexRule, tokenIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155TokensClaimedIterator{contract: _DropERC1155.contract, event: "TokensClaimed", logs: logs, sub: sub}, nil
}

// WatchTokensClaimed is a free log subscription operation binding the contract event 0x4f72e6585331094d368e469f11198272039d08cbddfcda1577e192687a83afb6.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, uint256 indexed tokenId, address indexed claimer, address receiver, uint256 quantityClaimed)
func (_DropERC1155 *DropERC1155Filterer) WatchTokensClaimed(opts *bind.WatchOpts, sink chan<- *DropERC1155TokensClaimed, claimConditionIndex []*big.Int, tokenId []*big.Int, claimer []common.Address) (event.Subscription, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "TokensClaimed", claimConditionIndexRule, tokenIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155TokensClaimed)
				if err := _DropERC1155.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
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

// ParseTokensClaimed is a log parse operation binding the contract event 0x4f72e6585331094d368e469f11198272039d08cbddfcda1577e192687a83afb6.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, uint256 indexed tokenId, address indexed claimer, address receiver, uint256 quantityClaimed)
func (_DropERC1155 *DropERC1155Filterer) ParseTokensClaimed(log types.Log) (*DropERC1155TokensClaimed, error) {
	event := new(DropERC1155TokensClaimed)
	if err := _DropERC1155.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155TokensLazyMintedIterator is returned from FilterTokensLazyMinted and is used to iterate over the raw logs and unpacked data for TokensLazyMinted events raised by the DropERC1155 contract.
type DropERC1155TokensLazyMintedIterator struct {
	Event *DropERC1155TokensLazyMinted // Event containing the contract specifics and raw log

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
func (it *DropERC1155TokensLazyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155TokensLazyMinted)
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
		it.Event = new(DropERC1155TokensLazyMinted)
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
func (it *DropERC1155TokensLazyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155TokensLazyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155TokensLazyMinted represents a TokensLazyMinted event raised by the DropERC1155 contract.
type DropERC1155TokensLazyMinted struct {
	StartTokenId *big.Int
	EndTokenId   *big.Int
	BaseURI      string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensLazyMinted is a free log retrieval operation binding the contract event 0x4e6c698792b8dfb7c94c60c7e9e91f82932832d5e1ec0870ed42cf674e6af445.
//
// Solidity: event TokensLazyMinted(uint256 startTokenId, uint256 endTokenId, string baseURI)
func (_DropERC1155 *DropERC1155Filterer) FilterTokensLazyMinted(opts *bind.FilterOpts) (*DropERC1155TokensLazyMintedIterator, error) {

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "TokensLazyMinted")
	if err != nil {
		return nil, err
	}
	return &DropERC1155TokensLazyMintedIterator{contract: _DropERC1155.contract, event: "TokensLazyMinted", logs: logs, sub: sub}, nil
}

// WatchTokensLazyMinted is a free log subscription operation binding the contract event 0x4e6c698792b8dfb7c94c60c7e9e91f82932832d5e1ec0870ed42cf674e6af445.
//
// Solidity: event TokensLazyMinted(uint256 startTokenId, uint256 endTokenId, string baseURI)
func (_DropERC1155 *DropERC1155Filterer) WatchTokensLazyMinted(opts *bind.WatchOpts, sink chan<- *DropERC1155TokensLazyMinted) (event.Subscription, error) {

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "TokensLazyMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155TokensLazyMinted)
				if err := _DropERC1155.contract.UnpackLog(event, "TokensLazyMinted", log); err != nil {
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

// ParseTokensLazyMinted is a log parse operation binding the contract event 0x4e6c698792b8dfb7c94c60c7e9e91f82932832d5e1ec0870ed42cf674e6af445.
//
// Solidity: event TokensLazyMinted(uint256 startTokenId, uint256 endTokenId, string baseURI)
func (_DropERC1155 *DropERC1155Filterer) ParseTokensLazyMinted(log types.Log) (*DropERC1155TokensLazyMinted, error) {
	event := new(DropERC1155TokensLazyMinted)
	if err := _DropERC1155.contract.UnpackLog(event, "TokensLazyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the DropERC1155 contract.
type DropERC1155TransferBatchIterator struct {
	Event *DropERC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *DropERC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155TransferBatch)
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
		it.Event = new(DropERC1155TransferBatch)
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
func (it *DropERC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155TransferBatch represents a TransferBatch event raised by the DropERC1155 contract.
type DropERC1155TransferBatch struct {
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
func (_DropERC1155 *DropERC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DropERC1155TransferBatchIterator, error) {

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

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155TransferBatchIterator{contract: _DropERC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_DropERC1155 *DropERC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *DropERC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155TransferBatch)
				if err := _DropERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseTransferBatch(log types.Log) (*DropERC1155TransferBatch, error) {
	event := new(DropERC1155TransferBatch)
	if err := _DropERC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the DropERC1155 contract.
type DropERC1155TransferSingleIterator struct {
	Event *DropERC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *DropERC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155TransferSingle)
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
		it.Event = new(DropERC1155TransferSingle)
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
func (it *DropERC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155TransferSingle represents a TransferSingle event raised by the DropERC1155 contract.
type DropERC1155TransferSingle struct {
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
func (_DropERC1155 *DropERC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*DropERC1155TransferSingleIterator, error) {

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

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155TransferSingleIterator{contract: _DropERC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_DropERC1155 *DropERC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *DropERC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155TransferSingle)
				if err := _DropERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseTransferSingle(log types.Log) (*DropERC1155TransferSingle, error) {
	event := new(DropERC1155TransferSingle)
	if err := _DropERC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the DropERC1155 contract.
type DropERC1155URIIterator struct {
	Event *DropERC1155URI // Event containing the contract specifics and raw log

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
func (it *DropERC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155URI)
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
		it.Event = new(DropERC1155URI)
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
func (it *DropERC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155URI represents a URI event raised by the DropERC1155 contract.
type DropERC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_DropERC1155 *DropERC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*DropERC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155URIIterator{contract: _DropERC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_DropERC1155 *DropERC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *DropERC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155URI)
				if err := _DropERC1155.contract.UnpackLog(event, "URI", log); err != nil {
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
func (_DropERC1155 *DropERC1155Filterer) ParseURI(log types.Log) (*DropERC1155URI, error) {
	event := new(DropERC1155URI)
	if err := _DropERC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC1155WalletClaimCountUpdatedIterator is returned from FilterWalletClaimCountUpdated and is used to iterate over the raw logs and unpacked data for WalletClaimCountUpdated events raised by the DropERC1155 contract.
type DropERC1155WalletClaimCountUpdatedIterator struct {
	Event *DropERC1155WalletClaimCountUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC1155WalletClaimCountUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC1155WalletClaimCountUpdated)
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
		it.Event = new(DropERC1155WalletClaimCountUpdated)
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
func (it *DropERC1155WalletClaimCountUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC1155WalletClaimCountUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC1155WalletClaimCountUpdated represents a WalletClaimCountUpdated event raised by the DropERC1155 contract.
type DropERC1155WalletClaimCountUpdated struct {
	TokenId *big.Int
	Wallet  common.Address
	Count   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWalletClaimCountUpdated is a free log retrieval operation binding the contract event 0x9260cdce30c9abdb65593c1a903e40c87feb886a28aa4335a6695547988aab2f.
//
// Solidity: event WalletClaimCountUpdated(uint256 tokenId, address indexed wallet, uint256 count)
func (_DropERC1155 *DropERC1155Filterer) FilterWalletClaimCountUpdated(opts *bind.FilterOpts, wallet []common.Address) (*DropERC1155WalletClaimCountUpdatedIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _DropERC1155.contract.FilterLogs(opts, "WalletClaimCountUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return &DropERC1155WalletClaimCountUpdatedIterator{contract: _DropERC1155.contract, event: "WalletClaimCountUpdated", logs: logs, sub: sub}, nil
}

// WatchWalletClaimCountUpdated is a free log subscription operation binding the contract event 0x9260cdce30c9abdb65593c1a903e40c87feb886a28aa4335a6695547988aab2f.
//
// Solidity: event WalletClaimCountUpdated(uint256 tokenId, address indexed wallet, uint256 count)
func (_DropERC1155 *DropERC1155Filterer) WatchWalletClaimCountUpdated(opts *bind.WatchOpts, sink chan<- *DropERC1155WalletClaimCountUpdated, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _DropERC1155.contract.WatchLogs(opts, "WalletClaimCountUpdated", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC1155WalletClaimCountUpdated)
				if err := _DropERC1155.contract.UnpackLog(event, "WalletClaimCountUpdated", log); err != nil {
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

// ParseWalletClaimCountUpdated is a log parse operation binding the contract event 0x9260cdce30c9abdb65593c1a903e40c87feb886a28aa4335a6695547988aab2f.
//
// Solidity: event WalletClaimCountUpdated(uint256 tokenId, address indexed wallet, uint256 count)
func (_DropERC1155 *DropERC1155Filterer) ParseWalletClaimCountUpdated(log types.Log) (*DropERC1155WalletClaimCountUpdated, error) {
	event := new(DropERC1155WalletClaimCountUpdated)
	if err := _DropERC1155.contract.UnpackLog(event, "WalletClaimCountUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
