package zyte_api_go

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Extractor struct {
	apiKey string
}

func NewExtractor(apiKey string) Extractor {
	return Extractor{apiKey: apiKey}
}

func (e Extractor) Extract(options Options) ResponseSchema {
	errs := validateOptions(options)
	if errs[0] != nil || errs[1] != nil {
		log.Fatalln(errs)
	}

	client := http.Client{}

	b, err := json.Marshal(options.Schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	req, err := http.NewRequest("POST", "https://api.zyte.com/v1/extract", strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(e.apiKey, "")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	encBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var resBody ResponseSchema

	err = json.Unmarshal(encBody, &resBody)
	if err != nil {
		log.Fatal(err)
	}

	if options.Schema.Screenshot {
		if options.ScreenshotFile != "" {
			// decode base64-encoded screenshot and write to file
			screenshotBytes, err := base64.StdEncoding.DecodeString(resBody.Screenshot)
			if err != nil {
				log.Fatal(err)
			}
			err = ioutil.WriteFile(options.ScreenshotFile, screenshotBytes, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return resBody
}
