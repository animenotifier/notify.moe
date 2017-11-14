package main

import (
	"fmt"
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/jikan"
	"github.com/fatih/color"
)

var jikanDB = arn.Node.Namespace("jikan")

func main() {
	color.Yellow("Importing jikan characters...")
	defer arn.Node.Close()

	for characterObj := range jikanDB.All("Character") {
		jikanCharacter := characterObj.(*jikan.Character)

		if jikanCharacter.Name != "Slaine Troyard" {
			continue
		}

		character := &arn.Character{
			ID:          arn.GenerateID("Character"),
			Description: jikanCharacter.About,
			Name: &arn.CharacterName{
				Romaji:   jikanCharacter.Name,
				Japanese: jikanCharacter.NameJapanese,
			},
			Image: jikanCharacter.Image,
			// Mappings: []*arn.Mapping{
			// 	&arn.Mapping{
			// 		Service: "myanimelist/character",
			// 		ServiceID: jikanCharacter.
			// 	}
			// },
		}

		if strings.HasPrefix(character.Name.Japanese, "(") {
			character.Name.Japanese = strings.TrimPrefix(character.Name.Japanese, "(")
			character.Name.Japanese = strings.TrimSuffix(character.Name.Japanese, ")")
		}

		lines := strings.Split(character.Description, "\n")
		finalLines := make([]string, 0, len(lines))

		for _, line := range lines {
			line = strings.TrimSpace(line)
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

		character.Description = strings.Join(finalLines, "\n")
		character.Description = strings.TrimSpace(character.Description)

		arn.PrettyPrint(character)
	}

	color.Green("Finished.")
}
