package utils

import "github.com/animenotifier/notify.moe/arn"

// AnimeWithRelatedAnime ...
type AnimeWithRelatedAnime struct {
	Anime   *arn.Anime
	Related *arn.Anime
}
