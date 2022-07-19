package thirdweb

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getMarketpaceNft() *NFTCollection {
	nft := getNft()
	nft.MintBatch([]*NFTMetadataInput{
		{
			Name: "Test 1",
		},
		{
			Name: "Test 2",
		},
	})
	return nft
}

func getMarketplaceEdition() *Edition {
	edition := getEdition()
	edition.MintBatch([]*EditionMetadataInput{
		{
			Metadata: &NFTMetadataInput{
				Name: "Test 1",
			},
			Supply: 100,
		},
		{
			Metadata: &NFTMetadataInput{
				Name: "Test 2",
			},
			Supply: 100,
		},
	})
	return edition
}

func getMarketplaceToken() *Token {
	token := getToken()
	token.MintBatchTo([]*TokenAmount{
		{
			ToAddress: adminWallet,
			Amount:    100,
		},
		{
			ToAddress: secondaryWallet,
			Amount:    100,
		},
		{
			ToAddress: tertiaryWallet,
			Amount:    100,
		},
	})

	return token
}

func getMarketplace() *Marketplace {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployMarketplace(&DeployMarketplaceMetadata{
		Name: "Marketplace",
	})
	marketplace, _ := sdk.GetMarketplace(address)

	return marketplace
}

func TestListNft(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketpaceNft()

	listingId, err := marketplace.CreateListing(&NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)

	assert.Equal(t, listingId, 0)

	listing, err := marketplace.GetListing(listingId)
	assert.Nil(t, err)
	assert.Equal(t, listing.Quantity, 1)
}
