package anime

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get anime page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return ctx.HTML(components.Anime(anime, user))
}
