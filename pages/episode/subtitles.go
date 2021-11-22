package episode

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	minio "github.com/minio/minio-go/v7"
)

// Subtitles returns the subtitles.
func Subtitles(ctx aero.Context) error {
	id := ctx.Get("id")
	language := ctx.Get("language")

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

	ctx.Response().SetHeader("Access-Control-Allow-Origin", "*")
	ctx.Response().SetHeader("Content-Type", "text/vtt; charset=utf-8")

	obj, err := arn.Spaces.GetObject(ctx.Request().Context(), "arn", fmt.Sprintf("videos/anime/%s/%d.%s.vtt", anime.ID, episode.Number, language), minio.GetObjectOptions{})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, err)
	}

	defer obj.Close()
	return ctx.ReadAll(obj)
}
