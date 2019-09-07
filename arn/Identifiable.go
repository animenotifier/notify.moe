package arn

// Identifiable applies to any type that has an ID and exposes it via GetID.
type Identifiable interface {
	GetID() string
}
