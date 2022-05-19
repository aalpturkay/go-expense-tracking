package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func Enum(fl validator.FieldLevel) bool {
	enumString := fl.Param()
	value := fl.Field().String()
	enumSlice := strings.Split(enumString, "_")
	for _, v := range enumSlice {
		if value == v {
			return true
		}
	}
	return false
}

func ValidateErrors(requestError error) string {
	return validate(requestError.(validator.ValidationErrors))
}

func validate(errors validator.ValidationErrors) string {
	resultErrors := ""
	for _, err := range errors {
		switch err.Tag() {
		case "required":
			resultErrors += err.Field() + " alanı gereklidir\n "
		case "email":
			resultErrors += err.Field() + " geçerli bir email olmalıdır\n "
		case "min":
			resultErrors += err.Field() + " en az " + err.Param() + " karakterden oluşmalı\n"
		case "Enum":
			replacer := *strings.NewReplacer("_", ",")
			resultErrors += err.Field() + " Bu alan bunlardan biri olmalı: " + replacer.Replace(err.Param())

		default:
			resultErrors += "Bu alanda hata: " + err.Tag()
		}
	}
	return resultErrors
}
