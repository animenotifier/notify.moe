package anime

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEpisodes = 26
const maxEpisodesLongSeries = 5
const maxDescriptionLength = 170

// Get anime page.
func Get(ctx *aero.Context) string {
	id := ctx.Get("id")
	user := utils.GetUser(ctx)
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	tracks, err := arn.GetSoundTracksByTag("anime:" + anime.ID)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Error fetching soundtracks", err)
	}

	episodesReversed := false

	if len(anime.Episodes().Items) > maxEpisodes {
		episodesReversed = true
		anime.Episodes().Items = anime.Episodes().Items[len(anime.Episodes().Items)-maxEpisodesLongSeries:]

		for i, j := 0, len(anime.Episodes().Items)-1; i < j; i, j = i+1, j-1 {
			anime.Episodes().Items[i], anime.Episodes().Items[j] = anime.Episodes().Items[j], anime.Episodes().Items[i]
		}
	}

	// Friends watching
	var friends []*arn.User

	if user != nil {
		friends = user.Follows().Users()

		deleted := 0
		for i := range friends {
			j := i - deleted
			if !friends[j].AnimeList().Contains(anime.ID) {
				friends = friends[:j+copy(friends[j:], friends[j+1:])]
				deleted++
			}
		}

		arn.SortUsersLastSeen(friends)
	}

	// Open Graph
	description := anime.Summary

	if len(description) > maxDescriptionLength {
		description = description[:maxDescriptionLength-3] + "..."
	}

	openGraph := &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":       anime.Title.Canonical,
			"og:image":       anime.Image.Large,
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

	return ctx.HTML(components.Anime(anime, friends, tracks, user, episodesReversed))
}
