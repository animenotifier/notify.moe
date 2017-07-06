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

	highPriority := []*arn.Anime{}
	lowPriority := []*arn.Anime{}

	for anime := range arn.MustStreamAnime() {
		if anime.GetMapping("shoboi/anime") == "" {
			continue
		}

		if anime.Status == "current" || anime.Status == "upcoming" {
			highPriority = append(highPriority, anime)
		} else {
			lowPriority = append(lowPriority, anime)
		}
	}

	color.Cyan("High priority queue:")
	refresh(highPriority)

	color.Cyan("Low priority queue:")
	refresh(lowPriority)

	color.Green("Finished.")
}

func refresh(queue []*arn.Anime) {
	for _, anime := range queue {
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
}
