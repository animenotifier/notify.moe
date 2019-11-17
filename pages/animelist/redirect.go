package animelist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Redirect to the full URL including the user nick.
func Redirect(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	return ctx.Redirect(http.StatusTemporaryRedirect, "/+"+user.Nick+ctx.Path())
}
