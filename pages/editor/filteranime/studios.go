package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Studios ...
func Studios(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without studios",
		func(anime *arn.Anime) bool {
			return len(anime.StudioIDs) == 0
		},
		nil,
	)
}
