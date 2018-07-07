package upload

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/utils"
)

// Avatar handles the avatar upload.
func Avatar(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in")
	}

	// Retrieve file from post body
	data, err := ctx.Request().Body().Bytes()

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Reading request body failed", err)
	}

	// Set avatar file
	err = user.SetAvatarBytes(data)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Invalid image format", err)
	}

	// Save avatar information
	user.Save()

	return user.AvatarLink("small")
}
