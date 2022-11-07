package thirdweb

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

// > This API is provided as a preview for developers and may change based on feedback that we receive. Do not use this API in a production environment.
//
// The wallet authenticator enables server-side applications to securely identify the
// connected wallet address of users on the client-side, and also enables users to authenticate
// to any backend using just their wallet. It implements the JSON Web Token (JWT) authentication
// standard.
//
// You can use the wallet authenticator as follows:
//
//	// First we specify the domain of the application to authenticate to
//	domain := "example.com"
//
//	// We can then generate a payload for the connected wallet to login
//	// This can also be done on the client side with the thirdweb TypeScript SDK
//	payload, err := sdk.Auth.Login(domain, nil)
//
//	// Then, on the server, we can securely verify the connected address that signed the payload
//	address, err := sdk.Auth.Verify(domain, payload, nil)
//
//	// And we can also generate an authentication token to send back to the original payload sender
//	token, err := sdk.Auth.GenerateAuthToken(domain, payload, nil)
//
//	// Finally, the token can be use dby the original payload sender to authenticate to the backend
//	// And the server can use the following function to authenticate the token and verify the address
//	address, err := sdk.Auth.Authenticate(domain, token)
type WalletAuthenticator struct {
	*ProviderHandler
}

func newWalletAuthenticator(provider *ethclient.Client, privateKey string) (*WalletAuthenticator, error) {
	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	return &WalletAuthenticator{handler}, nil
}

// Client-side function that allows the connected wallet to login to a server-side application.
// Generates a login payload that can be sent to the server-side for verification or authentication.
//
// domain: The domain of the application that you want to log in to
//
// options: Optional configuration options for the login payload
//
// returns: A login payload that can be sent to the server-side for verification or authentication
//
// Example
//
//	// Add the domain of the application that you want to log in to
//	domain := "example.com"
//
//	// Generate a signed login payload for the connected wallet to authenticate with
//	payload, err := sdk.Auth.Login(domain, nil)
func (auth *WalletAuthenticator) Login(
	domain string,
	options *WalletLoginOptions,
) (*WalletLoginPayload, error) {
	err := auth.requireSigner()
	if err != nil {
		return nil, err
	}

	signerAddress := auth.GetSignerAddress().String()
	payloadData := &WalletLoginPayloadData{
		Domain:         domain,
		ExpirationTime: time.Now().Add(time.Minute * 5),
		Address:        signerAddress,
		Nonce:          uuid.New().String(),
		ChainId:        0,
	}

	if options != nil {
		if !options.ExpirationTime.IsZero() {
			payloadData.ExpirationTime = options.ExpirationTime
		}

		if options.Nonce != "" {
			payloadData.Nonce = options.Nonce
		}

		if options.ChainId != 0 {
			payloadData.ChainId = options.ChainId
		}
	}

	message := auth.generateMessage(payloadData)
	signature, err := auth.signMessage(message)

	return &WalletLoginPayload{
		Payload:   payloadData,
		Signature: fmt.Sprintf("0x%s", hex.EncodeToString(signature)),
	}, nil
}

// Server-side function to securely verify the address of the logged in client-side wallet
// by validating the provided client-side login request.
//
// domain: The domain of the application to verify the login request for
//
// payload: The login payload to verify
//
// returns: The address of the logged in wallet that signed the payload
//
// Example
//
//	domain := "example.com"
//	payload, err := sdk.Auth.Login(domain, nil)
//
//	// Verify the login request
//	address, err := sdk.Auth.Verify(domain, payload, nil)
func (auth *WalletAuthenticator) Verify(
	domain string,
	payload *WalletLoginPayload,
	options *WalletVerifyOptions,
) (string, error) {
	// Check that the intended domain matches the domain of the payload
	if payload.Payload.Domain != domain {
		return "", fmt.Errorf(
			"Expected domain '%s' does not match domain on payload '%s'",
			domain,
			payload.Payload.Domain,
		)
	}

	// Check that the payload hasn't expired
	currentTime := time.Now()
	if payload.Payload.ExpirationTime.Before(currentTime) {
		return "", errors.New("Login request has expired")
	}

	// If chain ID is specified, check that it matches the chain ID of the signature
	if options != nil && options.ChainId != 0 && options.ChainId != payload.Payload.ChainId {
		return "", fmt.Errorf(
			"Chain ID '%d' does not match payload chain ID '%d'",
			options.ChainId,
			payload.Payload.ChainId,
		)
	}

	decodedSignature, err := hexutil.Decode(payload.Signature)
	if err != nil {
		return "", err
	}

	message := auth.generateMessage(payload.Payload)
	userAddress, err := auth.recoverAddress(message, decodedSignature)
	if err != nil {
		return "", err
	}

	if strings.ToLower(userAddress) != strings.ToLower(payload.Payload.Address) {
		return "", fmt.Errorf(
			"The intended payload address '%s' is not the payload signer",
			strings.ToLower(payload.Payload.Address),
		)
	}

	return userAddress, nil
}

