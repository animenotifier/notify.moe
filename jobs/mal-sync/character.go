package main

import (
	"strings"

	"github.com/animenotifier/arn"
)

func parseCharacterDescription(input string) (output string, attributes []*arn.CharacterAttribute) {
	// Parse attributes like these:
	// - Position: Club Manager
	// - Height: 162 cm (5' 4")
	// - Weight: 48 kg (106 lb)
	// - Birthday: November 24
	// - Hair color: Brown
	// - Eyes: Blue (anime), Green (manga)

	paragraphs := strings.Split(input, "\n\n")
	finalParagraphs := make([]string, 0, len(paragraphs))

	for _, paragraph := range paragraphs {
		// Is paragraph full of attributes?
		if strings.Contains(paragraph, "\n") {
			lines := strings.Split(paragraph, "\n")
			var lastAttribute *arn.CharacterAttribute

			for _, line := range lines {
				// line = strings.Replace(line, " (\n)", "", -1)

				// Remove all kinds of starting and ending parantheses.
				if strings.HasPrefix(line, "(") {
					line = strings.TrimPrefix(line, "(")
					line = strings.TrimSuffix(line, ")")
				}

				if !strings.Contains(line, ":") {
					// Remove list indicators
					line = strings.TrimPrefix(line, "- ")
					line = strings.TrimPrefix(line, "* ")

					// Add to previous attribute
					if lastAttribute != nil {
						if lastAttribute.Value != "" {
							lastAttribute.Value += ", "
						}

						lastAttribute.Value += line
					}

					continue
				}

				parts := strings.Split(line, ":")
				name := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				lastAttribute = &arn.CharacterAttribute{
					Name:  name,
					Value: value,
				}

				attributes = append(attributes, lastAttribute)
			}

			continue
		}

		// Remove all kinds of starting and ending parantheses.
		if strings.HasPrefix(paragraph, "(") {
			paragraph = strings.TrimPrefix(paragraph, "(")
			paragraph = strings.TrimSuffix(paragraph, ")")
		}

		// Replace source paragraph with an attribute
		if strings.HasPrefix(paragraph, "Source:") || strings.HasPrefix(paragraph, "source:") {
			source := paragraph[len("source:"):]
			source = strings.TrimSpace(source)

			attributes = append(attributes, &arn.CharacterAttribute{
				Name:  "Source",
				Value: source,
			})
			continue
		}

		paragraph = strings.TrimSpace(paragraph)

		// Skip paragraph if it's too short.
		if len(paragraph) < 30 {
			if !strings.HasSuffix(paragraph, ".") || strings.HasSuffix(paragraph, "...") {
				continue
			}
		}

		finalParagraphs = append(finalParagraphs, paragraph)

		// originalLine := line

		// line = strings.TrimSpace(line)

		// colonPos := strings.Index(line, ":")

		// // If a colon has not been found or the colon is too far,
		// // treat it as a normal line.
		// if colonPos == -1 || colonPos < 2 || colonPos > 25 {
		// 	finalLines = append(finalLines, originalLine)
		// 	continue
		// }

		// key := line[:colonPos]
		// value := line[colonPos+1:]

		// value = strings.TrimSpace(value)

		// if key == "source" {
		// 	key = "Source"
		// }

		// attributes = append(attributes, &arn.CharacterAttribute{
		// 	Name:  key,
		// 	Value: value,
		// })

		// fmt.Println(color.CyanString(key), color.YellowString(value))
	}

	output = strings.Join(finalParagraphs, "\n\n")
	output = strings.TrimSpace(output)

	return output, attributes
}
