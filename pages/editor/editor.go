package editor

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Redirect("/")
	}

	return ctx.HTML(components.Editor(ctx.URI(), user))
}
