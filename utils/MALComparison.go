package utils

import (
	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
)

// AnimeDiff describes a difference between two anime.
type AnimeDiff interface {
	String() string
	DetailsA() string
	DetailsB() string
}

// MALComparison encapsulates the difference between an ARN anime and a MAL anime.
type MALComparison struct {
	Anime       *arn.Anime
	MALAnime    *mal.Anime
	Differences []AnimeDiff
}
