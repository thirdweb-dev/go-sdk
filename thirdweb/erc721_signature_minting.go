package thirdweb

import (
	"bytes"
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
	"github.com/thirdweb-dev/go-sdk/internal/abi"
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

func (signature *ERC721SignatureMinting) Mint(signedPayload *SignedPayload721) (*types.Transaction, error) {
	message, err := signature.mapPayloadToContractStruct(signedPayload.Payload)
	if err != nil {
		return nil, err
	}

	txOpts := signature.helper.getTxOptions()
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

func (signature *ERC721SignatureMinting) MintBatch(signedPayloads []*SignedPayload721) (*types.Transaction, error) {
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
		tx, err := signature.abi.MintWithSignature(signature.helper.getTxOptions(), *payload, signedPayloads[i].Signature)
		if err != nil {
			return nil, err
		}

		encoded = append(encoded, tx.Data())
	}

	tx, err := signature.abi.Multicall(signature.helper.getTxOptions(), encoded)
	if err != nil {
		return nil, err
	}

	return signature.helper.awaitTx(tx.Hash())
}

func (signature *ERC721SignatureMinting) Verify(signedPayload *SignedPayload721) (bool, error) {
	mintRequest := signedPayload.Payload
	mintSignature := signedPayload.Signature
	message, err := signature.mapPayloadToContractStruct(mintRequest)

	fmt.Println("Sending payload...")
	fmt.Println(message)

	if err != nil {
		return false, err
	}

	verification, address, err := signature.abi.Verify(&bind.CallOpts{}, *message, mintSignature)
	fmt.Println("Address: ", address.String())
	fmt.Println("Signer: ", signature.helper.GetSignerAddress().String())

	return verification, err
}

func (signature *ERC721SignatureMinting) Generate(payloadToSign *Signature721PayloadInput) (*SignedPayload721, error) {
	payload, err := signature.GenerateBatch([]*Signature721PayloadInput{payloadToSign})
	if err != nil {
		return nil, err
	}

	return payload[0], nil
}

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
					{Name: "verifyingContract", Type: "string"},
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

		// // ============ Sketchy Sht ============

		// domainSeparator, err := typedData.EncodeData("EIP712Domain", typedData.Domain.Map(), 1)
		// if err != nil {
		// 	return nil, err
		// }

		// typedDataHash, err := typedData.EncodeData(typedData.PrimaryType, typedData.Message, 1)
		// if err != nil {
		// 	return nil, err
		// }

		// // ========== End Sketchy Sht ==========

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
		"uid":                    make([]byte, 32),
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
		Uid:                    [32]byte{},
	}, nil
}

func (signature *ERC721SignatureMinting) internalVerifySignature(signatureHash []byte, sigHash []byte) (bool, error) {
	// DANGER: Do not use this function! It mutates the signature objects to be wrong!
	if len(signatureHash) != 65 {
		return false, fmt.Errorf("invalid signature length: %d", len(signatureHash))
	}

	if signatureHash[64] != 27 && signatureHash[64] != 28 {
		return false, fmt.Errorf("invalid recovery id: %d", signatureHash[64])
	}
	signatureHash[64] -= 27

	pubKeyRaw, err := crypto.Ecrecover(sigHash, signatureHash)
	if err != nil {
		return false, fmt.Errorf("invalid signature: %s", err.Error())
	}

	pubKey, err := crypto.UnmarshalPubkey(pubKeyRaw)
	if err != nil {
		return false, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	if !bytes.Equal(recoveredAddr.Bytes(), signature.helper.GetSignerAddress().Bytes()) {
		return false, errors.New("Addresses do not match!")
	}

	fmt.Println(recoveredAddr.String())
	fmt.Println(signature.helper.GetSignerAddress().String())

	return true, nil
}
