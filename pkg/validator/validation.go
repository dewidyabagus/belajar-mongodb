package validator

import (
	"fmt"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
	once     sync.Once
)

func New() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
	})

	return validate
}

func ValidateStruct(s interface{}) error {
	if err := New().Struct(s); err != nil {
		return translateError(err)
	}

	return nil
}

func translateError(err error) error {
	if err == nil {
		return nil
	}

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(validate, trans)

	var errs []error
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		default:
			errs = append(errs, fmt.Errorf(e.Translate(trans)))

		case "required_if":
			errs = append(errs, fmt.Errorf("%s is a required field", e.Field()))

		case "boolean":
			errs = append(errs, fmt.Errorf("%s is a required field", e.Field()))

		}
	}

	return errs[0]
}
