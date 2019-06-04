package main

import (
	"fmt"
	"path"
	"strconv"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/kitsu"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Syncing characters with Kitsu DB")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	kitsuCharacters := kitsu.StreamCharacters()

	for kitsuCharacter := range kitsuCharacters {
		character := &arn.Character{
			Name: arn.CharacterName{
				Canonical: kitsuCharacter.Attributes.CanonicalName,
				English:   kitsuCharacter.Attributes.Names.En,
				Japanese:  kitsuCharacter.Attributes.Names.JaJp,
				Synonyms:  kitsuCharacter.Attributes.OtherNames,
			},
			Image: arn.Image{
				Extension: path.Ext(kitsu.FixImageURL(kitsuCharacter.Attributes.Image.Original)),
			},
			Description: kitsuCharacter.Attributes.Description,
			Attributes:  []*arn.CharacterAttribute{},
		}

		character.ID = kitsuCharacter.ID
		character.Mappings = []*arn.Mapping{
			{
				Service:   "kitsu/character",
				ServiceID: kitsuCharacter.ID,
			},
		}

		if kitsuCharacter.Attributes.MalID != 0 {
			character.Mappings = append(character.Mappings, &arn.Mapping{
				Service:   "myanimelist/character",
				ServiceID: strconv.Itoa(kitsuCharacter.Attributes.MalID),
			})
		}

		// We use markdown, so replace <br/> with two line breaks.
		character.Description = strings.Replace(character.Description, "<br/>", "\n\n", -1)

		// Parse attributes like these:
		// - Position: Club Manager
		// - Height: 162 cm (5' 4")
		// - Weight: 48 kg (106 lb)
		// - Birthday: November 24
		// - Hair color: Brown
		// - Eyes: Blue (anime), Green (manga)

		lines := strings.Split(character.Description, "\n\n")
		finalLines := make([]string, 0, len(lines))

		for _, line := range lines {
			originalLine := line

			if strings.HasPrefix(line, "(") {
				line = strings.TrimPrefix(line, "(")
				line = strings.TrimSuffix(line, ")")
			}

			line = strings.TrimSpace(line)

			colonPos := strings.Index(line, ":")

			if colonPos == -1 || colonPos < 2 || colonPos > 25 {
				finalLines = append(finalLines, originalLine)
				continue
			}

			key := line[:colonPos]
			value := line[colonPos+1:]

			value = strings.TrimSpace(value)

			if key == "source" {
				key = "Source"
			}

			character.Attributes = append(character.Attributes, &arn.CharacterAttribute{
				Name:  key,
				Value: value,
			})

			fmt.Println(color.CyanString(key), color.YellowString(value))
		}

		character.Description = strings.Join(finalLines, "\n\n")
		character.Description = strings.Trim(character.Description, "\n")
		character.Save()

		// Save Kitsu character in Kitsu DB
		arn.Kitsu.Set("Character", kitsuCharacter.ID, kitsuCharacter)

		// Log
		fmt.Printf("%s %s %s\n", color.GreenString("✔"), character.ID, character.Name)
	}
}
