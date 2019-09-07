package arn

// NickToUser stores the user ID by nickname.
type NickToUser struct {
	Nick   string `json:"nick" primary:"true"`
	UserID UserID `json:"userId"`
}
