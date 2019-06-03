package arn

// GoogleToUser stores the user ID by Google user ID.
type GoogleToUser struct {
	ID     string `json:"id"`
	UserID UserID `json:"userId"`
}
