package main

import "github.com/animenotifier/notify.moe/arn"

func main() {
	defer arn.Node.Close()

	for episodes := range arn.StreamAnimeEpisodes() {
		episodes.Sort()
		episodes.Save()
	}
}
