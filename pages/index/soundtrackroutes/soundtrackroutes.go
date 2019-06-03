package soundtrackroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/soundtrack"
	"github.com/animenotifier/notify.moe/pages/soundtracks"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Soundtracks
	page.Get(app, "/soundtracks", soundtracks.Latest)
	page.Get(app, "/soundtracks/from/:index", soundtracks.Latest)
	page.Get(app, "/soundtracks/best", soundtracks.Best)
	page.Get(app, "/soundtracks/best/from/:index", soundtracks.Best)
	page.Get(app, "/soundtracks/tag/:tag", soundtracks.FilterByTag)
	page.Get(app, "/soundtracks/tag/:tag/from/:index", soundtracks.FilterByTag)
	page.Get(app, "/soundtrack/:id", soundtrack.Get)
	page.Get(app, "/soundtrack/:id/lyrics", soundtrack.Lyrics)
	page.Get(app, "/soundtrack/:id/edit", soundtrack.Edit)
	page.Get(app, "/soundtrack/:id/history", soundtrack.History)
}
