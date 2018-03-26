package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
)

// Studios ...
func Studios(ctx *aero.Context) string {
	return editorList(
		ctx,
		"Anime without studios",
		func(anime *arn.Anime) bool {
			return len(anime.StudioIDs) == 0
		},
		nil,
	)
}
