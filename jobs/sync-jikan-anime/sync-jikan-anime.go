package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/jikan"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

const maxRetries = 3

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

	if !jikanDB.Exists("Anime", malID) {
		var anime *jikan.Anime
		var err error

		for try := 1; try <= maxRetries; try++ {
			time.Sleep(time.Second)
			anime, err = jikan.GetAnime(malID)

			if err == nil {
				jikanDB.Set("Anime", malID, anime)
				return
			}

			fmt.Printf("Error fetching %s on try %d: %v", malID, try, err)

			// Wait an additional second
			time.Sleep(time.Second)
		}
	}
}
