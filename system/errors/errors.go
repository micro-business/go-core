// Package errors defines the common errors
package errors

import (
	"fmt"
)

// ArgumentNilError indicates that the provided input argument is invalid.
type ArgumentNilError struct {
	ArgumentName string
	Message      string
	Err          error
}

// Error returns message for the TenantAlreadyExistsError error type
// Returns the error nessage
func (e ArgumentNilError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Argument \"%s\" is nil. Error message: %s.", e.ArgumentName, e.Message)
	}

	return fmt.Sprintf("Argument \"%s\" is nil. Error message: %s. Error: %v", e.ArgumentName, e.Message, e.Err)
}

func (e ArgumentNilError) Unwrap() error {
	return e.Err
}

func IsArgumentNilError(err error) bool {
	_, ok := err.(ArgumentNilError)

	return ok
}

// NewArgumentNilError creates a new ArgumentNilError error
func NewArgumentNilError(argumentName, message string) error {
	return ArgumentNilError{
		ArgumentName: argumentName,
		Message:      message,
	}
}

// NewArgumentNilErrorWithError creates a new ArgumentNilError error
func NewArgumentNilErrorWithError(argumentName, message string, err error) error {
	return ArgumentNilError{
		ArgumentName: argumentName,
		Message:      message,
		Err:          err,
	}
}

// ArgumentError indicates that the provided input argument is invalid.
type ArgumentError struct {
	ArgumentName string
	Message      string
	Err          error
}

// Error returns message for the TenantAlreadyExistsError error type
// Returns the error nessage
func (e ArgumentError) Error() string {
	if e.Err == nil {
		return fmt.Sprintf("Argument \"%s\" is invalid. Error message: %s.", e.ArgumentName, e.Message)
	}

	return fmt.Sprintf("Argument \"%s\" is invalid. Error message: %s. Error: %v", e.ArgumentName, e.Message, e.Err)
}

func (e ArgumentError) Unwrap() error {
	return e.Err
}

func IsArgumentError(err error) bool {
	_, ok := err.(ArgumentError)

	return ok
}

// NewArgumentError creates a new ArgumentError error
func NewArgumentError(argumentName, message string) error {
	return ArgumentError{
		ArgumentName: argumentName,
		Message:      message,
	}
}

// NewArgumentError creates a new ArgumentError error
func NewArgumentErrorWithError(argumentName, message string, err error) error {
	return ArgumentError{
		ArgumentName: argumentName,
		Message:      message,
		Err:          err,
	}
}
