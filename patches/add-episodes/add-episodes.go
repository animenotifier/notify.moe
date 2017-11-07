package main

import (
	"fmt"

	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		if anime.Episodes() != nil {
			continue
		}

		fmt.Println(anime)

		episodes := &arn.AnimeEpisodes{
			AnimeID: anime.ID,
			Items:   []*arn.AnimeEpisode{},
		}

		arn.DB.Set("AnimeEpisodes", anime.ID, episodes)
	}
}
