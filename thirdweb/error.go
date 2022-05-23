package thirdweb

import "fmt"

type notFoundError struct {
	identifier interface{}
}

func (m *notFoundError) Error() string {
	return fmt.Sprintf("Could not find with id %v", m.identifier)
}

type unmarshalError struct {
	body            string
	typeName        string
	UnderlyingError error
}

func (m *unmarshalError) Error() string {
	return fmt.Sprintf("Could not unmarshal %v with body %v", m.typeName, m.body)
}

type noSignerError struct {
	typeName string
	Err      error
}

func (m *noSignerError) Error() string {
	return fmt.Sprintf("Could not proceed with transaction in %v module, missing SigningMethod", m.typeName)
}

type noAddressError struct {
	typeName string
}

func (m *noAddressError) Error() string {
	return fmt.Sprintf("Could not proceed with transaction in %v module, missing or invalid signer address", m.typeName)
}

type unsupportedFunctionError struct {
	typeName string
	body     string
}

func (m *unsupportedFunctionError) Error() string {
	return fmt.Sprintf("The method you're executing in the %v module is not supported yet. %v", m.typeName, m.body)
}

type failedToUploadError struct {
	statusCode      int
	Payload         interface{}
	UnderlyingError error
}

func (m *failedToUploadError) Error() string {
	return fmt.Sprintf("Failed to upload, status code = %d", m.statusCode)
}
