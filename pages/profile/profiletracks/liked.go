package profiletracks

// import (
// 	"github.com/aerogo/aero"
// 	"github.com/animenotifier/notify.moe/arn"
// )

// // Liked shows all soundtracks liked by a particular user.
// func Liked(ctx aero.Context) error {
// 	return render(ctx, likedTracks)
// }

// // likedTracks returns all soundtracks that the user with the given user ID liked.
// func likedTracks(userID string) []*arn.SoundTrack {
// 	return arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
// 		return !track.IsDraft && len(track.Media) > 0 && track.LikedBy(userID)
// 	})
// }
