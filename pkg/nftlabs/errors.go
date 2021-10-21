package nftlabs

import "fmt"

type NotFoundError struct {
	identifier interface{}
	typeName   string
}

func (m *NotFoundError) Error() string {
	return fmt.Sprintf("Could not find %v with id %v", m.typeName, m.identifier)
}

type UnmarshalError struct {
	body            string
	typeName        string
	UnderlyingError error
}

func (m *UnmarshalError) Error() string {
	return fmt.Sprintf("Could not unmarshal %v with body %v", m.typeName, m.body)
}

type NoSignerError struct {
	typeName string
	Err      error
}

func (m *NoSignerError) Error() string {
	return fmt.Sprintf("Could not proceed with transaction in %v module, missing SigningMethod", m.typeName)
}

type NoAddressError struct {
	typeName string
}

func (m *NoAddressError) Error() string {
	return fmt.Sprintf("Could not proceed with transaction in %v module, missing or invalid signer address", m.typeName)
}

type UnsupportedFunctionError struct {
	typeName string
	body     string
}

func (m *UnsupportedFunctionError) Error() string {
	return fmt.Sprintf("The method you're executing in the %v module is not supported yet. %v", m.typeName, m.body)
}

type FailedToUploadError struct {
	statusCode int
	Payload interface{}
	UnderlyingError error
}

func (m *FailedToUploadError) Error() string {
	return fmt.Sprintf("Failed to upload, status code = %d", m.statusCode)
}
