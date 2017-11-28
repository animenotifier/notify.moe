package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/jikan"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var jikanDB = arn.Node.Namespace("jikan")

func main() {
	color.Yellow("Syncing with Jikan API")
	defer arn.Node.Close()

	count := 0

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID != "" {
			sync(anime, malID)
			count++
		}
	}

	color.Green("Finished syncing %d anime.", count)

	// Give OS some time to write buffers, just to be safe
	time.Sleep(10 * time.Second)
}

func sync(anime *arn.Anime, malID string) {
	fmt.Printf("%s %s (MAL: %s)\n", anime.ID, anime.Title.Canonical, malID)

	if jikanDB.Exists("Anime", malID) {
		return
	}

	time.Sleep(500 * time.Millisecond)
	jikanAnime, err := jikan.GetAnime(malID)

	if err == nil {
		jikanDB.Set("Anime", malID, jikanAnime)
		return
	}

	fmt.Printf("Error fetching %s: %v\n", malID, err)
}
