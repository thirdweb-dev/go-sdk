package thirdweb

import (
	"context"
	"fmt"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// This interface provides a way to query past events or listen for future events on any contract.
// It's currently support on all pre-built and custom contracts!
//
// Example
//
//	// First get an instance of your contract
//	contract, _ := sdk.GetContract("0x...");
//
//	// Now, get all Transfer events from a specific block range
//	contract.Events.GetEvents("Transfer", thirdweb.EventQueryOptions{
//	  FromBlock: 100000000,
//	  ToBlock:   100000001,
//	})
//
//	// And setup a listener to listen for future Transfer events
//	contract.Events.AddEventListener("Transfer", func (event thirdweb.ContractEvent) {
//	  fmt.Printf("%#v\n", event)
//	})
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

type EventSubscription struct {
	Err         func() <-chan error
	Unsubscribe func()
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

// Add a listener to listen in the background for any future events of a specific type.
//
// eventName: The name of the event to listen for
//
// listener: The listener function that will be called whenever a new event is received
//
// returns: An EventSubscription object that can be used to unsubscribe from the event or check for errors
//
// Example
//
//	// Define a listener function to be called whenever a new Transfer event is received
//	listener := func (event thirdweb.ContractEvent) {
//	  fmt.Printf("%#v\n", event)
//	}
//
//	// Add a new listener for the Transfer event
//	subscription := contract.Events.AddEventListener(context.Background(), "Transfer", listener)
//
//	// Unsubscribe from the Transfer event at some time in the future, closing the listener
//	subscription.Unsubscribe()
func (events *ContractEvents) AddEventListener(ctx context.Context, eventName string, listener func(event ContractEvent)) EventSubscription {
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)
	errors := make(chan error)

	go func() {
		nextBlockToCheck, err := events.helper.GetProvider().BlockNumber(ctx)
		if err != nil {
			errors <- err
			return
		}

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				currentBlockNumber, err := events.helper.GetProvider().BlockNumber(ctx)
				if err != nil {
					errors <- err
					continue
				}

				if currentBlockNumber >= nextBlockToCheck {
					recentEvents, err := events.GetEvents(ctx, eventName, EventQueryOptions{
						FromBlock: nextBlockToCheck,
						ToBlock:   &currentBlockNumber,
					})
					if err != nil {
						errors <- err
						continue
					}

					for _, event := range recentEvents {
						listener(event)
					}

					nextBlockToCheck = currentBlockNumber + 1
				}
			}
		}
	}()

	subscription := EventSubscription{
		Err: func() <-chan error {
			return errors
		},
		Unsubscribe: func() {
			done <- true
		},
	}

	return subscription
}

// Query past events of a specific type on the contract.
//
// eventName: The name of the event to query for
//
// options: The options to use when querying for events, including block range specifications and filters
//
// returns: a list of ContractEvent objects that match the query
//
// Example
//
//	// First we define a filter to only get Transfer events where the "from" address is "0x..."
//	// Note that you can only add filters for indexed parameters on events
//	filters := map[string]interface{} {
//	  "from": common.HexToAddress("0x...")
//	}
//
//	// Now we can define the query options, including the block range and the filter
//	queryOptions := thirdweb.EventQueryOptions{
//	  FromBlock: 100000000, // Defaults to block 0 if you don't specify this field
//	  ToBlock:   100000001, // Defaults to latest block if you don't specify this field
//	  Filters:   filters,
//	}
//
//	// Now we can query for the Transfer events
//	events, _ := contract.Events.GetEvents("Transfer", queryOptions)
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
			if input.Indexed {
				if value, ok := options.Filters[input.Name]; ok {
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
	if err := events.contract.UnpackLogIntoMap(parsedLog, eventName, log); err != nil {
		return ContractEvent{}, err
	}
	event := ContractEvent{eventName, parsedLog, log}

	return event, nil
}
