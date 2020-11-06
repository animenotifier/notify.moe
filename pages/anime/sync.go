package anime

import (
	"net/http"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
)

// SyncEpisodes syncs the episodes with an external site.
func SyncEpisodes(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	animeID := ctx.Get("id")

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	if user.Role != "editor" && user.Role != "admin" {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	anime, err := arn.GetAnime(animeID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	return anime.RefreshEpisodes()
}
