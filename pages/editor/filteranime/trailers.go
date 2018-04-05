package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Trailers ...
func Trailers(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime without trailers",
		func(anime *arn.Anime) bool {
			return len(anime.Trailers) == 0
		},
		nil,
	)
}
