package airing

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	animeList, err := arn.GetAiringAnime()

	if err != nil {
		return ctx.Error(500, "Failed fetching airing anime", err)
	}

	sort.Sort(arn.AnimeByPopularity(animeList))
	return ctx.HTML(components.Airing(animeList))
}
