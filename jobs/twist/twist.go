package main

import (
	"fmt"
	"os"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/twist"
)

var rateLimiter = time.NewTicker(500 * time.Millisecond)

func main() {
	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Replace this with ID list from twist.moe later
	twistAnime, err := twist.GetAnimeIndex()
	arn.PanicOnError(err)
	idList := arn.IDList(twistAnime.KitsuIDs())

	// Kitsu finder
	finder := arn.NewAnimeFinder("kitsu/anime")

	// Save index in cache
	arn.DB.Set("IDList", "animetwist index", &idList)

	color.Yellow("Refreshing twist.moe links for %d anime", len(idList))

	for count, kitsuID := range idList {
		anime := finder.GetAnime(kitsuID)

		if anime == nil {
			color.Red("Error fetching anime from the database with Kitsu ID %s", kitsuID)
			continue
		}

		// Log
		fmt.Fprintf(os.Stdout, "[%d / %d] ", count+1, len(idList))

		if anime.Status != "current" {
			fmt.Println("Not currently airing - skipping")
			continue
		}

		// Refresh
		err := anime.RefreshEpisodes()

		if err != nil {
			color.Red(err.Error())
			continue
		}

		// Ok
		color.Green("Found %d episodes for anime %s (Kitsu: %s)", len(anime.Episodes()), anime.ID, kitsuID)

		// Wait for rate limiter
		<-rateLimiter.C
	}
}
