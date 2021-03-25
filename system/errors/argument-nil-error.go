// Package errors defines the common errors
package errors

import "fmt"

// ArgumentNilError indicates that the provided input argument is invalid.
type ArgumentNilError struct {
	ArgumentName string
	Message      string
	Err          error
}

// Error returns message for the TenantAlreadyExistsError error type
// Returns the formatted error nessage
func (e ArgumentNilError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Argument \"%s\" is nil. Error message: %s.", e.ArgumentName, e.Message)
	}

	return fmt.Sprintf("Argument \"%s\" is nil. Error message: %s. Error: %v", e.ArgumentName, e.Message, e.Err)
}

// Unwrap returns the err if provided through NewArgumentNilErrorWithError function, otherwise returns nil
// Returns the unwrapped error if previosuly provided through NewArgumentNilErrorWithError, otherwise return false
func (e ArgumentNilError) Unwrap() error {
	return e.Err
}

// IsArgumentNilError indicates whether the error is of type ArgumentNilError
// err: The error to check whethe it is of ArgumentNilError type
// Returns true if the given err is of type ArgumentNilError, otherwise return false
func IsArgumentNilError(err error) bool {
	_, ok := err.(ArgumentNilError)

	return ok
}

// NewArgumentNilError creates a new ArgumentNilError error
// argumentName: The argument name
// message: The message to include
// Returns the newly created error
func NewArgumentNilError(argumentName, message string) error {
	return ArgumentNilError{
		ArgumentName: argumentName,
		Message:      message,
	}
}

// NewArgumentNilErrorWithError creates a new ArgumentNilError error
// argumentName: The argument name
// message: The message to include
// err: The error to wrap with the new created error
// Returns the newly created error
func NewArgumentNilErrorWithError(argumentName, message string, err error) error {
	return ArgumentNilError{
		ArgumentName: argumentName,
		Message:      message,
		Err:          err,
	}
}
