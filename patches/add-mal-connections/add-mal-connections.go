package main

import (
	"strconv"

	"github.com/animenotifier/arn"
	"github.com/fatih/color"
)

func main() {
	defer arn.Node.Close()

	for anime := range arn.StreamAnime() {
		malID := anime.GetMapping("myanimelist/anime")

		if malID == "" {
			continue
		}

		// Assure the string represents a number
		malNum, _ := strconv.Atoi(malID)
		normalizedID := strconv.Itoa(malNum)

		if malID != normalizedID {
			color.Red("%s does not match %d", malID, normalizedID)
			continue
		}

		// Save
		arn.DB.Set("MyAnimeListToAnime", malID, &arn.MyAnimeListToAnime{
			AnimeID:   anime.ID,
			ServiceID: malID,
			Edited:    arn.DateTimeUTC(),
			EditedBy:  "",
		})
	}
}
