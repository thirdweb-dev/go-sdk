package thirdweb

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type ContractEvents struct {
	abi      *abi.ABI
	contract *bind.BoundContract
	helper   *contractHelper
}

func newContractEvents(contractAbi string, helper *contractHelper) (*ContractEvents, error) {
	parsedAbi, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	contract := bind.NewBoundContract(helper.getAddress(), parsedAbi, helper.GetProvider(), helper.GetProvider(), helper.GetProvider())

	return &ContractEvents{
		abi:      &parsedAbi,
		contract: contract,
		helper:   helper,
	}, nil
}

func (events *ContractEvents) AddEventListener(eventName string) (chan types.Log, event.Subscription, error) {
	return events.contract.WatchLogs(&bind.WatchOpts{}, eventName)
}

func (events *ContractEvents) GetEvents(eventName string) (chan types.Log, event.Subscription, error) {
	return events.contract.FilterLogs(&bind.FilterOpts{Start: 0}, eventName)
}