// Package middleware implements different go-kit middlewares
package middleware

import "github.com/go-kit/kit/endpoint"

// MiddlewareProviderContract declares the methods to be implemented by the middleware provider service
// that is used to create different instances of go-kit middlewares
type MiddlewareProviderContract interface {
	// CreateLoggingMiddleware creates the logging middleware.
	// endpointName: Mandatory. The name of the endpoint
	// Returns the new endpoint with logging middleware added
	CreateLoggingMiddleware(endpointName string) endpoint.Middleware
}
