package soundtracks

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxTracks = 9

// Get renders the soundtracks page.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	tracks, err := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	arn.SortSoundTracksLatestFirst(tracks)

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.SoundTracks(tracks, user))
}

// From renders the soundtracks from the given index.
func From(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	index, err := ctx.GetInt("index")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid start index", err)
	}

	tracks, err := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	if index < 0 || index >= len(tracks) {
		return ctx.Error(http.StatusBadRequest, "Invalid start index (maximum is "+strconv.Itoa(len(tracks))+")", nil)
	}

	arn.SortSoundTracksLatestFirst(tracks)

	tracks = tracks[index:]

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.SoundTracksScrollable(tracks, user))
}
