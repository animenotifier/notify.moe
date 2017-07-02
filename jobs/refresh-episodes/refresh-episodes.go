package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Refreshing episode information for each anime.")

	for anime := range arn.MustStreamAnime() {
		episodeCount := len(anime.Episodes)

		err := anime.RefreshEpisodes()

		if err != nil {
			if strings.Contains(err.Error(), "missing a Shoboi ID") {
				continue
			}

			color.Red(err.Error())
		} else {
			fmt.Println(anime.ID, "|", anime.Title.Canonical, "+"+strconv.Itoa(len(anime.Episodes)-episodeCount))
		}
	}

	color.Green("Finished.")
}
