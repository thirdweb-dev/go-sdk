package thirdweb

import (
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mitchellh/mapstructure"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// This interface is currently support by the Edition and Edition Drop contracts.
// You can access all of its functions through an Edition or Edition Drop contract instance.
type ERC1155 struct {
	token     *abi.TokenERC1155
	drop      *abi.DropERC1155
	helper    *contractHelper
	storage   storage
	ClaimConditions *EditionDropClaimConditions
}

type EditionResult struct {
	nft *EditionMetadata
	err error
}

func newERC1155(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC1155, error) {
	token, err := abi.NewTokenERC1155(address, provider)
	if err != nil {
		return nil, err
	} 

	drop, err := abi.NewDropERC1155(address, provider)
	if err != nil {
		return nil, err
	}
	
	helper, err := newContractHelper(address, provider, privateKey)
	if err != nil {
		return nil, err
	} 

	claimConditions, err := newEditionDropClaimConditions(address, provider, helper, storage)
	if err != nil {
		return nil, err
	}
	
	return &ERC1155{
		token,
		drop,
		helper,
		storage,
		claimConditions,
	}, nil
	
}

// Get an NFT
//
// @extension: ERC1155
//
// tokenId: token ID of the token to get the metadata for
//
// returns: the metadata for the NFT and its supply
//
// Example
//
// 	nft, err := contract.Get(context.Background(), 0)
// 	supply := nft.Supply
// 	name := nft.Metadata.Name
func (erc1155 *ERC1155) Get(ctx context.Context, tokenId int) (*EditionMetadata, error) {
	supply := 0
	if totalSupply, err := erc1155.token.TotalSupply(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId))); err == nil {
		supply = int(totalSupply.Int64())
	}

	if metadata, err := erc1155.getTokenMetadata(ctx, tokenId); err != nil {
		return nil, err
	} else {
		metadataOwner := &EditionMetadata{
			Metadata: metadata,
			Supply:   supply,
		}
		return metadataOwner, nil
	}
}

// Get all NFTs
//
// @extension: ERC1155
//
// returns: the metadatas and supplies of all the NFTs on this contract
//
// Example
//
//	nfts, err := contract.GetAll(context.Background())
//	supplyOne := nfts[0].Supply
//	nameOne := nfts[0].Metadata.Name
func (erc1155 *ERC1155) GetAll(ctx context.Context) ([]*EditionMetadata, error) {
	if totalCount, err := erc1155.GetTotalCount(ctx); err != nil {
		return nil, err
	} else {
		tokenIds := []*big.Int{}
		for i := 0; i < totalCount; i++ {
			tokenIds = append(tokenIds, big.NewInt(int64(i)))
		}
		return fetchEditionsByTokenId(ctx, erc1155, tokenIds)
	}
}

// Get the total number of NFTs
//
// @extension: ERC1155Enumerable
//
// returns: the total number of NFTs on this contract
//
// Example
//
// 	totalCount, err := contract.GetTotalCount(context.Background())
func (erc1155 *ERC1155) GetTotalCount(ctx context.Context) (int, error) {
	count, err := erc1155.token.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return 0, err
	}

	return int(count.Int64()), nil
}

