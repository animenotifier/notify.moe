package animelist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Redirect to the full URL including the user nick.
func Redirect(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	return ctx.Redirect("/+" + user.Nick + ctx.URI())
}
