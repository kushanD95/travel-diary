package validator

import (
	"github.com/go-playground/validator"
)

var (
	validate *validator.Validate
	// trans    ut.Translator
)

func init() {
	validate = validator.New()
	// trans,_ = SetTransatorForStruct(validate)
}

// func SetTransatorForStruct(validate *validator.Validate) (ut.Translator, error) {
// 	uni := ut.New(en_US.New())
// 	translator,_:= uni.GetTranslator("en_US")
// 	validationErr:= translations.RegisterDefaultTranslation(validate,translator)
// 	return translator, validationErr
// }
