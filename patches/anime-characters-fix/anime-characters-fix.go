package main

import (
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()
	count := 0

	for anime := range arn.StreamAnime() {
		if anime.Characters() == nil {
			characters := &arn.AnimeCharacters{
				AnimeID: anime.ID,
				Items:   nil,
			}

			characters.Save()
			count++
		}
	}

	color.Green("Added %d missing AnimeCharacter objects", count)
}
