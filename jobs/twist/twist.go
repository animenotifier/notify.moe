package main

import (
	"fmt"
	"os"
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/twist"
	"github.com/fatih/color"
)

var rateLimiter = time.NewTicker(500 * time.Millisecond)

func main() {
	defer arn.Node.Close()

	// Replace this with ID list from twist.moe later
	twistAnime, err := twist.GetAnimeIndex()
	arn.PanicOnError(err)
	idList := arn.IDList(twistAnime.KitsuIDs())

	// Save index in cache
	arn.DB.Set("IDList", "animetwist index", &idList)

	color.Yellow("Refreshing twist.moe links for %d anime", len(idList))

	for count, animeID := range idList {
		anime, animeErr := arn.GetAnime(animeID)

		if animeErr != nil {
			color.Red("Error fetching anime from the database with ID %s: %v", animeID, animeErr)
			continue
		}

		// Log
		fmt.Fprintf(os.Stdout, "[%d / %d] ", count+1, len(idList))

		// Refresh
		anime.RefreshEpisodes()

		// Ok
		color.Green("Found %d episodes for anime %s", len(anime.Episodes().Items), animeID)

		// Wait for rate limiter
		<-rateLimiter.C
	}
}
