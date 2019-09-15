package main

import (
	"fmt"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	now := time.Now()
	futureThreshold := 8 * 7 * 24 * time.Hour

	for anime := range arn.StreamAnime() {
		// Try to find incorrect airing dates
		for _, episode := range anime.Episodes() {
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
			episode.Save()
		}
	}
}
