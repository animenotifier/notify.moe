package utils

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// GetUser returns the logged in user for the given context.
func GetUser(ctx aero.Context) *arn.User {
	return arn.GetUserFromContext(ctx)
}

// SameUser returns "true" or "false" depending on if the users are the same.
func SameUser(a *arn.User, b *arn.User) string {
	if a == nil {
		return "false"
	}

	if b == nil {
		return "false"
	}

	if a.ID == b.ID {
		return "true"
	}

	return "false"
}
