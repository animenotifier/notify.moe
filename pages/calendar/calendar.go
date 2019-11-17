package calendar

import (
	"sort"
	"time"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/arn/validate"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

var weekdayNames = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

// Get renders the calendar page.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	oneWeek := 7 * 24 * time.Hour

	now := time.Now()
	year, month, day := now.Date()
	now = time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	// Weekday index that we start with, Sunday is 0.
	weekdayIndex := int(now.Weekday())

	// Create days
	days := make([]*utils.CalendarDay, 7)

	for i := 0; i < 7; i++ {
		days[i] = &utils.CalendarDay{
			Name:    weekdayNames[(weekdayIndex+i)%7],
			Class:   "weekday",
			Entries: []*utils.CalendarEntry{},
		}

		if days[i].Name == "Saturday" || days[i].Name == "Sunday" {
			days[i].Class += " weekend"
		}
	}

	// Add anime episodes to the days
	for episode := range arn.StreamEpisodes() {
		if episode.Anime().Status == "finished" {
			continue
		}

		if !validate.DateTime(episode.AiringDate.Start) {
			continue
		}

		// Since we validated the date earlier, we can ignore the error value.
		airingDate, _ := time.Parse(time.RFC3339, episode.AiringDate.Start)

		// Subtract from the starting date offset.
		since := airingDate.Sub(now)

		// Ignore entries in the past and more than 1 week away.
		if since < 0 || since >= oneWeek {
			continue
		}

		dayIndex := int(since / (24 * time.Hour))

		entry := &utils.CalendarEntry{
			Anime:   episode.Anime(),
			Episode: episode,
			Added:   false,
		}

		if user != nil {
			animeListItem := user.AnimeList().Find(entry.Anime.ID)

			if animeListItem != nil && (animeListItem.Status == arn.AnimeListStatusWatching || animeListItem.Status == arn.AnimeListStatusPlanned) {
				entry.Added = true
			}
		}

		days[dayIndex].Entries = append(days[dayIndex].Entries, entry)
	}

	for i := 0; i < 7; i++ {
		// nolint:scopelint
		sort.Slice(days[i].Entries, func(a, b int) bool {
			airingA := days[i].Entries[a].Episode.AiringDate.Start
			airingB := days[i].Entries[b].Episode.AiringDate.Start

			if airingA == airingB {
				return days[i].Entries[a].Anime.Title.Canonical < days[i].Entries[b].Anime.Title.Canonical
			}

			return airingA < airingB
		})
	}

	return ctx.HTML(components.Calendar(days, user))
}
