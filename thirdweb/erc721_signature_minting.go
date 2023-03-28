package thirdweb

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	signerTypes "github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/google/uuid"

	"github.com/thirdweb-dev/go-sdk/v2/abi"
)

// You can access this interface from the NFT Collection contract under the
// signature interface.
type ERC721SignatureMinting struct {
	abi     *abi.TokenERC721
	Helper  *contractHelper
	storage storage
}

func newERC721SignatureMinting(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC721SignatureMinting, error) {
	if contractAbi, err := abi.NewTokenERC721(address, provider); err != nil {
		return nil, err
	} else if helper, err := newContractHelper(address, provider, privateKey); err != nil {
		return nil, err
	} else {
		return &ERC721SignatureMinting{
			contractAbi,
			helper,
			storage,
		}, nil
	}
}

// Mint a token with the data in given payload.
//
// signedPayload: the payload signed by the minters private key being used to mint
//
// returns: the transaction receipt of the mint
//
// Example
//
//	// Learn more about how to craft a payload in the Generate() function
//	signedPayload, err := contract.Signature.Generate(payload)
//	tx, err := contract.Signature.Mint(context.Background(), signedPayload)
func (signature *ERC721SignatureMinting) Mint(ctx context.Context, signedPayload SignedPayload721) (*types.Transaction, error) {
	txOpts, err := signature.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	if err := setErc20Allowance(
		ctx,
		signature.Helper,
		signedPayload.Payload.Price,
		signedPayload.Payload.Currency.String(),
		txOpts,
	); err != nil {
		return nil, err
	}

	signatureBytes, err := hex.DecodeString(signedPayload.Signature[2:])
	if err != nil {
		return nil, err
	}

	tx, err := signature.abi.MintWithSignature(txOpts, signedPayload.Payload, signatureBytes)
	if err != nil {
		return nil, err
	}

	return signature.Helper.AwaitTx(ctx, tx.Hash())
}

// Mint a batch of token with the data in given payload.
//
// signedPayload: the list of payloads signed by the minters private key being used to mint
//
// returns: the transaction receipt of the batch mint
//
// Example
//
//	// Learn more about how to craft multiple payloads in the GenerateBatch() function
//	signedPayloads, err := contract.Signature.GenerateBatch(payloads)
//	tx, err := contract.Signature.MintBatch(context.Background(), signedPayloads)
func (signature *ERC721SignatureMinting) MintBatch(ctx context.Context, signedPayloads []*SignedPayload721) (*types.Transaction, error) {
	var contractPayloads []abi.ITokenERC721MintRequest
	for _, signedPayload := range signedPayloads {
		price := signedPayload.Payload.Price
		if price.Cmp(big.NewInt(0)) == 1 {
			return nil, fmt.Errorf("Can only batch free mints. For mints with a price, use the Mint() function.")
		}
		contractPayloads = append(contractPayloads, signedPayload.Payload)
	}

	encoded := [][]byte{}
	for i, payload := range contractPayloads {
		txOpts, err := signature.Helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}

		signatureBytes, err := hex.DecodeString(signedPayloads[i].Signature[2:])
		if err != nil {
			return nil, err
		}

		tx, err := signature.abi.MintWithSignature(txOpts, payload, signatureBytes)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := signature.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := signature.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return signature.Helper.AwaitTx(ctx, tx.Hash())
}

// Verify that a signed payload is valid
//
// signedPayload: the payload to verify
//
// returns: true if the payload is valid, otherwise false.
//
// Example
//
//	// Learn more about how to craft a payload in the Generate() function
//	signedPayload, err := contract.Signature.Generate(payload)
//	isValid, err := contract.Signature.Verify(signedPayload)
func (signature *ERC721SignatureMinting) Verify(ctx context.Context, signedPayload SignedPayload721) (bool, error) {
	mintSignatureBytes, err := hex.DecodeString(signedPayload.Signature[2:])
	if err != nil {
		return false, err
	}

	if err != nil {
		return false, err
	}

	verification, _, err := signature.abi.Verify(&bind.CallOpts{Context: ctx}, signedPayload.Payload, mintSignatureBytes)
	return verification, err
}

// Generate a new payload from the given data
//
// payloadToSign: the payload containing the data for the signature mint
//
// returns: the payload signed by the minter's private key
//
// Example
//
//	payload := &thirdweb.Signature721PayloadInput{
//		To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6", // address to mint to
//		Price:                0,                                            // cost of minting
//		CurrencyAddress:      "0x0000000000000000000000000000000000000000", // currency to pay in order to mint
//		MintStartTime:        0,                                            // time where minting is allowed to start (epoch seconds)
//		MintEndTime:          100000000000000,                              // time when this signature expires (epoch seconds)
//		PrimarySaleRecipient: "0x0000000000000000000000000000000000000000", // address to receive the primary sales of this mint
//		Metadata: &thirdweb.NFTMetadataInput{																// metadata of the NFT to mint
//	 		Name:  "ERC721 Sigmint!",
//		},
//		RoyaltyRecipient: "0x0000000000000000000000000000000000000000",     // address to receive royalties of this mint
//		RoyaltyBps:       0,                                                // royalty cut of this mint in basis points
//	}
//
//	signedPayload, err := contract.Signature.Generate(payload)
func (signature *ERC721SignatureMinting) Generate(ctx context.Context, payloadToSign Signature721PayloadInput) (*SignedPayload721, error) {
	payload, err := signature.GenerateBatch(ctx, []Signature721PayloadInput{payloadToSign})
	if err != nil {
		return nil, err
	}

	return &payload[0], nil
}

