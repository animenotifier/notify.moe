package quoteroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/quote"
	"github.com/animenotifier/notify.moe/pages/quotes"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Quotes
	page.Get(app, "/quote/:id", quote.Get)
	page.Get(app, "/quote/:id/edit", quote.Edit)
	page.Get(app, "/quote/:id/history", quote.History)
	page.Get(app, "/quotes", quotes.Latest)
	page.Get(app, "/quotes/from/:index", quotes.Latest)
	page.Get(app, "/quotes/best", quotes.Best)
	page.Get(app, "/quotes/best/from/:index", quotes.Best)
}
