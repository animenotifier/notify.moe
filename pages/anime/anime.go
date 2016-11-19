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
		return ctx.Text("Anime not found")
	}

	return ctx.HTML(components.Anime(anime))
}
