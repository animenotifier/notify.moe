package anime

import (
	"net/http"
	"sort"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEpisodes = 26
const maxEpisodesLongSeries = 10
const maxDescriptionLength = 170

// Get anime page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	episodes := anime.Episodes().Items
	// episodesReversed := false

	if len(episodes) > maxEpisodes {
		// episodesReversed = true
		episodes = episodes[len(episodes)-maxEpisodesLongSeries:]

		for i, j := 0, len(episodes)-1; i < j; i, j = i+1, j-1 {
			episodes[i], episodes[j] = episodes[j], episodes[i]
		}
	}

	// Friends watching
	var friends []*arn.User
	friendsAnimeListItems := map[*arn.User]*arn.AnimeListItem{}

	if user != nil {
		friends = user.Follows().Users()

		deleted := 0
		for i := range friends {
			j := i - deleted
			friendAnimeList := friends[j].AnimeList()
			friendAnimeListItem := friendAnimeList.Find(anime.ID)

			if friendAnimeListItem == nil {
				friends = friends[:j+copy(friends[j:], friends[j+1:])]
				deleted++
			} else {
				friendsAnimeListItems[friends[j]] = friendAnimeListItem
			}
		}

		arn.SortUsersLastSeen(friends)
	}

	// Sort relations by start date
	relations := anime.Relations()

	if relations != nil {
		items := relations.Items

		sort.Slice(items, func(i, j int) bool {
			return items[i].Anime().StartDate < items[j].Anime().StartDate
		})
	}

	// Soundtracks
	tracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Media) > 0 && arn.Contains(track.Tags, "anime:"+anime.ID)
	})

	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Title < tracks[j].Title
	})

	// Open Graph
	description := anime.Summary

	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength-3] + "..."
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       anime.Title.Canonical,
			"og:image":       anime.Image("large"),
			"og:url":         "https://" + ctx.App.Config.Domain + anime.Link(),
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

	ctx.Data = openGraph

	return ctx.HTML(components.Anime(anime, tracks, episodes, friends, friendsAnimeListItems, user))
}
