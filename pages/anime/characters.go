package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Characters ...
func Characters(ctx *aero.Context) string {
	id := ctx.Get("id")
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.AnimeCharacters(anime))
}
