package admin

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// WebDev ...
func WebDev(ctx aero.Context) error {
	return ctx.HTML(components.WebDev())
}
