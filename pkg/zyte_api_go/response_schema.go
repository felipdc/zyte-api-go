package zyte_api_go

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
