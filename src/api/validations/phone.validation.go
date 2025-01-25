package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func IranMobileValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	result, err := regexp.MatchString(`((0?9)|(\+?989))\d{2}\W?\d{3}\W?\d{4}`, value)
	if err != nil {
		return false
	}

	return result
}
