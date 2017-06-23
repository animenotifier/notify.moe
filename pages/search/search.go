package search

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 9 * 7
const maxAnime = 9 * 7

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Get("term")

	userResults, animeResults := arn.Search(term, maxUsers, maxAnime)
	return ctx.HTML(components.SearchResults(userResults, animeResults))
}
