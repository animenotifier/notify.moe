package arn

// NickToUser stores the user ID by nickname.
type NickToUser struct {
	Nick   string `json:"nick" primary:"true"`
	UserID UserID `json:"userId"`
}

// GetID returns the primary key which is the nickname.
func (mapping *NickToUser) GetID() string {
	return mapping.Nick
}
