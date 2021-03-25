// Package errors defines the common errors
package errors

import "fmt"

// ArgumentError indicates that the provided input argument is invalid.
type ArgumentError struct {
	ArgumentName string
	Message      string
	Err          error
}

// Error returns message for the TenantAlreadyExistsError error type
// Returns the formatted error nessage
func (e ArgumentError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Argument \"%s\" is invalid. Error message: %s.", e.ArgumentName, e.Message)
	}

	return fmt.Sprintf("Argument \"%s\" is invalid. Error message: %s. Error: %v", e.ArgumentName, e.Message, e.Err)
}

// Unwrap returns the err if provided through NewArgumentErrorWithError function, otherwise returns nil
// Returns the unwrapped error if previosuly provided through NewArgumentErrorWithError, otherwise return false
func (e ArgumentError) Unwrap() error {
	return e.Err
}

// IsArgumentError indicates whether the error is of type ArgumentError
// err: The error to check whethe it is of ArgumentError type
// Returns true if the given err is of type ArgumentError, otherwise return false
func IsArgumentError(err error) bool {
	_, ok := err.(ArgumentError)

	return ok
}

// NewArgumentError creates a new ArgumentError error
// argumentName: The argument name
// message: The message to include
// Returns the newly created error
func NewArgumentError(argumentName, message string) error {
	return ArgumentError{
		ArgumentName: argumentName,
		Message:      message,
	}
}

// NewArgumentErrorWithError creates a new ArgumentError error
// argumentName: The argument name
// message: The message to include
// err: The error to wrap with the new created error
// Returns the newly created error
func NewArgumentErrorWithError(argumentName, message string, err error) error {
	return ArgumentError{
		ArgumentName: argumentName,
		Message:      message,
		Err:          err,
	}
}
