package tokenapi

import (
	"errors"

	"github.com/animenotifier/notify.moe/arn"
)

type AnimeParameters struct {
	Name    string
	ID      string
	Episode int
	Rating  arn.AnimeRatingCount
}

// AnimeUpdate parses parameters out of the JSON and then tries to integrate them into the database
func AnimeUpdate(request *TokenRequest) error {
	parameters := &AnimeParameters{}
	animeJSON := request.JSON.Get("anime")

	parameters.Name = animeJSON.Get("name").String()
	parameters.ID = animeJSON.Get("id").String()
	parameters.Episode = int(animeJSON.Get("episode").Int())

	if parameters.Name == "" && parameters.ID == "" {
		return errors.New("Neither ID nor Name of the anime has been supplied")
	}

	rating := animeJSON.Get("ratings")
	parameters.Rating.Overall = int(rating.Get("overall").Int())
	parameters.Rating.Story = int(rating.Get("story").Int())
	parameters.Rating.Visuals = int(rating.Get("visuals").Int())
	parameters.Rating.Soundtrack = int(rating.Get("soundtrack").Int())

	animeList := request.User.AnimeList()

	if !animeList.Contains(parameters.ID) {
		err := animeList.Add(parameters.ID)
		if err != nil {
			return err
		}
	}

	animeEntry := animeList.Find(parameters.ID)
	lastEpisode := animeEntry.Episodes
	// @TODO: Continue working here. See AnimeListItemAPI.go for similar implementation.

	return nil
}
