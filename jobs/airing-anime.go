package jobs

import (
	"sort"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

// AiringAnimeCache ...
type AiringAnimeCache struct {
	Anime []*arn.Anime `json:"anime"`
}

// AiringAnime ...
func AiringAnime() {
	animeList, err := arn.GetAiringAnime()

	if err != nil {
		color.Red("Failed fetching airing anime")
		color.Red(err.Error())
		return
	}

	sort.Sort(arn.AnimeByPopularity(animeList))

	saveErr := arn.SetObject("Cache", "airingAnime", &AiringAnimeCache{
		Anime: animeList,
	})

	if saveErr != nil {
		color.Red("Error saving airing anime")
		color.Red(saveErr.Error())
		return
	}
}
