// Package errors defines the common errors
package errors

import "fmt"

// UnknownError indicates that an unknown error has happened
type UnknownError struct {
	Message string
	Err     error
}

// Error returns message for the UnknownError error type
// Returns the formatted error nessage
func (e UnknownError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Unknown error occurred. Error message: %s.", e.Message)
	}

	return fmt.Sprintf("Unknown error occurred. Error message: %s. Error: %s", e.Message, e.Err.Error())
}

// Unwrap returns the err if provided through NewUnknownErrorWithError function, otherwise returns nil
// Returns the unwrapped error if previosuly provided through NewUnknownErrorWithError, otherwise return false
func (e UnknownError) Unwrap() error {
	return e.Err
}

// IsUnknownError indicates whether the error is of type UnknownError
// err: The error to check whethe it is of UnknownError type
// Returns true if the given err is of type UnknownError, otherwise return false
func IsUnknownError(err error) bool {
	_, ok := err.(UnknownError)

	return ok
}

// NewUnknownError creates a new UnknownError error
// Returns the newly created error
func NewUnknownError(message string) error {
	return UnknownError{
		Message: message,
	}
}

// NewUnknownErrorWithError creates a new UnknownError error
// message: The additional message to include
// err: The error to wrap with the new created error
// Returns the newly created error
func NewUnknownErrorWithError(message string, err error) error {
	return UnknownError{
		Message: message,
		Err:     err,
	}
}
