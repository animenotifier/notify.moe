package main

import "github.com/animenotifier/arn"

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		anime.Rating.Reset()
		anime.Save()
	}
}
