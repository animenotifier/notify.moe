package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Licensors ...
func Licensors(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime without licensors",
		func(anime *arn.Anime) bool {
			return len(anime.LicensorIDs) == 0
		},
		nil,
	)
}
