package utils

import "github.com/animenotifier/notify.moe/arn"

// CalendarDay is a calendar day.
type CalendarDay struct {
	Name    string
	Class   string
	Entries []*CalendarEntry
}

// CalendarEntry is a calendar entry.
type CalendarEntry struct {
	Anime   *arn.Anime
	Episode *arn.Episode
	Added   bool
}
