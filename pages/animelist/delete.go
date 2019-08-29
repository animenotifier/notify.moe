package animelist

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// DeleteConfirmation shows the confirmation page before deleting an anime list.
func DeleteConfirmation(ctx aero.Context) error {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	return ctx.HTML(components.DeleteAnimeList(user))
}

// Delete deletes your entire anime list.
func Delete(ctx aero.Context) error {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	animeList := user.AnimeList()
	animeList.Lock()
	animeList.Items = nil
	animeList.Unlock()

	return ctx.String("ok")
}
