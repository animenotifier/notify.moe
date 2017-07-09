package search

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 6 * 6
const maxAnime = 5 * 6

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Query("q")

	userResults, animeResults := arn.Search(term, maxUsers, maxAnime)
	return ctx.HTML(components.SearchResults(term, userResults, animeResults))
}
