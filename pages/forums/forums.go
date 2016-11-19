package forums

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/forum"
)

// Get ...
func Get(ctx *aero.Context) string {
	return forum.Get(ctx)
}
