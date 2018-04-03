package user

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/editform"
)

// Edit user.
func Edit(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user := utils.GetUser(ctx)

	if user == nil || user.Role != "admin" {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this user", nil)
	}

	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	return ctx.HTML(editform.Render(viewUser, "Edit user", user))
}
