package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Episodes ...
func Episodes(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")

	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.AnimeEpisodes(anime.Episodes().Items, user))
}
