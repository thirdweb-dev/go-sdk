package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// This interface is currently accessible from the Edition Drop contract contract type
// via the ClaimConditions property.
type EditionDropClaimConditions struct {
	abi     *abi.DropERC1155
	helper  *contractHelper
	storage storage
}

func newEditionDropClaimConditions(address common.Address, provider *ethclient.Client, helper *contractHelper, storage storage) (*EditionDropClaimConditions, error) {
	if contractAbi, err := abi.NewDropERC1155(address, provider); err != nil {
		return nil, err
	} else {
		claimConditions := &EditionDropClaimConditions{
			contractAbi,
			helper,
			storage,
		}
		return claimConditions, err
	}
}

// Get the currently active claim condition for a given token
//
// tokenId: the token ID of the token to get the active claim condition for
//
// returns: the currently active claim condition metadata
//
// Example
//
//	tokenId := 0
//	condition, err := contract.ClaimConditions.GetActive(tokenId)
//
//	// Now you have access to all the claim condition metadata
//	fmt.Println("Start Time:", condition.StartTime)
//	fmt.Println("Available:", condition.AvailableSupply)
//	fmt.Println("Quantity:", condition.MaxQuantity)
//	fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
//	fmt.Println("Price:", condition.Price)
//	fmt.Println("Wait In Seconds", condition.WaitInSeconds)
func (claim *EditionDropClaimConditions) GetActive(tokenId int) (*ClaimConditionOutput, error) {
	id, err := claim.abi.GetActiveClaimConditionId(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	if err != nil {
		return nil, err
	}

	mc, err := claim.abi.GetClaimConditionById(&bind.CallOpts{}, big.NewInt(int64(tokenId)), id)
	if err != nil {
		return nil, err
	}

	provider := claim.helper.GetProvider()
	claimCondition, err := transformResultToClaimCondition(
		&mc,
		provider,
		claim.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimCondition, nil
}

// Get all claim conditions on this contract for a given token
//
// tokenId: the token ID of the token to get the claim conditions for
//
// returns: the metadata for all the claim conditions on this contract
//
// Example
//
//	tokenId := 0
//	conditions, err := contract.ClaimConditions.GetAll(tokenId)
//
//	// Now you have access to all the claim condition metadata
//	condition := conditions[0]
//	fmt.Println("Start Time:", condition.StartTime)
//	fmt.Println("Available:", condition.AvailableSupply)
//	fmt.Println("Quantity:", condition.MaxQuantity)
//	fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
//	fmt.Println("Price:", condition.Price)
//	fmt.Println("Wait In Seconds", condition.WaitInSeconds)
func (claim *EditionDropClaimConditions) GetAll(tokenId int) ([]*ClaimConditionOutput, error) {
	condition, err := claim.abi.ClaimCondition(&bind.CallOpts{}, big.NewInt(int64(tokenId)))
	if err != nil {
		return nil, err
	}

	provider := claim.helper.GetProvider()
	startId := condition.CurrentStartId.Int64()
	count := condition.Count.Int64()

	conditions := []*ClaimConditionOutput{}
	for i := startId; i < count; i++ {
		mc, err := claim.abi.GetClaimConditionById(&bind.CallOpts{}, big.NewInt(int64(tokenId)), big.NewInt(i))
		if err != nil {
			return nil, err
		}

		claimCondition, err := transformResultToClaimCondition(
			&mc,
			provider,
			claim.storage,
		)
		if err != nil {
			return nil, err
		}

		conditions = append(conditions, claimCondition)
	}

	return conditions, nil
}
