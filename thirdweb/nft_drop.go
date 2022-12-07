package thirdweb

import (
	"context"
	"encoding/hex"
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
	*ERC721
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

	totalCount, err := nft.abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	tokenIds := []*big.Int{}
	for i := 0; i < int(totalCount.Int64()); i++ {
		owner, err := nft.abi.OwnerOf(&bind.CallOpts{Context: ctx}, big.NewInt(int64(i)))
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
	if maxId, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{Context: ctx}); err != nil {
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
	maxId, err := drop.Abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	unmintedId, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{Context: ctx})
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
func (drop *NFTDrop) TotalClaimedSupply(ctx context.Context) (int, error) {
	claimed, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	return int(claimed.Int64()), nil
}

// Get the total number of NFTs that have not yet been claimed.
func (drop *NFTDrop) TotalUnclaimedSupply(ctx context.Context) (int, error) {
	claimed, err := drop.Abi.NextTokenIdToClaim(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	total, err := drop.Abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	unclaimed := big.NewInt(0).Sub(total, claimed)
	return int(unclaimed.Int64()), nil
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
	claimVerification, err := drop.prepareClaim(ctx, address, 0, false)
	if err != nil {
		return nil, err
	}

	active, err := drop.ClaimConditions.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	activeConditionIndex, err := drop.Abi.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	totalClaimedInPhase, err := drop.Abi.GetSupplyClaimedByWallet(&bind.CallOpts{Context: ctx}, activeConditionIndex, common.HexToAddress(address))
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

func (drop *NFTDrop) GetClaimIneligibilityReasons(ctx context.Context, quantity int, addressToCheck string) ([]ClaimEligibility, error) {
	reasons := []ClaimEligibility{}

	active, err := drop.ClaimConditions.GetActive(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "!CONDITION") || strings.Contains(err.Error(), "no active mint condition") {
			reasons = append(reasons, NoClaimConditionSet)
			return reasons, nil
		}

		return reasons, err
	}

	activeConditionIndex, err := drop.Abi.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	totalClaimedInPhase, err := drop.Abi.GetSupplyClaimedByWallet(&bind.CallOpts{Context: ctx}, activeConditionIndex, common.HexToAddress(addressToCheck))
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
		allowlistEntry, err = drop.ClaimConditions.GetClaimerProofs(ctx, addressToCheck)
		if err != nil {
			return reasons, err
		}

		if allowlistEntry != nil {
			claimVerification, err := drop.prepareClaim(
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

			activeConditionIndex, err := drop.Abi.GetActiveClaimConditionId(&bind.CallOpts{Context: ctx})
			if err != nil {
				return reasons, err
			}

			proof := abi.IDropAllowlistProof{
				Proof:                  claimVerification.Proofs,
				QuantityLimitPerWallet: claimVerification.MaxClaimable,
				PricePerToken:          claimVerification.PriceInProof,
				Currency:               common.HexToAddress(claimVerification.CurrencyAddressInProof),
			}

			isValid, err := drop.Abi.VerifyClaim(
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
		balance, err := drop.Helper.GetProvider().BalanceAt(ctx, common.HexToAddress(addressToCheck), nil)
		if err != nil {
			return reasons, err
		}

		if balance.Cmp(totalPrice) < 0 {
			reasons = append(reasons, InsufficientBalance)
			return reasons, nil
		}
	} else {
		provider := drop.Helper.GetProvider()
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
	startNumber, err := drop.Abi.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
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
	if err := mapstructure.Decode(data, &dataToUpload); err != nil {
		return nil, err
	}

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
	tx, err := drop.Abi.LazyMint(
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
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
func (drop *NFTDrop) Claim(ctx context.Context, quantity int) (*types.Transaction, error) {
	address := drop.Helper.GetSignerAddress().String()
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
	addressToClaim := drop.helper.GetSignerAddress().Hex()

	claimVerification, err := drop.prepareClaim(ctx, addressToClaim, quantity, true)
	if err != nil {
		return nil, err
	}

	txOpts, err := drop.Helper.GetTxOptions(ctx)
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

	tx, err := drop.Abi.Claim(
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

	return drop.Helper.AwaitTx(tx.Hash())
}

func (drop *NFTDrop) GetClaimArguments(
	ctx context.Context,
	destinationAddress string,
	quantity int,
) (
	*ClaimArguments,
	error,
) {
	claimVerification, err := drop.prepareClaim(ctx, destinationAddress, quantity, false)
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

func (drop *NFTDrop) prepareClaim(ctx context.Context, addressToClaim string, quantity int, handleApproval bool) (*ClaimVerification, error) {
	active, err := drop.ClaimConditions.GetActive(ctx)
	if err != nil {
		return nil, err
	}

	merkleMetadata, err := drop.ClaimConditions.getMerkleMetadata()
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		ctx,
		addressToClaim,
		quantity,
		active,
		merkleMetadata,
		drop.Helper,
		drop.storage,
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
					drop.Helper,
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
