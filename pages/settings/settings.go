package settings

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Get settings.
func Get(component func(*arn.User) string) func(aero.Context) error {
	return func(ctx aero.Context) error {
		user := arn.GetUserFromContext(ctx)

		if user == nil {
			return ctx.Error(http.StatusUnauthorized, "Not logged in")
		}

		return ctx.HTML(component(user))
	}
}