// Generate a batch of new payload from the given data
//
// payloadToSign: the payloads containing the data for the signature mint
//
// returns: the payloads signed by the minter's private key
//
// Example
//
//	payload := []*thirdweb.Signature721PayloadInput{
//		&thirdweb.Signature721PayloadInput{
//			To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
//			Price:                0,
//			CurrencyAddress:      "0x0000000000000000000000000000000000000000",
//			MintStartTime:        0,
//			MintEndTime:          100000000000000,
//			PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
//			Metadata: &thirdweb.NFTMetadataInput{
//	 			Name:  "ERC721 Sigmint!",
//	 			Image: imageFile,
//			},
//			RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
//			RoyaltyBps:       0,
//		},
//		&thirdweb.Signature721PayloadInput{
//			To:                   "0x9e1b8A86fFEE4a7175DAE4bDB1cC12d111Dcb3D6",
//			Price:                0,
//			CurrencyAddress:      "0x0000000000000000000000000000000000000000",
//			MintStartTime:        0,
//			MintEndTime:          100000000000000,
//			PrimarySaleRecipient: "0x0000000000000000000000000000000000000000",
//			Metadata: &thirdweb.NFTMetadataInput{
//	 			Name:  "ERC721 Sigmint!",
//	 			Image: imageFile,
//			},
//			RoyaltyRecipient: "0x0000000000000000000000000000000000000000",
//			RoyaltyBps:       0,
//		},
//	}
//
//	signedPayload, err := contract.Signature.GenerateBatch(context.Background(), payload)
func (signature *ERC721SignatureMinting) GenerateBatch(ctx context.Context, payloadsToSign []Signature721PayloadInput) ([]SignedPayload721, error) {
	// TODO: Verify roles and return error

	metadatas := []*NFTMetadataInput{}
	for _, payload := range payloadsToSign {
		metadatas = append(metadatas, payload.Metadata)
	}

	uris, err := uploadOrExtractUris(ctx, metadatas, signature.storage)
	if err != nil {
		return nil, err
	}

	chainId, err := signature.Helper.GetChainID(ctx)
	if err != nil {
		return nil, err
	}

	var signedPayloads []SignedPayload721
	verifyingContract := signature.Helper.getAddress()
	signer := func(in []byte) ([]byte, error) {
		return crypto.Sign(in, signature.Helper.GetPrivateKey())
	}

	for i, uri := range uris {
		p := payloadsToSign[i]
		provider := signature.Helper.GetProvider()
		price, err := normalizePriceValue(ctx, provider, p.Price, p.CurrencyAddress)
		if err != nil {
			return nil, err
		}

		payload := Signature721Payload{
			To:                   common.HexToAddress(p.To),
			Price:                price,
			CurrencyAddress:      common.HexToAddress(p.CurrencyAddress),
			MintStartTime:        time.Unix(int64(p.MintStartTime), 0),
			MintEndTime:          time.Unix(int64(p.MintEndTime), 0),
			PrimarySaleRecipient: common.HexToAddress(p.PrimarySaleRecipient),
			RoyaltyRecipient:     common.HexToAddress(p.RoyaltyRecipient),
			RoyaltyBps:           p.RoyaltyBps,
			MetadataUri:          uri,
		}

		signed, err := GenerateAndSign(chainId, verifyingContract, payload, signer)
		if err != nil {
			return nil, err
		}
		signedPayloads = append(signedPayloads, *signed)
	}

	return signedPayloads, nil
}