// Get owned NFTs
//
// @extension: ERC1155Enumerable
//
// address: the address of the owner of the NFTs
//
// returns: the metadatas and supplies of all the NFTs owned by the address
//
// Example
//
//	owner := "{{wallet_address}}"
//	nfts, err := contract.GetOwned(context.Background(), owner)
//	name := nfts[0].Metadata.Name
func (erc1155 *ERC1155) GetOwned(ctx context.Context, address string) ([]*EditionMetadataOwner, error) {
	if address == "" {
		address = erc1155.helper.GetSignerAddress().String()
	}

	maxId, err := erc1155.token.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}

	owners := []common.Address{}
	ids := []*big.Int{}
	for i := 0; i < int(maxId.Int64()); i++ {
		owners = append(owners, common.HexToAddress(address))
		ids = append(ids, big.NewInt(int64(i)))
	}

	balances, err := erc1155.token.BalanceOfBatch(&bind.CallOpts{Context: ctx}, owners, ids)
	if err != nil {
		return nil, err
	}

	metadataOwners := []*EditionMetadataOwner{}
	metadatas, err := fetchEditionsByTokenId(ctx, erc1155, ids)
	if err != nil {
		return nil, err
	}
	for index, balance := range balances {
		metadata := metadatas[index]
		if err == nil {
			metadataOwner := &EditionMetadataOwner{
				Metadata:      metadata.Metadata,
				Supply:        metadata.Supply,
				Owner:         address,
				QuantityOwned: int(balance.Int64()),
			}
			metadataOwners = append(metadataOwners, metadataOwner)
		}
	}

	return metadataOwners, nil
}

// Get the total supply of an NFT
//
// @extension: ERC1155
//
// tokenId: the token ID to check the total supply of
//
// returns: the supply of NFTs on the specified token ID
//
// Example
//
// 	tokenId := 0
//
// 	totalSupply, err := contract.TotalSupply(context.Background, tokenId)
func (erc1155 *ERC1155) TotalSupply(ctx context.Context, tokenId int) (int, error) {
	supply, err := erc1155.token.TotalSupply(&bind.CallOpts{Context: ctx}, big.NewInt(int64(tokenId)))
	if err != nil {
		return 0, err
	}

	return int(supply.Int64()), nil
}

// Get NFT balance
//
// @extension: ERC1155
//
// tokenId: the token ID of a specific token to check the balance of
//
// returns: the number of NFTs of the specified token ID owned by the connected wallet
func (erc1155 *ERC1155) Balance(ctx context.Context, tokenId int) (int, error) {
	address := erc1155.helper.GetSignerAddress().String()
	return erc1155.BalanceOf(ctx, address, tokenId)
}

// Get NFT balance of a specific wallet
//
// @extension: ERC1155
//
// address: the address of the wallet to get the NFT balance of
//
// returns: the number of NFTs of the specified token ID owned by the specified wallet
//
// Example
//
//	address := "{{wallet_address}}"
//	tokenId := 0
//	balance, err := contract.BalanceOf(context.Background(), address, tokenId)
func (erc1155 *ERC1155) BalanceOf(ctx context.Context, address string, tokenId int) (int, error) {
	balance, err := erc1155.token.BalanceOf(&bind.CallOpts{Context: ctx}, common.HexToAddress(address), big.NewInt(int64(tokenId)))
	if err != nil {
		return 0, err
	}

	return int(balance.Int64()), nil
}

// Check NFT approval
//
// @extension: ERC1155
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
// 	isApproved, err := contract.IsApproved(context.Background, owner, operator)
func (erc1155 *ERC1155) IsApproved(ctx context.Context, owner string, operator string) (bool, error) {
	return erc1155.token.IsApprovedForAll(&bind.CallOpts{Context: ctx}, common.HexToAddress(owner), common.HexToAddress(operator))
}

// Transfer NFTs
//
// @extension: ERC1155
//
// to: wallet address to transfer the tokens to
//
// tokenId: the token ID of the NFT to transfer
//
// amount: number of NFTs of the token ID to transfer
//
// returns: the transaction of the NFT transfer
//
// Example
//
//	to := "0x..."
//	tokenId := 0
//	amount := 1
//
//	tx, err := contract.Transfer(context.Background(), to, tokenId, amount)
func (erc1155 *ERC1155) Transfer(ctx context.Context, to string, tokenId int, amount int) (*types.Transaction, error) {
	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc1155.token.SafeTransferFrom(
		txOpts,
		erc1155.helper.GetSignerAddress(),
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(amount)),
		[]byte{},
	); err != nil {
		return nil, err
	} else {
		return erc1155.helper.AwaitTx(ctx, tx.Hash())
	}
}

