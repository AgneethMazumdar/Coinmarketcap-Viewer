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

func HandleRequest(get string) []byte {
	response, err := http.Get(get)

	if err != nil {
		fmt.Printf("The HTTP request failed with err %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	return data
}

func Response(get string) Listings {

	var listing Listings

	data := HandleRequest(get)
	json.Unmarshal(data, &listing)

	return listing
}

func main() {
	fmt.Println("Starting the application...")
	data := Response("https://api.coinmarketcap.com/v2/listings/")

	for _, value := range data.Data {
		fmt.Println("ID: ", value.ID, " Name: ", value.Name)
	}
}