// Server-side function that generates a JWT token from the provided login request that the
// client-side wallet can use to authenticate to the server-side application.
//
// domain: The domain of the application to authenticate to
//
// payload: The login payload to authenticate with
//
// options: Optional configuration options for the authentication token
//
// returns:  An authentication token that can be used to make authenticated requests to the server
//
// Example
//
//	domain := "example.com"
//	payload, err := sdk.Auth.Login(domain, nil)
//
//	// Generate an authentication token for the logged in wallet
//	token, err := sdk.Auth.GenerateAuthToken(domain, payload, nil)
func (auth *WalletAuthenticator) GenerateAuthToken(
	domain string,
	payload *WalletLoginPayload,
	options *WalletAuthenticationOptions,
) (string, error) {
	err := auth.requireSigner()
	if err != nil {
		return "", err
	}

	userAddress, err := auth.Verify(domain, payload, nil)
	if err != nil {
		return "", err
	}
	adminAddress := auth.GetSignerAddress().String()
	payloadData := &WalletAuthenticationPayloadData{
		Iss: adminAddress,
		Sub: userAddress,
		Aud: domain,
		Nbf: time.Now().Unix(),
		Exp: time.Now().Add(time.Hour * 5).Unix(),
		Iat: time.Now().Unix(),
		Jti: uuid.New().String(),
	}

	if options != nil {
		if !options.ExpirationTime.IsZero() {
			payloadData.Exp = options.ExpirationTime.Unix()
		}

		if !options.InvalidBefore.IsZero() {
			payloadData.Nbf = options.InvalidBefore.Unix()
		}
	}

	message, err := json.Marshal(payloadData)
	if err != nil {
		return "", err
	}

	signature, err := auth.signMessage(string(message))
	if err != nil {
		return "", err
	}

	// Header used for JWT token specifying hash algorithm
	header := map[string]string{
		// Specify ECDSA with SHA-256 for hashing algorithm
		"alg": "ES256",
		"typ": "JWT",
	}
	headerBytes, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	encodedHeader := auth.base64Encode(string(headerBytes))
	encodedData := auth.base64Encode(string(message))
	encodedSignature := auth.base64Encode(string(signature))

	// Generate a JWT token with base64 encoded header, payload, and signature
	token := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedData, encodedSignature)

	return token, nil
}

