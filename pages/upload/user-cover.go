package upload

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// UserCover handles the cover image upload.
func UserCover(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	if !user.IsPro() {
		return ctx.Error(http.StatusUnauthorized, "Only available for PRO users")
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

	return nil
}
