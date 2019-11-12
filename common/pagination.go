// Package common defines the common types
package common

// The pagination information compatible with graphql-relay connection definition, for more information visit:
// https://facebook.github.io/relay/graphql/connections.htm
type Pagination struct {
	After  *string
	First  *int
	Before *string
	Last   *int
}
