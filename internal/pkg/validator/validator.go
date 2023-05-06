package validator

import (
	"regexp"
	"sync"

	goValidator "github.com/go-playground/validator/v10"
)

var singleValidator sync.Once

var (
	validator *goValidator.Validate
)

func init() {
	// do only once
	singleValidator.Do(func() {
		validator = goValidator.New()
		validator.RegisterValidation("date_iso8601", ValidateDateIso8601)
	})
}

func GetValidator() *goValidator.Validate {
	return validator
}

func ValidateDateIso8601(fl goValidator.FieldLevel) bool {
	var re = regexp.MustCompile(`^\d{4}\-(0[1-9]|1[012])\-(0[1-9]|[12][0-9]|3[01])$`)
	return re.MatchString(fl.Field().String())
}
