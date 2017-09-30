package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	for anime := range arn.MustStreamAnime() {
		if anime.Popularity != nil {
			continue
		}

		anime.Popularity = &arn.AnimePopularity{}
		arn.PanicOnError(anime.Save())
	}
}
