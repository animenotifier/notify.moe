package admin

import "github.com/aerogo/aero"
import "github.com/animenotifier/notify.moe/components"

// WebDev ...
func WebDev(ctx *aero.Context) string {
	return ctx.HTML(components.WebDev())
}
