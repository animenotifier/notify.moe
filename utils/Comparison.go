package utils

import "github.com/animenotifier/notify.moe/arn"

// Comparison of an anime between 2 users.
type Comparison struct {
	Anime *arn.Anime
	ItemA *arn.AnimeListItem
	ItemB *arn.AnimeListItem
}
