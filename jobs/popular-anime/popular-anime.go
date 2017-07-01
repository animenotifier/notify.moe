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

	// Fetch all anime
	animeList, err := arn.AllAnime()
	arn.PanicOnError(err)

	// Overall
	sort.Slice(animeList, func(i, j int) bool {
		return animeList[i].Rating.Overall > animeList[j].Rating.Overall
	})

	saveAs(animeList[:maxPopularAnime], "best anime overall")

	// Story
	sort.Slice(animeList, func(i, j int) bool {
		return animeList[i].Rating.Story > animeList[j].Rating.Story
	})

	saveAs(animeList[:maxPopularAnime], "best anime story")

	// Visuals
	sort.Slice(animeList, func(i, j int) bool {
		return animeList[i].Rating.Visuals > animeList[j].Rating.Visuals
	})

	saveAs(animeList[:maxPopularAnime], "best anime visuals")

	// Soundtrack
	sort.Slice(animeList, func(i, j int) bool {
		return animeList[i].Rating.Soundtrack > animeList[j].Rating.Soundtrack
	})

	saveAs(animeList[:maxPopularAnime], "best anime soundtrack")

	// Done.
	color.Green("Finished.")
}

// Convert to ListOfIDs and save in cache.
func saveAs(list []*arn.Anime, cacheKey string) {
	cache := &arn.ListOfIDs{}

	for _, anime := range list {
		cache.IDList = append(cache.IDList, anime.ID)
	}

	arn.PanicOnError(arn.DB.Set("Cache", cacheKey, cache))
}
