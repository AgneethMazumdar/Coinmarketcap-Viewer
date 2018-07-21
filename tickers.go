package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func handleRequest(get string) []byte {
	response, err := http.Get(get)

	if err != nil {
		fmt.Printf("The HTTP request failed with err %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return data
}

func TickerRes(t Tickers) Tickers {

	GET := "https://api.coinmarketcap.com/v2/ticker/?structure=array"
	json.Unmarshal(handleRequest(GET), &t)

	return t
}

func BTC() CoinData {
	t := TickerRes(Tickers{})

	btc := CoinData{
		Timestamp: t.Metadata.Timestamp,
		Symbol:    t.Data[0].Symbol,
		Name:      t.Data[0].Name,
		Price:     t.Data[0].Quotes.USD.Price,
		Volume:    t.Data[0].Quotes.USD.Volume24H,
		MarketCap: t.Data[0].Quotes.USD.MarketCap,
	}
	return btc
}
