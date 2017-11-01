package main

import (
	"github.com/animenotifier/arn"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		providerID := anime.GetMapping("anilist/anime")
		arn.DB.Delete("AniListToAnime", providerID)
		anime.RemoveMapping("anilist/anime", providerID)
		anime.Save()
	}
}
