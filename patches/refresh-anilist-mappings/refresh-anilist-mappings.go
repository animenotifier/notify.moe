package main

import (
	"github.com/animenotifier/anilist"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Iterating through AniList anime to generate new mappings")
	defer arn.Node.Close()

	err := anilist.Authorize()
	arn.PanicOnError(err)
	color.Green(anilist.AccessToken)

	allAnime := arn.AllAnime()
	count := 0

	for aniListAnime := range anilist.StreamAnime() {
		println(aniListAnime.TitleRomaji)

		anime := arn.FindAniListAnime(aniListAnime, allAnime)

		if anime != nil {
			color.Green("%s %s", anime.ID, aniListAnime.TitleRomaji)
			count++
		} else {
			color.Red("Not found")
		}
	}

	color.Green("%d anime are connected with AniList", count)
}
