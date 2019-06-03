package arn

// Linkable defines an object that can be linked.
type Linkable interface {
	Link() string
}
