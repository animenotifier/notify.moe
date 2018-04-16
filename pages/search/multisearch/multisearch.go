package multisearch

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

// Anime search page.
func Anime(ctx *aero.Context) string {
	return ctx.HTML(components.MultiSearch())
}
