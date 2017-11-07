package search

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 25
const maxAnime = 25
const maxPosts = 3
const maxThreads = 3
const maxTracks = 5

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Query("q")

	users, animes, posts, threads, tracks := arn.Search(term, maxUsers, maxAnime, maxPosts, maxThreads, maxTracks)
	return ctx.HTML(components.SearchResults(term, users, animes, posts, threads, tracks))
}
