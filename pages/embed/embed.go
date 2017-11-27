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
		return utils.AllowEmbed(ctx, ctx.HTML(components.Login("_blank")))
	}

	// Extension is enabled as long as the site isn't finished yet.
	// ---
	// if !user.IsPro() && user.TimeSinceRegistered() > 14*24*time.Hour {
	// 	return utils.AllowEmbed(ctx, ctx.HTML(components.EmbedProNotice(user)))
	// }

	animeList := user.AnimeList()

	if animeList == nil {
		return ctx.Error(http.StatusNotFound, "Anime list not found", nil)
	}

	watchingList := animeList.Watching()
	watchingList.Sort()

	return utils.AllowEmbed(ctx, ctx.HTML(components.AnimeList(watchingList, animeList.User(), user)))
}
