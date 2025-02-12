package validator

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

var ValidatorInstance = validator.New()

func ValidateStruct(s interface{}) error {
	err := ValidatorInstance.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				return fmt.Errorf("field '%s' validation failed on the '%s' tag", fieldError.Field(), fieldError.Tag())
			}
		}
		return err
	}
	return nil
}
