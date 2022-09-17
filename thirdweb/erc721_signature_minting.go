package thirdweb

import (
	"context"
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
	"github.com/thirdweb-dev/go-sdk/abi"
)

// You can access this interface from the NFT Collection contract under the
// signature interface.
type ERC721SignatureMinting struct {
	abi     *abi.TokenERC721
	helper  *contractHelper
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
func (signature *ERC721SignatureMinting) Mint(ctx context.Context, signedPayload *SignedPayload721) (*types.Transaction, error) {
	message, err := signature.mapPayloadToContractStruct(signedPayload.Payload)
	if err != nil {
		return nil, err
	}

	txOpts, err := signature.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	setErc20Allowance(
		signature.helper,
		message.Price,
		message.Currency.String(),
		txOpts,
	)

	tx, err := signature.abi.MintWithSignature(txOpts, *message, signedPayload.Signature)
	if err != nil {
		return nil, err
	}

	return signature.helper.awaitTx(tx.Hash())
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
	contractPayloads := []*abi.ITokenERC721MintRequest{}
	for _, signedPayload := range signedPayloads {
		if signedPayload.Payload.Price > 0 {
			return nil, fmt.Errorf("Can only batch free mints. For mints with a price, use the Mint() function.")
		}

		payload, err := signature.mapPayloadToContractStruct(signedPayload.Payload)
		if err != nil {
			return nil, err
		}

		contractPayloads = append(contractPayloads, payload)
	}

	encoded := [][]byte{}
	for i, payload := range contractPayloads {
		txOpts, err := signature.helper.getEncodedTxOptions(ctx)
		if err != nil {
			return nil, err
		}
		tx, err := signature.abi.MintWithSignature(txOpts, *payload, signedPayloads[i].Signature)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	txOpts, err := signature.helper.getTxOptions(ctx)
	if err != nil {
		return nil, err
	}
	tx, err := signature.abi.Multicall(txOpts, encoded)
	if err != nil {
		return nil, err
	}

	return signature.helper.awaitTx(tx.Hash())
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
func (signature *ERC721SignatureMinting) Verify(signedPayload *SignedPayload721) (bool, error) {
	mintRequest := signedPayload.Payload
	mintSignature := signedPayload.Signature
	message, err := signature.mapPayloadToContractStruct(mintRequest)

	if err != nil {
		return false, err
	}

	verification, _, err := signature.abi.Verify(&bind.CallOpts{}, *message, mintSignature)
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
func (signature *ERC721SignatureMinting) Generate(payloadToSign *Signature721PayloadInput) (*SignedPayload721, error) {
	payload, err := signature.GenerateBatch([]*Signature721PayloadInput{payloadToSign})
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
//	signedPayload, err := contract.Signature.GenerateBatch(payload)
func (signature *ERC721SignatureMinting) GenerateBatch(payloadsToSign []*Signature721PayloadInput) ([]*SignedPayload721, error) {
	// TODO: Verify roles and return error

	metadatas := []*NFTMetadataInput{}
	for _, payload := range payloadsToSign {
		metadatas = append(metadatas, payload.Metadata)
	}

	uris, err := uploadOrExtractUris(metadatas, signature.storage)
	if err != nil {
		return nil, err
	}

	chainId, err := signature.helper.GetChainID()
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

		payload := &Signature721PayloadOutput{
			To:                   p.To,
			Price:                p.Price,
			CurrencyAddress:      p.CurrencyAddress,
			MintStartTime:        p.MintStartTime,
			MintEndTime:          p.MintEndTime,
			PrimarySaleRecipient: p.PrimarySaleRecipient,
			Metadata:             p.Metadata,
			RoyaltyRecipient:     p.RoyaltyRecipient,
			RoyaltyBps:           p.RoyaltyBps,
			Uri:                  uri,
			Uid:                  id,
		}

		mappedPayload, err := signature.generateMessage(payload)
		if err != nil {
			return nil, err
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
				VerifyingContract: signature.helper.getAddress().String(),
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

		privateKey := signature.helper.GetPrivateKey()
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
			Signature: signatureHash,
		})
	}

	return signedPayloads, nil
}

func (signature *ERC721SignatureMinting) generateMessage(mintRequest *Signature721PayloadOutput) (signerTypes.TypedDataMessage, error) {
	provider := signature.helper.GetProvider()
	price, err := normalizePriceValue(provider, mintRequest.Price, mintRequest.CurrencyAddress)
	if err != nil {
		return nil, err
	}

	message := signerTypes.TypedDataMessage{
		"to":                     mintRequest.To,
		"royaltyRecipient":       mintRequest.RoyaltyRecipient,
		"royaltyBps":             fmt.Sprintf("%v", mintRequest.RoyaltyBps),
		"primarySaleRecipient":   mintRequest.PrimarySaleRecipient,
		"uri":                    mintRequest.Uri,
		"price":                  fmt.Sprintf("%v", int(price.Int64())),
		"currency":               mintRequest.CurrencyAddress,
		"validityStartTimestamp": fmt.Sprintf("%v", mintRequest.MintStartTime),
		"validityEndTimestamp":   fmt.Sprintf("%v", mintRequest.MintEndTime),
		"uid":                    mintRequest.Uid[:],
	}

	return message, nil
}

func (signature *ERC721SignatureMinting) mapPayloadToContractStruct(mintRequest *Signature721PayloadOutput) (*abi.ITokenERC721MintRequest, error) {
	provider := signature.helper.GetProvider()
	price, err := normalizePriceValue(provider, mintRequest.Price, mintRequest.CurrencyAddress)
	if err != nil {
		return nil, err
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
