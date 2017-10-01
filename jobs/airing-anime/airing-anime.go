package main

import (
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const (
	currentlyAiringBonus     = 5.0
	popularityThreshold      = 5
	popularityPenalty        = 8.0
	watchingPopularityWeight = 0.3
	plannedPopularityWeight  = 0.2
)

func main() {
	color.Yellow("Caching airing anime")

	animeList, err := arn.GetAiringAnime()

	if err != nil {
		color.Red("Failed fetching airing anime")
		color.Red(err.Error())
		return
	}

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

		return scoreA > scoreB
	})

	// Convert to small anime list
	cache := &arn.ListOfIDs{}

	for _, anime := range animeList {
		cache.IDList = append(cache.IDList, anime.ID)
	}

	println(len(cache.IDList))

	saveErr := arn.DB.Set("Cache", "airing anime", cache)

	if saveErr != nil {
		color.Red("Error saving airing anime")
		color.Red(saveErr.Error())
		return
	}

	color.Green("Finished.")
}
