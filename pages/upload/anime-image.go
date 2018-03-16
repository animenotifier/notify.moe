package upload

import (
	"net/http"

	"github.com/animenotifier/arn"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// AnimeImage handles the anime image upload.
func AnimeImage(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	animeID := ctx.Get("id")

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized", nil)
	}

	anime, err := arn.GetAnime(animeID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	// Retrieve file from post body
	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	// Set anime image file
	err = anime.SetImageBytes(data)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid image format", err)
	}

	// Save image information
	anime.Save()

	return "ok"
}
