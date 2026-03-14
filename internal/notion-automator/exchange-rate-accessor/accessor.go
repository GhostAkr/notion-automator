package exchangerateaccessor

type ExchangeRateGetter interface {
	GetExchangeRate(baseCurrency string, quoteCurrency string) float32
}
