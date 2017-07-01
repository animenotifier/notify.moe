package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	for anime := range arn.MustStreamAnime() {
		providerID := anime.GetMapping("anilist/anime")
		_, err := arn.DB.Delete("AniListToAnime", providerID)
		arn.PanicOnError(err)
		anime.RemoveMapping("anilist/anime", providerID)
		arn.PanicOnError(anime.Save())
	}
}
