package utils

import "github.com/animenotifier/arn"

// AnimeWithRelatedAnime ...
type AnimeWithRelatedAnime struct {
	Anime   *arn.Anime
	Related *arn.Anime
}
