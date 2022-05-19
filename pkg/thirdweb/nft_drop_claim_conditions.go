package thirdweb

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type nftDropClaimConditions struct {
	abi     *abi.DropERC721
	helper  *contractHelper
	storage storage
}

func newNFTDropClaimConditions(address common.Address, provider *ethclient.Client, helper *contractHelper, storage storage) (*nftDropClaimConditions, error) {
	if contractAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		claimConditions := &nftDropClaimConditions{
			contractAbi,
			helper,
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
