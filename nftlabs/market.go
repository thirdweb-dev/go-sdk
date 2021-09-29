package nftlabs

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nftlabs/nftlabs-sdk-go/abi"
)

type Market interface {
	CommonModule
	GetListing(listingId *big.Int) (Listing, error)
	GetAll(filter ListingFilter) ([]Listing, error)
	List(args NewListingArgs) (Listing, error)
	UnlistAll(listingId *big.Int) error
	Unlist(listingId *big.Int, quantity *big.Int) error
	Buy(listingId *big.Int, quantity *big.Int) error
	GetMarketFeeBps() (*big.Int, error)
	SetMarketFeeBps(fee *big.Int) error
}

type MarketModule struct {
	Client  *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	module  *abi.Market

	privateKey    *ecdsa.PrivateKey
	rawPrivateKey string
	signerAddress common.Address
}

func NewMarketSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*MarketModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	module, err := abi.NewMarket(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	gw := NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &MarketModule{
		Client:  client,
		Address: address,
		Options: opt,
		gateway: gw,
		module:  module,
	}, nil
}

func (sdk *MarketModule) GetMarketFeeBps() (*big.Int, error) {
	return sdk.module.MarketCaller.MarketFeeBps(&bind.CallOpts{})
}

func (sdk *MarketModule) SetMarketFeeBps(fee *big.Int) error {
	if tx, err := sdk.module.SetMarketFeeBps(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, fee); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *MarketModule) GetListing(listingId *big.Int) (Listing, error) {
	if result, err := sdk.module.MarketCaller.Listings(&bind.CallOpts{}, listingId); err != nil {
		return Listing{}, err
	} else {
		return sdk.transformResultToListing(result)
	}
}

func (sdk *MarketModule) GetAll(filter ListingFilter) ([]Listing, error) {
	listings := make([]abi.MarketListing, 0)

	hasFilter := filter.TokenContract != "" || filter.TokenId != nil || filter.Seller != ""

	if !hasFilter {
		// TODO: fetch all
		result, err := sdk.module.GetAllListings(&bind.CallOpts{})
		if err != nil {
			return nil, err
		}
		for _, l := range result {
			listings = append(listings, l)
		}
	} else {
		if filter.TokenContract != "" && filter.TokenId != nil {
			result, err := sdk.module.GetListingsByAsset(&bind.CallOpts{}, common.HexToAddress(filter.TokenContract), filter.TokenId)
			if err != nil {
				return nil, err
			}
			for _, l := range result {
				listings = append(listings, l)
			}
		} else if filter.Seller != "" {
			result, err := sdk.module.GetListingsBySeller(&bind.CallOpts{}, common.HexToAddress(filter.Seller))
			if err != nil {
				return nil, err
			}
			for _, l := range result {
				listings = append(listings, l)
			}
		} else if filter.TokenContract != "" {
			result, err := sdk.module.GetListingsByAssetContract(&bind.CallOpts{}, common.HexToAddress(filter.TokenContract))
			if err != nil {
				return nil, err
			}
			for _, l := range result {
				listings = append(listings, l)
			}
		}
	}

	availableListings := make([]Listing, 0)
	for _, listing := range listings {
		if listing.Quantity.Int64() == 0 {
			continue
		}

		if !hasFilter {
			if transformed, err := sdk.transformResultToListing(listing); err != nil {
				return nil, err
			} else {
				availableListings = append(availableListings, transformed)
			}
			continue
		}

		if filter.Seller != "" && strings.ToLower(filter.Seller) != strings.ToLower(listing.Seller.String()) {
			continue
		}

		if filter.TokenContract != "" && strings.ToLower(filter.TokenContract) != strings.ToLower(listing.AssetContract.String()) {
			continue
		}

		if filter.TokenId.String() != "" && strings.ToLower(filter.TokenId.String()) != strings.ToLower(listing.TokenId.String()) {
			continue
		}

		if transformed, err := sdk.transformResultToListing(listing); err != nil {
			return nil, err
		} else {
			availableListings = append(availableListings, transformed)
		}
	}

	return availableListings, nil
}

