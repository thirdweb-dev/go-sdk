package nftlabs

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func getPublicAddress(key *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := key.Public()

	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		// TODO: return better error
		return [20]byte{}, errors.New("Failed to decode public key from private key, maybe wrong format")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
