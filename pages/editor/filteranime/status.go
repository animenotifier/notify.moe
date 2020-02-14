package filteranime

import (
	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/arn"
)

// Status ...
func Status(ctx aero.Context) error {
	return editorList(
		ctx,
		"Anime with an incorrect status",
		func(anime *arn.Anime) bool {
			return anime.Status != anime.CalculatedStatus()
		},
		nil,
	)
}
