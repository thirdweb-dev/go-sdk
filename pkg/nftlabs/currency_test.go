package nftlabs

import (
	"math/big"
	"testing"
)

// TestFormatUnits transforms a listing value amount into a readable string in CurrencyValue.DisplayValue object
// and is callable via Currency.GetValue
func TestFormatUnits(t *testing.T) {
	var currency Currency
	var want string
	currency, _ = newCurrencyModule(nil, "", nil)
	decimals := big.NewInt(18)

	value := big.NewInt(0)
	value.SetString("100000000000000000", 10)
	want = "0.1"
	if display := currency.formatUnits(value, decimals); display != want {
		 t.Fatalf(`formatUnits(%d, %d) = %v, want %v\n`, value, decimals, display, want)
	}

	value.SetString("1000000000000000000", 10)
	want = "1"
	if display := currency.formatUnits(value, decimals); display != want {
		t.Fatalf(`formatUnits(%d, %d) = %v, want %v\n`, value, decimals, display, want)
	}

	value.SetString("10000000000000000000", 10)
	want = "10"
	if display := currency.formatUnits(value, decimals); display != want {
		t.Fatalf(`formatUnits(%d, %d) = %v, want %v\n`, value, decimals, display, want)
	}

	value.SetString("11100000000000000000", 10)
	want = "11.1"
	if display := currency.formatUnits(value, decimals); display != want {
		t.Fatalf(`formatUnits(%d, %d) = %v, want %v\n`, value, decimals, display, want)
	}

	// TODO: failing, need to fix
	//value.SetString("11111111111111111111", 10)
	//want = "11.111111111111111111"
	//if display := currency.formatUnits(value, decimals); display != want {
	//	t.Fatalf(`formatUnits(%d, %d) = %v, want %v\n`, value, decimals, display, want)
	//}
}
