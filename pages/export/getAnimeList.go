package export

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

func getAnimeList(ctx aero.Context) (*arn.AnimeList, error) {
	nick := ctx.Get("nick")
	user := arn.GetUserFromContext(ctx)
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return nil, ctx.Error(http.StatusNotFound, "User not found", err)
	}

	// Fetch all eligible items
	animeList := viewUser.AnimeList()

	if animeList == nil {
		return nil, ctx.Error(http.StatusNotFound, "Anime list not found")
	}

	// Filter private items
	if user == nil || user.ID != viewUser.ID {
		animeList = animeList.WithoutPrivateItems()
	}

	return animeList, nil
}
