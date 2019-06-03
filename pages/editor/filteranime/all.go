package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// All ...
func All(ctx aero.Context) error {
	return editorList(
		ctx,
		"All anime",
		func(anime *arn.Anime) bool {
			return true
		},
		func(anime *arn.Anime) string {
			return "https://www.google.com/search?q=" + anime.Title.Canonical
		},
	)
}
