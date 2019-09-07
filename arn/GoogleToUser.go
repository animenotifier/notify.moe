package arn

// GoogleToUser stores the user ID by Google user ID.
type GoogleToUser struct {
	ID     string `json:"id" primary:"true"`
	UserID UserID `json:"userId"`
}

// GetID returns the ID.
func (mapping *GoogleToUser) GetID() string {
	return mapping.ID
}
