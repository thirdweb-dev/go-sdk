package thirdweb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNftDrop() *NFTDrop {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployNFTDrop(&DeployNFTDropMetadata{
		Name: "NFT Drop",
	})
	drop, _ := sdk.GetNFTDrop(address)

	return drop
}

func TestCreateBatchNftDrop(t *testing.T) {
	drop := getNftDrop()

	balance, _ := drop.Balance()
	assert.Equal(t, 0, balance)

	_, err := drop.CreateBatch(
		[]*NFTMetadataInput{
			{
				Name: "NFT 1",
			},
			{
				Name: "NFT 2",
			},
		},
	)
	assert.Nil(t, err)

	nfts, _ := drop.GetAllUnclaimed()
	assert.Equal(t, 2, len(nfts))
	assert.Equal(t, nfts[0].Name, "NFT 1")
	assert.Equal(t, nfts[1].Name, "NFT 2")
}
