package tokenapi

import (
	"errors"

	"github.com/animenotifier/notify.moe/arn"
)

type AnimeParameters struct {
	AnimeName    string
	AnimeID      string
	AnimeEpisode int
	AnimeRating  arn.AnimeRatingCount
}

func AnimeUpdate(request *TokenRequest) error {
	parameters := &AnimeParameters{}
	animeJSON := request.JSON.Get("anime")

	parameters.AnimeName = animeJSON.Get("name").String()
	parameters.AnimeID = animeJSON.Get("id").String()
	parameters.AnimeEpisode = int(animeJSON.Get("episode").Int())

	if parameters.AnimeName == "" && parameters.AnimeID == "" {
		return errors.New("Neither ID nor Name of the anime has been supplied")
	}

	rating := animeJSON.Get("ratings")
	parameters.AnimeRating.Overall = int(rating.Get("overall").Int())
	parameters.AnimeRating.Story = int(rating.Get("story").Int())
	parameters.AnimeRating.Visuals = int(rating.Get("visuals").Int())
	parameters.AnimeRating.Soundtrack = int(rating.Get("soundtrack").Int())

	return nil
}
