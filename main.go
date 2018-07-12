package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Listings struct {
	Data []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Symbol      string `json:"symbol"`
		WebsiteSlug string `json:"website_slug"`
	} `json:"data"`
	Metadata struct {
		Timestamp           int         `json:"timestamp"`
		NumCryptocurrencies int         `json:"num_cryptocurrencies"`
		Error               interface{} `json:"error"`
	} `json:"metadata"`
}

type Tickers struct {
	Data []struct {
		ID                int     `json:"id"`
		Name              string  `json:"name"`
		Symbol            string  `json:"symbol"`
		WebsiteSlug       string  `json:"website_slug"`
		Rank              int     `json:"rank"`
		CirculatingSupply float64 `json:"circulating_supply"`
		TotalSupply       float64 `json:"total_supply"`
		MaxSupply         float64 `json:"max_supply"`
		Quotes            struct {
			USD struct {
				Price            float64 `json:"price"`
				Volume24H        float64 `json:"volume_24h"`
				MarketCap        float64 `json:"market_cap"`
				PercentChange1H  float64 `json:"percent_change_1h"`
				PercentChange24H float64 `json:"percent_change_24h"`
				PercentChange7D  float64 `json:"percent_change_7d"`
			} `json:"USD"`
		} `json:"quotes"`
		LastUpdated int `json:"last_updated"`
	} `json:"data"`
	Metadata struct {
		Timestamp           int         `json:"timestamp"`
		NumCryptocurrencies int         `json:"num_cryptocurrencies"`
		Error               interface{} `json:"error"`
	} `json:"metadata"`
}

func HandleRequest(get string) []byte {
	response, err := http.Get(get)

	if err != nil {
		fmt.Printf("The HTTP request failed with err %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return data
}

func ListingsRes() Listings {

	var listing Listings
	json.Unmarshal(HandleRequest("https://api.coinmarketcap.com/v2/listings/"), &listing)

	return listing
}

func TickerRes() Tickers {

	var ticker Tickers
	json.Unmarshal(HandleRequest("https://api.coinmarketcap.com/v2/ticker/?structure=array"), &ticker)

	return ticker
}

func main() {
	fmt.Println("Starting the application...")
	data := ListingsRes()

	for _, value := range data.Data {
		fmt.Println("ID: ", value.ID, " Name: ", value.Name)
	}
}

// next step
// get a 'shell' database started
// then be able to retrieve data from there (ids & names)
// get a web page up that'll let you pick tokens to get
// constant live updates from
// & perform basic calculations on them? (basic ta indicators)
