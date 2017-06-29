package embed

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get anime list in the browser extension.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return utils.AllowEmbed(ctx, ctx.HTML(components.Login()))
	}

	animeList := user.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", nil)
	}

	animeList.Sort()

	return utils.AllowEmbed(ctx, ctx.HTML(components.AnimeList(animeList, user)))
}
