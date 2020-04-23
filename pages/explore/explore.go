package explore

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get the explore page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	return ctx.HTML(components.Explore(user))
}
