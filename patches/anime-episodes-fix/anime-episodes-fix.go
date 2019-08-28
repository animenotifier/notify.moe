package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		anime.EpisodeIDs = nil
		anime.Save()
	}

	for episode := range arn.StreamEpisodes() {
		anime := episode.Anime()
		anime.EpisodeIDs = append(anime.EpisodeIDs, episode.ID)
	}

	for anime := range arn.StreamAnime() {
		color.Yellow(anime.Title.Canonical)
		episodes := anime.Episodes()
		episodes.Sort()

		for _, episode := range episodes {
			fmt.Println(episode.Number, episode.ID)
			anime.EpisodeIDs = append(anime.EpisodeIDs, episode.ID)
		}

		fmt.Println()
		anime.Save()
	}
}