// TODO: change args to struct
func (sdk *MarketModule) List(args NewListingArgs) (Listing, error) {
	if sdk.signerAddress == common.HexToAddress("0") {
		return Listing{}, &NoSignerError{typeName: "nft"}
	}

	erc165Module, err := newErc165SdkModule(sdk.Client, args.AssetContractAddress, &SdkOptions{})
	if err != nil {
		return Listing{}, err
	}

	isERC721, err := erc165Module.module.ERC165Caller.SupportsInterface(&bind.CallOpts{}, InterfaceIdErc721)
	if err != nil {
		return Listing{}, err
	}

	if isERC721 {
		log.Printf("Contract %v is an erc721 contract", args.AssetContractAddress)
		return sdk.listErc721(args)
	} else {
		log.Printf("Contract %v is not a erc721 contract", args.AssetContractAddress)
		return Listing{}, &UnsupportedFunctionError{
			typeName: "market",
			body:     "Asset must be an ERC721 contract. Other types will be supported soon.",
		}
	}
}

func (sdk *MarketModule) listErc721(args NewListingArgs) (Listing, error) {
	packAddress := common.HexToAddress(args.AssetContractAddress)
	currencyAddress := common.HexToAddress(args.CurrencyContractAddress)

	log.Println(args)

	log.Printf("Creating erc721 module, at address %v\n", args.AssetContractAddress)
	erc721Module, err := newErc721SdkModule(sdk.Client, args.AssetContractAddress, &SdkOptions{})
	if err != nil {
		// TODO: return better error
		return Listing{}, err
	}
	if err := erc721Module.SetPrivateKey(sdk.rawPrivateKey); err != nil {
		return Listing{}, err
	}

	log.Printf("Checking if caller (%v) is approved\n", sdk.signerAddress)
	if isApproved, err := erc721Module.module.ERC721Caller.IsApprovedForAll(&bind.CallOpts{}, sdk.signerAddress, common.HexToAddress(sdk.Address)); err != nil {
		return Listing{}, err
	} else {
		log.Printf("Caller is not approved, setting approval on marketplace %v\n", sdk.Address)
		if !isApproved {
			if _, err := erc721Module.module.ERC721Transactor.SetApprovalForAll(&bind.TransactOpts{
				NoSend: false,
				From:   sdk.signerAddress,
				Signer: erc721Module.getSigner(),
			}, common.HexToAddress(sdk.Address), true); err != nil {
				// return better error describing that "Failed to Gran approval on PackMetadata contract"
				return Listing{}, err
			}
		}
	}

	log.Printf("Caller %v has been approved from %v\n", sdk.Address, sdk.signerAddress.String())

	result, err := sdk.module.MarketTransactor.List(&bind.TransactOpts{
		NoSend:  false,
		Signer:  sdk.getSigner(),
		From:    sdk.signerAddress,
		Context: context.Background(),
	}, packAddress, args.TokenId, currencyAddress, args.Price, args.Quantity, args.RewardsPerOpen, args.SecondsUntilOpenStart, args.SecondsUntilOpenEnd)
	if err != nil {
		return Listing{}, err
	}

	log.Printf("List call completed, result  = %v\n", result.Hash())

	if err := waitForTx(sdk.Client, result.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts); err != nil {
		// TODO: return tx failed err
		return Listing{}, err
	}

	receipt, err := sdk.Client.TransactionReceipt(context.Background(), result.Hash())
	if err != nil {
		log.Printf("Failed to lookup transaction receipt with hash %v\n", result.Hash().String())
		return Listing{}, err
	}

	log.Printf("Got receipt %v for tx %v\n", receipt.TxHash, result.Hash())

	if newListing, err := sdk.getNewMarketListing(receipt.Logs); err != nil {
		return Listing{}, err
	} else {
		return sdk.transformResultToListing(*newListing)
	}
}

func (sdk *MarketModule) getNewMarketListing(logs []*types.Log) (*abi.MarketListing, error) {
	var listing abi.MarketListing
	for _, l := range logs {
		iterator, err := sdk.module.MarketFilterer.ParseNewListing(*l)
		if err != nil {
			return nil, err
		}

		if iterator.Listing.ListingId != nil {
			fmt.Printf("Listing id = %v, listing = %v\n", iterator.ListingId, iterator.Listing)
			listing = iterator.Listing
			break
		}
	}

	return &listing, nil
}

