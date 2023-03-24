package thirdweb

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mitchellh/mapstructure"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// This interface is currently support by the NFT Collection and NFT Drop contracts.
// You can access all of its functions through an NFT Collection or NFT Drop contract instance.
type ERC721 struct {
	token   		*abi.TokenERC721
	drop    		*abi.DropERC721
	helper  		*contractHelper
	storage 		storage
	ClaimConditions *NFTDropClaimConditions
}

type NFTResult struct {
	nft *NFTMetadataOwner
	err error
}

func newERC721(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC721, error) {
	token, err := abi.NewTokenERC721(address, provider)
	if err != nil {
		return nil, err
	}

	drop, err := abi.NewDropERC721(address, provider)
	if err != nil {
		return nil, err
	}
	
	helper, err := newContractHelper(address, provider, privateKey)
	if err != nil {
		return nil, err
	}

	claimConditions, err := newNFTDropClaimConditions(address, provider, helper, storage)
	if err != nil {
		return nil, err
	}

	return &ERC721{
		token,
		drop,
		helper,
		storage,
		claimConditions,
	}, nil
}

// Get an NFT
//
// @extension: ERC721
//
// tokenId: token ID of the token to get the metadata for
//
// returns: the metadata for the NFT and its owner
//
// Example
//
// 	nft, err := contract.ERC721.Get(context.Background(), 0)
// 	owner := nft.Owner
// 	name := nft.Metadata.Name
func (erc721 *ERC721) Get(ctx context.Context, tokenId int) (*NFTMetadataOwner, error) {
	owner := "0x0000000000000000000000000000000000000000"
	if address, err := erc721.OwnerOf(ctx, tokenId); err == nil {
		owner = address
	}

	if metadata, err := erc721.getTokenMetadata(ctx, tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &NFTMetadataOwner{
			Metadata: metadata,
			Owner:    owner,
		}
		return metadataOwner, nil
	}
}

// Get all NFTs
//
// @extension: ERC721Supply | ERC721Enumerable
//
// returns: the metadata of all the NFTs on this contract
//
// Example
//
//	nfts, err := contract.ERC721.GetAll(context.Background())
//	ownerOne := nfts[0].Owner
//	nameOne := nfts[0].Metadata.Name
func (erc721 *ERC721) GetAll(ctx context.Context) ([]*NFTMetadataOwner, error) {
	if totalCount, err := erc721.GetTotalCount(ctx); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}
		for i := 0; i < totalCount; i++ {
			tokenIds = append(tokenIds, big.NewInt(int64(i)))
		}
		return erc721.fetchNFTsByTokenId(ctx, tokenIds)
	}
}

// Get the total number of NFTs
//
// @extension: ERC721ClaimCustom | ERC721ClaimPhasesV2 | ERC721ClaimConditionsV2
//
// returns: the total number of NFTs on this contract
//
// Example
//
// 	totalCount, err := contract.ERC721.GetTotalCount(context.Background())
func (erc721 *ERC721) GetTotalCount(ctx context.Context) (int, error) {
	count, err := erc721.token.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	return int(count.Int64()), nil
}

// Get all claimed NFTs
//
// returns: a list of the metadatas of the claimed NFTs
//
// Example
//
//	claimedNfts, err := contract.ERC721.GetAllClaimed(context.Background())
//	firstOwner := claimedNfts[0].Owner
func (erc721 *ERC721) GetAllClaimed(ctx context.Context) ([]*NFTMetadataOwner, error) {
	if maxId, err := erc721.drop.NextTokenIdToClaim(&bind.CallOpts{Context: ctx}); err != nil {
		return nil, err
	} else {
		nfts := []*NFTMetadataOwner{}

		for i := 0; i < int(maxId.Int64()); i++ {
			if nft, err := erc721.Get(ctx, i); err == nil {
				nfts = append(nfts, nft)
			}
		}

		return nfts, nil
	}
}

