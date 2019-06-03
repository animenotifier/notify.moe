package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Fixing broken anime relations")

	defer arn.Node.Close()

	count := 0

	for relations := range arn.StreamAnimeRelations() {
		brokenIDs := []string{}

		for _, item := range relations.Items {
			_, err := arn.GetAnime(item.AnimeID)

			if err != nil {
				brokenIDs = append(brokenIDs, item.AnimeID)
			}
		}

		for _, brokenID := range brokenIDs {
			relations.Remove(brokenID)
			count++
		}

		relations.Save()
	}

	color.Green("Finished removing %d broken relations.", count)
}
