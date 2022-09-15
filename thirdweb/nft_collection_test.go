package thirdweb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hardhat wallet addresses
const adminWallet = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const adminPrivateKey = "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
const secondaryWallet = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
const secondaryPrivateKey = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
const tertiaryWallet = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"
const tertiaryPrivateKey = "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"

func getSDK() *ThirdwebSDK {
	sdk, _ := NewThirdwebSDK("http://localhost:8545", &SDKOptions{
		PrivateKey: adminPrivateKey,
	})

	return sdk
}

func getNft() *NFTCollection {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployNFTCollection(context.Background(), &DeployNFTCollectionMetadata{
		Name: "NFT Collection",
	})
	nft, _ := sdk.GetNFTCollection(address)

	return nft
}

func TestMintNft(t *testing.T) {
	nft := getNft()

	balance, _ := nft.Balance(context.Background())
	assert.Equal(t, 0, balance)

	_, err := nft.Mint(context.Background(), &NFTMetadataInput{
		Name: "NFT",
	})
	assert.Nil(t, err)

	balance, _ = nft.Balance(context.Background())
	assert.Equal(t, 1, balance)
}

func TestBatchMintNft(t *testing.T) {
	nft := getNft()

	balance, _ := nft.Balance(context.Background())
	assert.Equal(t, 0, balance)

	_, err := nft.MintBatch(context.Background(), []*NFTMetadataInput{
		{
			Name: "NFT 1",
		},
		{
			Name: "NFT 2",
		},
	})
	assert.Nil(t, err)

	balance, _ = nft.Balance(context.Background())
	assert.Equal(t, 2, balance)

	nfts, _ := nft.GetAll(context.Background())
	assert.Equal(t, 2, len(nfts))
	assert.Equal(t, "NFT 1", nfts[0].Metadata.Name)
	assert.Equal(t, "NFT 2", nfts[1].Metadata.Name)

	nfts, _ = nft.GetOwned(context.Background(), nft.helper.GetSignerAddress().String())
	assert.Equal(t, 2, len(nfts))
	assert.Equal(t, "NFT 1", nfts[0].Metadata.Name)
	assert.Equal(t, "NFT 2", nfts[1].Metadata.Name)
}

func TestBurnNft(t *testing.T) {
	nft := getNft()

	nft.Mint(context.Background(), &NFTMetadataInput{
		Name: "NFT",
	})

	balance, _ := nft.Balance(context.Background())
	assert.Equal(t, 1, balance)

	_, err := nft.Burn(context.Background(), 0)
	assert.Nil(t, err)

	balance, _ = nft.Balance(context.Background())
	assert.Equal(t, 0, balance)
}

func TestTransferNft(t *testing.T) {
	nft := getNft()

	nft.Mint(context.Background(), &NFTMetadataInput{
		Name: "NFT",
	})

	balance, _ := nft.Balance(context.Background())
	assert.Equal(t, 1, balance)

	_, err := nft.Transfer(context.Background(), secondaryWallet, 0)
	assert.Nil(t, err)

	balance, _ = nft.Balance(context.Background())
	assert.Equal(t, 0, balance)
}

func TestSignatureMintNft(t *testing.T) {
	nft := getNft()

	balance, _ := nft.Balance(context.Background())
	assert.Equal(t, 0, balance)

	payload, err := nft.Signature.Generate(
		&Signature721PayloadInput{
			To:                   nft.helper.GetSignerAddress().String(),
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
		},
	)
	assert.Nil(t, err)

	valid, err := nft.Signature.Verify(payload)
	assert.Nil(t, err)

	assert.True(t, valid)

	_, err = nft.Signature.Mint(context.Background(), payload)
	assert.Nil(t, err)

	balance, _ = nft.Balance(context.Background())
	assert.Equal(t, 1, balance)

	metadata, _ := nft.Get(context.Background(), 0)
	assert.Equal(t, "Sigmint", metadata.Metadata.Name)
}
