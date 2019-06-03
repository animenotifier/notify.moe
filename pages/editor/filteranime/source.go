package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Source ...
func Source(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without a source",
		func(anime *arn.Anime) bool {
			return anime.Source == ""
		},
		nil,
	)
}
