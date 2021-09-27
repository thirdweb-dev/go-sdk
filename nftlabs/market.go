package nftlabs

import (
	"context"
	"crypto/ecdsa"
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

type MarketSdk interface {
	CommonModule
	GetListing(listingId *big.Int) (Listing, error)
	GetAll() ([]Listing, error)
	List(
		assetContract string,
		tokenId *big.Int,
		currencyContractAddress string,
		price *big.Int,
		quantity *big.Int,
		secondsUntilStart *big.Int,
		secondsUntilEnd *big.Int) (Listing, error)
	UnlistAll(listingId *big.Int) error
	Unlist(listingId *big.Int, quantity *big.Int) error
	Buy(listingId *big.Int, quantity *big.Int) error
}

type MarketSdkModule struct {
	Client *ethclient.Client
	Address string
	Options *SdkOptions
	gateway Gateway
	caller *abi.MarketCaller
	transactor *abi.MarketTransactor
	filterer *abi.MarketFilterer

	privateKey *ecdsa.PrivateKey
	rawPrivateKey string
	signerAddress common.Address
}

func NewMarketSdkModule(client *ethclient.Client, address string, opt *SdkOptions) (*MarketSdkModule, error) {
	if opt.IpfsGatewayUrl == "" {
		opt.IpfsGatewayUrl = "https://cloudflare-ipfs.com/ipfs/"
	}

	caller, err := abi.NewMarketCaller(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	transactor, err := abi.NewMarketTransactor(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	filterer, err := abi.NewMarketFilterer(common.HexToAddress(address), client)
	if err != nil {
		return nil, err
	}

	// internally we force this gw, but could allow an override for testing
	gw := NewCloudflareGateway(opt.IpfsGatewayUrl)

	return &MarketSdkModule{
		Client: client,
		Address: address,
		Options: opt,
		gateway: gw,
		caller: caller,
		transactor: transactor,
		filterer: filterer,
	}, nil
}

func (sdk *MarketSdkModule) GetListing(listingId *big.Int) (Listing, error) {
	if result, err := sdk.caller.Listings(&bind.CallOpts{}, listingId); err != nil {
		return Listing{}, err
	} else {
		return sdk.transformResultToListing(result)
	}
}

func (sdk *MarketSdkModule) GetAll() ([]Listing, error) {
	panic("implement me")
}

func (sdk *MarketSdkModule) List(
	assetContractAddress string,
	tokenId *big.Int,
	currencyContractAddress string,
	pricePerToken *big.Int,
	quantity *big.Int,
	secondsUntilStart *big.Int,
	secondsUntilEnd *big.Int) (Listing, error) {
	if sdk.signerAddress == common.HexToAddress("0") {
		return Listing{}, &NoSignerError{typeName: "nft"}
	}

	erc165Module, err := newErc165SdkModule(sdk.Client, assetContractAddress, &SdkOptions{})
	if err != nil {
		return Listing{}, err
	}

	isERC721, err := erc165Module.module.ERC165Caller.SupportsInterface(&bind.CallOpts{}, InterfaceIdErc721)
	if err != nil {
		return  Listing{}, err
	}

	if isERC721 {
		log.Printf("Contract %v is an erc721 contract", assetContractAddress)
		return sdk.listErc721(
			assetContractAddress,
			tokenId,
			currencyContractAddress,
			pricePerToken,
			quantity,
			secondsUntilStart,
			secondsUntilEnd)
	} else {
		log.Printf("Contract %v is not a erc721 contract", assetContractAddress)
		return Listing{}, &UnsupportedFunctionError{
			typeName: "market",
			body:     "Asset must be an ERC721 contract. Other types will be supported soon.",
		}
	}

	//query := ethereum.FilterQuery{
	//	FromBlock: receipt.BlockNumber,
	//	ToBlock:   receipt.BlockNumber,
	//	Addresses: []common.Address{
	//		common.HexToAddress(sdk.Address),
	//	},
	//})

	//logs, err := sdk.Client.FilterLogs(context.Background(), query)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//var event interface{}
	//for _, log := range logs {
	//	if log.DecodeRLP()
	//}
}

func (sdk *MarketSdkModule) listErc721(
	assetContractAddress string,
	tokenId *big.Int,
	currencyContractAddress string,
	pricePerToken *big.Int,
	quantity *big.Int,
	secondsUntilStart *big.Int,
	secondsUntilEnd *big.Int) (Listing, error) {

	packAddress := common.HexToAddress(assetContractAddress)
	currencyAddress := common.HexToAddress(currencyContractAddress)

	log.Printf("Creating erc721 module, at address %v\n", assetContractAddress)
	erc721Module, err := newErc721SdkModule(sdk.Client, assetContractAddress, &SdkOptions{})
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
				// return better error describing that "Failed to Gran approval on Pack contract"
				return Listing{}, err
			}
		}
	}

	log.Printf("Caller %v has been approved from %v\n", sdk.Address, sdk.signerAddress.String())

	result, err := sdk.transactor.List(&bind.TransactOpts{
		NoSend: false,
		Signer: sdk.getSigner(),
		From: sdk.signerAddress,
	}, packAddress, tokenId, currencyAddress, pricePerToken, quantity, secondsUntilStart, secondsUntilEnd)
	if err != nil {
		return Listing{}, err
	}

	log.Printf("List call completed, result  = %v\n", result.Hash())

	_, err = sdk.Client.TransactionReceipt(context.Background(), result.Hash())
	if err != nil {
		log.Printf("Failed to lookup transaction receipt with hash %v\n", result.Hash().String())
		return Listing{}, err
	}


	panic("implement me")
}



func (sdk *MarketSdkModule) UnlistAll(listingId *big.Int) error {
	panic("implement me")
}

func (sdk *MarketSdkModule) Unlist(listingId *big.Int, quantity *big.Int) error {
	panic("implement me")
}

func (sdk *MarketSdkModule) Buy(listingId *big.Int, quantity *big.Int) error {
	panic("implement me")
}

func (sdk *MarketSdkModule) transformResultToListing(listing abi.MarketListing) (Listing, error) {
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
		log.Printf("Getting nft module at %v", listing.AssetContract)
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
		time.Unix(listing.SaleStart.Int64() * 1000, 0)
	} else {
		saleStart = nil
	}

	var saleEnd *time.Time
	// TODO: should I be doing Int64() here ??? is there data loss ???
	if listing.SaleEnd.Int64() > 0 && listing.SaleEnd.Int64() < big.MaxExp - 1 {
		time.Unix(listing.SaleEnd.Int64() * 1000, 0)
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

func (sdk *MarketSdkModule) SetPrivateKey(privateKey string) error {
	if pKey, publicAddress, err := processPrivateKey(privateKey); err != nil {
		return err
	} else {
		sdk.rawPrivateKey = privateKey
		sdk.privateKey = pKey
		sdk.signerAddress = publicAddress
	}
	return nil
}

func (sdk *MarketSdkModule) getSigner() func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
	return func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
		ctx := context.Background()
		chainId, _ := sdk.Client.ChainID(ctx)
		return types.SignTx(transaction, types.NewEIP155Signer(chainId), sdk.privateKey)
	}
}

func (sdk *MarketSdkModule) getSignerAddress() common.Address {
	if sdk.signerAddress == common.HexToAddress("0") {
		return common.HexToAddress(sdk.Address)
	} else {
		return sdk.signerAddress
	}
}
