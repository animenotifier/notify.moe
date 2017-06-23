package main

import (
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)


const maxPopularAnime = 10

// Note this is using the airing-anime as a template with modfications
// made to it.
func main() {
	color.Yellow("Caching popular anime")

	animeChan, err := arn.AllAnime()

	if err != nil {
		color.Red("Failed fetching anime channel")
		color.Red(err.Error())
		return
	}

	var animeList []*arn.Anime

	for anime := range animeChan {
		animeList = append(animeList, anime)
	}

	sort.Slice(animeList, func(i, j int) bool {
		return animeList[i].Rating.Overall > animeList[j].Rating.Overall
	})
	
	// Change size of anime list to 10
	animeList = animeList[:maxPopularAnime]

	// Convert to small anime list
	cache := &arn.ListOfIDs{}

	for _, anime := range animeList {
		cache.IDList = append(cache.IDList, anime.ID)
	}

	println(len(cache.IDList))

	saveErr := arn.DB.Set("Cache", "popular anime", cache)

	if saveErr != nil {
		color.Red("Error saving popular anime")
		color.Red(saveErr.Error())
		return
	}

	color.Green("Finished.")
}
