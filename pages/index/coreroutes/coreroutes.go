package coreroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/activity"
	"github.com/animenotifier/notify.moe/pages/calendar"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/statistics"
	"github.com/animenotifier/notify.moe/pages/terms"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Front page
	l.Page("/", home.Get)

	// Login
	l.Page("/login", login.Get)

	// Activity
	l.Page("/activity", activity.Get)

	// Calendar
	l.Page("/calendar", calendar.Get)

	// Statistics
	l.Page("/statistics", statistics.Get)
	l.Page("/statistics/anime", statistics.Anime)

	// Legal stuff
	l.Page("/terms", terms.Get)

	// Browser extension
	l.Page("/extension/embed", embed.Get)
}
