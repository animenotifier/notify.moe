package utils

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// GetUser ...
func GetUser(ctx *aero.Context) *arn.User {
	userID := ctx.Session().GetString("userId")

	if userID == "" {
		return nil
	}

	user, err := arn.GetUser(userID)

	if err != nil {
		return nil
	}

	return user
}
