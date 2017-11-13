package main

import (
	"fmt"
	"time"

	"github.com/animenotifier/jikan"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

var jikanDB = arn.Node.Namespace("jikan")

func main() {
	color.Yellow("Syncing characters with Jikan API")
	defer arn.Node.Close()

	allAnime := jikanDB.All("Anime")

	count := 0

	for animeObj := range allAnime {
		anime := animeObj.(*jikan.Anime)

		if len(anime.Character) == 0 {
			continue
		}

		fmt.Println(anime.Title)

		for _, character := range anime.Character {
			characterID := jikan.GetCharacterIDFromURL(character.URL)

			if characterID == "" {
				fmt.Println("Invalid character ID")
				continue
			}

			fetchCharacter(characterID)
		}
	}

	color.Green("Finished syncing %d characters.", count)
}

func fetchCharacter(malCharacterID string) {
	fmt.Printf("Fetching character ID %s\n", malCharacterID)

	if jikanDB.Exists("Character", malCharacterID) {
		return
	}

	time.Sleep(time.Second)
	character, err := jikan.GetCharacter(malCharacterID)

	if err == nil {
		jikanDB.Set("Character", malCharacterID, character)
		return
	}

	fmt.Printf("Error fetching %s: %v", malCharacterID, err)
}
