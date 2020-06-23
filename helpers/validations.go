package helpers


import (
	"../dto"
	"../langs"
	"github.com/go-playground/validator"
)

func GenerateValidationResponse(err error) (response dto.ValidationResponse){
	response.Success = false

	var validations []dto.Validation

	validationErrors := err.(validator.ValidationErrors)

	for _, value := range validationErrors {
		field, rule := value.Field(), value.Tag()

		validation := dto.Validation{Field : field, Message: langs.GenerateValidationMessage(field, rule)}

		validations = append(validations, validation)
	}

	response.Validations = validations

	return response
}