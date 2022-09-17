package thirdweb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getNftDrop() *NFTDrop {
	sdk := getSDK()
	address, _ := sdk.Deployer.DeployNFTDrop(context.Background(), &DeployNFTDropMetadata{
		Name: "NFT Drop",
	})
	drop, _ := sdk.GetNFTDrop(address)

	return drop
}

func TestCreateBatchNftDrop(t *testing.T) {
	drop := getNftDrop()

	balance, _ := drop.Balance(context.Background())
	assert.Equal(t, 0, balance)

	_, err := drop.CreateBatch(
		context.Background(),
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

	nfts, _ := drop.GetAllUnclaimed(context.Background())
	assert.Equal(t, 2, len(nfts))
	assert.Equal(t, nfts[0].Name, "NFT 1")
	assert.Equal(t, nfts[1].Name, "NFT 2")
}
