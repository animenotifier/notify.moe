package main

import (
	"github.com/animenotifier/anilist"
	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	arn.PanicOnError(anilist.Authorize())
	color.Green(anilist.AccessToken)

	allAnime, err := arn.AllAnime()
	arn.PanicOnError(err)

	count := 0
	stream := anilist.StreamAnime()

	for aniListAnime := range stream {
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
