package thirdweb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getToken() *Token {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployToken(&DeployTokenMetadata{
		Name: "Token",
	})
	token, _ := sdk.GetToken(address)

	return token
}

func TestMintToken(t *testing.T) {
	token := getToken()

	balance, _ := token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)

	_, err := token.Mint(0.1)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(0.1), balance.DisplayValue)
}

func TestBurnToken(t *testing.T) {
	token := getToken()

	token.Mint(10)

	balance, _ := token.Balance()
	assert.Equal(t, float64(10), balance.DisplayValue)

	_, err := token.Burn(10)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)
}

func TestTransferToken(t *testing.T) {
	token := getToken()

	token.Mint(10)

	balance, _ := token.Balance()
	assert.Equal(t, float64(10), balance.DisplayValue)

	_, err := token.Transfer(secondaryWallet, 10)
	assert.Nil(t, err)

	balance, _ = token.Balance()
	assert.Equal(t, float64(0), balance.DisplayValue)
}
