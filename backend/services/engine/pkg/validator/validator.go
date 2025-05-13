// internal/pkg/validator/validator.go
package validator

import (
	"errors"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validator is the interface for the validation service
type Validator interface {
	ValidateStruct(data interface{}) error
}

// CustomValidator implements the Validator interface
type CustomValidator struct {
	validate *validator.Validate
}

// NewValidator creates a new instance of the CustomValidator
func NewValidator() Validator {
	v := &CustomValidator{
		validate: validator.New(),
	}
	return v
}

// ValidateStruct validates any struct with validation tags
func (v *CustomValidator) ValidateStruct(data interface{}) error {
	if data == nil {
		return errors.New("data is nil")
	}

	// Ensure data is a struct or pointer to a struct
	value := reflect.ValueOf(data)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}
	if value.Kind() != reflect.Struct {
		return errors.New("data must be a struct or pointer to a struct")
	}

	err := v.validate.Struct(data)
	if err == nil {
		return nil
	}

	var valErrors ValidationErrors
	for _, err := range err.(validator.ValidationErrors) {
		field := err.Field()
		var errorMsg string
		switch err.Tag() {
		case "required":
			errorMsg = field + " is required"
		case "email":
			errorMsg = field + " must be a valid email address"
		case "min":
			errorMsg = field + " must be at least " + err.Param() + " characters long"
		case "oneof":
			errorMsg = field + " must be one of " + err.Param()
		case "required_if":
			errorMsg = field + " is required when " + err.Param()
		default:
			errorMsg = field + " is invalid"
		}
		valErrors = append(valErrors, ValidationError{Field: field, Error: errorMsg})
	}

	return valErrors
}

// ValidationError represents a single validation error
type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// ValidationErrors is a collection of validation errors
type ValidationErrors []ValidationError

// Error implements the error interface for ValidationErrors
func (ve ValidationErrors) Error() string {
	var messages []string
	for _, err := range ve {
		messages = append(messages, err.Field+": "+err.Error)
	}
	return strings.Join(messages, "; ")
}
