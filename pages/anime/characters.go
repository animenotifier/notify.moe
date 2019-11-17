package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/components"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Characters ...
func Characters(ctx aero.Context) error {
	id := ctx.Get("id")
	user := arn.GetUserFromContext(ctx)
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.AnimeCharacters(anime, user, true))
}
