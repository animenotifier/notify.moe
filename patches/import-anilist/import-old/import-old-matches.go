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
	ID            int     `json:"id"`
	ProviderID    int     `json:"providerId"`
	Title         string  `json:"title"`
	ProviderTitle string  `json:"providerTitle"`
	Similarity    float64 `json:"similarity"`
	Edited        string  `json:"edited"`
	EditedBy      string  `json:"editedBy"`
}

// ProviderMatch ...
type ProviderMatch struct {
	AnimeID    string `json:"animeId"`
	ProviderID string `json:"providerId"`
	Edited     string `json:"edited"`
	EditedBy   string `json:"editedBy"`
}

// AniListToAnime ...
type AniListToAnime ProviderMatch

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
		newMatch := &ProviderMatch{
			AnimeID:    strconv.Itoa(match.ProviderID),
			ProviderID: strconv.Itoa(match.ID),
			Edited:     match.Edited,
			EditedBy:   match.EditedBy,
		}

		// Get anime
		anime, err := arn.GetAnime(newMatch.AnimeID)

		if err != nil {
			continue
		}

		anime.Mappings = append(anime.Mappings, &arn.Mapping{
			Service:   "anilist/anime",
			ServiceID: newMatch.ProviderID,
			Created:   newMatch.Edited,
			CreatedBy: newMatch.EditedBy,
		})

		// Save
		fmt.Println(anime.Title.Canonical)
		arn.PanicOnError(anime.Save())
		arn.PanicOnError(arn.DB.Set("AniListToAnime", newMatch.ProviderID, newMatch))
	}
}

// AnilistToAnime
/*
AnimeID
ProviderID

*/
