package search

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxUsers = 9 * 4
const maxAnime = 9 * 4

// Get search page.
func Get(ctx *aero.Context) string {
	term := ctx.Query("q")

	userResults, animeResults := arn.Search(term, maxUsers, maxAnime)
	return ctx.HTML(components.SearchResults(term, userResults, animeResults))
}
