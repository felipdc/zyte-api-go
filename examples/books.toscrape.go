package main

import (
	"fmt"

	"github.com/felipdc/zyte-api-go/pkg/zyte_api_go"
)

func main() {
	e := zyte_api_go.NewExtractor("90f32f1034b5436888ad46d54868f98a")

	// Set options
	schema := zyte_api_go.RequestSchema{
		Url:              "http://books.toscrape.com/",
		HttpResponseBody: true,
	}

	options := zyte_api_go.Options{
		Schema: schema,
	}

	// Scrape a website
	response := e.Extract(options)

	// Print the response
	fmt.Println(response)
}
