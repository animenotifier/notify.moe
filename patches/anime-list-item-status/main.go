package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Setting list item status to correct value")

	// Get a stream of all anime lists
	allAnimeLists, err := arn.StreamAnimeLists()

	if err != nil {
		panic(err)
	}

	// Iterate over the stream
	for animeList := range allAnimeLists {
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

		err := animeList.Save()
		arn.PanicOnError(err)
	}

	color.Green("Finished.")
}