// Get all unclaimed NFTs
//
// returns: a list of the metadatas of the unclaimed NFTs
//
// Example
//
//	unclaimedNfts, err := contract.ERC721.GetAllUnclaimed(context.Background())
//	firstNftName := unclaimedNfts[0].Name
func (erc721 *ERC721) GetAllUnclaimed(ctx context.Context) ([]*NFTMetadata, error) {
	maxId, err := erc721.drop.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	unmintedId, err := erc721.drop.NextTokenIdToClaim(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	nfts := []*NFTMetadata{}
	for i := int(unmintedId.Int64()); i < int(maxId.Int64()); i++ {
		if nft, err := erc721.getTokenMetadata(ctx, i); err == nil {
			nfts = append(nfts, nft)
		}
	}

	return nfts, nil
}

// Get the number of claimed NFTs
//
// @extension: ERC721ClaimCustom | ERC721ClaimPhasesV2 | ERC721ClaimConditionsV2
//
// Example
//
// 	totalClaimed, err := contract.ERC721.TotalClaimedSupply(context.Background())
func (erc721 *ERC721) TotalClaimedSupply(ctx context.Context) (int, error) {
	claimed, err := erc721.drop.NextTokenIdToClaim(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	return int(claimed.Int64()), nil
}

// Get the number of unclaimed NFTs
//
// @extension: ERC721ClaimCustom | ERC721ClaimPhasesV2 | ERC721ClaimConditionsV2
//
// Example
//
// 	totalUnclaimed, err := contract.ERC721.TotalUnclaimedSupply(context.Background())
func (erc721 *ERC721) TotalUnclaimedSupply(ctx context.Context) (int, error) {
	claimed, err := erc721.drop.NextTokenIdToClaim(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	total, err := erc721.drop.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	unclaimed := big.NewInt(0).Sub(total, claimed)
	return int(unclaimed.Int64()), nil
}

// Get the owner of an NFT
//
// @extension: ERC721
//
// tokenId: the token ID of the NFT to get the owner of
//
// returns: the owner of the NFT
//
// Example
//
// 	tokenId := 0
// 	owner, err := contract.ERC721.OwnerOf(context.Background(), tokenId)
func (erc721 *ERC721) OwnerOf(ctx context.Context, tokenId int) (string, error) {
	if address, err := erc721.token.OwnerOf(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(tokenId))); err != nil {
		return "", err
	} else {
		return address.String(), nil
	}
}

// Get the total number of NFTs
//
// @extension: ERC721
//
// returns: the supply of NFTs on this contract
//
// Example
//
// 	supply, err := contract.ERC721.TotalSupply(context.Background)
func (erc721 *ERC721) TotalSupply(ctx context.Context) (int, error) {
	supply, err := erc721.token.TotalSupply(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return 0, err
	}

	return int(supply.Int64()), nil
}

// Get NFT balance
//
// @extension: ERC721
//
// returns: the number of NFTs on this contract owned by the connected wallet
//
// Example
//
// 	balance, err := contract.ERC721.Balance(context.Background())
func (erc721 *ERC721) Balance(ctx context.Context) (int, error) {
	return erc721.BalanceOf(ctx, erc721.helper.GetSignerAddress().String())
}

// Get NFT balance of a specific wallet
//
// @extension: ERC721
//
// address: the address of the wallet to get the NFT balance of
//
// returns: the number of NFTs on this contract owned by the specified wallet
//
// Example
//
//	address := "{{wallet_address}}"
//	balance, err := contract.ERC721.BalanceOf(context.Background(), address)
func (erc721 *ERC721) BalanceOf(ctx context.Context, address string) (int, error) {
	balance, err := erc721.token.BalanceOf(&bind.CallOpts{
		Context: ctx,
	}, common.HexToAddress(address))
	if err != nil {
		return 0, err
	}

	return int(balance.Int64()), nil
}

// Check NFT approval
//
// @extension: ERC721
//
// address: the address whose assets are to be checked
//
// operator: the address of the operator to check
//
// returns: true if the operator is approved for all operations of the assets, otherwise false
//
// Example
//
// 	owner := "{{wallet_address}}"
// 	operator := "0x..."
//
// 	isApproved, err := contract.ERC721.IsApproved(ctx, owner, operator)
func (erc721 *ERC721) IsApproved(ctx context.Context, owner string, operator string) (bool, error) {
	return erc721.token.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner), common.HexToAddress(operator))
}

