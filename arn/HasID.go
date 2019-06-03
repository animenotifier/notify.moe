package arn

// hasID includes an object ID.
type hasID struct {
	ID string `json:"id"`
}

// GetID returns the ID.
func (obj *hasID) GetID() string {
	return obj.ID
}
