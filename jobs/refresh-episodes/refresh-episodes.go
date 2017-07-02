package main

import (
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing episode information for each anime.")

	for anime := range arn.MustStreamAnime() {
		err := anime.RefreshEpisodes()

		if err != nil {
			color.Red(err.Error())
		}
	}

	color.Green("Finished.")
}
