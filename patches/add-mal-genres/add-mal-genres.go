package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	count := 0

	for anime := range arn.StreamAnime() {
		if len(anime.Genres) > 0 {
			continue
		}

		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		malAnimeObj, err := arn.MAL.Get("Anime", malID)

		if err != nil {
			continue
		}

		malAnime := malAnimeObj.(*mal.Anime)

		if len(malAnime.Genres) == 0 {
			continue
		}

		anime.Genres = malAnime.Genres
		anime.Save()

		count++

		fmt.Println(anime.ID, anime, anime.Genres)
	}

	color.Green("Added genres to %d anime", count)
}
