package exchangerateapi

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ExchangeRateApiLatestResponse struct {
	Result             string             `json:"result"`
	Documentation      string             `json:"documentation"`
	TermsOfUse         string             `json:"terms_of_use"`
	TimeLastUpdateUnix int64              `json:"time_last_update_unix"`
	TimeLastUpdateUTC  string             `json:"time_last_update_utc"`
	TimeNextUpdateUnix int64              `json:"time_next_update_unix"`
	TimeNextUpdateUTC  string             `json:"time_next_update_utc"`
	BaseCode           string             `json:"base_code"`
	ConversionRates    map[string]float32 `json:"conversion_rates"`
}

type ExchangeRateApi struct {
	baseUrl string
	apiKey  string
}

func NewExchangeRateApi() *ExchangeRateApi {
	godotenv.Load()

	return &ExchangeRateApi{
		baseUrl: "https://v6.exchangerate-api.com/v6",
		apiKey:  os.Getenv("EXCHANGERATE_API_KEY"),
	}
}

func (api ExchangeRateApi) GetExchangeRate(baseCurrency string, quoteCurrency string) float32 {
	requestUrl := api.baseUrl + "/" + api.apiKey + "/latest/" + baseCurrency

	resp, err := http.Get(requestUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return -1.0
	} else if resp.StatusCode != http.StatusOK {
		panic("Unexpected status: " + resp.Status)
	}

	var latestResponse ExchangeRateApiLatestResponse
	err = json.NewDecoder(resp.Body).Decode(&latestResponse)
	if err != nil {
		panic(err)
	}

	var quoteCurrencies map[string]float32 = latestResponse.ConversionRates
	rate, ok := quoteCurrencies[quoteCurrency]

	if ok {
		return rate
	} else {
		return -1.0
	}
}
