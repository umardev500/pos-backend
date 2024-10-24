package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/umardev500/pos-backend/contracts"
	"github.com/umardev500/pos-backend/models"
)

type validatorUtil struct {
	validate *validator.Validate
}

func NewValidatorUtil(v *validator.Validate) contracts.ValidatorInterface {
	return &validatorUtil{
		validate: v,
	}
}

// parseTag is method to parse tag to message
func (v *validatorUtil) parseTag(tag, field string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s is not a valid email", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, tag)
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, tag)
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, tag)
	case "oneof":
		return fmt.Sprintf("The value for %s must be one of the allowed values", tag)
	case "gt":
		return fmt.Sprintf("%s must be greater than %s", field, tag)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, tag)
	case "lt":
		return fmt.Sprintf("%s must be less than %s", field, tag)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, tag)
	}

	return ""
}

func (v *validatorUtil) Struct(payload interface{}) []models.ErrItem {
	err := v.validate.Struct(payload)
	if err == nil {
		return nil
	}

	var validationErrs []models.ErrItem

	// Handler validation errors
	for _, validationErr := range err.(validator.ValidationErrors) {
		field := validationErr.Field()
		tag := validationErr.Tag()
		validationErr := models.ErrItem{
			Field:   field,
			Message: v.parseTag(tag, field),
			Case:    tag,
		}

		validationErrs = append(validationErrs, validationErr)
	}

	return validationErrs
}
