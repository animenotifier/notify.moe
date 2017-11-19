package soundtracks

import (
	"net/http"
	"strconv"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Best renders the soundtracks page.
func Best(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0
	})

	arn.SortSoundTracksPopularFirst(tracks)

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	return ctx.HTML(components.SoundTracks(tracks, maxTracks, user))
}

// BestFrom renders the soundtracks from the given index.
func BestFrom(ctx *aero.Context) string {
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

	arn.SortSoundTracksPopularFirst(allTracks)

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
