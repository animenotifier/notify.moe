package utils

import "github.com/animenotifier/arn"

// UserList is a named list of users.
type UserList struct {
	Name  string
	Users []*arn.User
}
