package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/animenotifier/arn"
)

// OldMatch ...
type OldMatch struct {
	ID           int     `json:"id"`
	ServiceID    int     `json:"providerId"`
	Title        string  `json:"title"`
	ServiceTitle string  `json:"providerTitle"`
	Similarity   float64 `json:"similarity"`
	Edited       string  `json:"edited"`
	EditedBy     string  `json:"editedBy"`
}

func main() {
	matches := []OldMatch{}
	data, _ := ioutil.ReadFile("MatchKitsu.json")
	json.Unmarshal(data, &matches)

	for _, match := range matches {
		// Custom anime in 3.0
		if match.ID >= 1000000 {
			continue
		}

		// New match type
		newMatch := &arn.AniListToAnime{
			AnimeID:    strconv.Itoa(match.ServiceID),
			ServiceID:  strconv.Itoa(match.ID),
			Similarity: match.Similarity,
			Edited:     match.Edited,
			EditedBy:   match.EditedBy,
		}

		// Get anime
		anime, err := arn.GetAnime(newMatch.AnimeID)

		if err != nil {
			continue
		}

		if anime.GetMapping("anilist/anime") != "" {
			continue
		}

		anime.Mappings = append(anime.Mappings, &arn.Mapping{
			Service:   "anilist/anime",
			ServiceID: newMatch.ServiceID,
			Created:   newMatch.Edited,
			CreatedBy: newMatch.EditedBy,
		})

		// Save
		fmt.Println(anime.Title.Canonical)
		arn.PanicOnError(anime.Save())
		arn.PanicOnError(arn.DB.Set("AniListToAnime", newMatch.ServiceID, newMatch))
	}
}

// AnilistToAnime
/*
AnimeID
ServiceID

*/
