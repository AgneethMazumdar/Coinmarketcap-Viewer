package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	c := CoinData{}
	db := initDB()

	t, _ := template.ParseFiles("viewer.html")
	t.Execute(w, displayCoins(c, db))

	db.Close()
}

func displayCoins(c CoinData, db *sql.DB) []CoinData {

	query := `SELECT * FROM btc`

	rows, err := db.Query(query)
	checkErr(err)

	defer rows.Close()
	checkErr(rows.Err())

	fmt.Println("queryAll rows")

	var coins []CoinData

	for rows.Next() {
		fmt.Println("Test 1")
		err = rows.Scan(&c.Timestamp, &c.Symbol, &c.Name,
			&c.Price, &c.Volume, &c.MarketCap)
		checkErr(err)

		// insert anonymous structs into a slice
		coins = append(coins,
			CoinData{
				Timestamp: c.Timestamp,
				Symbol:    c.Symbol,
				Name:      c.Name,
				Price:     c.Price,
				Volume:    c.Volume,
				MarketCap: c.MarketCap,
			})

		fmt.Println(c.Timestamp, " ", c.Symbol)
	}
	fmt.Println("End queryAll")

	return coins
}

func main() {

	c := CoinData{}

	db := initDB()

	// insert(BTC(), db)
	displayCoins(c, db)

	db.Close()

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8000", nil)

}

// We need to be able to retrieve the data from the database and
// loop it onto a a web page