func (erc721 *ERC721) GetClaimInfo(ctx context.Context, address string) (*ClaimInfo, error) {
	claimVerification, err := erc721.prepareClaim(ctx, address, 0, false)
	if err != nil {
		return nil, err
	}

	active, err := erc721.ClaimConditions.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	activeConditionIndex, err := erc721.drop.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	totalClaimedInPhase, err := erc721.drop.GetSupplyClaimedByWallet(&bind.CallOpts{Context: ctx}, activeConditionIndex, common.HexToAddress(address))
	if err != nil {
		return nil, err
	}

	remainingClaimable := big.NewInt(0).Sub(claimVerification.MaxClaimable, totalClaimedInPhase)
	if remainingClaimable.Cmp(big.NewInt(0)) < 0 {
		remainingClaimable = big.NewInt(0)
	}
	if active.AvailableSupply.Cmp(remainingClaimable) < 0 {
		remainingClaimable = active.AvailableSupply
	}

	return &ClaimInfo{
		PricePerToken:      claimVerification.Price,
		RemainingClaimable: remainingClaimable,
		CurrencyAddress:    common.HexToAddress(claimVerification.CurrencyAddress),
	}, nil
}

func (erc721 *ERC721) GetClaimIneligibilityReasons(ctx context.Context, quantity int, addressToCheck string) ([]ClaimEligibility, error) {
	reasons := []ClaimEligibility{}

	active, err := erc721.ClaimConditions.GetActive(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "!CONDITION") || strings.Contains(err.Error(), "no active mint condition") {
			reasons = append(reasons, NoClaimConditionSet)
			return reasons, nil
		}

		return reasons, err
	}

	activeConditionIndex, err := erc721.drop.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	totalClaimedInPhase, err := erc721.drop.GetSupplyClaimedByWallet(&bind.CallOpts{Context: ctx}, activeConditionIndex, common.HexToAddress(addressToCheck))
	if err != nil {
		return nil, err
	}

	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	if active.AvailableSupply.Cmp(MaxUint256) != 0 {
		if active.AvailableSupply.Cmp(big.NewInt(int64(quantity))) < 0 {
			reasons = append(reasons, NotEnoughSupply)
			return reasons, nil
		}
	}

	hasAllowlistEntry := !strings.HasPrefix(hex.EncodeToString(active.MerkleRootHash[:]), zeroAddress)
	var allowlistEntry *SnapshotEntryWithProof
	if hasAllowlistEntry {
		allowlistEntry, err = erc721.ClaimConditions.GetClaimerProofs(ctx, addressToCheck)
		if err != nil {
			return reasons, err
		}

		if allowlistEntry != nil {
			claimVerification, err := erc721.prepareClaim(
				ctx,
				addressToCheck,
				quantity,
				false,
			)
			if err != nil {
				return reasons, err
			}

			if (active.MaxClaimablePerWallet.Cmp(big.NewInt(0)) == 0 &&
				claimVerification.MaxClaimable.Cmp(MaxUint256) == 0) ||
				claimVerification.MaxClaimable.Cmp(big.NewInt(0)) == 0 {
				reasons = append(reasons, AddressNotAllowed)
				return reasons, nil
			} else if totalClaimedInPhase.Add(totalClaimedInPhase, big.NewInt(int64(quantity))).Cmp(claimVerification.MaxClaimable) > 0 {
				reasons = append(reasons, ExceedsMaxClaimable)
				return reasons, nil
			}

			activeConditionIndex, err := erc721.drop.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
			if err != nil {
				return reasons, err
			}

			proof := abi.IDropAllowlistProof{
				Proof:                  claimVerification.Proofs,
				QuantityLimitPerWallet: claimVerification.MaxClaimable,
				PricePerToken:          claimVerification.PriceInProof,
				Currency:               common.HexToAddress(claimVerification.CurrencyAddressInProof),
			}

			isValid, err := erc721.drop.VerifyClaim(
				&bind.CallOpts{Context: ctx},
				activeConditionIndex,
				common.HexToAddress(addressToCheck),
				big.NewInt(int64(quantity)),
				common.HexToAddress(claimVerification.CurrencyAddress),
				claimVerification.Price,
				proof,
			)

			if err != nil || !isValid {
				reasons = append(reasons, AddressNotAllowed)
				return reasons, nil
			}
		}
	}

	if !hasAllowlistEntry || allowlistEntry == nil {
		if active.MaxClaimablePerWallet.Cmp(big.NewInt(0)) == 0 {
			reasons = append(reasons, AddressNotAllowed)
			return reasons, nil
		} else {
			if totalClaimedInPhase.Add(totalClaimedInPhase, big.NewInt(int64(quantity))).Cmp(active.MaxClaimablePerWallet) > 0 {
				reasons = append(reasons, ExceedsMaxClaimable)
				return reasons, nil
			}
		}
	}

	totalPrice := active.Price.Mul(active.Price, big.NewInt(int64(quantity)))
	if isNativeToken(active.CurrencyAddress) {
		balance, err := erc721.helper.GetProvider().BalanceAt(ctx, common.HexToAddress(addressToCheck), nil)
		if err != nil {
			return reasons, err
		}

		if balance.Cmp(totalPrice) < 0 {
			reasons = append(reasons, InsufficientBalance)
			return reasons, nil
		}
	} else {
		provider := erc721.helper.GetProvider()
		erc20, err := abi.NewIERC20(common.HexToAddress(active.CurrencyAddress), provider)
		if err != nil {
			return reasons, err
		}

		balance, err := erc20.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(addressToCheck))
		if err != nil {
			return reasons, err
		}

		if balance.Cmp(totalPrice) < 0 {
			reasons = append(reasons, InsufficientBalance)
			return reasons, nil
		}
	}

	return reasons, nil
}

