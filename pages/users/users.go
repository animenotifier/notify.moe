package users

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	users, err := arn.FilterUsers(func(user *arn.User) bool {
		return user.IsActive() && user.Avatar != ""
	})

	sort.Slice(users, func(i, j int) bool {
		return users[i].Registered < users[j].Registered
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Could not fetch user data", err)
	}

	return ctx.HTML(components.Users(users))
}
