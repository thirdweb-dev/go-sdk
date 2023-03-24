package thirdweb

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// You can access the Edition Drop interface from the SDK as follows:
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
//	contract, err := sdk.GetEditionDrop("{{contract_address}}")
type EditionDrop struct {
	*ERC1155Standard
	abi             *abi.DropERC1155
	Helper          *contractHelper
	ClaimConditions *EditionDropClaimConditions
	Encoder         *ContractEncoder
	Events          *ContractEvents
}

func newEditionDrop(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*EditionDrop, error) {
	if contractAbi, err := abi.NewDropERC1155(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			if erc1155, err := newERC1155Standard(provider, address, privateKey, storage); err != nil {
				return nil, err
			} else {
				claimConditions, err := newEditionDropClaimConditions(address, provider, helper, storage)
				if err != nil {
					return nil, err
				}

				encoder, err := newContractEncoder(abi.DropERC1155ABI, helper)
				if err != nil {
					return nil, err
				}

				events, err := newContractEvents(abi.DropERC1155ABI, helper)
				if err != nil {
					return nil, err
				}

				edition := &EditionDrop{
					erc1155,
					contractAbi,
					helper,
					claimConditions,
					encoder,
					events,
				}
				return edition, nil
			}
		}
	}
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
//	metadatasWithSupply := []*thirdweb.EditionMetadataInput{
//		&thirdweb.EditionMetadataInput{
//			Metadata: &thirdweb.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//				Image: image0,
//			},
//			Supply: 100,
//		},
//		&thirdweb.EditionMetadataInput{
//			Metadata: &thirdweb.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//				Image: image1,
//			},
//			Supply: 100,
//		},
//	}
//
//	tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatasWithSupply)
func (drop *EditionDrop) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	return drop.erc1155.CreateBatch(ctx, metadatas)
}

// Claim NFTs from this contract to the connect wallet.
//
// tokenId: the token ID of the NFT to claim
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *EditionDrop) Claim(ctx context.Context, tokenId int, quantity int) (*types.Transaction, error) {
	return drop.erc1155.Claim(ctx, tokenId, quantity)
}

// Claim NFTs from this contract to the connect wallet.
//
// tokenId: the token ID of the NFT to claim
//
// destinationAddress: the address of the wallet to claim the NFTs to
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
//
// Example
//
//	address = "{{wallet_address}}"
//	tokenId = 0
//	quantity = 1
//
//	tx, err := contract.ClaimTo(context.Background(), address, tokenId, quantity)
func (drop *EditionDrop) ClaimTo(ctx context.Context, destinationAddress string, tokenId int, quantity int) (*types.Transaction, error) {
	return drop.erc1155.ClaimTo(ctx, destinationAddress, tokenId, quantity)
}