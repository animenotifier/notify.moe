package welcome

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get returns the welcome page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	return ctx.HTML(components.Welcome(user))
}
