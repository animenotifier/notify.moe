package anime

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// RedirectByMapping redirects to the anime with the given mapping ID.
func RedirectByMapping(mappingName string) func(aero.Context) error {
	return func(ctx aero.Context) error {
		id := ctx.Get("id")
		finder := arn.NewAnimeFinder(mappingName)
		anime := finder.GetAnime(id)

		if anime == nil {
			return ctx.Error(http.StatusNotFound, "Anime not found")
		}

		return utils.SmartRedirect(ctx, "/anime/"+anime.ID)
	}
}