// Transfer an NFT
//
// @extension: ERC721
//
// to: wallet address to transfer the tokens to
//
// tokenId: the token ID of the NFT to transfer
//
// returns: the transaction of the NFT transfer
//
// Example
//
//	to := "0x..."
//	tokenId := 0
//
//	tx, err := contract.ERC721.Transfer(context.Background(), to, tokenId)
func (erc721 *ERC721) Transfer(ctx context.Context, to string, tokenId int) (*types.Transaction, error) {
	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.token.SafeTransferFrom(txOpts, erc721.helper.GetSignerAddress(), common.HexToAddress(to), big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.helper.AwaitTx(ctx, tx.Hash())
	}
}

// Burna an NFT
//
// @extension: ERC721Burnable
//
// tokenId: tokenID of the token to burn
//
// returns: the transaction receipt of the burn
//
// Example
//
//	tokenId := 0
//	tx, err := contract.ERC721.Burn(context.Background(), tokenId)
func (erc721 *ERC721) Burn(ctx context.Context, tokenId int) (*types.Transaction, error) {
	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.token.Burn(txOpts, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.helper.AwaitTx(ctx, tx.Hash())
	}
}

// Set approval for all NFTs
//
// @extension: ERC721
//
// address: the address whose assets are to be approved
//
// operator: the address of the operator to set the approval for
//
// approved: true if the operator is approved for all operations of the assets, otherwise false
//
// returns: the transaction receipt of the approval
//
// Example
//
// 	operator := "{{wallet_address}}"
// 	approved := true
//
// 	tx, err := contract.ERC721.SetApprovalForAll(context.Background(), operator, approved)
func (erc721 *ERC721) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error) {
	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.token.SetApprovalForAll(txOpts, common.HexToAddress(operator), approved); err != nil {
		return nil, err
	} else {
		return erc721.helper.AwaitTx(ctx, tx.Hash())
	}
}

// Set approval for a specific NFT
//
// @extension: ERC721
//
// operator: the address of the operator to approve
//
// tokenId: the token ID of the NFT to approve
//
// returns: the transaction receipt of the approval
//
// Example
//
// 	operator := "{{wallet_address}}"
// 	approved := "0x..."
// 	tokenId := 0
//
// 	tx, err := contract.ERC721.SetApprovalForToken(context.Background(), operator, approved, tokenId)
func (erc721 *ERC721) SetApprovalForToken(ctx context.Context, operator string, tokenId int) (*types.Transaction, error) {
	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc721.token.Approve(txOpts, common.HexToAddress(operator), big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		return erc721.helper.AwaitTx(ctx, tx.Hash())
	}
}


// Mint an NFT
//
// @extension: ERC721Mintable
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	image, err := os.Open("path/to/image.jpg")
//	defer image.Close()
//
//	metadata := &thirdweb.NFTMetadataInput{
//		Name: "Cool NFT",
//		Description: "This is a cool NFT",
//		Image: image,
//	}
//
// 	tx, err := contract.ERC721.Mint(context.Background(), metadata)
func (erc721 *ERC721) Mint(ctx context.Context, metadata *NFTMetadataInput) (*types.Transaction, error) {
	address := erc721.helper.GetSignerAddress().String()
	return erc721.MintTo(ctx, address, metadata)
}

