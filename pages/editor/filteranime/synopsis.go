package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Synopsis ...
func Synopsis(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without a long synopsis",
		func(anime *arn.Anime) bool {
			return len(anime.Summary) < 170
		},
		nil,
	)
}
