package utils

import "github.com/animenotifier/notify.moe/arn"

// UserList is a named list of users.
type UserList struct {
	Name  string
	Users []*arn.User
}
