package errors

import "fmt"

// CustomError is a struct that holds custom error information.
type CustomError struct {
	Code    int
	Message string
	Err     error
}

// New creates a new CustomError.
func New(code int, message string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Code: %d, Message: %s, Error: %s", e.Code, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

// Unwrap returns the wrapped error.
func (e *CustomError) Unwrap() error {
	return e.Err
}
