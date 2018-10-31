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
	allAnime := arn.FilterAnime(func(anime *arn.Anime) bool {
		return anime.GetMapping("myanimelist/anime") != ""
	})

	arn.SortAnimeByQuality(allAnime)
	color.Yellow("%d anime found", len(allAnime))

	for _, anime := range allAnime {
		syncAnime(anime, anime.GetMapping("myanimelist/anime"))
	}

	// Sync the most important ones first
	allCharacters := arn.FilterCharacters(func(character *arn.Character) bool {
		return character.GetMapping("myanimelist/character") != ""
	})

	arn.SortCharactersByLikes(allCharacters)
	color.Yellow("%d characters found", len(allCharacters))

	for _, character := range allCharacters {
		syncCharacter(character, character.GetMapping("myanimelist/character"))
	}
}

func syncAnime(anime *arn.Anime, malID string) {
	obj, err := malDB.Get("Anime", malID)

	if err != nil {
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
		return
	}

	// Skip manually created characters
	if character.CreatedBy != "" {
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
