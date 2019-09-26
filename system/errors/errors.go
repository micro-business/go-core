// Package errors defines the common errors
package errors

import "fmt"

// ArgumentNilError indicates that the provided input argument is invalid.
type ArgumentNilError struct {
	ArgumentName string
	ErrorMessage string
	message      string
}

// Error returns message for the TenantAlreadyExistsError error type
// Returns the error nessage
func (e ArgumentNilError) Error() string {
	return e.message
}

// NewArgumentNilError creates a new ArgumentNilError error
func NewArgumentNilError(argumentName, errorMessage string) error {
	return ArgumentNilError{
		ArgumentName: argumentName,
		ErrorMessage: errorMessage,
		message:      fmt.Sprintf("Argument \"%s\" is nil. Error message: %s", argumentName, errorMessage),
	}
}

// ArgumentError indicates that the provided input argument is invalid.
type ArgumentError struct {
	ArgumentName string
	ErrorMessage string
	message      string
}

// Error returns message for the TenantAlreadyExistsError error type
// Returns the error nessage
func (e ArgumentError) Error() string {
	return e.message
}

// NewArgumentError creates a new ArgumentError error
func NewArgumentError(argumentName, errorMessage string) error {
	return ArgumentError{
		ArgumentName: argumentName,
		ErrorMessage: errorMessage,
		message:      fmt.Sprintf("Argument \"%s\" is invalid. Error message: %s", argumentName, errorMessage),
	}
}
