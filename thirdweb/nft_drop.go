package thirdweb

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

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
	*ERC721Standard
	Abi             *abi.DropERC721
	Helper          *contractHelper
	ClaimConditions *NFTDropClaimConditions
	Encoder         *NFTDropEncoder
	Events          *ContractEvents
}

func newNFTDrop(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*NFTDrop, error) {
	if contractAbi, err := abi.NewDropERC721(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc721, err := newERC721Standard(provider, address, privateKey, storage); err != nil {
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

				events, err := newContractEvents(abi.DropERC721ABI, helper)
				if err != nil {
					return nil, err
				}

				nftCollection := &NFTDrop{
					erc721,
					contractAbi,
					helper,
					claimConditions,
					encoder,
					events,
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
		address = nft.Helper.GetSignerAddress().String()
	}

	if tokenIds, err := nft.GetOwnedTokenIDs(ctx, address); err != nil {
		return nil, err
	} else {
		return nft.erc721.fetchNFTsByTokenId(ctx, tokenIds)
	}
}

// Get the tokenIds of all the NFTs owned by a specific address.
//
// address: the address of the owner of the NFTs
//
// returns: the tokenIds of all the NFTs owned by the address
func (nft *NFTDrop) GetOwnedTokenIDs(ctx context.Context, address string) ([]*big.Int, error) {
	if address == "" {
		address = nft.Helper.GetSignerAddress().String()
	}

	totalCount, err := nft.erc721.drop.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	tokenIds := []*big.Int{}
	for i := 0; i < int(totalCount.Int64()); i++ {
		owner, err := nft.erc721.drop.OwnerOf(&bind.CallOpts{Context: ctx}, big.NewInt(int64(i)))
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
	return drop.erc721.GetAllClaimed(ctx)
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
	return drop.erc721.GetAllUnclaimed(ctx)
}

// Get the total number of NFTs that have been claimed.
func (drop *NFTDrop) TotalClaimedSupply(ctx context.Context) (int, error) {
	return drop.erc721.TotalClaimedSupply(ctx)
}

// Get the total number of NFTs that have not yet been claimed.
func (drop *NFTDrop) TotalUnclaimedSupply(ctx context.Context) (int, error) {
	return drop.erc721.TotalUnclaimedSupply(ctx)
}

func (drop *NFTDrop) GetTotalClaimed(ctx context.Context, address string) (*big.Int, error) {
	events, err := drop.Events.GetEvents(ctx, "TokensClaimed", EventQueryOptions{
		Filters: map[string]interface{}{
			"receiver": common.HexToAddress(address),
		},
	})
	if err != nil {
		return nil, err
	}

	totalClaimed := big.NewInt(0)
	for _, event := range events {
		totalClaimed.Add(totalClaimed, event.Data["quantityClaimed"].(*big.Int))
	}

	return totalClaimed, nil
}

func (drop *NFTDrop) GetClaimInfo(ctx context.Context, address string) (*ClaimInfo, error) {
	return drop.erc721.GetClaimInfo(ctx, address)
}

func (drop *NFTDrop) GetClaimIneligibilityReasons(ctx context.Context, quantity int, addressToCheck string) ([]ClaimEligibility, error) {
	return drop.erc721.GetClaimIneligibilityReasons(ctx, quantity, addressToCheck)
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
	return drop.erc721.CreateBatch(ctx, metadatas)
}

// Claim NFTs from this contract to the connect wallet.
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *NFTDrop) Claim(ctx context.Context, quantity int) (*types.Transaction, error) {
	return drop.erc721.Claim(ctx, quantity)
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
	return drop.erc721.ClaimTo(ctx, destinationAddress, quantity)
}

func (drop *NFTDrop) GetClaimArguments(
	ctx context.Context,
	destinationAddress string,
	quantity int,
) (
	*ClaimArguments,
	error,
) {
	return drop.erc721.GetClaimArguments(ctx, destinationAddress, quantity)
}
