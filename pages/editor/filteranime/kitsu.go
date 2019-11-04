package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Kitsu ...
func Kitsu(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without Kitsu mappings",
		func(anime *arn.Anime) bool {
			return anime.GetMapping("kitsu/anime") == ""
		},
		func(anime *arn.Anime) string {
			return "https://kitsu.io/anime?text=" + anime.Title.Canonical
		},
	)
}
