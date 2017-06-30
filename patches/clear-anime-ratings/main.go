package main

import "github.com/animenotifier/arn"

func main() {
	for anime := range arn.MustStreamAnime() {
		anime.Rating.Reset()
		anime.MustSave()
	}
}
