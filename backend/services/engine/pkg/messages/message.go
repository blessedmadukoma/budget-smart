package messages

import "errors"

// Custom error definitions
var (
	ErrNotFound              = errors.New("resource not found")
	ErrConflict              = errors.New("resource conflict")       // conflict due to resource state or concurrent changes or incompatible changes
	ErrExists                = errors.New("resource already exists") // cannot be created again
	ErrNotExists                = errors.New("does not exist")
	ErrUnauthorized          = errors.New("unauthorized access")
	ErrForbidden             = errors.New("access forbidden")
	ErrBadRequest            = errors.New("bad request")
	ErrMissingRequestBody    = errors.New("missing request body")
	ErrInternalServer        = errors.New("internal server error")
	ErrValidation            = errors.New("validation error")
	ErrHashPassword          = errors.New("failed to hash password")
	ErrInvalidCredentials    = errors.New("invalid credentials")
	ErrInvalidPassword    = errors.New("invalid password")
	ErrThrottleTryAgainLater = errors.New("please wait for few minutes, try again later")
)

// WrapError allows adding context to an existing error
func WrapError(err error, message string) error {
	return errors.New(message + ": " + err.Error())
}

var (
	OperationWasSuccessful = "Operation was successful!"
	AccountRejected        = "Your account has been rejected."
	AccountPending         = "Your account is pending approval."
	AccountIsLocked        = "Your account is locked."
	AccountIsRestricted    = "Your account is restricted."
)
