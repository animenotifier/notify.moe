package episode

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get renders the anime episode.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	episodeNumber, err := ctx.GetInt("episode-number")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Episode is not a number", err)
	}

	// Get anime
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	// Get anime episodes
	animeEpisodes, err := arn.GetAnimeEpisodes(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime episodes not found", err)
	}

	episode, episodeIndex := animeEpisodes.Find(episodeNumber)

	if episode == nil {
		return ctx.Error(http.StatusNotFound, "Anime episode not found")
	}

	return ctx.HTML(components.AnimeEpisode(anime, episode, user, episodeIndex))
}
