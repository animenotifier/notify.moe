package search

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 25
const maxAnime = 25
const maxPosts = 2
const maxThreads = 2
const maxTracks = 4
const maxCharacters = 22

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Get("term")
	term = strings.TrimPrefix(term, "/")

	users, animes, posts, threads, tracks, characters := arn.Search(term, maxUsers, maxAnime, maxPosts, maxThreads, maxTracks, maxCharacters)
	return ctx.HTML(components.SearchResults(term, users, animes, posts, threads, tracks, characters))
}
