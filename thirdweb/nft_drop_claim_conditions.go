package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/abi"
)

// This interface is currently accessible from the NFT Drop contract contract type
// via the ClaimConditions property.
type NFTDropClaimConditions struct {
	abi     *abi.DropERC721
	helper  *contractHelper
	storage storage
}

func newNFTDropClaimConditions(address common.Address, provider *ethclient.Client, helper *contractHelper, storage storage) (*NFTDropClaimConditions, error) {
	if contractAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		claimConditions := &NFTDropClaimConditions{
			contractAbi,
			helper,
			storage,
		}
		return claimConditions, err
	}
}

// Get the currently active claim condition
//
// returns: the currently active claim condition metadata
//
// Example
//
// 	condition, err := contract.ClaimConditions.GetActive()
//
// 	// Now you have access to all the claim condition metadata
// 	fmt.Println("Start Time:", condition.StartTime)
// 	fmt.Println("Available:", condition.AvailableSupply)
// 	fmt.Println("Quantity:", condition.MaxQuantity)
// 	fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
// 	fmt.Println("Price:", condition.Price)
// 	fmt.Println("Wait In Seconds", condition.WaitInSeconds)
func (claim *NFTDropClaimConditions) GetActive() (*ClaimConditionOutput, error) {
	id, err := claim.abi.GetActiveClaimConditionId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	mc, err := claim.abi.GetClaimConditionById(&bind.CallOpts{}, id)
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

// Get all claim conditions on this contract
//
// returns: the metadata for all the claim conditions on this contract
//
// Example
//
// 	conditions, err := contract.ClaimConditions.GetAll()
//
// 	// Now you have access to all the claim condition metadata
// 	condition := conditions[0]
// 	fmt.Println("Start Time:", condition.StartTime)
// 	fmt.Println("Available:", condition.AvailableSupply)
// 	fmt.Println("Quantity:", condition.MaxQuantity)
// 	fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
// 	fmt.Println("Price:", condition.Price)
// 	fmt.Println("Wait In Seconds", condition.WaitInSeconds)
func (claim *NFTDropClaimConditions) GetAll() ([]*ClaimConditionOutput, error) {
	condition, err := claim.abi.ClaimCondition(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	provider := claim.helper.GetProvider()
	startId := condition.CurrentStartId.Int64()
	count := condition.Count.Int64()

	conditions := []*ClaimConditionOutput{}
	for i := startId; i < count; i++ {
		mc, err := claim.abi.GetClaimConditionById(&bind.CallOpts{}, big.NewInt(i))
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
