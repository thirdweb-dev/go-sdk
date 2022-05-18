package thirdweb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type NFTDropClaimConditions struct {
	contractWrapper *ContractWrapper[*abi.DropERC721]
	storage         Storage
}

func NewNFTDropClaimConditions(contractWrapper *ContractWrapper[*abi.DropERC721], storage Storage) *NFTDropClaimConditions {
	claimConditions := &NFTDropClaimConditions{
		contractWrapper,
		storage,
	}
	return claimConditions
}

func (claim *NFTDropClaimConditions) GetActive() (*ClaimConditionOutput, error) {
	id, err := claim.contractWrapper.abi.GetActiveClaimConditionId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	mc, err := claim.contractWrapper.abi.GetClaimConditionById(&bind.CallOpts{}, id)
	if err != nil {
		return nil, err
	}

	provider := claim.contractWrapper.GetProvider()
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
