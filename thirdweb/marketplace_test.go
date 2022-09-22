package thirdweb

import (
	"context"
	"encoding/hex"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getMarketplaceNft() *NFTCollection {
	nft := getNft()
	nft.MintBatch(context.Background(), []*NFTMetadataInput{
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
	edition.MintBatch(context.Background(), []*EditionMetadataInput{
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
	token.MintBatchTo(context.Background(), []*TokenAmount{
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
	address, _ := sdk.Deployer.DeployMarketplace(context.Background(), &DeployMarketplaceMetadata{
		Name: "Marketplace",
	})
	marketplace, _ := sdk.GetMarketplace(address)

	return marketplace
}

func TestListNft(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketplaceNft()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
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

	listing, err := marketplace.GetListing(context.Background(), listingId)
	assert.Nil(t, err)
	assert.Equal(t, listing.Quantity, 1)
	assert.Equal(t, listing.AssetContractAddress, nft.helper.getAddress().Hex())
}

func TestListEdition(t *testing.T) {
	marketplace := getMarketplace()
	edition := getMarketplaceEdition()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
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
	listing, err := marketplace.GetListing(context.Background(), listingId)
	assert.Nil(t, err)
	assert.Equal(t, listing.Quantity, 2)
	assert.Equal(t, listing.AssetContractAddress, edition.helper.getAddress().Hex())
}

func TestListToken(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketplaceNft()
	token := getMarketplaceToken()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
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

	listing, err := marketplace.GetListing(context.Background(), listingId)
	assert.Nil(t, err)
	assert.Equal(t, listing.CurrencyContractAddress, token.helper.getAddress().Hex())
}

func TestBuyoutListingNft(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketplaceNft()
	token := getMarketplaceToken()

	adminBalance, _ := nft.BalanceOf(context.Background(), adminWallet)
	secondaryBalance, _ := nft.BalanceOf(context.Background(), secondaryWallet)
	assert.Equal(t, adminBalance, 2)
	assert.Equal(t, secondaryBalance, 0)

	adminTokens, _ := token.BalanceOf(adminWallet)
	secondaryTokens, _ := token.BalanceOf(secondaryWallet)
	assert.Equal(t, adminTokens.DisplayValue, float64(100))
	assert.Equal(t, secondaryTokens.DisplayValue, float64(100))

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      10.0,
	})
	assert.Nil(t, err)

	marketplace.helper.UpdatePrivateKey(secondaryPrivateKey)
	_, err = marketplace.BuyoutListing(context.Background(), listingId, 1)
	assert.Nil(t, err)

	adminBalance, _ = nft.BalanceOf(context.Background(), adminWallet)
	secondaryBalance, _ = nft.BalanceOf(context.Background(), secondaryWallet)
	assert.Equal(t, adminBalance, 1)
	assert.Equal(t, secondaryBalance, 1)

	adminTokens, _ = token.BalanceOf(adminWallet)
	secondaryTokens, _ = token.BalanceOf(secondaryWallet)
	assert.Equal(t, adminTokens.DisplayValue, float64(110))
	assert.Equal(t, secondaryTokens.DisplayValue, float64(90))
}

func TestBuyoutListingEdition(t *testing.T) {
	marketplace := getMarketplace()
	token := getMarketplaceToken()
	edition := getMarketplaceEdition()

	adminBalance, _ := edition.BalanceOf(context.Background(), adminWallet, 0)
	secondaryBalance, _ := edition.BalanceOf(context.Background(), secondaryWallet, 0)
	assert.Equal(t, adminBalance, 100)
	assert.Equal(t, secondaryBalance, 0)

	adminTokens, _ := token.BalanceOf(adminWallet)
	secondaryTokens, _ := token.BalanceOf(secondaryWallet)
	assert.Equal(t, adminTokens.DisplayValue, float64(100))
	assert.Equal(t, secondaryTokens.DisplayValue, float64(100))

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
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
	_, err = marketplace.BuyoutListing(context.Background(), listingId, 10)
	assert.Nil(t, err)

	adminBalance, _ = edition.BalanceOf(context.Background(), adminWallet, 0)
	secondaryBalance, _ = edition.BalanceOf(context.Background(), secondaryWallet, 0)
	assert.Equal(t, adminBalance, 90)
	assert.Equal(t, secondaryBalance, 10)

	adminTokens, _ = token.BalanceOf(adminWallet)
	secondaryTokens, _ = token.BalanceOf(secondaryWallet)
	assert.Equal(t, adminTokens.DisplayValue, float64(110))
	assert.Equal(t, secondaryTokens.DisplayValue, float64(90))
}

func TestGetListingFilters(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketplaceNft()
	edition := getMarketplaceEdition()
	token := getMarketplaceToken()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      0,
	})
	assert.Nil(t, err)
	assert.Equal(t, listingId, 0)

	listingId, err = marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  1,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)
	assert.Equal(t, listingId, 1)

	listingId, err = marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     edition.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 10,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)
	assert.Equal(t, listingId, 2)

	listings, err := marketplace.GetAllListings(context.Background(), nil)
	assert.Nil(t, err)
	assert.Equal(t, len(listings), 3)

	listings, err = marketplace.GetAllListings(context.Background(), &MarketplaceFilter{
		TokenContract: nft.helper.getAddress().Hex(),
	})
	assert.Nil(t, err)
	assert.Equal(t, len(listings), 2)

	listings, err = marketplace.GetAllListings(context.Background(), &MarketplaceFilter{
		Seller: adminWallet,
	})
	assert.Nil(t, err)
	assert.Equal(t, len(listings), 3)

	listings, err = marketplace.GetAllListings(context.Background(), &MarketplaceFilter{
		Seller: secondaryWallet,
	})
	assert.Nil(t, err)
	assert.Equal(t, len(listings), 0)

	marketplace.helper.UpdatePrivateKey(secondaryPrivateKey)
	_, err = marketplace.BuyoutListing(context.Background(), 2, 10)
	assert.Nil(t, err)

	listings, err = marketplace.GetActiveListings(context.Background(), nil)
	assert.Nil(t, err)
	assert.Equal(t, len(listings), 2)
}

func TestCreateListingUnapprovedEncoder(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketplaceNft()

	_, err := marketplace.Encoder.CreateListing(context.Background(), adminWallet, &NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      1.0,
	})
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), marketplace.helper.getAddress().Hex())
	assert.Contains(t, err.Error(), nft.helper.getAddress().Hex())
	assert.Contains(t, err.Error(), adminWallet)
	assert.Contains(t, err.Error(), "marketplace.Encoder.ApproveCreateListing")
}

