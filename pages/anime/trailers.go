package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/utils"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Trailers Display the trailers of an anime.
func Trailers(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.AnimeTrailers(anime, user))
}
