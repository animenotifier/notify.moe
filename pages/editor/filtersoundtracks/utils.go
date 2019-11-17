package filtersoundtracks

import (
	"net/http"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
)

const maxSoundTrackEntries = 70

// editorList renders the soundtrack list with the given title and filter.
func editorList(ctx aero.Context, title string, filter func(*arn.SoundTrack) bool, searchLink func(*arn.SoundTrack) string) error {
	user := arn.GetUserFromContext(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Error(http.StatusUnauthorized, "Not authorized")
	}

	tracks, count := filterSoundTracks(filter)
	url := strings.TrimPrefix(ctx.Path(), "/_")

	return ctx.HTML(components.SoundTracksEditorListFull(
		title,
		tracks,
		count,
		url,
		searchLink,
		user,
	))
}

// filterSoundTracks filters soundtracks by the given filter function.
func filterSoundTracks(filter func(*arn.SoundTrack) bool) ([]*arn.SoundTrack, int) {
	// Filter
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && filter(track)
	})

	// Sort
	arn.SortSoundTracksPopularFirst(tracks)

	// Limit
	count := len(tracks)

	if count > maxSoundTrackEntries {
		tracks = tracks[:maxSoundTrackEntries]
	}

	return tracks, count
}
