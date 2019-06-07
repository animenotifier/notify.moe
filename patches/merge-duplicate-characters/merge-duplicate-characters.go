package main

import (
	"fmt"
	"sort"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Merging duplicate characters")

	defer color.Green("Finished")
	defer arn.Node.Close()

	malIDToCharacters := map[string][]*arn.Character{}

	for character := range arn.StreamCharacters() {
		malID := character.GetMapping("myanimelist/character")

		if malID != "" {
			malIDToCharacters[malID] = append(malIDToCharacters[malID], character)
		}
	}

	for _, characters := range malIDToCharacters {
		if len(characters) > 1 {
			// nolint:scopelint
			sort.Slice(characters, func(i, j int) bool {
				return len(characters[i].Likes) > len(characters[j].Likes)
			})

			for index, character := range characters {
				if index == 0 {
					continue
				}

				fmt.Printf("Merging '%s' with '%s' (%s to %s)\n", color.YellowString(character.String()), color.YellowString(characters[0].String()), character.ID, characters[0].ID)
				character.Merge(characters[0])
			}
		}
	}
}
