package episode

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
	minio "github.com/minio/minio-go"
)

var spaces *minio.Client

func init() {
	if arn.APIKeys.S3.ID == "" || arn.APIKeys.S3.Secret == "" {
		return
	}

	go func() {
		var err error
		endpoint := "sfo2.digitaloceanspaces.com"
		ssl := true

		// Initiate a client using DigitalOcean Spaces.
		spaces, err = minio.New(endpoint, arn.APIKeys.S3.ID, arn.APIKeys.S3.Secret, ssl)

		if err != nil {
			log.Fatal(err)
		}
	}()
}

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

	// Does the episode exist?
	uploaded := false

	if spaces != nil {
		stat, err := spaces.StatObject("arn", fmt.Sprintf("videos/anime/%s/%d.webm", anime.ID, episodeNumber), minio.StatObjectOptions{})
		uploaded = (err == nil) && (stat.Size > 0)
	}

	episode, episodeIndex := animeEpisodes.Find(episodeNumber)

	if episode == nil {
		return ctx.Error(http.StatusNotFound, "Anime episode not found")
	}

	return ctx.HTML(components.AnimeEpisode(anime, episode, episodeIndex, uploaded, user))
}
