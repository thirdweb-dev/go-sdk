package thirdweb

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
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
//	condition, err := contract.ClaimConditions.GetActive()
//
//	// Now you have access to all the claim condition metadata
//	fmt.Println("Start Time:", condition.StartTime)
//	fmt.Println("Available:", condition.AvailableSupply)
//	fmt.Println("Quantity:", condition.MaxQuantity)
//	fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
//	fmt.Println("Price:", condition.Price)
//	fmt.Println("Wait In Seconds", condition.WaitInSeconds)
func (claim *NFTDropClaimConditions) GetActive(ctx context.Context) (*ClaimConditionOutput, error) {
	id, err := claim.abi.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	active, err := claim.abi.GetClaimConditionById(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}

	provider := claim.helper.GetProvider()
	claimCondition, err := transformResultToClaimCondition(
		ctx,
		&active,
		provider,
	)
	if err != nil {
		return nil, err
	}

	return claimCondition, nil
}

func (claim *NFTDropClaimConditions) Get(ctx context.Context, claimConditionId int) (*ClaimConditionOutput, error) {
	condition, err := claim.abi.GetClaimConditionById(&bind.CallOpts{Context: ctx}, big.NewInt(int64(claimConditionId)))
	if err != nil {
		return nil, err
	}

	provider := claim.helper.GetProvider()

	claimCondition, err := transformResultToClaimCondition(
		ctx,
		&condition,
		provider,
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
//	conditions, err := contract.ClaimConditions.GetAll()
//
//	// Now you have access to all the claim condition metadata
//	condition := conditions[0]
//	fmt.Println("Start Time:", condition.StartTime)
//	fmt.Println("Available:", condition.AvailableSupply)
//	fmt.Println("Quantity:", condition.MaxQuantity)
//	fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
//	fmt.Println("Price:", condition.Price)
//	fmt.Println("Wait In Seconds", condition.WaitInSeconds)
func (claim *NFTDropClaimConditions) GetAll(ctx context.Context) ([]*ClaimConditionOutput, error) {
	condition, err := claim.abi.ClaimCondition(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	provider := claim.helper.GetProvider()
	startId := condition.CurrentStartId.Int64()
	count := condition.Count.Int64()

	conditions := []*ClaimConditionOutput{}
	for i := startId; i < count; i++ {
		mc, err := claim.abi.GetClaimConditionById(&bind.CallOpts{Context: ctx}, big.NewInt(i))
		if err != nil {
			return nil, err
		}

		claimCondition, err := transformResultToClaimCondition(
			ctx,
			&mc,
			provider,
		)
		if err != nil {
			return nil, err
		}

		conditions = append(conditions, claimCondition)
	}

	return conditions, nil
}

func (claim *NFTDropClaimConditions) getMerkleMetadata(ctx context.Context) (*map[string]string, error) {
	uri, err := claim.abi.InternalContractURI(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	body, err := claim.storage.Get(ctx, uri)
	if err != nil {
		return nil, err
	}

	var rawMetadata struct {
		Merkle map[string]string `json:"merkle"`
	}
	if err := json.Unmarshal(body, &rawMetadata); err != nil {
		return nil, err
	}

	return &rawMetadata.Merkle, nil
}

func (claim *NFTDropClaimConditions) GetClaimerProofs(
	ctx context.Context,
	claimerAddress string,
) (*SnapshotEntryWithProof, error) {
	claimCondition, err := claim.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	if !strings.HasPrefix(hex.EncodeToString(claimCondition.MerkleRootHash[:]), zeroAddress) {
		merkleMetadata, err := claim.getMerkleMetadata(ctx)
		if err != nil {
			return nil, err
		}

		return fetchSnapshotEntryForAddress(
			ctx,
			common.HexToAddress(claimerAddress),
			claimCondition.MerkleRootHash,
			merkleMetadata,
			claim.helper.GetProvider(),
			claim.storage,
		)
	} else {
		return nil, nil
	}
}

/**
func (claim *NFTDropClaimConditions) CanClaim(
	quantity int,
	addressToCheck string,
) (bool, error) {
	reasons, err := claim.GetClaimIneligibilityReasons(quantity, addressToCheck)
	if err != nil {
		return false, err
	}

	if len(reasons) > 0 {
		return false, nil
	} else {
		return true, nil
	}
}

func (claim *NFTDropClaimConditions) GetClaimIneligibilityReasons(
	quantity int,
	addressToCheck string,
) ([]*ClaimEligibility, error) {
	return nil, nil
}

func (claim *NFTDropClaimConditions) Set(
	claimConditionInputs []*ClaimConditionInput,
	resetClaimEligibilityForAll bool,
) (*types.Transaction, error) {
	return nil, nil
}
**/
