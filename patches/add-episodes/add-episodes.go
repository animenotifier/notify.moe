package main

import (
	"fmt"

	"github.com/animenotifier/arn"
)

func main() {
	count := 0

	for anime := range arn.MustStreamAnime() {
		episodes := anime.Episodes()

		if episodes == nil {
			episodes = &arn.AnimeEpisodes{
				AnimeID: anime.ID,
				Items:   []*arn.AnimeEpisode{},
			}

			if episodes.Save() == nil {
				count++
			}
		}
	}

	fmt.Println("Added empty anime episodes to", count, "anime.")
}
