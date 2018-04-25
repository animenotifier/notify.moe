package quoteroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/quote"
	"github.com/animenotifier/notify.moe/pages/quotes"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Quotes
	l.Page("/quote/:id", quote.Get)
	l.Page("/quote/:id/edit", quote.Edit)
	l.Page("/quote/:id/history", quote.History)
	l.Page("/quotes", quotes.Latest)
	l.Page("/quotes/from/:index", quotes.Latest)
	l.Page("/quotes/best", quotes.Best)
	l.Page("/quotes/best/from/:index", quotes.Best)
}
