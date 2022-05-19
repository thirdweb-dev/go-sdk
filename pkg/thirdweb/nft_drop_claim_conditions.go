package thirdweb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type nftDropClaimConditions struct {
	contractWrapper *contractWrapper[*abi.DropERC721]
	storage         storage
}

func newNFTDropClaimConditions(contractWrapper *contractWrapper[*abi.DropERC721], storage storage) *nftDropClaimConditions {
	claimConditions := &nftDropClaimConditions{
		contractWrapper,
		storage,
	}
	return claimConditions
}

func (claim *nftDropClaimConditions) GetActive() (*ClaimConditionOutput, error) {
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
