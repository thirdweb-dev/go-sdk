package thirdweb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const domain = "thirdweb.com"

func TestVerify(t *testing.T) {
	sdk := getSDK()
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	address, _ := sdk.Auth.Verify(domain, payload, nil)

	assert.Equal(t, address, adminWallet)
}
