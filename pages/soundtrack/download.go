package soundtrack

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils"
)

// Download tries to refresh the soundtrack file.
func Download(ctx aero.Context) error {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "editor" && user.Role != "admin") {
		return ctx.Error(http.StatusUnauthorized, "Not logged in or not auhorized to edit this soundtrack")
	}

	track, err := arn.GetSoundTrack(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	track.Download()
	return nil
}
