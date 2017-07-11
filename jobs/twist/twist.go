package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/twist"
	"github.com/fatih/color"
)

var rateLimiter = time.NewTicker(500 * time.Millisecond)

func main() {
	// Replace this with ID list from twist.moe later
	twistAnime, err := twist.GetAnimeIndex()
	arn.PanicOnError(err)
	idList := twistAnime.KitsuIDs()

	color.Yellow("Refreshing twist.moe links for %d anime", len(idList))

	for count, animeID := range idList {
		// Wait for rate limiter
		<-rateLimiter.C

		anime, animeErr := arn.GetAnime(animeID)

		if animeErr != nil {
			color.Red("Error fetching anime from the database with ID %s: %v", animeID, animeErr)
			continue
		}

		// Log
		fmt.Fprintf(os.Stdout, "[%d / %d] ", count+1, len(idList))

		// Get twist.moe feed
		feed, err := twist.GetFeedByKitsuID(animeID)

		if err != nil {
			color.Red("Error querying ID %s: %v", animeID, err)
			continue
		}

		episodes := feed.Episodes

		// // Sort by episode number
		// sort.Slice(episodes, func(a, b int) bool {
		// 	return episodes[a].Number < episodes[b].Number
		// })

		for _, episode := range episodes {
			arnEpisode := anime.EpisodeByNumber(episode.Number)

			if arnEpisode == nil {
				color.Red("Anime %s Episode %d not found", anime.ID, episode.Number)
				continue
			}

			if arnEpisode.Links == nil {
				arnEpisode.Links = map[string]string{}
			}

			arnEpisode.Links["twist.moe"] = strings.Replace(episode.Link, "https://test.twist.moe/", "https://twist.moe/", 1)
		}

		arn.PanicOnError(anime.Episodes().Save())
		color.Green("Found %d episodes for anime %s", len(episodes), animeID)
	}
}
