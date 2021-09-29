package nftlabs

type CurrencyMetadata struct {
	Name string
	Symbol string
	Decimals uint8
}

type CurrencyValue struct {
	CurrencyMetadata
	Value string
	DisplayValue uint64
}
