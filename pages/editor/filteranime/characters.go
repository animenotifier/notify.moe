package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Characters ...
func Characters(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime without characters",
		func(anime *arn.Anime) bool {
			return len(anime.Characters().Items) == 0
		},
		nil,
	)
}
