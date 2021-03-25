// Package errors defines the common errors
package errors

import "fmt"

// NotFoundError indicates that the  with the given email address does not exist
type NotFoundError struct {
	Err error
}

// Error returns message for the NotFoundError error type
// Returns the formatted error nessage
func (e NotFoundError) Error() string {
	if e.Err == nil {
		return "Not found."
	}

	return fmt.Sprintf("Not found. Error: %s", e.Err.Error())
}

// Unwrap returns the err if provided through NotFoundError function, otherwise returns nil
// Returns the unwrapped error if previosuly provided through NotFoundError, otherwise return false
func (e NotFoundError) Unwrap() error {
	return e.Err
}

// IsNotFoundError indicates whether the error is of type NotFoundError
// err: The error to check whethe it is of NotFoundError type
// Returns true if the given err is of type NotFoundError, otherwise return false
func IsNotFoundError(err error) bool {
	_, ok := err.(NotFoundError)

	return ok
}

// NewNotFoundError creates a new NotFoundError error
// Returns the newly created error
func NewNotFoundError() error {
	return NotFoundError{}
}

// NewNotFoundErrorWithError creates a new NotFoundError error
// err: The error to wrap with the new created error
// Returns the newly created error
func NewNotFoundErrorWithError(err error) error {
	return NotFoundError{
		Err: err,
	}
}
