
## Wallet Authenticator

\> This API is provided as a preview for developers and may change based on feedback that we receive\. Do not use this API in a production environment\.

The wallet authenticator enables server\-side applications to securely identify the connected wallet address of users on the client\-side, and also enables users to authenticate to any backend using just their wallet\. It implements the JSON Web Token \(JWT\) authentication standard\.

You can use the wallet authenticator as follows:

```
// First we specify the domain of the application to authenticate to
domain := "example.com"

// We can then generate a payload for the connected wallet to login
// This can also be done on the client side with the thirdweb TypeScript SDK
payload, err := sdk.Auth.Login(domain, nil)

// Then, on the server, we can securely verify the connected address that signed the payload
address, err := sdk.Auth.Verify(domain, payload, nil)

// And we can also generate an authentication token to send back to the original payload sender
token, err := sdk.Auth.GenerateAuthToken(domain, payload, nil)

// Finally, the token can be use dby the original payload sender to authenticate to the backend
// And the server can use the following function to authenticate the token and verify the address
address, err := sdk.Auth.Authenticate(domain, token)
```

```go
type WalletAuthenticator struct {
    *ProviderHandler
}
```

### func \(\*WalletAuthenticator\) [Authenticate](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/wallet_authenticator.go#L280-L283>)

```go
func (auth *WalletAuthenticator) Authenticate(domain string, token string) (string, error)
```

Server\-side function that authenticates the provided JWT token\. This function verifies that the provided authentication token is valid and returns the address of the authenticated wallet\.

domain: The domain of the application to authenticate the token to

token: The authentication token to authenticate with

returns: The address of the authenticated wallet

#### Example

```
domain := "example.com"
payload, err := sdk.Auth.Login(domain)
token, err := sdk.Auth.GenerateAuthToken(domain, payload)

// Authenticate the token and get the address of the authenticating wallet
address, err := sdk.Auth.Authenticate(domain, token)
```

### func \(\*WalletAuthenticator\) [GenerateAuthToken](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/wallet_authenticator.go#L197-L201>)

```go
func (auth *WalletAuthenticator) GenerateAuthToken(domain string, payload *WalletLoginPayload, options *WalletAuthenticationOptions) (string, error)
```

Server\-side function that generates a JWT token from the provided login request that the client\-side wallet can use to authenticate to the server\-side application\.

domain: The domain of the application to authenticate to

payload: The login payload to authenticate with

options: Optional configuration options for the authentication token

returns:  An authentication token that can be used to make authenticated requests to the server

#### Example

```
domain := "example.com"
payload, err := sdk.Auth.Login(domain, nil)

// Generate an authentication token for the logged in wallet
token, err := sdk.Auth.GenerateAuthToken(domain, payload, nil)
```

### func \(\*WalletAuthenticator\) [Login](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/wallet_authenticator.go#L72-L75>)

```go
func (auth *WalletAuthenticator) Login(domain string, options *WalletLoginOptions) (*WalletLoginPayload, error)
```

Client\-side function that allows the connected wallet to login to a server\-side application\. Generates a login payload that can be sent to the server\-side for verification or authentication\.

domain: The domain of the application that you want to log in to

options: Optional configuration options for the login payload

returns: A login payload that can be sent to the server\-side for verification or authentication

#### Example

```
// Add the domain of the application that you want to log in to
domain := "example.com"

// Generate a signed login payload for the connected wallet to authenticate with
payload, err := sdk.Auth.Login(domain, nil)
```

### func \(\*WalletAuthenticator\) [Verify](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/wallet_authenticator.go#L129-L133>)

```go
func (auth *WalletAuthenticator) Verify(domain string, payload *WalletLoginPayload, options *WalletVerifyOptions) (string, error)
```

Server\-side function to securely verify the address of the logged in client\-side wallet by validating the provided client\-side login request\.

domain: The domain of the application to verify the login request for

payload: The login payload to verify

returns: The address of the logged in wallet that signed the payload

#### Example

```
domain := "example.com"
payload, err := sdk.Auth.Login(domain, nil)

// Verify the login request
address, err := sdk.Auth.Verify(domain, payload, nil)
```

## type [WalletLoginOptions](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L218-L222>)

```go
type WalletLoginOptions struct {
    Nonce          string
    ExpirationTime time.Time
    ChainId        int
}
```

## type [WalletLoginPayload](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L232-L235>)

```go
type WalletLoginPayload struct {
    Payload   *WalletLoginPayloadData `json:"payload"`
    Signature string                  `json:"signature"`
}
```

## type [WalletLoginPayloadData](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L224-L230>)

```go
type WalletLoginPayloadData struct {
    Domain         string    `json:"domain"`
    Address        string    `json:"address"`
    Nonce          string    `json:"nonce"`
    ExpirationTime time.Time `json:"expiration_time"`
    ChainId        int       `json:"chain_id"`
}
```

## type [WalletVerifyOptions](<https://github.com/thirdweb-dev/go-sdk/blob/main/thirdweb/types.go#L237-L239>)

```go
type WalletVerifyOptions struct {
    ChainId int
}
```
