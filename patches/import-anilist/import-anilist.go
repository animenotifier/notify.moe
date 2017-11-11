package main

import (
	"github.com/animenotifier/anilist"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	arn.PanicOnError(anilist.Authorize())
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
