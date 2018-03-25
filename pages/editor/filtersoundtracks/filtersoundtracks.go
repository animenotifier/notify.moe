package filtersoundtracks

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

const maxEntries = 70

// NoLinks ...
func NoLinks(ctx *aero.Context) string {
	user := utils.GetUser(ctx)

	if user == nil || (user.Role != "admin" && user.Role != "editor") {
		return ctx.Redirect("/")
	}

	soundTracks := arn.FilterSoundTracks(func(track *arn.SoundTrack) bool {
		return !track.IsDraft && len(track.Links) == 0
	})

	arn.SortSoundTracksPopularFirst(soundTracks)

	count := len(soundTracks)

	if count > maxEntries {
		soundTracks = soundTracks[:maxEntries]
	}

	return ctx.HTML(components.SoundTracksEditorList(soundTracks, count, ctx.URI(), user))
}
