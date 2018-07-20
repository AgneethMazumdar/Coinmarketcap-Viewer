package main

import (
	_ "github.com/lib/pq"
)

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

func main() {

	//c := CoinData{}

	db := initDB()

	insert(BTC(), db)

	//queryAll(c, db)

	db.Close()
}

// We need to be able to retrieve the data from the database and
// loop it onto a a web page

// Symbol - Name - Price - 24HR Volume - MarketCap

/*

CREATE TABLE btc (
	timestamp int,
	symbol VARCHAR(20),
	name VARCHAR(20),
	price VARCHAR(20),
	volume VARCHAR(20),
	marketcap VARCHAR(20)
)

*/
