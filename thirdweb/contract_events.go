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
)

type ContractEvents struct {
	abi      *abi.ABI
	contract *bind.BoundContract
	helper   *contractHelper
}

type EventQueryOptions struct {
	FromBlock uint64
	ToBlock   *uint64
	Filters   map[string]interface{}
}

type ContractEvent struct {
	EventName   string
	Data        map[string]interface{}
	Transaction types.Log
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

func (events *ContractEvents) AddEventListener(eventName string, listener func (event ContractEvent)) (func (), error) {
	logs, subscription, err := events.contract.WatchLogs(&bind.WatchOpts{}, eventName)
	if err != nil {
		fmt.Println("Error in watch logs")
		return func () {}, err
	}

	removeListener := func () {
		subscription.Unsubscribe()
		close(logs)
	}

	go func() {
		for log := range logs {
			event, err := events.transformEvent(eventName, log)
			if err != nil {
				return
			}

			listener(event)
		}
	}()

	return removeListener, nil
}

func (events *ContractEvents) GetEvents(ctx context.Context, eventName string, options EventQueryOptions) ([]ContractEvent, error) {
	eventSignature, ok := events.abi.Events[eventName]
	if !ok {
		return nil, fmt.Errorf("Event with name '%s' not found", eventName)
	}

	query := [][]interface{}{{eventSignature.ID}}
	if options.Filters != nil {
		// args := []interface{}{}
		for _, input := range eventSignature.Inputs {
			// For all indexed inputs, check if a filter is provided
			if value, ok := options.Filters[input.Name]; ok && input.Indexed {
				// If a filter is provided, check if the type is correct
				if reflect.TypeOf(value) != input.Type.GetType() {
					return nil, fmt.Errorf(
						"Filter for indexed input '%s' is of wrong type, expected type '%s'", 
						input.Name, 
						input.Type.GetType().String(),
					)
				}

				// If the type is correct, add the filter to the query
				query = append(query, []interface{}{value})
			} else {
				// If no filter is provided, add an empty value
				query = append(query, []interface{}{})
			}
		}
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

	parsedLogs := []ContractEvent{}
	for _, log := range logs {
		event, err := events.transformEvent(eventName, log)
		if err != nil {
			return nil, err
		}

		parsedLogs = append(parsedLogs, event)
	}

	return parsedLogs, nil
}

func (events *ContractEvents) transformEvent(eventName string, log types.Log) (ContractEvent, error) {
	parsedLog := map[string]interface{}{}
	events.contract.UnpackLogIntoMap(parsedLog, eventName, log)
	event := ContractEvent{ eventName, parsedLog, log }

	return event, nil
}