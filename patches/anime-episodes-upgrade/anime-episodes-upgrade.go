package main

import "github.com/animenotifier/notify.moe/arn"

func main() {
	defer arn.Node.Close()

	for episodes := range arn.StreamAnimeEpisodes() {
		anime := episodes.Anime()
		anime.EpisodeIDs = nil

		for _, episode := range episodes.Items {
			episode.ID = arn.GenerateID("Episode")
			episode.AnimeID = anime.ID
			episode.Save()

			anime.EpisodeIDs = append(anime.EpisodeIDs, episode.ID)
		}

		anime.Save()
	}
}
