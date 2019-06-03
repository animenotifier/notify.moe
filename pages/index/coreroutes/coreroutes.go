package coreroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/activity"
	"github.com/animenotifier/notify.moe/pages/calendar"
	"github.com/animenotifier/notify.moe/pages/embed"
	"github.com/animenotifier/notify.moe/pages/home"
	"github.com/animenotifier/notify.moe/pages/login"
	"github.com/animenotifier/notify.moe/pages/statistics"
	"github.com/animenotifier/notify.moe/pages/terms"
	"github.com/animenotifier/notify.moe/pages/welcome"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Front page
	page.Get(app, "/", home.Get)

	// Login
	page.Get(app, "/login", login.Get)

	// Welcome
	page.Get(app, "/welcome", welcome.Get)

	// Activity
	page.Get(app, "/activity", activity.Global)
	page.Get(app, "/activity/from/:index", activity.Global)
	page.Get(app, "/activity/followed", activity.Followed)
	page.Get(app, "/activity/followed/from/:index", activity.Followed)

	// Calendar
	page.Get(app, "/calendar", calendar.Get)

	// Statistics
	page.Get(app, "/statistics", statistics.Get)
	page.Get(app, "/statistics/anime", statistics.Anime)

	// Legal stuff
	page.Get(app, "/terms", terms.Get)

	// Browser extension
	page.Get(app, "/extension/embed", embed.Get)
}
