package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Moving Kitsu IDs to new IDs")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		kitsuID := anime.GetMapping("kitsu/anime")

		if kitsuID == "" {
			continue
		}

		anime.MoveImageFiles(kitsuID, anime.ID)
	}
}
