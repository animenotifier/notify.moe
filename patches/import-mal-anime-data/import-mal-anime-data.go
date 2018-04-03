package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/fatih/color"
)

var malDB = arn.Node.Namespace("mal").RegisterTypes((*mal.Anime)(nil))

func main() {
	defer arn.Node.Close()
	color.Yellow("Importing MAL data")

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		fmt.Println(anime.Title.Canonical, malID)
		sync(anime, malID)
	}

	color.Green("Finished importing MAL data")
}

func sync(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	malAnime := obj.(*mal.Anime)

	if len(anime.Genres) == 0 {
		anime.Genres = malAnime.Genres
	}

	if anime.EpisodeCount == 0 {
		anime.EpisodeCount = malAnime.EpisodeCount
	}

	if anime.EpisodeLength == 0 {
		anime.EpisodeLength = malAnime.EpisodeLength
	}

	if anime.StartDate == "" {
		anime.StartDate = malAnime.StartDate
	}

	if anime.EndDate == "" {
		anime.EndDate = malAnime.EndDate
	}

	if anime.Source == "" {
		anime.Source = malAnime.Source
	}

	if anime.Title.Japanese == "" {
		anime.Title.Japanese = malAnime.JapaneseTitle
	}

	if anime.Title.English == "" {
		anime.Title.English = malAnime.EnglishTitle
	}

	anime.Save()
}
