package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ScreenshotOptions struct {
	Format   string `json:"format,omitempty" validate:"omitempty,eq=PNG|eq=JPEG"`
	FullPage bool   `json:"fullPage,omitempty"`
}

type RequestSchema struct {
	Url                 string             `json:"url" validate:"required"`
	HttpResponseBody    bool               `json:"httpResponseBody,omitempty" validate:"required_without_all=HttpRequestText HttpResponseHeaders Screenshot BrowserHtml"`
	HttpResponseHeaders bool               `json:"httpResponseHeaders,omitempty" validate:"required_without_all=HttpRequestText HttpResponseBody Screenshot BrowserHtml"`
	HttpRequestText     string             `json:"httpRequestText,omitempty" validate:"required_without_all=HttpResponseHeaders HttpResponseBody Screenshot BrowserHtml"`
	BrowserHtml         bool               `json:"browserHtml,omitempty" validate:"required_without_all=HttpResponseHeaders HttpResponseBody Screenshot HttpRequestText"`
	Screenshot          bool               `json:"screenshot,omitempty" validate:"required_without_all=HttpResponseHeaders HttpResponseBody BrowserHtml HttpRequestText"`
	HttpRequestMethod   string             `json:"httpRequestMethod,omitempty" validate:"omitempty,eq=POST|eq=GET"`
	HttpRequestBody     []byte             `json:"httpRequestBody,omitempty"`
	Geolocation         string             `json:"geolocation,omitempty" validate:"omitempty,iso3166_1_alpha2"`
	Javascript          bool               `json:"javascript,omitempty"`
	JobId               string             `json:"jobId,omitempty"`
	ScreenshotOptions   *ScreenshotOptions `json:"screenshotOptions,omitempty"`
}

type Options struct {
	schema RequestSchema `validate:"required"`
}

type Extractor struct {
	apiKey string
}

type ResponseSchema struct {
	Url              string `json:"url"`
	StatusCode       uint   `json:"statusCode"`
	HttpResponseBody string `json:"httpResponseBody"`
	Screenshot       string `json:"screenshot"`
	Type             string `json:"string"`
	Title            string `json:"title"`
	Status           int32  `json:"status"`
	Detail           string `json:"detail"`
}

func validateOptions(options Options) [2]error {
	validate := validator.New()
	errsOptions := validate.Struct(options)
	errsSchema := validate.Struct(options.schema)
	errs := [...]error{errsOptions, errsSchema}
	return errs
}

func (e Extractor) Extract(options Options) ResponseSchema {

	errs := validateOptions(options)
	if errs[0] != nil || errs[1] != nil {
		log.Fatalln(errs)
	}

	client := http.Client{}

	b, err := json.Marshal(options.schema)
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

	return resBody
}

func NewExtractor(apiKey string) Extractor {
	return Extractor{apiKey: apiKey}
}
