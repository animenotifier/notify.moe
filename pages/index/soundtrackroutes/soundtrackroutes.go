package soundtrackroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/soundtrack"
	"github.com/animenotifier/notify.moe/pages/soundtracks"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Soundtracks
	l.Page("/soundtracks", soundtracks.Latest)
	l.Page("/soundtracks/from/:index", soundtracks.Latest)
	l.Page("/soundtracks/best", soundtracks.Best)
	l.Page("/soundtracks/best/from/:index", soundtracks.Best)
	l.Page("/soundtracks/tag/:tag", soundtracks.FilterByTag)
	l.Page("/soundtracks/tag/:tag/from/:index", soundtracks.FilterByTag)
	l.Page("/soundtrack/:id", soundtrack.Get)
	l.Page("/soundtrack/:id/lyrics", soundtrack.Lyrics)
	l.Page("/soundtrack/:id/edit", soundtrack.Edit)
	l.Page("/soundtrack/:id/history", soundtrack.History)
}
