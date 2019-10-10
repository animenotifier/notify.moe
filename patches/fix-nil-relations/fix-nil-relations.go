package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Fixing non-existing anime relations")
	defer arn.Node.Close()

	count := 0

	for anime := range arn.StreamAnime() {
		relations := anime.Relations()

		if relations == nil {
			relations = &arn.AnimeRelations{
				AnimeID: anime.ID,
			}

			relations.Save()
			count++
		}
	}

	color.Green("Finished adding %d anime relations objects.", count)
}
