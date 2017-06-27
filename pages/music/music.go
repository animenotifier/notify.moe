package music

import "github.com/aerogo/aero"
import "github.com/animenotifier/notify.moe/components"

// Get renders the music page.
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.Music())
}
