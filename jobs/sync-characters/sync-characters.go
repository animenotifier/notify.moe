package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/kitsu"
	"github.com/fatih/color"
)

func main() {
	color.Yellow("Syncing characters with Kitsu DB")
	defer arn.Node.Close()

	kitsuCharacters := kitsu.StreamCharacters()

	for kitsuCharacter := range kitsuCharacters {
		character := &arn.Character{
			ID:          kitsuCharacter.ID,
			Name:        kitsuCharacter.Attributes.Name,
			Image:       kitsu.FixImageURL(kitsuCharacter.Attributes.Image.Original),
			Description: kitsuCharacter.Attributes.Description,
		}

		// We use markdown, so replace <br/> with two line breaks.
		character.Description = strings.Replace(character.Description, "<br/>", "\n\n", -1)
		character.Save()

		fmt.Printf("%s %s %s\n", color.GreenString("✔"), character.ID, character.Name)
	}

	color.Green("Finished.")
}
