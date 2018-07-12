package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request interface {
	Response()
}

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

func Response(get string) Request {

	req := Request{}
	response, err := http.Get(get)

	if err != nil {
		fmt.Printf("The HTTP request failed with err %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		json.Unmarshal(data, &req)
	}
	defer response.Body.Close()

	return req
}

func main() {
	fmt.Println("Starting the application...")
	Response("https://api.coinmarketcap.com/v2/listings/")
}
