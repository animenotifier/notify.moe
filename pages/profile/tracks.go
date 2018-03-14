package profile

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	"github.com/animenotifier/notify.moe/utils/infinitescroll"
)

// GetSoundTracksAddedByUser shows all soundtracks added by a particular user.
func GetSoundTracksAddedByUser(ctx *aero.Context) string {
	return getSoundTracks(ctx, addedTracks)
}

// GetSoundTracksLikedByUser shows all soundtracks liked by a particular user.
func GetSoundTracksLikedByUser(ctx *aero.Context) string {
	return getSoundTracks(ctx, likedTracks)
}

// addedTracks returns all soundtracks that the user with the given userID published.
func addedTracks(userID string) []*arn.SoundTrack {
	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.CreatedBy == userID
	})
}

// likedTracks returns all soundtracks that the user with the given userID liked.
func likedTracks(userID string) []*arn.SoundTrack {
	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && track.LikedBy(userID)
	})
}

// getSoundTracks is the request handler for profile soundtrack pages.
func getSoundTracks(ctx *aero.Context, fetch func(userID string) []*arn.SoundTrack) string {
	nick := ctx.Get("nick")
	index, _ := ctx.GetInt("index")
	user := utils.GetUser(ctx)
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "User not found", err)
	}

	// Fetch all eligible tracks
	allTracks := fetch(viewUser.ID)

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
