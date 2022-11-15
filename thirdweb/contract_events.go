package thirdweb

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

type ContractEvents struct {
	abi      *abi.ABI
	contract *bind.BoundContract
	helper   *contractHelper
}

type EventQueryOptions struct {
	FromBlock uint64
	ToBlock   *uint64
	Filters   *map[string]interface{}
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

func (events *ContractEvents) GetEvents(ctx context.Context, eventName string, options EventQueryOptions) ([]map[string]interface{}, error) {
	eventSignature, ok := events.abi.Events[eventName]
	if !ok {
		return nil, fmt.Errorf("Event with name '%s' not found", eventName)
	}



	query := [][]interface{}{{eventSignature.ID}}
	if options.Filters != nil {
		args := []interface{}{}
		for _, input := range eventSignature.Inputs {
			// For all indexed inputs, check if a filter is provided
			if value, ok := (*options.Filters)[input.Name]; ok && input.Indexed {
				// If a filter is provided, check if the type is correct
				if reflect.TypeOf(value) != input.Type.GetType() {
					return nil, fmt.Errorf(
						"Filter for indexed input '%s' is of wrong type, expected type '%s'", 
						input.Name, 
						input.Type.GetType().String(),
					)
				}

				// If the type is correct, add the filter to the query
				args = append(args, value)
			}
		}

		query = append(query, args)
	}
	
	topics, err := abi.MakeTopics(query...)
	if err != nil {
		return nil, err
	}

	config := ethereum.FilterQuery{
		Addresses: []common.Address{events.helper.getAddress()},
		Topics:    topics,
		FromBlock: new(big.Int).SetUint64(options.FromBlock),
	}
	if options.ToBlock != nil {
		config.ToBlock = new(big.Int).SetUint64(*options.ToBlock)
	}

	logs, err := events.helper.GetProvider().FilterLogs(ctx, config)
	if err != nil {
		return nil, err
	}

	parsedLogs := []map[string]interface{}{}
	for _, log := range logs {
		event := map[string]interface{}{}
		events.contract.UnpackLogIntoMap(event, eventName, log)
	}

	return parsedLogs, nil
}