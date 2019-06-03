package main

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Deleting unused characters")

	defer color.Green("Finished")
	defer arn.Node.Close()

	used := map[string]bool{}

	// Check quotes
	for quote := range arn.StreamQuotes() {
		used[quote.CharacterID] = true
	}

	// Check log
	for entry := range arn.StreamEditLogEntries() {
		if entry.ObjectType != "Character" {
			continue
		}

		used[entry.ObjectID] = true
	}

	// Check anime characters
	for list := range arn.StreamAnimeCharacters() {
		for _, animeCharacter := range list.Items {
			used[animeCharacter.CharacterID] = true
		}
	}

	characters := []*arn.Character{}

	// Delete unused characters
	for character := range arn.StreamCharacters() {
		if used[character.ID] {
			characters = append(characters, character)
		} else {
			fmt.Println("Deleting", character.ID, character)
		}
	}

	arn.DB.Clear("Character")

	for _, character := range characters {
		character.Save()
	}

	fmt.Println(len(used), len(characters))
}
