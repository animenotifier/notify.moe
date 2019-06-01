package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Genres ...
func Genres(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without genres",
		func(anime *arn.Anime) bool {
			return len(anime.Genres) == 0
		},
		nil,
	)
}
