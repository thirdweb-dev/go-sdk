package thirdweb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getToken() *Token {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployToken(context.Background(), &DeployTokenMetadata{
		Name: "Token",
	})
	token, _ := sdk.GetToken(address)

	return token
}

func TestMintToken(t *testing.T) {
	token := getToken()

	balance, _ := token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)

	_, err := token.Mint(context.Background(), 0.1)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(0.1), balance.DisplayValue)
}

func TestBatchMintToken(t *testing.T) {
	token := getToken()

	balance, _ := token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)

	_, err := token.MintBatchTo(
		context.Background(),
		[]*TokenAmount{
			{
				ToAddress: token.helper.GetSignerAddress().String(),
				Amount:    1.5,
			},
			{
				ToAddress: secondaryWallet,
				Amount:    2.5,
			},
		},
	)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(1.5), balance.DisplayValue)

	supply, _ := token.TotalSupply()
	assert.Equal(t, float64(4), supply.DisplayValue)
}

func TestBurnToken(t *testing.T) {
	token := getToken()

	token.Mint(context.Background(), 10)

	balance, _ := token.Balance()
	assert.Equal(t, float64(10), balance.DisplayValue)

	_, err := token.Burn(context.Background(), 10)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)
}

func TestTransferToken(t *testing.T) {
	token := getToken()

	token.Mint(context.Background(), 10)

	balance, _ := token.Balance()
	assert.Equal(t, float64(10), balance.DisplayValue)

	_, err := token.Transfer(context.Background(), secondaryWallet, 10)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)
}
