package nftlabs

type Currency struct {
	Name string
	Symbol string
	Decimals uint8
}

type CurrencyValue struct {
	Currency
	Value string
	DisplayValue uint64
}