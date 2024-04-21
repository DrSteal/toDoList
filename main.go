package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// type ExchLen struct {
// 	Exch float64 `json:"Value"`
// }

// type AZN struct {
// 	Currency ExchLen `json:"AZN"`
// }

// type Response struct {
// 	Valute AZN `json:"Valute"`
// }

type Product struct {
	Name  string
	Price int
}

var products = []Product{
	{
		Name:  "milk",
		Price: 10000,
	},
	{
		Name:  "tea",
		Price: 80000,
	},
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE  product (
			name text,
			price integer
		)
	`)
	if err != nil {
		panic(err)
	}
}

func main() {

	db, err := sql.Open("sqlite3", "product.db")

	if err != nil {
		panic(err)
	}

	createTable(db)

	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {

		rows, err := db.Query("SELECT name, price FROM product ORDER BY name")

		if err != nil {
			log.Println("query products error,", err)
			return
		}

		var productList []Product

		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.Name, &p.Price); err != nil {
				log.Println("scan products error,", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			productList = append(productList, p)

		}

		json.NewEncoder(w).Encode(productList)
	})

	log.Println("start server")
	http.ListenAndServe(":8000", nil)

	// 	url := "https://www.cbr-xml-daily.ru/daily_json.js"

	// 	resp, err := http.Get(url)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	var response Response

	// 	err = json.NewDecoder(resp.Body).Decode(&response)

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(response.Valute.Currency.Exch)

}
