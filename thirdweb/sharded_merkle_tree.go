package thirdweb

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"golang.org/x/crypto/sha3"

	merkletree "github.com/thirdweb-dev/go-sdk/v2/merkle"
)

type ShardedMerkleTree struct {
	storage            storage
	baseUri            string
	originalEntriesUri string
	shardNybbles       int
	tokenDecimals      int
	shards             map[string]ShardData
	trees              map[string]*merkletree.MerkleTree
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
	maxClaimable, err := convertQuantityToBigNumber(entry.MaxClaimable, tokenDecimals)
	if err != nil {
		return "", err
	}

	entryPrice := entry.Price
	if entryPrice == "" {
		entryPrice = "unlimited"
	}

	price, err := convertQuantityToBigNumber(entryPrice, currencyDecimals)
	if err != nil {
		return "", err
	}

	currencyAddress := entry.CurrencyAddress
	if currencyAddress == "" {
		currencyAddress = zeroAddress
	}

	hash := solsha3.SoliditySHA3(
		[]string{"address", "uint256", "uint256", "address"},
		[]interface{}{
			entry.Address,
			maxClaimable.String(),
			price.String(),
			currencyAddress,
		},
	)

	return hex.EncodeToString(hash), nil
}

func (tree *ShardedMerkleTree) GetProof(
	ctx context.Context,
	address string,
	provider *ethclient.Client,
) (*SnapshotEntryWithProof, error) {
	shardId := strings.ToLower(address[2 : tree.shardNybbles+2])
	shard, exists := tree.shards[shardId]

	currencyDecimalMap := make(map[string]int)

	if !exists {
		var shardData ShardData
		uri := tree.baseUri + `/` + shardId + `.json`
		raw, err := tree.storage.Get(ctx, uri)
		if err != nil {
			fmt.Println("[warning] specified address is not in merkle tree")
			return nil, nil
		}

		err = json.Unmarshal(raw, &shardData)
		if err != nil {
			return nil, err
		}

		tree.shards[shardId] = shardData
		shard = tree.shards[shardId]

		hashedEntries := []merkletree.DataBlock{}
		for _, entry := range shard.Entries {
			currencyDecimals, err := tree.FetchAndCacheDecimals(ctx, currencyDecimalMap, provider, entry.CurrencyAddress)
			if err != nil {
				return nil, err
			}

			hash, err := tree.HashEntry(&entry, tree.tokenDecimals, currencyDecimals)
			if err != nil {
				return nil, err
			}

			hashBytes, err := hex.DecodeString(hash)
			if err != nil {
				return nil, err
			}

			hashedEntries = append(hashedEntries, &MerkleNode{data: hashBytes})
		}

		// Define our own hash function so we can match merkletreejs
		calculateHash := func(data []byte) ([]byte, error) {
			// Avoid hashing the leaf nodes to match merkletreejs implementation
			for _, node := range hashedEntries {
				node, err := node.Serialize()
				if err != nil {
					return nil, err
				}

				if reflect.DeepEqual(node, data) {
					return data, nil
				}
			}

			h := sha3.NewLegacyKeccak256()

			if _, err := h.Write(data); err != nil {
				return nil, err
			}

			return h.Sum(nil), nil
		}

		if len(hashedEntries) > 1 {
			config := &merkletree.Config{
				Mode:       merkletree.ModeProofGenAndTreeBuild,
				HashFunc:   calculateHash,
				SortLeaves: true,
				SortPairs:  true,
			}

			merkleTree, err := merkletree.New(config, hashedEntries)
			if err != nil {
				return nil, err
			}

			tree.trees[shardId] = merkleTree
		}
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

	leafBytes, err := hex.DecodeString(leaf)
	if err != nil {
		return nil, err
	}

	proof := [][32]byte{}
	currentTree := tree.trees[shardId]
	if currentTree != nil {
		proofData, err := currentTree.GenerateProof(&MerkleNode{data: leafBytes})
		if err != nil {
			return nil, err
		}

		for _, p := range proofData.Siblings {
			var proofItem [32]byte
			copy(proofItem[:], p)
			proof = append(proof, proofItem)
		}
	}

	for _, p := range shard.Proofs {
		var proofItem [32]byte
		proofBytes, err := hex.DecodeString(p[2:])
		if err != nil {
			return nil, err
		}

		copy(proofItem[:], proofBytes)
		proof = append(proof, proofItem)
	}

	return &SnapshotEntryWithProof{
		Address:         entry.Address,
		MaxClaimable:    entry.MaxClaimable,
		Price:           entry.Price,
		CurrencyAddress: entry.CurrencyAddress,
		Proof:           proof,
	}, nil
}

type MerkleNode struct {
	data []byte
}

func (t *MerkleNode) Serialize() ([]byte, error) {
	return t.data, nil
}
