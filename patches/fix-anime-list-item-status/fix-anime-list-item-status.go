package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Setting list item status to correct value")
	defer arn.Node.Close()

	// Iterate over the stream
	for animeList := range arn.StreamAnimeLists() {
		fmt.Println(animeList.User().Nick)

		for _, item := range animeList.Items {
			if item.Status == arn.AnimeListStatusPlanned && item.Episodes > 0 {
				item.Status = arn.AnimeListStatusWatching
			}

			if item.Anime().Status == "finished" && item.Anime().EpisodeCount != 0 && item.Episodes >= item.Anime().EpisodeCount {
				item.Status = arn.AnimeListStatusCompleted
				item.Episodes = item.Anime().EpisodeCount
			}
		}

		animeList.Save()
	}

	color.Green("Finished.")
}
