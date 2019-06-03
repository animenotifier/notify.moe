package main

import (
	"fmt"
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
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

				name, value := parseAttribute(line)

				if name != "" && value != "" {
					lastAttribute = &arn.CharacterAttribute{
						Name:  name,
						Value: value,
					}

					attributes = append(attributes, lastAttribute)
				}
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
		if len(paragraph) < 30 && !strings.HasSuffix(paragraph, ".") || strings.HasSuffix(paragraph, "...") {
			continue
		}

		// Is it an attribute?
		name, value := parseAttribute(paragraph)

		if name != "" && value != "" {
			attributes = append(attributes, &arn.CharacterAttribute{
				Name:  name,
				Value: value,
			})
			continue
		}

		finalParagraphs = append(finalParagraphs, paragraph)
	}

	output = strings.Join(finalParagraphs, "\n\n")
	output = strings.TrimSpace(output)

	return output, attributes
}

func parseAttribute(line string) (string, string) {
	if !strings.Contains(line, ":") {
		return "", ""
	}

	parts := strings.Split(line, ":")
	name := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	// Remove list indicators
	name = strings.TrimPrefix(name, "- ")
	name = strings.TrimPrefix(name, "* ")

	if strings.HasPrefix(name, "~") && strings.HasSuffix(value, "~") {
		name = strings.TrimPrefix(name, "~")
		value = strings.TrimSuffix(value, "~")
	}

	if strings.HasPrefix(name, "[") && strings.HasSuffix(value, "]") {
		name = strings.TrimPrefix(name, "[")
		value = strings.TrimSuffix(value, "]")
	}

	if strings.HasPrefix(name, "(") && strings.HasSuffix(value, ")") {
		name = strings.TrimPrefix(name, "(")
		value = strings.TrimSuffix(value, ")")
	}

	if name == "source" || name == "sources" {
		name = "Source"
	}

	if len(name) > 25 || len(value) > 50 || strings.HasSuffix(value, ".") {
		return "", ""
	}

	fmt.Println(color.GreenString(name), color.YellowString(value))
	return name, value
}
