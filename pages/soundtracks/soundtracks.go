package soundtracks

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxTracks = 12

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

	return ctx.HTML(components.SoundTracks(tracks, maxTracks, user))
}

// From renders the soundtracks from the given index.
func From(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	index, err := ctx.GetInt("index")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid start index", err)
	}

	allTracks, err := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, "Error fetching soundtracks", err)
	}

	if index < 0 || index >= len(allTracks) {
		return ctx.Error(http.StatusBadRequest, "Invalid start index (maximum is "+strconv.Itoa(len(allTracks))+")", nil)
	}

	arn.SortSoundTracksLatestFirst(allTracks)

	tracks := allTracks[index:]

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	nextIndex := index + maxTracks

	if nextIndex >= len(allTracks) {
		// End of data - no more scrolling
		ctx.Response().Header().Set("X-LoadMore-Index", "-1")
	} else {
		// Send the index for the next request
		ctx.Response().Header().Set("X-LoadMore-Index", strconv.Itoa(nextIndex))
	}

	return ctx.HTML(components.SoundTracksScrollable(tracks, user))
}
