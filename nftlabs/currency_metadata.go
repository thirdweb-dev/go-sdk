package nftlabs

type Currency struct {
	Name string
	Symbol string
	Decimals uint64
}

type CurrencyValue struct {
	Currency
	Value string
	DisplayValue uint64
}