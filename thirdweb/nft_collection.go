package thirdweb

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

// You can access the NFT Collection interface from the SDK as follows:
//
// 	import (
// 		"github.com/thirdweb-dev/go-sdk/thirdweb"
// 	)
//
// 	privateKey = "..."
//
// 	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
// 	})
//
//	contract, err := sdk.GetNFTCollection("{{contract_address}}")
type NFTCollection struct {
	abi    *abi.TokenERC721
	helper *contractHelper
	*ERC721
	Signature *ERC721SignatureMinting
}

func newNFTCollection(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTCollection, error) {
	if contractAbi, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				signature, err := newERC721SignatureMinting(provider, address, privateKey, storage)
				if err != nil {
					return nil, err
				}

				nftCollection := &NFTCollection{
					contractAbi,
					helper,
					erc721,
					signature,
				}
				return nftCollection, nil
			}
		}
	}
}

// Get the metadatas of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the metadata of all the NFTs owned by the address
//
// Example
//
// 	owner := "{{wallet_address}}"
// 	nfts, err := contract.GetOwned(owner)
// 	name := nfts[0].Metadata.Name
func (nft *NFTCollection) GetOwned(address string) ([]*NFTMetadataOwner, error) {
	if address == "" {
		address = nft.helper.GetSignerAddress().String()
	}

	if tokenIds, err := nft.GetOwnedTokenIDs(address); err != nil {
		return nil, err
	} else {
		return nft.fetchNFTsByTokenId(tokenIds)
	}
}

// Get the tokenIds of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the tokenIds of all the NFTs owned by the address
func (nft *NFTCollection) GetOwnedTokenIDs(address string) ([]*big.Int, error) {
	if address == "" {
		address = nft.helper.GetSignerAddress().String()
	}

	if balance, err := nft.abi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(address)); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}

		for i := 0; i < int(balance.Int64()); i++ {
			if tokenId, err := nft.abi.TokenOfOwnerByIndex(&bind.CallOpts{}, common.HexToAddress(address), big.NewInt(int64(i))); err == nil {
				tokenIds = append(tokenIds, tokenId)
			}
		}

		return tokenIds, nil
	}
}

// Mint a new NFT to the connected wallet.
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) Mint(metadata *NFTMetadataInput) (*types.Transaction, error) {
	address := nft.helper.GetSignerAddress().String()
	return nft.MintTo(address, metadata)
}

// Mint a new NFT to the specified wallet.
//
// address: the wallet address to mint to
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	image, err := os.Open("path/to/image.jpg")
// 	defer image.Close()
//
// 	metadata := &thirdweb.NFTMetadataInput{
// 		Name: "Cool NFT",
// 		Description: "This is a cool NFT",
// 		Image: image,
// 	}
//
// 	tx, err := contract.MintTo("{{wallet_address}}", metadata)
func (nft *NFTCollection) MintTo(address string, metadata *NFTMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(metadata, nft.storage)
	if err != nil {
		return nil, err
	}

	tx, err := nft.abi.MintTo(
		nft.helper.getTxOptions(),
		common.HexToAddress(address),
		uri,
	)
	if err != nil {
		return nil, err
	}

	return nft.helper.awaitTx(tx.Hash())
}

// Mint a batch of new NFTs to the connected wallet.
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
func (nft *NFTCollection) MintBatch(metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	address := nft.helper.GetSignerAddress().String()
	return nft.MintBatchTo(address, metadatas)
}

// Mint a batch of new NFTs to the specified wallet.
//
// to: the wallet address to mint to
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	metadatas := []*thirdweb.NFTMetadataInput{
// 		&thirdweb.NFTMetadataInput{
// 			Name: "Cool NFT",
// 			Description: "This is a cool NFT",
// 		}
// 		&thirdweb.NFTMetadataInput{
// 			Name: "Cool NFT 2",
// 			Description: "This is also a cool NFT",
// 		}
// 	}
//
// 	tx, err := contract.MintBatchTo("{{wallet_address}}", metadatas)
func (nft *NFTCollection) MintBatchTo(address string, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	uris, err := uploadOrExtractUris(metadatas, nft.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	for _, uri := range uris {
		tx, err := nft.abi.MintTo(
			nft.helper.getTxOptions(),
			common.HexToAddress(address), uri,
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := nft.abi.Multicall(nft.helper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return nft.helper.awaitTx(tx.Hash())
}
