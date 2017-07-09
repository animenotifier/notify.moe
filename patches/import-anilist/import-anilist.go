package main

import (
	"fmt"

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
		anime := arn.FindAniListAnime(aniListAnime, allAnime)

		if anime == nil {
			fmt.Println(anime.ID, aniListAnime.TitleRomaji)
		}

		count++
	}
}
