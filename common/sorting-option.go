// Package common defines the common types
package common

type SortingDirection int

const (
	Ascending SortingDirection = iota
	Descending
)

type SortingOptionPair struct {
	Name      string
	Direction SortingDirection
}
