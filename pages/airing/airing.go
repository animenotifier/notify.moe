package airing

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get ...
func Get(ctx *aero.Context) string {
	var animeList []*arn.Anime

	scan := make(chan *arn.Anime)
	arn.Scan("Anime", scan)

	for anime := range scan {
		if anime.AiringStatus != "currently airing" || anime.Adult {
			continue
		}

		animeList = append(animeList, anime)
	}

	sort.Sort(arn.AnimeByPopularity(animeList))

	return ctx.HTML(components.Airing(animeList))
}
