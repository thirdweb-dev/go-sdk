
## NFT Drop Claim Conditions

This interface is currently accessible from the NFT Drop contract contract type via the ClaimConditions property.

```go
type NFTDropClaimConditions struct {}
```

### func \(\*NFTDropClaimConditions\) [GetActive](<https://github.com/ricebin/go-sdk/blob/main/thirdweb/nft_drop_claim_conditions.go#L49>)

```go
func (claim *NFTDropClaimConditions) GetActive() (*ClaimConditionOutput, error)
```

#### Get the currently active claim condition

returns: the currently active claim condition metadata

#### Example

```
condition, err := contract.ClaimConditions.GetActive()

// Now you have access to all the claim condition metadata
fmt.Println("Start Time:", condition.StartTime)
fmt.Println("Available:", condition.AvailableSupply)
fmt.Println("Quantity:", condition.MaxQuantity)
fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
fmt.Println("Price:", condition.Price)
fmt.Println("Wait In Seconds", condition.WaitInSeconds)
```

### func \(\*NFTDropClaimConditions\) [GetAll](<https://github.com/ricebin/go-sdk/blob/main/thirdweb/nft_drop_claim_conditions.go#L90>)

```go
func (claim *NFTDropClaimConditions) GetAll() ([]*ClaimConditionOutput, error)
```

#### Get all claim conditions on this contract

returns: the metadata for all the claim conditions on this contract

#### Example

```
conditions, err := contract.ClaimConditions.GetAll()

// Now you have access to all the claim condition metadata
condition := conditions[0]
fmt.Println("Start Time:", condition.StartTime)
fmt.Println("Available:", condition.AvailableSupply)
fmt.Println("Quantity:", condition.MaxQuantity)
fmt.Println("Quantity Limit:", condition.QuantityLimitPerTransaction)
fmt.Println("Price:", condition.Price)
fmt.Println("Wait In Seconds", condition.WaitInSeconds)
```
