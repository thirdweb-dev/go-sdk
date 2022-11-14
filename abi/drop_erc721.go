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

// IClaimConditionClaimCondition is an auto generated low-level Go binding around an user-defined struct.
type IClaimConditionClaimCondition struct {
	StartTimestamp         *big.Int
	MaxClaimableSupply     *big.Int
	SupplyClaimed          *big.Int
	QuantityLimitPerWallet *big.Int
	MerkleRoot             [32]byte
	PricePerToken          *big.Int
	Currency               common.Address
	Metadata               string
}

// IDropAllowlistProof is an auto generated low-level Go binding around an user-defined struct.
type IDropAllowlistProof struct {
	Proof                  [][32]byte
	QuantityLimitPerWallet *big.Int
	PricePerToken          *big.Int
	Currency               common.Address
}

// DropERC721MetaData contains all meta data concerning the DropERC721 contract.
var DropERC721MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ApprovalCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApprovalToCurrentOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ApproveToCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceQueryForZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MintZeroQuantity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerQueryForNonexistentToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferCallerNotOwnerNorApproved\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferFromIncorrectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToNonERC721ReceiverImplementer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransferToZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"URIQueryForNonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIClaimCondition.ClaimCondition[]\",\"name\":\"claimConditions\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"resetEligibility\",\"type\":\"bool\"}],\"name\":\"ClaimConditionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"prevURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"ContractURIUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRoyaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRoyaltyBps\",\"type\":\"uint256\"}],\"name\":\"DefaultRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"MaxTotalSupplyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"platformFeeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"platformFeeBps\",\"type\":\"uint256\"}],\"name\":\"PlatformFeeInfoUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"PrimarySaleRecipientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"royaltyRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyBps\",\"type\":\"uint256\"}],\"name\":\"RoyaltyForToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"revealedURI\",\"type\":\"string\"}],\"name\":\"TokenURIRevealed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"claimConditionIndex\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantityClaimed\",\"type\":\"uint256\"}],\"name\":\"TokensClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedBaseURI\",\"type\":\"bytes\"}],\"name\":\"TokensLazyMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIDrop.AllowlistProof\",\"name\":\"_allowlistProof\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimCondition\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentStartId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractVersion\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"key\",\"type\":\"bytes\"}],\"name\":\"encryptDecrypt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"encryptedData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveClaimConditionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseURICount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getBatchIdAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"}],\"name\":\"getClaimConditionById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structIClaimCondition.ClaimCondition\",\"name\":\"condition\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultRoyaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPlatformFeeInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_key\",\"type\":\"bytes\"}],\"name\":\"getRevealURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"revealedURI\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"member\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getRoyaltyInfoForToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"}],\"name\":\"getSupplyClaimedByWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyClaimedByWallet\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRoleWithSwitch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_contractURI\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_trustedForwarders\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_saleRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_royaltyBps\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_platformFeeBps\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_batchId\",\"type\":\"uint256\"}],\"name\":\"isEncryptedBatch\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"isTrustedForwarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_baseURIForTokens\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"lazyMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenIdToClaim\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenIdToMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"primarySaleRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_key\",\"type\":\"bytes\"}],\"name\":\"reveal\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"revealedURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salePrice\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxClaimableSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyClaimed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"internalType\":\"structIClaimCondition.ClaimCondition[]\",\"name\":\"_conditions\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"_resetClaimEligibility\",\"type\":\"bool\"}],\"name\":\"setClaimConditions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"setContractURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_royaltyRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_royaltyBps\",\"type\":\"uint256\"}],\"name\":\"setDefaultRoyaltyInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxTotalSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_platformFeeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_platformFeeBps\",\"type\":\"uint256\"}],\"name\":\"setPlatformFeeInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_saleRecipient\",\"type\":\"address\"}],\"name\":\"setPrimarySaleRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_bps\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyInfoForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalMinted\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_conditionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pricePerToken\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"quantityLimitPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerToken\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"}],\"internalType\":\"structIDrop.AllowlistProof\",\"name\":\"_allowlistProof\",\"type\":\"tuple\"}],\"name\":\"verifyClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOverride\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DropERC721ABI is the input ABI used to generate the binding from.
// Deprecated: Use DropERC721MetaData.ABI instead.
var DropERC721ABI = DropERC721MetaData.ABI

// DropERC721 is an auto generated Go binding around an Ethereum contract.
type DropERC721 struct {
	DropERC721Caller     // Read-only binding to the contract
	DropERC721Transactor // Write-only binding to the contract
	DropERC721Filterer   // Log filterer for contract events
}

// DropERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type DropERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DropERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DropERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DropERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DropERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DropERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DropERC721Session struct {
	Contract     *DropERC721       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DropERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DropERC721CallerSession struct {
	Contract *DropERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DropERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DropERC721TransactorSession struct {
	Contract     *DropERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DropERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type DropERC721Raw struct {
	Contract *DropERC721 // Generic contract binding to access the raw methods on
}

// DropERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DropERC721CallerRaw struct {
	Contract *DropERC721Caller // Generic read-only contract binding to access the raw methods on
}

// DropERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DropERC721TransactorRaw struct {
	Contract *DropERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDropERC721 creates a new instance of DropERC721, bound to a specific deployed contract.
func NewDropERC721(address common.Address, backend bind.ContractBackend) (*DropERC721, error) {
	contract, err := bindDropERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DropERC721{DropERC721Caller: DropERC721Caller{contract: contract}, DropERC721Transactor: DropERC721Transactor{contract: contract}, DropERC721Filterer: DropERC721Filterer{contract: contract}}, nil
}

// NewDropERC721Caller creates a new read-only instance of DropERC721, bound to a specific deployed contract.
func NewDropERC721Caller(address common.Address, caller bind.ContractCaller) (*DropERC721Caller, error) {
	contract, err := bindDropERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DropERC721Caller{contract: contract}, nil
}

// NewDropERC721Transactor creates a new write-only instance of DropERC721, bound to a specific deployed contract.
func NewDropERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*DropERC721Transactor, error) {
	contract, err := bindDropERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DropERC721Transactor{contract: contract}, nil
}

// NewDropERC721Filterer creates a new log filterer instance of DropERC721, bound to a specific deployed contract.
func NewDropERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*DropERC721Filterer, error) {
	contract, err := bindDropERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DropERC721Filterer{contract: contract}, nil
}

