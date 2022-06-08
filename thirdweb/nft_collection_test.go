package thirdweb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Hardhat wallet addresses
const adminWallet = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
const secondaryWallet = "0x70997970C51812dc3A010C7d01b50e0d17dc79C8"
const tertiaryWallet = "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"

func getSDK() *ThirdwebSDK {
	sdk, _ := NewThirdwebSDK("http://localhost:8545", &SDKOptions{
		PrivateKey: "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
	})

	return sdk
}

func getNft() *NFTCollection {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployNFTCollection(&DeployNFTCollectionMetadata{
		Name: "NFT Collection",
	})
	nft, _ := sdk.GetNFTCollection(address)

	return nft
}

func TestMintNft(t *testing.T) {
	nft := getNft()

	balance, _ := nft.Balance()
	assert.Equal(t, 0, balance)

	nft.Mint(&NFTMetadataInput{
		Name: "NFT",
	})

	balance, _ = nft.Balance()
	assert.Equal(t, 1, balance)
}

func TestBurnNft(t *testing.T) {
	nft := getNft()

	nft.Mint(&NFTMetadataInput{
		Name: "NFT",
	})

	balance, _ := nft.Balance()
	assert.Equal(t, 1, balance)

	_, err := nft.Burn(0)
	assert.Nil(t, err)

	balance, _ = nft.Balance()
	assert.Equal(t, 0, balance)
}

func TestTransferNft(t *testing.T) {
	nft := getNft()

	nft.Mint(&NFTMetadataInput{
		Name: "NFT",
	})

	balance, _ := nft.Balance()
	assert.Equal(t, 1, balance)

	_, err := nft.Transfer(secondaryWallet, 0)
	assert.Nil(t, err)

	balance, _ = nft.Balance()
	assert.Equal(t, 0, balance)
}
