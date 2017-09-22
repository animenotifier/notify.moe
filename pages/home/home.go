package home

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

// Get the anime list or the frontpage when logged out.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return frontpage.Get(ctx)
	}

	return ctx.Redirect("/animelist/watching")
	//return AnimeList(ctx, user, arn.AnimeListStatusWatching)
}
