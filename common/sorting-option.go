// Package common defines the common types
package common

type SortingDirection int

const (
	Acsending SortingDirection = iota
	Descending
)

type SortingOptionPair struct {
	Name      string
	Direction SortingDirection
}
