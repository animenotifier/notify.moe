package main

import "github.com/animenotifier/arn"

func main() {
	defer arn.Node.Close()

	for episodes := range arn.StreamAnimeEpisodes() {
		episodes.Sort()
		episodes.Save()
	}
}
