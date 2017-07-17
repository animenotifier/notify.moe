package main

import (
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const currentlyAiringBonus = 4.0

func main() {
	color.Yellow("Caching airing anime")

	animeList, err := arn.GetAiringAnime()

	if err != nil {
		color.Red("Failed fetching airing anime")
		color.Red(err.Error())
		return
	}

	sort.Slice(animeList, func(i, j int) bool {
		scoreA := animeList[i].Rating.Overall
		scoreB := animeList[j].Rating.Overall

		if animeList[i].Status == "current" {
			scoreA += currentlyAiringBonus
		}

		if animeList[j].Status == "current" {
			scoreB += currentlyAiringBonus
		}

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
