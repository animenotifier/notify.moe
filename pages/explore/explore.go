package explore

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

// Get search page.
func Get(ctx *aero.Context) string {
	animeList, err := arn.GetPopularAnimeCached()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching popular anime", err)
	}

	return ctx.HTML(components.ExploreAnime(nil, nil, animeList))
}
