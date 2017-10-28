package best

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
)

const maxEntries = 7

// Get search page.
func Get(ctx *aero.Context) string {
	return ctx.HTML(components.BestAnime(nil, nil, nil, nil, nil))
}
