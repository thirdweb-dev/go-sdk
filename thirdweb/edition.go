package thirdweb

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// You can access the Edition interface from the SDK as follows:
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
//	contract, err := sdk.GetEdition("{{contract_address}}")
type Edition struct {
	*ERC1155Standard
	abi       *abi.TokenERC1155
	Helper    *contractHelper
	Signature *ERC1155SignatureMinting
	Encoder   *ContractEncoder
	Events    *ContractEvents
}

func newEdition(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*Edition, error) {
	if contractAbi, err := abi.NewTokenERC1155(address, provider); err != nil {
		return nil, err
	} else {
		if helper, err := newContractHelper(address, provider, privateKey); err != nil {
			return nil, err
		} else {
			erc1155, err := newERC1155Standard(provider, address, privateKey, storage)
			if err != nil {
				return nil, err
			}

			signature, err := newERC1155SignatureMinting(provider, address, privateKey, storage)
			if err != nil {
				return nil, err
			}

			encoder, err := newContractEncoder(abi.TokenERC1155ABI, helper)
			if err != nil {
				return nil, err
			}

			events, err := newContractEvents(abi.TokenERC1155ABI, helper)
			if err != nil {
				return nil, err
			}

			edition := &Edition{
				erc1155,
				contractAbi,
				helper,
				signature,
				encoder,
				events,
			}
			return edition, nil
		}
	}
}

// Mint an NFT to the connected wallet.
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) Mint(ctx context.Context, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	return edition.erc1155.Mint(ctx, metadataWithSupply)
}

// Mint a new NFT to the specified wallet.
//
// address: the wallet address to mint the NFT to
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//		image, err := os.Open("path/to/image.jpg")
//		defer image.Close()
//
//		metadataWithSupply := &thirdweb.EditionMetadataInput{
//	        context.Background(),
//			Metadata: &thirdweb.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//				Image: image,
//			},
//			Supply: 100,
//		}
//
//		tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", metadataWithSupply)
func (edition *Edition) MintTo(ctx context.Context, address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	return edition.erc1155.MintTo(ctx, address, metadataWithSupply)
}

// Mint additionaly supply of a token to the connected wallet.
//
// tokenId: token ID to mint additional supply of
//
// additionalSupply: additional supply to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintAdditionalSupply(ctx context.Context, tokenId int, additionalSupply int) (*types.Transaction, error) {
	return edition.erc1155.MintAdditionalSupply(ctx, tokenId, additionalSupply)
}

// Mint additional supply of a token to the specified wallet.
//
// to: address of the wallet to mint NFTs to
//
// tokenId: token Id to mint additional supply of
//
// additionalySupply: additional supply to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintAdditionalSupplyTo(ctx context.Context, to string, tokenId int, additionalSupply int) (*types.Transaction, error) {
	return edition.erc1155.MintAdditionalSupplyTo(ctx, to, tokenId, additionalSupply)
}

// Mint a batch of NFTs to the connected wallet.
//
// metadatasWithSupply: list of NFT metadatas with supplies to mint
//
// returns: the transaction receipt of the mint
func (edition *Edition) MintBatch(ctx context.Context, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	return edition.erc1155.MintBatch(ctx, metadatasWithSupply)
}

// Mint a batch of NFTs to a specific wallet.
//
// to: address of the wallet to mint NFTs to
//
// metadatasWithSupply: list of NFT metadatas with supplies to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	metadatasWithSupply := []*thirdweb.EditionMetadataInput{
//		&thirdweb.EditionMetadataInput{
//			Metadata: &thirdweb.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//			},
//			Supply: 100,
//		},
//		&thirdweb.EditionMetadataInput{
//			Metadata: &thirdweb.NFTMetadataInput{
//				Name: "Cool NFT",
//				Description: "This is a cool NFT",
//			},
//			Supply: 100,
//		},
//	}
//
//	tx, err := contract.MintBatchTo(context.Background(), "{{wallet_address}}", metadatasWithSupply)
func (edition *Edition) MintBatchTo(ctx context.Context, to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	return edition.erc1155.MintBatchTo(ctx, to, metadatasWithSupply)
}
