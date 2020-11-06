package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// EpisodeCount ...
func EpisodeCount(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without an episode count",
		func(anime *arn.Anime) bool {
			return anime.EpisodeCount == 0
		},
		nil,
	)
}
