package main

import (
	"strconv"
	"time"

	"github.com/akyoto/color"
	"github.com/animenotifier/notify.moe/arn"
)

func main() {
	color.Yellow("Adding creation dates")

	defer color.Green("Finished")
	defer arn.Node.Close()

	baseTime := time.Now().Add(-6 * 30 * 24 * time.Hour)
	irregular := time.Duration(0)

	for character := range arn.StreamCharacters() {
		// Skip manually created characters
		if character.CreatedBy != "" {
			continue
		}

		malID := character.GetMapping("myanimelist/character")

		if malID != "" {
			malIDNumber, err := strconv.Atoi(malID)

			if err != nil {
				panic(err)
			}

			character.Created = baseTime.Add(time.Duration(malIDNumber) * time.Minute).Format(time.RFC3339)
		} else {
			irregular++
			character.Created = baseTime.Add(-irregular * time.Minute).Format(time.RFC3339)
		}

		character.Save()
	}
}
