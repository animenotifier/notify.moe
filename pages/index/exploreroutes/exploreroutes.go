package exploreroutes

import (
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/explore"
	"github.com/animenotifier/notify.moe/pages/explore/explorecolor"
	"github.com/animenotifier/notify.moe/pages/explore/halloffame"
)

// Register registers the page routes.
func Register(l *layout.Layout) {
	// Explore
	l.Page("/explore", explore.Filter)
	l.Page("/explore/anime/:year/:season/:status/:type", explore.Filter)
	l.Page("/explore/color/:color/anime", explorecolor.AnimeByAverageColor)
	l.Page("/explore/color/:color/anime/from/:index", explorecolor.AnimeByAverageColor)
	l.Page("/halloffame", halloffame.Get)
}
