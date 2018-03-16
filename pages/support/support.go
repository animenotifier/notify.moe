package support

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get support page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	profileLink := "/"

	if user != nil {
		profileLink = "/+" + user.Nick
	}

	return ctx.HTML(components.Support(profileLink, user))
}
