package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Shoboi ...
func Shoboi(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without Shoboi mappings",
		func(anime *arn.Anime) bool {
			return anime.GetMapping("shoboi/anime") == ""
		},
		func(anime *arn.Anime) string {
			return "http://cal.syoboi.jp/find?type=quick&sd=1&kw=" + anime.Title.Japanese
		},
	)
}
