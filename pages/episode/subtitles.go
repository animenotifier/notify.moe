package episode

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	minio "github.com/minio/minio-go/v6"
)

// Subtitles returns the subtitles.
func Subtitles(ctx aero.Context) error {
	id := ctx.Get("id")
	language := ctx.Get("language")
	episodeNumber, err := ctx.GetInt("episode-number")

	if err != nil {
		return ctx.Error(http.StatusBadRequest, "Episode is not a number", err)
	}

	// Get anime
	anime, err := arn.GetAnime(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Anime not found", err)
	}

	ctx.Response().SetHeader("Access-Control-Allow-Origin", "*")
	ctx.Response().SetHeader("Content-Type", "text/vtt; charset=utf-8")

	obj, err := arn.Spaces.GetObject("arn", fmt.Sprintf("videos/anime/%s/%d.%s.vtt", anime.ID, episodeNumber, language), minio.GetObjectOptions{})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, err)
	}

	defer obj.Close()
	return ctx.ReadAll(obj)
}
