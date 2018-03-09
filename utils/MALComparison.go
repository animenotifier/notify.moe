package utils

import (
	"strings"

	"github.com/animenotifier/arn"
	"github.com/animenotifier/mal"
)

type AnimeDiff interface {
	String() string
	DetailsA() string
	DetailsB() string
}

type AnimeGenresDiff struct {
	GenresA []string
	GenresB []string
}

func (diff *AnimeGenresDiff) String() string {
	return "Genres are different"
}

func (diff *AnimeGenresDiff) DetailsA() string {
	return strings.Join(diff.GenresA, ", ")
}

func (diff *AnimeGenresDiff) DetailsB() string {
	return strings.Join(diff.GenresB, ", ")
}

type MALComparison struct {
	Anime       *arn.Anime
	MALAnime    *mal.Anime
	Differences []AnimeDiff
}
