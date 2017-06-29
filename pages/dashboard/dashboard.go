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
	var posts []*arn.Post
	var userList interface{}
	var followingList []*arn.User
	var soundTracks []*arn.SoundTrack
	var upcomingEpisodes []*arn.UpcomingEpisode

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
		animeList, err := arn.GetAnimeList(user)

		if err != nil {
			return
		}

		var keys []string

		for _, item := range animeList.Items {
			keys = append(keys, item.AnimeID)
		}

		objects, getErr := arn.DB.GetMany("Anime", keys)

		if getErr != nil {
			return
		}

		allAnimeInList := objects.([]*arn.Anime)

		for _, anime := range allAnimeInList {
			if len(upcomingEpisodes) >= maxScheduleItems {
				break
			}

			futureEpisodes := anime.UpcomingEpisodes()

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

	return ctx.HTML(components.Dashboard(upcomingEpisodes, posts, soundTracks, followingList))
}
