// Package errors defines the common errors
package errors

import "fmt"

// AlreadyExistsError indicates that the  with the given information already exists
type AlreadyExistsError struct {
	Err error
}

// Error returns message for the AlreadyExistsError error type
// Returns the formatted error nessage
func (e AlreadyExistsError) Error() string {
	if e.Err == nil {
		return "Already exists."
	}

	return fmt.Sprintf("Already exists. Error: %s", e.Err.Error())
}

// Unwrap returns the err if provided through NewAlreadyExistsErrorWithError function, otherwise returns nil
// Returns the unwrapped error if previosuly provided through NewAlreadyExistsErrorWithError, otherwise return false
func (e AlreadyExistsError) Unwrap() error {
	return e.Err
}

// IsAlreadyExistsError indicates whether the error is of type AlreadyExistsError
// err: The error to check whethe it is of AlreadyExistsError type
// Returns true if the given err is of type AlreadyExistsError, otherwise return false
func IsAlreadyExistsError(err error) bool {
	_, ok := err.(AlreadyExistsError)

	return ok
}

// NewAlreadyExistsError creates a new AlreadyExistsError error
// Returns the newly created error
func NewAlreadyExistsError() error {
	return AlreadyExistsError{}
}

// NewAlreadyExistsErrorWithError creates a new AlreadyExistsError error
// err: The error to wrap with the new created error
// Returns the newly created error
func NewAlreadyExistsErrorWithError(err error) error {
	return AlreadyExistsError{
		Err: err,
	}
}
