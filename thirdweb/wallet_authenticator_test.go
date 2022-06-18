package thirdweb

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const domain = "thirdweb.com"

func TestVerify(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	address, _ := sdk.Auth.Verify(domain, payload, nil)

	assert.Equal(t, address, adminWallet)
}

func TestVerifyWithParams(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, &WalletLoginOptions{
		ExpirationTime: time.Now().Add(time.Hour * 5),
		ChainId:        137,
	})

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	address, _ := sdk.Auth.Verify(domain, payload, &WalletVerifyOptions{
		ChainId: 137,
	})

	assert.Equal(t, address, adminWallet)
}

func TestRejectIncorrectDomain(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	_, err := sdk.Auth.Verify("test.thirdweb.com", payload, nil)

	assert.Equal(t, err.Error(), "Expected domain 'test.thirdweb.com' does not match domain on payload 'thirdweb.com'")
}

func TestRejectExpiredPayload(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, &WalletLoginOptions{
		ExpirationTime: time.Now().Add(-time.Hour * 5),
	})

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	_, err := sdk.Auth.Verify(domain, payload, nil)

	assert.Equal(t, err.Error(), "Login request has expired")
}

func TestRejectIncorrectChainId(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, &WalletLoginOptions{
		ChainId: 1,
	})

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	_, err := sdk.Auth.Verify(domain, payload, &WalletVerifyOptions{
		ChainId: 137,
	})

	assert.Equal(t, err.Error(), "Chain ID '137' does not match payload chain ID '1'")
}

func TestRejectIncorrectSigner(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)
	payload.Payload.Address = secondaryWallet

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	_, err := sdk.Auth.Verify(domain, payload, nil)

	assert.Equal(
		t,
		err.Error(),
		fmt.Sprintf(
			"The intended payload address '%s' is not the payload signer",
			strings.ToLower(secondaryWallet),
		),
	)
}

func TestGenerateToken(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	token, _ := sdk.Auth.GenerateAuthToken(domain, payload, nil)
	address, _ := sdk.Auth.Authenticate(domain, token)

	assert.Equal(t, address, adminWallet)
}

func TestRejectTokenIncorrectDoain(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	token, _ := sdk.Auth.GenerateAuthToken(domain, payload, nil)
	_, err := sdk.Auth.Authenticate("test.thirdweb.com", token)

	assert.Equal(t, err.Error(), "Expected token to be for the domain 'test.thirdweb.com', but found token with domain 'thirdweb.com'")
}

func TestRejectTokenBeforeInvalid(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	token, _ := sdk.Auth.GenerateAuthToken(domain, payload, &WalletAuthenticationOptions{
		InvalidBefore: time.Now().Add(time.Hour * 5),
	})
	_, err := sdk.Auth.Authenticate(domain, token)

	assert.Contains(t, err.Error(), "This token is invalid before")
}

func TestRejectExpiredToken(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	token, _ := sdk.Auth.GenerateAuthToken(domain, payload, &WalletAuthenticationOptions{
		ExpirationTime: time.Now().Add(-time.Hour * 5),
	})
	_, err := sdk.Auth.Authenticate(domain, token)

	assert.Contains(t, err.Error(), "This token expired")
}

func TestRejectAdminNotConnected(t *testing.T) {
	sdk := getSDK()
	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	payload, _ := sdk.Auth.Login(domain, nil)

	sdk.Auth.UpdatePrivateKey(secondaryPrivateKey)
	token, _ := sdk.Auth.GenerateAuthToken(domain, payload, nil)

	sdk.Auth.UpdatePrivateKey(adminPrivateKey)
	_, err := sdk.Auth.Authenticate(domain, token)

	assert.Contains(t, err.Error(), "Expected the connected wallet address")
}
