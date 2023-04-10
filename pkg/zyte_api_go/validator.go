package zyte_api_go

import (
	"github.com/go-playground/validator/v10"
)

func validateOptions(options Options) [2]error {
	validate := validator.New()
	errsOptions := validate.Struct(options)
	errsSchema := validate.Struct(options.Schema)
	errs := [...]error{errsOptions, errsSchema}
	return errs
}
