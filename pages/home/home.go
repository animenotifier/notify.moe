package home

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

// Get the anime list or the frontpage when logged out.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return frontpage.Get(ctx)
	}

	if !user.HasBasicInfo() {
		return utils.SmartRedirect(ctx, "/welcome")
	}

	return utils.SmartRedirect(ctx, "/+"+user.Nick+"/animelist/watching")
}
