package arn

// HasDraft includes a boolean indicating whether the object is a draft.
type hasDraft struct {
	IsDraft bool `json:"isDraft" editable:"true"`
}

// GetIsDraft tells you whether the object is a draft or not.
func (obj *hasDraft) GetIsDraft() bool {
	return obj.IsDraft
}

// SetIsDraft sets the draft state for this object.
func (obj *hasDraft) SetIsDraft(isDraft bool) {
	obj.IsDraft = isDraft
}
