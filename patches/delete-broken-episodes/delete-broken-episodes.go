package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Fixing non-existing anime relations")
	defer arn.Node.Close()

	count := 0

	for episode := range arn.StreamEpisodes() {
		anime := episode.Anime()

		if anime == nil {
			color.Yellow(episode.AnimeID)
			_ = episode.Delete()
			count++
		}
	}

	color.Green("Finished deleting %d episodes.", count)
}
