
## Edition Drop

This interface is currently accessible from the Edition Drop contract contract type via the ClaimConditions property.

```go
type EditionDropClaimConditions struct {}
```

### func \(\*EditionDropClaimConditions\) [GetActive](<https://github.com/ricebin/go-sdk/blob/main/thirdweb/edition_drop_claim_conditions.go#L52>)

```go
func (claim *EditionDropClaimConditions) GetActive(ctx context.Context, tokenId int) (*ClaimConditionOutput, error)
```

#### Get the currently active claim condition for a given token

tokenId: the token ID of the token to get the active claim condition for

returns: the currently active claim condition metadata

#### Example

```
tokenId := 0
condition, err := contract.ClaimConditions.GetActive(context.Background(), tokenId)

// Now you have access to all the claim condition metadata
fmt.Println("Start Time:", condition.StartTime)
fmt.Println("Available:", condition.AvailableSupply)
fmt.Println("Quantity:", condition.MaxQuantity)
fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
fmt.Println("Price:", condition.Price)
fmt.Println("Wait In Seconds", condition.WaitInSeconds)
```

### func \(\*EditionDropClaimConditions\) [GetAll](<https://github.com/ricebin/go-sdk/blob/main/thirdweb/edition_drop_claim_conditions.go#L96>)

```go
func (claim *EditionDropClaimConditions) GetAll(ctx context.Context, tokenId int) ([]*ClaimConditionOutput, error)
```

#### Get all claim conditions on this contract for a given token

tokenId: the token ID of the token to get the claim conditions for

returns: the metadata for all the claim conditions on this contract

#### Example

```
tokenId := 0
conditions, err := contract.ClaimConditions.GetAll(context.Background(), tokenId)

// Now you have access to all the claim condition metadata
condition := conditions[0]
fmt.Println("Start Time:", condition.StartTime)
fmt.Println("Available:", condition.AvailableSupply)
fmt.Println("Quantity:", condition.MaxQuantity)
fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
fmt.Println("Price:", condition.Price)
fmt.Println("Wait In Seconds", condition.WaitInSeconds)
```