// Burn NFTs
//
// @extension: ERC1155Burnable
//
// tokenId: tokenID of the token to burn
//
// amount: number of NFTs of the token ID to burn
//
// returns: the transaction receipt of the burn
//
// Example
//
//	tokenId := 0
//	amount := 1
//	tx, err := contract.Burn(context.Background(), tokenId, amount)
func (erc1155 *ERC1155) Burn(ctx context.Context, tokenId int, amount int) (*types.Transaction, error) {
	address := erc1155.helper.GetSignerAddress()
	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc1155.token.Burn(
		txOpts,
		address,
		big.NewInt(int64(tokenId)),
		big.NewInt(int64(amount)),
	); err != nil {
		return nil, err
	} else {
		return erc1155.helper.AwaitTx(ctx, tx.Hash())
	}
}

// Set approval for all NFTs
//
// @extension: ERC1155
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
// 	tx, err := contract.SetApprovalForAll(context.Background(), operator, approved)
func (erc1155 *ERC1155) SetApprovalForAll(ctx context.Context, operator string, approved bool) (*types.Transaction, error) {
	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if tx, err := erc1155.token.SetApprovalForAll(
		txOpts,
		common.HexToAddress(operator),
		approved,
	); err != nil {
		return nil, err
	} else {
		return erc1155.helper.AwaitTx(ctx, tx.Hash())
	}
}

// Mint an NFT
//
// @extension: ERC1155Mintable
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	image, err := os.Open("path/to/image.jpg")
// 	defer image.Close()
//
// 	metadataWithSupply := &thirdweb.EditionMetadataInput{
// 		Metadata: &thirdweb.NFTMetadataInput{
// 			Name: "Cool NFT",
//			Description: "This is a cool NFT",
//			Image: image,
//		},
//		Supply: 100,
// 	}
//
// 	tx, err := contract.Mint(context.Background(), metadataWithSupply)
func (erc1155 *ERC1155) Mint(ctx context.Context, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	address := erc1155.helper.GetSignerAddress().String()
	return erc1155.MintTo(ctx, address, metadataWithSupply)
}

// Mint an NFT to a specific wallet
//
// @extension: ERC1155Mintable
//
// address: the wallet address to mint the NFT to
//
// metadataWithSupply: nft metadata with supply of the NFT to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	image, err := os.Open("path/to/image.jpg")
// 	defer image.Close()
//
// 	metadataWithSupply := &thirdweb.EditionMetadataInput{
// 		Metadata: &thirdweb.NFTMetadataInput{
// 			Name: "Cool NFT",
//			Description: "This is a cool NFT",
//			Image: image,
//		},
//		Supply: 100,
// 	}
//
// 	tx, err := contract.MintTo(context.Background(), "{{wallet_address}}", metadataWithSupply)
func (erc1155 *ERC1155) MintTo(ctx context.Context, address string, metadataWithSupply *EditionMetadataInput) (*types.Transaction, error) {
	uri, err := uploadOrExtractUri(ctx, metadataWithSupply.Metadata, erc1155.storage)
	if err != nil {
		return nil, err
	}

	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc1155.token.MintTo(
		txOpts,
		common.HexToAddress(address),
		MaxUint256,
		uri,
		big.NewInt(int64(metadataWithSupply.Supply)),
	)
	if err != nil {
		return nil, err
	}

	return erc1155.helper.AwaitTx(ctx, tx.Hash())
}

// Mint additionaly supply of an NFT
//
// @extension: ERC1155Mintable
//
// tokenId: token ID to mint additional supply of
//
// additionalSupply: additional supply to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	tokenId := 0
// 	additionalSupply := 100
//
// 	tx, err := contract.MintAdditionalSupply(context.Background(), tokenId, additionalSupply)
func (erc1155 *ERC1155) MintAdditionalSupply(ctx context.Context, tokenId int, additionalSupply int) (*types.Transaction, error) {
	address := erc1155.helper.GetSignerAddress().String()
	return erc1155.MintAdditionalSupplyTo(ctx, address, tokenId, additionalSupply)
}

