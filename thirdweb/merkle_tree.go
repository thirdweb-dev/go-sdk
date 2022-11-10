package thirdweb

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/cbergoon/merkletree"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/sha3"
)

type ShardedMerkleTree struct {
	storage storage
	baseUri string
	originalEntriesUri string
	shardNybbles int
	tokenDecimals int
	shards map[string]ShardData
	trees map[string]*merkletree.MerkleTree
}

func newShardedMerkleTree(
	storage storage,
	baseUri string,
	originalEntriesUri string,
	shardNybbles int,
	tokenDecimals int,
) *ShardedMerkleTree {
	return &ShardedMerkleTree{
		storage,
		baseUri,
		originalEntriesUri,
		shardNybbles,
		tokenDecimals,
		make(map[string]ShardData),
		make(map[string]*merkletree.MerkleTree),
	}
}

func shardedMerkleTreeFromInfo(metadata *ShardedMerkleTreeInfo, storage storage) *ShardedMerkleTree {
	return newShardedMerkleTree(
		storage,
		metadata.BaseUri,
		metadata.OriginalEntriesUri,
		metadata.ShardNybbles,
		metadata.TokenDecimals,
	)
}

func (tree *ShardedMerkleTree) FetchAndCacheDecimals(
	ctx context.Context,
	cache map[string]int,
	provider *ethclient.Client,
	currencyAddress string,
) (int, error) {
	decimals, exists := cache[currencyAddress]
	if !exists {
		currencyMetadata, err := fetchCurrencyMetadata(ctx, provider, currencyAddress)
		if err != nil {
			return 0, err
		}

		decimals = currencyMetadata.Decimals
		cache[currencyAddress] = decimals
	}

	return decimals, nil
}

func (tree *ShardedMerkleTree) HashEntry(
	entry *SnapshotEntry,
	tokenDecimals int,
	currencyDecimals int,
) (string, error) {
	maxClaimable, err := parseUnits(float64(entry.MaxClaimable), tokenDecimals)
	if err != nil {
		return "", err 
	}

	price, err := parseUnits(float64(entry.Price), currencyDecimals)
	if err != nil {
		return "", err 
	}

	hash := solsha3.SoliditySHA3(
		[]string{"address", "uint256", "uint256", "address"},
		[]interface{}{
			entry.Address,
			maxClaimable.String(),
			price.String(),
			entry.CurrencyAddress,
		},
	)

	return hex.EncodeToString(hash), nil
}

func (tree *ShardedMerkleTree) GetProof(
	ctx context.Context, 
	address string,
	provider *ethclient.Client,
) (*SnapshotEntryWithProof, error) {
	shardId := strings.ToLower(address[2:tree.shardNybbles + 2])
	shard, exists := tree.shards[shardId]

	currencyDecimalMap := make(map[string]int)

	if !exists {
		var shardData ShardData
		raw, err := tree.storage.Get(tree.baseUri + `/` + shardId + `.json`)
		if err != nil {
			return nil, err
		}
		
		err = json.Unmarshal(raw, &shardData)
		if err != nil {
			return nil, err
		}

		tree.shards[shardId] = shardData
		shard = tree.shards[shardId]

		hashedEntries := make([]merkletree.Content, len(shard.Entries))
		for _, entry := range shard.Entries {
			currencyDecimals, err := tree.FetchAndCacheDecimals(ctx, currencyDecimalMap, provider, entry.CurrencyAddress)
			if err != nil {
				return nil, err
			}

			hash, err := tree.HashEntry(&entry, tree.tokenDecimals, currencyDecimals)
			if err != nil {
				return nil, err
			}

			hashedEntries = append(hashedEntries, MerkleNode{Hash: hash})
		}
		
		merkleTree, err := merkletree.NewTree(hashedEntries)
		if err != nil {
			return nil, err
		}

		tree.trees[shardId] = merkleTree
	}

	var entry *SnapshotEntry
	for _, e := range shard.Entries {
		if strings.ToLower(e.Address) == strings.ToLower(address) {
			entry = &e
			break
		}
	}

	if entry == nil {
		return nil, nil
	}

	currencyDecimals, err := tree.FetchAndCacheDecimals(ctx, currencyDecimalMap, provider, entry.CurrencyAddress)
	if err != nil {
		return nil, err
	}

	leaf, err := tree.HashEntry(entry, tree.tokenDecimals, currencyDecimals)
	if err != nil {
		return nil, err
	}

	proofData, _, err := tree.trees[shardId].GetMerklePath(MerkleNode{Hash: leaf})
	if err != nil {
		return nil, err
	}

	var proof [][32]byte;
	for _, p := range proofData {
		var proofItem [32]byte
		copy(proofItem[:], p)
		proof = append(proof, proofItem)
	}

	return &SnapshotEntryWithProof{
		Address: entry.Address,
		MaxClaimable: big.NewInt(int64(entry.MaxClaimable)),
		Price: big.NewInt(int64(entry.Price)),
		CurrencyAddress: entry.CurrencyAddress,
		Proof: proof,
	}, nil
}

type MerkleNode struct {
	Hash string
}

func (t MerkleNode) CalculateHash() ([]byte, error) {
	h := sha3.NewLegacyKeccak256()

  if _, err := h.Write([]byte(t.Hash)); err != nil {
    return nil, err
  }

  return h.Sum(nil), nil
}

func (t MerkleNode) Equals(other merkletree.Content) (bool, error) {
  return t.Hash == other.(MerkleNode).Hash, nil
}