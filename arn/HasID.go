package arn

// hasID includes an object ID.
type hasID struct {
	ID string `json:"id" primary:"true"`
}

// GetID returns the ID.
func (obj *hasID) GetID() string {
	return obj.ID
}
