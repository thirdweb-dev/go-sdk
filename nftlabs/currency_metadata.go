package nftlabs

import "math/big"

type CurrencyMetadata struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals uint8 `json:"decimals"`
}

type CurrencyValue struct {
	CurrencyMetadata
	Value        *big.Int `json:"value"`
	DisplayValue string `json:"displayValue"`
}
