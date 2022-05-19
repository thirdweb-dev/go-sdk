package thirdweb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type nftDropClaimConditions struct {
	abi             *abi.DropERC721
	contractWrapper *contractWrapper
	storage         storage
}

func newNFTDropClaimConditions(address common.Address, provider *ethclient.Client, contractWrapper *contractWrapper, storage storage) (*nftDropClaimConditions, error) {
	if abi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		claimConditions := &nftDropClaimConditions{
			abi,
			contractWrapper,
			storage,
		}
		return claimConditions, err
	}
}

func (claim *nftDropClaimConditions) GetActive() (*ClaimConditionOutput, error) {
	id, err := claim.abi.GetActiveClaimConditionId(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	mc, err := claim.abi.GetClaimConditionById(&bind.CallOpts{}, id)
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
