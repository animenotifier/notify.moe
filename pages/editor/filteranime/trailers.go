package filteranime

import (
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Trailers ...
func Trailers(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without trailers",
		func(anime *arn.Anime) bool {
			return len(anime.Trailers) == 0
		},
		func(anime *arn.Anime) string {
			title := anime.Title.Canonical

			if anime.Title.Japanese != "" {
				title = anime.Title.Japanese
			}

			return "https://www.youtube.com/results?search_query=" + strings.Replace(title+" PV", " ", "+", -1)
		},
	)
}
