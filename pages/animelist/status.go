package animelist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// FilterByStatus returns a handler for the given anime list item status.
func FilterByStatus(status string) aero.Handle {
	return func(ctx *aero.Context) string {
		user := utils.GetUser(ctx)
		list, response := statusList(ctx, status)

		if response != "" {
			return response
		}

		return ctx.HTML(components.ProfileAnimeListFilteredByStatus(list, list.User(), user, status, ctx.URI()))
	}
}

// statusList handles the request for an anime list with a given status.
func statusList(ctx *aero.Context, status string) (*arn.AnimeList, string) {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return nil, ctx.Error(http.StatusNotFound, "User not found", err)
	}

	animeList := viewUser.AnimeList()

	if animeList == nil {
		return nil, ctx.Error(http.StatusNotFound, "Anime list not found", nil)
	}

	watchingList := animeList.FilterStatus(status)
	watchingList.PrefetchAnime()
	watchingList.Sort()

	return watchingList, ""
}
