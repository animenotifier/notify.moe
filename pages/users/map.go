package users

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// Map shows a map of all users.
func Map(ctx aero.Context) error {
	return ctx.HTML(components.UserMap(ctx.Path()))
}
