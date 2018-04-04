package anime

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// RedirectByMapping redirects to the anime with the given mapping ID.
func RedirectByMapping(mappingName string) func(*aero.Context) string {
	return func(ctx *aero.Context) string {
		id := ctx.Get("id")
		finder := arn.NewAnimeFinder(mappingName)
		anime := finder.GetAnime(id)

		if anime == nil {
			return ctx.Error(http.StatusNotFound, "Anime not found", nil)
		}

		return utils.SmartRedirect(ctx, "/anime/"+anime.ID)
	}
}
