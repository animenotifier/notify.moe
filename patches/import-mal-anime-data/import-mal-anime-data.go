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

		fmt.Printf("%s %s\n", color.CyanString(anime.Title.Canonical), malID)
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

	if len(anime.Genres) == 0 && len(malAnime.Genres) > 0 {
		fmt.Println("Genres:", malAnime.Genres)
		anime.Genres = malAnime.Genres
	}

	if anime.EpisodeCount == 0 && malAnime.EpisodeCount != 0 {
		fmt.Println("EpisodeCount:", malAnime.EpisodeCount)
		anime.EpisodeCount = malAnime.EpisodeCount
	}

	if anime.EpisodeLength == 0 && malAnime.EpisodeLength != 0 {
		fmt.Println("EpisodeLength:", malAnime.EpisodeLength)
		anime.EpisodeLength = malAnime.EpisodeLength
	}

	if anime.StartDate == "" && malAnime.StartDate != "" {
		fmt.Println("StartDate:", malAnime.StartDate)
		anime.StartDate = malAnime.StartDate
	}

	if anime.EndDate == "" && malAnime.EndDate != "" {
		fmt.Println("EndDate:", malAnime.EndDate)
		anime.EndDate = malAnime.EndDate
	}

	if anime.Source == "" && malAnime.Source != "" {
		fmt.Println("Source:", malAnime.Source)
		anime.Source = malAnime.Source
	}

	if anime.Title.Japanese == "" && malAnime.JapaneseTitle != "" {
		fmt.Println("JapaneseTitle:", malAnime.JapaneseTitle)
		anime.Title.Japanese = malAnime.JapaneseTitle
	}

	if anime.Title.English == "" && malAnime.EnglishTitle != "" {
		fmt.Println("EnglishTitle:", malAnime.EnglishTitle)
		anime.Title.English = malAnime.EnglishTitle
	}

	anime.Save()
}
