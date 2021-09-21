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