// Mint an NFT to a specific wallet
//
// @extension: ERC721Mintable
//
// address: the wallet address to mint to
//
// metadata: metadata of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	image, err := os.Open("path/to/image.jpg")
//	defer image.Close()
//
//	metadata := &thirdweb.NFTMetadataInput{
//		Name: "Cool NFT",
//		Description: "This is a cool NFT",
//		Image: image,
//	}
//
//	tx, err := contract.ERC721.MintTo(context.Background(), "{{wallet_address}}", metadata)
func (erc721 *ERC721) MintTo(ctx context.Context, address string, metadata *NFTMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(ctx, metadata, erc721.storage)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc721.token.MintTo(
		txOpts,
		common.HexToAddress(address),
		uri,
	)
	if err != nil {
		return nil, err
	}

	return erc721.helper.AwaitTx(ctx, tx.Hash())
}

// Mint many NFTs
//
// @extension: ERC721BatchMintable
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	metadatas := []*thirdweb.NFTMetadataInput{
//		&thirdweb.NFTMetadataInput{
//			Name: "Cool NFT",
//			Description: "This is a cool NFT",
//		}
//		&thirdweb.NFTMetadataInput{
//			Name: "Cool NFT 2",
//			Description: "This is also a cool NFT",
//		}
//	}
//
//	tx, err := contract.ERC721.MintBatchTo(context.Background(), metadatas)
func (erc721 *ERC721) MintBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	address := erc721.helper.GetSignerAddress().String()
	return erc721.MintBatchTo(ctx, address, metadatas)
}

