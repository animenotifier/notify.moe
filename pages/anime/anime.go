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

	if len(anime.Episodes) > maxEpisodes {
		episodesReversed = true
		anime.Episodes = anime.Episodes[len(anime.Episodes)-maxEpisodesLongSeries-1:]

		for i, j := 0, len(anime.Episodes)-1; i < j; i, j = i+1, j-1 {
			anime.Episodes[i], anime.Episodes[j] = anime.Episodes[j], anime.Episodes[i]
		}
	}

	return ctx.HTML(components.Anime(anime, tracks, user, episodesReversed))
}
