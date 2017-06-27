package profile

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/flow"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5
const maxTracks = 5

// Get user profile page.
func Get(ctx *aero.Context) string {
	nick := ctx.Get("nick")
	viewUser, err := arn.GetUserByNick(nick)

	if err != nil {
		return ctx.Error(404, "User not found", err)
	}

	return Profile(ctx, viewUser)
}

// Profile renders the user profile page of the given viewUser.
func Profile(ctx *aero.Context, viewUser *arn.User) string {
	var user *arn.User
	var threads []*arn.Thread
	var animeList *arn.AnimeList
	var tracks []*arn.SoundTrack
	var posts []*arn.Post

	flow.Parallel(func() {
		user = utils.GetUser(ctx)
	}, func() {
		animeList = viewUser.AnimeList()
	}, func() {
		threads = viewUser.Threads()

		arn.SortThreadsByDate(threads)

		if len(threads) > maxPosts {
			threads = threads[:maxPosts]
		}
	}, func() {
		posts = viewUser.Posts()
		arn.SortPostsLatestFirst(posts)

		if len(posts) > maxPosts {
			posts = posts[:maxPosts]
		}
	}, func() {
		tracks = viewUser.SoundTracks()
		arn.SortSoundTracksLatestFirst(tracks)

		if len(tracks) > maxTracks {
			tracks = tracks[:maxTracks]
		}
	})

	return ctx.HTML(components.Profile(viewUser, user, animeList, threads, posts, tracks))
}
