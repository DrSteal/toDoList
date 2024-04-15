package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExchLen struct {
	Exch float64 `json:"Value"`
}

type AZN struct {
	Currency ExchLen `json:"AZN"`
}

type Response struct {
	Valute AZN `json:"Valute"`
}

// type Product struct {
// 	Name  string
// 	Price int
// }

// var products = []Product{
// 	{
// 		Name:  "Milk",
// 		Price: 4200,
// 	},
// 	{
// 		Name:  "Coffe",
// 		Price: 5000,
// 	},
// }

func main() {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	var response Response

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.Valute.Currency.Exch)

}

// http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
// 	err := json.NewEncoder(w).Encode(products)
// 	if err != nil {
// 		panic(err)
// 	}
// })

// log.Println("server start")
// http.ListenAndServe(":8000", nil)
