package utils

import "github.com/animenotifier/arn"

// HallOfFameEntry is an entry in the hall of fame.
type HallOfFameEntry struct {
	Year  int
	Anime *arn.Anime
}
