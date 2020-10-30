package tokenapi

import (
	"github.com/animenotifier/notify.moe/arn"
)

type AnimeParameters struct {
	AnimeName    string
	AnimeID      int
	AnimeEpisode int
	AnimeRating  arn.AnimeRating
}

func AnimeUpdate(request *TokenRequest, parameters *AnimeParameters) error {
	return nil
}
