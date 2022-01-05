package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const port = ":8080"

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"cost"`
}

func main() {

	var list []Product

	list = append(list, Product{
		Name:  "Apple",
		Price: 20,
	})
	list = append(list, Product{
		Name:  "Banana",
		Price: 24,
	})

	http.HandleFunc("/products", func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("content-type", "application/json; charset=utf-8")
		data, err := json.Marshal(list)

		if err != nil {
			log.Fatal(err)
		}

		response.Write(data)
	})

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
