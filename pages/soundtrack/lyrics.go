package soundtrack

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/server/middleware"
	"github.com/animenotifier/notify.moe/utils"
)

// Lyrics of a soundtrack.
func Lyrics(ctx aero.Context) error {
	id := ctx.Get("id")
	track, err := arn.GetSoundTrack(id)
	user := arn.GetUserFromContext(ctx)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Track not found", err)
	}

	openGraph := getOpenGraph(track)

	if track.Lyrics.Native != "" {
		openGraph.Tags["og:description"] = utils.CutLongDescription(track.Lyrics.Native)
	}

	if track.Lyrics.Romaji != "" {
		openGraph.Tags["og:description"] = utils.CutLongDescription(track.Lyrics.Romaji)
	}

	customCtx := ctx.(*middleware.OpenGraphContext)
	customCtx.OpenGraph = openGraph
	return ctx.HTML(components.SoundTrackLyricsPage(track, user))
}
