package utils

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// GetUser returns the logged in user for the given context.
func GetUser(ctx *aero.Context) *arn.User {
	return arn.GetUserFromContext(ctx)
}
