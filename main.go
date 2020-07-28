package main

import (
	"net/http"

	"github.com/gerfg/go-druid/handlers"
)

func main() {
	http.HandleFunc("/", handlers.DruidQuery)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
