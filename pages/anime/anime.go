package anime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	id, _ := ctx.GetInt("id")
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(404, "Anime not found", err)
	}

	return ctx.HTML(components.Anime(anime))
}
