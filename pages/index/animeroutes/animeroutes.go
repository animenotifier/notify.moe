package animeroutes

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/layout"
	"github.com/animenotifier/notify.moe/pages/anime"
	"github.com/animenotifier/notify.moe/pages/anime/editanime"
	"github.com/animenotifier/notify.moe/pages/episode"
	"github.com/animenotifier/notify.moe/pages/genre"
	"github.com/animenotifier/notify.moe/pages/genres"
)

// Register registers the page routes.
func Register(l *layout.Layout, app *aero.Application) {
	// Anime
	l.Page("/anime/:id", anime.Get)
	l.Page("/anime/:id/episodes", anime.Episodes)
	l.Page("/anime/:id/characters", anime.Characters)
	l.Page("/anime/:id/tracks", anime.Tracks)
	l.Page("/anime/:id/relations", anime.Relations)
	l.Page("/anime/:id/comments", anime.Comments)
	l.Page("/anime/:id/episode/:episode-number", episode.Get)
	app.Get("/anime/:id/episode/:episode-number/subtitles/:language", episode.Subtitles)

	// Anime redirects
	l.Page("/kitsu/anime/:id", anime.RedirectByMapping("kitsu/anime"))
	l.Page("/mal/anime/:id", anime.RedirectByMapping("myanimelist/anime"))
	l.Page("/anilist/anime/:id", anime.RedirectByMapping("anilist/anime"))

	// Edit anime
	l.Page("/anime/:id/edit", editanime.Main)
	l.Page("/anime/:id/edit/images", editanime.Images)
	l.Page("/anime/:id/edit/characters", editanime.Characters)
	l.Page("/anime/:id/edit/relations", editanime.Relations)
	l.Page("/anime/:id/edit/episodes", editanime.Episodes)
	l.Page("/anime/:id/edit/history", editanime.History)

	// Genres
	l.Page("/genres", genres.Get)
	l.Page("/genre/:name", genre.Get)
}
