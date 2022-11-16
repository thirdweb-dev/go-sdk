
## Contract Events

This interface provides a way to query past events or listen for future events on any contract\. It's currently support on all pre\-built and custom contracts\!

### Example

```
// First get an instance of your contract
contract, _ := sdk.GetContract("0x...");

// Now, get all Transfer events from a specific block range
contract.Events.GetEvents("Transfer", thirdweb.EventQueryOptions{
  FromBlock: 100000000,
  ToBlock:   100000001,
})

// And setup a listener to listen for future Transfer events
contract.Events.AddEventListener("Transfer", func (event thirdweb.ContractEvent) {
  fmt.Printf("%#v\n", event)
})
```

```go
type ContractEvents struct {}
```

### func \(\*ContractEvents\) [AddEventListener](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_events.go#L95>)

```go
func (events *ContractEvents) AddEventListener(ctx context.Context, eventName string, listener func(event ContractEvent)) EventSubscription
```

Add a listener to listen in the background for any future events of a specific type\.

eventName: The name of the event to listen for

listener: The listener function that will be called whenever a new event is received

returns: An EventSubscription object that can be used to unsubscribe from the event or check for errors

#### Example

```
// Define a listener function to be called whenever a new Transfer event is received
listener := func (event thirdweb.ContractEvent) {
  fmt.Printf("%#v\n", event)
}

// Add a new listener for the Transfer event
subscription := contract.Events.AddEventListener(context.Background(), "Transfer", listener)

// Unsubscribe from the Transfer event at some time in the future, closing the listener
subscription.Unsubscribe()
```

### func \(\*ContractEvents\) [GetEvents](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/contract_events.go#L175>)

```go
func (events *ContractEvents) GetEvents(ctx context.Context, eventName string, options EventQueryOptions) ([]ContractEvent, error)
```

Query past events of a specific type on the contract\.

eventName: The name of the event to query for

options: The options to use when querying for events, including block range specifications and filters

returns: a list of ContractEvent objects that match the query

#### Example

```
// First we define a filter to only get Transfer events where the "from" address is "0x..."
// Note that you can only add filters for indexed parameters on events
filters := map[string]interface{} {
  "from": common.HexToAddress("0x...")
}

// Now we can define the query options, including the block range and the filter
queryOptions := thirdweb.EventQueryOptions{
  FromBlock: 100000000, // Defaults to block 0 if you don't specify this field
  ToBlock:   100000001, // Defaults to latest block if you don't specify this field
  Filters:   filters,
}

// Now we can query for the Transfer events
events, _ := contract.Events.GetEvents("Transfer", queryOptions)
```
