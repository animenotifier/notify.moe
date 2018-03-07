package upload

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Cover handles the cover image upload.
func Cover(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	// Retrieve file from post body
	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	// Set cover image file
	err = user.SetCoverBytes(data)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid image format", err)
	}

	// Save cover image information
	user.Save()

	return "ok"
}
