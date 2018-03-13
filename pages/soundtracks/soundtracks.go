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

// Latest renders the soundtracks page.
func Latest(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

	arn.SortSoundTracksLatestFirst(tracks)

	// Limit the number of displayed tracks
	loadMoreIndex := 0

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
		loadMoreIndex = maxTracks
	}

	return ctx.HTML(components.SoundTracks(tracks, loadMoreIndex, "", user))
}

// LatestFrom renders the soundtracks from the given index.
func LatestFrom(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	index, err := ctx.GetInt("index")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Invalid start index", err)
	}

	allTracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

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
