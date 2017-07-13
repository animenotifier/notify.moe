package main

import (
	"fmt"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing characters with Kitsu DB")

	kitsuCharacters := kitsu.StreamCharacters()

	for kitsuCharacter := range kitsuCharacters {
		character := &arn.Character{
			ID:          kitsuCharacter.ID,
			Name:        kitsuCharacter.Attributes.Name,
			Image:       kitsu.FixImageURL(kitsuCharacter.Attributes.Image.Original),
			Description: kitsuCharacter.Attributes.Description,
		}

		fmt.Printf("%s %s\n", character.ID, character.Name)

		arn.PanicOnError(character.Save())
	}

	color.Green("Finished.")
}
