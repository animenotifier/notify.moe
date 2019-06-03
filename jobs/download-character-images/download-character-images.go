package main

import (
	"fmt"
	"net/http"
	"sort"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/aerogo/http/client"
	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Updating anime ratings")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	characters := arn.FilterCharacters(func(character *arn.Character) bool {
		return character.Image.LastModified == 0
	})

	sort.Slice(characters, func(i, j int) bool {
		return characters[i].ID < characters[j].ID
	})

	for index, character := range characters {
		fmt.Printf("[%d / %d] %s %s\n", index+1, len(characters), character.ID, color.CyanString(character.String()))
		download(character.ID)
	}
}

func download(characterID string) {
	character, err := arn.GetCharacter(characterID)

	if err != nil {
		color.Red(err.Error())
		return
	}

	url := fmt.Sprintf("https://media.kitsu.io/characters/images/%s/original%s", character.GetMapping("kitsu/character"), character.Image.Extension)
	response, err := client.Get(url).End()

	if err != nil {
		color.Red(err.Error())
		return
	}

	if response.StatusCode() != http.StatusOK {
		color.Red("Status: %d", response.StatusCode())
		return
	}

	err = character.SetImageBytes(response.Bytes())

	if err != nil {
		color.Red(err.Error())
		return
	}

	character.Save()
}
