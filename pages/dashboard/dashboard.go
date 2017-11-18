package dashboard

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/aerogo/flow"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxForumActivity = 5
const maxFollowing = 5
const maxSoundTracks = 5
const maxScheduleItems = 5

// Get the dashboard.
func Get(ctx *aero.Context) string {
	var forumActivity []arn.Postable
	var followingList []*arn.User
	var soundTracks []*arn.SoundTrack
	var upcomingEpisodes []*arn.UpcomingEpisode

	user := utils.GetUser(ctx)

	if user == nil {
		return ctx.Error(http.StatusUnauthorized, "Not logged in", nil)
	}

	flow.Parallel(func() {
		posts := arn.AllPosts()
		threads := arn.AllThreads()

		arn.SortPostsLatestFirst(posts)
		arn.SortThreadsLatestFirst(threads)

		posts = arn.FilterPostsWithUniqueThreads(posts, maxForumActivity)

		postPostables := arn.ToPostables(posts)
		threadPostables := arn.ToPostables(threads)

		allPostables := append(postPostables, threadPostables...)

		arn.SortPostablesLatestFirst(allPostables)
		forumActivity = arn.FilterPostablesWithUniqueThreads(allPostables, maxForumActivity)
	}, func() {
		animeList, err := arn.GetAnimeList(user.ID)

		if err != nil {
			return
		}

		animeList = animeList.Watching()

		for _, item := range animeList.Items {
			futureEpisodes := item.Anime().UpcomingEpisodes()

			if len(futureEpisodes) == 0 {
				continue
			}

			upcomingEpisodes = append(upcomingEpisodes, futureEpisodes...)
		}

		sort.Slice(upcomingEpisodes, func(i, j int) bool {
			return upcomingEpisodes[i].Episode.AiringDate.Start < upcomingEpisodes[j].Episode.AiringDate.Start
		})

		if len(upcomingEpisodes) >= maxScheduleItems {
			upcomingEpisodes = upcomingEpisodes[:maxScheduleItems]
		}
	}, func() {
		soundTracks = arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
			return !track.IsDraft && len(track.Media) > 0
		})

		arn.SortSoundTracksLatestFirst(soundTracks)

		if len(soundTracks) > maxSoundTracks {
			soundTracks = soundTracks[:maxSoundTracks]
		}
	}, func() {
		followingList = user.Follows().Users()
		arn.SortUsersLastSeen(followingList)

		if len(followingList) > maxFollowing {
			followingList = followingList[:maxFollowing]
		}
	})

	return ctx.HTML(components.Dashboard(upcomingEpisodes, forumActivity, soundTracks, followingList, user))
}
