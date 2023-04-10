package main

import (
	"fmt"

	"github.com/felipdc/zyte-api-go/pkg/zyte_api_go"
)

func main() {
	e := zyte_api_go.NewExtractor("API_KEY")

	// Set options
	schema := zyte_api_go.RequestSchema{
		Url:        "http://books.toscrape.com/",
		Screenshot: true,
		Actions: []zyte_api_go.Action{
			{
				Action: "click",
				Selector: zyte_api_go.Selector{
					Type:  "css",
					Value: "#default > div > div > div > div > section > div:nth-child(2) > ol > li:nth-child(1) > article > div.image_container > a",
				},
			},
		},
	}

	options := zyte_api_go.Options{
		Schema:         schema,
		ScreenshotFile: "screenshot.png",
	}

	// Scrape a website
	response := e.Extract(options)

	// Print the response
	fmt.Println(response)
}
