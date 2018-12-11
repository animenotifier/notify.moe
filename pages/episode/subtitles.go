package episode

import (
	"fmt"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	minio "github.com/minio/minio-go"
)

// Subtitles returns the subtitles.
func Subtitles(ctx *aero.Context) string {
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

	ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Response().Header().Set("Content-Type", "text/vtt; charset=utf-8")

	obj, err := spaces.GetObject("arn", fmt.Sprintf("videos/anime/%s/%d.%s.vtt", anime.ID, episodeNumber, language), minio.GetObjectOptions{})

	if err != nil {
		return ctx.Error(http.StatusInternalServerError, err)
	}

	defer obj.Close()

	data := make([]byte, 0, 65535)
	buffer := make([]byte, 4096)
	n := 0

	for err == nil {
		n, err = obj.Read(buffer)
		data = append(data, buffer[:n]...)
	}

	return string(data)
}
