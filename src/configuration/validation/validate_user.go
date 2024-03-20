package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator" // as ut
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en" // as en_translations
	"github.com/luiisp/go-uplift/src/configuration/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}

}
func ValidateUsrError(
	val_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonVarErrVal validator.ValidationErrors

	if errors.As(val_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid Field Type.")
	} else if errors.As(val_err, &jsonVarErrVal) {
		errorCauses := []rest_err.Causes{}
		for _, e := range val_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Messages: e.Translate(transl),
				Field:    e.Field(),
			}
			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestErrorValidationError("Some fiels are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error to try convert fields")
	}
}
