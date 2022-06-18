package thirdweb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getEdition() *Edition {
	sdk := getSDK()
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

func TestBatchMintEdition(t *testing.T) {
	edition := getEdition()

	balance, _ := edition.Balance(0)
	assert.Equal(t, 0, balance)

	_, err := edition.MintBatchTo(
		edition.helper.GetSignerAddress().String(),
		[]*EditionMetadataInput{
			{
				Metadata: &NFTMetadataInput{
					Name: "NFT 1",
				},
				Supply: 1,
			},
			{
				Metadata: &NFTMetadataInput{
					Name: "NFT 2",
				},
				Supply: 2,
			},
		},
	)
	assert.Nil(t, err)

	balance, _ = edition.Balance(0)
	assert.Equal(t, 1, balance)

	balance, _ = edition.Balance(1)
	assert.Equal(t, 2, balance)

	nfts, _ := edition.GetAll()
	assert.Equal(t, 2, len(nfts))
	assert.Equal(t, "NFT 1", nfts[0].Metadata.Name)
	assert.Equal(t, "NFT 2", nfts[1].Metadata.Name)
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

func TestSignatureMint(t *testing.T) {
	edition := getEdition()

	payload, err := edition.Signature.Generate(
		&Signature1155PayloadInput{
			To:                   edition.helper.GetSignerAddress().String(),
			Price:                0,
			CurrencyAddress:      "0x0000000000000000000000000000000000000000",
			MintStartTime:        0,
			MintEndTime:          100000000000000,
			PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
			Metadata: &NFTMetadataInput{
				Name: "Sigmint",
			},
			RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
			RoyaltyBps:       0,
			Quantity:         1,
		},
	)
	assert.Nil(t, err)

	valid, err := edition.Signature.Verify(payload)
	assert.Nil(t, err)
	assert.True(t, valid)

	_, err = edition.Signature.Mint(payload)
	assert.Nil(t, err)

	balance, _ := edition.Balance(0)
	assert.Equal(t, 1, balance)
}
