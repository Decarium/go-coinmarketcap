package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
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

func main() {

	now := time.Now().UTC()

	fmt.Println(now)

	test, err2 := GetTickers()

	fmt.Printf("%v", len(test))
	fmt.Printf("%v", err2)

	url := "https://api.coinmarketcap.com/v1/ticker/"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)

	fmt.Printf("%+v", resp.Body)

	tickers := []Ticker{}

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&tickers); err != nil {
		log.Println(err)
	}

	fmt.Printf("%+v", tickers[0])
	fmt.Println("vim-go")
}
