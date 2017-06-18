package admin

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get admin page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || user.Role != "admin" {
		return ctx.Redirect("/")
	}

	return ctx.HTML(components.Admin(user))
}
