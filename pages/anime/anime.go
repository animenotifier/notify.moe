package anime

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/assets"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

const (
	maxEpisodes           = 26
	maxEpisodesLongSeries = 12
	maxDescriptionLength  = 170
	maxFriendsPerEpisode  = 9
)

// Get anime page.
func Get(ctx aero.Context) error {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	// Anime list item
	var animeListItem *arn.AnimeListItem

	if user != nil {
		animeListItem = user.AnimeList().Find(anime.ID)
	}

	// Episodes
	episodes := anime.Episodes()

	if len(episodes) > maxEpisodes {
		episodes = episodes[len(episodes)-maxEpisodesLongSeries:]
	}

	// Friends watching
	var friends []*arn.User
	friendsAnimeListItems := map[*arn.User]*arn.AnimeListItem{}
	episodeToFriends := map[int][]*arn.User{}

	if user != nil {
		friends = user.Follows().Users()
		deleted := 0

		if animeListItem != nil {
			episodeToFriends[animeListItem.Episodes] = append(episodeToFriends[animeListItem.Episodes], user)
		}

		for i := range friends {
			j := i - deleted
			friend := friends[j]
			friendAnimeList := friend.AnimeList()
			friendAnimeListItem := friendAnimeList.Find(anime.ID)

			if friendAnimeListItem == nil || friendAnimeListItem.Private {
				friends = friends[:j+copy(friends[j:], friends[j+1:])]
				deleted++
			} else {
				friendsAnimeListItems[friend] = friendAnimeListItem

				if len(episodeToFriends[friendAnimeListItem.Episodes]) < maxFriendsPerEpisode {
					episodeToFriends[friendAnimeListItem.Episodes] = append(episodeToFriends[friendAnimeListItem.Episodes], friend)
				}
			}
		}

		arn.SortUsersLastSeenFirst(friends)
	}

	// Sort relations by start date
	relations := anime.Relations()

	if relations != nil {
		relations.SortByStartDate()
	}

	// Soundtracks
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && arn.Contains(track.Tags, "anime:"+anime.ID)
	})

	sort.Slice(tracks, func(i, j int) bool {
		if len(tracks[i].Likes) == len(tracks[j].Likes) {
			return tracks[i].Title.ByUser(user) < tracks[j].Title.ByUser(user)
		}

		return len(tracks[i].Likes) > len(tracks[j].Likes)
	})

	// AMVs
	amvs := []*arn.AMV{}
	amvAppearances := []*arn.AMV{}

	for amv := range arn.StreamAMVs() {
		if amv.IsDraft {
			continue
		}

		if amv.MainAnimeID == anime.ID {
			amvs = append(amvs, amv)
		} else if arn.Contains(amv.ExtraAnimeIDs, anime.ID) {
			amvAppearances = append(amvAppearances, amv)
		}
	}

	sort.Slice(amvs, func(i, j int) bool {
		if len(amvs[i].Likes) == len(amvs[j].Likes) {
			return amvs[i].Title.ByUser(user) < amvs[j].Title.ByUser(user)
		}

		return len(amvs[i].Likes) > len(amvs[j].Likes)
	})

	// Open Graph
	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = getOpenGraph(anime)

	return ctx.HTML(components.Anime(anime, animeListItem, tracks, amvs, amvAppearances, episodes, friends, friendsAnimeListItems, episodeToFriends, user))
}

func getOpenGraph(anime *arn.Anime) *arn.OpenGraph {
	description := anime.Summary

	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength-3] + "..."
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       anime.Title.Canonical,
			"og:image":       "https:" + anime.ImageLink("large"),
			"og:url":         "https://" + assets.Domain + anime.Link(),
			"og:site_name":   "notify.moe",
			"og:description": description,
		},
		Meta: map[string]string{
			"description": description,
			"keywords":    anime.Title.Canonical + ",anime",
		},
	}

	switch anime.Type {
	case "tv":
		openGraph.Tags["og:type"] = "video.tv_show"
	case "movie":
		openGraph.Tags["og:type"] = "video.movie"
	}

	return openGraph
}
