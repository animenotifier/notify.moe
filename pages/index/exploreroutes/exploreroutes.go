package exploreroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/explore"
	"github.com/animenotifier/notify.moe/pages/explore/explorecolor"
	"github.com/animenotifier/notify.moe/pages/explore/halloffame"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Explore
	page.Get(app, "/explore", explore.Filter)
	page.Get(app, "/explore/anime/:year/:season/:status/:type", explore.Filter)
	page.Get(app, "/explore/color/:color/anime", explorecolor.AnimeByAverageColor)
	page.Get(app, "/explore/color/:color/anime/from/:index", explorecolor.AnimeByAverageColor)
	page.Get(app, "/halloffame", halloffame.Get)
}
