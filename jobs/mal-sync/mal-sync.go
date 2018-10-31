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

		syncAnime(anime, malID)
	}

	// Sync the most important ones first
	allCharacters := arn.AllCharacters()
	arn.SortCharactersByLikes(allCharacters)

	for _, character := range allCharacters {
		malID := character.GetMapping("myanimelist/character")

		if malID == "" {
			continue
		}

		syncCharacter(character, malID)
	}
}

func syncAnime(anime *arn.Anime, malID string) {
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

func syncCharacter(character *arn.Character, malID string) {
	obj, err := malDB.Get("Character", malID)

	if err != nil {
		fmt.Println(err)
		return
	}

	malCharacter := obj.(*mal.Character)

	description, attributes := parseCharacterDescription(malCharacter.Description)
	character.Description = description
	character.Attributes = attributes

	if character.Name.Japanese == "" && malCharacter.JapaneseName != "" {
		character.Name.Japanese = malCharacter.JapaneseName
	}

	// Save in database
	character.Save()
}
