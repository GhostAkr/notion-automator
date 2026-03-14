package exchangerateapi

import "testing"

// Important: this is an integration test
func TestGetExchangeRateReturnsValidRate(t *testing.T) {
	// Arrange
	exchangeRateApi := NewExchangeRateApi()

	cases := []struct {
		baseCurrency, quoteCurrency string
	}{
		{"USD", "RSD"},
		{"USD", "TRY"},
		{"EUR", "USD"},
		{"EUR", "RSD"},
		{"EUR", "TRY"},
	}

	// Act & Assert
	for _, c := range cases {
		var rate float32 = exchangeRateApi.GetExchangeRate(c.baseCurrency, c.quoteCurrency)

		if rate == -1.0 {
			t.Errorf("GetExchangeRate(%q, %q) == %f, want a valid rate", 
				c.baseCurrency,
				c.quoteCurrency,
				rate)
		}
	}
}

// Important: this is an integration test
func TestGetExchangeRateReturnsInvalidRateForInvalidCurrency(t *testing.T) {
	// Arrange
	exchangeRateApi := NewExchangeRateApi()

	cases := []struct {
		baseCurrency, quoteCurrency string
	}{
		{"USD", "YYY"},
		{"XXX", "RSD"},
	}

	// Act & Assert
	for _, c := range cases {
		var rate float32 = exchangeRateApi.GetExchangeRate(c.baseCurrency, c.quoteCurrency)

		if rate != -1.0 {
			t.Errorf("GetExchangeRate(%q, %q) == %f, want -1.0", 
				c.baseCurrency,
				c.quoteCurrency,
				rate)
		}
	}
}
