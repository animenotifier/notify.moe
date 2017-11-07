package explore

import (
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const (
	currentlyAiringBonus      = 5.0
	popularityThreshold       = 5
	popularityPenalty         = 8.0
	watchingPopularityWeight  = 0.3
	completedPopularityWeight = 0.3
	plannedPopularityWeight   = 0.2
)

// Get ...
func Get(ctx *aero.Context) string {
	year := "2017"
	status := "current"
	typ := "tv"
	results := filterAnime(year, status, typ)

	return ctx.HTML(components.ExploreAnime(results, year, status, typ))
}

// Filter ...
func Filter(ctx *aero.Context) string {
	year := ctx.Get("year")
	status := ctx.Get("status")
	typ := ctx.Get("type")

	results := filterAnime(year, status, typ)

	return ctx.HTML(components.ExploreAnime(results, year, status, typ))
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

	sortAnime(results)
	return results
}

func sortAnime(animeList []*arn.Anime) {
	sort.Slice(animeList, func(i, j int) bool {
		a := animeList[i]
		b := animeList[j]

		scoreA := a.Rating.Overall
		scoreB := b.Rating.Overall

		if a.Status == "current" {
			scoreA += currentlyAiringBonus
		}

		if b.Status == "current" {
			scoreB += currentlyAiringBonus
		}

		if a.Popularity.Total() < popularityThreshold {
			scoreA -= popularityPenalty
		}

		if b.Popularity.Total() < popularityThreshold {
			scoreB -= popularityPenalty
		}

		scoreA += float64(a.Popularity.Watching) * watchingPopularityWeight
		scoreB += float64(b.Popularity.Watching) * watchingPopularityWeight

		scoreA += float64(a.Popularity.Planned) * plannedPopularityWeight
		scoreB += float64(b.Popularity.Planned) * plannedPopularityWeight

		scoreA += float64(a.Popularity.Completed) * completedPopularityWeight
		scoreB += float64(b.Popularity.Completed) * completedPopularityWeight

		return scoreA > scoreB
	})
}
