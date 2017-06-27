package dashboard

import (
	"github.com/aerogo/aero"
	"github.com/aerogo/flow"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/pages/frontpage"
	"github.com/animenotifier/notify.moe/utils"
)

const maxPosts = 5
const maxFollowing = 5
const maxSoundTracks = 5

// Get the dashboard or the frontpage when logged out.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil {
		return frontpage.Get(ctx)
	}

	return dashboard(ctx)
}

// Render the dashboard.
func dashboard(ctx *aero.Context) string {
	var posts []*arn.Post
	var userList interface{}
	var followingList []*arn.User
	var soundTracks []*arn.SoundTrack

	user := utils.GetUser(ctx)

	flow.Parallel(func() {
		var err error
		posts, err = arn.AllPosts()

		if err != nil {
			return
		}

		arn.SortPostsLatestFirst(posts)
		posts = arn.FilterPostsWithUniqueThreads(posts, maxPosts)
	}, func() {
		var err error
		soundTracks, err = arn.AllSoundTracks()

		if err != nil {
			return
		}

		arn.SortSoundTracksLatestFirst(soundTracks)

		if len(soundTracks) > maxSoundTracks {
			soundTracks = soundTracks[:maxSoundTracks]
		}
	}, func() {
		var err error
		userList, err = arn.DB.GetMany("User", user.Following)

		if err != nil {
			return
		}

		followingList = userList.([]*arn.User)
		followingList = arn.SortUsersLastSeen(followingList)

		if len(followingList) > maxFollowing {
			followingList = followingList[:maxFollowing]
		}
	})

	return ctx.HTML(components.Dashboard(posts, soundTracks, followingList))
}
