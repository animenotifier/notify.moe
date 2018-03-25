package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// TBA ...
func TBA(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime to be announced",
		func(anime *arn.Anime) bool {
			return anime.Status == "tba"
		},
		func(anime *arn.Anime) string {
			return "https://www.google.com/search?q=" + anime.Title.Canonical
		},
	)
}