func TestBuyoutListingUnapprovedEncoder(t *testing.T) {
	marketplace := getMarketplace()
	token := getMarketplaceToken()
	edition := getMarketplaceEdition()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     edition.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 10,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)

	_, err = marketplace.Encoder.BuyoutListing(context.Background(), adminWallet, listingId, 10, adminWallet)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), marketplace.helper.getAddress().Hex())
	assert.Contains(t, err.Error(), token.helper.getAddress().Hex())
	assert.Contains(t, err.Error(), adminWallet)
	assert.Contains(t, err.Error(), "marketplace.Encoder.ApproveBuyoutListing")
}

func TestCreateListingApprovedEncoder(t *testing.T) {
	marketplace := getMarketplace()
	nft := getMarketplaceNft()

	tx, err := marketplace.Encoder.ApproveCreateListing(context.Background(), adminWallet, &NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)
	toAddress := tx.To()
	assert.Equal(t, toAddress.Hex(), nft.helper.getAddress().Hex())

	isApproved, err := nft.IsApproved(context.Background(), adminWallet, marketplace.helper.getAddress().Hex())
	assert.Equal(t, isApproved, false)

	nft.SetApprovalForAll(context.Background(), marketplace.helper.getAddress().Hex(), true)

	isApproved, err = nft.IsApproved(context.Background(), adminWallet, marketplace.helper.getAddress().Hex())
	assert.Equal(t, isApproved, true)

	tx, err = marketplace.Encoder.CreateListing(context.Background(), adminWallet, &NewDirectListing{
		AssetContractAddress:     nft.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 1,
		CurrencyContractAddress:  "0x0000000000000000000000000000000000000000",
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)
	toAddress = tx.To()
	assert.Equal(t, toAddress.Hex(), marketplace.helper.getAddress().Hex())
}

func TestBuyoutListingApprovedEncoder(t *testing.T) {
	marketplace := getMarketplace()
	token := getMarketplaceToken()
	edition := getMarketplaceEdition()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     edition.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 10,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)

	tx, err := marketplace.Encoder.ApproveBuyoutListing(context.Background(), adminWallet, listingId, 10, adminWallet)
	assert.Nil(t, err)
	toAddress := tx.To()
	assert.Equal(t, toAddress.Hex(), token.helper.getAddress().Hex())

	token.SetAllowance(context.Background(), marketplace.helper.getAddress().Hex(), 1000000)

	tx, err = marketplace.Encoder.BuyoutListing(context.Background(), adminWallet, listingId, 10, adminWallet)
	assert.Nil(t, err)
	toAddress = tx.To()
	assert.Equal(t, toAddress.Hex(), marketplace.helper.getAddress().Hex())
}

func TestGenericEncoder(t *testing.T) {
	marketplace := getMarketplace()
	token := getMarketplaceToken()
	edition := getMarketplaceEdition()

	listingId, err := marketplace.CreateListing(context.Background(), &NewDirectListing{
		AssetContractAddress:     edition.helper.getAddress().Hex(),
		TokenId:                  0,
		StartTimeInEpochSeconds:  int(time.Now().Unix()) - 1000,
		ListingDurationInSeconds: 10000,
		Quantity:                 10,
		CurrencyContractAddress:  token.helper.getAddress().Hex(),
		BuyoutPricePerToken:      1.0,
	})
	assert.Nil(t, err)

	tx, err := marketplace.Encoder.Encode(context.Background(), adminWallet, "cancelDirectListing", listingId)
	assert.Nil(t, err)
	toAddress := tx.To()
	assert.Equal(t, toAddress.Hex(), marketplace.helper.getAddress().Hex())
	assert.Equal(t, hex.EncodeToString(tx.Data()), "7506c84a0000000000000000000000000000000000000000000000000000000000000000")
}
