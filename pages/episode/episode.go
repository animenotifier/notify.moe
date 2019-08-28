package episode

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	minio "github.com/minio/minio-go/v6"
)

// Get renders the anime episode.
func Get(ctx aero.Context) error {
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

	// Does the episode exist?
	uploaded := false

	if arn.Spaces != nil {
		stat, err := arn.Spaces.StatObject("arn", fmt.Sprintf("videos/anime/%s/%d.webm", anime.ID, episodeNumber), minio.StatObjectOptions{})
		uploaded = (err == nil) && (stat.Size > 0)
	}

	episode, episodeIndex := animeEpisodes.Find(episodeNumber)

	if episode == nil {
		return ctx.Error(http.StatusNotFound, "Anime episode not found")
	}

	return ctx.HTML(components.AnimeEpisode(anime, episode, episodeIndex, uploaded, user))
}
