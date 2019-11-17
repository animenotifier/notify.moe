package profiletracks

// import (
// 	"net/http"

// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// 	"github.com/animenotifier/notify.moe/components"
// 	"github.com/animenotifier/notify.moe/utils"
// 	"github.com/animenotifier/notify.moe/utils/infinitescroll"
// )

// const (
// 	tracksFirstLoad = 12
// 	tracksPerScroll = 9
// )

// // render renders the soundtracks on user profiles.
// func render(ctx aero.Context, fetch func(userID string) []*arn.SoundTrack) string {
// 	nick := ctx.Get("nick")
// 	index, _ := ctx.GetInt("index")
// 	user := arn.GetUserFromContext(ctx)
// 	viewUser, err := arn.GetUserByNick(nick)

// 	if err != nil {
// 		return ctx.Error(http.StatusNotFound, "User not found", err)
// 	}

// 	// Fetch all eligible tracks
// 	allTracks := fetch(viewUser.ID)

// 	// Sort the tracks by publication date
// 	arn.SortSoundTracksLatestFirst(allTracks)

// 	// Slice the part that we need
// 	tracks := allTracks[index:]
// 	maxLength := tracksFirstLoad

// 	if index > 0 {
// 		maxLength = tracksPerScroll
// 	}

// 	if len(tracks) > maxLength {
// 		tracks = tracks[:maxLength]
// 	}

// 	// Next index
// 	nextIndex := infinitescroll.NextIndex(ctx, len(allTracks), maxLength, index)

// 	// In case we're scrolling, send soundtracks only (without the page frame)
// 	if index > 0 {
// 		return ctx.HTML(components.SoundTracksScrollable(tracks, user))
// 	}

// 	// Otherwise, send the full page
// 	return ctx.HTML(components.ProfileSoundTracks(tracks, viewUser, nextIndex, user, ctx.Path()))
// }
