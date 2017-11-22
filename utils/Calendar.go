package utils

import "github.com/animenotifier/arn"

// CalendarDay is a calendar day.
type CalendarDay struct {
	Name    string
	Entries []*CalendarEntry
}

// CalendarEntry is a calendar entry.
type CalendarEntry struct {
	Anime   *arn.Anime
	Episode *arn.AnimeEpisode
}
