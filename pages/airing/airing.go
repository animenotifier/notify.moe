package airing

import (
	"github.com/aerogo/aero"
)

// Get ...
func Get(ctx *aero.Context) string {
	// airingAnimeCache := new(arn.AiringAnimeCache)
	// err := arn.GetObject("Cache", "airingAnime", airingAnimeCache)

	// if err != nil {
	// 	return ctx.Error(500, "Couldn't fetch airing anime", err)
	// }

	// return ctx.HTML(components.Airing(airingAnimeCache.Anime))
	return ctx.HTML("Coming soon.")
}
