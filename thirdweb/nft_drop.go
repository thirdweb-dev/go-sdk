package thirdweb

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mitchellh/mapstructure"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// You can access the NFT Drop interface from the SDK as follows:
//
//	import (
//		"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
//	)
//
//	privateKey = "..."
//
//	sdk, err := thirdweb.NewThirdwebSDK("mumbai", &thirdweb.SDKOptions{
//		PrivateKey: privateKey,
//	})
//
//	contract, err := sdk.GetNFTDrop("{{contract_address}}")
type NFTDrop struct {
	Abi    *abi.DropERC721
	helper *contractHelper
	*ERC721
	ClaimConditions *NFTDropClaimConditions
	Encoder         *NFTDropEncoder
}

func newNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTDrop, error) {
	if contractAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				claimConditions, err := newNFTDropClaimConditions(address, provider, helper, storage)
				if err != nil {
					return nil, err
				}

				encoder, err := newNFTDropEncoder(contractAbi, helper, claimConditions, storage)
				if err != nil {
					return nil, err
				}

				nftCollection := &NFTDrop{
					contractAbi,
					helper,
					erc721,
					claimConditions,
					encoder,
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
//	owner := "{{wallet_address}}"
//	nfts, err := contract.GetOwned(context.Background(), owner)
//	name := nfts[0].Metadata.Name
func (nft *NFTDrop) GetOwned(ctx context.Context, address string) ([]*NFTMetadataOwner, error) {
	if address == "" {
		address = nft.helper.GetSignerAddress().String()
	}

	if tokenIds, err := nft.GetOwnedTokenIDs(ctx, address); err != nil {
		return nil, err
	} else {
		return nft.fetchNFTsByTokenId(ctx, tokenIds)
	}
}

// Get the tokenIds of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the tokenIds of all the NFTs owned by the address
func (nft *NFTDrop) GetOwnedTokenIDs(ctx context.Context, address string) ([]*big.Int, error) {
	if address == "" {
		address = nft.helper.GetSignerAddress().String()
	}

	totalCount, err := nft.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	tokenIds := []*big.Int{}
	for i := 0; i < int(totalCount.Int64()); i++ {
		owner, err := nft.abi.OwnerOf(&bind.CallOpts{}, big.NewInt(int64(i)))
		if err != nil {
			return nil, err
		} 

		if strings.ToLower(owner.String()) == strings.ToLower(address) {
			tokenIds = append(tokenIds, big.NewInt(int64(i)))
		}
	}

	return tokenIds, nil
}


// Get a list of all the NFTs that have been claimed from this contract.
//
// returns: a list of the metadatas of the claimed NFTs
//
// Example
//
//	claimedNfts, err := contract.GetAllClaimed(context.Background())
//	firstOwner := claimedNfts[0].Owner
func (drop *NFTDrop) GetAllClaimed(ctx context.Context) ([]*NFTMetadataOwner, error) {
	if maxId, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{}); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for i := 0; i < int(maxId.Int64()); i++ {
			if nft, err := drop.Get(ctx, i); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

// Get a list of all the NFTs on this contract that have not yet been claimed.
//
// returns: a list of the metadatas of the unclaimed NFTs
//
// Example
//
//	unclaimedNfts, err := contract.GetAllUnclaimed(context.Background())
//	firstNftName := unclaimedNfts[0].Name
func (drop *NFTDrop) GetAllUnclaimed(ctx context.Context) ([]*NFTMetadata, error) {
	maxId, err := drop.Abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	unmintedId, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}

	nfts := []*NFTMetadata{}
	for i := int(unmintedId.Int64()); i < int(maxId.Int64()); i++ {
		if nft, err := drop.getTokenMetadata(ctx, i); err == nil {
			nfts = append(nfts, nft)
		}
	}

	return nfts, nil
}

// Get the total number of NFTs that have been claimed.
func (drop *NFTDrop) TotalClaimedSupply() (int, error) {
	claimed, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	return int(claimed.Int64()), nil
}

// Get the total number of NFTs that have not yet been claimed.
func (drop *NFTDrop) TotalUnclaimedSupply() (int, error) {
	claimed, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	total, err := drop.Abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return 0, err
	}

	unclaimed := big.NewInt(0).Sub(total, claimed)
	return int(unclaimed.Int64()), nil
}

// Create a batch of NFTs on this contract.
//
// metadatas: a list of the metadatas of the NFTs to create
//
// returns: the transaction receipt of the batch creation
//
// Example
//
//	image0, err := os.Open("path/to/image/0.jpg")
//	defer image0.Close()
//
//	image1, err := os.Open("path/to/image/1.jpg")
//	defer image1.Close()
//
//	metadatas := []*thirdweb.NFTMetadataInput{
//		&thirdweb.NFTMetadataInput{
//			Name: "Cool NFT",
//			Description: "This is a cool NFT",
//			Image: image1
//		}
//		&thirdweb.NFTMetadataInput{
//			Name: "Cool NFT 2",
//			Description: "This is also a cool NFT",
//			Image: image2
//		}
//	}
//
//	tx, err := contract.CreateBatch(context.Background(), metadatas)
func (drop *NFTDrop) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	startNumber, err := drop.Abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	fileStartNumber := int(startNumber.Int64())

	contractAddress := drop.helper.getAddress().String()
	signerAddress := drop.helper.GetSignerAddress().String()

	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}

	dataToUpload := []map[string]interface{}{}
	mapstructure.Decode(data, &dataToUpload)

	batch, err := drop.storage.UploadBatch(
		dataToUpload,
		fileStartNumber,
		contractAddress,
		signerAddress,
	)

	txOpts, err := drop.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := drop.Abi.LazyMint(
		txOpts,
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.helper.awaitTx(tx.Hash())
}

// Claim NFTs from this contract to the connect wallet.
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *NFTDrop) Claim(ctx context.Context, quantity int) (*types.Transaction, error) {
	address := drop.helper.GetSignerAddress().String()
	return drop.ClaimTo(ctx, address, quantity)
}

// Claim NFTs from this contract to the connect wallet.
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
//
// Example
//
//	address := "{{wallet_address}}"
//	quantity = 1
//
//	tx, err := contract.ClaimTo(context.Background(), address, quantity)
func (drop *NFTDrop) ClaimTo(ctx context.Context, destinationAddress string, quantity int) (*types.Transaction, error) {
	active, err := drop.ClaimConditions.GetActive()
	if err != nil {
		return nil, err
	}
	
	claimVerification, err := drop.prepareClaim(ctx, quantity)
	if err != nil {
		return nil, err
	}

	txOpts, err := drop.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	txOpts.Value = claimVerification.Value

	proof := abi.IDropAllowlistProof{
		Proof: claimVerification.Proofs,
		QuantityLimitPerWallet: claimVerification.MaxClaimable,
		PricePerToken: claimVerification.Price,
		Currency: common.HexToAddress(claimVerification.CurrencyAddress),
	}

	tx, err := drop.Abi.Claim(
		txOpts,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(active.CurrencyAddress),
		active.Price,
		proof,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.helper.awaitTx(tx.Hash())
}

func (drop *NFTDrop) prepareClaim(ctx context.Context, quantity int) (*ClaimVerification, error) {
	active, err := drop.ClaimConditions.GetActive()
	if err != nil {
		return nil, err
	}

	merkleMetadata, err := drop.ClaimConditions.GetMerkleMetadata()
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		ctx,
		quantity,
		active,
		merkleMetadata,
		drop.helper,
		drop.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}
