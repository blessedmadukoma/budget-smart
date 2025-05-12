package messages

import "errors"

// Custom error definitions
var (
	ErrNotFound           = errors.New("resource not found")
	ErrConflict           = errors.New("resource conflict")       // conflict due to resource state or concurrent changes or incompatible changes
	ErrExists             = errors.New("resource already exists") // cannot be created again
	ErrUnauthorized       = errors.New("unauthorized access")
	ErrForbidden          = errors.New("access forbidden")
	ErrBadRequest         = errors.New("bad request")
	ErrMissingRequestBody = errors.New("missing request body")
	ErrInternalServer     = errors.New("internal server error")
	ErrValidation         = errors.New("validation error")
	ErrWrongPassword      = errors.New("wrong password")
)

// WrapError allows adding context to an existing error
func WrapError(err error, message string) error {
	return errors.New(message + ": " + err.Error())
}
