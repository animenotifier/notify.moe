package users

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Users())
}
