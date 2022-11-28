package thirdweb

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mitchellh/mapstructure"
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
	*ERC1155
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
			if erc1155, err := newERC1155(provider, address, privateKey, storage); err != nil {
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
	startNumber, err := drop.abi.NextTokenIdToMint(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	fileStartNumber := int(startNumber.Int64())

	contractAddress := drop.Helper.getAddress().String()
	signerAddress := drop.Helper.GetSignerAddress().String()

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

	txOpts, err := drop.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := drop.abi.LazyMint(
		txOpts,
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.Helper.AwaitTx(tx.Hash())
}

// Claim NFTs from this contract to the connect wallet.
//
// tokenId: the token ID of the NFT to claim
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *EditionDrop) Claim(ctx context.Context, tokenId int, quantity int) (*types.Transaction, error) {
	address := drop.Helper.GetSignerAddress().String()
	return drop.ClaimTo(ctx, address, tokenId, quantity)
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
	claimVerification, err := drop.prepareClaim(ctx, tokenId, quantity)
	if err != nil {
		return nil, err
	}

	active, err := drop.ClaimConditions.GetActive(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	txOpts, err := drop.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	txOpts.Value = claimVerification.Value

	proof := abi.IDrop1155AllowlistProof{
		Proof: claimVerification.Proofs,
		QuantityLimitPerWallet: claimVerification.MaxClaimable,
		PricePerToken: claimVerification.Price,
		Currency: common.HexToAddress(claimVerification.CurrencyAddress),
	}

	tx, err := drop.abi.Claim(
		txOpts,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(quantity)),
		common.HexToAddress(active.CurrencyAddress),
		active.Price,
		proof,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return drop.Helper.AwaitTx(tx.Hash())
}

func (drop *EditionDrop) prepareClaim(ctx context.Context, tokenId int, quantity int) (*ClaimVerification, error) {
	addressToClaim := drop.helper.GetSignerAddress().Hex()
	claimCondition, err := drop.ClaimConditions.GetActive(ctx, tokenId)
	if err != nil {
		return nil, err
	}


	merkleMetadata, err := drop.ClaimConditions.GetMerkleMetadata()
	if err != nil {
		return nil, err
	}


	claimVerification, err := prepareClaim(
		ctx,
		addressToClaim,
		quantity,
		claimCondition,
		merkleMetadata,
		drop.Helper,
		drop.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}
