package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// GetSoundTracksByUser shows all soundtracks of a particular user.
func GetSoundTracksByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user := utils.GetUser(ctx)
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	tracks, err := arn.GetSoundTracksByUser(viewUser)

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	arn.SortSoundTracksLatestFirst(tracks)

	return ctx.HTML(components.TrackList(tracks, viewUser, user))

}
