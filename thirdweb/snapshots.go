package thirdweb

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/cbergoon/merkletree"
	"github.com/mitchellh/mapstructure"
)

type SnapshotInput struct {
	Address      string
	MaxClaimable int
}

type SnapshotClaim struct {
	Address      string   `json:"address"`
	MaxClaimable int      `json:"maxClaimable"`
	Proof        []string `json:"proof"`
}

type SnapshotInfo struct {
	MerkleRoot string          `json:"merkleRoot"`
	Claims     []SnapshotClaim `json:"claims"`
}

type SnapshotInfos struct {
	Snapshot    SnapshotInfo
	MerkleRoot  string
	SnapshotUri string
}

type MerkleContent struct {
	address string
}

func (t MerkleContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.address)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (t MerkleContent) Equals(other merkletree.Content) (bool, error) {
	return t.address == other.(MerkleContent).address, nil
}

func createSnapshot(
	snapshotInput []*SnapshotInput,
	tokenDecimals int,
	storage storage,
) (*SnapshotInfos, error) {
	addresses := []string{}
	for _, s := range snapshotInput {
		addresses = append(addresses, s.Address)
	}

	hasDuplicates := make(map[string]bool)
	for _, address := range addresses {
		if hasDuplicates[address] {
			return nil, fmt.Errorf("DUPLICATE_LEAFS: Address %s is duplicated in snapshot", address)
		}
		hasDuplicates[address] = true
	}

	merkleContent := []merkletree.Content{}
	for _, address := range addresses {
		merkleContent = append(merkleContent, MerkleContent{address: address})
	}

	tree, err := merkletree.NewTree(merkleContent)
	if err != nil {
		return nil, err
	}

	merkleRoot := hex.EncodeToString(tree.MerkleRoot())
	claims := []SnapshotClaim{}
	for _, s := range snapshotInput {
		proof, _, _ := tree.GetMerklePath(MerkleContent{address: s.Address})
		encodedProofs := []string{}
		for _, p := range proof {
			encodedProofs = append(encodedProofs, hex.EncodeToString(p))
		}

		claims = append(claims, SnapshotClaim{
			Address:      s.Address,
			MaxClaimable: s.MaxClaimable,
			Proof:        encodedProofs,
		})
	}

	snapshot := SnapshotInfo{
		MerkleRoot: merkleRoot,
		Claims:     claims,
	}

	// TODO: Hash address and max claimable into content

	// TODO: Upload metadata to IPFS
	snapshotToUpload := map[string]interface{}{}
	if err := mapstructure.Decode(snapshot, &snapshotToUpload); err != nil {
		return nil, err
	}
	uri, err := storage.Upload(context.Background(), snapshotToUpload, "", "")
	if err != nil {
		return nil, err
	}

	return &SnapshotInfos{
		Snapshot:    snapshot,
		MerkleRoot:  merkleRoot,
		SnapshotUri: uri,
	}, nil
}
