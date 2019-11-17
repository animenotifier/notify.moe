package user

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/pages/profile"
)

// Get redirects /+ to /+UserName
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusBadRequest, "Not logged in")
	}

	if user.Nick == "" {
		return ctx.Error(http.StatusInternalServerError, "User did not set a nickname")
	}

	return profile.Profile(ctx, user)
}
