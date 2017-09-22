package home

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

// FilterByStatus returns a handler for the given anime list item status.
func FilterByStatus(status string) aero.Handle {
	return func(ctx *aero.Context) string {
		user := utils.GetUser(ctx)

		if user == nil {
			return frontpage.Get(ctx)
		}

		return AnimeList(ctx, user, status)
	}
}

// AnimeList sends the anime list with the given status for given user.
func AnimeList(ctx *aero.Context, user *arn.User, status string) string {
	viewUser := user
	animeList := viewUser.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", nil)
	}

	animeList.PrefetchAnime()
	animeList.Sort()

	return ctx.HTML(components.Home(animeList.FilterStatus(status), viewUser, user, status))
}
