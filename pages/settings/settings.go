package settings

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// Get settings.
func Get(component func(*arn.User) string) func(*aero.Context) string {
	return func(ctx *aero.Context) string {
		user := utils.GetUser(ctx)

		if user == nil {
			return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
		}

		return ctx.HTML(component(user))
	}
}
