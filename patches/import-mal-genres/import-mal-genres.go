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
	color.Yellow("Importing genres")

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		fmt.Println(anime.Title.Canonical, malID)
		sync(anime, malID)
	}

	color.Green("Finished importing genres")
}

func sync(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	malAnime := obj.(*mal.Anime)
	anime.Genres = malAnime.Genres
	anime.Save()
}
