package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Relations ...
func Relations(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without relations",
		func(anime *arn.Anime) bool {
			return len(anime.Relations().Items) == 0
		},
		func(anime *arn.Anime) string {
			return "http://notify.moe/search/" + anime.Title.Canonical
		},
	)
}
