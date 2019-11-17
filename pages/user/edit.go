package user

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit user.
func Edit(ctx aero.Context) error {
	nick := ctx.Get("nick")
	user := arn.GetUserFromContext(ctx)

	if user == nil || user.Role != "admin" {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this user")
	}

	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	return ctx.HTML(editform.Render(viewUser, "Edit user", user))
}
