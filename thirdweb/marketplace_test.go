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
	assert.Equal(t, listing.AssetContractAddress, nft.helper.getAddress().Hex())
}

func TestListEdition(t *testing.T) {
	marketplace := getMarketplace()
	edition := getMarketplaceEdition()

	listingId, err := marketplace.CreateListing(&NewDirectListing{
		AssetContractAddress:     edition.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 2,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)

	assert.Equal(t, listingId, 0)
	listing, err := marketplace.GetListing(listingId)
	assert.Nil(t, err)
	assert.Equal(t, listing.Quantity, 2)
	assert.Equal(t, listing.AssetContractAddress, edition.helper.getAddress().Hex())
}

func TestListToken(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketpaceNft()
	token := getMarketplaceToken()

	listingId, err := marketplace.CreateListing(&NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)

	assert.Equal(t, listingId, 0)

	listing, err := marketplace.GetListing(listingId)
	assert.Nil(t, err)
	assert.Equal(t, listing.CurrencyContractAddress, token.helper.getAddress().Hex())
}

func TestBuyoutListing(t *testing.T) {
	marketplace := getMarketplace()
	token := getMarketplaceToken()
	edition := getMarketplaceEdition()

	adminBalance, _ := edition.BalanceOf(adminWallet, 0)
	secondaryBalance, _ := edition.BalanceOf(secondaryWallet, 0)
	assert.Equal(t, adminBalance, 100)
	assert.Equal(t, secondaryBalance, 0)

	adminTokens, _ := token.BalanceOf(adminWallet)
	secondaryTokens, _ := token.BalanceOf(secondaryWallet)
	assert.Equal(t, adminTokens.DisplayValue, 100)
	assert.Equal(t, secondaryTokens.DisplayValue, 100)

	listingId, err := marketplace.CreateListing(&NewDirectListing{
		AssetContractAddress:     edition.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 10,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      1.0,
	})

	assert.Nil(t, err)

	marketplace.helper.UpdatePrivateKey(secondaryPrivateKey)
	_, err = marketplace.BuyoutListing(listingId, 10)
	assert.Nil(t, err)

	adminBalance, _ = edition.BalanceOf(adminWallet, 0)
	secondaryBalance, _ = edition.BalanceOf(secondaryWallet, 0)
	assert.Equal(t, adminBalance, 90)
	assert.Equal(t, secondaryBalance, 10)

	adminTokens, _ = token.BalanceOf(adminWallet)
	secondaryTokens, _ = token.BalanceOf(secondaryWallet)
	assert.Equal(t, adminTokens.DisplayValue, 110)
	assert.Equal(t, secondaryTokens.DisplayValue, 90)
}
