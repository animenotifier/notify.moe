package animelist

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get anime list.
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	animeList := viewUser.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", err)
	}

	sort.Slice(animeList.Items, func(i, j int) bool {
		return animeList.Items[i].FinalRating() < animeList.Items[j].FinalRating()
	})

	return ctx.HTML(components.AnimeList(animeList))
}
