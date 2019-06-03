package searchroutes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/pages/search"
	"github.com/animenotifier/notify.moe/pages/search/multisearch"
	"github.com/animenotifier/notify.moe/utils/page"
)

// Register registers the page routes.
func Register(app *aero.Application) {
	// Search
	page.Get(app, "/search/*term", search.Get)
	page.Get(app, "/empty-search", search.GetEmptySearch)
	page.Get(app, "/anime-search/*term", search.Anime)
	page.Get(app, "/character-search/*term", search.Characters)
	page.Get(app, "/posts-search/*term", search.Posts)
	page.Get(app, "/threads-search/*term", search.Threads)
	page.Get(app, "/soundtrack-search/*term", search.SoundTracks)
	page.Get(app, "/user-search/*term", search.Users)
	page.Get(app, "/amv-search/*term", search.AMVs)
	page.Get(app, "/company-search/*term", search.Companies)

	// Multi-search
	page.Get(app, "/multisearch/anime", multisearch.Anime)
}