// bindDropERC721 binds a generic wrapper to an already deployed contract.
func bindDropERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DropERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DropERC721 *DropERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DropERC721.Contract.DropERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DropERC721 *DropERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DropERC721.Contract.DropERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DropERC721 *DropERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DropERC721.Contract.DropERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DropERC721 *DropERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DropERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DropERC721 *DropERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DropERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DropERC721 *DropERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DropERC721.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DropERC721 *DropERC721Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DropERC721 *DropERC721Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _DropERC721.Contract.DEFAULTADMINROLE(&_DropERC721.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_DropERC721 *DropERC721CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _DropERC721.Contract.DEFAULTADMINROLE(&_DropERC721.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DropERC721 *DropERC721Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DropERC721 *DropERC721Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DropERC721.Contract.BalanceOf(&_DropERC721.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DropERC721.Contract.BalanceOf(&_DropERC721.CallOpts, owner)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_DropERC721 *DropERC721Caller) ClaimCondition(opts *bind.CallOpts) (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "claimCondition")

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

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_DropERC721 *DropERC721Session) ClaimCondition() (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _DropERC721.Contract.ClaimCondition(&_DropERC721.CallOpts)
}

// ClaimCondition is a free data retrieval call binding the contract method 0xd637ed59.
//
// Solidity: function claimCondition() view returns(uint256 currentStartId, uint256 count)
func (_DropERC721 *DropERC721CallerSession) ClaimCondition() (struct {
	CurrentStartId *big.Int
	Count          *big.Int
}, error) {
	return _DropERC721.Contract.ClaimCondition(&_DropERC721.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_DropERC721 *DropERC721Caller) ContractType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "contractType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_DropERC721 *DropERC721Session) ContractType() ([32]byte, error) {
	return _DropERC721.Contract.ContractType(&_DropERC721.CallOpts)
}

// ContractType is a free data retrieval call binding the contract method 0xcb2ef6f7.
//
// Solidity: function contractType() pure returns(bytes32)
func (_DropERC721 *DropERC721CallerSession) ContractType() ([32]byte, error) {
	return _DropERC721.Contract.ContractType(&_DropERC721.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_DropERC721 *DropERC721Caller) InternalContractURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "contractURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_DropERC721 *DropERC721Session) InternalContractURI() (string, error) {
	return _DropERC721.Contract.InternalContractURI(&_DropERC721.CallOpts)
}

// InternalContractURI is a free data retrieval call binding the contract method 0xe8a3d485.
//
// Solidity: function contractURI() view returns(string)
func (_DropERC721 *DropERC721CallerSession) InternalContractURI() (string, error) {
	return _DropERC721.Contract.InternalContractURI(&_DropERC721.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_DropERC721 *DropERC721Caller) ContractVersion(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "contractVersion")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_DropERC721 *DropERC721Session) ContractVersion() (uint8, error) {
	return _DropERC721.Contract.ContractVersion(&_DropERC721.CallOpts)
}

// ContractVersion is a free data retrieval call binding the contract method 0xa0a8e460.
//
// Solidity: function contractVersion() pure returns(uint8)
func (_DropERC721 *DropERC721CallerSession) ContractVersion() (uint8, error) {
	return _DropERC721.Contract.ContractVersion(&_DropERC721.CallOpts)
}

// EncryptDecrypt is a free data retrieval call binding the contract method 0xe7150322.
//
// Solidity: function encryptDecrypt(bytes data, bytes key) pure returns(bytes result)
func (_DropERC721 *DropERC721Caller) EncryptDecrypt(opts *bind.CallOpts, data []byte, key []byte) ([]byte, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "encryptDecrypt", data, key)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncryptDecrypt is a free data retrieval call binding the contract method 0xe7150322.
//
// Solidity: function encryptDecrypt(bytes data, bytes key) pure returns(bytes result)
func (_DropERC721 *DropERC721Session) EncryptDecrypt(data []byte, key []byte) ([]byte, error) {
	return _DropERC721.Contract.EncryptDecrypt(&_DropERC721.CallOpts, data, key)
}

// EncryptDecrypt is a free data retrieval call binding the contract method 0xe7150322.
//
// Solidity: function encryptDecrypt(bytes data, bytes key) pure returns(bytes result)
func (_DropERC721 *DropERC721CallerSession) EncryptDecrypt(data []byte, key []byte) ([]byte, error) {
	return _DropERC721.Contract.EncryptDecrypt(&_DropERC721.CallOpts, data, key)
}

// EncryptedData is a free data retrieval call binding the contract method 0xa05112fc.
//
// Solidity: function encryptedData(uint256 ) view returns(bytes)
func (_DropERC721 *DropERC721Caller) EncryptedData(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "encryptedData", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncryptedData is a free data retrieval call binding the contract method 0xa05112fc.
//
// Solidity: function encryptedData(uint256 ) view returns(bytes)
func (_DropERC721 *DropERC721Session) EncryptedData(arg0 *big.Int) ([]byte, error) {
	return _DropERC721.Contract.EncryptedData(&_DropERC721.CallOpts, arg0)
}

// EncryptedData is a free data retrieval call binding the contract method 0xa05112fc.
//
// Solidity: function encryptedData(uint256 ) view returns(bytes)
func (_DropERC721 *DropERC721CallerSession) EncryptedData(arg0 *big.Int) ([]byte, error) {
	return _DropERC721.Contract.EncryptedData(&_DropERC721.CallOpts, arg0)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_DropERC721 *DropERC721Caller) GetActiveClaimConditionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getActiveClaimConditionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_DropERC721 *DropERC721Session) GetActiveClaimConditionId() (*big.Int, error) {
	return _DropERC721.Contract.GetActiveClaimConditionId(&_DropERC721.CallOpts)
}

// GetActiveClaimConditionId is a free data retrieval call binding the contract method 0xc68907de.
//
// Solidity: function getActiveClaimConditionId() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) GetActiveClaimConditionId() (*big.Int, error) {
	return _DropERC721.Contract.GetActiveClaimConditionId(&_DropERC721.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DropERC721 *DropERC721Caller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DropERC721 *DropERC721Session) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DropERC721.Contract.GetApproved(&_DropERC721.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DropERC721 *DropERC721CallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DropERC721.Contract.GetApproved(&_DropERC721.CallOpts, tokenId)
}

// GetBaseURICount is a free data retrieval call binding the contract method 0x63b45e2d.
//
// Solidity: function getBaseURICount() view returns(uint256)
func (_DropERC721 *DropERC721Caller) GetBaseURICount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getBaseURICount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseURICount is a free data retrieval call binding the contract method 0x63b45e2d.
//
// Solidity: function getBaseURICount() view returns(uint256)
func (_DropERC721 *DropERC721Session) GetBaseURICount() (*big.Int, error) {
	return _DropERC721.Contract.GetBaseURICount(&_DropERC721.CallOpts)
}

// GetBaseURICount is a free data retrieval call binding the contract method 0x63b45e2d.
//
// Solidity: function getBaseURICount() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) GetBaseURICount() (*big.Int, error) {
	return _DropERC721.Contract.GetBaseURICount(&_DropERC721.CallOpts)
}

// GetBatchIdAtIndex is a free data retrieval call binding the contract method 0x2419f51b.
//
// Solidity: function getBatchIdAtIndex(uint256 _index) view returns(uint256)
func (_DropERC721 *DropERC721Caller) GetBatchIdAtIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getBatchIdAtIndex", _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBatchIdAtIndex is a free data retrieval call binding the contract method 0x2419f51b.
//
// Solidity: function getBatchIdAtIndex(uint256 _index) view returns(uint256)
func (_DropERC721 *DropERC721Session) GetBatchIdAtIndex(_index *big.Int) (*big.Int, error) {
	return _DropERC721.Contract.GetBatchIdAtIndex(&_DropERC721.CallOpts, _index)
}

// GetBatchIdAtIndex is a free data retrieval call binding the contract method 0x2419f51b.
//
// Solidity: function getBatchIdAtIndex(uint256 _index) view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) GetBatchIdAtIndex(_index *big.Int) (*big.Int, error) {
	return _DropERC721.Contract.GetBatchIdAtIndex(&_DropERC721.CallOpts, _index)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_DropERC721 *DropERC721Caller) GetClaimConditionById(opts *bind.CallOpts, _conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getClaimConditionById", _conditionId)

	if err != nil {
		return *new(IClaimConditionClaimCondition), err
	}

	out0 := *abi.ConvertType(out[0], new(IClaimConditionClaimCondition)).(*IClaimConditionClaimCondition)

	return out0, err

}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_DropERC721 *DropERC721Session) GetClaimConditionById(_conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	return _DropERC721.Contract.GetClaimConditionById(&_DropERC721.CallOpts, _conditionId)
}

// GetClaimConditionById is a free data retrieval call binding the contract method 0x6f8934f4.
//
// Solidity: function getClaimConditionById(uint256 _conditionId) view returns((uint256,uint256,uint256,uint256,bytes32,uint256,address,string) condition)
func (_DropERC721 *DropERC721CallerSession) GetClaimConditionById(_conditionId *big.Int) (IClaimConditionClaimCondition, error) {
	return _DropERC721.Contract.GetClaimConditionById(&_DropERC721.CallOpts, _conditionId)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_DropERC721 *DropERC721Caller) GetDefaultRoyaltyInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getDefaultRoyaltyInfo")

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
func (_DropERC721 *DropERC721Session) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _DropERC721.Contract.GetDefaultRoyaltyInfo(&_DropERC721.CallOpts)
}

