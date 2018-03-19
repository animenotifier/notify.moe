package profilequotes

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Liked shows all quotes liked by a particular user.
func Liked(ctx *aero.Context) string {
	return render(ctx, likedQuotes)
}

// likedQuotes returns all quotes that the user with the given user ID liked.
func likedQuotes(userID string) []*arn.Quote {
	return arn.FilterQuotes(func(track *arn.Quote) bool {
		return !track.IsDraft && len(track.Text.English) > 0 && track.LikedBy(userID)
	})
}
