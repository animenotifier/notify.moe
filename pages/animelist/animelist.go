package animelist

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get anime list.
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user := utils.GetUser(ctx)
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	animeList := viewUser.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", nil)
	}

	sort.Slice(animeList.Items, func(i, j int) bool {
		return animeList.Items[i].FinalRating() > animeList.Items[j].FinalRating()
	})

	return ctx.HTML(components.AnimeList(animeList, user))
}
