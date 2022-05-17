package thirdweb

import (
	"encoding/json"
	"math/big"
)

// NFT

func fetchTokenMetadata(tokenId int, uri string, storage Storage) (*NFTMetadata, error) {
	if body, err := storage.Get(uri); err != nil {
		return nil, err
	} else {
		metadata := &NFTMetadata{
			Id: big.NewInt(int64(tokenId)),
		}
		if err := json.Unmarshal(body, &metadata); err != nil {
			return nil, &UnmarshalError{body: string(body), typeName: "nft", UnderlyingError: err}
		}

		return metadata, nil
	}
}

func uploadOrExtractUri(metadata *NFTMetadataInput, storage Storage) (string, error) {
	return storage.Upload(metadata, "", "")
}

func uploadOrExtractUris(metadatas []*NFTMetadataInput, storage Storage) ([]string, error) {
	// Why is this necessary?
	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}

	return storage.UploadBatch(data, "", "")
}
