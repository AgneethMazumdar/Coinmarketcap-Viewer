package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/*

SQL Create Table

CREATE TABLE btc (
	timestamp int,
	symbol VARCHAR(20),
	name VARCHAR(20),
	price VARCHAR(20),
	volume VARCHAR(20),
	marketcap VARCHAR(20)
)

*/

type CoinData struct {
	Timestamp int
	Symbol    string
	Name      string
	Price     float64
	Volume    float64
	MarketCap float64
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "coins"
)

func initDB() *sql.DB {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	checkErr(err)

	return db
}

func insert(c CoinData, db *sql.DB) {

	sqlStatement :=
		`INSERT INTO btc (timestamp, symbol, name, price, volume, marketcap) 
		VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(sqlStatement, c.Timestamp, c.Symbol,
		c.Name, c.Price, c.Volume, c.MarketCap)

	checkErr(err)
}

func queryRowID(c CoinData, db *sql.DB, id int) {

	query :=
		`SELECT * FROM characters WHERE character_id = $1`

	stmt, err := db.Prepare(query)
	checkErr(err)

	err = stmt.QueryRow(id).Scan(&c.Timestamp, &c.Symbol, &c.Name,
		&c.Price, &c.Volume, &c.MarketCap)

	checkErr(err)

	fmt.Println(&c.Timestamp, &c.Symbol, &c.Name,
		&c.Price, &c.Volume, &c.MarketCap)
}

func queryAll(c CoinData, db *sql.DB) *sql.Rows {

	query :=
		`SELECT * FROM btc`

	rows, err := db.Query(query)
	checkErr(err)

	defer rows.Close()
	// logAll(rows, c)

	checkErr(rows.Err())

	return rows
}

func logAll(rows *sql.Rows, c CoinData) {

	for rows.Next() {
		err := rows.Scan(&c.Timestamp, &c.Symbol, &c.Name,
			&c.Price, &c.Volume, &c.MarketCap)

		checkErr(err)

		// log.Println(c.Timestamp, c.Symbol, c.Name,
		// 	c.Price, c.Volume, c.MarketCap)
	}
}

func checkErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// change query all to receive a query
// change query all to create and append objects to a slice
