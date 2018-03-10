package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	count := 0

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		malAnimeObj, err := arn.MAL.Get("Anime", malID)

		if err != nil {
			continue
		}

		malAnime := malAnimeObj.(*mal.Anime)

		fmt.Println(anime.Source)

		anime.Source = strings.ToLower(malAnime.Source)

		if anime.Source == "Unknown" {
			anime.Source = ""
		}

		anime.Save()

		count++

		fmt.Println(anime.ID, anime, anime.Source)
	}

	color.Green("Processed %d anime", count)
}
