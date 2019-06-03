package profiletracks

// import (
// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// )

// // Added shows all soundtracks added by a particular user.
// func Added(ctx aero.Context) error {
// 	return render(ctx, addedTracks)
// }

// // addedTracks returns all soundtracks that the user with the given user ID published.
// func addedTracks(userID string) []*arn.SoundTrack {
// 	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
// 		return !track.IsDraft && len(track.Media) > 0 && track.CreatedBy == userID
// 	})
// }