func (sdk *MarketModule) UnlistAll(listingId *big.Int) error {
	listing, err := sdk.GetListing(listingId)
	if err != nil {
		return err
	}
	return sdk.Unlist(listingId, listing.Quantity)
}

func (sdk *MarketModule) Unlist(listingId *big.Int, quantity *big.Int) error {
	if tx, err := sdk.module.Unlist(&bind.TransactOpts{
		NoSend: false,
		From:   sdk.getSignerAddress(),
		Signer: sdk.getSigner(),
	}, listingId, quantity); err != nil {
		return err
	} else {
		return waitForTx(sdk.Client, tx.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
	}
}

func (sdk *MarketModule) Buy(listingId *big.Int, quantity *big.Int) error {
	listing, err := sdk.GetListing(listingId)
	if err != nil {
		return err
	}
	totalPrice := big.Int{}
	totalPrice.Set(listing.Price)
	totalPrice.Mul(&totalPrice, quantity)

	return nil
}

func (sdk *MarketModule) transformResultToListing(listing abi.MarketListing) (Listing, error) {
	listingCurrency := listing.Currency

	var currencyMetadata *CurrencyValue
	if strings.HasPrefix(listingCurrency.String(), "0x000000000000") {
		currencyMetadata = nil
	} else {
		// TODO: this is bad, don't want to create an instance of the module every time but idk how else to get it in here
		// damages testability
		log.Printf("Getting listing currency at address %v\n", listingCurrency)
		currency, err := NewCurrencySdkModule(sdk.Client, listingCurrency.Hex())
		if err != nil {
			// TODO: return better error
			return Listing{}, err
		}

		if currencyValue, err := currency.GetValue(listing.PricePerToken); err != nil {
			// TODO: return better error
			return Listing{}, err
		} else {
			currencyMetadata = currencyValue
		}
	}

	var nftMetadata *NftMetadata
	if !strings.HasPrefix(listing.AssetContract.String(), "0x000000000000") {
		// TODO: again, bad, need to create this in the function because we don't know the nft contract when we get here
		// damages testability
		nftModule, err := NewNftSdkModule(sdk.Client, listing.AssetContract.Hex(), &SdkOptions{})
		if err != nil {
			// TODO: return better error
			return Listing{}, err
		}

		if meta, err := nftModule.Get(listing.TokenId); err != nil {
			// TODO: return better error
			return Listing{}, err
		} else {
			nftMetadata = &meta
		}
	}

	var saleStart *time.Time
	// TODO: should I be doing Int64() here ??? is there data loss ???
	if listing.SaleStart.Int64() > 0 {
		time.Unix(listing.SaleStart.Int64()*1000, 0)
	} else {
		saleStart = nil
	}

	var saleEnd *time.Time
	// TODO: should I be doing Int64() here ??? is there data loss ???
	if listing.SaleEnd.Int64() > 0 && listing.SaleEnd.Int64() < big.MaxExp-1 {
		time.Unix(listing.SaleEnd.Int64()*1000, 0)
	} else {
		saleEnd = nil
	}

	return Listing{
		Id:               listing.ListingId,
		Seller:           listing.Seller,
		TokenContract:    listing.AssetContract,
		TokenId:          listing.TokenId,
		TokenMetadata:    nftMetadata,
		Quantity:         listing.Quantity,
		CurrentContract:  listingCurrency,
		CurrencyMetadata: currencyMetadata,
		Price:            listing.PricePerToken,
		SaleStart:        saleStart,
		SaleEnd:          saleEnd,
	}, nil
}

func (sdk *MarketModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.rawPrivateKey = privateKey
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}

func (sdk *MarketModule) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}

func (sdk *MarketModule) getSignerAddress() common.Address {
	if sdk.signerAddress == common.HexToAddress("0") {
		return common.HexToAddress(sdk.Address)
	} else {
		return sdk.signerAddress
	}
}
