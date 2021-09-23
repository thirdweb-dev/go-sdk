package nftlabs

import "fmt"

type NotFoundError struct{
	identifier interface{}
	typeName string
}

func (m *NotFoundError) Error() string {
	return fmt.Sprintf("Could not find %v with id %v", m.typeName, m.identifier)
}

type UnmarshalError struct{
	body string
	typeName string
	underlyingError error
}

func (m *UnmarshalError) Error() string {
	return fmt.Sprintf("Could not unmarshal %v with body %v", m.typeName, m.body)
}

type NoSignerError struct{
	typeName string
}

func (m *NoSignerError) Error() string {
	return fmt.Sprintf("Could not proceed with transaction in %v module, missing SigningMethod", m.typeName)
}

type NoAddressError struct{
	typeName string
}

func (m *NoAddressError) Error() string {
	return fmt.Sprintf("Could not proceed with transaction in %v module, missing or invalid signer address", m.typeName)
}
