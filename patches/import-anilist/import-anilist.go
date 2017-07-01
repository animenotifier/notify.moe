package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	arn.PanicOnError(arn.AniList.Authorize())
	color.Green(arn.AniList.AccessToken)

	allAnime, err := arn.AllAnime()
	arn.PanicOnError(err)

	count := 0
	stream := arn.AniList.StreamAnime()

	for aniListAnime := range stream {
		anime := arn.FindAniListAnime(aniListAnime, allAnime)

		if anime != nil {
			fmt.Println(aniListAnime.TitleRomaji, "=>", anime.Title.Canonical)
		}

		count++
	}
}