// Server-side function that authenticates the provided JWT token. This function verifies that
// the provided authentication token is valid and returns the address of the authenticated wallet.
//
// domain: The domain of the application to authenticate the token to
//
// token: The authentication token to authenticate with
//
// returns: The address of the authenticated wallet
//
// Example
//
//	domain := "example.com"
//	payload, err := sdk.Auth.Login(domain)
//	token, err := sdk.Auth.GenerateAuthToken(domain, payload)
//
//	// Authenticate the token and get the address of the authenticating wallet
//	address, err := sdk.Auth.Authenticate(domain, token)
func (auth *WalletAuthenticator) Authenticate(
	domain string,
	token string,
) (string, error) {
	err := auth.requireSigner()
	if err != nil {
		return "", err
	}

	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		return "", fmt.Errorf("Invalid authentication token format")
	}

	encodedPayload := tokenParts[1]
	encodedSignature := tokenParts[2]

	decodedPayload, err := auth.base64Decode(encodedPayload)
	if err != nil {
		return "", err
	}

	decodedSignature, err := auth.base64Decode(encodedSignature)
	if err != nil {
		return "", err
	}

	payload := &WalletAuthenticationPayloadData{}
	err = json.Unmarshal([]byte(decodedPayload), &payload)
	if err != nil {
		return "", err
	}

	// Check that the intended audience matches the domain
	if payload.Aud != domain {
		return "", fmt.Errorf(
			"Expected token to be for the domain '%s', but found token with domain '%s'",
			domain,
			payload.Aud,
		)
	}

	// Check tht the token is past the invalid before time
	if payload.Nbf > time.Now().Unix() {
		return "", fmt.Errorf(
			"This token is invalid before epoch time '%d', current epoch time is '%d'",
			payload.Nbf,
			time.Now().Unix(),
		)
	}

	// Check that the token hasn't expired
	if payload.Exp < time.Now().Unix() {
		return "", fmt.Errorf(
			"This token expired at epoch time '%d', current epoch time is '%d'",
			payload.Exp,
			time.Now().Unix(),
		)
	}

	// Check that the connected wallet matches the token issuer
	signerAddress := auth.GetSignerAddress().String()
	if strings.ToLower(signerAddress) != strings.ToLower(payload.Iss) {
		return "", fmt.Errorf(
			"Expected the connected wallet address '%s' to match the token issuer address '%s'",
			strings.ToLower(signerAddress),
			strings.ToLower(payload.Iss),
		)
	}

	// Check that the connected wallet signed the token
	message, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	adminAddress, err := auth.recoverAddress(string(message), []byte(decodedSignature))
	if err != nil {
		return "", err
	}

	if strings.ToLower(adminAddress) != strings.ToLower(signerAddress) {
		return "", fmt.Errorf(
			"The connected wallet address '%s' did not sign the token",
			strings.ToLower(signerAddress),
		)
	}

	return payload.Sub, nil
}

func (auth *WalletAuthenticator) generateMessage(
	payload *WalletLoginPayloadData,
) string {
	message := ""

	// Add the domain and login address for transparency
	message += fmt.Sprintf("%s wants you to sign in with your account:\n%s\n\n", payload.Domain, payload.Address)

	// Prompt user to make sure that domain is correct to prevent phishing attacks
	message += "Make sure that the requesting domain above matches the URL of the current website.\n\n"

	// Add data fields in compliance with the EIP-4361 standard
	if payload.ChainId != 0 {
		message += fmt.Sprintf("Chain ID: %d\n", payload.ChainId)
	}

	message += fmt.Sprintf("Nonce: %s\n", payload.Nonce)
	message += fmt.Sprintf("Expiration Time: %s\n", payload.ExpirationTime.Format("2006-01-02T15:04:05.999Z"))

	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	return fullMessage
}

func (auth *WalletAuthenticator) recoverAddress(message string, signature []byte) (string, error) {
	// Support both formats of recovery bit (27/28 or 0/1)
	if signature[64] == 27 || signature[64] == 28 {
		signature[64] -= 27
	}

	messageByte := []byte(message)
	messageHash := crypto.Keccak256Hash(messageByte)
	publicKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		return "", err
	}

	address := crypto.PubkeyToAddress(*publicKey)
	return address.String(), nil
}

func (auth *WalletAuthenticator) requireSigner() error {
	if auth.GetPrivateKey() == nil {
		return errors.New("This action requires a connected wallet. Please pass a valid private key to the SDK")
	}

	return nil
}

func (auth *WalletAuthenticator) signMessage(message string) ([]byte, error) {
	err := auth.requireSigner()
	if err != nil {
		return nil, err
	}

	messageByte := []byte(message)
	messageHash := crypto.Keccak256Hash(messageByte)

	privateKey := auth.GetPrivateKey()
	signatureHash, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	return signatureHash, nil
}

func (auth *WalletAuthenticator) base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func (auth *WalletAuthenticator) base64Decode(data string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