// Mint many NFTs to a specific wallet
//
// @extension: ERC721BatchMintable
//
// to: the wallet address to mint to
//
// metadatas: list of metadata of the NFTs to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	metadatas := []*thirdweb.NFTMetadataInput{
//		&thirdweb.NFTMetadataInput{
//			Name: "Cool NFT",
//			Description: "This is a cool NFT",
//		}
//		&thirdweb.NFTMetadataInput{
//			Name: "Cool NFT 2",
//			Description: "This is also a cool NFT",
//		}
//	}
//
//	tx, err := contract.ERC721.MintBatchTo(context.Background(), "{{wallet_address}}", metadatas)
func (erc721 *ERC721) MintBatchTo(ctx context.Context, address string, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	uris, err := uploadOrExtractUris(ctx, metadatas, erc721.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	for _, uri := range uris {
		txOpts, err := erc721.helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := erc721.token.MintTo(
			txOpts,
			common.HexToAddress(address),
			uri,
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc721.token.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return erc721.helper.AwaitTx(ctx, tx.Hash())
}

// Lazy mint NFTs
//
// @extension: ERC721LazyMintable
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
//	tx, err := contract.ERC721.CreateBatch(context.Background(), metadatas)
func (erc721 *ERC721) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	startNumber, err := erc721.drop.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	fileStartNumber := int(startNumber.Int64())

	contractAddress := erc721.helper.getAddress().String()
	signerAddress := erc721.helper.GetSignerAddress().String()

	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}

	dataToUpload := []map[string]interface{}{}
	if err := mapstructure.Decode(data, &dataToUpload); err != nil {
		return nil, err
	}

	batch, err := erc721.storage.UploadBatch(
		ctx,
		dataToUpload,
		fileStartNumber,
		contractAddress,
		signerAddress,
	)

	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc721.drop.LazyMint(
		txOpts,
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return erc721.helper.AwaitTx(ctx, tx.Hash())
}

// Claim an NFT
//
// @extension: ERC721ClaimCustom | ERC721ClaimPhasesV2 | ERC721ClaimConditionsV2
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
//
// Example
//
//	quantity = 1
//
//	tx, err := contract.ERC721.Claim(context.Background(), quantity)
func (erc721 *ERC721) Claim(ctx context.Context, quantity int) (*types.Transaction, error) {
	address := erc721.helper.GetSignerAddress().String()
	return erc721.ClaimTo(ctx, address, quantity)
}

// Claim NFTs to a specific wallet
//
// @extension: ERC721ClaimCustom | ERC721ClaimPhasesV2 | ERC721ClaimConditionsV2
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
//	tx, err := contract.ERC721.ClaimTo(context.Background(), address, quantity)
func (erc721 *ERC721) ClaimTo(ctx context.Context, destinationAddress string, quantity int) (*types.Transaction, error) {
	addressToClaim := erc721.helper.GetSignerAddress().Hex()

	claimVerification, err := erc721.prepareClaim(ctx, addressToClaim, quantity, true)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc721.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	txOpts.Value = claimVerification.Value

	proof := abi.IDropAllowlistProof{
		Proof:                  claimVerification.Proofs,
		QuantityLimitPerWallet: claimVerification.MaxClaimable,
		PricePerToken:          claimVerification.PriceInProof,
		Currency:               common.HexToAddress(claimVerification.CurrencyAddressInProof),
	}

	tx, err := erc721.drop.Claim(
		txOpts,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(claimVerification.CurrencyAddress),
		claimVerification.Price,
		proof,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return erc721.helper.AwaitTx(ctx, tx.Hash())
}

func (erc721 *ERC721) GetClaimArguments(
	ctx context.Context,
	destinationAddress string,
	quantity int,
) (
	*ClaimArguments,
	error,
) {
	claimVerification, err := erc721.prepareClaim(ctx, destinationAddress, quantity, false)
	if err != nil {
		return nil, err
	}

	proof := abi.IDropAllowlistProof{
		Proof:                  claimVerification.Proofs,
		QuantityLimitPerWallet: claimVerification.MaxClaimable,
		PricePerToken:          claimVerification.PriceInProof,
		Currency:               common.HexToAddress(claimVerification.CurrencyAddressInProof),
	}

	return &ClaimArguments{
		claimVerification.Value,
		common.HexToAddress(destinationAddress),
		big.NewInt(int64(quantity)),
		common.HexToAddress(claimVerification.CurrencyAddress),
		claimVerification.Price,
		proof,
		[]byte{},
	}, nil
}

func (erc721 *ERC721) getTokenMetadata(ctx context.Context, tokenId int) (*NFTMetadata, error) {
	if uri, err := erc721.token.TokenURI(&bind.CallOpts{
		Context: ctx,
	}, big.NewInt(int64(tokenId))); err != nil {
		return nil, err
	} else {
		if nft, err := fetchTokenMetadata(ctx, tokenId, uri, erc721.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}

func (erc721 *ERC721) fetchNFTsByTokenId(ctx context.Context, tokenIds []*big.Int) ([]*NFTMetadataOwner, error) {
	total := len(tokenIds)

	ch := make(chan *NFTResult)
	// fetch all nfts in parallel
	for i := 0; i < total; i++ {
		go func(id int) {
			if nft, err := erc721.Get(ctx, id); err == nil {
				ch <- &NFTResult{nft, nil}
			} else {
				fmt.Println(err)
				ch <- &NFTResult{nil, err}
			}
		}(i)
	}
	// wait for all goroutines to emit
	results := make([]*NFTResult, total)
	for i := range results {
		results[i] = <-ch
	}
	// filter out errors
	nfts := []*NFTMetadataOwner{}
	for _, res := range results {
		if res.nft != nil {
			nfts = append(nfts, res.nft)
		}
	}
	// Sort by ID
	sort.SliceStable(nfts, func(i, j int) bool {
		return nfts[i].Metadata.Id.Cmp(nfts[j].Metadata.Id) < 0
	})
	return nfts, nil
}


func (erc721 *ERC721) prepareClaim(ctx context.Context, addressToClaim string, quantity int, handleApproval bool) (*ClaimVerification, error) {
	active, err := erc721.ClaimConditions.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	merkleMetadata, err := erc721.ClaimConditions.getMerkleMetadata(ctx)
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		ctx,
		addressToClaim,
		quantity,
		active,
		merkleMetadata,
		erc721.helper,
		erc721.storage,
	)
	if err != nil {
		return nil, err
	}

	// Handle approval for ERC20
	if handleApproval {
		MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)

		var pricePerToken *big.Int
		if claimVerification.Price.Cmp(MaxUint256) == 0 {
			pricePerToken = active.Price
		} else {
			pricePerToken = claimVerification.Price
		}

		var currencyAddress string
		if claimVerification.CurrencyAddress != zeroAddress {
			currencyAddress = claimVerification.CurrencyAddress
		} else {
			currencyAddress = active.CurrencyAddress
		}

		if pricePerToken.Cmp(big.NewInt(0)) > 0 {
			if !isNativeToken(currencyAddress) {
				err := approveErc20Allowance(
					ctx,
					erc721.helper,
					currencyAddress,
					pricePerToken,
					quantity,
				)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return claimVerification, nil
}
