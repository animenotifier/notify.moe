package main

import (
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
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
		return animeList[i].Rating.Overall > animeList[j].Rating.Overall
	})

	// Convert to small anime list
	cache := &arn.ListOfIDs{}

	for _, anime := range animeList {
		cache.IDList = append(cache.IDList, anime.ID)
	}

	println(len(cache.IDList))

	saveErr := arn.SetObject("Cache", "airing anime", cache)

	if saveErr != nil {
		color.Red("Error saving airing anime")
		color.Red(saveErr.Error())
		return
	}

	color.Green("Finished.")
}
