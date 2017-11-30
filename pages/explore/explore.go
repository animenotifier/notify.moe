package explore

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get ...
func Get(ctx *aero.Context) string {
	year := "2017"
	status := "current"
	typ := "tv"
	results := filterAnime(year, status, typ)
	user := utils.GetUser(ctx)

	return ctx.HTML(components.ExploreAnime(results, year, status, typ, user))
}

// Filter ...
func Filter(ctx *aero.Context) string {
	year := ctx.Get("year")
	status := ctx.Get("status")
	typ := ctx.Get("type")
	user := utils.GetUser(ctx)

	results := filterAnime(year, status, typ)

	return ctx.HTML(components.ExploreAnime(results, year, status, typ, user))
}

func filterAnime(year, status, typ string) []*arn.Anime {
	var results []*arn.Anime

	for anime := range arn.StreamAnime() {
		if len(anime.StartDate) < 4 {
			continue
		}

		if anime.StartDate[:4] != year {
			continue
		}

		if anime.Status != status {
			continue
		}

		if anime.Type != typ {
			continue
		}

		results = append(results, anime)
	}

	arn.SortAnimeByQuality(results, status)
	return results
}
