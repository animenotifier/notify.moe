package arn

// EmailToUser stores the user ID for an email address.
type EmailToUser struct {
	Email  string `json:"email" primary:"true"`
	UserID UserID `json:"userId"`
}
