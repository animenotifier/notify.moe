package airing

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	airingAnimeCache := new(arn.AiringAnimeCache)
	err := arn.GetObject("Cache", "airingAnime", airingAnimeCache)

	if err != nil {
		return ctx.Error(500, "Couldn't fetch airing anime", err)
	}

	return ctx.HTML(components.Airing(airingAnimeCache.Anime))
}