// GetDefaultRoyaltyInfo is a free data retrieval call binding the contract method 0xb24f2d39.
//
// Solidity: function getDefaultRoyaltyInfo() view returns(address, uint16)
func (_DropERC721 *DropERC721CallerSession) GetDefaultRoyaltyInfo() (common.Address, uint16, error) {
	return _DropERC721.Contract.GetDefaultRoyaltyInfo(&_DropERC721.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_DropERC721 *DropERC721Caller) GetPlatformFeeInfo(opts *bind.CallOpts) (common.Address, uint16, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getPlatformFeeInfo")

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
func (_DropERC721 *DropERC721Session) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _DropERC721.Contract.GetPlatformFeeInfo(&_DropERC721.CallOpts)
}

// GetPlatformFeeInfo is a free data retrieval call binding the contract method 0xd45573f6.
//
// Solidity: function getPlatformFeeInfo() view returns(address, uint16)
func (_DropERC721 *DropERC721CallerSession) GetPlatformFeeInfo() (common.Address, uint16, error) {
	return _DropERC721.Contract.GetPlatformFeeInfo(&_DropERC721.CallOpts)
}

// GetRevealURI is a free data retrieval call binding the contract method 0x9fc4d68f.
//
// Solidity: function getRevealURI(uint256 _batchId, bytes _key) view returns(string revealedURI)
func (_DropERC721 *DropERC721Caller) GetRevealURI(opts *bind.CallOpts, _batchId *big.Int, _key []byte) (string, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getRevealURI", _batchId, _key)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetRevealURI is a free data retrieval call binding the contract method 0x9fc4d68f.
//
// Solidity: function getRevealURI(uint256 _batchId, bytes _key) view returns(string revealedURI)
func (_DropERC721 *DropERC721Session) GetRevealURI(_batchId *big.Int, _key []byte) (string, error) {
	return _DropERC721.Contract.GetRevealURI(&_DropERC721.CallOpts, _batchId, _key)
}

// GetRevealURI is a free data retrieval call binding the contract method 0x9fc4d68f.
//
// Solidity: function getRevealURI(uint256 _batchId, bytes _key) view returns(string revealedURI)
func (_DropERC721 *DropERC721CallerSession) GetRevealURI(_batchId *big.Int, _key []byte) (string, error) {
	return _DropERC721.Contract.GetRevealURI(&_DropERC721.CallOpts, _batchId, _key)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DropERC721 *DropERC721Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DropERC721 *DropERC721Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DropERC721.Contract.GetRoleAdmin(&_DropERC721.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_DropERC721 *DropERC721CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _DropERC721.Contract.GetRoleAdmin(&_DropERC721.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_DropERC721 *DropERC721Caller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_DropERC721 *DropERC721Session) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DropERC721.Contract.GetRoleMember(&_DropERC721.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address member)
func (_DropERC721 *DropERC721CallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _DropERC721.Contract.GetRoleMember(&_DropERC721.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_DropERC721 *DropERC721Caller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_DropERC721 *DropERC721Session) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DropERC721.Contract.GetRoleMemberCount(&_DropERC721.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256 count)
func (_DropERC721 *DropERC721CallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _DropERC721.Contract.GetRoleMemberCount(&_DropERC721.CallOpts, role)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_DropERC721 *DropERC721Caller) GetRoyaltyInfoForToken(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, uint16, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getRoyaltyInfoForToken", _tokenId)

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
func (_DropERC721 *DropERC721Session) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _DropERC721.Contract.GetRoyaltyInfoForToken(&_DropERC721.CallOpts, _tokenId)
}

// GetRoyaltyInfoForToken is a free data retrieval call binding the contract method 0x4cc157df.
//
// Solidity: function getRoyaltyInfoForToken(uint256 _tokenId) view returns(address, uint16)
func (_DropERC721 *DropERC721CallerSession) GetRoyaltyInfoForToken(_tokenId *big.Int) (common.Address, uint16, error) {
	return _DropERC721.Contract.GetRoyaltyInfoForToken(&_DropERC721.CallOpts, _tokenId)
}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_DropERC721 *DropERC721Caller) GetSupplyClaimedByWallet(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "getSupplyClaimedByWallet", _conditionId, _claimer)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_DropERC721 *DropERC721Session) GetSupplyClaimedByWallet(_conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	return _DropERC721.Contract.GetSupplyClaimedByWallet(&_DropERC721.CallOpts, _conditionId, _claimer)
}

// GetSupplyClaimedByWallet is a free data retrieval call binding the contract method 0xad1eefc5.
//
// Solidity: function getSupplyClaimedByWallet(uint256 _conditionId, address _claimer) view returns(uint256 supplyClaimedByWallet)
func (_DropERC721 *DropERC721CallerSession) GetSupplyClaimedByWallet(_conditionId *big.Int, _claimer common.Address) (*big.Int, error) {
	return _DropERC721.Contract.GetSupplyClaimedByWallet(&_DropERC721.CallOpts, _conditionId, _claimer)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DropERC721 *DropERC721Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DropERC721 *DropERC721Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DropERC721.Contract.HasRole(&_DropERC721.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_DropERC721 *DropERC721CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _DropERC721.Contract.HasRole(&_DropERC721.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_DropERC721 *DropERC721Caller) HasRoleWithSwitch(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "hasRoleWithSwitch", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_DropERC721 *DropERC721Session) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _DropERC721.Contract.HasRoleWithSwitch(&_DropERC721.CallOpts, role, account)
}

// HasRoleWithSwitch is a free data retrieval call binding the contract method 0xa32fa5b3.
//
// Solidity: function hasRoleWithSwitch(bytes32 role, address account) view returns(bool)
func (_DropERC721 *DropERC721CallerSession) HasRoleWithSwitch(role [32]byte, account common.Address) (bool, error) {
	return _DropERC721.Contract.HasRoleWithSwitch(&_DropERC721.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DropERC721 *DropERC721Caller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DropERC721 *DropERC721Session) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DropERC721.Contract.IsApprovedForAll(&_DropERC721.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_DropERC721 *DropERC721CallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _DropERC721.Contract.IsApprovedForAll(&_DropERC721.CallOpts, owner, operator)
}

// IsEncryptedBatch is a free data retrieval call binding the contract method 0x492e224b.
//
// Solidity: function isEncryptedBatch(uint256 _batchId) view returns(bool)
func (_DropERC721 *DropERC721Caller) IsEncryptedBatch(opts *bind.CallOpts, _batchId *big.Int) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "isEncryptedBatch", _batchId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEncryptedBatch is a free data retrieval call binding the contract method 0x492e224b.
//
// Solidity: function isEncryptedBatch(uint256 _batchId) view returns(bool)
func (_DropERC721 *DropERC721Session) IsEncryptedBatch(_batchId *big.Int) (bool, error) {
	return _DropERC721.Contract.IsEncryptedBatch(&_DropERC721.CallOpts, _batchId)
}

// IsEncryptedBatch is a free data retrieval call binding the contract method 0x492e224b.
//
// Solidity: function isEncryptedBatch(uint256 _batchId) view returns(bool)
func (_DropERC721 *DropERC721CallerSession) IsEncryptedBatch(_batchId *big.Int) (bool, error) {
	return _DropERC721.Contract.IsEncryptedBatch(&_DropERC721.CallOpts, _batchId)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_DropERC721 *DropERC721Caller) IsTrustedForwarder(opts *bind.CallOpts, forwarder common.Address) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "isTrustedForwarder", forwarder)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_DropERC721 *DropERC721Session) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _DropERC721.Contract.IsTrustedForwarder(&_DropERC721.CallOpts, forwarder)
}

// IsTrustedForwarder is a free data retrieval call binding the contract method 0x572b6c05.
//
// Solidity: function isTrustedForwarder(address forwarder) view returns(bool)
func (_DropERC721 *DropERC721CallerSession) IsTrustedForwarder(forwarder common.Address) (bool, error) {
	return _DropERC721.Contract.IsTrustedForwarder(&_DropERC721.CallOpts, forwarder)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_DropERC721 *DropERC721Caller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "maxTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_DropERC721 *DropERC721Session) MaxTotalSupply() (*big.Int, error) {
	return _DropERC721.Contract.MaxTotalSupply(&_DropERC721.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) MaxTotalSupply() (*big.Int, error) {
	return _DropERC721.Contract.MaxTotalSupply(&_DropERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DropERC721 *DropERC721Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DropERC721 *DropERC721Session) Name() (string, error) {
	return _DropERC721.Contract.Name(&_DropERC721.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DropERC721 *DropERC721CallerSession) Name() (string, error) {
	return _DropERC721.Contract.Name(&_DropERC721.CallOpts)
}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_DropERC721 *DropERC721Caller) NextTokenIdToClaim(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "nextTokenIdToClaim")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_DropERC721 *DropERC721Session) NextTokenIdToClaim() (*big.Int, error) {
	return _DropERC721.Contract.NextTokenIdToClaim(&_DropERC721.CallOpts)
}

// NextTokenIdToClaim is a free data retrieval call binding the contract method 0xacd083f8.
//
// Solidity: function nextTokenIdToClaim() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) NextTokenIdToClaim() (*big.Int, error) {
	return _DropERC721.Contract.NextTokenIdToClaim(&_DropERC721.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_DropERC721 *DropERC721Caller) NextTokenIdToMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "nextTokenIdToMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_DropERC721 *DropERC721Session) NextTokenIdToMint() (*big.Int, error) {
	return _DropERC721.Contract.NextTokenIdToMint(&_DropERC721.CallOpts)
}

// NextTokenIdToMint is a free data retrieval call binding the contract method 0x3b1475a7.
//
// Solidity: function nextTokenIdToMint() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) NextTokenIdToMint() (*big.Int, error) {
	return _DropERC721.Contract.NextTokenIdToMint(&_DropERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DropERC721 *DropERC721Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DropERC721 *DropERC721Session) Owner() (common.Address, error) {
	return _DropERC721.Contract.Owner(&_DropERC721.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DropERC721 *DropERC721CallerSession) Owner() (common.Address, error) {
	return _DropERC721.Contract.Owner(&_DropERC721.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DropERC721 *DropERC721Caller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DropERC721 *DropERC721Session) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DropERC721.Contract.OwnerOf(&_DropERC721.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DropERC721 *DropERC721CallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DropERC721.Contract.OwnerOf(&_DropERC721.CallOpts, tokenId)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_DropERC721 *DropERC721Caller) PrimarySaleRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "primarySaleRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_DropERC721 *DropERC721Session) PrimarySaleRecipient() (common.Address, error) {
	return _DropERC721.Contract.PrimarySaleRecipient(&_DropERC721.CallOpts)
}

// PrimarySaleRecipient is a free data retrieval call binding the contract method 0x079fe40e.
//
// Solidity: function primarySaleRecipient() view returns(address)
func (_DropERC721 *DropERC721CallerSession) PrimarySaleRecipient() (common.Address, error) {
	return _DropERC721.Contract.PrimarySaleRecipient(&_DropERC721.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_DropERC721 *DropERC721Caller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "royaltyInfo", tokenId, salePrice)

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
func (_DropERC721 *DropERC721Session) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _DropERC721.Contract.RoyaltyInfo(&_DropERC721.CallOpts, tokenId, salePrice)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 salePrice) view returns(address receiver, uint256 royaltyAmount)
func (_DropERC721 *DropERC721CallerSession) RoyaltyInfo(tokenId *big.Int, salePrice *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _DropERC721.Contract.RoyaltyInfo(&_DropERC721.CallOpts, tokenId, salePrice)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DropERC721 *DropERC721Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DropERC721 *DropERC721Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DropERC721.Contract.SupportsInterface(&_DropERC721.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DropERC721 *DropERC721CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DropERC721.Contract.SupportsInterface(&_DropERC721.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DropERC721 *DropERC721Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DropERC721 *DropERC721Session) Symbol() (string, error) {
	return _DropERC721.Contract.Symbol(&_DropERC721.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DropERC721 *DropERC721CallerSession) Symbol() (string, error) {
	return _DropERC721.Contract.Symbol(&_DropERC721.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DropERC721 *DropERC721Caller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "tokenURI", _tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DropERC721 *DropERC721Session) TokenURI(_tokenId *big.Int) (string, error) {
	return _DropERC721.Contract.TokenURI(&_DropERC721.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 _tokenId) view returns(string)
func (_DropERC721 *DropERC721CallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _DropERC721.Contract.TokenURI(&_DropERC721.CallOpts, _tokenId)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_DropERC721 *DropERC721Caller) TotalMinted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "totalMinted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_DropERC721 *DropERC721Session) TotalMinted() (*big.Int, error) {
	return _DropERC721.Contract.TotalMinted(&_DropERC721.CallOpts)
}

// TotalMinted is a free data retrieval call binding the contract method 0xa2309ff8.
//
// Solidity: function totalMinted() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) TotalMinted() (*big.Int, error) {
	return _DropERC721.Contract.TotalMinted(&_DropERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DropERC721 *DropERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DropERC721 *DropERC721Session) TotalSupply() (*big.Int, error) {
	return _DropERC721.Contract.TotalSupply(&_DropERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DropERC721 *DropERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _DropERC721.Contract.TotalSupply(&_DropERC721.CallOpts)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_DropERC721 *DropERC721Caller) VerifyClaim(opts *bind.CallOpts, _conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	var out []interface{}
	err := _DropERC721.contract.Call(opts, &out, "verifyClaim", _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_DropERC721 *DropERC721Session) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	return _DropERC721.Contract.VerifyClaim(&_DropERC721.CallOpts, _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)
}

// VerifyClaim is a free data retrieval call binding the contract method 0x23a2902b.
//
// Solidity: function verifyClaim(uint256 _conditionId, address _claimer, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof) view returns(bool isOverride)
func (_DropERC721 *DropERC721CallerSession) VerifyClaim(_conditionId *big.Int, _claimer common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof) (bool, error) {
	return _DropERC721.Contract.VerifyClaim(&_DropERC721.CallOpts, _conditionId, _claimer, _quantity, _currency, _pricePerToken, _allowlistProof)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721Transactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721Session) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.Approve(&_DropERC721.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721TransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.Approve(&_DropERC721.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_DropERC721 *DropERC721Transactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_DropERC721 *DropERC721Session) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.Burn(&_DropERC721.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_DropERC721 *DropERC721TransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.Burn(&_DropERC721.TransactOpts, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_DropERC721 *DropERC721Transactor) Claim(opts *bind.TransactOpts, _receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "claim", _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_DropERC721 *DropERC721Session) Claim(_receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.Claim(&_DropERC721.TransactOpts, _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// Claim is a paid mutator transaction binding the contract method 0x84bb1e42.
//
// Solidity: function claim(address _receiver, uint256 _quantity, address _currency, uint256 _pricePerToken, (bytes32[],uint256,uint256,address) _allowlistProof, bytes _data) payable returns()
func (_DropERC721 *DropERC721TransactorSession) Claim(_receiver common.Address, _quantity *big.Int, _currency common.Address, _pricePerToken *big.Int, _allowlistProof IDropAllowlistProof, _data []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.Claim(&_DropERC721.TransactOpts, _receiver, _quantity, _currency, _pricePerToken, _allowlistProof, _data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.GrantRole(&_DropERC721.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.GrantRole(&_DropERC721.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_DropERC721 *DropERC721Transactor) Initialize(opts *bind.TransactOpts, _defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "initialize", _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_DropERC721 *DropERC721Session) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.Initialize(&_DropERC721.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xe1591634.
//
// Solidity: function initialize(address _defaultAdmin, string _name, string _symbol, string _contractURI, address[] _trustedForwarders, address _saleRecipient, address _royaltyRecipient, uint128 _royaltyBps, uint128 _platformFeeBps, address _platformFeeRecipient) returns()
func (_DropERC721 *DropERC721TransactorSession) Initialize(_defaultAdmin common.Address, _name string, _symbol string, _contractURI string, _trustedForwarders []common.Address, _saleRecipient common.Address, _royaltyRecipient common.Address, _royaltyBps *big.Int, _platformFeeBps *big.Int, _platformFeeRecipient common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.Initialize(&_DropERC721.TransactOpts, _defaultAdmin, _name, _symbol, _contractURI, _trustedForwarders, _saleRecipient, _royaltyRecipient, _royaltyBps, _platformFeeBps, _platformFeeRecipient)
}

// LazyMint is a paid mutator transaction binding the contract method 0xd37c353b.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens, bytes _data) returns(uint256 batchId)
func (_DropERC721 *DropERC721Transactor) LazyMint(opts *bind.TransactOpts, _amount *big.Int, _baseURIForTokens string, _data []byte) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "lazyMint", _amount, _baseURIForTokens, _data)
}

// LazyMint is a paid mutator transaction binding the contract method 0xd37c353b.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens, bytes _data) returns(uint256 batchId)
func (_DropERC721 *DropERC721Session) LazyMint(_amount *big.Int, _baseURIForTokens string, _data []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.LazyMint(&_DropERC721.TransactOpts, _amount, _baseURIForTokens, _data)
}

// LazyMint is a paid mutator transaction binding the contract method 0xd37c353b.
//
// Solidity: function lazyMint(uint256 _amount, string _baseURIForTokens, bytes _data) returns(uint256 batchId)
func (_DropERC721 *DropERC721TransactorSession) LazyMint(_amount *big.Int, _baseURIForTokens string, _data []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.LazyMint(&_DropERC721.TransactOpts, _amount, _baseURIForTokens, _data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_DropERC721 *DropERC721Transactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_DropERC721 *DropERC721Session) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DropERC721.Contract.Multicall(&_DropERC721.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_DropERC721 *DropERC721TransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _DropERC721.Contract.Multicall(&_DropERC721.TransactOpts, data)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.RenounceRole(&_DropERC721.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.RenounceRole(&_DropERC721.TransactOpts, role, account)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _index, bytes _key) returns(string revealedURI)
func (_DropERC721 *DropERC721Transactor) Reveal(opts *bind.TransactOpts, _index *big.Int, _key []byte) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "reveal", _index, _key)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _index, bytes _key) returns(string revealedURI)
func (_DropERC721 *DropERC721Session) Reveal(_index *big.Int, _key []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.Reveal(&_DropERC721.TransactOpts, _index, _key)
}

// Reveal is a paid mutator transaction binding the contract method 0xce805642.
//
// Solidity: function reveal(uint256 _index, bytes _key) returns(string revealedURI)
func (_DropERC721 *DropERC721TransactorSession) Reveal(_index *big.Int, _key []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.Reveal(&_DropERC721.TransactOpts, _index, _key)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.RevokeRole(&_DropERC721.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_DropERC721 *DropERC721TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.RevokeRole(&_DropERC721.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721Session) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SafeTransferFrom(&_DropERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721TransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SafeTransferFrom(&_DropERC721.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DropERC721 *DropERC721Transactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DropERC721 *DropERC721Session) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.SafeTransferFrom0(&_DropERC721.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_DropERC721 *DropERC721TransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _DropERC721.Contract.SafeTransferFrom0(&_DropERC721.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DropERC721 *DropERC721Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DropERC721 *DropERC721Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DropERC721.Contract.SetApprovalForAll(&_DropERC721.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DropERC721 *DropERC721TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DropERC721.Contract.SetApprovalForAll(&_DropERC721.TransactOpts, operator, approved)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_DropERC721 *DropERC721Transactor) SetClaimConditions(opts *bind.TransactOpts, _conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setClaimConditions", _conditions, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_DropERC721 *DropERC721Session) SetClaimConditions(_conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _DropERC721.Contract.SetClaimConditions(&_DropERC721.TransactOpts, _conditions, _resetClaimEligibility)
}

// SetClaimConditions is a paid mutator transaction binding the contract method 0x74bc7db7.
//
// Solidity: function setClaimConditions((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] _conditions, bool _resetClaimEligibility) returns()
func (_DropERC721 *DropERC721TransactorSession) SetClaimConditions(_conditions []IClaimConditionClaimCondition, _resetClaimEligibility bool) (*types.Transaction, error) {
	return _DropERC721.Contract.SetClaimConditions(&_DropERC721.TransactOpts, _conditions, _resetClaimEligibility)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_DropERC721 *DropERC721Transactor) SetContractURI(opts *bind.TransactOpts, _uri string) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setContractURI", _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_DropERC721 *DropERC721Session) SetContractURI(_uri string) (*types.Transaction, error) {
	return _DropERC721.Contract.SetContractURI(&_DropERC721.TransactOpts, _uri)
}

// SetContractURI is a paid mutator transaction binding the contract method 0x938e3d7b.
//
// Solidity: function setContractURI(string _uri) returns()
func (_DropERC721 *DropERC721TransactorSession) SetContractURI(_uri string) (*types.Transaction, error) {
	return _DropERC721.Contract.SetContractURI(&_DropERC721.TransactOpts, _uri)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_DropERC721 *DropERC721Transactor) SetDefaultRoyaltyInfo(opts *bind.TransactOpts, _royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setDefaultRoyaltyInfo", _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_DropERC721 *DropERC721Session) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetDefaultRoyaltyInfo(&_DropERC721.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetDefaultRoyaltyInfo is a paid mutator transaction binding the contract method 0x600dd5ea.
//
// Solidity: function setDefaultRoyaltyInfo(address _royaltyRecipient, uint256 _royaltyBps) returns()
func (_DropERC721 *DropERC721TransactorSession) SetDefaultRoyaltyInfo(_royaltyRecipient common.Address, _royaltyBps *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetDefaultRoyaltyInfo(&_DropERC721.TransactOpts, _royaltyRecipient, _royaltyBps)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_DropERC721 *DropERC721Transactor) SetMaxTotalSupply(opts *bind.TransactOpts, _maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setMaxTotalSupply", _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_DropERC721 *DropERC721Session) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetMaxTotalSupply(&_DropERC721.TransactOpts, _maxTotalSupply)
}

// SetMaxTotalSupply is a paid mutator transaction binding the contract method 0x3f3e4c11.
//
// Solidity: function setMaxTotalSupply(uint256 _maxTotalSupply) returns()
func (_DropERC721 *DropERC721TransactorSession) SetMaxTotalSupply(_maxTotalSupply *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetMaxTotalSupply(&_DropERC721.TransactOpts, _maxTotalSupply)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_DropERC721 *DropERC721Transactor) SetOwner(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setOwner", _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_DropERC721 *DropERC721Session) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.SetOwner(&_DropERC721.TransactOpts, _newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address _newOwner) returns()
func (_DropERC721 *DropERC721TransactorSession) SetOwner(_newOwner common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.SetOwner(&_DropERC721.TransactOpts, _newOwner)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_DropERC721 *DropERC721Transactor) SetPlatformFeeInfo(opts *bind.TransactOpts, _platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setPlatformFeeInfo", _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_DropERC721 *DropERC721Session) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetPlatformFeeInfo(&_DropERC721.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPlatformFeeInfo is a paid mutator transaction binding the contract method 0x1e7ac488.
//
// Solidity: function setPlatformFeeInfo(address _platformFeeRecipient, uint256 _platformFeeBps) returns()
func (_DropERC721 *DropERC721TransactorSession) SetPlatformFeeInfo(_platformFeeRecipient common.Address, _platformFeeBps *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetPlatformFeeInfo(&_DropERC721.TransactOpts, _platformFeeRecipient, _platformFeeBps)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_DropERC721 *DropERC721Transactor) SetPrimarySaleRecipient(opts *bind.TransactOpts, _saleRecipient common.Address) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setPrimarySaleRecipient", _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_DropERC721 *DropERC721Session) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.SetPrimarySaleRecipient(&_DropERC721.TransactOpts, _saleRecipient)
}

// SetPrimarySaleRecipient is a paid mutator transaction binding the contract method 0x6f4f2837.
//
// Solidity: function setPrimarySaleRecipient(address _saleRecipient) returns()
func (_DropERC721 *DropERC721TransactorSession) SetPrimarySaleRecipient(_saleRecipient common.Address) (*types.Transaction, error) {
	return _DropERC721.Contract.SetPrimarySaleRecipient(&_DropERC721.TransactOpts, _saleRecipient)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_DropERC721 *DropERC721Transactor) SetRoyaltyInfoForToken(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "setRoyaltyInfoForToken", _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_DropERC721 *DropERC721Session) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetRoyaltyInfoForToken(&_DropERC721.TransactOpts, _tokenId, _recipient, _bps)
}

// SetRoyaltyInfoForToken is a paid mutator transaction binding the contract method 0x9bcf7a15.
//
// Solidity: function setRoyaltyInfoForToken(uint256 _tokenId, address _recipient, uint256 _bps) returns()
func (_DropERC721 *DropERC721TransactorSession) SetRoyaltyInfoForToken(_tokenId *big.Int, _recipient common.Address, _bps *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.SetRoyaltyInfoForToken(&_DropERC721.TransactOpts, _tokenId, _recipient, _bps)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721Session) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.TransferFrom(&_DropERC721.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DropERC721 *DropERC721TransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DropERC721.Contract.TransferFrom(&_DropERC721.TransactOpts, from, to, tokenId)
}

// DropERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DropERC721 contract.
type DropERC721ApprovalIterator struct {
	Event *DropERC721Approval // Event containing the contract specifics and raw log

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
func (it *DropERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721Approval)
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
		it.Event = new(DropERC721Approval)
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
func (it *DropERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721Approval represents a Approval event raised by the DropERC721 contract.
type DropERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DropERC721 *DropERC721Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DropERC721ApprovalIterator, error) {

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

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721ApprovalIterator{contract: _DropERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DropERC721 *DropERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DropERC721Approval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721Approval)
				if err := _DropERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseApproval(log types.Log) (*DropERC721Approval, error) {
	event := new(DropERC721Approval)
	if err := _DropERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DropERC721 contract.
type DropERC721ApprovalForAllIterator struct {
	Event *DropERC721ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DropERC721ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721ApprovalForAll)
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
		it.Event = new(DropERC721ApprovalForAll)
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
func (it *DropERC721ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721ApprovalForAll represents a ApprovalForAll event raised by the DropERC721 contract.
type DropERC721ApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DropERC721 *DropERC721Filterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DropERC721ApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721ApprovalForAllIterator{contract: _DropERC721.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DropERC721 *DropERC721Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DropERC721ApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721ApprovalForAll)
				if err := _DropERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseApprovalForAll(log types.Log) (*DropERC721ApprovalForAll, error) {
	event := new(DropERC721ApprovalForAll)
	if err := _DropERC721.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721ClaimConditionsUpdatedIterator is returned from FilterClaimConditionsUpdated and is used to iterate over the raw logs and unpacked data for ClaimConditionsUpdated events raised by the DropERC721 contract.
type DropERC721ClaimConditionsUpdatedIterator struct {
	Event *DropERC721ClaimConditionsUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC721ClaimConditionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721ClaimConditionsUpdated)
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
		it.Event = new(DropERC721ClaimConditionsUpdated)
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
func (it *DropERC721ClaimConditionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721ClaimConditionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721ClaimConditionsUpdated represents a ClaimConditionsUpdated event raised by the DropERC721 contract.
type DropERC721ClaimConditionsUpdated struct {
	ClaimConditions  []IClaimConditionClaimCondition
	ResetEligibility bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimConditionsUpdated is a free log retrieval operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_DropERC721 *DropERC721Filterer) FilterClaimConditionsUpdated(opts *bind.FilterOpts) (*DropERC721ClaimConditionsUpdatedIterator, error) {

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "ClaimConditionsUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC721ClaimConditionsUpdatedIterator{contract: _DropERC721.contract, event: "ClaimConditionsUpdated", logs: logs, sub: sub}, nil
}

// WatchClaimConditionsUpdated is a free log subscription operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_DropERC721 *DropERC721Filterer) WatchClaimConditionsUpdated(opts *bind.WatchOpts, sink chan<- *DropERC721ClaimConditionsUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "ClaimConditionsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721ClaimConditionsUpdated)
				if err := _DropERC721.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
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

// ParseClaimConditionsUpdated is a log parse operation binding the contract event 0xbf4016fceeaaa4ac5cf4be865b559ff85825ab4ca7aa7b661d16e2f544c03098.
//
// Solidity: event ClaimConditionsUpdated((uint256,uint256,uint256,uint256,bytes32,uint256,address,string)[] claimConditions, bool resetEligibility)
func (_DropERC721 *DropERC721Filterer) ParseClaimConditionsUpdated(log types.Log) (*DropERC721ClaimConditionsUpdated, error) {
	event := new(DropERC721ClaimConditionsUpdated)
	if err := _DropERC721.contract.UnpackLog(event, "ClaimConditionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721ContractURIUpdatedIterator is returned from FilterContractURIUpdated and is used to iterate over the raw logs and unpacked data for ContractURIUpdated events raised by the DropERC721 contract.
type DropERC721ContractURIUpdatedIterator struct {
	Event *DropERC721ContractURIUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC721ContractURIUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721ContractURIUpdated)
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
		it.Event = new(DropERC721ContractURIUpdated)
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
func (it *DropERC721ContractURIUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721ContractURIUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721ContractURIUpdated represents a ContractURIUpdated event raised by the DropERC721 contract.
type DropERC721ContractURIUpdated struct {
	PrevURI string
	NewURI  string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractURIUpdated is a free log retrieval operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_DropERC721 *DropERC721Filterer) FilterContractURIUpdated(opts *bind.FilterOpts) (*DropERC721ContractURIUpdatedIterator, error) {

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC721ContractURIUpdatedIterator{contract: _DropERC721.contract, event: "ContractURIUpdated", logs: logs, sub: sub}, nil
}

// WatchContractURIUpdated is a free log subscription operation binding the contract event 0xc9c7c3fe08b88b4df9d4d47ef47d2c43d55c025a0ba88ca442580ed9e7348a16.
//
// Solidity: event ContractURIUpdated(string prevURI, string newURI)
func (_DropERC721 *DropERC721Filterer) WatchContractURIUpdated(opts *bind.WatchOpts, sink chan<- *DropERC721ContractURIUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "ContractURIUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721ContractURIUpdated)
				if err := _DropERC721.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseContractURIUpdated(log types.Log) (*DropERC721ContractURIUpdated, error) {
	event := new(DropERC721ContractURIUpdated)
	if err := _DropERC721.contract.UnpackLog(event, "ContractURIUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721DefaultRoyaltyIterator is returned from FilterDefaultRoyalty and is used to iterate over the raw logs and unpacked data for DefaultRoyalty events raised by the DropERC721 contract.
type DropERC721DefaultRoyaltyIterator struct {
	Event *DropERC721DefaultRoyalty // Event containing the contract specifics and raw log

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
func (it *DropERC721DefaultRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721DefaultRoyalty)
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
		it.Event = new(DropERC721DefaultRoyalty)
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
func (it *DropERC721DefaultRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721DefaultRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721DefaultRoyalty represents a DefaultRoyalty event raised by the DropERC721 contract.
type DropERC721DefaultRoyalty struct {
	NewRoyaltyRecipient common.Address
	NewRoyaltyBps       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterDefaultRoyalty is a free log retrieval operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_DropERC721 *DropERC721Filterer) FilterDefaultRoyalty(opts *bind.FilterOpts, newRoyaltyRecipient []common.Address) (*DropERC721DefaultRoyaltyIterator, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721DefaultRoyaltyIterator{contract: _DropERC721.contract, event: "DefaultRoyalty", logs: logs, sub: sub}, nil
}

// WatchDefaultRoyalty is a free log subscription operation binding the contract event 0x90d7ec04bcb8978719414f82e52e4cb651db41d0e6f8cea6118c2191e6183adb.
//
// Solidity: event DefaultRoyalty(address indexed newRoyaltyRecipient, uint256 newRoyaltyBps)
func (_DropERC721 *DropERC721Filterer) WatchDefaultRoyalty(opts *bind.WatchOpts, sink chan<- *DropERC721DefaultRoyalty, newRoyaltyRecipient []common.Address) (event.Subscription, error) {

	var newRoyaltyRecipientRule []interface{}
	for _, newRoyaltyRecipientItem := range newRoyaltyRecipient {
		newRoyaltyRecipientRule = append(newRoyaltyRecipientRule, newRoyaltyRecipientItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "DefaultRoyalty", newRoyaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721DefaultRoyalty)
				if err := _DropERC721.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseDefaultRoyalty(log types.Log) (*DropERC721DefaultRoyalty, error) {
	event := new(DropERC721DefaultRoyalty)
	if err := _DropERC721.contract.UnpackLog(event, "DefaultRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DropERC721 contract.
type DropERC721InitializedIterator struct {
	Event *DropERC721Initialized // Event containing the contract specifics and raw log

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
func (it *DropERC721InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721Initialized)
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
		it.Event = new(DropERC721Initialized)
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
func (it *DropERC721InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721Initialized represents a Initialized event raised by the DropERC721 contract.
type DropERC721Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DropERC721 *DropERC721Filterer) FilterInitialized(opts *bind.FilterOpts) (*DropERC721InitializedIterator, error) {

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DropERC721InitializedIterator{contract: _DropERC721.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DropERC721 *DropERC721Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DropERC721Initialized) (event.Subscription, error) {

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721Initialized)
				if err := _DropERC721.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DropERC721 *DropERC721Filterer) ParseInitialized(log types.Log) (*DropERC721Initialized, error) {
	event := new(DropERC721Initialized)
	if err := _DropERC721.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721MaxTotalSupplyUpdatedIterator is returned from FilterMaxTotalSupplyUpdated and is used to iterate over the raw logs and unpacked data for MaxTotalSupplyUpdated events raised by the DropERC721 contract.
type DropERC721MaxTotalSupplyUpdatedIterator struct {
	Event *DropERC721MaxTotalSupplyUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC721MaxTotalSupplyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721MaxTotalSupplyUpdated)
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
		it.Event = new(DropERC721MaxTotalSupplyUpdated)
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
func (it *DropERC721MaxTotalSupplyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721MaxTotalSupplyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721MaxTotalSupplyUpdated represents a MaxTotalSupplyUpdated event raised by the DropERC721 contract.
type DropERC721MaxTotalSupplyUpdated struct {
	MaxTotalSupply *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMaxTotalSupplyUpdated is a free log retrieval operation binding the contract event 0xf2672935fc79f5237559e2e2999dbe743bf65430894ac2b37666890e7c69e1af.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 maxTotalSupply)
func (_DropERC721 *DropERC721Filterer) FilterMaxTotalSupplyUpdated(opts *bind.FilterOpts) (*DropERC721MaxTotalSupplyUpdatedIterator, error) {

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "MaxTotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return &DropERC721MaxTotalSupplyUpdatedIterator{contract: _DropERC721.contract, event: "MaxTotalSupplyUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxTotalSupplyUpdated is a free log subscription operation binding the contract event 0xf2672935fc79f5237559e2e2999dbe743bf65430894ac2b37666890e7c69e1af.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 maxTotalSupply)
func (_DropERC721 *DropERC721Filterer) WatchMaxTotalSupplyUpdated(opts *bind.WatchOpts, sink chan<- *DropERC721MaxTotalSupplyUpdated) (event.Subscription, error) {

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "MaxTotalSupplyUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721MaxTotalSupplyUpdated)
				if err := _DropERC721.contract.UnpackLog(event, "MaxTotalSupplyUpdated", log); err != nil {
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

// ParseMaxTotalSupplyUpdated is a log parse operation binding the contract event 0xf2672935fc79f5237559e2e2999dbe743bf65430894ac2b37666890e7c69e1af.
//
// Solidity: event MaxTotalSupplyUpdated(uint256 maxTotalSupply)
func (_DropERC721 *DropERC721Filterer) ParseMaxTotalSupplyUpdated(log types.Log) (*DropERC721MaxTotalSupplyUpdated, error) {
	event := new(DropERC721MaxTotalSupplyUpdated)
	if err := _DropERC721.contract.UnpackLog(event, "MaxTotalSupplyUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721OwnerUpdatedIterator is returned from FilterOwnerUpdated and is used to iterate over the raw logs and unpacked data for OwnerUpdated events raised by the DropERC721 contract.
type DropERC721OwnerUpdatedIterator struct {
	Event *DropERC721OwnerUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC721OwnerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721OwnerUpdated)
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
		it.Event = new(DropERC721OwnerUpdated)
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
func (it *DropERC721OwnerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721OwnerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721OwnerUpdated represents a OwnerUpdated event raised by the DropERC721 contract.
type DropERC721OwnerUpdated struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdated is a free log retrieval operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_DropERC721 *DropERC721Filterer) FilterOwnerUpdated(opts *bind.FilterOpts, prevOwner []common.Address, newOwner []common.Address) (*DropERC721OwnerUpdatedIterator, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721OwnerUpdatedIterator{contract: _DropERC721.contract, event: "OwnerUpdated", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdated is a free log subscription operation binding the contract event 0x8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d76.
//
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_DropERC721 *DropERC721Filterer) WatchOwnerUpdated(opts *bind.WatchOpts, sink chan<- *DropERC721OwnerUpdated, prevOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var prevOwnerRule []interface{}
	for _, prevOwnerItem := range prevOwner {
		prevOwnerRule = append(prevOwnerRule, prevOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "OwnerUpdated", prevOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721OwnerUpdated)
				if err := _DropERC721.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
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
// Solidity: event OwnerUpdated(address indexed prevOwner, address indexed newOwner)
func (_DropERC721 *DropERC721Filterer) ParseOwnerUpdated(log types.Log) (*DropERC721OwnerUpdated, error) {
	event := new(DropERC721OwnerUpdated)
	if err := _DropERC721.contract.UnpackLog(event, "OwnerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721PlatformFeeInfoUpdatedIterator is returned from FilterPlatformFeeInfoUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeInfoUpdated events raised by the DropERC721 contract.
type DropERC721PlatformFeeInfoUpdatedIterator struct {
	Event *DropERC721PlatformFeeInfoUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC721PlatformFeeInfoUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721PlatformFeeInfoUpdated)
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
		it.Event = new(DropERC721PlatformFeeInfoUpdated)
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
func (it *DropERC721PlatformFeeInfoUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721PlatformFeeInfoUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721PlatformFeeInfoUpdated represents a PlatformFeeInfoUpdated event raised by the DropERC721 contract.
type DropERC721PlatformFeeInfoUpdated struct {
	PlatformFeeRecipient common.Address
	PlatformFeeBps       *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeInfoUpdated is a free log retrieval operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_DropERC721 *DropERC721Filterer) FilterPlatformFeeInfoUpdated(opts *bind.FilterOpts, platformFeeRecipient []common.Address) (*DropERC721PlatformFeeInfoUpdatedIterator, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721PlatformFeeInfoUpdatedIterator{contract: _DropERC721.contract, event: "PlatformFeeInfoUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeInfoUpdated is a free log subscription operation binding the contract event 0xe2497bd806ec41a6e0dd992c29a72efc0ef8fec9092d1978fd4a1e00b2f18304.
//
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_DropERC721 *DropERC721Filterer) WatchPlatformFeeInfoUpdated(opts *bind.WatchOpts, sink chan<- *DropERC721PlatformFeeInfoUpdated, platformFeeRecipient []common.Address) (event.Subscription, error) {

	var platformFeeRecipientRule []interface{}
	for _, platformFeeRecipientItem := range platformFeeRecipient {
		platformFeeRecipientRule = append(platformFeeRecipientRule, platformFeeRecipientItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "PlatformFeeInfoUpdated", platformFeeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721PlatformFeeInfoUpdated)
				if err := _DropERC721.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
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
// Solidity: event PlatformFeeInfoUpdated(address indexed platformFeeRecipient, uint256 platformFeeBps)
func (_DropERC721 *DropERC721Filterer) ParsePlatformFeeInfoUpdated(log types.Log) (*DropERC721PlatformFeeInfoUpdated, error) {
	event := new(DropERC721PlatformFeeInfoUpdated)
	if err := _DropERC721.contract.UnpackLog(event, "PlatformFeeInfoUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721PrimarySaleRecipientUpdatedIterator is returned from FilterPrimarySaleRecipientUpdated and is used to iterate over the raw logs and unpacked data for PrimarySaleRecipientUpdated events raised by the DropERC721 contract.
type DropERC721PrimarySaleRecipientUpdatedIterator struct {
	Event *DropERC721PrimarySaleRecipientUpdated // Event containing the contract specifics and raw log

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
func (it *DropERC721PrimarySaleRecipientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721PrimarySaleRecipientUpdated)
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
		it.Event = new(DropERC721PrimarySaleRecipientUpdated)
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
func (it *DropERC721PrimarySaleRecipientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721PrimarySaleRecipientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721PrimarySaleRecipientUpdated represents a PrimarySaleRecipientUpdated event raised by the DropERC721 contract.
type DropERC721PrimarySaleRecipientUpdated struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPrimarySaleRecipientUpdated is a free log retrieval operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_DropERC721 *DropERC721Filterer) FilterPrimarySaleRecipientUpdated(opts *bind.FilterOpts, recipient []common.Address) (*DropERC721PrimarySaleRecipientUpdatedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721PrimarySaleRecipientUpdatedIterator{contract: _DropERC721.contract, event: "PrimarySaleRecipientUpdated", logs: logs, sub: sub}, nil
}

// WatchPrimarySaleRecipientUpdated is a free log subscription operation binding the contract event 0x299d17e95023f496e0ffc4909cff1a61f74bb5eb18de6f900f4155bfa1b3b333.
//
// Solidity: event PrimarySaleRecipientUpdated(address indexed recipient)
func (_DropERC721 *DropERC721Filterer) WatchPrimarySaleRecipientUpdated(opts *bind.WatchOpts, sink chan<- *DropERC721PrimarySaleRecipientUpdated, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "PrimarySaleRecipientUpdated", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721PrimarySaleRecipientUpdated)
				if err := _DropERC721.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParsePrimarySaleRecipientUpdated(log types.Log) (*DropERC721PrimarySaleRecipientUpdated, error) {
	event := new(DropERC721PrimarySaleRecipientUpdated)
	if err := _DropERC721.contract.UnpackLog(event, "PrimarySaleRecipientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the DropERC721 contract.
type DropERC721RoleAdminChangedIterator struct {
	Event *DropERC721RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *DropERC721RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721RoleAdminChanged)
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
		it.Event = new(DropERC721RoleAdminChanged)
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
func (it *DropERC721RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721RoleAdminChanged represents a RoleAdminChanged event raised by the DropERC721 contract.
type DropERC721RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DropERC721 *DropERC721Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*DropERC721RoleAdminChangedIterator, error) {

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

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721RoleAdminChangedIterator{contract: _DropERC721.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_DropERC721 *DropERC721Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *DropERC721RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721RoleAdminChanged)
				if err := _DropERC721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseRoleAdminChanged(log types.Log) (*DropERC721RoleAdminChanged, error) {
	event := new(DropERC721RoleAdminChanged)
	if err := _DropERC721.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the DropERC721 contract.
type DropERC721RoleGrantedIterator struct {
	Event *DropERC721RoleGranted // Event containing the contract specifics and raw log

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
func (it *DropERC721RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721RoleGranted)
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
		it.Event = new(DropERC721RoleGranted)
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
func (it *DropERC721RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721RoleGranted represents a RoleGranted event raised by the DropERC721 contract.
type DropERC721RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC721 *DropERC721Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DropERC721RoleGrantedIterator, error) {

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

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721RoleGrantedIterator{contract: _DropERC721.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC721 *DropERC721Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *DropERC721RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721RoleGranted)
				if err := _DropERC721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseRoleGranted(log types.Log) (*DropERC721RoleGranted, error) {
	event := new(DropERC721RoleGranted)
	if err := _DropERC721.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the DropERC721 contract.
type DropERC721RoleRevokedIterator struct {
	Event *DropERC721RoleRevoked // Event containing the contract specifics and raw log

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
func (it *DropERC721RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721RoleRevoked)
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
		it.Event = new(DropERC721RoleRevoked)
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
func (it *DropERC721RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721RoleRevoked represents a RoleRevoked event raised by the DropERC721 contract.
type DropERC721RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC721 *DropERC721Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*DropERC721RoleRevokedIterator, error) {

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

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721RoleRevokedIterator{contract: _DropERC721.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_DropERC721 *DropERC721Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *DropERC721RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721RoleRevoked)
				if err := _DropERC721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseRoleRevoked(log types.Log) (*DropERC721RoleRevoked, error) {
	event := new(DropERC721RoleRevoked)
	if err := _DropERC721.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721RoyaltyForTokenIterator is returned from FilterRoyaltyForToken and is used to iterate over the raw logs and unpacked data for RoyaltyForToken events raised by the DropERC721 contract.
type DropERC721RoyaltyForTokenIterator struct {
	Event *DropERC721RoyaltyForToken // Event containing the contract specifics and raw log

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
func (it *DropERC721RoyaltyForTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721RoyaltyForToken)
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
		it.Event = new(DropERC721RoyaltyForToken)
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
func (it *DropERC721RoyaltyForTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721RoyaltyForTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721RoyaltyForToken represents a RoyaltyForToken event raised by the DropERC721 contract.
type DropERC721RoyaltyForToken struct {
	TokenId          *big.Int
	RoyaltyRecipient common.Address
	RoyaltyBps       *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRoyaltyForToken is a free log retrieval operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_DropERC721 *DropERC721Filterer) FilterRoyaltyForToken(opts *bind.FilterOpts, tokenId []*big.Int, royaltyRecipient []common.Address) (*DropERC721RoyaltyForTokenIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721RoyaltyForTokenIterator{contract: _DropERC721.contract, event: "RoyaltyForToken", logs: logs, sub: sub}, nil
}

// WatchRoyaltyForToken is a free log subscription operation binding the contract event 0x7365cf4122f072a3365c20d54eff9b38d73c096c28e1892ec8f5b0e403a0f12d.
//
// Solidity: event RoyaltyForToken(uint256 indexed tokenId, address indexed royaltyRecipient, uint256 royaltyBps)
func (_DropERC721 *DropERC721Filterer) WatchRoyaltyForToken(opts *bind.WatchOpts, sink chan<- *DropERC721RoyaltyForToken, tokenId []*big.Int, royaltyRecipient []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var royaltyRecipientRule []interface{}
	for _, royaltyRecipientItem := range royaltyRecipient {
		royaltyRecipientRule = append(royaltyRecipientRule, royaltyRecipientItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "RoyaltyForToken", tokenIdRule, royaltyRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721RoyaltyForToken)
				if err := _DropERC721.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseRoyaltyForToken(log types.Log) (*DropERC721RoyaltyForToken, error) {
	event := new(DropERC721RoyaltyForToken)
	if err := _DropERC721.contract.UnpackLog(event, "RoyaltyForToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721TokenURIRevealedIterator is returned from FilterTokenURIRevealed and is used to iterate over the raw logs and unpacked data for TokenURIRevealed events raised by the DropERC721 contract.
type DropERC721TokenURIRevealedIterator struct {
	Event *DropERC721TokenURIRevealed // Event containing the contract specifics and raw log

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
func (it *DropERC721TokenURIRevealedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721TokenURIRevealed)
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
		it.Event = new(DropERC721TokenURIRevealed)
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
func (it *DropERC721TokenURIRevealedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721TokenURIRevealedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721TokenURIRevealed represents a TokenURIRevealed event raised by the DropERC721 contract.
type DropERC721TokenURIRevealed struct {
	Index       *big.Int
	RevealedURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenURIRevealed is a free log retrieval operation binding the contract event 0x6df1d8db2a036436ffe0b2d1833f2c5f1e624818dfce2578c0faa4b83ef9998d.
//
// Solidity: event TokenURIRevealed(uint256 indexed index, string revealedURI)
func (_DropERC721 *DropERC721Filterer) FilterTokenURIRevealed(opts *bind.FilterOpts, index []*big.Int) (*DropERC721TokenURIRevealedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "TokenURIRevealed", indexRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721TokenURIRevealedIterator{contract: _DropERC721.contract, event: "TokenURIRevealed", logs: logs, sub: sub}, nil
}

// WatchTokenURIRevealed is a free log subscription operation binding the contract event 0x6df1d8db2a036436ffe0b2d1833f2c5f1e624818dfce2578c0faa4b83ef9998d.
//
// Solidity: event TokenURIRevealed(uint256 indexed index, string revealedURI)
func (_DropERC721 *DropERC721Filterer) WatchTokenURIRevealed(opts *bind.WatchOpts, sink chan<- *DropERC721TokenURIRevealed, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "TokenURIRevealed", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721TokenURIRevealed)
				if err := _DropERC721.contract.UnpackLog(event, "TokenURIRevealed", log); err != nil {
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

// ParseTokenURIRevealed is a log parse operation binding the contract event 0x6df1d8db2a036436ffe0b2d1833f2c5f1e624818dfce2578c0faa4b83ef9998d.
//
// Solidity: event TokenURIRevealed(uint256 indexed index, string revealedURI)
func (_DropERC721 *DropERC721Filterer) ParseTokenURIRevealed(log types.Log) (*DropERC721TokenURIRevealed, error) {
	event := new(DropERC721TokenURIRevealed)
	if err := _DropERC721.contract.UnpackLog(event, "TokenURIRevealed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721TokensClaimedIterator is returned from FilterTokensClaimed and is used to iterate over the raw logs and unpacked data for TokensClaimed events raised by the DropERC721 contract.
type DropERC721TokensClaimedIterator struct {
	Event *DropERC721TokensClaimed // Event containing the contract specifics and raw log

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
func (it *DropERC721TokensClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721TokensClaimed)
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
		it.Event = new(DropERC721TokensClaimed)
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
func (it *DropERC721TokensClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721TokensClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721TokensClaimed represents a TokensClaimed event raised by the DropERC721 contract.
type DropERC721TokensClaimed struct {
	ClaimConditionIndex *big.Int
	Claimer             common.Address
	Receiver            common.Address
	StartTokenId        *big.Int
	QuantityClaimed     *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterTokensClaimed is a free log retrieval operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_DropERC721 *DropERC721Filterer) FilterTokensClaimed(opts *bind.FilterOpts, claimConditionIndex []*big.Int, claimer []common.Address, receiver []common.Address) (*DropERC721TokensClaimedIterator, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "TokensClaimed", claimConditionIndexRule, claimerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721TokensClaimedIterator{contract: _DropERC721.contract, event: "TokensClaimed", logs: logs, sub: sub}, nil
}

// WatchTokensClaimed is a free log subscription operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_DropERC721 *DropERC721Filterer) WatchTokensClaimed(opts *bind.WatchOpts, sink chan<- *DropERC721TokensClaimed, claimConditionIndex []*big.Int, claimer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var claimConditionIndexRule []interface{}
	for _, claimConditionIndexItem := range claimConditionIndex {
		claimConditionIndexRule = append(claimConditionIndexRule, claimConditionIndexItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "TokensClaimed", claimConditionIndexRule, claimerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721TokensClaimed)
				if err := _DropERC721.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
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

// ParseTokensClaimed is a log parse operation binding the contract event 0xfa76a4010d9533e3e964f2930a65fb6042a12fa6ff5b08281837a10b0be7321e.
//
// Solidity: event TokensClaimed(uint256 indexed claimConditionIndex, address indexed claimer, address indexed receiver, uint256 startTokenId, uint256 quantityClaimed)
func (_DropERC721 *DropERC721Filterer) ParseTokensClaimed(log types.Log) (*DropERC721TokensClaimed, error) {
	event := new(DropERC721TokensClaimed)
	if err := _DropERC721.contract.UnpackLog(event, "TokensClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721TokensLazyMintedIterator is returned from FilterTokensLazyMinted and is used to iterate over the raw logs and unpacked data for TokensLazyMinted events raised by the DropERC721 contract.
type DropERC721TokensLazyMintedIterator struct {
	Event *DropERC721TokensLazyMinted // Event containing the contract specifics and raw log

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
func (it *DropERC721TokensLazyMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721TokensLazyMinted)
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
		it.Event = new(DropERC721TokensLazyMinted)
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
func (it *DropERC721TokensLazyMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721TokensLazyMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721TokensLazyMinted represents a TokensLazyMinted event raised by the DropERC721 contract.
type DropERC721TokensLazyMinted struct {
	StartTokenId     *big.Int
	EndTokenId       *big.Int
	BaseURI          string
	EncryptedBaseURI []byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokensLazyMinted is a free log retrieval operation binding the contract event 0x2a0365091ef1a40953c670dce28177e37520648a6fdc91506bffac0ab045570d.
//
// Solidity: event TokensLazyMinted(uint256 indexed startTokenId, uint256 endTokenId, string baseURI, bytes encryptedBaseURI)
func (_DropERC721 *DropERC721Filterer) FilterTokensLazyMinted(opts *bind.FilterOpts, startTokenId []*big.Int) (*DropERC721TokensLazyMintedIterator, error) {

	var startTokenIdRule []interface{}
	for _, startTokenIdItem := range startTokenId {
		startTokenIdRule = append(startTokenIdRule, startTokenIdItem)
	}

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "TokensLazyMinted", startTokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721TokensLazyMintedIterator{contract: _DropERC721.contract, event: "TokensLazyMinted", logs: logs, sub: sub}, nil
}

// WatchTokensLazyMinted is a free log subscription operation binding the contract event 0x2a0365091ef1a40953c670dce28177e37520648a6fdc91506bffac0ab045570d.
//
// Solidity: event TokensLazyMinted(uint256 indexed startTokenId, uint256 endTokenId, string baseURI, bytes encryptedBaseURI)
func (_DropERC721 *DropERC721Filterer) WatchTokensLazyMinted(opts *bind.WatchOpts, sink chan<- *DropERC721TokensLazyMinted, startTokenId []*big.Int) (event.Subscription, error) {

	var startTokenIdRule []interface{}
	for _, startTokenIdItem := range startTokenId {
		startTokenIdRule = append(startTokenIdRule, startTokenIdItem)
	}

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "TokensLazyMinted", startTokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721TokensLazyMinted)
				if err := _DropERC721.contract.UnpackLog(event, "TokensLazyMinted", log); err != nil {
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

// ParseTokensLazyMinted is a log parse operation binding the contract event 0x2a0365091ef1a40953c670dce28177e37520648a6fdc91506bffac0ab045570d.
//
// Solidity: event TokensLazyMinted(uint256 indexed startTokenId, uint256 endTokenId, string baseURI, bytes encryptedBaseURI)
func (_DropERC721 *DropERC721Filterer) ParseTokensLazyMinted(log types.Log) (*DropERC721TokensLazyMinted, error) {
	event := new(DropERC721TokensLazyMinted)
	if err := _DropERC721.contract.UnpackLog(event, "TokensLazyMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DropERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DropERC721 contract.
type DropERC721TransferIterator struct {
	Event *DropERC721Transfer // Event containing the contract specifics and raw log

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
func (it *DropERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DropERC721Transfer)
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
		it.Event = new(DropERC721Transfer)
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
func (it *DropERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DropERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DropERC721Transfer represents a Transfer event raised by the DropERC721 contract.
type DropERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DropERC721 *DropERC721Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DropERC721TransferIterator, error) {

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

	logs, sub, err := _DropERC721.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DropERC721TransferIterator{contract: _DropERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DropERC721 *DropERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DropERC721Transfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _DropERC721.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DropERC721Transfer)
				if err := _DropERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_DropERC721 *DropERC721Filterer) ParseTransfer(log types.Log) (*DropERC721Transfer, error) {
	event := new(DropERC721Transfer)
	if err := _DropERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
