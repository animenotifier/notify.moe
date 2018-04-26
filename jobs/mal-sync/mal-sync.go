package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
	"github.com/fatih/color"
)

const imageWidthThreshold = 225

var (
	malDB           = arn.Node.Namespace("mal").RegisterTypes((*mal.Anime)(nil))
	characterFinder = arn.NewCharacterFinder("myanimelist/character")
)

func main() {
	color.Yellow("Syncing with MAL DB")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	// Invoke via parameters
	if InvokeShellArgs() {
		return
	}

	// Sync the most important ones first
	allAnime := arn.AllAnime()
	arn.SortAnimeByQuality(allAnime)

	for _, anime := range allAnime {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		sync(anime, malID)
	}
}

func sync(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	malAnime := obj.(*mal.Anime)

	// Log title
	fmt.Printf("%s %s\n", color.CyanString(anime.Title.Canonical), malID)

	type syncFunction func(*arn.Anime, *mal.Anime)

	syncFunctions := []syncFunction{
		syncTitles,
		syncDates,
		syncEpisodes,
		syncOthers,
		syncImage,
		syncCharacters,
	}

	for _, syncFunction := range syncFunctions {
		syncFunction(anime, malAnime)
	}

	// Save in database
	anime.Save()
}
