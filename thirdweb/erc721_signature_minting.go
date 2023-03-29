package thirdweb

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

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
	legacy    *abi.TokenERC721
	extension *abi.SignatureMintERC721
	Helper    *contractHelper
	storage   storage
}

func newERC721SignatureMinting(provider *ethclient.Client, address common.Address, privateKey string, storage storage) (*ERC721SignatureMinting, error) {
	legacy, err := abi.NewTokenERC721(address, provider)
	if err != nil {
		return nil, err
	}

	extension, err := abi.NewSignatureMintERC721(address, provider)
	if err != nil {
		return nil, err
	}

	helper, err := newContractHelper(address, provider, privateKey)
	if err != nil {
		return nil, err
	}

	return &ERC721SignatureMinting{
		legacy,
		extension,
		helper,
		storage,
	}, nil
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

// Deprecated: use MintAndAwait
func (signature *ERC721SignatureMinting) Mint(ctx context.Context, signedPayload *SignedPayload721) (*types.Transaction, error) {
	return signature.MintAndAwait(ctx, signedPayload)
}

func (signature *ERC721SignatureMinting) MintAndAwait(ctx context.Context, signedPayload *SignedPayload721) (*types.Transaction, error) {
	txOpts, err := signature.Helper.GetTxOptions(ctx)
	if err != nil {
		return nil, err
	}

	tx, err := signature.MintWithOpts(ctx, signedPayload, txOpts)

	if err != nil {
		return nil, err
	}

	return signature.Helper.AwaitTx(ctx, tx.Hash())
}

func (signature *ERC721SignatureMinting) MintWithOpts(ctx context.Context, signedPayload *SignedPayload721, txOpts *bind.TransactOpts) (*types.Transaction, error) {
	if signature.isLegacyContract(ctx) {
		message, err := mapLegacyPayloadToContractStruct(signedPayload.Payload)
		if err != nil {
			return nil, err
		}

		if err := setErc20Allowance(
			ctx,
			signature.Helper,
			message.Price,
			message.Currency.String(),
			txOpts,
		); err != nil {
			return nil, err
		}

		signatureBytes, err := hex.DecodeString(signedPayload.Signature[2:])
		if err != nil {
			return nil, err
		}

		return signature.legacy.MintWithSignature(txOpts, *message, signatureBytes)
	} else {
		message, err := mapPayloadToContractStruct(signedPayload.Payload)
		if err != nil {
			return nil, err
		}

		if err := setErc20Allowance(
			ctx,
			signature.Helper,
			message.PricePerToken,
			message.Currency.String(),
			txOpts,
		); err != nil {
			return nil, err
		}

		signatureBytes, err := hex.DecodeString(signedPayload.Signature[2:])
		if err != nil {
			return nil, err
		}

		return signature.extension.MintWithSignature(txOpts, *message, signatureBytes)
	}
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
	if signature.isLegacyContract(ctx) {
		contractPayloads := []*abi.ITokenERC721MintRequest{}
		for _, signedPayload := range signedPayloads {
			price, ok := big.NewInt(0).SetString(signedPayload.Payload.Price, 10)
			if !ok {
				return nil, errors.New("Specified price was not a valid big.Int")
			}

			if price.Cmp(big.NewInt(0)) == 1 {
				return nil, fmt.Errorf("Can only batch free mints. For mints with a price, use the Mint() function.")
			}

			payload, err := mapLegacyPayloadToContractStruct(signedPayload.Payload)
			if err != nil {
				return nil, err
			}

			contractPayloads = append(contractPayloads, payload)
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

			tx, err := signature.legacy.MintWithSignature(txOpts, *payload, signatureBytes)
			if err != nil {
				return nil, err
			}

			encoded = append(encoded, tx.Data())
		}

		txOpts, err := signature.Helper.GetTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := signature.legacy.Multicall(txOpts, encoded)
		if err != nil {
			return nil, err
		}

		return signature.Helper.AwaitTx(ctx, tx.Hash())
	} else {
		contractPayloads := []*abi.ISignatureMintERC721MintRequest{}
		for _, signedPayload := range signedPayloads {
			price, ok := big.NewInt(0).SetString(signedPayload.Payload.Price, 10)
			if !ok {
				return nil, errors.New("Specified price was not a valid big.Int")
			}

			if price.Cmp(big.NewInt(0)) == 1 {
				return nil, fmt.Errorf("Can only batch free mints. For mints with a price, use the Mint() function.")
			}

			payload, err := mapPayloadToContractStruct(signedPayload.Payload)
			if err != nil {
				return nil, err
			}

			contractPayloads = append(contractPayloads, payload)
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

			tx, err := signature.extension.MintWithSignature(txOpts, *payload, signatureBytes)
			if err != nil {
				return nil, err
			}

			encoded = append(encoded, tx.Data())
		}

		txOpts, err := signature.Helper.GetTxOptions(ctx)
		if err != nil {
			return nil, err
		}

		// TODO: This should throw if you don't have multicall on contract
		tx, err := signature.legacy.Multicall(txOpts, encoded)
		if err != nil {
			return nil, err
		}

		return signature.Helper.AwaitTx(ctx, tx.Hash())
	}
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
func (signature *ERC721SignatureMinting) Verify(ctx context.Context, signedPayload *SignedPayload721) (bool, error) {
	mintRequest := signedPayload.Payload
	mintSignatureBytes, err := hex.DecodeString(signedPayload.Signature[2:])
	if err != nil {
		return false, err
	}

	if signature.isLegacyContract(ctx) {
		message, err := mapLegacyPayloadToContractStruct(mintRequest)

		if err != nil {
			return false, err
		}

		verification, _, err := signature.legacy.Verify(&bind.CallOpts{Context: ctx}, *message, mintSignatureBytes)
		return verification, err
	} else {
		message, err := mapPayloadToContractStruct(mintRequest)

		if err != nil {
			return false, err
		}

		verification, err := signature.extension.Verify(&bind.CallOpts{Context: ctx}, *message, mintSignatureBytes)
		return verification.Success, err
	}
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
func (signature *ERC721SignatureMinting) Generate(ctx context.Context, payloadToSign *Signature721PayloadInput) (*SignedPayload721, error) {
	payload, err := signature.GenerateBatch(ctx, []*Signature721PayloadInput{payloadToSign})
	if err != nil {
		return nil, err
	}

	return payload[0], nil
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
func (signature *ERC721SignatureMinting) GenerateBatch(ctx context.Context, payloadsToSign []*Signature721PayloadInput) ([]*SignedPayload721, error) {
	// TODO: Verify roles and return error

	metadatas := []*NFTMetadataInput{}
	for _, payload := range payloadsToSign {
		metadatas = append(metadatas, payload.Metadata)
	}

	uris, err := uploadOrExtractUris(ctx, metadatas, signature.storage)
	if err != nil {
		return nil, err
	}

	signedPayloads := []*SignedPayload721{}

	for i, uri := range uris {
		p := payloadsToSign[i]

		generatedId := uuid.New()
		id := [32]byte{}
		for i := 0; i < 16; i++ {
			id[16+i] = generatedId[i]
		}

		provider := signature.Helper.GetProvider()
		price, err := normalizePriceValue(ctx, provider, p.Price, p.CurrencyAddress)
		if err != nil {
			return nil, err
		}

		payload := &Signature721PayloadOutput{
			To:                   p.To,
			Price:                price.String(),
			CurrencyAddress:      p.CurrencyAddress,
			MintStartTime:        p.MintStartTime,
			MintEndTime:          p.MintEndTime,
			PrimarySaleRecipient: p.PrimarySaleRecipient,
			RoyaltyRecipient:     p.RoyaltyRecipient,
			RoyaltyBps:           p.RoyaltyBps,
			Uri:                  uri,
			Uid:                  id,
		}

		typedData, err := signature.generateMessage(ctx, payload)
		if err != nil {
			return nil, err
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

		privateKey := signature.Helper.GetPrivateKey()
		signatureHash, err := crypto.Sign(sigHash, privateKey)
		if err != nil {
			return nil, err
		}

		// We need this to correct v = 0,1 to v = 27,28 - or else all will break
		if signatureHash[64] == 0 || signatureHash[64] == 1 {
			signatureHash[64] += 27
		}

		signedPayloads = append(signedPayloads, &SignedPayload721{
			Payload:   payload,
			Signature: "0x" + hex.EncodeToString(signatureHash),
		})
	}

	return signedPayloads, nil
}

func (signature *ERC721SignatureMinting) GenerateBatchWithUris(ctx context.Context, payloadsToSign []*Signature721PayloadInputWithUri) ([]*SignedPayload721, error) {
	// TODO: Verify roles and return error
	signedPayloads := []*SignedPayload721{}

	for _, p := range payloadsToSign {

		generatedId := uuid.New()
		id := [32]byte{}
		for i := 0; i < 16; i++ {
			id[16+i] = generatedId[i]
		}

		provider := signature.Helper.GetProvider()
		price, err := normalizePriceValue(ctx, provider, p.Price, p.CurrencyAddress)
		if err != nil {
			return nil, err
		}

		payload := &Signature721PayloadOutput{
			To:                   p.To,
			Price:                price.String(),
			CurrencyAddress:      p.CurrencyAddress,
			MintStartTime:        p.MintStartTime,
			MintEndTime:          p.MintEndTime,
			PrimarySaleRecipient: p.PrimarySaleRecipient,
			RoyaltyRecipient:     p.RoyaltyRecipient,
			RoyaltyBps:           p.RoyaltyBps,
			Uri:                  p.MetadataUri,
			Uid:                  id,
		}

		typedData, err := signature.generateMessage(ctx, payload)
		if err != nil {
			return nil, err
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

		privateKey := signature.Helper.GetPrivateKey()
		signatureHash, err := crypto.Sign(sigHash, privateKey)
		if err != nil {
			return nil, err
		}

		// We need this to correct v = 0,1 to v = 27,28 - or else all will break
		if signatureHash[64] == 0 || signatureHash[64] == 1 {
			signatureHash[64] += 27
		}

		signedPayloads = append(signedPayloads, &SignedPayload721{
			Payload:   payload,
			Signature: "0x" + hex.EncodeToString(signatureHash),
		})
	}

	return signedPayloads, nil
}

func (signature *ERC721SignatureMinting) generateMessage(ctx context.Context, mintRequest *Signature721PayloadOutput) (*signerTypes.TypedData, error) {
	chainId, err := signature.Helper.GetChainID(ctx)
	if err != nil {
		return nil, err
	}

	if signature.isLegacyContract(ctx) {
		message := signerTypes.TypedDataMessage{
			"to":                     mintRequest.To,
			"royaltyRecipient":       mintRequest.RoyaltyRecipient,
			"royaltyBps":             fmt.Sprintf("%v", mintRequest.RoyaltyBps),
			"primarySaleRecipient":   mintRequest.PrimarySaleRecipient,
			"uri":                    mintRequest.Uri,
			"price":                  fmt.Sprintf("%v", mintRequest.Price),
			"currency":               mintRequest.CurrencyAddress,
			"validityStartTimestamp": fmt.Sprintf("%v", mintRequest.MintStartTime),
			"validityEndTimestamp":   fmt.Sprintf("%v", mintRequest.MintEndTime),
			"uid":                    mintRequest.Uid[:],
		}

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
				VerifyingContract: signature.Helper.getAddress().String(),
			},
			Message: message,
		}

		return &typedData, nil
	} else {
		message := signerTypes.TypedDataMessage{
			"to":                     mintRequest.To,
			"royaltyRecipient":       mintRequest.RoyaltyRecipient,
			"royaltyBps":             fmt.Sprintf("%v", mintRequest.RoyaltyBps),
			"quantity":               fmt.Sprintf("%v", 1),
			"primarySaleRecipient":   mintRequest.PrimarySaleRecipient,
			"uri":                    mintRequest.Uri,
			"pricePerToken":          fmt.Sprintf("%v", mintRequest.Price),
			"currency":               mintRequest.CurrencyAddress,
			"validityStartTimestamp": fmt.Sprintf("%v", mintRequest.MintStartTime),
			"validityEndTimestamp":   fmt.Sprintf("%v", mintRequest.MintEndTime),
			"uid":                    mintRequest.Uid[:],
		}

		typedData := signerTypes.TypedData{
			Types: signerTypes.Types{
				"MintRequest": []signerTypes.Type{
					{Name: "to", Type: "address"},
					{Name: "royaltyRecipient", Type: "address"},
					{Name: "royaltyBps", Type: "uint256"},
					{Name: "primarySaleRecipient", Type: "address"},
					{Name: "uri", Type: "string"},
					{Name: "quantity", Type: "uint256"},
					{Name: "pricePerToken", Type: "uint256"},
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
				Name:              "SignatureMintERC721",
				Version:           "1",
				ChainId:           math.NewHexOrDecimal256(chainId.Int64()),
				VerifyingContract: signature.Helper.getAddress().String(),
			},
			Message: message,
		}

		return &typedData, nil
	}
}

func (signature *ERC721SignatureMinting) isLegacyContract(ctx context.Context) bool {
	contractType, err := signature.legacy.ContractType(&bind.CallOpts{
		Context: ctx,
	})
	if err != nil {
		return false
	}

	// If contractType = TokenERC721
	return hex.EncodeToString(contractType[:]) == "546f6b656e455243373231000000000000000000000000000000000000000000"
}

func mapLegacyPayloadToContractStruct(mintRequest *Signature721PayloadOutput) (*abi.ITokenERC721MintRequest, error) {
	price, ok := big.NewInt(0).SetString(mintRequest.Price, 10)
	if !ok {
		return nil, errors.New("Specified price was not a valid big.Int")
	}

	return &abi.ITokenERC721MintRequest{
		To:                     common.HexToAddress(mintRequest.To),
		RoyaltyRecipient:       common.HexToAddress(mintRequest.RoyaltyRecipient),
		RoyaltyBps:             big.NewInt(int64(mintRequest.RoyaltyBps)),
		PrimarySaleRecipient:   common.HexToAddress(mintRequest.PrimarySaleRecipient),
		Uri:                    mintRequest.Uri,
		Price:                  price,
		Currency:               common.HexToAddress(mintRequest.CurrencyAddress),
		ValidityStartTimestamp: big.NewInt(int64(mintRequest.MintStartTime)),
		ValidityEndTimestamp:   big.NewInt(int64(mintRequest.MintEndTime)),
		Uid:                    mintRequest.Uid,
	}, nil
}

func mapPayloadToContractStruct(mintRequest *Signature721PayloadOutput) (*abi.ISignatureMintERC721MintRequest, error) {
	price, ok := big.NewInt(0).SetString(mintRequest.Price, 10)
	if !ok {
		return nil, errors.New("Specified price was not a valid big.Int")
	}

	return &abi.ISignatureMintERC721MintRequest{
		To:                     common.HexToAddress(mintRequest.To),
		RoyaltyRecipient:       common.HexToAddress(mintRequest.RoyaltyRecipient),
		RoyaltyBps:             big.NewInt(int64(mintRequest.RoyaltyBps)),
		PrimarySaleRecipient:   common.HexToAddress(mintRequest.PrimarySaleRecipient),
		Uri:                    mintRequest.Uri,
		Quantity:               big.NewInt(1), // Always has quantity of 1
		PricePerToken:          price,
		Currency:               common.HexToAddress(mintRequest.CurrencyAddress),
		ValidityStartTimestamp: big.NewInt(int64(mintRequest.MintStartTime)),
		ValidityEndTimestamp:   big.NewInt(int64(mintRequest.MintEndTime)),
		Uid:                    mintRequest.Uid,
	}, nil
}
