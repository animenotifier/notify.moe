package utils

import (
	"github.com/animenotifier/mal"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/utils/animediff"
)

// MALComparison encapsulates the difference between an ARN anime and a MAL anime.
type MALComparison struct {
	Anime       *arn.Anime
	MALAnime    *mal.Anime
	Differences []animediff.Difference
}
