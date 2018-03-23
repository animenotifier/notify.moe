package utils

import "github.com/animenotifier/arn"

// MaxAnime limits the number of anime that will maximally be returned.
func MaxAnime(animes []*arn.Anime, maxLength int) []*arn.Anime {
	if len(animes) > maxLength {
		return animes[:maxLength]
	}

	return animes
}
