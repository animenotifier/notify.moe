package search

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 36
const maxAnime = 26
const maxPosts = 3
const maxThreads = 3

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Query("q")

	userResults, animeResults, postResults, threadResults := arn.Search(term, maxUsers, maxAnime, maxPosts, maxThreads)
	return ctx.HTML(components.SearchResults(term, userResults, animeResults, postResults, threadResults))
}
