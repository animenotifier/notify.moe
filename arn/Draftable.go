package arn

// Draftable describes a type where drafts can be created.
type Draftable interface {
	GetIsDraft() bool
	SetIsDraft(bool)
}
