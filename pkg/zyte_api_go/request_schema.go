package zyte_api_go

type ScreenshotOptions struct {
	Format   string `json:"format,omitempty" validate:"omitempty,eq=PNG|eq=JPEG"`
	FullPage bool   `json:"fullPage,omitempty"`
}

type Selector struct {
	Type  string `json:"type" validate:"required,eq=css|xpath"`
	Value string `json:"value" validate:"required"`
}

type Action struct {
	Action      string   `json:"action" validate:"required"`
	Selector    Selector `json:"selector,omitempty"`
	Text        string   `json:"text,omitempty"`
	ElapsedTime int      `json:"elapsedTime,omitempty"`
	Status      string   `json:"status,omitempty"`
	Error       string   `json:"error,omitempty"`
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
	Actions             []Action           `json:"actions,omitempty" validate:"omitempty"`
}
