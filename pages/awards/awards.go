package awards

import "github.com/aerogo/aero"
import "github.com/animenotifier/notify.moe/components"

// Get ...
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Awards())
}
