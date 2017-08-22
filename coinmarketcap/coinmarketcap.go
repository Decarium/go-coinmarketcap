package coinmarketcap

import (
	"encoding/json"
	"log"
	"net/http"
)

var (
	url = "https://api.coinmarketcap.com/v1/"
)

//Ticker is the structure of a returned slice of time
type Ticker struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

type GlobalData struct {
	TotalMarketCapUsd            float64 `json:"total_market_cap_usd"`
	Total24HVolumeUsd            float64 `json:"total_24h_volume_usd"`
	BitcoinPercentageOfMarketCap float64 `json:"bitcoin_percentage_of_market_cap"`
	ActiveCurrencies             int     `json:"active_currencies"`
	ActiveAssets                 int     `json:"active_assets"`
	ActiveMarkets                int     `json:"active_markets"`
}

func GetGlobal() (GlobalData, error) {

	globalURL := url + "global/"

	req, err := http.NewRequest(http.MethodGet, tickerURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	global := GlobalData{}

	err = json.NewDecoder(resp.Body).Decode(&global)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return global, err
}

//GetTickers test comment
func GetTickers() ([]Ticker, error) {

	tickerURL := url + "ticker/"

	req, err := http.NewRequest(http.MethodGet, tickerURL, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return nil, err
	}

	defer resp.Body.Close()

	//Define the object to return
	tickers := []Ticker{}

	// Use json.Decode for reading streams of JSON data
	err = json.NewDecoder(resp.Body).Decode(&tickers)
	if err != nil {
		log.Println(err)
	}

	return tickers, err

}
