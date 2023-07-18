package thirdweb

import (
	"crypto/sha256"
	"encoding/hex"
)

func deriveClientId(secretKey string) string {
    hasher := sha256.New()
    hasher.Write([]byte(secretKey))
    hashed := hex.EncodeToString(hasher.Sum(nil))
    return hashed[:32]
}