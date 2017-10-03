package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	now := time.Now()
	futureThreshold := 8 * 7 * 24 * time.Hour

	for anime := range arn.MustStreamAnime() {
		modified := false

		// Try to find incorrect airing dates
		for _, episode := range anime.Episodes().Items {
			if episode.AiringDate.Start == "" {
				continue
			}

			startTime, err := time.Parse(time.RFC3339, episode.AiringDate.Start)

			if err == nil && startTime.Sub(now) < futureThreshold {
				continue
			}

			// Definitely wrong airing date on this episode
			fmt.Printf("%s | %s | Ep %d | %s\n", anime.ID, color.YellowString(anime.Title.Canonical), episode.Number, episode.AiringDate.Start)

			// Delete the wrong airing date
			episode.AiringDate.Start = ""
			episode.AiringDate.End = ""

			modified = true
		}

		if modified == true {
			arn.PanicOnError(anime.Episodes().Save())
		}
	}
}
