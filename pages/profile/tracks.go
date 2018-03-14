package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

// GetSoundTracksByUser shows all soundtracks of a particular user.
func GetSoundTracksByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user := utils.GetUser(ctx)
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	index, _ := ctx.GetInt("index")

	// Fetch all eligible tracks
	allTracks := fetchAllTracksByUser(viewUser.ID)

	// Sort the tracks by publication date
	arn.SortSoundTracksLatestFirst(allTracks)

	// Slice the part that we need
	tracks := allTracks[index:]

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allTracks), maxTracks, index)

	// In case we're scrolling, send soundtracks only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.SoundTracksScrollable(tracks, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.TrackList(tracks, viewUser, nextIndex, user, ctx.URI()))

}

// GetSoundTracksLikedByUser shows all soundtracks liked by a particular user.
func GetSoundTracksLikedByUser(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	user := utils.GetUser(ctx)
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	index, _ := ctx.GetInt("index")

	// Fetch all eligible tracks
	allTracks := fetchAllTracksLikedByUser(viewUser.ID)

	// Sort the tracks by publication date
	arn.SortSoundTracksLatestFirst(allTracks)

	// Slice the part that we need
	tracks := allTracks[index:]

	if len(tracks) > maxTracks {
		tracks = tracks[:maxTracks]
	}

	// Next index
	nextIndex := infinitescroll.NextIndex(ctx, len(allTracks), maxTracks, index)

	// In case we're scrolling, send soundtracks only (without the page frame)
	if index > 0 {
		return ctx.HTML(components.SoundTracksScrollable(tracks, user))
	}

	// Otherwise, send the full page
	return ctx.HTML(components.TrackList(tracks, viewUser, nextIndex, user, ctx.URI()))
}

// fetchAllTracksByUser returns all soundtracks that the user with userID published
func fetchAllTracksByUser(userID string) []*arn.SoundTrack {
	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.CreatedBy == userID
	})
}

// fetchAllTracksLikedByUser returns all soundtracks that the user with userID liked
func fetchAllTracksLikedByUser(userID string) []*arn.SoundTrack {
	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.LikedBy(userID)
	})
}
