package thirdweb

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerkleTreeSmall(t *testing.T) {
	sdk, err := NewThirdwebSDK("http://localhost:8545", nil)
	if err != nil {
		panic(err)
	}

	uri := "ipfs://QmeAx8aRvsYXN6mzky72b9V1HWokb271FoBmDu4tatC8hE/0"
	storage := newIpfsStorage("", defaultIpfsGatewayUrl, http.DefaultClient)

	body, err := storage.Get(context.Background(), uri)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)

	var info ShardedMerkleTreeInfo
	if err := json.Unmarshal(body, &info); err != nil {
		panic(err)
	}

	tree := shardedMerkleTreeFromInfo(&info, storage)
	proof, err := tree.GetProof(context.Background(), "0x0000000000000000000000000000000000000000", sdk.GetProvider())
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)

	fmt.Printf("%#v", proof.Proof)
}

func TestMerkleTreeEdgeCase(t *testing.T) {
	sdk, err := NewThirdwebSDK("http://localhost:8545", nil)
	if err != nil {
		panic(err)
	}

	uri := "ipfs://QmacDnA4i7Za19LpE3pngwLfUtakn71ghaKKjTkM2Phzj8/0"
	storage := newIpfsStorage("", defaultIpfsGatewayUrl, http.DefaultClient)

	body, err := storage.Get(context.Background(), uri)
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)

	var info ShardedMerkleTreeInfo
	if err := json.Unmarshal(body, &info); err != nil {
		panic(err)
	}

	tree := shardedMerkleTreeFromInfo(&info, storage)
	proof, err := tree.GetProof(context.Background(), "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6", sdk.GetProvider())
	if err != nil {
		panic(err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, proof)

	correctProofs := []string{
		"36f4f9e91158fce37ae8b1f3443025384d220b25db67976898893f3418722bee",
		"35a5350f8ef7d3351bad08b2476dba6ec6939d501b889d98e726ae6a3e822ef4",
		"8cbf083257ff0aa97fb2b28f0b9a07c3dc6e9820f89b584ef0137d50f82b7b71",
		"e0be2bf7a799f716120eb3f15b6e2139e701c070b7333d5d19baf2154c2c9f95",
		"91d36657024683736d0c306fcbe60f236f24dcd3c03c5c39dc48c1c0fb08418d",
	}

	for i, correctProof := range correctProofs {
		proofBytes, err := hex.DecodeString(correctProof)
		assert.Nil(t, err)

		assert.Equal(t, bytes.Compare(proof.Proof[i][:], proofBytes), 0)
	}
}
