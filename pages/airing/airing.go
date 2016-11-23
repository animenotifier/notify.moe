package airing

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/jobs"
)

// Get ...
func Get(ctx *aero.Context) string {
	var airingAnimeCache jobs.AiringAnimeCache
	arn.GetObject("Cache", "airingAnime", &airingAnimeCache)
	return ctx.HTML(components.Airing(airingAnimeCache.Anime))
}
