package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Fixing non-existing anime relations")
	defer color.Green("Finished.")
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		relations := anime.Relations()

		if relations == nil {
			relations = &arn.AnimeRelations{
				AnimeID: anime.ID,
			}

			relations.Save()
		}
	}
}
