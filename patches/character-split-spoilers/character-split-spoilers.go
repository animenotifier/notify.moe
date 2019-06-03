package main

import (
	"strings"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Splitting character spoiler paragraphs into separate spoilers")

	defer color.Green("Finished.")
	defer arn.Node.Close()

	for character := range arn.StreamCharacters() {
		spoilers := character.Spoilers
		finalSpoilers := []arn.Spoiler{}

		for _, spoiler := range spoilers {
			if !strings.Contains(strings.TrimSpace(spoiler.Text), "\n\n") {
				finalSpoilers = append(finalSpoilers, spoiler)
				continue
			}

			paragraphs := strings.Split(spoiler.Text, "\n\n")

			for _, paragraph := range paragraphs {
				finalSpoilers = append(finalSpoilers, arn.Spoiler{
					Text: paragraph,
				})
			}
		}

		if len(finalSpoilers) != len(spoilers) {
			character.Spoilers = finalSpoilers
			character.Save()
		}
	}
}
