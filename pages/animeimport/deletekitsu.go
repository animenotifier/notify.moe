package animeimport

import (
	"net/http"

	"github.com/animenotifier/notify.moe/arn"

	"github.com/aerogo/aero"
)

// DeleteKitsu marks an anime for deletion.
func DeleteKitsu(ctx aero.Context) error {
	id := ctx.Get("id")

	// Is the user allowed to delete?
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	// Check that the anime really exists
	kitsuAnimeObj, err := arn.Kitsu.Get("Anime", id)

	if kitsuAnimeObj == nil {
		return ctx.Error(http.StatusNotFound, "Kitsu anime not found", err)
	}

	// Add to deleted IDs list
	deletedKitsuAnime, err := arn.GetIDList("deleted kitsu anime")

	if err != nil {
		deletedKitsuAnime = arn.IDList{}
	}

	deletedKitsuAnime = deletedKitsuAnime.Append(id)

	// Save in database
	arn.DB.Set("IDList", "deleted kitsu anime", &deletedKitsuAnime)

	return nil
}
