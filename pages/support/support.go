package support

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get support page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	profileLink := "/"

	if user != nil {
		profileLink = "/+" + user.Nick
	}

	return ctx.HTML(components.Support(profileLink, user))
}
