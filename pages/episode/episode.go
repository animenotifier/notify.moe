package episode

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	minio "github.com/minio/minio-go/v7"
)

// Get renders the anime episode.
func Get(ctx aero.Context) error {
	user := arn.GetUserFromContext(ctx)
	id := ctx.Get("id")

	// Get episode
	episode, err := arn.GetEpisode(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Episode not found", err)
	}

	// Get anime
	anime := episode.Anime()

	if anime == nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	// Does the episode exist?
	uploaded := false

	if arn.Spaces != nil {
		stat, err := arn.Spaces.StatObject(ctx.Request().Context(), "arn", fmt.Sprintf("videos/anime/%s/%d.webm", anime.ID, episode.Number), minio.StatObjectOptions{})
		uploaded = (err == nil) && (stat.Size > 0)
	}

	_, episodeIndex := anime.Episodes().Find(episode.Number)

	if episode == nil {
		return ctx.Error(http.StatusNotFound, "Anime episode not found")
	}

	return ctx.HTML(components.Episode(anime, episode, episodeIndex, uploaded, user))
}