func (signature *ERC721SignatureMinting) GenerateBatchWithUris(ctx context.Context, payloadsToSign []*Signature721PayloadInputWithUri) ([]SignedPayload721, error) {
	// TODO: Verify roles and return error

	chainId, err := signature.Helper.GetChainID(ctx)
	if err != nil {
		return nil, err
	}

	var signedPayloads []SignedPayload721

	verifyingContract := signature.Helper.getAddress()
	signer := func(in []byte) ([]byte, error) {
		return crypto.Sign(in, signature.Helper.GetPrivateKey())
	}
	for _, p := range payloadsToSign {
		provider := signature.Helper.GetProvider()
		price, err := normalizePriceValue(ctx, provider, p.Price, p.CurrencyAddress)
		if err != nil {
			return nil, err
		}
		payload := Signature721Payload{
			To:                   common.HexToAddress(p.To),
			Price:                price,
			CurrencyAddress:      common.HexToAddress(p.CurrencyAddress),
			MintStartTime:        time.Unix(int64(p.MintStartTime), 0),
			MintEndTime:          time.Unix(int64(p.MintEndTime), 0),
			PrimarySaleRecipient: common.HexToAddress(p.PrimarySaleRecipient),
			RoyaltyRecipient:     common.HexToAddress(p.RoyaltyRecipient),
			RoyaltyBps:           p.RoyaltyBps,
			MetadataUri:          p.MetadataUri,
		}

		signed, err := GenerateAndSign(chainId, verifyingContract, payload, signer)
		if err != nil {
			return nil, err
		}

		signedPayloads = append(signedPayloads, *signed)
	}

	return signedPayloads, nil
}

type SignFunc func([]byte) ([]byte, error)

// copied from GenerateBatch (https://github.com/fractalwagmi/fractal-api-server/blob/32e0fc8a99042d8603065929aa4d0c1d65fe8be6/vendor/github.com/thirdweb-dev/go-sdk/v2/thirdweb/erc721_signature_minting.go#L375)
func GenerateAndSign(chainId *big.Int, contract common.Address, p Signature721Payload, signFunc SignFunc) (*SignedPayload721, error) {
	// TODO: Verify roles and return error

	generatedId := uuid.New()
	id := [32]byte{}
	for i := 0; i < 16; i++ {
		id[16+i] = generatedId[i]
	}

	mintRequest := mapPayloadToContractStruct(id, p)
	mappedPayload := generateMessage(id, p)

	typedData := signerTypes.TypedData{
		Types: signerTypes.Types{
			"MintRequest": []signerTypes.Type{
				{Name: "to", Type: "address"},
				{Name: "royaltyRecipient", Type: "address"},
				{Name: "royaltyBps", Type: "uint256"},
				{Name: "primarySaleRecipient", Type: "address"},
				{Name: "uri", Type: "string"},
				{Name: "price", Type: "uint256"},
				{Name: "currency", Type: "address"},
				{Name: "validityStartTimestamp", Type: "uint128"},
				{Name: "validityEndTimestamp", Type: "uint128"},
				{Name: "uid", Type: "bytes32"},
			},
			"EIP712Domain": []signerTypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "MintRequest",
		Domain: signerTypes.TypedDataDomain{
			Name:              "TokenERC721",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
			VerifyingContract: contract.String(),
		},
		Message: mappedPayload,
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return nil, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	sigHash := crypto.Keccak256(rawData)

	signatureHash, err := signFunc(sigHash)
	if err != nil {
		return nil, err
	}

	// We need this to correct v = 0,1 to v = 27,28 - or else all will break
	if signatureHash[64] == 0 || signatureHash[64] == 1 {
		signatureHash[64] += 27
	}

	return &SignedPayload721{
		Payload:   mintRequest,
		Signature: "0x" + hex.EncodeToString(signatureHash),
	}, nil
}

func generateMessage(uid [32]byte, mintRequest Signature721Payload) signerTypes.TypedDataMessage {
	message := signerTypes.TypedDataMessage{
		"to":                     mintRequest.To.String(),
		"royaltyRecipient":       mintRequest.RoyaltyRecipient.String(),
		"royaltyBps":             fmt.Sprintf("%v", mintRequest.RoyaltyBps),
		"primarySaleRecipient":   mintRequest.PrimarySaleRecipient.String(),
		"uri":                    mintRequest.MetadataUri,
		"price":                  mintRequest.Price.String(),
		"currency":               mintRequest.CurrencyAddress.String(),
		"validityStartTimestamp": fmt.Sprintf("%d", mintRequest.MintStartTime.Unix()),
		"validityEndTimestamp":   fmt.Sprintf("%d", mintRequest.MintEndTime.Unix()),
		"uid":                    uid[:],
	}
	return message
}

func mapPayloadToContractStruct(uid [32]byte, mintRequest Signature721Payload) abi.ITokenERC721MintRequest {
	return abi.ITokenERC721MintRequest{
		To:                     mintRequest.To,
		RoyaltyRecipient:       mintRequest.RoyaltyRecipient,
		RoyaltyBps:             big.NewInt(int64(mintRequest.RoyaltyBps)),
		PrimarySaleRecipient:   mintRequest.PrimarySaleRecipient,
		Uri:                    mintRequest.MetadataUri,
		Price:                  mintRequest.Price,
		Currency:               mintRequest.CurrencyAddress,
		ValidityStartTimestamp: big.NewInt(mintRequest.MintStartTime.Unix()),
		ValidityEndTimestamp:   big.NewInt(mintRequest.MintEndTime.Unix()),
		Uid:                    uid,
	}
}
