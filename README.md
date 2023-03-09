# Zyte API Go Wrapper
This is a Go wrapper for the Zyte API. Zyte API is a web scraping API that allows you to extract data from websites using antibot and proxy capabilities. This wrapper aims to provide an easy way to call the Zyte API using Go.

**Please note that this wrapper is still under development and may contain bugs.**

## Getting Started
### Prerequisites
- Go 1.16 or higher
- A Zyte API account and API key

### Installation
To install the wrapper, run the following command:

```bash
go get github.com/felipdc/zyte-api-go
```
### Usage
To use the wrapper, you first need to create a client:

```go
import (
    "github.com/felipdc/zyte-api-go"
)

func main() {
    e := NewExtractor("YOUR_API_KEY")
}
```
You can then use the client to make API requests:

```go
// Set options
schema := RequestSchema{
    Url:              "http://books.toscrape.com/",
    HttpResponseBody: true,
}

// Scrape a website
response := e.Extract(options)

// Print the response
fmt.Println(response)
```

## Parameters development
Below is a table describing which parameter is currently supported.

| Parameter                | Done |
|--------------------------|------|
| requestHeaders           | ❌    |
| httpRequestMethod        | ✅    |
| httpRequestBody          | ✅    |
| httpRequestText          | ✅    |
| customHttpRequestHeaders | ❌    |
| httpResponseBody         | ✅    |
| httpResponseHeaders      | ✅    |
| browserHtml              | ✅    |
| screenshot               | ✅    |
| screenshotOptions        | ✅    |
| geolocation              | ✅    |
| javascript               | ✅    |
| actions                  | ❌    |
| jobId                    | ✅    |
| echoData                 | ❌    |
| experimental             | ❌    |

For more information about ZyteAPI, please refer to the [documentation](https://docs.zyte.com/zyte-api/reference/http.html#).