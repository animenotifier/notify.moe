package arn

// hasID includes an object ID.
type hasID struct {
	ID ID `json:"id" primary:"true"`
}

// GetID returns the ID.
func (obj *hasID) GetID() ID {
	return obj.ID
}
