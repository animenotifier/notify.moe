package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Characters ...
func Characters(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without characters",
		func(anime *arn.Anime) bool {
			return len(anime.Characters().Items) == 0
		},
		nil,
	)
}
