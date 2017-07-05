package dashboard

import (
	"sort"

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
const maxScheduleItems = 5

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
	var forumActivity []arn.Postable
	var userList interface{}
	var followingList []*arn.User
	var soundTracks []*arn.SoundTrack
	var upcomingEpisodes []*arn.UpcomingEpisode

	user := utils.GetUser(ctx)

	flow.Parallel(func() {
		forumActivity, _ = arn.GetForumActivityCached()
	}, func() {
		animeList, err := arn.GetAnimeList(user)

		if err != nil {
			return
		}

		animeList = animeList.Watching()
		animeList.PrefetchAnime()

		for _, item := range animeList.Items {
			if len(upcomingEpisodes) >= maxScheduleItems {
				break
			}

			futureEpisodes := item.Anime().UpcomingEpisodes()

			if len(futureEpisodes) == 0 {
				continue
			}

			upcomingEpisodes = append(upcomingEpisodes, futureEpisodes...)
		}

		sort.Slice(upcomingEpisodes, func(i, j int) bool {
			return upcomingEpisodes[i].Episode.AiringDate.Start < upcomingEpisodes[j].Episode.AiringDate.Start
		})
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

	return ctx.HTML(components.Dashboard(upcomingEpisodes, forumActivity, soundTracks, followingList))
}