// Mint additional supply of an NFT to a specific wallet
//
// @extension: ERC1155Mintable
//
// to: address of the wallet to mint NFTs to
//
// tokenId: token Id to mint additional supply of
//
// additionalySupply: additional supply to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
// 	to := "{{wallet_address}}"
// 	tokenId := 0
// 	additionalSupply := 100
//
// 	tx, err := contract.MintAdditionalSupplyTo(context.Background(), to, tokenId, additionalSupply)
func (erc1155 *ERC1155) MintAdditionalSupplyTo(ctx context.Context, to string, tokenId int, additionalSupply int) (*types.Transaction, error) {
	metadata, err := erc1155.getTokenMetadata(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc1155.token.MintTo(
		txOpts,
		common.HexToAddress(to),
		big.NewInt(int64(tokenId)),
		metadata.Uri,
		big.NewInt(int64(additionalSupply)),
	)
	if err != nil {
		return nil, err
	}

	return erc1155.helper.AwaitTx(ctx, tx.Hash())
}

// Mint many NFTs
//
// @extension: ERC1155BatchMintable
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
//	tx, err := contract.MintBatch(context.Background(), metadatasWithSupply)
func (erc1155 *ERC1155) MintBatch(ctx context.Context, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	return erc1155.MintBatchTo(ctx, erc1155.helper.GetSignerAddress().String(), metadatasWithSupply)
}

// Mint many NFTs to a specific wallet
//
// @extension: ERC1155BatchMintable
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
func (erc1155 *ERC1155) MintBatchTo(ctx context.Context, to string, metadatasWithSupply []*EditionMetadataInput) (*types.Transaction, error) {
	metadatas := []*NFTMetadataInput{}
	for _, metadataWithSupply := range metadatasWithSupply {
		metadatas = append(metadatas, metadataWithSupply.Metadata)
	}

	supplies := []int{}
	for _, metadataWithSupply := range metadatasWithSupply {
		supplies = append(supplies, metadataWithSupply.Supply)
	}

	uris, err := uploadOrExtractUris(ctx, metadatas, erc1155.storage)
	if err != nil {
		return nil, err
	}

	encoded := [][]byte{}
	MaxUint256 := new(big.Int).Sub(new(big.Int).Lsh(common.Big1, 256), common.Big1)
	for index, uri := range uris {
		txOpts, err := erc1155.helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := erc1155.token.MintTo(
			txOpts,
			common.HexToAddress(to),
			MaxUint256,
			uri,
			big.NewInt(int64(supplies[index])),
		)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc1155.token.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return erc1155.helper.AwaitTx(ctx, tx.Hash())
}

// Lazy mint NFTs
//
// @extension: ERC1155LazyMintableV2
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
//	tx, err := contract.CreateBatch(context.Background(), metadatasWithSupply)
func (erc1155 *ERC1155) CreateBatch(ctx context.Context, metadatas []*NFTMetadataInput) (*types.Transaction, error) {
	startNumber, err := erc1155.drop.NextTokenIdToMint(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, err
	}
	fileStartNumber := int(startNumber.Int64())

	contractAddress := erc1155.helper.getAddress().String()
	signerAddress := erc1155.helper.GetSignerAddress().String()

	data := []interface{}{}
	for _, metadata := range metadatas {
		data = append(data, metadata)
	}
	dataToUpload := []map[string]interface{}{}
	if err := mapstructure.Decode(data, &dataToUpload); err != nil {
		return nil, err
	}

	batch, err := erc1155.storage.UploadBatch(
		ctx,
		dataToUpload,
		fileStartNumber,
		contractAddress,
		signerAddress,
	)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := erc1155.drop.LazyMint(
		txOpts,
		big.NewInt(int64(len(batch.uris))),
		batch.baseUri,
		[]byte{},
	)
	if err != nil {
		return nil, err
	}

	return erc1155.helper.AwaitTx(ctx, tx.Hash())
}

// Claim an NFT
//
// @extension: ERC1155ClaimCustom | ERC1155ClaimPhasesV2 | ERC1155ClaimConditionsV2
//
// tokenId: the token ID of the NFT to claim
//
// quantity: the number of NFTs to claim
//
// returns: the transaction receipt of the claim
//
// Example
//
//	tokenId = 0
//	quantity = 1
//
//	tx, err := contract.ClaimTo(context.Background(), tokenId, quantity)
func (erc1155 *ERC1155) Claim(ctx context.Context, tokenId int, quantity int) (*types.Transaction, error) {
	address := erc1155.helper.GetSignerAddress().String()
	return erc1155.ClaimTo(ctx, address, tokenId, quantity)
}

// Claim an NFT to a specific wallet
//
// @extension: ERC1155ClaimCustom | ERC1155ClaimPhasesV2 | ERC1155ClaimConditionsV2
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
func (erc1155 *ERC1155) ClaimTo(ctx context.Context, destinationAddress string, tokenId int, quantity int) (*types.Transaction, error) {
	claimVerification, err := erc1155.prepareClaim(ctx, tokenId, quantity)
	if err != nil {
		return nil, err
	}

	active, err := erc1155.ClaimConditions.GetActive(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	txOpts, err := erc1155.helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	txOpts.Value = claimVerification.Value

	proof := abi.IDrop1155AllowlistProof{
		Proof:                  claimVerification.Proofs,
		QuantityLimitPerWallet: claimVerification.MaxClaimable,
		PricePerToken:          claimVerification.Price,
		Currency:               common.HexToAddress(claimVerification.CurrencyAddress),
	}

	tx, err := erc1155.drop.Claim(
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

	return erc1155.helper.AwaitTx(ctx, tx.Hash())
}

func (erc1155 *ERC1155) prepareClaim(ctx context.Context, tokenId int, quantity int) (*ClaimVerification, error) {
	addressToClaim := erc1155.helper.GetSignerAddress().Hex()
	claimCondition, err := erc1155.ClaimConditions.GetActive(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	merkleMetadata, err := erc1155.ClaimConditions.GetMerkleMetadata(ctx)
	if err != nil {
		return nil, err
	}

	claimVerification, err := prepareClaim(
		ctx,
		addressToClaim,
		quantity,
		claimCondition,
		merkleMetadata,
		erc1155.helper,
		erc1155.storage,
	)
	if err != nil {
		return nil, err
	}

	return claimVerification, nil
}


func (erc1155 *ERC1155) getTokenMetadata(ctx context.Context, tokenId int) (*NFTMetadata, error) {
	if uri, err := erc1155.token.Uri(
		&bind.CallOpts{Context: ctx},
		big.NewInt(int64(tokenId)),
	); err != nil {
		return nil, &notFoundError{
			tokenId,
		}
	} else {
		if nft, err := fetchTokenMetadata(ctx, tokenId, uri, erc1155.storage); err != nil {
			return nil, err
		} else {
			return nft, nil
		}
	}
}

func fetchEditionsByTokenId(ctx context.Context, erc1155 *ERC1155, tokenIds []*big.Int) ([]*EditionMetadata, error) {
	total := len(tokenIds)

	ch := make(chan *EditionResult)
	// fetch all nfts in parallel
	for i := 0; i < total; i++ {
		go func(id int) {
			if nft, err := erc1155.Get(ctx, id); err == nil {
				ch <- &EditionResult{nft, nil}
			} else {
				fmt.Println(err)
				ch <- &EditionResult{nil, err}
			}
		}(i)
	}
	// wait for all goroutines to emit
	results := make([]*EditionResult, total)
	for i := range results {
		results[i] = <-ch
	}
	// filter out errors
	nfts := []*EditionMetadata{}
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
