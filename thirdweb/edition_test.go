package thirdweb

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getEdition() *Edition {
	sdk := getSDK()
	fmt.Println("Deploying edition...")
	address, _ := sdk.Deployer.DeployEdition(&DeployEditionMetadata{
		Name: "Edition",
	})
	edition, _ := sdk.GetEdition(address)

	return edition
}

func TestMintEdition(t *testing.T) {
	edition := getEdition()

	balance, _ := edition.Balance(0)
	assert.Equal(t, 0, balance)

	_, err := edition.Mint(&EditionMetadataInput{
		Metadata: &NFTMetadataInput{
			Name: "NFT",
		},
		Supply: 10,
	})
	assert.Nil(t, err)

	balance, _ = edition.Balance(0)
	assert.Equal(t, 10, balance)
}

func TestBurnEdition(t *testing.T) {
	edition := getEdition()

	edition.Mint(&EditionMetadataInput{
		Metadata: &NFTMetadataInput{
			Name: "NFT",
		},
		Supply: 10,
	})

	balance, _ := edition.Balance(0)
	assert.Equal(t, 10, balance)

	_, err := edition.Burn(0, 10)
	assert.Nil(t, err)

	balance, _ = edition.Balance(0)
	assert.Equal(t, 0, balance)
}

func TestTransferEdition(t *testing.T) {
	edition := getEdition()

	edition.Mint(&EditionMetadataInput{
		Metadata: &NFTMetadataInput{
			Name: "NFT",
		},
		Supply: 10,
	})

	balance, _ := edition.Balance(0)
	assert.Equal(t, 10, balance)

	_, err := edition.Transfer(secondaryWallet, 0, 10)
	assert.Nil(t, err)

	balance, _ = edition.Balance(0)
	assert.Equal(t, 0, balance)
}
