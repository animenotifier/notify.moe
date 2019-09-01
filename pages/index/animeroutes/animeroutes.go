package animeroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/anime/editanime"
	"github.com/animenotifier/notify.moe/pages/calendar"
	"github.com/animenotifier/notify.moe/pages/episode"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
	"github.com/animenotifier/notify.moe/pages/statistics"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Anime
	page.Get(app, "/anime/:id", anime.Get)
	page.Get(app, "/anime/:id/episodes", anime.Episodes)
	page.Get(app, "/anime/:id/characters", anime.Characters)
	page.Get(app, "/anime/:id/tracks", anime.Tracks)
	page.Get(app, "/anime/:id/relations", anime.Relations)
	page.Get(app, "/anime/:id/comments", anime.Comments)
	page.Get(app, "/episode/:id", episode.Get)
	app.Get("/episode/:id/subtitles/:language", episode.Subtitles)

	// Anime redirects
	page.Get(app, "/kitsu/anime/:id", anime.RedirectByMapping("kitsu/anime"))
	page.Get(app, "/mal/anime/:id", anime.RedirectByMapping("myanimelist/anime"))
	page.Get(app, "/anilist/anime/:id", anime.RedirectByMapping("anilist/anime"))

	// Edit anime
	page.Get(app, "/anime/:id/edit", editanime.Main)
	page.Get(app, "/anime/:id/edit/images", editanime.Images)
	page.Get(app, "/anime/:id/edit/characters", editanime.Characters)
	page.Get(app, "/anime/:id/edit/relations", editanime.Relations)
	page.Get(app, "/anime/:id/edit/episodes", editanime.Episodes)
	page.Get(app, "/anime/:id/edit/history", editanime.History)

	// Genres
	page.Get(app, "/genres", genres.Get)
	page.Get(app, "/genre/:name", genre.Get)

	// Calendar
	page.Get(app, "/calendar", calendar.Get)

	// Statistics
	page.Get(app, "/statistics", statistics.Get)
	page.Get(app, "/statistics/anime", statistics.Anime)
}
