package thirdweb

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestMerkleTree(t *testing.T) {
	sdk, err := NewThirdwebSDK("http://localhost:8545", nil)
	if err != nil {
		panic(err)
	}

	uri := "ipfs://QmeAx8aRvsYXN6mzky72b9V1HWokb271FoBmDu4tatC8hE/0"
	storage := newIpfsStorage(defaultIpfsGatewayUrl)

	body, err := storage.Get(uri)
	if err != nil {
		panic(err)
	}

	var info ShardedMerkleTreeInfo
	if err := json.Unmarshal(body, &info); err != nil {
		fmt.Println("Failed here...")
		panic(err)
	}

	tree := shardedMerkleTreeFromInfo(&info, storage)
	proof, err := tree.GetProof(context.Background(), "0x0000000000000000000000000000000000000000", sdk.GetProvider())
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", proof.Proof)
}