package login

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx aero.Context) error {
	return ctx.HTML(components.Login(""))
}
