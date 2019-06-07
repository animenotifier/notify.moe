package main

import (
	"fmt"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
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

	// Sync anime
	if objectType == "all" || objectType == "anime" {
		allAnime := arn.FilterAnime(func(anime *arn.Anime) bool {
			return anime.GetMapping("myanimelist/anime") != ""
		})

		arn.SortAnimeByQuality(allAnime)
		color.Yellow("%d anime found", len(allAnime))

		for _, anime := range allAnime {
			syncAnime(anime, anime.GetMapping("myanimelist/anime"))
		}
	}

	// Sync characters
	if objectType == "all" || objectType == "character" {
		allCharacters := arn.FilterCharacters(func(character *arn.Character) bool {
			return character.GetMapping("myanimelist/character") != ""
		})

		arn.SortCharactersByLikes(allCharacters)
		color.Yellow("%d characters found", len(allCharacters))

		for _, character := range allCharacters {
			syncCharacter(character, character.GetMapping("myanimelist/character"))
		}
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

	malCharacter := obj.(*mal.Character)

	if character.Name.Japanese == "" && malCharacter.JapaneseName != "" {
		character.Name.Japanese = malCharacter.JapaneseName
	}

	if character.Name.Canonical == "" && malCharacter.Name != "" {
		character.Name.Canonical = malCharacter.Name
	}

	allowUpdating := character.CreatedBy == "" && character.EditedBy == ""
	description, attributes := parseCharacterDescription(malCharacter.Description)

	if strings.Contains(character.Description, "No biography written.") {
		character.Description = ""
	}

	if (allowUpdating || character.Description == "") && description != "" {
		character.Description = description
	}

	if allowUpdating {
		character.Attributes = attributes
		character.Spoilers = []arn.Spoiler{}

		for _, spoilerText := range malCharacter.Spoilers {
			if !strings.Contains(strings.TrimSpace(spoilerText), "\n\n") {
				character.Spoilers = append(character.Spoilers, arn.Spoiler{
					Text: spoilerText,
				})

				continue
			}

			paragraphs := strings.Split(spoilerText, "\n\n")

			for _, paragraph := range paragraphs {
				character.Spoilers = append(character.Spoilers, arn.Spoiler{
					Text: paragraph,
				})
			}
		}
	}

	// Save in database
	character.Save()
}
